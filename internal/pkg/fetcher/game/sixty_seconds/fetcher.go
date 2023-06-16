package sixty_seconds

import (
	"net/http"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/clients"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage"
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

	leagueID = 3
)

// GamesFetcher ...
type GamesFetcher struct {
	client                   http.Client
	openLeagueGamesListPath  string
	firstLeagueGamesListPath string
	name                     string
	placeStorage             storage.PlaceStorage
	registratorServiceClient clients.RegistratorServiceClient
	url                      string
}

// Config ...
type Config struct {
	OpenLeagueGamesListPath  string
	FirstLeagueGamesListPath string
	Name                     string
	PlaceStorage             storage.PlaceStorage
	RegistratorServiceClient clients.RegistratorServiceClient
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
		registratorServiceClient: cfg.RegistratorServiceClient,
		url:                      cfg.URL,
	}
}

// GetName ...
func (w *GamesFetcher) GetName() string {
	return w.name
}
