//go:generate mockery --case underscore --name GameTypeMatchStorage --with-expecter

package squiz

import (
	"context"
	"net/http"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/fetcher/clients"
	pkgmodel "github.com/nikita5637/quiz-registrator-api/pkg/model"
)

const (
	// FetcherName ...
	FetcherName = "squiz"
	// GamesListPath ...
	GamesListPath = "/#schedule"
	// URL ...
	URL = "https://spb.squiz.ru"

	leagueID   = pkgmodel.LeagueSquiz
	maxPlayers = 8
)

// GameTypeMatchStorage ...
type GameTypeMatchStorage interface {
	GetGameTypeByDescription(ctx context.Context, description string) (int32, error)
}

// GamesFetcher ...
type GamesFetcher struct {
	client                   http.Client
	gamesListPath            string
	gameTypeMatchStorage     GameTypeMatchStorage
	name                     string
	placesCache              map[string]int32
	registratorServiceClient clients.RegistratorServiceClient
	url                      string
}

// Config ...
type Config struct {
	GamesListPath            string
	GameTypeMatchStorage     GameTypeMatchStorage
	Name                     string
	RegistratorServiceClient clients.RegistratorServiceClient
	URL                      string
}

// NewGamesFetcher ...
func NewGamesFetcher(cfg Config) *GamesFetcher {
	return &GamesFetcher{
		client:                   *http.DefaultClient,
		gamesListPath:            cfg.GamesListPath,
		gameTypeMatchStorage:     cfg.GameTypeMatchStorage,
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
