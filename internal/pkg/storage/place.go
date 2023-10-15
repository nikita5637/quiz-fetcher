//go:generate mockery --case underscore --name PlaceStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage/mysql"
	database "github.com/nikita5637/quiz-fetcher/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/tx"
)

// PlaceStorage ...
type PlaceStorage interface {
	GetPlaceByNameAndAddress(ctx context.Context, name, address string) (database.Place, error)
}

// NewPlaceStorage ...
func NewPlaceStorage(driver string, txManager *tx.Manager) PlaceStorage {
	switch driver {
	case mysql.DriverName:
		return mysql.NewPlaceStorageAdapter(txManager)
	}

	return nil
}
