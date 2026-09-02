package enums

import (
	"fmt"
	"testing"

	assert2 "github.com/stretchr/testify/assert"
)

func TestBasicEnum(t *testing.T) {

	t.Run("testing trr basir Enum", func(t *testing.T) {
		s := Active

		assert2.Equal(t, 1, int(s))
		assert2.Equal(t, Active, s)
		fmt.Printf("value %v \n", s)
	})
}
