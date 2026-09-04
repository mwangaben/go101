package errors

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestErrors(t *testing.T) {

	t.Run("it tests the Add function", func(t *testing.T) {

		add, err := Add(2)

		assert2.Equal(t, -1, add)
		assert2.Error(t, err)
		assert2.ErrorContains(t, err, "can't work")
	})

	t.Run("it tests MakeTea can not make more than two cups ata once", func(t *testing.T) {
		result, err := MakeTea(3, Medium)

		assert2.Equal(t, 0, result)
		assert2.ErrorContains(t, err, "you can not make")
	})

	t.Run("it tests MakeTea can not make veryHot tea due to power is at the moment", func(t *testing.T) {
		result, err := MakeTea(2, VeryHot)

		assert2.Equal(t, 0, result)
		assert2.Error(t, err)
		assert2.ErrorContains(t, err, "no enough power")
	})

	t.Run("it makes tea when conditions are met", func(t *testing.T) {
		result, err := MakeTea(2, Hot)

		assert2.Equal(t, 1, result)
		assert2.NoError(t, err)
	})
}
