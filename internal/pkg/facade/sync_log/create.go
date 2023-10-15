package synclog

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
)

// CreateSyncLog ...
func (f *Facade) CreateSyncLog(ctx context.Context, syncLog model.SyncLog) (model.SyncLog, error) {
	createdSyncLog := syncLog
	err := f.db.RunTX(ctx, "CreateSyncLog", func(ctx context.Context) error {
		id, err := f.syncLogStorage.CreateSyncLog(ctx, convertModelSyncLogToDBSyncLog(syncLog))
		if err != nil {
			return fmt.Errorf("creating sync log error: %w", err)
		}

		createdSyncLog.ID = id

		return nil
	})
	if err != nil {
		return model.SyncLog{}, fmt.Errorf("CreateSyncLog error: %w", err)
	}

	return createdSyncLog, nil
}
