package quiz_please

import (
	"net/http"

	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
)

const (
	fetcherName    = "quiz, please!"
	gameResultPath = "/game-page?id=%d"
	leagueID       = int32(leaguepb.LeagueID_QUIZ_PLEASE)
	url            = "https://spb.quizplease.ru"
)

// Fetcher ...
type Fetcher struct {
	client         http.Client
	gameResultPath string
	leagueID       int32
	name           string
	team           string
	url            string
}

// Config ...
type Config struct {
	Team string
}

// New ...
func New(cfg Config) *Fetcher {
	return &Fetcher{
		client:         *http.DefaultClient,
		gameResultPath: gameResultPath,
		leagueID:       leagueID,
		name:           fetcherName,
		team:           cfg.Team,
		url:            url,
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
