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
	"github.com/stretchr/testify/assert"
)

func TestFetcher_GetGameResult(t *testing.T) {
	t.Run("ok 18612", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var r *strings.Reader
			switch req.URL.Query().Get("game_id") {
			case "18612":
				r = strings.NewReader(json18612)
			}
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		fetcher := Fetcher{
			client:         *http.DefaultClient,
			gameResultPath: gameResultPath,
			name:           fetcherName,
			team:           "Улица плохих снов",
			url:            svr.URL,
		}

		gameResult, err := fetcher.GetGameResult(ctx, 18612)
		assert.Equal(t, model.GameResult{
			ResultPlace: 6,
			RoundPoints: maybe.Just(`{"Сумма":25,"1":8,"2":9,"3":8,"Матрица":230}`),
		}, gameResult)
		assert.NoError(t, err)
	})

	t.Run("ok 18619", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var r *strings.Reader
			switch req.URL.Query().Get("game_id") {
			case "18619":
				r = strings.NewReader(json18619)
			}
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		fetcher := Fetcher{
			client:         *http.DefaultClient,
			gameResultPath: gameResultPath,
			name:           fetcherName,
			team:           "Улица плохих снов",
			url:            svr.URL,
		}

		gameResult, err := fetcher.GetGameResult(ctx, 18619)
		assert.Equal(t, model.GameResult{
			ResultPlace: 6,
			RoundPoints: maybe.Just(`{"Сумма":23,"1":9,"2":4,"3":10,"Матрица":380}`),
		}, gameResult)
		assert.NoError(t, err)
	})

	t.Run("ok 21281", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var r *strings.Reader
			switch req.URL.Query().Get("game_id") {
			case "21281":
				r = strings.NewReader(json21281)
			}
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		fetcher := Fetcher{
			client:         *http.DefaultClient,
			gameResultPath: gameResultPath,
			name:           fetcherName,
			team:           "Улица плохих снов",
			url:            svr.URL,
		}

		gameResult, err := fetcher.GetGameResult(ctx, 21281)
		assert.Equal(t, model.GameResult{
			ResultPlace: 20,
			RoundPoints: maybe.Just(`{"Сумма":18,"1":7,"2":5,"3":6,"Матрица":80}`),
		}, gameResult)
		assert.NoError(t, err)
	})
}
