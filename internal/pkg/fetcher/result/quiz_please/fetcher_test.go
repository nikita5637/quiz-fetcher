package quiz_please

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
	f := Fetcher{
		name: "name",
	}
	got := f.GetName()
	assert.Equal(t, "name", got)
}

func TestFetcher_GetLeagueID(t *testing.T) {
	f := Fetcher{
		leagueID: 1,
	}
	got := f.GetLeagueID()
	assert.Equal(t, int32(1), got)
}
