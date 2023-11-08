//go:generate mockery --case underscore --name Syncer --with-expecter
//go:generate mockery --case underscore --name SyncLogsFacade --with-expecter

package synchronizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		s := New(Config{})
		assert.NotNil(t, s)
	})
}
