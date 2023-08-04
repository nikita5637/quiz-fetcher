package squiz

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage/mocks"
	database "github.com/nikita5637/quiz-fetcher/internal/pkg/storage/mysql"
	commonpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/common"
	"github.com/stretchr/testify/assert"
)

func TestGamesFetcher_GetGamesList(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			r := strings.NewReader(json1)
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		mockPlaceStorage := mocks.NewPlaceStorage(t)

		fx := tearUp(t)

		fetcher := GamesFetcher{
			gameTypeMatchStorage: fx.gameTypeMatchStorage,
			placeStorage:         mockPlaceStorage,
			url:                  svr.URL,
		}

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "BarBQ Night", "ул. Ломоносова, 16").Times(6).Return(database.Place{
			ID:         11,
			ExternalID: 9,
		}, nil)

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "Parkking", "Александровский парк, 4, корп. 3").Times(3).Return(database.Place{
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
				LeagueID:    leagueID,
				Type:        int32(commonpb.GameType_GAME_TYPE_CLASSIC),
				Number:      "364.2",
				Name:        "",
				PlaceID:     9,
				DateTime:    convertTime("2023-08-04 16:30"),
				Price:       500,
				PaymentType: "cash",
				MaxPlayers:  maxPlayers,
			},
			{
				LeagueID:    leagueID,
				Type:        int32(commonpb.GameType_GAME_TYPE_THEMATIC),
				Number:      "12",
				Name:        "Гарри Поттер",
				PlaceID:     9,
				DateTime:    convertTime("2023-08-05 13:00"),
				Price:       500,
				PaymentType: "cash",
				MaxPlayers:  maxPlayers,
			},
			{
				LeagueID:    leagueID,
				Type:        int32(commonpb.GameType_GAME_TYPE_THEMATIC),
				Number:      "15",
				Name:        "18+",
				PlaceID:     2,
				DateTime:    convertTime("2023-08-10 16:30"),
				Price:       500,
				PaymentType: "cash",
				MaxPlayers:  maxPlayers,
			},
			{
				LeagueID:    leagueID,
				Type:        int32(commonpb.GameType_GAME_TYPE_CLASSIC),
				Number:      "365.1",
				Name:        "Битва городов",
				PlaceID:     9,
				DateTime:    convertTime("2023-08-11 16:30"),
				Price:       500,
				PaymentType: "cash",
				MaxPlayers:  maxPlayers,
			},
			{
				LeagueID:    leagueID,
				Type:        int32(commonpb.GameType_GAME_TYPE_CLASSIC),
				Number:      "365.2",
				Name:        "",
				PlaceID:     9,
				DateTime:    convertTime("2023-08-12 13:00"),
				Price:       500,
				PaymentType: "cash",
				MaxPlayers:  maxPlayers,
			},
			{
				LeagueID:    leagueID,
				Type:        int32(commonpb.GameType_GAME_TYPE_CLASSIC),
				Number:      "366.1",
				Name:        "",
				PlaceID:     2,
				DateTime:    convertTime("2023-08-17 16:30"),
				Price:       500,
				PaymentType: "cash",
				MaxPlayers:  maxPlayers,
			},
			{
				LeagueID:    leagueID,
				Type:        int32(commonpb.GameType_GAME_TYPE_CLASSIC),
				Number:      "366.2",
				Name:        "",
				PlaceID:     9,
				DateTime:    convertTime("2023-08-18 16:30"),
				Price:       500,
				PaymentType: "cash",
				MaxPlayers:  maxPlayers,
			},
			{
				LeagueID:    leagueID,
				Type:        int32(commonpb.GameType_GAME_TYPE_THEMATIC),
				Number:      "14",
				Name:        "ТВ-шоу",
				PlaceID:     9,
				DateTime:    convertTime("2023-08-19 13:00"),
				Price:       500,
				PaymentType: "cash",
				MaxPlayers:  maxPlayers,
			},
			{
				LeagueID:    leagueID,
				Type:        int32(commonpb.GameType_GAME_TYPE_CLASSIC),
				Number:      "367",
				Name:        "",
				PlaceID:     2,
				DateTime:    convertTime("2023-08-24 16:30"),
				Price:       500,
				PaymentType: "cash",
				MaxPlayers:  maxPlayers,
			},
		}

		assert.ElementsMatch(t, expect, got)
		assert.NoError(t, err)
	})
}
