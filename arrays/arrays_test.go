package arrays

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestArray(t *testing.T) {

	t.Run("the array testing", func(t *testing.T) {
		var a [5]int

		assert.Equal(t, 5, len(a))
		a[1] = 2
		a[0] = 100
		a[4] = 12
		assert.Equal(t, 2, a[1])
		assert.Equal(t, 100, a[0])
		assert.Equal(t, 12, a[4])
	})

	t.Run("assigned and use", func(t *testing.T) {
		b := [5]int{1, 2, 3, 4, 5}

		assert.Equal(t, 2, b[1])
	})

	t.Run("assigned and use and count", func(t *testing.T) {
		b := [...]int{1, 2, 5, 4, 5, 6}

		assert.Equal(t, 5, b[2])
		assert.Equal(t, 6, len(b))
	})
}
