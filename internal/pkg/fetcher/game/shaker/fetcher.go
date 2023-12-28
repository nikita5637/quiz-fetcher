package shaker

import (
	"net/http"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage"
)

const (
	fetcherName   = "shaker"
	gamesListPath = "/api/v1/event/published"
	leagueID      = int32(4)
	url           = "https://backend.shakerquiz.ru"
)

// Fetcher ...
type Fetcher struct {
	client               http.Client
	gamesListPath        string
	gameTypeMatchStorage storage.GameTypeMatchStorage
	leagueID             int32
	name                 string
	placeStorage         storage.PlaceStorage
	url                  string
}

// Config ...
type Config struct {
	GameTypeMatchStorage storage.GameTypeMatchStorage
	PlaceStorage         storage.PlaceStorage
}

// New ...
func New(cfg Config) *Fetcher {
	return &Fetcher{
		client:               *http.DefaultClient,
		gamesListPath:        gamesListPath,
		gameTypeMatchStorage: cfg.GameTypeMatchStorage,
		leagueID:             leagueID,
		name:                 fetcherName,
		placeStorage:         cfg.PlaceStorage,
		url:                  url,
	}
}

// GetName ...
func (f *Fetcher) GetName() string {
	return f.name
}

// GetLeagueID ...
func (f *Fetcher) GetLeagueID() int32 {
	return f.leagueID
}
