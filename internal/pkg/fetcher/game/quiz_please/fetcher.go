package quiz_please

import (
	"net/http"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage"
)

const (
	fetcherName        = "quiz, please!"
	gameInfoPathFormat = "/ajax/scope-game?id=%d"
	gamesListPath      = "/schedule"
	leagueID           = int32(1)
	url                = "https://spb.quizplease.ru"
)

// Fetcher ...
type Fetcher struct {
	client             http.Client
	gameInfoPathFormat string
	gamesListPath      string
	leagueID           int32
	name               string
	placeStorage       storage.PlaceStorage
	url                string
}

// Config ...
type Config struct {
	PlaceStorage storage.PlaceStorage
}

// New ...
func New(cfg Config) *Fetcher {
	return &Fetcher{
		client:             *http.DefaultClient,
		gameInfoPathFormat: gameInfoPathFormat,
		gamesListPath:      gamesListPath,
		leagueID:           leagueID,
		name:               fetcherName,
		placeStorage:       cfg.PlaceStorage,
		url:                url,
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
