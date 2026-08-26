package variables

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestVariables(t *testing.T) {

	t.Run("it test the variables", func(t *testing.T) {
		assert.Equal(t, FirstName, "Mylan")
		assert.Equal(t, Age, 3)
	})
}
