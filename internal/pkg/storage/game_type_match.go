package storage

import (
	"context"

	"github.com/nikita5637/quiz-fetcher/internal/config"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/tx"
)

// GameTypeMatchStorage ...
type GameTypeMatchStorage interface {
	GetGameTypeByDescription(ctx context.Context, description string) (int32, error)
}

// NewGameTypeMatchStorage ...
func NewGameTypeMatchStorage(txManager *tx.Manager) GameTypeMatchStorage {
	switch config.GetValue("Driver").String() {
	case config.DriverMySQL:
		return mysql.NewGameTypeMatchStorageAdapter(txManager)
	}

	return nil
}
