package squiz

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/clients/mocks"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"github.com/stretchr/testify/assert"
)

func TestGamesFetcher_GetGamesList(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			r := strings.NewReader(html1)
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		mockRegistratorServiceClient := mocks.NewRegistratorServiceClient(t)

		fetcher := GamesFetcher{
			placesCache:              make(map[string]int32, 0),
			registratorServiceClient: mockRegistratorServiceClient,
			url:                      svr.URL,
		}

		mockRegistratorServiceClient.EXPECT().GetPlaceByNameAndAddress(ctx, &registrator.GetPlaceByNameAndAddressRequest{
			Address: "ул. Ломоносова, 16",
			Name:    "BarBQ Night",
		}).Once().Return(&registrator.GetPlaceByNameAndAddressResponse{
			Place: &registrator.Place{
				Id: 11,
			},
		}, nil)

		mockRegistratorServiceClient.EXPECT().GetPlaceByNameAndAddress(ctx, &registrator.GetPlaceByNameAndAddressRequest{
			Address: "Александровский парк, 4, корп. 3",
			Name:    "Parkking",
		}).Once().Return(&registrator.GetPlaceByNameAndAddressResponse{
			Place: &registrator.Place{
				Id: 2,
			},
		}, nil)

		got, err := fetcher.GetGamesList(context.Background())
		assert.Len(t, got, 10)

		expect := []model.Game{
			{
				ExternalID:  231222,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "311.2",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2022-12-23 16:30"),
				Price:       1000,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  241222,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "312",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2022-12-24 13:00"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  120123,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "317",
				Name:        "",
				PlaceID:     2,
				DateTime:    convertTime("2023-01-12 16:30"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  70123,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "316.2",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2023-01-07 13:00"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  130123,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "318.1",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2023-01-13 16:30"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  281222,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "313",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2022-12-28 16:30"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  291222,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "314.1",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2022-12-29 16:30"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  301222,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "314.2",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2022-12-30 16:30"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  50123,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "315",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2023-01-05 16:30"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  60123,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "316.1",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2023-01-06 16:30"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
		}

		assert.ElementsMatch(t, expect, got)
		assert.NoError(t, err)
	})

	t.Run("test case 2", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			r := strings.NewReader(html2)
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		mockRegistratorServiceClient := mocks.NewRegistratorServiceClient(t)

		fetcher := GamesFetcher{
			placesCache:              make(map[string]int32, 0),
			registratorServiceClient: mockRegistratorServiceClient,
			url:                      svr.URL,
		}

		mockRegistratorServiceClient.EXPECT().GetPlaceByNameAndAddress(ctx, &registrator.GetPlaceByNameAndAddressRequest{
			Address: "ул. Ломоносова, 16",
			Name:    "BarBQ Night",
		}).Once().Return(&registrator.GetPlaceByNameAndAddressResponse{
			Place: &registrator.Place{
				Id: 11,
			},
		}, nil)

		mockRegistratorServiceClient.EXPECT().GetPlaceByNameAndAddress(ctx, &registrator.GetPlaceByNameAndAddressRequest{
			Address: "Александровский парк, 4, корп. 3",
			Name:    "Parkking",
		}).Once().Return(&registrator.GetPlaceByNameAndAddressResponse{
			Place: &registrator.Place{
				Id: 2,
			},
		}, nil)

		got, err := fetcher.GetGamesList(context.Background())
		assert.Len(t, got, 10)

		expect := []model.Game{
			{
				ExternalID:  140123,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "318.2",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2023-01-14 13:00"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  241222,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "312",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2022-12-24 13:00"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  120123,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "317",
				Name:        "",
				PlaceID:     2,
				DateTime:    convertTime("2023-01-12 16:30"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  70123,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "316.2",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2023-01-07 13:00"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  130123,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "318.1",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2023-01-13 16:30"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  281222,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "313",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2022-12-28 16:30"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  291222,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "314.1",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2022-12-29 16:30"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  301222,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "314.2",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2022-12-30 16:30"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  50123,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "315",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2023-01-05 16:30"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  60123,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "316.1",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2023-01-06 16:30"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
		}

		assert.ElementsMatch(t, expect, got)
		assert.NoError(t, err)
	})

	t.Run("test case 3", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			r := strings.NewReader(html3)
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		mockRegistratorServiceClient := mocks.NewRegistratorServiceClient(t)

		fetcher := GamesFetcher{
			placesCache:              make(map[string]int32, 0),
			registratorServiceClient: mockRegistratorServiceClient,
			url:                      svr.URL,
		}

		mockRegistratorServiceClient.EXPECT().GetPlaceByNameAndAddress(ctx, &registrator.GetPlaceByNameAndAddressRequest{
			Address: "ул. Ломоносова, 16",
			Name:    "BarBQ Night",
		}).Once().Return(&registrator.GetPlaceByNameAndAddressResponse{
			Place: &registrator.Place{
				Id: 11,
			},
		}, nil)

		mockRegistratorServiceClient.EXPECT().GetPlaceByNameAndAddress(ctx, &registrator.GetPlaceByNameAndAddressRequest{
			Address: "Александровский парк, 4, корп. 3",
			Name:    "Parkking",
		}).Once().Return(&registrator.GetPlaceByNameAndAddressResponse{
			Place: &registrator.Place{
				Id: 2,
			},
		}, nil)

		got, err := fetcher.GetGamesList(context.Background())
		assert.Len(t, got, 7)

		expect := []model.Game{
			{
				ExternalID:  120123,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "317",
				Name:        "",
				PlaceID:     2,
				DateTime:    convertTime("2023-01-12 16:30"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  130123,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "318.1",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2023-01-13 16:30"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  140123,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "318.2",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2023-01-14 13:00"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  200123,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "319.1",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2023-01-20 16:30"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  210123,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "319.2",
				Name:        "",
				PlaceID:     11,
				DateTime:    convertTime("2023-01-21 13:00"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  220123,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_THEMATIC),
				Number:      "7",
				Name:        "ТВ-шоу",
				PlaceID:     11,
				DateTime:    convertTime("2023-01-22 13:00"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
			{
				ExternalID:  260123,
				LeagueID:    leagueID,
				Type:        int32(registrator.GameType_GAME_TYPE_CLASSIC),
				Number:      "320.1",
				Name:        "",
				PlaceID:     2,
				DateTime:    convertTime("2023-01-26 16:30"),
				Price:       400,
				PaymentType: "cash",
				MaxPlayers:  8,
			},
		}

		assert.ElementsMatch(t, expect, got)
		assert.NoError(t, err)
	})
}
