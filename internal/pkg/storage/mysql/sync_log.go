package mysql

import (
	"context"
	"errors"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/tx"
)

// SyncLogStorageAdapter ...
type SyncLogStorageAdapter struct {
	syncLogStorage *SyncLogStorage
}

// NewSyncLogStorageAdapter ...
func NewSyncLogStorageAdapter(txManager *tx.Manager) *SyncLogStorageAdapter {
	return &SyncLogStorageAdapter{
		syncLogStorage: NewSyncLogStorage(txManager),
	}
}

// FindLastSync ...
func (a *SyncLogStorageAdapter) FindLastSync(ctx context.Context, name string) (model.SyncLog, error) {
	syncLogsDB, err := a.syncLogStorage.Find(ctx, builder.And(
		builder.Eq{
			"name": name,
		},
	), "-id")
	if err != nil {
		return model.SyncLog{}, err
	}

	if len(syncLogsDB) == 0 {
		return model.SyncLog{}, errors.New("sync log not found")
	}

	return convertDBSyncLogToModelSyncLog(syncLogsDB[0]), nil
}

// Insert ...
func (a *SyncLogStorageAdapter) Insert(ctx context.Context, syncLog model.SyncLog) (int, error) {
	return a.syncLogStorage.Insert(ctx, convertModelSyncLogToDBSyncLog(syncLog))
}

// Update ...
func (a *SyncLogStorageAdapter) Update(ctx context.Context, syncLog model.SyncLog) error {
	return a.syncLogStorage.Update(ctx, convertModelSyncLogToDBSyncLog(syncLog))
}

func convertDBSyncLogToModelSyncLog(syncLog SyncLog) model.SyncLog {
	return model.SyncLog{
		ID:         syncLog.ID,
		Name:       syncLog.Name,
		LastSyncAt: syncLog.LastSyncAt,
		Status:     model.SyncStatus(syncLog.Status),
	}
}

func convertModelSyncLogToDBSyncLog(syncLog model.SyncLog) SyncLog {
	return SyncLog{
		ID:         syncLog.ID,
		Name:       syncLog.Name,
		LastSyncAt: syncLog.LastSyncAt,
		Status:     int(syncLog.Status),
	}
}
