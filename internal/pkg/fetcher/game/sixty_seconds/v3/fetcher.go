package sixty_seconds

import (
	"net/http"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage"
)

const (
	fetcherName = "60 seconds"
	leagueID    = int32(3)

	// SchedulePath ...
	SchedulePath = "/quizgames/schedule/71"
	// URL ...
	URL = "https://club60sec.ru"
)

// Fetcher ...
type Fetcher struct {
	client                 http.Client
	leagueID               int32
	name                   string
	needToFetchOpenLeague  bool
	needToFetchFirstLeague bool
	placeStorage           storage.PlaceStorage

	schedulePath string
	url          string
}

// Config ...
type Config struct {
	PlaceStorage storage.PlaceStorage

	NeedToFetchOpenLeague  bool
	NeedToFetchFirstLeague bool
	SchedulePath           string
	URL                    string
}

// New ...
func New(cfg Config) *Fetcher {
	return &Fetcher{
		client:                 *http.DefaultClient,
		schedulePath:           cfg.SchedulePath,
		leagueID:               leagueID,
		needToFetchOpenLeague:  cfg.NeedToFetchOpenLeague,
		needToFetchFirstLeague: cfg.NeedToFetchFirstLeague,
		name:                   fetcherName,
		placeStorage:           cfg.PlaceStorage,
		url:                    cfg.URL,
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
