package squiz

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
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
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

		mockPlaceStorage := mocks.NewPlaceStorage(t)

		fx := tearUp(t)

		fetcher := Fetcher{
			gameTypeMatchStorage: fx.gameTypeMatchStorage,
			placeStorage:         mockPlaceStorage,
			url:                  svr.URL,
		}

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "BarBQ Night", "ул. Ломоносова, 16").Times(9).Return(database.Place{
			ID:         11,
			ExternalID: 9,
		}, nil)

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "Parkking", "Александровский парк, 4, корп. 3").Once().Return(database.Place{
			ID:         2,
			ExternalID: 2,
		}, nil)

		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Игра на общие темы. Самый популярный и массовый вариант.").Return(1, nil)

		got, err := fetcher.GetGamesList(context.Background())
		assert.Len(t, got, 10)

		expect := []model.Game{
			{
				ExternalID:  maybe.Just(int32(231222)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "311.2",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2022-12-23 16:30"),
				Price:       1000,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(241222)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "312",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2022-12-24 13:00"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(120123)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "317",
				Name:        maybe.Nothing[string](),
				PlaceID:     2,
				DateTime:    convertTime("2023-01-12 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(70123)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "316.2",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-01-07 13:00"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(130123)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "318.1",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-01-13 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(281222)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "313",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2022-12-28 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(291222)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "314.1",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2022-12-29 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(301222)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "314.2",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2022-12-30 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(50123)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "315",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-01-05 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(60123)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "316.1",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-01-06 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
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

		mockPlaceStorage := mocks.NewPlaceStorage(t)

		fx := tearUp(t)

		fetcher := Fetcher{
			gameTypeMatchStorage: fx.gameTypeMatchStorage,
			placeStorage:         mockPlaceStorage,
			url:                  svr.URL,
		}

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "BarBQ Night", "ул. Ломоносова, 16").Times(9).Return(database.Place{
			ID:         11,
			ExternalID: 9,
		}, nil)

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "Parkking", "Александровский парк, 4, корп. 3").Once().Return(database.Place{
			ID:         2,
			ExternalID: 2,
		}, nil)

		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Игра на общие темы. Самый популярный и массовый вариант.").Return(1, nil)

		got, err := fetcher.GetGamesList(context.Background())
		assert.Len(t, got, 10)

		expect := []model.Game{
			{
				ExternalID:  maybe.Just(int32(140123)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "318.2",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-01-14 13:00"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(241222)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "312",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2022-12-24 13:00"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(120123)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "317",
				Name:        maybe.Nothing[string](),
				PlaceID:     2,
				DateTime:    convertTime("2023-01-12 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(70123)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "316.2",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-01-07 13:00"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(130123)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "318.1",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-01-13 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(281222)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "313",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2022-12-28 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(291222)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "314.1",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2022-12-29 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(301222)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "314.2",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2022-12-30 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(50123)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "315",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-01-05 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(60123)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "316.1",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-01-06 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
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

		mockPlaceStorage := mocks.NewPlaceStorage(t)

		fx := tearUp(t)

		fetcher := Fetcher{
			gameTypeMatchStorage: fx.gameTypeMatchStorage,
			placeStorage:         mockPlaceStorage,
			url:                  svr.URL,
		}

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "BarBQ Night", "ул. Ломоносова, 16").Times(5).Return(database.Place{
			ID:         11,
			ExternalID: 9,
		}, nil)

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "Parkking", "Александровский парк, 4, корп. 3").Twice().Return(database.Place{
			ID:         2,
			ExternalID: 2,
		}, nil)

		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Игра на общие темы. Самый популярный и массовый вариант.").Return(1, nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Фановая игра с раундами по мотивам популярных ТВ-шоу.").Return(2, nil)

		got, err := fetcher.GetGamesList(context.Background())
		assert.Len(t, got, 7)

		expect := []model.Game{
			{
				ExternalID:  maybe.Just(int32(120123)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "317",
				Name:        maybe.Nothing[string](),
				PlaceID:     2,
				DateTime:    convertTime("2023-01-12 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(130123)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "318.1",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-01-13 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(140123)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "318.2",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-01-14 13:00"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(200123)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "319.1",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-01-20 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(210123)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "319.2",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-01-21 13:00"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(220123)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_THEMATIC),
				Number:      "7",
				Name:        maybe.Just("ТВ-шоу"),
				PlaceID:     9,
				DateTime:    convertTime("2023-01-22 13:00"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(260123)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "320.1",
				Name:        maybe.Nothing[string](),
				PlaceID:     2,
				DateTime:    convertTime("2023-01-26 16:30"),
				Price:       400,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
		}

		assert.ElementsMatch(t, expect, got)
		assert.NoError(t, err)
	})

	t.Run("test case 4", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			r := strings.NewReader(html4)
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		mockPlaceStorage := mocks.NewPlaceStorage(t)

		fx := tearUp(t)

		fetcher := Fetcher{
			gameTypeMatchStorage: fx.gameTypeMatchStorage,
			placeStorage:         mockPlaceStorage,
			url:                  svr.URL,
		}

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "BarBQ Night", "ул. Ломоносова, 16").Times(6).Return(database.Place{
			ID:         11,
			ExternalID: 9,
		}, nil)

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "Parkking", "Александровский парк, 4, корп. 3").Times(4).Return(database.Place{
			ID:         2,
			ExternalID: 2,
		}, nil)

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "Цинь", "16-я линия В.О., 83").Once().Return(database.Place{
			ID:         23,
			ExternalID: 1,
		}, nil)

		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Игра на общие темы. Самый популярный и массовый вариант.").Return(int32(gamepb.GameType_GAME_TYPE_CLASSIC), nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Игра на общие темы. Самый популярный и массовый вариант. Подходит всем, идеально для новичков.").Return(int32(gamepb.GameType_GAME_TYPE_CLASSIC), nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Вопросы для взрослых, черный юмор, без цензуры. Приходите, только если уверены в себе.").Return(int32(gamepb.GameType_GAME_TYPE_THEMATIC), nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Игра, посвященная музыке девяностых.").Return(int32(gamepb.GameType_GAME_TYPE_THEMATIC_MOVIES_AND_MUSIC), nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Игра на общие темы. Пакет из очень простых вопросов - идеально для новичков.").Return(int32(gamepb.GameType_GAME_TYPE_CLASSIC), nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Вопросы только на темы кино, сериалов и музыки. Зарубежное и российское, современное и классическое. Кинофаны и меломаны - вэлкам!").Return(int32(gamepb.GameType_GAME_TYPE_MOVIES_AND_MUSIC), nil)

		got, err := fetcher.GetGamesList(context.Background())
		assert.Len(t, got, 11)

		expect := []model.Game{
			{
				ExternalID:  maybe.Just(int32(80623)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_MOVIES_AND_MUSIC),
				Number:      "40",
				Name:        maybe.Just("Сериалы. Кино. Музыка"),
				PlaceID:     2,
				DateTime:    convertTime("2023-06-08 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(90623)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "353.1",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-06-09 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(100623)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "353.2",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-06-10 13:00"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(110623)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_THEMATIC),
				Number:      "13",
				Name:        maybe.Just("18+"),
				PlaceID:     2,
				DateTime:    convertTime("2023-06-11 12:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(150623)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "10",
				Name:        maybe.Just("Лига Новичков"),
				PlaceID:     2,
				DateTime:    convertTime("2023-06-15 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(160623)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "354.1",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-06-16 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(170623)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "354.2",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-06-17 13:00"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(220623)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "355",
				Name:        maybe.Nothing[string](),
				PlaceID:     2,
				DateTime:    convertTime("2023-06-22 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(230623)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "356",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-06-23 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(0)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLOSED),
				Number:      "Сезон 9",
				Name:        maybe.Just(finalGameName),
				PlaceID:     1,
				DateTime:    convertTime("2023-06-24 12:00"),
				Price:       0,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Just(int32(250623)),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_THEMATIC_MOVIES_AND_MUSIC),
				Number:      "2",
				Name:        maybe.Just("Музыка 90-х"),
				PlaceID:     9,
				DateTime:    convertTime("2023-06-25 12:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
		}

		assert.ElementsMatch(t, expect, got)
		assert.NoError(t, err)
	})
}
