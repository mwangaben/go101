package datatypes

import (
	"github.com/go-playground/assert/v2"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestChangeName(t *testing.T) {

	t.Run("It set Value to variable", func(t *testing.T) {
		Age = 3
		assert.IsEqual(Age, 3)
		Age = 12
		assert.IsEqual(Age, 12)
	})

	t.Run("it set bytes number", func(t *testing.T) {
		Alphabet = 'A'
		assert.IsEqual(Alphabet, 'A')

		Alphabet = 'S'
		assert.IsEqual(Alphabet, 'S')
	})

	t.Run("it set rune", func(t *testing.T) {
		Alpha = 'D'

		assert2.Equal(t, Alpha, 'D')
	})
}
