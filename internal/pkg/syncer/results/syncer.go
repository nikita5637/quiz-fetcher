//go:generate mockery --case underscore --name Fetcher --with-expecter
//go:generate mockery --case underscore --name GameServiceClient --with-expecter
//go:generate mockery --case underscore --name GameResultManagerServiceClient --with-expecter

package results

import (
	"context"
	"time"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	gameresultmanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Fetcher ...
type Fetcher interface {
	GetGameResult(ctx context.Context, externalID int32) (model.GameResult, error)
	GetName() string
	GetLeagueID() int32
}

// GameServiceClient ...
type GameServiceClient interface {
	GetGame(ctx context.Context, in *gamepb.GetGameRequest, opts ...grpc.CallOption) (*gamepb.Game, error)
}

// GameResultManagerServiceClient ...
type GameResultManagerServiceClient interface {
	ListGameResults(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*gameresultmanagerpb.ListGameResultsResponse, error)
	PatchGameResult(ctx context.Context, in *gameresultmanagerpb.PatchGameResultRequest, opts ...grpc.CallOption) (*gameresultmanagerpb.GameResult, error)
}

// Syncer ...
type Syncer struct {
	disabledResultsFetchers        []string
	enabled                        bool
	fetchers                       []Fetcher
	gameServiceClient              GameServiceClient
	gameResultManagerServiceClient GameResultManagerServiceClient
	name                           string
	period                         time.Duration
}

// Config ...
type Config struct {
	DisabledResultsFetchers        []string
	Enabled                        bool
	Fetchers                       []Fetcher
	GameServiceClient              GameServiceClient
	GameResultManagerServiceClient GameResultManagerServiceClient
	Name                           string
	Period                         time.Duration
}

// New ...
func New(cfg Config) *Syncer {
	return &Syncer{
		disabledResultsFetchers:        cfg.DisabledResultsFetchers,
		enabled:                        cfg.Enabled,
		fetchers:                       cfg.Fetchers,
		gameServiceClient:              cfg.GameServiceClient,
		gameResultManagerServiceClient: cfg.GameResultManagerServiceClient,
		name:                           cfg.Name,
		period:                         cfg.Period,
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
