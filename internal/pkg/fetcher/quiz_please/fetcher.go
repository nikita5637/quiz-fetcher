package quiz_please

import (
	"net/http"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/clients"
)

const (
	// FetcherName ...
	FetcherName = "quiz, please!"
	// GamesListPath ...
	GamesListPath = "/schedule"
	// GameInfoPathFormat ...
	GameInfoPathFormat = "/ajax/scope-game?id=%d"
	// URL ...
	URL = "https://spb.quizplease.ru"

	leagueID = 1
)

// GamesFetcher ...
type GamesFetcher struct {
	client                   http.Client
	gameInfoPathFormat       string
	gamesListPath            string
	name                     string
	placesCache              map[string]int32
	registratorServiceClient clients.RegistratorServiceClient
	url                      string
}

// Config ...
type Config struct {
	GameInfoPathFormat       string
	GamesListPath            string
	Name                     string
	RegistratorServiceClient clients.RegistratorServiceClient
	URL                      string
}

// NewGamesFetcher ...
func NewGamesFetcher(cfg Config) *GamesFetcher {
	return &GamesFetcher{
		client:                   *http.DefaultClient,
		gameInfoPathFormat:       cfg.GameInfoPathFormat,
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
