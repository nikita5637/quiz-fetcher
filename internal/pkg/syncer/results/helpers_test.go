package results

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/syncer/results/mocks"
)

type fixture struct {
	ctx context.Context

	fetcher1                       *mocks.Fetcher
	fetcher2                       *mocks.Fetcher
	gameServiceClient              *mocks.GameServiceClient
	gameResultManagerServiceClient *mocks.GameResultManagerServiceClient

	syncer *Syncer
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		fetcher1:                       mocks.NewFetcher(t),
		fetcher2:                       mocks.NewFetcher(t),
		gameServiceClient:              mocks.NewGameServiceClient(t),
		gameResultManagerServiceClient: mocks.NewGameResultManagerServiceClient(t),
	}

	fx.syncer = &Syncer{
		fetchers: []Fetcher{
			fx.fetcher1,
			fx.fetcher2,
		},
		gameServiceClient:              fx.gameServiceClient,
		gameResultManagerServiceClient: fx.gameResultManagerServiceClient,
	}

	t.Cleanup(func() {})

	return fx
}
