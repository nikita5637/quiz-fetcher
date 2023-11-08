package synchronizer

import (
	"errors"
	"testing"
	"time"

	synclog "github.com/nikita5637/quiz-fetcher/internal/pkg/facade/sync_log"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	timeutils "github.com/nikita5637/quiz-fetcher/utils/time"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

func TestSynchronizer_sync(t *testing.T) {
	timeutils.TimeNow = func() time.Time {
		return timeutils.ConvertTime("2006-01-02 15:04")
	}

	t.Run("context done", func(t *testing.T) {
		fx := tearUp(t)
		ctx, cancel := context.WithCancel(fx.ctx)
		cancel()

		fx.synchronizer.sync(ctx)
	})

	t.Run("syncers disabled", func(t *testing.T) {
		fx := tearUp(t)

		fx.syncer1.EXPECT().Enabled().Return(false)
		fx.syncer2.EXPECT().Enabled().Return(false)

		fx.synchronizer.sync(fx.ctx)
	})

	t.Run("error: get last sync internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.syncer1.EXPECT().Enabled().Return(true)
		fx.syncer2.EXPECT().Enabled().Return(false)

		fx.syncer1.EXPECT().GetName().Return("syncer")

		fx.syncLogsFacade.EXPECT().GetLastSync(mock.Anything, "syncer").Return(model.SyncLog{}, errors.New("some error"))

		fx.synchronizer.sync(fx.ctx)
	})

	t.Run("error: create sync log error", func(t *testing.T) {
		fx := tearUp(t)

		fx.syncer1.EXPECT().Enabled().Return(true)
		fx.syncer2.EXPECT().Enabled().Return(false)

		fx.syncer1.EXPECT().GetName().Return("syncer")

		fx.syncLogsFacade.EXPECT().GetLastSync(mock.Anything, "syncer").Return(model.SyncLog{}, synclog.ErrSyncLogNotFound)
		fx.syncLogsFacade.EXPECT().CreateSyncLog(mock.Anything, model.SyncLog{
			Name:       "syncer",
			LastSyncAt: timeutils.TimeNow(),
		}).Return(model.SyncLog{}, errors.New("some error"))

		fx.synchronizer.sync(fx.ctx)
	})
}
