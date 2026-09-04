package errors

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomErr(t *testing.T) {

	t.Run("it test custom err", func(t *testing.T) {
		result, err := DivideWithCustomError(3, 0)

		assert2.Equal(t, float64(0), result)

		expectedError := DivisionErr{
			Dividend: 3,
			Divisor:  0,
			Message:  "can not divide by zero",
		}
		assert2.ErrorIs(t, err, expectedError)
	})
}
