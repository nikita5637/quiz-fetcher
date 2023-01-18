package syncer

import (
	"context"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
)

// Insert ...
func (f *Facade) Insert(ctx context.Context, syncLog model.SyncLog) (int, error) {
	return f.syncLogStorage.Insert(ctx, syncLog)
}
