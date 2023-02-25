package synchronizer

import (
	"context"
	"runtime/debug"
	"time"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/logger"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	time_utils "github.com/nikita5637/quiz-fetcher/utils/time"
)

// Syncer ...
type Syncer interface {
	Enabled() bool
	GetLastSyncAt() time.Time
	GetPeriod() time.Duration
	GetSyncerName() string
	GetSyncStatus() model.SyncStatus
	Sync(ctx context.Context) error
}

// Synchronizer ...
type Synchronizer struct {
	syncers []Syncer
}

// NewSynchronizer ...
func NewSynchronizer(syncers ...Syncer) *Synchronizer {
	s := make([]Syncer, 0, len(syncers))
	for _, syncer := range syncers {
		s = append(s, syncer)
	}

	return &Synchronizer{
		syncers: s,
	}
}

// Start ...
func (s *Synchronizer) Start(ctx context.Context) error {
	go func(ctx context.Context) {
		for range time.Tick(1 * time.Second) {
			for _, syncer := range s.syncers {
				if !syncer.Enabled() {
					continue
				}

				if (time_utils.TimeNow().After(syncer.GetLastSyncAt().Add(syncer.GetPeriod()*time.Second)) && syncer.GetSyncStatus() != model.SyncStatusInProgress) ||
					syncer.GetSyncStatus() == model.SyncStatusNotSynced {
					go func(ctx context.Context, syncer Syncer) {
						defer func() {
							if r := recover(); r != nil {
								logger.ErrorKV(ctx, "panic recovered", "r", r, "name", syncer.GetSyncerName(), "stack", string(debug.Stack()))
							}
						}()

						logger.InfoKV(ctx, "sync started", "name", syncer.GetSyncerName())

						if err := syncer.Sync(ctx); err != nil {
							logger.ErrorKV(ctx, "sync error", "syncer name", syncer.GetSyncerName(), "error", err)
						} else {
							logger.InfoKV(ctx, "sync done", "name", syncer.GetSyncerName())
						}
					}(ctx, syncer)
				}
			}
		}
	}(ctx)

	<-ctx.Done()

	logger.Info(ctx, "synchronizer gracefully stopped")
	return nil
}
