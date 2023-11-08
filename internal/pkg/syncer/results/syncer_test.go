package results

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		s := New(Config{})
		assert.NotNil(t, s)
	})
}

func TestSyncer_Enabled(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		s := Syncer{
			enabled: true,
		}
		assert.NotNil(t, s)
		assert.True(t, s.Enabled())
	})
}

func TestSyncer_GetPeriod(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		s := Syncer{
			period: 60 * time.Second,
		}
		assert.Equal(t, 60*time.Second, s.GetPeriod())
	})
}

func TestSyncer_GetName(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		s := Syncer{
			name: "name",
		}
		assert.Equal(t, "name", s.GetName())
	})
}
