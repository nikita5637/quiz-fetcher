package mysql

import (
	"context"
	"database/sql"

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
	gameTypeMatchDB, err := a.gameTypeMatchStorage.GetGameTypeMatchByDescription(ctx, sql.NullString{
		String: description,
		Valid:  description != "",
	})
	if err != nil {
		return 0, err
	}

	return int32(gameTypeMatchDB.GameType), nil
}

// GetGameTypeByName ...
func (a *GameTypeMatchStorageAdapter) GetGameTypeByName(ctx context.Context, name string) (int32, error) {
	gameTypeMatchDB, err := a.gameTypeMatchStorage.GetGameTypeMatchByName(ctx, sql.NullString{
		String: name,
		Valid:  name != "",
	})
	if err != nil {
		return 0, err
	}

	return int32(gameTypeMatchDB.GameType), nil
}
