package quiz_please

import (
	"net/http"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/clients"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage"
	pkgmodel "github.com/nikita5637/quiz-registrator-api/pkg/model"
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

	leagueID = pkgmodel.LeagueQuizPlease
)

// Fetcher ...
type Fetcher struct {
	client                   http.Client
	gameInfoPathFormat       string
	gamesListPath            string
	name                     string
	placeStorage             storage.PlaceStorage
	registratorServiceClient clients.RegistratorServiceClient
	url                      string
}

// Config ...
type Config struct {
	GameInfoPathFormat       string
	GamesListPath            string
	Name                     string
	PlaceStorage             storage.PlaceStorage
	RegistratorServiceClient clients.RegistratorServiceClient
	URL                      string
}

// New ...
func New(cfg Config) *Fetcher {
	return &Fetcher{
		client:                   *http.DefaultClient,
		gameInfoPathFormat:       cfg.GameInfoPathFormat,
		gamesListPath:            cfg.GamesListPath,
		name:                     cfg.Name,
		placeStorage:             cfg.PlaceStorage,
		registratorServiceClient: cfg.RegistratorServiceClient,
		url:                      cfg.URL,
	}
}

// GetName ...
func (f *Fetcher) GetName() string {
	return f.name
}
