package sixty_seconds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		f := New(Config{})
		assert.NotNil(t, f)
	})
}

func TestFetcher_GetName(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		f := &Fetcher{
			name: "name",
		}
		assert.Equal(t, "name", f.GetName())
	})
}

func TestFetcher_GetLeagueID(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		f := &Fetcher{
			leagueID: 1,
		}
		assert.Equal(t, int32(1), f.GetLeagueID())
	})
}
