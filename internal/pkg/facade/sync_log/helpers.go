package synclog

import (
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	database "github.com/nikita5637/quiz-fetcher/internal/pkg/storage/mysql"
)

func convertDBSyncLogToModelSyncLog(syncLog database.SyncLog) model.SyncLog {
	return model.SyncLog{
		ID:         syncLog.ID,
		Name:       syncLog.Name,
		LastSyncAt: syncLog.LastSyncAt,
		Status:     model.SyncStatus(syncLog.Status),
	}
}

func convertModelSyncLogToDBSyncLog(syncLog model.SyncLog) database.SyncLog {
	return database.SyncLog{
		ID:         syncLog.ID,
		Name:       syncLog.Name,
		LastSyncAt: syncLog.LastSyncAt,
		Status:     int(syncLog.Status),
	}
}
