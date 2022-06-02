package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_divide(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		resp, err := divide(2, 1)
		assert.Equal(t, 2, resp)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		resp, err := divide(2, 0)
		assert.Equal(t, 0, resp)
		assert.Error(t, err)
	})
}
