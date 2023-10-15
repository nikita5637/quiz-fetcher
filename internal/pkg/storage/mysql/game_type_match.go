package mysql

import (
	"context"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/tx"
)

// GameTypeMatchStorageAdapter ...
type GameTypeMatchStorageAdapter struct {
	gameTypeMatchStorage *GameTypeMatchStorage
}

// NewGameTypeMatchStorageAdapter ...
func NewGameTypeMatchStorageAdapter(txManager *tx.Manager) *GameTypeMatchStorageAdapter {
	return &GameTypeMatchStorageAdapter{
		gameTypeMatchStorage: NewGameTypeMatchStorage(txManager),
	}
}

// GetGameTypeByDescription ...
func (a *GameTypeMatchStorageAdapter) GetGameTypeByDescription(ctx context.Context, description string) (int32, error) {
	gameTypeMatchDB, err := a.gameTypeMatchStorage.GetGameTypeMatchByDescription(ctx, description)
	if err != nil {
		return 0, err
	}

	return int32(gameTypeMatchDB.GameType), nil
}
