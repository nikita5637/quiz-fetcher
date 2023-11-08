package synchronizer

import (
	"context"
	"testing"
	"time"

	"github.com/nikita5637/quiz-fetcher/internal/app/synchronizer/mocks"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
)

type fixture struct {
	ctx            context.Context
	syncer1        *mocks.Syncer
	syncer2        *mocks.Syncer
	syncLogsFacade *mocks.SyncLogsFacade

	synchronizer *Synchronizer
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx:            context.Background(),
		syncer1:        mocks.NewSyncer(t),
		syncer2:        mocks.NewSyncer(t),
		syncLogsFacade: mocks.NewSyncLogsFacade(t),
	}

	fx.synchronizer = &Synchronizer{
		syncers: []Syncer{
			fx.syncer1,
			fx.syncer2,
		},
		syncLogs:       make(map[string]model.SyncLog, 2),
		syncLogsFacade: fx.syncLogsFacade,
		timeout:        600 * time.Second,
	}

	t.Cleanup(func() {})

	return fx
}
