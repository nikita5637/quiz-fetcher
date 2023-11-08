package games

import (
	"errors"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	timeutils "github.com/nikita5637/quiz-fetcher/utils/time"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestSyncer_Sync(t *testing.T) {
	t.Run("ok: all fetchers is disabled", func(t *testing.T) {
		fx := tearUp(t)

		fx.syncer.disabledGamesFetchers = []string{
			"fetcher1",
			"fetcher2",
		}

		fx.gamesFetcher1.EXPECT().GetName().Return("fetcher1")
		fx.gamesFetcher2.EXPECT().GetName().Return("fetcher2")

		err := fx.syncer.Sync(fx.ctx)
		assert.NoError(t, err)
	})

	t.Run("error: getting games list from a master", func(t *testing.T) {
		fx := tearUp(t)

		fx.syncer.disabledGamesFetchers = []string{
			"fetcher2",
		}

		fx.gamesFetcher1.EXPECT().GetName().Return("fetcher1")
		fx.gamesFetcher2.EXPECT().GetName().Return("fetcher2")

		fx.gamesFetcher1.EXPECT().GetGamesList(mock.Anything).Return(nil, errors.New("some error"))

		err := fx.syncer.Sync(fx.ctx)
		assert.NoError(t, err)
	})

	t.Run("error: getting games by league ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.syncer.disabledGamesFetchers = []string{
			"fetcher2",
		}

		fx.gamesFetcher1.EXPECT().GetName().Return("fetcher1")
		fx.gamesFetcher2.EXPECT().GetName().Return("fetcher2")

		fx.gamesFetcher1.EXPECT().GetGamesList(mock.Anything).Return([]model.Game{
			{
				ExternalID:  maybe.Just(int32(222)),
				LeagueID:    1,
				Name:        maybe.Just("new name"),
				DateTime:    timeutils.ConvertTime("2023-01-02 16:30"),
				PaymentType: maybe.Just("cash"),
			},
			{
				ExternalID:  maybe.Just(int32(333)),
				LeagueID:    1,
				Name:        maybe.Just("new name"),
				DateTime:    timeutils.ConvertTime("2023-01-03 16:30"),
				PaymentType: maybe.Just("cash"),
			},
			{
				ExternalID:  maybe.Just(int32(444)),
				LeagueID:    1,
				Name:        maybe.Nothing[string](),
				DateTime:    timeutils.ConvertTime("2023-01-04 16:30"),
				PaymentType: maybe.Just("cash"),
			},
			{
				ExternalID:  maybe.Just(int32(555)),
				LeagueID:    1,
				Name:        maybe.Nothing[string](),
				DateTime:    timeutils.ConvertTime("2023-01-05 16:30"),
				PaymentType: maybe.Just("cash"),
			},
		}, nil)

		fx.gamesFetcher1.EXPECT().GetLeagueID().Return(1)

		fx.gameServiceClient.EXPECT().SearchGamesByLeagueID(mock.Anything, &gamepb.SearchGamesByLeagueIDRequest{
			Id:       1,
			Page:     1,
			PageSize: 50,
		}).Return(nil, errors.New("some error"))

		err := fx.syncer.Sync(fx.ctx)
		assert.NoError(t, err)
	})

	t.Run("error: creating game error", func(t *testing.T) {
		fx := tearUp(t)

		fx.syncer.disabledGamesFetchers = []string{
			"fetcher2",
		}

		fx.gamesFetcher1.EXPECT().GetName().Return("fetcher1")
		fx.gamesFetcher2.EXPECT().GetName().Return("fetcher2")

		fx.gamesFetcher1.EXPECT().GetGamesList(mock.Anything).Return([]model.Game{
			{
				ExternalID:  maybe.Just(int32(444)),
				LeagueID:    1,
				Name:        maybe.Nothing[string](),
				DateTime:    timeutils.ConvertTime("2023-01-04 16:30"),
				PaymentType: maybe.Just("cash"),
			},
		}, nil)

		fx.gamesFetcher1.EXPECT().GetLeagueID().Return(1)

		fx.gameServiceClient.EXPECT().SearchGamesByLeagueID(mock.Anything, &gamepb.SearchGamesByLeagueIDRequest{
			Id:       1,
			Page:     1,
			PageSize: 50,
		}).Return(&gamepb.SearchGamesByLeagueIDResponse{
			Games: []*gamepb.Game{},
			Total: 0,
		}, nil)

		fx.gameServiceClient.EXPECT().CreateGame(mock.Anything, &gamepb.CreateGameRequest{
			Game: &gamepb.Game{
				ExternalId:  wrapperspb.Int32(444),
				LeagueId:    1,
				Date:        timestamppb.New(timeutils.ConvertTime("2023-01-04 16:30")),
				PaymentType: wrapperspb.String("cash"),
			},
		}).Return(nil, errors.New("some error"))

		err := fx.syncer.Sync(fx.ctx)
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.syncer.disabledGamesFetchers = []string{
			"fetcher2",
		}

		fx.gamesFetcher1.EXPECT().GetName().Return("fetcher1")
		fx.gamesFetcher2.EXPECT().GetName().Return("fetcher2")

		fx.gamesFetcher1.EXPECT().GetGamesList(mock.Anything).Return([]model.Game{
			{
				ExternalID:  maybe.Just(int32(222)),
				LeagueID:    1,
				Name:        maybe.Just("new name"),
				DateTime:    timeutils.ConvertTime("2023-01-02 16:30"),
				PaymentType: maybe.Just("cash"),
			},
			{
				ExternalID:  maybe.Just(int32(333)),
				LeagueID:    1,
				Name:        maybe.Just("new name"),
				DateTime:    timeutils.ConvertTime("2023-01-03 16:30"),
				PaymentType: maybe.Just("cash"),
			},
			{
				ExternalID:  maybe.Just(int32(444)),
				LeagueID:    1,
				Name:        maybe.Nothing[string](),
				DateTime:    timeutils.ConvertTime("2023-01-04 16:30"),
				PaymentType: maybe.Just("cash"),
			},
			{
				ExternalID:  maybe.Just(int32(555)),
				LeagueID:    1,
				Name:        maybe.Nothing[string](),
				DateTime:    timeutils.ConvertTime("2023-01-05 16:30"),
				PaymentType: maybe.Just("cash"),
			},
		}, nil)

		fx.gamesFetcher1.EXPECT().GetLeagueID().Return(1)

		fx.gameServiceClient.EXPECT().SearchGamesByLeagueID(mock.Anything, &gamepb.SearchGamesByLeagueIDRequest{
			Id:       1,
			Page:     1,
			PageSize: 50,
		}).Return(&gamepb.SearchGamesByLeagueIDResponse{
			Games: []*gamepb.Game{
				{
					ExternalId:  wrapperspb.Int32(111),
					LeagueId:    1,
					Date:        timestamppb.New(timeutils.ConvertTime("2023-01-01 16:30")),
					PaymentType: wrapperspb.String("cash"),
					IsInMaster:  true,
				},
				{
					ExternalId:  wrapperspb.Int32(222),
					LeagueId:    1,
					Name:        wrapperspb.String("old name"),
					Date:        timestamppb.New(timeutils.ConvertTime("2023-01-02 16:30")),
					PaymentType: wrapperspb.String("cash"),
				},
				{
					ExternalId:  wrapperspb.Int32(333),
					LeagueId:    1,
					Name:        wrapperspb.String("old name"),
					Date:        timestamppb.New(timeutils.ConvertTime("2023-01-03 16:30")),
					PaymentType: wrapperspb.String("cash"),
				},
				{
					ExternalId:  wrapperspb.Int32(666),
					LeagueId:    1,
					Date:        timestamppb.New(timeutils.ConvertTime("2023-01-06 16:30")),
					PaymentType: wrapperspb.String("cash"),
					IsInMaster:  true,
				},
			},
			Total: 4,
		}, nil)

		fx.gameServiceClient.EXPECT().SearchGamesByLeagueID(mock.Anything, &gamepb.SearchGamesByLeagueIDRequest{
			Id:       1,
			Page:     2,
			PageSize: 50,
		}).Return(&gamepb.SearchGamesByLeagueIDResponse{
			Games: []*gamepb.Game{},
			Total: 4,
		}, nil)

		fx.gameServiceClient.EXPECT().CreateGame(mock.Anything, &gamepb.CreateGameRequest{
			Game: &gamepb.Game{
				ExternalId:  wrapperspb.Int32(444),
				LeagueId:    1,
				Date:        timestamppb.New(timeutils.ConvertTime("2023-01-04 16:30")),
				PaymentType: wrapperspb.String("cash"),
			},
		}).Return(&gamepb.Game{
			ExternalId:  wrapperspb.Int32(444),
			LeagueId:    1,
			Date:        timestamppb.New(timeutils.ConvertTime("2023-01-04 16:30")),
			PaymentType: wrapperspb.String("cash"),
		}, nil)

		fx.gameServiceClient.EXPECT().CreateGame(mock.Anything, &gamepb.CreateGameRequest{
			Game: &gamepb.Game{
				ExternalId:  wrapperspb.Int32(555),
				LeagueId:    1,
				Date:        timestamppb.New(timeutils.ConvertTime("2023-01-05 16:30")),
				PaymentType: wrapperspb.String("cash"),
			},
		}).Return(&gamepb.Game{
			ExternalId:  wrapperspb.Int32(555),
			LeagueId:    1,
			Date:        timestamppb.New(timeutils.ConvertTime("2023-01-05 16:30")),
			PaymentType: wrapperspb.String("cash"),
		}, nil)

		fx.gameServiceClient.EXPECT().PatchGame(mock.Anything, &gamepb.PatchGameRequest{
			Game: &gamepb.Game{
				ExternalId:  wrapperspb.Int32(222),
				LeagueId:    1,
				Name:        wrapperspb.String("new name"),
				Date:        timestamppb.New(timeutils.ConvertTime("2023-01-02 16:30")),
				PaymentType: wrapperspb.String("cash"),
			},
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
		}).Return(&gamepb.Game{
			ExternalId:  wrapperspb.Int32(222),
			LeagueId:    1,
			Name:        wrapperspb.String("new name"),
			Date:        timestamppb.New(timeutils.ConvertTime("2023-01-02 16:30")),
			PaymentType: wrapperspb.String("cash"),
		}, nil)

		fx.gameServiceClient.EXPECT().PatchGame(mock.Anything, &gamepb.PatchGameRequest{
			Game: &gamepb.Game{
				ExternalId:  wrapperspb.Int32(333),
				LeagueId:    1,
				Name:        wrapperspb.String("new name"),
				Date:        timestamppb.New(timeutils.ConvertTime("2023-01-03 16:30")),
				PaymentType: wrapperspb.String("cash"),
			},
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
		}).Return(&gamepb.Game{
			ExternalId:  wrapperspb.Int32(333),
			LeagueId:    1,
			Name:        wrapperspb.String("new name"),
			Date:        timestamppb.New(timeutils.ConvertTime("2023-01-03 16:30")),
			PaymentType: wrapperspb.String("cash"),
		}, nil)

		fx.gameServiceClient.EXPECT().PatchGame(mock.Anything, &gamepb.PatchGameRequest{
			Game: &gamepb.Game{
				ExternalId:  wrapperspb.Int32(111),
				LeagueId:    1,
				Date:        timestamppb.New(timeutils.ConvertTime("2023-01-01 16:30")),
				PaymentType: wrapperspb.String("cash"),
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"is_in_master",
				},
			},
		}).Return(&gamepb.Game{
			ExternalId:  wrapperspb.Int32(111),
			LeagueId:    1,
			Date:        timestamppb.New(timeutils.ConvertTime("2023-01-01 16:30")),
			PaymentType: wrapperspb.String("cash"),
		}, nil)

		fx.gameServiceClient.EXPECT().PatchGame(mock.Anything, &gamepb.PatchGameRequest{
			Game: &gamepb.Game{
				ExternalId:  wrapperspb.Int32(666),
				LeagueId:    1,
				Date:        timestamppb.New(timeutils.ConvertTime("2023-01-06 16:30")),
				PaymentType: wrapperspb.String("cash"),
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"is_in_master",
				},
			},
		}).Return(&gamepb.Game{
			ExternalId:  wrapperspb.Int32(666),
			LeagueId:    1,
			Date:        timestamppb.New(timeutils.ConvertTime("2023-01-06 16:30")),
			PaymentType: wrapperspb.String("cash"),
		}, nil)

		err := fx.syncer.Sync(fx.ctx)
		assert.NoError(t, err)
	})
}

func TestSyncer_handleGames(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameServiceClient.EXPECT().CreateGame(fx.ctx, &gamepb.CreateGameRequest{
			Game: &gamepb.Game{
				ExternalId:  wrapperspb.Int32(444),
				Date:        timestamppb.New(timeutils.ConvertTime("2023-01-04 16:30")),
				PaymentType: wrapperspb.String("cash"),
			},
		}).Return(&gamepb.Game{
			ExternalId:  wrapperspb.Int32(444),
			Date:        timestamppb.New(timeutils.ConvertTime("2023-01-04 16:30")),
			PaymentType: wrapperspb.String("cash"),
		}, nil)

		fx.gameServiceClient.EXPECT().CreateGame(fx.ctx, &gamepb.CreateGameRequest{
			Game: &gamepb.Game{
				ExternalId:  wrapperspb.Int32(555),
				Date:        timestamppb.New(timeutils.ConvertTime("2023-01-05 16:30")),
				PaymentType: wrapperspb.String("cash"),
			},
		}).Return(nil, errors.New("some error"))

		fx.gameServiceClient.EXPECT().PatchGame(fx.ctx, &gamepb.PatchGameRequest{
			Game: &gamepb.Game{
				ExternalId:  wrapperspb.Int32(222),
				Name:        wrapperspb.String("new name"),
				Date:        timestamppb.New(timeutils.ConvertTime("2023-01-02 16:30")),
				PaymentType: wrapperspb.String("cash"),
			},
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
		}).Return(&gamepb.Game{
			ExternalId:  wrapperspb.Int32(222),
			Name:        wrapperspb.String("new name"),
			Date:        timestamppb.New(timeutils.ConvertTime("2023-01-02 16:30")),
			PaymentType: wrapperspb.String("cash"),
		}, nil)

		fx.gameServiceClient.EXPECT().PatchGame(fx.ctx, &gamepb.PatchGameRequest{
			Game: &gamepb.Game{
				ExternalId:  wrapperspb.Int32(333),
				Name:        wrapperspb.String("new name"),
				Date:        timestamppb.New(timeutils.ConvertTime("2023-01-03 16:30")),
				PaymentType: wrapperspb.String("cash"),
			},
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
		}).Return(nil, errors.New("some error"))

		fx.gameServiceClient.EXPECT().PatchGame(fx.ctx, &gamepb.PatchGameRequest{
			Game: &gamepb.Game{
				ExternalId:  wrapperspb.Int32(111),
				Date:        timestamppb.New(timeutils.ConvertTime("2023-01-01 16:30")),
				PaymentType: wrapperspb.String("cash"),
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"is_in_master",
				},
			},
		}).Return(&gamepb.Game{
			ExternalId:  wrapperspb.Int32(111),
			Date:        timestamppb.New(timeutils.ConvertTime("2023-01-01 16:30")),
			PaymentType: wrapperspb.String("cash"),
		}, nil)

		fx.gameServiceClient.EXPECT().PatchGame(fx.ctx, &gamepb.PatchGameRequest{
			Game: &gamepb.Game{
				ExternalId:  wrapperspb.Int32(666),
				Date:        timestamppb.New(timeutils.ConvertTime("2023-01-06 16:30")),
				PaymentType: wrapperspb.String("cash"),
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"is_in_master",
				},
			},
		}).Return(nil, errors.New("some error"))

		fx.syncer.handleGames(fx.ctx, []model.Game{
			{
				ExternalID:  maybe.Just(int32(222)),
				Name:        maybe.Just("new name"),
				DateTime:    timeutils.ConvertTime("2023-01-02 16:30"),
				PaymentType: maybe.Just("cash"),
			},
			{
				ExternalID:  maybe.Just(int32(333)),
				Name:        maybe.Just("new name"),
				DateTime:    timeutils.ConvertTime("2023-01-03 16:30"),
				PaymentType: maybe.Just("cash"),
			},
			{
				ExternalID:  maybe.Just(int32(444)),
				Name:        maybe.Nothing[string](),
				DateTime:    timeutils.ConvertTime("2023-01-04 16:30"),
				PaymentType: maybe.Just("cash"),
			},
			{
				ExternalID:  maybe.Just(int32(555)),
				Name:        maybe.Nothing[string](),
				DateTime:    timeutils.ConvertTime("2023-01-05 16:30"),
				PaymentType: maybe.Just("cash"),
			},
		}, []model.Game{
			{
				ExternalID:  maybe.Just(int32(111)),
				Name:        maybe.Nothing[string](),
				DateTime:    timeutils.ConvertTime("2023-01-01 16:30"),
				PaymentType: maybe.Just("cash"),
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(222)),
				Name:        maybe.Just("old name"),
				DateTime:    timeutils.ConvertTime("2023-01-02 16:30"),
				PaymentType: maybe.Just("cash"),
			},
			{
				ExternalID:  maybe.Just(int32(333)),
				Name:        maybe.Just("old name"),
				DateTime:    timeutils.ConvertTime("2023-01-03 16:30"),
				PaymentType: maybe.Just("cash"),
			},
			{
				ExternalID:  maybe.Just(int32(666)),
				Name:        maybe.Nothing[string](),
				DateTime:    timeutils.ConvertTime("2023-01-06 16:30"),
				PaymentType: maybe.Just("cash"),
				IsInMaster:  true,
			},
		})
	})
}

func Test_isEqual(t *testing.T) {
	type args struct {
		game1 model.Game
		game2 model.Game
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal",
			args: args{
				game1: model.Game{
					ExternalID: maybe.Just(int32(1)),
					LeagueID:   2,
					Number:     "number",
					Name:       maybe.Just("name1"),
					PlaceID:    3,
					DateTime:   timeutils.ConvertTime("2023-01-01 16:30"),
				},
				game2: model.Game{
					ExternalID: maybe.Just(int32(1)),
					LeagueID:   2,
					Number:     "number",
					Name:       maybe.Just("name2"),
					PlaceID:    3,
					DateTime:   timeutils.ConvertTime("2023-01-01 16:30"),
				},
			},
			want: true,
		},
		{
			name: "not equal",
			args: args{
				game1: model.Game{
					ExternalID: maybe.Just(int32(1)),
					LeagueID:   2,
					Number:     "number",
					PlaceID:    4,
					DateTime:   timeutils.ConvertTime("2023-01-01 16:30"),
				},
				game2: model.Game{
					ExternalID: maybe.Just(int32(1)),
					LeagueID:   2,
					Number:     "number",
					PlaceID:    3,
					DateTime:   timeutils.ConvertTime("2023-01-01 16:30"),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEqual(tt.args.game1, tt.args.game2); got != tt.want {
				t.Errorf("isEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isDeepEqual(t *testing.T) {
	type args struct {
		game1 model.Game
		game2 model.Game
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal",
			args: args{
				game1: model.Game{
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    2,
					Type:        3,
					Number:      "number",
					Name:        maybe.Just("name"),
					PlaceID:     4,
					DateTime:    timeutils.ConvertTime("2023-01-01 16:30"),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					IsInMaster:  true,
				},
				game2: model.Game{
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    2,
					Type:        3,
					Number:      "number",
					Name:        maybe.Just("name"),
					PlaceID:     4,
					DateTime:    timeutils.ConvertTime("2023-01-01 16:30"),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					IsInMaster:  true,
				},
			},
			want: true,
		},
		{
			name: "not equal",
			args: args{
				game1: model.Game{
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    2,
					Type:        3,
					Number:      "number",
					Name:        maybe.Just("name"),
					PlaceID:     4,
					DateTime:    timeutils.ConvertTime("2023-01-01 16:30"),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					IsInMaster:  true,
				},
				game2: model.Game{
					ExternalID:  maybe.Just(int32(1)),
					LeagueID:    2,
					Type:        3,
					Number:      "number2",
					Name:        maybe.Just("name"),
					PlaceID:     4,
					DateTime:    timeutils.ConvertTime("2023-01-01 16:30"),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  9,
					IsInMaster:  true,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDeepEqual(tt.args.game1, tt.args.game2); got != tt.want {
				t.Errorf("isDeepEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
