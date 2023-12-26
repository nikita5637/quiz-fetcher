package sixty_seconds

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage/mocks"
	database "github.com/nikita5637/quiz-fetcher/internal/pkg/storage/mysql"
	time_utils "github.com/nikita5637/quiz-fetcher/utils/time"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"github.com/stretchr/testify/assert"
)

const (
	// first league
	game23284 = "/quizgames/game/23284/"
)

func TestFetcher_GetGamesList(t *testing.T) {
	t.Run("test case 1 (first league only)", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var r io.Reader
			switch req.URL.Path {
			case firstLeagueGamesListPath:
				r = strings.NewReader(html1)
			case game23284:
				r = strings.NewReader(html23284)
			}
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		mockPlaceStorage := mocks.NewPlaceStorage(t)

		fetcher := Fetcher{
			firstLeagueGamesListPath: firstLeagueGamesListPath,
			client:                   *http.DefaultClient,
			placeStorage:             mockPlaceStorage,
			url:                      svr.URL,
		}

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "Дворец «Олимпия»", "Литейный пр., д. 14").Return(database.Place{
			ID:         1,
			ExternalID: 1,
		}, nil)

		got, err := fetcher.GetGamesList(ctx)
		assert.Equal(t, []model.Game{
			{
				ExternalID:  maybe.Just(int32(23284)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "#4",
				Name:        maybe.Just("Первая лига"),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2024-01-09 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
		}, got)
		assert.NoError(t, err)
	})
}

func Test_getExternalID(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr bool
	}{
		{
			name: "tc1",
			args: args{
				path: "/quizgames/game/23284/",
			},
			want:    23284,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getExternalID(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("getExternalID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getExternalID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getName(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tc1",
			args: args{
				text: "Первая лига | Игра #4",
			},
			want: "Первая лига",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getName(tt.args.text); got != tt.want {
				t.Errorf("getName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNumber(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tc1",
			args: args{
				text: "Первая лига | Игра #4",
			},
			want: "#4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumber(tt.args.text); got != tt.want {
				t.Errorf("getNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPrice(t *testing.T) {
	type args struct {
		price string
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				price: "					     1500 руб. с команды",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				price: "1500 руб. с команды",
			},
			want:    1500,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPrice(tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
