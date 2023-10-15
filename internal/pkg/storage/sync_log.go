//go:generate mockery --case underscore --name SyncLogStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage/mysql"
	database "github.com/nikita5637/quiz-fetcher/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/tx"
)

// SyncLogStorage ...
type SyncLogStorage interface {
	CreateSyncLog(ctx context.Context, sync database.SyncLog) (int, error)
	FindLastSync(ctx context.Context, name string) ([]database.SyncLog, error)
	PatchSyncLog(ctx context.Context, sync database.SyncLog) error
}

// NewSyncLogStorage ...
func NewSyncLogStorage(driver string, txManager *tx.Manager) SyncLogStorage {
	switch driver {
	case mysql.DriverName:
		return mysql.NewSyncLogStorageAdapter(txManager)
	}

	return nil
}
