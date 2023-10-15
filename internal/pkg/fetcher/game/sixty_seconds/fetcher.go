package sixty_seconds

import (
	"net/http"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
)

const (
	// FetcherName ...
	FetcherName = "60 seconds"
	// OpenLeagueGamesListPath ...
	OpenLeagueGamesListPath = "/league/117/"
	// FirstLeagueGamesListPath ...
	FirstLeagueGamesListPath = "/league/118/"
	// URL ...
	URL = "https://60sec.online"

	leagueID = int32(leaguepb.LeagueID_SIXTY_SECONDS)
)

// GamesFetcher ...
type GamesFetcher struct {
	client                   http.Client
	openLeagueGamesListPath  string
	firstLeagueGamesListPath string
	name                     string
	placeStorage             storage.PlaceStorage
	url                      string
}

// Config ...
type Config struct {
	OpenLeagueGamesListPath  string
	FirstLeagueGamesListPath string
	Name                     string
	PlaceStorage             storage.PlaceStorage
	URL                      string
}

// NewGamesFetcher ...
func NewGamesFetcher(cfg Config) *GamesFetcher {
	return &GamesFetcher{
		client:                   *http.DefaultClient,
		openLeagueGamesListPath:  cfg.OpenLeagueGamesListPath,
		firstLeagueGamesListPath: cfg.FirstLeagueGamesListPath,
		name:                     cfg.Name,
		placeStorage:             cfg.PlaceStorage,
		url:                      cfg.URL,
	}
}

// GetName ...
func (f *GamesFetcher) GetName() string {
	return f.name
}

// GetLeagueID ...
func (f *GamesFetcher) GetLeagueID() int32 {
	return leagueID
}
