package quiz_please

import (
	"net/http"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
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

	leagueID = int32(leaguepb.LeagueID_QUIZ_PLEASE)
)

// Fetcher ...
type Fetcher struct {
	client             http.Client
	gameInfoPathFormat string
	gamesListPath      string
	name               string
	placeStorage       storage.PlaceStorage
	url                string
}

// Config ...
type Config struct {
	GameInfoPathFormat string
	GamesListPath      string
	Name               string
	PlaceStorage       storage.PlaceStorage
	URL                string
}

// New ...
func New(cfg Config) *Fetcher {
	return &Fetcher{
		client:             *http.DefaultClient,
		gameInfoPathFormat: cfg.GameInfoPathFormat,
		gamesListPath:      cfg.GamesListPath,
		name:               cfg.Name,
		placeStorage:       cfg.PlaceStorage,
		url:                cfg.URL,
	}
}

// GetName ...
func (f *Fetcher) GetName() string {
	return f.name
}

// GetLeagueID ...
func (f *Fetcher) GetLeagueID() int32 {
	return leagueID
}
