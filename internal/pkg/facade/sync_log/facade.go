package synclog

import (
	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/tx"
)

// Facade ...
type Facade struct {
	syncLogStorage storage.SyncLogStorage

	db *tx.Manager
}

// Config ...
type Config struct {
	SyncLogStorage storage.SyncLogStorage

	TxManager *tx.Manager
}

// New ...
func New(cfg Config) *Facade {
	return &Facade{
		syncLogStorage: cfg.SyncLogStorage,
		db:             cfg.TxManager,
	}
}
