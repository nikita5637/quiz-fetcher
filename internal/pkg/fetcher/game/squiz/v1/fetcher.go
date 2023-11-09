//go:generate mockery --case underscore --name GameTypeMatchStorage --with-expecter

package squiz

import (
	"context"
	"net/http"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage"
)

const (
	fetcherName   = "squiz"
	gamesListPath = "/#schedule"
	leagueID      = int32(2)
	maxPlayers    = 8
	url           = "https://spb.squiz.ru"
)

// GameTypeMatchStorage ...
type GameTypeMatchStorage interface {
	GetGameTypeByDescription(ctx context.Context, description string) (int32, error)
}

// Fetcher ...
type Fetcher struct {
	client               http.Client
	gamesListPath        string
	gameTypeMatchStorage GameTypeMatchStorage
	name                 string
	placeStorage         storage.PlaceStorage
	url                  string
}

// Config ...
type Config struct {
	GameTypeMatchStorage GameTypeMatchStorage
	PlaceStorage         storage.PlaceStorage
}

// New ...
func New(cfg Config) *Fetcher {
	return &Fetcher{
		client:               *http.DefaultClient,
		gamesListPath:        gamesListPath,
		gameTypeMatchStorage: cfg.GameTypeMatchStorage,
		name:                 fetcherName,
		placeStorage:         cfg.PlaceStorage,
		url:                  url,
	}
}

// GetName ...
func (f *Fetcher) GetName() string {
	return f.name
}
