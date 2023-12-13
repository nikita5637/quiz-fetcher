package sixty_seconds

import (
	"net/http"
)

const (
	fetcherName    = "60 seconds"
	gameResultPath = "/quizgames/ajax/get_game_results/"
	leagueID       = int32(3)
	url            = "https://club60sec.ru"
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
