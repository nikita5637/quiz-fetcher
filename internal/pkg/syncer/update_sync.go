package syncer

import (
	"context"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
)

// Update ...
func (f *Facade) Update(ctx context.Context, syncLog model.SyncLog) error {
	return f.syncLogStorage.Update(ctx, syncLog)
}
