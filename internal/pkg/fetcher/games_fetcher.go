package fetcher

import (
	"context"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
)

// GamesFetcher ...
type GamesFetcher interface {
	GetGamesList(ctx context.Context) ([]model.Game, error)
	GetName() string
}
