package mysql

import (
	"context"

	"github.com/go-xorm/builder"
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

// CreateSyncLog ...
func (a *SyncLogStorageAdapter) CreateSyncLog(ctx context.Context, syncLog SyncLog) (int, error) {
	return a.syncLogStorage.Insert(ctx, syncLog)
}

// FindLastSync ...
func (a *SyncLogStorageAdapter) FindLastSync(ctx context.Context, name string) ([]SyncLog, error) {
	return a.syncLogStorage.Find(ctx, builder.And(
		builder.Eq{
			"name": name,
		},
	), "-id")
}

// PatchSyncLog ...
func (a *SyncLogStorageAdapter) PatchSyncLog(ctx context.Context, syncLog SyncLog) error {
	return a.syncLogStorage.Update(ctx, syncLog)
}
