package enums

import (
	"fmt"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicEnum(t *testing.T) {

	t.Run("testing the basic Enum", func(t *testing.T) {
		s := Active

		assert2.Equal(t, 1, int(s))
		assert2.Equal(t, Active, s)
		fmt.Printf("value %v \n", s)
	})
}
