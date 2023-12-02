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
	// open league
	game17630 = "/game/17630/"
	game17631 = "/game/17631/"
	game17632 = "/game/17632/"
	game17633 = "/game/17633/"

	// open league
	game22678 = "/game/22678/"
	game22679 = "/game/22679/"

	// first league
	game18615 = "/game/18615/"
	game18616 = "/game/18616/"
	game18617 = "/game/18617/"

	// first league
	game21281 = "/game/21281/"
	game21282 = "/game/21282/"
	game21283 = "/game/21283/"
)

func TestGamesFetcher_GetGamesList(t *testing.T) {
	t.Run("test case 1 (open league only)", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var r io.Reader
			switch req.URL.Path {
			case openLeagueGamesListPath:
				r = strings.NewReader(html1)
			case game17630:
				r = strings.NewReader(html17630)
			case game17631:
				r = strings.NewReader(html17631)
			case game17632:
				r = strings.NewReader(html17632)
			}
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		mockPlaceStorage := mocks.NewPlaceStorage(t)

		fetcher := Fetcher{
			openLeagueGamesListPath: openLeagueGamesListPath,
			client:                  *http.DefaultClient,
			placeStorage:            mockPlaceStorage,
			url:                     svr.URL,
		}

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "Дворец «Олимпия»", "Литейный пр., д. 14").Return(database.Place{
			ID:         1,
			ExternalID: 1,
		}, nil)

		got, err := fetcher.GetGamesList(ctx)
		assert.Equal(t, []model.Game{
			{
				ExternalID:  maybe.Just(int32(17630)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "#8",
				Name:        maybe.Just("Открытая лига"),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-02-06 16:30"),
				Price:       1200,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(17631)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "#9",
				Name:        maybe.Just("Открытая лига"),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-02-13 16:30"),
				Price:       1200,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(17632)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "#10",
				Name:        maybe.Just("Открытая лига"),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-02-20 16:30"),
				Price:       1200,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
		}, got)
		assert.NoError(t, err)
	})

	t.Run("test case 2 (open league only)", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var r io.Reader
			switch req.URL.Path {
			case openLeagueGamesListPath:
				r = strings.NewReader(html2)
			case game17631:
				r = strings.NewReader(html17631)
			case game17632:
				r = strings.NewReader(html17632)
			case game17633:
				r = strings.NewReader(html17633)
			}
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		mockPlaceStorage := mocks.NewPlaceStorage(t)

		fetcher := Fetcher{
			openLeagueGamesListPath: openLeagueGamesListPath,
			client:                  *http.DefaultClient,
			placeStorage:            mockPlaceStorage,
			url:                     svr.URL,
		}

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "Дворец «Олимпия»", "Литейный пр., д. 14").Return(database.Place{
			ID:         1,
			ExternalID: 1,
		}, nil)

		got, err := fetcher.GetGamesList(ctx)
		assert.Equal(t, []model.Game{
			{
				ExternalID:  maybe.Just(int32(17631)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "#9",
				Name:        maybe.Just("Открытая лига"),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-02-13 16:30"),
				Price:       1200,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(17632)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "#10",
				Name:        maybe.Just("Открытая лига"),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-02-20 16:30"),
				Price:       1200,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(17633)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "Финал",
				Name:        maybe.Just("Открытая лига"),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-02-27 16:30"),
				Price:       1200,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
		}, got)
		assert.NoError(t, err)
	})

	t.Run("test case 3 (first league only)", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var r io.Reader
			switch req.URL.Path {
			case firstLeagueGamesListPath:
				r = strings.NewReader(html3)
			case game18615:
				r = strings.NewReader(html18615)
			case game18616:
				r = strings.NewReader(html18616)
			case game18617:
				r = strings.NewReader(html18617)
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
				ExternalID:  maybe.Just(int32(18615)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "#4",
				Name:        maybe.Just("Первая лига"),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-04-04 16:30"),
				Price:       1200,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(18616)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "#5",
				Name:        maybe.Just("Первая лига"),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-04-11 16:30"),
				Price:       1200,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(18617)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "#6",
				Name:        maybe.Just("Первая лига"),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-04-18 16:30"),
				Price:       1200,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
		}, got)
		assert.NoError(t, err)
	})

	t.Run("test case 4 (first league only)", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var r io.Reader
			switch req.URL.Path {
			case firstLeagueGamesListPath:
				r = strings.NewReader(html4)
			case game21281:
				r = strings.NewReader(html21281)
			case game21282:
				r = strings.NewReader(html21282)
			case game21283:
				r = strings.NewReader(html21283)
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

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "Rossi's Club", "ул. Зодчего Росси, 1-3").Return(database.Place{
			ID:         1,
			ExternalID: 1,
		}, nil)

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "Дворец «Олимпия»", "Литейный пр., д. 14").Return(database.Place{
			ID:         2,
			ExternalID: 2,
		}, nil)

		got, err := fetcher.GetGamesList(ctx)
		assert.Equal(t, []model.Game{
			{
				ExternalID:  maybe.Just(int32(21281)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "#8",
				Name:        maybe.Just("Первая лига"),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-10-31 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(21282)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "#9",
				Name:        maybe.Just("Первая лига"),
				PlaceID:     2,
				DateTime:    time_utils.ConvertTime("2023-11-07 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(21283)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "#10",
				Name:        maybe.Just("Первая лига"),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-11-14 16:30"),
				Price:       1500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
		}, got)
		assert.NoError(t, err)
	})

	t.Run("test case 5 (open league only)", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var r io.Reader
			switch req.URL.Path {
			case openLeagueGamesListPath:
				r = strings.NewReader(html5)
			case game22678:
				r = strings.NewReader(html22678)
			case game22679:
				r = strings.NewReader(html22679)
			}
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		mockPlaceStorage := mocks.NewPlaceStorage(t)

		fetcher := Fetcher{
			openLeagueGamesListPath: openLeagueGamesListPath,
			client:                  *http.DefaultClient,
			placeStorage:            mockPlaceStorage,
			url:                     svr.URL,
		}

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "Дворец «Олимпия»", "Литейный пр., д. 14").Return(database.Place{
			ID:         1,
			ExternalID: 1,
		}, nil)

		got, err := fetcher.GetGamesList(ctx)
		assert.Equal(t, []model.Game{
			{
				ExternalID:  maybe.Just(int32(22678)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "#1",
				Name:        maybe.Just("Открытая лига"),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-12-04 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(22679)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "#2",
				Name:        maybe.Just("Открытая лига"),
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-12-18 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  6,
				IsInMaster:  true,
			},
		}, got)
		assert.NoError(t, err)
	})
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
			name: "tc1",
			args: args{
				price: string([]byte{194, 160, 52, 48, 48, 32, 209, 128, 209, 131, 208, 177, 46, 32, 209, 129, 32, 209, 135, 208, 181, 208, 187, 208, 190, 208, 178, 208, 181, 208, 186, 208, 176}),
			},
			want:    400,
			wantErr: false,
		},
		{
			name: "tc2",
			args: args{
				price: string([]byte{194, 160, 49, 53, 48, 48, 32, 209, 128, 209, 131, 208, 177, 46, 32, 209, 129, 32, 208, 186, 208, 190, 208, 188, 208, 176, 208, 189, 208, 180, 209, 139}),
			},
			want:    1500,
			wantErr: false,
		},
		{
			name: "invalid string",
			args: args{
				price: "invalid price",
			},
			want:    0,
			wantErr: true,
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
