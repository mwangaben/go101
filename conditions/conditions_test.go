package conditions

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestConditionsStatement(t *testing.T) {

	t.Run("It checks the Even number", func(t *testing.T) {
		assert.Equal(t, IsOddNumber(6), false)
		assert.Equal(t, IsOddNumber(3), true)
	})

	t.Run("It checks the Odd number", func(t *testing.T) {
		assert.Equal(t, IsEvenNumber(12), true)
	})

	t.Run("It sort  odd number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}
		expected := []int{1, 3, 5, 7}

		assert.Equal(t, expected, SortOddNumber(numbers))
	})
}
