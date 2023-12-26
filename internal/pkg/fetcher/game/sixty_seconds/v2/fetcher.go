package sixty_seconds

import (
	"net/http"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage"
)

const (
	fetcherName              = "60 seconds"
	firstLeagueGamesListPath = "/quizgames/league/118/"
	leagueID                 = int32(3)
	url                      = "https://club60sec.ru"
)

// Fetcher ...
type Fetcher struct {
	client                   http.Client
	firstLeagueGamesListPath string
	leagueID                 int32
	name                     string
	placeStorage             storage.PlaceStorage
	url                      string
}

// Config ...
type Config struct {
	PlaceStorage storage.PlaceStorage
}

// New ...
func New(cfg Config) *Fetcher {
	return &Fetcher{
		client:                   *http.DefaultClient,
		firstLeagueGamesListPath: firstLeagueGamesListPath,
		leagueID:                 leagueID,
		name:                     fetcherName,
		placeStorage:             cfg.PlaceStorage,
		url:                      url,
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
