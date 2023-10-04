//go:generate mockery --case underscore --name SyncLogStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/tx"
)

// SyncLogStorage ...
type SyncLogStorage interface {
	FindLastSync(ctx context.Context, name string) (model.SyncLog, error)
	Insert(ctx context.Context, sync model.SyncLog) (int, error)
	Update(ctx context.Context, sync model.SyncLog) error
}

// NewSyncLogStorage ...
func NewSyncLogStorage(driver string, txManager *tx.Manager) SyncLogStorage {
	switch driver {
	case mysql.DriverName:
		return mysql.NewSyncLogStorageAdapter(txManager)
	}

	return nil
}
