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
	GamesListPath = "/#schedule"
	// URL ...
	URL = "https://spb.squiz.ru"

	leagueID   = int32(leaguepb.LeagueID_SQUIZ)
	maxPlayers = 8
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
func (w *GamesFetcher) GetName() string {
	return w.name
}
