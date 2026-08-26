package swithes

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestSwitchStatement(t *testing.T) {

	t.Run("it Sort Odd numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}
		expected := []int{1, 3, 5, 7}
		actual := SortNumbers(numbers, "odd")
		assert.Equal(t, expected, actual)
	})

	t.Run("It sort Even numbers", func(t *testing.T) {
		numbers := []int{1, 2, 4, 5, 6, 7}
		expected := []int{2, 4, 6}
		actual := SortNumbers(numbers, "even")

		assert.Equal(t, expected, actual)
	})

	t.Run("It fails to so unknown type", func(t *testing.T) {
		numbers := []int{1, 2, 4, 5, 6, 7}
		var ints []int
		assert.Equal(t, ints, SortNumbers(numbers, "prime"))
	})

	t.Run("It says the type", func(t *testing.T) {
		assert.Equal(t, "bool", GiveType(true))
		assert.Equal(t, "int", GiveType(1))
		assert.Equal(t, "string", GiveType("Name"))
	})

}
