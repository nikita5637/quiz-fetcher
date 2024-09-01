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
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			r := strings.NewReader(json1)
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

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(fx.ctx, "BarBQ Night", "ул. Ломоносова, 16").Times(6).Return(database.Place{
			ID:         11,
			ExternalID: 9,
		}, nil)

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(fx.ctx, "Parkking", "Александровский парк, 4, корп. 3").Times(3).Return(database.Place{
			ID:         2,
			ExternalID: 2,
		}, nil)

		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Игра на общие темы. Самый популярный вариант.").Return(1, nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Игра для любителей книг и фильмов о «Мальчике, который выжил». Вас ждут волшебные раунды на внимательность, логику и знания.").Return(2, nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Вопросы для взрослых, черный юмор, без цензуры.").Return(2, nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Игра на общие темы. Один пакет вопросов играем одновременно в разных городах.").Return(1, nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Игра с раундами по мотивам легендарных ТВ-передач.").Return(2, nil)

		got, err := fetcher.GetGamesList(context.Background())
		assert.Len(t, got, 9)

		expect := []model.Game{
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "364.2",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-08-04 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_THEMATIC),
				Number:      "12",
				Name:        maybe.Just("Гарри Поттер"),
				PlaceID:     9,
				DateTime:    convertTime("2023-08-05 13:00"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_THEMATIC),
				Number:      "15",
				Name:        maybe.Just("18+"),
				PlaceID:     2,
				DateTime:    convertTime("2023-08-10 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "365.1",
				Name:        maybe.Just("Битва городов"),
				PlaceID:     9,
				DateTime:    convertTime("2023-08-11 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "365.2",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-08-12 13:00"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "366.1",
				Name:        maybe.Nothing[string](),
				PlaceID:     2,
				DateTime:    convertTime("2023-08-17 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "366.2",
				Name:        maybe.Nothing[string](),
				PlaceID:     9,
				DateTime:    convertTime("2023-08-18 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_THEMATIC),
				Number:      "14",
				Name:        maybe.Just("ТВ-шоу"),
				PlaceID:     9,
				DateTime:    convertTime("2023-08-19 13:00"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "367",
				Name:        maybe.Nothing[string](),
				PlaceID:     2,
				DateTime:    convertTime("2023-08-24 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
		}

		assert.ElementsMatch(t, expect, got)
		assert.NoError(t, err)
	})

	t.Run("test case 2", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			r := strings.NewReader(json2)
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

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(fx.ctx, "BarBQ Night", "Ломоносова, 16").Times(7).Return(database.Place{
			ID:         11,
			ExternalID: 9,
		}, nil)

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(fx.ctx, "Parkking", "Александровский парк, 4, корп. 3").Times(5).Return(database.Place{
			ID:         2,
			ExternalID: 2,
		}, nil)

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(fx.ctx, "Rossi's Club", "Зодчего Росси, 1-3").Times(1).Return(database.Place{
			ID:         10,
			ExternalID: 8,
		}, nil)

		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Квиз с вопросами про все-все-все.").Return(1, nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Классический Squiz для новичков и опытных игроков.").Return(1, nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Игра на логику и эрудицию на общие темы.").Return(1, nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Самый популярный&nbsp;формат с раундами на разные темы.").Return(1, nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Классика с вопросами про все на свете.").Return(1, nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Игра с раундами по мотивам легендарных ТВ-передач.").Return(2, nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Участвуют только команды без опыта или с опытом не выше тройного белого пояса. Повтор вопросов классической игры #469 от 24 и 29 августа.").Return(1, nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Квиз для всех и обо всем.").Return(1, nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Универсальная игра для игроков всех уровней.").Return(1, nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "100 мультимедийных вопросов по сериалам, кино и музыке.").Return(5, nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Squiz для всех: от новичка до эрудита").Return(1, nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Игра для каждого на самые разные темы.").Return(1, nil)
		fx.gameTypeMatchStorage.EXPECT().GetGameTypeByDescription(fx.ctx, "Сквиз для команд любого уровня.").Return(1, nil)

		got, err := fetcher.GetGamesList(fx.ctx)
		assert.Len(t, got, 13)

		expect := []model.Game{
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "471.2",
				Name:        maybe.Just("Классическая игра"),
				PlaceID:     9,
				DateTime:    convertTime("2024-09-01 13:00"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "472.1",
				Name:        maybe.Just("Классическая игра"),
				PlaceID:     2,
				DateTime:    convertTime("2024-09-03 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "472.2",
				Name:        maybe.Just("Классическая игра"),
				PlaceID:     9,
				DateTime:    convertTime("2024-09-04 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "473.1",
				Name:        maybe.Just("Классическая игра"),
				PlaceID:     2,
				DateTime:    convertTime("2024-09-05 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "473.2",
				Name:        maybe.Just("Классическая игра"),
				PlaceID:     9,
				DateTime:    convertTime("2024-09-06 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_THEMATIC),
				Number:      "30",
				Name:        maybe.Just("ТВ-шоу"),
				PlaceID:     9,
				DateTime:    convertTime("2024-09-07 13:00"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "40",
				Name:        maybe.Just("Лига новичков"),
				PlaceID:     9,
				DateTime:    convertTime("2024-09-08 13:00"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "474.1",
				Name:        maybe.Just("Классическая игра"),
				PlaceID:     2,
				DateTime:    convertTime("2024-09-10 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "474.2",
				Name:        maybe.Just("Классическая игра"),
				PlaceID:     8,
				DateTime:    convertTime("2024-09-11 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_MOVIES_AND_MUSIC),
				Number:      "58",
				Name:        maybe.Just("Сериалы. Кино. Музыка"),
				PlaceID:     2,
				DateTime:    convertTime("2024-09-12 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "475.1",
				Name:        maybe.Just("Классическая игра"),
				PlaceID:     9,
				DateTime:    convertTime("2024-09-13 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "475.2",
				Name:        maybe.Just("Классическая игра"),
				PlaceID:     9,
				DateTime:    convertTime("2024-09-14 13:00"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  maxPlayers,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "475.3",
				Name:        maybe.Just("Классическая игра"),
				PlaceID:     2,
				DateTime:    convertTime("2024-09-15 12:30"),
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
