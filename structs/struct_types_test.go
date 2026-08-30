package structs

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestStructTypes(t *testing.T) {

	c := NewCircle(21)
	ca := c.Area()
	cp := c.Perimeter()

	sq := NewSquare(10)
	sqA := sq.Area()
	sqP := sq.Perimeter()

	t.Run("it show the shape area and perimter", func(t *testing.T) {
		assert2.Equal(t, 1385.442360233099, ca)
		assert2.Equal(t, 131.94689145077132, cp)

		assert2.Equal(t, float64(100), sqA)
		assert2.Equal(t, float64(40), sqP)

		expectCircleInfo := "Area:1385.442360233099 Perimeter:131.94689145077132"
		expectSquareinfo := "Area:100 Perimeter:40"

		assert2.Equal(t, expectCircleInfo, ShareInfo(c))
		assert2.Equal(t, expectSquareinfo, ShareInfo(sq))

	})
}
