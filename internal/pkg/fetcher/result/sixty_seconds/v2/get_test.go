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
	t.Run("ok json 22682", func(t *testing.T) {
		ctx := context.Background()

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			r := strings.NewReader(json22682)
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

		gameResult, err := fetcher.GetGameResult(ctx, 22682)
		assert.Equal(t, model.GameResult{
			ResultPlace: 9,
			RoundPoints: maybe.Just(`{"Сумма":23,"1":11,"2":6,"3":6,"Матрица":350}`),
		}, gameResult)
		assert.NoError(t, err)
	})
}
