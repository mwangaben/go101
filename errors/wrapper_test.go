package errors

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestTest(t *testing.T) {

	t.Run("testing the DatabaseErro ", func(t *testing.T) {

		err := ProcessQuery("SELECT * FROM users")

		assert2.Error(t, err)
		assert2.ErrorContains(t, err, "SELECT * FROM")
	})
}
