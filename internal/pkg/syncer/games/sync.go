package games

import (
	"context"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/logger"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	timeutils "github.com/nikita5637/quiz-fetcher/utils/time"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Sync ...
func (s *Syncer) Sync(ctx context.Context) error {
	begin := timeutils.TimeNow().UTC()

	logger.Debug(ctx, "syncing has started")

	for _, fetcher := range s.fetchers {
		fetcherName := fetcher.GetName()

		fetcherCtx := logger.ToContext(ctx, logger.FromContext(ctx).WithOptions(zap.Fields(
			zap.String("fetcher", fetcherName),
		)))

		gamesFetcherDisabled := false
		for _, disabledGamesFetcher := range s.disabledGamesFetchers {
			if disabledGamesFetcher == fetcherName {
				gamesFetcherDisabled = true
				break
			}
		}

		if gamesFetcherDisabled {
			logger.Info(fetcherCtx, "games fetcher disabled. skip...")
			continue
		}

		var masterGames []model.Game
		masterGames, err := fetcher.GetGamesList(fetcherCtx)
		if err != nil {
			logger.ErrorKV(fetcherCtx, "getting games list from a master error", zap.Error(err))
			continue
		}
		logger.Debugf(fetcherCtx, "there are %d games in a master", len(masterGames))

		pbLeagueGames := make([]*gamepb.Game, 0)
		page := uint64(1)
		for {
			var leagueGamesResp *gamepb.SearchGamesByLeagueIDResponse
			leagueGamesResp, err = s.gameServiceClient.SearchGamesByLeagueID(fetcherCtx, &gamepb.SearchGamesByLeagueIDRequest{
				Id:       fetcher.GetLeagueID(),
				Page:     page,
				PageSize: 50,
			})
			if err != nil || len(leagueGamesResp.GetGames()) == 0 {
				break
			}

			pbLeagueGames = append(pbLeagueGames, leagueGamesResp.GetGames()...)
			page++
		}
		if err != nil {
			logger.ErrorKV(fetcherCtx, "getting league games error", zap.Error(err))
			continue
		}

		leagueGames := make([]model.Game, 0, len(pbLeagueGames))
		for _, pbLeagueGame := range pbLeagueGames {
			leagueGames = append(leagueGames, convertProtoGameToModelGame(pbLeagueGame))
		}
		logger.Debugf(fetcherCtx, "there are %d games in a database", len(leagueGames))

		s.handleGames(fetcherCtx, masterGames, leagueGames)
	}

	logger.DebugKV(ctx, "syncing finished", zap.Int64("duration(ms)", timeutils.TimeNow().UTC().Sub(begin).Milliseconds()))
	return nil
}

func (s *Syncer) handleGames(ctx context.Context, masterGames, dbGames []model.Game) {
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

	logger.Debugf(ctx, "there are %d games to create", len(gamesToCreate))
	logger.Debugf(ctx, "there are %d games in both the master and the database", len(gamesInBoth))

	for _, gameToCreate := range gamesToCreate {
		logger.InfoKV(ctx, "creating new game", zap.Reflect("game", gameToCreate))
		_, err := s.gameServiceClient.CreateGame(ctx, &gamepb.CreateGameRequest{
			Game: convertModelGameToProtoGame(gameToCreate),
		})
		if err != nil {
			logger.ErrorKV(ctx, "create game error", zap.Error(err))
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

	logger.Debugf(ctx, "there are %d games to patch", len(gamesToPatch))

	for _, gameToPatch := range gamesToPatch {
		logger.InfoKV(ctx, "patching existing game", zap.Reflect("game", gameToPatch))
		_, err := s.gameServiceClient.PatchGame(ctx, &gamepb.PatchGameRequest{
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
			logger.ErrorKV(ctx, "patch game error", zap.Error(err))
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

	logger.Debugf(ctx, "there are %d games to delete", len(gamesToDelete))

	for _, gameToDelete := range gamesToDelete {
		logger.InfoKV(ctx, "deleting existing game", zap.Reflect("game", gameToDelete))
		_, err := s.gameServiceClient.PatchGame(ctx, &gamepb.PatchGameRequest{
			Game: convertModelGameToProtoGame(gameToDelete),
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"is_in_master",
				},
			},
		})
		if err != nil {
			logger.ErrorKV(ctx, "patch game error", zap.Error(err))
		}
	}
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
