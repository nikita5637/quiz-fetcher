package game

import (
	"context"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
)

// Fetcher ...
type Fetcher interface {
	GetGamesList(ctx context.Context) ([]model.Game, error)
	GetName() string
}
