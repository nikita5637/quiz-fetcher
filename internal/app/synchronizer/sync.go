package synchronizer

import (
	"context"
	"errors"
	"runtime/debug"
	"time"

	"github.com/nikita5637/quiz-fetcher/internal/app/info"
	synclog "github.com/nikita5637/quiz-fetcher/internal/pkg/facade/sync_log"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/logger"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	timeutils "github.com/nikita5637/quiz-fetcher/utils/time"
)

// Sync ...
func (s *Synchronizer) Sync(ctx context.Context) error {
	timeTicker := time.NewTicker(time.Second)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-timeTicker.C:
			ctx, _ := context.WithTimeout(ctx, s.timeout) // nolint
			ch := make(chan struct{})
			go func() {
				defer func() {
					close(ch)
					if r := recover(); r != nil {
						logger.ErrorKV(ctx, "panic recovered", "r", r, "stack", string(debug.Stack()))
					}
				}()

				s.sync(ctx)
			}()

			select {
			case <-ctx.Done():
				logger.Errorf(ctx, "synchronization interrupted: %s", ctx.Err().Error())
				if errors.Is(ctx.Err(), context.DeadlineExceeded) {
					// don't stop app when context deadline exceeded
					continue
				}

				// waiting until sync ended
				select {
				case <-ch:
					return nil
				}
			case <-ch:
			}
		}
	}
}

func (s *Synchronizer) sync(ctx context.Context) {
	for _, syncer := range s.syncers {
		select {
		case <-ctx.Done():
			return
		default:
		}

		if !syncer.Enabled() {
			continue
		}

		var err error
		syncerName := syncer.GetName()
		syncerCtx := logger.ToContext(ctx, logger.FromContext(ctx).With("syncer", syncerName))

		lastSyncLog, ok := s.syncLogs[syncerName]
		if !ok {
			// first iteration after starting app
			if lastSyncLog, err = s.syncLogsFacade.GetLastSync(syncerCtx, syncerName); err != nil {
				if errors.Is(err, synclog.ErrSyncLogNotFound) {
					if lastSyncLog, err = s.syncLogsFacade.CreateSyncLog(syncerCtx, model.SyncLog{
						Name:       syncerName,
						LastSyncAt: timeutils.TimeNow(),
						Status:     model.SyncStatusNotSynced,
					}); err != nil {
						logger.Errorf(syncerCtx, "creating sync log error: %s", err.Error())
						continue
					}
					s.syncLogs[syncerName] = lastSyncLog
				} else {
					logger.Errorf(syncerCtx, "getting last sync error: %s", err.Error())
					continue
				}
			}
		}

		if (info.StartedAt.After(lastSyncLog.LastSyncAt) && lastSyncLog.Status == model.SyncStatusInProgress) ||
			timeutils.TimeNow().After(lastSyncLog.LastSyncAt.Add(syncer.GetPeriod())) ||
			lastSyncLog.Status == model.SyncStatusNotSynced {
			logger.Info(syncerCtx, "syncing has started")

			if lastSyncLog.Status == model.SyncStatusNotSynced {
				lastSyncLog.Status = model.SyncStatusInProgress
				if _, err = s.syncLogsFacade.PatchSyncLog(syncerCtx, lastSyncLog); err != nil {
					logger.Errorf(syncerCtx, "patching sync log error: %s", err.Error())
					continue
				}
			} else {
				if lastSyncLog, err = s.syncLogsFacade.CreateSyncLog(syncerCtx, model.SyncLog{
					Name:       syncerName,
					LastSyncAt: timeutils.TimeNow(),
					Status:     model.SyncStatusInProgress,
				}); err != nil {
					logger.Errorf(syncerCtx, "creating sync log error: %s", err.Error())
					continue
				}
			}
			s.syncLogs[syncerName] = lastSyncLog

			if err := syncer.Sync(syncerCtx); err != nil {
				lastSyncLog.Status = model.SyncStatusFailed
				if _, err2 := s.syncLogsFacade.PatchSyncLog(syncerCtx, lastSyncLog); err2 != nil {
					logger.Errorf(syncerCtx, "patching sync log error: %s", err2.Error())
					continue
				}
				s.syncLogs[syncerName] = lastSyncLog

				logger.Errorf(syncerCtx, "syncing error: %s", err.Error())
				continue
			}

			lastSyncLog.Status = model.SyncStatusOK
			if _, err := s.syncLogsFacade.PatchSyncLog(syncerCtx, lastSyncLog); err != nil {
				logger.Errorf(syncerCtx, "patching sync log error: %s", err.Error())
				continue
			}
			s.syncLogs[syncerName] = lastSyncLog

			logger.Info(syncerCtx, "syncing finished")
		}
	}
}
