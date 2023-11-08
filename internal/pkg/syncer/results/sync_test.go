package results

import (
	"errors"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	gameresultmanager "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestSyncer_Sync(t *testing.T) {
	t.Run("error: list game results error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultManagerServiceClient.EXPECT().ListGameResults(fx.ctx, &emptypb.Empty{}).Return(nil, errors.New("some error"))

		err := fx.syncer.Sync(fx.ctx)
		assert.Error(t, err)
	})

	t.Run("ok: there are no game results without round points", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultManagerServiceClient.EXPECT().ListGameResults(fx.ctx, &emptypb.Empty{}).Return(&gameresultmanager.ListGameResultsResponse{
			GameResults: []*gameresultmanager.GameResult{
				{
					Id:          1,
					RoundPoints: "points",
				},
				{
					Id:          2,
					RoundPoints: "points",
				},
			},
		}, nil)

		err := fx.syncer.Sync(fx.ctx)
		assert.NoError(t, err)
	})

	t.Run("error: get game error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultManagerServiceClient.EXPECT().ListGameResults(fx.ctx, &emptypb.Empty{}).Return(&gameresultmanager.ListGameResultsResponse{
			GameResults: []*gameresultmanager.GameResult{
				{
					Id:     1,
					GameId: 1,
				},
				{
					Id:     2,
					GameId: 2,
				},
			},
		}, nil)

		fx.gameServiceClient.EXPECT().GetGame(mock.Anything, &gamepb.GetGameRequest{
			Id: 1,
		}).Return(nil, errors.New("some error"))

		fx.gameServiceClient.EXPECT().GetGame(mock.Anything, &gamepb.GetGameRequest{
			Id: 2,
		}).Return(nil, errors.New("some error"))

		err := fx.syncer.Sync(fx.ctx)

		assert.NoError(t, err)
	})

	t.Run("ok: fetcher for league not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultManagerServiceClient.EXPECT().ListGameResults(fx.ctx, &emptypb.Empty{}).Return(&gameresultmanager.ListGameResultsResponse{
			GameResults: []*gameresultmanager.GameResult{
				{
					GameId: 1,
				},
			},
		}, nil)

		fx.gameServiceClient.EXPECT().GetGame(mock.Anything, &gamepb.GetGameRequest{
			Id: 1,
		}).Return(&gamepb.Game{
			Id:       1,
			LeagueId: 1,
		}, nil)

		fx.fetcher1.EXPECT().GetLeagueID().Return(2)
		fx.fetcher2.EXPECT().GetLeagueID().Return(3)

		err := fx.syncer.Sync(fx.ctx)
		assert.NoError(t, err)
	})

	t.Run("ok: game without external ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultManagerServiceClient.EXPECT().ListGameResults(fx.ctx, &emptypb.Empty{}).Return(&gameresultmanager.ListGameResultsResponse{
			GameResults: []*gameresultmanager.GameResult{
				{
					GameId: 1,
				},
			},
		}, nil)

		fx.gameServiceClient.EXPECT().GetGame(mock.Anything, &gamepb.GetGameRequest{
			Id: 1,
		}).Return(&gamepb.Game{
			Id:       1,
			LeagueId: 1,
		}, nil)

		fx.fetcher1.EXPECT().GetLeagueID().Return(1)
		fx.fetcher1.EXPECT().GetName().Return("league")

		err := fx.syncer.Sync(fx.ctx)
		assert.NoError(t, err)
	})

	t.Run("error: get game result from master error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultManagerServiceClient.EXPECT().ListGameResults(fx.ctx, &emptypb.Empty{}).Return(&gameresultmanager.ListGameResultsResponse{
			GameResults: []*gameresultmanager.GameResult{
				{
					GameId: 1,
				},
			},
		}, nil)

		fx.gameServiceClient.EXPECT().GetGame(mock.Anything, &gamepb.GetGameRequest{
			Id: 1,
		}).Return(&gamepb.Game{
			Id:         1,
			ExternalId: wrapperspb.Int32(1),
			LeagueId:   1,
		}, nil)

		fx.fetcher1.EXPECT().GetLeagueID().Return(1)
		fx.fetcher1.EXPECT().GetName().Return("league")
		fx.fetcher1.EXPECT().GetGameResult(mock.Anything, int32(1)).Return(model.GameResult{}, errors.New("some error"))

		err := fx.syncer.Sync(fx.ctx)
		assert.NoError(t, err)
	})

	t.Run("error: patch game result error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultManagerServiceClient.EXPECT().ListGameResults(fx.ctx, &emptypb.Empty{}).Return(&gameresultmanager.ListGameResultsResponse{
			GameResults: []*gameresultmanager.GameResult{
				{
					Id:     777,
					GameId: 1,
				},
			},
		}, nil)

		fx.gameServiceClient.EXPECT().GetGame(mock.Anything, &gamepb.GetGameRequest{
			Id: 1,
		}).Return(&gamepb.Game{
			Id:         1,
			ExternalId: wrapperspb.Int32(1),
			LeagueId:   1,
		}, nil)

		fx.fetcher1.EXPECT().GetLeagueID().Return(1)
		fx.fetcher1.EXPECT().GetName().Return("league")
		fx.fetcher1.EXPECT().GetGameResult(mock.Anything, int32(1)).Return(model.GameResult{
			ResultPlace: 1,
			RoundPoints: maybe.Just("{}"),
		}, nil)

		fx.gameResultManagerServiceClient.EXPECT().PatchGameResult(mock.Anything, &gameresultmanager.PatchGameResultRequest{
			GameResult: &gameresultmanager.GameResult{
				Id:          777,
				ResultPlace: 1,
				RoundPoints: "{}",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"result_place",
					"round_points",
				},
			},
		}).Return(nil, errors.New("some error"))

		err := fx.syncer.Sync(fx.ctx)
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameResultManagerServiceClient.EXPECT().ListGameResults(fx.ctx, &emptypb.Empty{}).Return(&gameresultmanager.ListGameResultsResponse{
			GameResults: []*gameresultmanager.GameResult{
				{
					Id:     777,
					GameId: 1,
				},
			},
		}, nil)

		fx.gameServiceClient.EXPECT().GetGame(mock.Anything, &gamepb.GetGameRequest{
			Id: 1,
		}).Return(&gamepb.Game{
			Id:         1,
			ExternalId: wrapperspb.Int32(1),
			LeagueId:   1,
		}, nil)

		fx.fetcher1.EXPECT().GetLeagueID().Return(1)
		fx.fetcher1.EXPECT().GetName().Return("league")
		fx.fetcher1.EXPECT().GetGameResult(mock.Anything, int32(1)).Return(model.GameResult{
			ResultPlace: 1,
			RoundPoints: maybe.Just("{}"),
		}, nil)

		fx.gameResultManagerServiceClient.EXPECT().PatchGameResult(mock.Anything, &gameresultmanager.PatchGameResultRequest{
			GameResult: &gameresultmanager.GameResult{
				Id:          777,
				ResultPlace: 1,
				RoundPoints: "{}",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"result_place",
					"round_points",
				},
			},
		}).Return(&gameresultmanager.GameResult{
			Id:          777,
			ResultPlace: 1,
			RoundPoints: "{}",
		}, nil)

		err := fx.syncer.Sync(fx.ctx)
		assert.NoError(t, err)
	})
}
