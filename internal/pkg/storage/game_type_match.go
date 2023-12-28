//go:generate mockery --case underscore --name GameTypeMatchStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/tx"
)

// GameTypeMatchStorage ...
type GameTypeMatchStorage interface {
	GetGameTypeByDescription(ctx context.Context, description string) (int32, error)
	GetGameTypeByName(ctx context.Context, name string) (int32, error)
}

// NewGameTypeMatchStorage ...
func NewGameTypeMatchStorage(driver string, txManager *tx.Manager) GameTypeMatchStorage {
	switch driver {
	case mysql.DriverName:
		return mysql.NewGameTypeMatchStorageAdapter(txManager)
	}

	return nil
}
