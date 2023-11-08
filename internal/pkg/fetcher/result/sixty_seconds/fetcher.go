package sixty_seconds

import (
	"net/http"

	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
)

const (
	fetcherName    = "60 seconds"
	gameResultPath = "/get_game_results/?csrfmiddlewaretoken=gpR7OPhHMBehi8efDiGSLK6ICFkoL9WwKhOUsKRr8EUymyfhZ8gMEzqtAQQvA4Nf&game_id=%d&key=1000"
	leagueID       = int32(leaguepb.LeagueID_SIXTY_SECONDS)
	url            = "https://60sec.online"
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
