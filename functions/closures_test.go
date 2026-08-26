package functions

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func SeqInt() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func TestClosures(t *testing.T) {

	t.Run("It display the sequence", func(t *testing.T) {
		nextInt := SeqInt()
		assert.Equal(t, 1, nextInt())
		assert.Equal(t, 2, nextInt())
		assert.Equal(t, 3, nextInt())
	})
}
