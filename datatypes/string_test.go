package datatypes

import (
	"github.com/go-playground/assert/v2"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {

	t.Run("it check the default of string", func(t *testing.T) {
		assert.IsEqual(Name, "")
		assert2.Empty(t, Name)
	})

	t.Run("it checks the string length", func(t *testing.T) {
		Name = "Benedict"
		assert.Equal(t, Name, "Benedict")
		assert.Equal(t, 8, len(Name))
	})

	t.Run("it assigned string by backtick", func(t *testing.T) {
		Message = `Message was sent`

		assert.IsEqual(Message, "Message was sent")
	})
}
