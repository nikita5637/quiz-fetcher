package sixty_seconds

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/clients/mocks"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	time_utils "github.com/nikita5637/quiz-fetcher/utils/time"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"github.com/stretchr/testify/assert"
)

const (
	// open league
	game17630 = "/game/17630/"
	game17631 = "/game/17631/"
	game17632 = "/game/17632/"
	game17633 = "/game/17633/"

	// first league
	game18615 = "/game/18615/"
	game18616 = "/game/18616/"
	game18617 = "/game/18617/"
)

func TestGamesFetcher_GetGamesList(t *testing.T) {
	t.Run("test case 1 (open league only)", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var r io.Reader
			switch req.URL.Path {
			case OpenLeagueGamesListPath:
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

		mockRegistratorServiceClient := mocks.NewRegistratorServiceClient(t)

		fetcher := GamesFetcher{
			openLeagueGamesListPath:  OpenLeagueGamesListPath,
			client:                   *http.DefaultClient,
			placesCache:              make(map[string]int32, 0),
			registratorServiceClient: mockRegistratorServiceClient,
			url:                      svr.URL,
		}

		mockRegistratorServiceClient.EXPECT().GetPlaceByNameAndAddress(ctx, &registrator.GetPlaceByNameAndAddressRequest{
			Address: "Литейный пр., д. 14",
			Name:    "Дворец «Олимпия»",
		}).Return(&registrator.GetPlaceByNameAndAddressResponse{
			Place: &registrator.Place{
				Id: 1,
			},
		}, nil)

		got, err := fetcher.GetGamesList(ctx)
		assert.Equal(t, []model.Game{
			{
				ExternalID:  17630,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "#8",
				Name:        "Открытая лига",
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-02-06 16:30"),
				Price:       1200,
				PaymentType: "cash",
				MaxPlayers:  6,
			},
			{
				ExternalID:  17631,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "#9",
				Name:        "Открытая лига",
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-02-13 16:30"),
				Price:       1200,
				PaymentType: "cash",
				MaxPlayers:  6,
			},
			{
				ExternalID:  17632,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "#10",
				Name:        "Открытая лига",
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-02-20 16:30"),
				Price:       1200,
				PaymentType: "cash",
				MaxPlayers:  6,
			},
		}, got)
		assert.NoError(t, err)
	})

	t.Run("test case 2 (open league only)", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var r io.Reader
			switch req.URL.Path {
			case OpenLeagueGamesListPath:
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

		mockRegistratorServiceClient := mocks.NewRegistratorServiceClient(t)

		fetcher := GamesFetcher{
			openLeagueGamesListPath:  OpenLeagueGamesListPath,
			client:                   *http.DefaultClient,
			placesCache:              make(map[string]int32, 0),
			registratorServiceClient: mockRegistratorServiceClient,
			url:                      svr.URL,
		}

		mockRegistratorServiceClient.EXPECT().GetPlaceByNameAndAddress(ctx, &registrator.GetPlaceByNameAndAddressRequest{
			Address: "Литейный пр., д. 14",
			Name:    "Дворец «Олимпия»",
		}).Return(&registrator.GetPlaceByNameAndAddressResponse{
			Place: &registrator.Place{
				Id: 1,
			},
		}, nil)

		got, err := fetcher.GetGamesList(ctx)
		assert.Equal(t, []model.Game{
			{
				ExternalID:  17631,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "#9",
				Name:        "Открытая лига",
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-02-13 16:30"),
				Price:       1200,
				PaymentType: "cash",
				MaxPlayers:  6,
			},
			{
				ExternalID:  17632,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "#10",
				Name:        "Открытая лига",
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-02-20 16:30"),
				Price:       1200,
				PaymentType: "cash",
				MaxPlayers:  6,
			},
			{
				ExternalID:  17633,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "Финал",
				Name:        "Открытая лига",
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-02-27 16:30"),
				Price:       1200,
				PaymentType: "cash",
				MaxPlayers:  6,
			},
		}, got)
		assert.NoError(t, err)
	})

	t.Run("test case 3 (first league only)", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var r io.Reader
			switch req.URL.Path {
			case FirstLeagueGamesListPath:
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

		mockRegistratorServiceClient := mocks.NewRegistratorServiceClient(t)

		fetcher := GamesFetcher{
			firstLeagueGamesListPath: FirstLeagueGamesListPath,
			client:                   *http.DefaultClient,
			placesCache:              make(map[string]int32, 0),
			registratorServiceClient: mockRegistratorServiceClient,
			url:                      svr.URL,
		}

		mockRegistratorServiceClient.EXPECT().GetPlaceByNameAndAddress(ctx, &registrator.GetPlaceByNameAndAddressRequest{
			Address: "Литейный пр., д. 14",
			Name:    "Дворец «Олимпия»",
		}).Return(&registrator.GetPlaceByNameAndAddressResponse{
			Place: &registrator.Place{
				Id: 1,
			},
		}, nil)

		got, err := fetcher.GetGamesList(ctx)
		assert.Equal(t, []model.Game{
			{
				ExternalID:  18615,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "#4",
				Name:        "Первая лига",
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-04-04 16:30"),
				Price:       1200,
				PaymentType: "cash",
				MaxPlayers:  6,
			},
			{
				ExternalID:  18616,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "#5",
				Name:        "Первая лига",
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-04-11 16:30"),
				Price:       1200,
				PaymentType: "cash",
				MaxPlayers:  6,
			},
			{
				ExternalID:  18617,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "#6",
				Name:        "Первая лига",
				PlaceID:     1,
				DateTime:    time_utils.ConvertTime("2023-04-18 16:30"),
				Price:       1200,
				PaymentType: "cash",
				MaxPlayers:  6,
			},
		}, got)
		assert.NoError(t, err)
	})
}
