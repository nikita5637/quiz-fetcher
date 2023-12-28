package shaker

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

func TestFetcher_GetGamesList(t *testing.T) {
	t.Run("ok 1", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var r io.Reader
			switch req.URL.Path {
			case gamesListPath:
				r = strings.NewReader(json1)
			}
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		mockGameTypeMatchStorage := mocks.NewGameTypeMatchStorage(t)
		mockPlaceStorage := mocks.NewPlaceStorage(t)

		fetcher := Fetcher{
			client:               *http.DefaultClient,
			gamesListPath:        gamesListPath,
			gameTypeMatchStorage: mockGameTypeMatchStorage,
			placeStorage:         mockPlaceStorage,
			url:                  svr.URL,
		}

		mockGameTypeMatchStorage.EXPECT().GetGameTypeByName(ctx, "ИРОНИЯ СУДЬБЫ, ИЛИ С ЛЁГКИМ ПАРОМ!").Return(2, nil)
		mockGameTypeMatchStorage.EXPECT().GetGameTypeByName(ctx, "#ОБОВСЁМ [НОВОГОДНИЙ]").Return(1, nil)
		mockGameTypeMatchStorage.EXPECT().GetGameTypeByName(ctx, "ГАРРИ ПОТТЕР [ФИЛЬМЫ И КНИГИ. С НОВОГОДНИМ ТУРОМ]").Return(2, nil)

		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "бар \"BarBQ night\"", "Ломоносова 16").Return(database.Place{
			ID:         1,
			ExternalID: 9,
		}, nil)
		mockPlaceStorage.EXPECT().GetPlaceByNameAndAddress(ctx, "ресторан \"Parkking\"", "Александровский парк 4 лит В").Return(database.Place{
			ID:         1,
			ExternalID: 2,
		}, nil)
		expected := []model.Game{
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_THEMATIC),
				Number:      "1",
				Name:        maybe.Just("ИРОНИЯ СУДЬБЫ, ИЛИ С ЛЁГКИМ ПАРОМ!"),
				PlaceID:     9,
				DateTime:    time_utils.ConvertTime("2023-12-26 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  10,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_CLASSIC),
				Number:      "5.1",
				Name:        maybe.Just("#ОБОВСЁМ [НОВОГОДНИЙ]"),
				PlaceID:     9,
				DateTime:    time_utils.ConvertTime("2023-12-27 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  10,
				IsInMaster:  true,
			},
			{
				ExternalID:  maybe.Nothing[int32](),
				LeagueID:    leagueID,
				Type:        int32(gamepb.GameType_GAME_TYPE_THEMATIC),
				Number:      "1.1",
				Name:        maybe.Just("ГАРРИ ПОТТЕР [ФИЛЬМЫ И КНИГИ. С НОВОГОДНИМ ТУРОМ]"),
				PlaceID:     2,
				DateTime:    time_utils.ConvertTime("2023-12-27 16:30"),
				Price:       500,
				PaymentType: maybe.Just("cash"),
				MaxPlayers:  10,
				IsInMaster:  true,
			},
		}

		got, err := fetcher.GetGamesList(ctx)
		assert.Equal(t, expected, got)
		assert.NoError(t, err)
	})
}
