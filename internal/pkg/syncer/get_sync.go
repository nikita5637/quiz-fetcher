package syncer

import (
	"context"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
)

// GetLastSync ...
func (f *Facade) GetLastSync(ctx context.Context, name string) (model.SyncLog, error) {
	syncLog, err := f.syncLogStorage.FindLastSync(ctx, name)
	if err != nil {
		return model.SyncLog{}, err
	}

	return syncLog, nil
}
