//go:generate mockery --case underscore --name Syncer --with-expecter
//go:generate mockery --case underscore --name SyncLogsFacade --with-expecter

package synchronizer

import (
	"context"
	"time"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
)

// Syncer ...
type Syncer interface {
	Enabled() bool
	GetName() string
	GetPeriod() time.Duration
	Sync(ctx context.Context) error
}

// SyncLogsFacade ...
type SyncLogsFacade interface {
	CreateSyncLog(ctx context.Context, syncLog model.SyncLog) (model.SyncLog, error)
	GetLastSync(ctx context.Context, name string) (model.SyncLog, error)
	PatchSyncLog(ctx context.Context, syncLog model.SyncLog) (model.SyncLog, error)
}

// Synchronizer ...
type Synchronizer struct {
	syncers        []Syncer
	syncLogs       map[string]model.SyncLog
	syncLogsFacade SyncLogsFacade
	timeout        time.Duration
}

// Config ...
type Config struct {
	Syncers        []Syncer
	SyncLogsfacade SyncLogsFacade
	Timeout        time.Duration
}

// New ...
func New(cfg Config) *Synchronizer {
	return &Synchronizer{
		syncers:        cfg.Syncers,
		syncLogs:       make(map[string]model.SyncLog, len(cfg.Syncers)),
		syncLogsFacade: cfg.SyncLogsfacade,
		timeout:        cfg.Timeout,
	}
}
