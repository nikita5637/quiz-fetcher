package quiz_please

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
	t.Run("ok. number of rounds eq 7", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var r *strings.Reader
			switch req.URL.Query()["id"][0] {
			case "57178":
				r = strings.NewReader(html57178)
			}
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		fetcher := Fetcher{
			client:         *http.DefaultClient,
			gameResultPath: gameResultPath,
			name:           fetcherName,
			team:           "Жизнь и грабля",
			url:            svr.URL,
		}

		gameResult, err := fetcher.GetGameResult(ctx, 57178)
		assert.Equal(t, model.GameResult{
			ResultPlace: 4,
			RoundPoints: maybe.Just(`{"1 раунд":6,"2 раунд":6,"3 раунд":3,"4 раунд":10,"5 раунд":6,"6 раунд":6,"7 раунд":15,"итого":52}`),
		}, gameResult)
		assert.NoError(t, err)
	})

	t.Run("ok. number of rounds eq 8. music party online", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var r *strings.Reader
			switch req.URL.Query()["id"][0] {
			case "65277":
				r = strings.NewReader(html65277)
			}
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		fetcher := Fetcher{
			client:         *http.DefaultClient,
			gameResultPath: gameResultPath,
			name:           fetcherName,
			team:           "Название команды 3",
			url:            svr.URL,
		}

		gameResult, err := fetcher.GetGameResult(ctx, 65277)
		assert.Equal(t, model.GameResult{
			ResultPlace: 5,
			RoundPoints: maybe.Just(`{"1 раунд":10,"2 раунд":10,"3 раунд":10,"4 раунд":10,"5 раунд":18,"6 раунд":9,"7 раунд":6,"8 раунд":17,"итого":90}`),
		}, gameResult)
		assert.NoError(t, err)
	})

	t.Run("ok. number of rounds eq 8. music party offline", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var r *strings.Reader
			switch req.URL.Query()["id"][0] {
			case "65259":
				r = strings.NewReader(html65259)
			}
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		fetcher := Fetcher{
			client:         *http.DefaultClient,
			gameResultPath: gameResultPath,
			name:           fetcherName,
			team:           "Идеальные незнакомцы",
			url:            svr.URL,
		}

		gameResult, err := fetcher.GetGameResult(ctx, 65259)
		assert.Equal(t, model.GameResult{
			ResultPlace: 2,
			RoundPoints: maybe.Just(`{"1 раунд":10,"2 раунд":10,"3 раунд":8,"4 раунд":9,"5 раунд":11,"6 раунд":10,"7 раунд":10,"8 раунд":11,"итого":79}`),
		}, gameResult)
		assert.NoError(t, err)
	})

	t.Run("ok. number of rounds eq 8. with points", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var r *strings.Reader
			switch req.URL.Query()["id"][0] {
			case "63328":
				r = strings.NewReader(html63328)
			}
			_, err := io.Copy(w, r)
			assert.NoError(t, err)
		}))
		defer svr.Close()

		fetcher := Fetcher{
			client:         *http.DefaultClient,
			gameResultPath: gameResultPath,
			name:           fetcherName,
			team:           "Funny cucumber",
			url:            svr.URL,
		}

		gameResult, err := fetcher.GetGameResult(ctx, 63328)
		assert.Equal(t, model.GameResult{
			ResultPlace: 15,
			RoundPoints: maybe.Just(`{"1 раунд":10,"2 раунд":6,"3 раунд":3,"4 раунд":8,"5 раунд":11,"6 раунд":9.5,"7 раунд":10,"8 раунд":9,"итого":66.5}`),
		}, gameResult)
		assert.NoError(t, err)
	})
}
