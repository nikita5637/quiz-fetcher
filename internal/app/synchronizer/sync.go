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
	"go.uber.org/zap"
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
						logger.ErrorKV(ctx, "panic recovered", zap.Reflect("r", r), zap.String("stack", string(debug.Stack())))
					}
				}()

				s.sync(ctx)
			}()

			select {
			case <-ctx.Done():
				logger.ErrorKV(ctx, "synchronization interrupted", zap.Error(ctx.Err()))
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
						logger.ErrorKV(syncerCtx, "creating sync log error", zap.Error(err))
						continue
					}
					s.syncLogs[syncerName] = lastSyncLog
				} else {
					logger.ErrorKV(syncerCtx, "getting last sync error", zap.Error(err))
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
					logger.ErrorKV(syncerCtx, "patching sync log error", zap.Error(err))
					continue
				}
			} else {
				if lastSyncLog, err = s.syncLogsFacade.CreateSyncLog(syncerCtx, model.SyncLog{
					Name:       syncerName,
					LastSyncAt: timeutils.TimeNow(),
					Status:     model.SyncStatusInProgress,
				}); err != nil {
					logger.ErrorKV(syncerCtx, "creating sync log error", zap.Error(err))
					continue
				}
			}
			s.syncLogs[syncerName] = lastSyncLog

			if err := syncer.Sync(syncerCtx); err != nil {
				lastSyncLog.Status = model.SyncStatusFailed
				if _, err2 := s.syncLogsFacade.PatchSyncLog(syncerCtx, lastSyncLog); err2 != nil {
					logger.ErrorKV(syncerCtx, "patching sync log error", zap.Error(err2))
					continue
				}
				s.syncLogs[syncerName] = lastSyncLog

				logger.ErrorKV(syncerCtx, "syncing error", zap.Error(err))
				continue
			}

			lastSyncLog.Status = model.SyncStatusOK
			if _, err := s.syncLogsFacade.PatchSyncLog(syncerCtx, lastSyncLog); err != nil {
				logger.ErrorKV(syncerCtx, "patching sync log error", zap.Error(err))
				continue
			}
			s.syncLogs[syncerName] = lastSyncLog

			logger.Info(syncerCtx, "syncing finished")
		}
	}
}
