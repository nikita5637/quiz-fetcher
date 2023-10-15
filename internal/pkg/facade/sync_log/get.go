package synclog

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
)

// GetLastSync ...
func (f *Facade) GetLastSync(ctx context.Context, name string) (model.SyncLog, error) {
	var modelSyncLog model.SyncLog
	err := f.db.RunTX(ctx, "GetLastSync", func(ctx context.Context) error {
		dbSyncLogs, err := f.syncLogStorage.FindLastSync(ctx, name)
		if err != nil {
			return fmt.Errorf("finding last sync error: %w", err)
		}

		if len(dbSyncLogs) == 0 {
			return ErrSyncLogNotFound
		}

		modelSyncLog = convertDBSyncLogToModelSyncLog(dbSyncLogs[0])
		return nil
	})
	if err != nil {
		return model.SyncLog{}, fmt.Errorf("GetLastSync error: %w", err)
	}

	return modelSyncLog, nil
}
