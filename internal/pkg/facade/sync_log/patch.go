package synclog

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
)

// PatchSyncLog ...
func (f *Facade) PatchSyncLog(ctx context.Context, syncLog model.SyncLog) (model.SyncLog, error) {
	err := f.db.RunTX(ctx, "PatchSyncLog", func(ctx context.Context) error {
		return f.syncLogStorage.PatchSyncLog(ctx, convertModelSyncLogToDBSyncLog(syncLog))
	})
	if err != nil {
		return model.SyncLog{}, fmt.Errorf("PatchSyncLog error: %w", err)
	}

	return syncLog, nil
}
