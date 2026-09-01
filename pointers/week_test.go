package pointers

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestWeekDaysAndWeekEndEnums(t *testing.T) {

	t.Run("It checks the week days", func(t *testing.T) {

		today := Tuesday

		assert2.Equal(t, false, today.IsWeekEnd())
		assert2.Equal(t, true, today.IsWeekDay())

	})
}
