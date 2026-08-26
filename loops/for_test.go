package loops

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForLoop(t *testing.T) {

	t.Run("it check if numbers were passed", func(t *testing.T) {
		actual := LoopOne()
		expected := []int{1, 2, 3, 4}
		assert.Equal(t, expected, actual)
	})

	t.Run("Single line for loop ", func(t *testing.T) {
		actual := LoopTwo()
		expected := []int{1, 2, 3, 4}

		assert.Equal(t, expected, actual)
	})

	t.Run("It uses range to loop through", func(t *testing.T) {
		actual := LoopThree()
		expected := []int{0, 1, 2, 3}
		assert.Equal(t, expected, actual)
	})

	t.Run("Show alphabet of the word", func(t *testing.T) {
		actual := ShowAlphabet("Benny")
		expected := []rune{'B', 'e', 'n', 'n', 'y'}
		assert.Equal(t, expected, actual)
	})

}
