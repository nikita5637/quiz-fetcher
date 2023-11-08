//go:generate mockery --case underscore --name Fetcher --with-expecter
//go:generate mockery --case underscore --name GameServiceClient --with-expecter

package games

import (
	"context"
	"time"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"google.golang.org/grpc"
)

// Fetcher ...
type Fetcher interface {
	GetGamesList(ctx context.Context) ([]model.Game, error)
	GetName() string
	GetLeagueID() int32
}

// GameServiceClient ...
type GameServiceClient interface {
	CreateGame(ctx context.Context, in *gamepb.CreateGameRequest, opts ...grpc.CallOption) (*gamepb.Game, error)
	PatchGame(ctx context.Context, in *gamepb.PatchGameRequest, opts ...grpc.CallOption) (*gamepb.Game, error)
	SearchGamesByLeagueID(ctx context.Context, in *gamepb.SearchGamesByLeagueIDRequest, opts ...grpc.CallOption) (*gamepb.SearchGamesByLeagueIDResponse, error)
}

// Syncer ...
type Syncer struct {
	disabledGamesFetchers []string
	enabled               bool
	fetchers              []Fetcher
	gameServiceClient     GameServiceClient
	name                  string
	period                time.Duration
}

// Config ...
type Config struct {
	DisabledGamesFetchers []string
	Enabled               bool
	Fetchers              []Fetcher
	GameServiceClient     GameServiceClient
	Name                  string
	Period                time.Duration
}

// New ...
func New(cfg Config) *Syncer {
	return &Syncer{
		disabledGamesFetchers: cfg.DisabledGamesFetchers,
		enabled:               cfg.Enabled,
		fetchers:              cfg.Fetchers,
		gameServiceClient:     cfg.GameServiceClient,
		name:                  cfg.Name,
		period:                cfg.Period,
	}
}

// Enabled ...
func (s *Syncer) Enabled() bool {
	return s.enabled
}

// GetPeriod ...
func (s *Syncer) GetPeriod() time.Duration {
	return s.period
}

// GetName ...
func (s *Syncer) GetName() string {
	return s.name
}
