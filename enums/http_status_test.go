package enums

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestHttpStatus_IsClientError(t *testing.T) {
	t.Run("it test the status options", func(t *testing.T) {
		status := NotFound
		assert2.Equal(t, false, status.IsSuccess())
		assert2.True(t, status.IsClientError())
		assert2.False(t, status.IsServerError())

		status = OK
		assert2.Equal(t, "200 Ok", status.String())
	})
}
