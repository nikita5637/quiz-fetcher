package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlaceStorageAdapter(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got := NewPlaceStorageAdapter(nil)
		assert.NotNil(t, got)
	})
}
