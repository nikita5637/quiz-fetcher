package results

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/logger"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	gameresultmanager "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Sync ...
func (s *Syncer) Sync(ctx context.Context) error {
	resp, err := s.gameResultManagerServiceClient.ListGameResults(ctx, &emptypb.Empty{})
	if err != nil {
		return fmt.Errorf("listing game results error: %w", err)
	}

	// patching
	for _, gameResult := range resp.GetGameResults() {
		if gameResult.GetRoundPoints() != "" {
			continue
		}

		gameResp, err := s.gameServiceClient.GetGame(ctx, &gamepb.GetGameRequest{
			Id: gameResult.GetGameId(),
		})
		if err != nil {
			logger.ErrorKV(ctx, "getting game error", zap.Error(err), zap.Int32("gameID", gameResult.GameId))
			continue
		}

		for _, fetcher := range s.fetchers {
			if gameResp.GetLeagueId() != fetcher.GetLeagueID() {
				continue
			}

			fetchCtx := logger.ToContext(ctx, logger.FromContext(ctx).WithOptions(zap.Fields(
				zap.String("fetcher", fetcher.GetName()),
			)))

			if externalID := gameResp.GetExternalId(); externalID != nil {
				fetchedGameResult, err := fetcher.GetGameResult(fetchCtx, externalID.GetValue())
				if err != nil {
					logger.ErrorKV(fetchCtx, "getting game result from master error", zap.Error(err), zap.Int32("externalID", externalID.GetValue()))
					break
				}

				_, err = s.gameResultManagerServiceClient.PatchGameResult(fetchCtx, &gameresultmanager.PatchGameResultRequest{
					GameResult: &gameresultmanager.GameResult{
						Id:          gameResult.GetId(),
						ResultPlace: fetchedGameResult.ResultPlace,
						RoundPoints: fetchedGameResult.RoundPoints.Value(),
					},
					UpdateMask: &fieldmaskpb.FieldMask{
						Paths: []string{
							"result_place",
							"round_points",
						},
					},
				})
				if err != nil {
					logger.ErrorKV(fetchCtx, "patching game result error", zap.Error(err), zap.Int32("gameResultID", gameResult.GetId()))
					break
				}

				logger.InfoKV(fetchCtx, "game result has been patched", zap.Int32("gameResultID", gameResult.GetId()))
			}

			break
		}
	}

	// creating
	// todo ...

	return nil
}
