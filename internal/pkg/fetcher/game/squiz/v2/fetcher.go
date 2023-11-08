//go:generate mockery --case underscore --name GameTypeMatchStorage --with-expecter

package squiz

import (
	"context"
	"net/http"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
)

const (
	fetcherName    = "squiz"
	gamesListPath  = "/api/getproductslist/"
	gamesListQuery = "?storepartuid=111979372401&getparts=true&getoptions=true&slice=1&size=100"
	leagueID       = int32(leaguepb.LeagueID_SQUIZ)
	maxPlayers     = 8
	url            = "https://store.tildacdn.com"
)

// GameTypeMatchStorage ...
type GameTypeMatchStorage interface {
	GetGameTypeByDescription(ctx context.Context, description string) (int32, error)
}

// Fetcher ...
type Fetcher struct {
	client               http.Client
	gamesListPath        string
	gameTypeMatchStorage GameTypeMatchStorage
	leagueID             int32
	name                 string
	placeStorage         storage.PlaceStorage
	url                  string
}

// Config ...
type Config struct {
	GameTypeMatchStorage GameTypeMatchStorage
	PlaceStorage         storage.PlaceStorage
}

// New ...
func New(cfg Config) *Fetcher {
	return &Fetcher{
		client:               *http.DefaultClient,
		gamesListPath:        gamesListPath,
		gameTypeMatchStorage: cfg.GameTypeMatchStorage,
		leagueID:             leagueID,
		name:                 fetcherName,
		placeStorage:         cfg.PlaceStorage,
		url:                  url,
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
