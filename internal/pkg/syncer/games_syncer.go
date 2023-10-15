//go:generate mockery --case underscore --name GameServiceClient --with-expecter
//go:generate mockery --case underscore --name SyncLogFacade --with-expecter

package syncer

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	synclog "github.com/nikita5637/quiz-fetcher/internal/pkg/facade/sync_log"
	game_fetcher "github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/game"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/logger"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	time_utils "github.com/nikita5637/quiz-fetcher/utils/time"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	fetcherNameKey = "fetcher name"
	syncerName     = "games"
)

var (
	errJobAlreadyInProgress = errors.New("job already in progress")
)

// GameServiceClient ...
type GameServiceClient interface {
	CreateGame(ctx context.Context, in *gamepb.CreateGameRequest, opts ...grpc.CallOption) (*gamepb.Game, error)
	PatchGame(ctx context.Context, in *gamepb.PatchGameRequest, opts ...grpc.CallOption) (*gamepb.Game, error)
	SearchGamesByLeagueID(ctx context.Context, in *gamepb.SearchGamesByLeagueIDRequest, opts ...grpc.CallOption) (*gamepb.SearchGamesByLeagueIDResponse, error)
}

// SyncLogFacade ...
type SyncLogFacade interface {
	CreateSyncLog(ctx context.Context, syncLog model.SyncLog) (model.SyncLog, error)
	GetLastSync(ctx context.Context, name string) (model.SyncLog, error)
	PatchSyncLog(ctx context.Context, syncLog model.SyncLog) (model.SyncLog, error)
}

// GamesSyncer ...
type GamesSyncer struct {
	enabled               bool
	gameServiceClient     GameServiceClient
	gamesFetchers         []game_fetcher.Fetcher
	disabledGamesFetchers []string
	lastSyncAt            time.Time
	lastSyncID            int
	m                     sync.Mutex
	name                  string
	period                time.Duration
	syncLogFacade         SyncLogFacade
	syncStatus            model.SyncStatus
}

// GamesSyncerConfig ...
type GamesSyncerConfig struct {
	Enabled               bool
	GameServiceClient     GameServiceClient
	GamesFetchers         []game_fetcher.Fetcher
	DisabledGamesFetchers []string
	Period                time.Duration
	SyncLogFacade         SyncLogFacade
}

// NewGamesSyncer ...
func NewGamesSyncer(ctx context.Context, cfg GamesSyncerConfig) (*GamesSyncer, error) {
	logger.InfoKV(ctx, "initialize syncer", "name", syncerName)
	gs := &GamesSyncer{
		enabled:               cfg.Enabled,
		gameServiceClient:     cfg.GameServiceClient,
		gamesFetchers:         cfg.GamesFetchers,
		disabledGamesFetchers: cfg.DisabledGamesFetchers,
		m:                     sync.Mutex{},
		name:                  syncerName,
		period:                cfg.Period,
		syncStatus:            model.SyncStatusNotSynced,
		syncLogFacade:         cfg.SyncLogFacade,
	}

	lastSync, err := gs.syncLogFacade.GetLastSync(ctx, gs.name)
	if err != nil {
		if errors.Is(err, synclog.ErrSyncLogNotFound) {
			now := time_utils.TimeNow().UTC()
			var createdSyncLog model.SyncLog
			createdSyncLog, err = gs.syncLogFacade.CreateSyncLog(ctx, model.SyncLog{
				Name:       gs.name,
				LastSyncAt: now,
				Status:     gs.syncStatus,
			})
			if err != nil {
				return nil, fmt.Errorf("creating sync log error: %w", err)
			}

			gs.lastSyncAt = now
			gs.lastSyncID = createdSyncLog.ID
		} else {
			return nil, fmt.Errorf("getting last sync error: %w", err)
		}
	}

	gs.syncStatus = lastSync.Status
	gs.lastSyncAt = lastSync.LastSyncAt
	gs.lastSyncID = lastSync.ID

	logger.InfoKV(ctx, "initialize syncer done", "name", gs.name, "status", gs.syncStatus.String(), "sync period", cfg.Period, "enabled", cfg.Enabled)
	return gs, nil
}

// Enabled ...
func (g *GamesSyncer) Enabled() bool {
	return g.enabled
}

// GetLastSyncAt ...
func (g *GamesSyncer) GetLastSyncAt() time.Time {
	return g.lastSyncAt
}

// GetPeriod ...
func (g *GamesSyncer) GetPeriod() time.Duration {
	return g.period
}

// GetSyncerName ...
func (g *GamesSyncer) GetSyncerName() string {
	return g.name
}

// GetSyncStatus ...
func (g *GamesSyncer) GetSyncStatus() model.SyncStatus {
	return g.syncStatus
}

// Sync ...
func (g *GamesSyncer) Sync(ctx context.Context) error {
	begin := time_utils.TimeNow().UTC()

	err := g.prepare(ctx)
	if err != nil {
		if err == errJobAlreadyInProgress {
			logger.WarnKV(ctx, "games fetch job already in progress")
			return nil
		}
		return err
	}

	err = func() error {
		for _, gamesFetcher := range g.gamesFetchers {
			gamesFetcherDisabled := false
			for _, disabledGamesFetcher := range g.disabledGamesFetchers {
				if disabledGamesFetcher == gamesFetcher.GetName() {
					gamesFetcherDisabled = true
					break
				}
			}

			if gamesFetcherDisabled {
				logger.InfoKV(ctx, "games fetcher disabled. skip...", fetcherNameKey, gamesFetcher.GetName())
				continue
			}

			var masterGames []model.Game
			masterGames, err = gamesFetcher.GetGamesList(ctx)
			if err != nil {
				logger.ErrorKV(ctx, "games fetch error", fetcherNameKey, gamesFetcher.GetName(), "error", err)
				return err
			}
			logger.DebugKV(ctx, fmt.Sprintf("there are %d games in a master", len(masterGames)), fetcherNameKey, gamesFetcher.GetName())

			pbLeagueGames := make([]*gamepb.Game, 0)
			page := uint64(1)
			for {
				var leagueGamesResp *gamepb.SearchGamesByLeagueIDResponse
				leagueGamesResp, err = g.gameServiceClient.SearchGamesByLeagueID(ctx, &gamepb.SearchGamesByLeagueIDRequest{
					Id:       leaguepb.LeagueID(gamesFetcher.GetLeagueID()),
					Page:     page,
					PageSize: 50,
				})
				if err != nil {
					logger.ErrorKV(ctx, "get league games error", fetcherNameKey, gamesFetcher.GetName(), "error", err)
					return err
				}

				if len(leagueGamesResp.GetGames()) == 0 {
					break
				}

				pbLeagueGames = append(pbLeagueGames, leagueGamesResp.GetGames()...)
				page++
			}

			leagueGames := make([]model.Game, 0, len(pbLeagueGames))
			for _, pbLeagueGame := range pbLeagueGames {
				leagueGames = append(leagueGames, convertProtoGameToModelGame(pbLeagueGame))
			}
			logger.DebugKV(ctx, fmt.Sprintf("there are %d games in a database", len(leagueGames)), fetcherNameKey, gamesFetcher.GetName())

			if err = g.handleGames(ctx, gamesFetcher.GetName(), masterGames, leagueGames); err != nil {
				logger.ErrorKV(ctx, "handle games error", fetcherNameKey, gamesFetcher.GetName(), "error", err)
				return err
			}
		}

		return nil
	}()
	if err != nil {
		now := time_utils.TimeNow().UTC()
		if _, err = g.syncLogFacade.PatchSyncLog(ctx, model.SyncLog{
			ID:         g.lastSyncID,
			Name:       g.name,
			LastSyncAt: now,
			Status:     model.SyncStatusFailed,
		}); err != nil {
			return err
		}

		g.lastSyncAt = now
		g.syncStatus = model.SyncStatusFailed

		return err
	}

	now := time_utils.TimeNow().UTC()
	_, err = g.syncLogFacade.PatchSyncLog(ctx, model.SyncLog{
		ID:         g.lastSyncID,
		Name:       g.name,
		LastSyncAt: now,
		Status:     model.SyncStatusOK,
	})
	if err != nil {
		return err
	}
	g.lastSyncAt = now
	g.syncStatus = model.SyncStatusOK

	logger.DebugKV(ctx, "sync done", "name", g.name, "duration(ms)", time_utils.TimeNow().UTC().Sub(begin).Milliseconds())
	return nil
}

func (g *GamesSyncer) handleGames(ctx context.Context, gamesFetcherName string, masterGames, dbGames []model.Game) error {
	var gamesToCreate []model.Game
	var gamesInBoth []model.Game
	for _, masterGame := range masterGames {
		found := false
		for _, dbGame := range dbGames {
			if isEqual(masterGame, dbGame) {
				found = true
				break
			}
		}
		if !found {
			gamesToCreate = append(gamesToCreate, masterGame)
		} else {
			gamesInBoth = append(gamesInBoth, masterGame)
		}
	}

	logger.DebugKV(ctx, fmt.Sprintf("there are %d games to create", len(gamesToCreate)), fetcherNameKey, gamesFetcherName)
	logger.DebugKV(ctx, fmt.Sprintf("there are %d games in both the master and the database", len(gamesInBoth)), fetcherNameKey, gamesFetcherName)

	for _, gameToCreate := range gamesToCreate {
		logger.InfoKV(ctx, "creating new game", "game", gameToCreate)
		_, err := g.gameServiceClient.CreateGame(ctx, &gamepb.CreateGameRequest{
			Game: convertModelGameToProtoGame(gameToCreate),
		})
		if err != nil {
			logger.ErrorKV(ctx, "create game error", fetcherNameKey, gamesFetcherName, "error", err)
		}
	}

	var gamesToPatch []model.Game
	for _, gameInBoth := range gamesInBoth {
		for _, dbGame := range dbGames {
			if isEqual(gameInBoth, dbGame) {
				if !isDeepEqual(gameInBoth, dbGame) {
					g := gameInBoth
					g.ID = dbGame.ID
					gamesToPatch = append(gamesToPatch, g)
				}
			}
		}
	}

	logger.DebugKV(ctx, fmt.Sprintf("there are %d games to patch", len(gamesToPatch)), fetcherNameKey, gamesFetcherName)

	for _, gameToPatch := range gamesToPatch {
		logger.InfoKV(ctx, "patching existing game", "game", gameToPatch)
		_, err := g.gameServiceClient.PatchGame(ctx, &gamepb.PatchGameRequest{
			Game: convertModelGameToProtoGame(gameToPatch),
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"type",
					"name",
					"price",
					"payment_type",
					"max_players",
					"is_in_master",
				},
			},
		})
		if err != nil {
			logger.ErrorKV(ctx, "patch game error", fetcherNameKey, gamesFetcherName, "error", err)
		}
	}

	var gamesToDelete []model.Game
	for _, dbGame := range dbGames {
		found := false
		for _, masterGame := range masterGames {
			if isEqual(dbGame, masterGame) {
				found = true
				break
			}
		}
		if !found && dbGame.IsInMaster {
			g := dbGame
			g.IsInMaster = false
			gamesToDelete = append(gamesToDelete, g)
		}
	}

	logger.DebugKV(ctx, fmt.Sprintf("there are %d games to delete", len(gamesToDelete)), fetcherNameKey, gamesFetcherName)

	for _, gameToDelete := range gamesToDelete {
		logger.InfoKV(ctx, "deleting existing game", "game", gameToDelete)
		_, err := g.gameServiceClient.PatchGame(ctx, &gamepb.PatchGameRequest{
			Game: convertModelGameToProtoGame(gameToDelete),
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"is_in_master",
				},
			},
		})
		if err != nil {
			logger.ErrorKV(ctx, "patch game error", fetcherNameKey, gamesFetcherName, "error", err)
		}
	}
	return nil
}

func isEqual(game1, game2 model.Game) bool {
	return game1.ExternalID == game2.ExternalID &&
		game1.LeagueID == game2.LeagueID &&
		game1.Number == game2.Number &&
		game1.PlaceID == game2.PlaceID &&
		game1.DateTime == game2.DateTime
}

func isDeepEqual(game1, game2 model.Game) bool {
	return game1.ExternalID == game2.ExternalID &&
		game1.LeagueID == game2.LeagueID &&
		game1.Type == game2.Type &&
		game1.Number == game2.Number &&
		game1.Name == game2.Name &&
		game1.PlaceID == game2.PlaceID &&
		game1.DateTime == game2.DateTime &&
		game1.Price == game2.Price &&
		game1.PaymentType == game2.PaymentType &&
		game1.MaxPlayers == game2.MaxPlayers &&
		game1.IsInMaster == game2.IsInMaster
}

func (g *GamesSyncer) prepare(ctx context.Context) error {
	g.m.Lock()
	defer g.m.Unlock()

	if g.syncStatus == model.SyncStatusOK || g.syncStatus == model.SyncStatusFailed {
		now := time_utils.TimeNow().UTC()
		createdSyncLog, err := g.syncLogFacade.CreateSyncLog(ctx, model.SyncLog{
			Name:       g.name,
			LastSyncAt: now,
			Status:     model.SyncStatusNotSynced,
		})
		if err != nil {
			return err
		}

		g.lastSyncAt = now
		g.lastSyncID = createdSyncLog.ID
		g.syncStatus = model.SyncStatusNotSynced
	} else if g.syncStatus == model.SyncStatusInProgress {
		// skip if status is "In progress"
		return errJobAlreadyInProgress
	}

	if _, err := g.syncLogFacade.PatchSyncLog(ctx, model.SyncLog{
		ID:         g.lastSyncID,
		Name:       g.name,
		LastSyncAt: g.lastSyncAt,
		Status:     model.SyncStatusInProgress,
	}); err != nil {
		return err
	}

	g.syncStatus = model.SyncStatusInProgress

	return nil
}
