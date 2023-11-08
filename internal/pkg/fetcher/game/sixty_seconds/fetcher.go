package sixty_seconds

import (
	"net/http"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
)

const (
	fetcherName              = "60 seconds"
	firstLeagueGamesListPath = "/league/118/"
	leagueID                 = int32(leaguepb.LeagueID_SIXTY_SECONDS)
	openLeagueGamesListPath  = "/league/117/"
	url                      = "https://60sec.online"
)

// Fetcher ...
type Fetcher struct {
	client                   http.Client
	openLeagueGamesListPath  string
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
		openLeagueGamesListPath:  openLeagueGamesListPath,
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
