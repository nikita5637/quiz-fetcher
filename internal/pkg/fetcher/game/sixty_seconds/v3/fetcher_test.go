package sixty_seconds

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// TestFetcher ...
func TestFetcher(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(GetSuite))
}
