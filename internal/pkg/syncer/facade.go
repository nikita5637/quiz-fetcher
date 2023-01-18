package syncer

import "github.com/nikita5637/quiz-fetcher/internal/pkg/storage"

// Facade ...
type Facade struct {
	syncLogStorage storage.SyncLogStorage
}

// Config ...
type Config struct {
	SyncLogStorage storage.SyncLogStorage
}

// NewFacade ...
func NewFacade(cfg Config) *Facade {
	return &Facade{
		syncLogStorage: cfg.SyncLogStorage,
	}
}
