package mysql

import (
	"context"
	"database/sql"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/tx"
)

// PlaceStorageAdapter ...
type PlaceStorageAdapter struct {
	placeStorage *PlaceStorage
}

// NewPlaceStorageAdapter ...
func NewPlaceStorageAdapter(txManager *tx.Manager) *PlaceStorageAdapter {
	return &PlaceStorageAdapter{
		placeStorage: NewPlaceStorage(txManager),
	}
}

// GetPlaceByNameAndAddress ...
func (a *PlaceStorageAdapter) GetPlaceByNameAndAddress(ctx context.Context, name, address string) (Place, error) {
	dbPlaces, err := a.placeStorage.Find(ctx, builder.Eq{
		"name":    name,
		"address": address,
	}, "id")
	if err != nil {
		return Place{}, err
	}

	if len(dbPlaces) == 0 {
		return Place{}, sql.ErrNoRows
	}

	return dbPlaces[0], nil
}
