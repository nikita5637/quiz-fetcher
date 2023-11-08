package games

import (
	"context"
	"reflect"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/syncer/games/mocks"
	timeutils "github.com/nikita5637/quiz-fetcher/utils/time"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type fixture struct {
	ctx               context.Context
	gameServiceClient *mocks.GameServiceClient
	gamesFetcher1     *mocks.Fetcher
	gamesFetcher2     *mocks.Fetcher

	syncer *Syncer
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx:               context.Background(),
		gameServiceClient: mocks.NewGameServiceClient(t),
		gamesFetcher1:     mocks.NewFetcher(t),
		gamesFetcher2:     mocks.NewFetcher(t),
	}

	fx.syncer = &Syncer{
		fetchers: []Fetcher{
			fx.gamesFetcher1,
			fx.gamesFetcher2,
		},
		gameServiceClient: fx.gameServiceClient,
	}

	t.Cleanup(func() {})

	return fx
}

func Test_convertModelGameToProtoGame(t *testing.T) {
	type args struct {
		game model.Game
	}
	tests := []struct {
		name string
		args args
		want *gamepb.Game
	}{
		{
			name: "tc1",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(2)),
					LeagueID:    3,
					Type:        4,
					Number:      "number",
					Name:        maybe.Just("name"),
					PlaceID:     5,
					DateTime:    timeutils.ConvertTime("2023-01-01 16:30"),
					Price:       400,
					PaymentType: maybe.Just("cash"),
					MaxPlayers:  8,
					IsInMaster:  true,
				},
			},
			want: &gamepb.Game{
				Id:          1,
				ExternalId:  wrapperspb.Int32(2),
				LeagueId:    3,
				Type:        4,
				Number:      "number",
				Name:        wrapperspb.String("name"),
				PlaceId:     5,
				Date:        timestamppb.New(timeutils.ConvertTime("2023-01-01 16:30")),
				Price:       400,
				PaymentType: wrapperspb.String("cash"),
				MaxPlayers:  8,
				IsInMaster:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertModelGameToProtoGame(tt.args.game); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertModelGameToProtoGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertProtoGameToModelGame(t *testing.T) {
	type args struct {
		game *gamepb.Game
	}
	tests := []struct {
		name string
		args args
		want model.Game
	}{
		{
			name: "tc1",
			args: args{
				game: &gamepb.Game{
					Id:          1,
					ExternalId:  wrapperspb.Int32(2),
					LeagueId:    3,
					Type:        4,
					Number:      "number",
					Name:        wrapperspb.String("name"),
					PlaceId:     5,
					Date:        timestamppb.New(timeutils.ConvertTime("2023-01-01 16:30")),
					Price:       400,
					PaymentType: wrapperspb.String("cash"),
					MaxPlayers:  9,
					IsInMaster:  true,
				},
			},
			want: model.Game{
				ID:          1,
				ExternalID:  maybe.Just(int32(2)),
				LeagueID:    3,
				Type:        4,
				Number:      "number",
				Name:        maybe.Just("name"),
				PlaceID:     5,
				DateTime:    timeutils.ConvertTime("2023-01-01 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  9,
				IsInMaster:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertProtoGameToModelGame(tt.args.game); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertProtoGameToModelGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
