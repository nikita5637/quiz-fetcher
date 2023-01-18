package syncer

import (
	"context"
	"time"

	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/logger"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	time_utils "github.com/nikita5637/quiz-fetcher/utils/time"
)

const (
	syncerName = "games syncer"
)

// GamesSyncer ...
type GamesSyncer struct {
	enabled                  bool
	gamesFetchers            []fetcher.GamesFetcher
	lastSyncAt               time.Time
	lastSyncID               int
	name                     string
	period                   time.Duration
	registratorServiceClient registrator.RegistratorServiceClient
	syncStatus               model.SyncStatus
	syncerFacade             Facade
}

// GamesSyncerConfig ...
type GamesSyncerConfig struct {
	Enabled                  bool
	GamesFetchers            []fetcher.GamesFetcher
	Period                   uint64
	RegistratorServiceClient registrator.RegistratorServiceClient
	SyncerFacade             Facade
}

// NewGamesSyncer ...
func NewGamesSyncer(ctx context.Context, cfg GamesSyncerConfig) (*GamesSyncer, error) {
	logger.InfoKV(ctx, "initialize syncer", "name", syncerName)
	gs := &GamesSyncer{
		enabled:                  cfg.Enabled,
		gamesFetchers:            cfg.GamesFetchers,
		name:                     syncerName,
		period:                   time.Duration(cfg.Period),
		registratorServiceClient: cfg.RegistratorServiceClient,
		syncerFacade:             cfg.SyncerFacade,
		syncStatus:               model.SyncStatusNotSynced,
	}

	lastSync, err := gs.syncerFacade.GetLastSync(ctx, gs.name)
	if err != nil {
		return nil, err
	}

	// TODO change to error
	if lastSync.Name == "" {
		now := time_utils.TimeNow()
		lastSyncID, err := gs.syncerFacade.Insert(ctx, model.SyncLog{
			Name:       gs.name,
			LastSyncAt: now,
			Status:     gs.syncStatus,
		})
		if err != nil {
			return nil, err
		}

		gs.lastSyncAt = now
		gs.lastSyncID = lastSyncID
	} else {
		gs.syncStatus = lastSync.Status
		gs.lastSyncAt = lastSync.LastSyncAt
		gs.lastSyncID = lastSync.ID
	}

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
	begin := time_utils.TimeNow()

	if g.syncStatus == model.SyncStatusOK || g.syncStatus == model.SyncStatusFailed {
		now := time_utils.TimeNow()
		lastSyncID, err := g.syncerFacade.Insert(ctx, model.SyncLog{
			Name:       g.name,
			LastSyncAt: now,
			Status:     model.SyncStatusNotSynced,
		})
		if err != nil {
			return err
		}

		g.lastSyncAt = now
		g.lastSyncID = lastSyncID
		g.syncStatus = model.SyncStatusNotSynced
	}

	var err error
	err = g.syncerFacade.Update(ctx, model.SyncLog{
		ID:         g.lastSyncID,
		Name:       g.name,
		LastSyncAt: g.lastSyncAt,
		Status:     model.SyncStatusInProgress,
	})
	if err != nil {
		return err
	}

	g.syncStatus = model.SyncStatusInProgress

	err = func() error {
		for _, gamesFetcher := range g.gamesFetchers {
			var games []model.Game
			games, err = gamesFetcher.GetGamesList(ctx)
			if err != nil {
				logger.ErrorKV(ctx, "games fetch error", "fetcher name", gamesFetcher.GetName(), "error", err)
				return err
			}

			pbGames := make([]*registrator.AddGamesRequest_Game, 0, len(games))
			for _, game := range games {
				pbGame := &registrator.AddGamesRequest_Game{
					ExternalId:  game.ExternalID,
					LeagueId:    game.LeagueID,
					Type:        registrator.GameType(game.Type),
					Number:      game.Number,
					Name:        game.Name,
					PlaceId:     game.PlaceID,
					Date:        timestamppb.New(game.DateTime),
					Price:       game.Price,
					PaymentType: game.PaymentType,
					MaxPlayers:  game.MaxPlayers,
				}

				pbGames = append(pbGames, pbGame)
			}

			if _, err2 := g.registratorServiceClient.AddGames(ctx, &registrator.AddGamesRequest{
				Games: pbGames,
			}); err2 != nil {
				return err2
			}
		}

		return nil
	}()
	if err != nil {
		now := time_utils.TimeNow()
		updateErr := g.syncerFacade.Update(ctx, model.SyncLog{
			ID:         g.lastSyncID,
			Name:       g.name,
			LastSyncAt: now,
			Status:     model.SyncStatusFailed,
		})
		if updateErr != nil {
			return updateErr
		}

		g.lastSyncAt = now
		g.syncStatus = model.SyncStatusFailed

		return err
	}

	now := time_utils.TimeNow()
	err = g.syncerFacade.Update(ctx, model.SyncLog{
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

	logger.DebugKV(ctx, "sync done", "name", g.name, "duration(ms)", time_utils.TimeNow().Sub(begin).Milliseconds())
	return nil
}
