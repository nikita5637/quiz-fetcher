package squiz

import (
	"net/http"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/clients"
)

const (
	// FetcherName ...
	FetcherName = "squiz"
	// GamesListPath ...
	GamesListPath = "/#schedule"
	// URL ...
	URL = "https://spb.squiz.ru"

	leagueID = 2
)

// GamesFetcher ...
type GamesFetcher struct {
	client                   http.Client
	gamesListPath            string
	name                     string
	placesCache              map[string]int32
	registratorServiceClient clients.RegistratorServiceClient
	url                      string
}

// Config ...
type Config struct {
	GamesListPath            string
	Name                     string
	RegistratorServiceClient clients.RegistratorServiceClient
	URL                      string
}

// NewGamesFetcher ...
func NewGamesFetcher(cfg Config) *GamesFetcher {
	return &GamesFetcher{
		client:                   *http.DefaultClient,
		gamesListPath:            cfg.GamesListPath,
		name:                     cfg.Name,
		placesCache:              make(map[string]int32, 0),
		registratorServiceClient: cfg.RegistratorServiceClient,
		url:                      cfg.URL,
	}
}

// GetName ...
func (w *GamesFetcher) GetName() string {
	return w.name
}
