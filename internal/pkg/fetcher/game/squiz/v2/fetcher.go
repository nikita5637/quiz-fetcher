//go:generate mockery --case underscore --name GameTypeMatchStorage --with-expecter

package squiz

import (
	"context"
	"net/http"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
)

const (
	// FetcherName ...
	FetcherName = "squiz"
	// GamesListPath ...
	GamesListPath = "/api/getproductslist/"
	// URL ...
	URL = "https://store.tildacdn.com"

	gamesListQuery = "?storepartuid=111979372401&getparts=true&getoptions=true&slice=1&size=100"
	leagueID       = int32(leaguepb.LeagueID_SQUIZ)
	maxPlayers     = 8
)

// GameTypeMatchStorage ...
type GameTypeMatchStorage interface {
	GetGameTypeByDescription(ctx context.Context, description string) (int32, error)
}

// GamesFetcher ...
type GamesFetcher struct {
	client               http.Client
	gamesListPath        string
	gameTypeMatchStorage GameTypeMatchStorage
	name                 string
	placeStorage         storage.PlaceStorage
	url                  string
}

// Config ...
type Config struct {
	GamesListPath        string
	GameTypeMatchStorage GameTypeMatchStorage
	Name                 string
	PlaceStorage         storage.PlaceStorage
	URL                  string
}

// NewGamesFetcher ...
func NewGamesFetcher(cfg Config) *GamesFetcher {
	return &GamesFetcher{
		client:               *http.DefaultClient,
		gamesListPath:        cfg.GamesListPath,
		gameTypeMatchStorage: cfg.GameTypeMatchStorage,
		name:                 cfg.Name,
		placeStorage:         cfg.PlaceStorage,
		url:                  cfg.URL,
	}
}

// GetName ...
func (f *GamesFetcher) GetName() string {
	return f.name
}

// GetLeagueID ...
func (f *GamesFetcher) GetLeagueID() int32 {
	return leagueID
}
