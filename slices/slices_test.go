package slices

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func AddElements(elements []int) int {
	var sum int

	for _, n := range elements {
		sum = sum + n
	}

	return sum
}

func TestSlices(t *testing.T) {

	t.Run("testing the slice making", func(t *testing.T) {
		var s []string
		assert.Equal(t, 0, len(s))
		assert.Equal(t, nil, s)
	})

	t.Run("Making non zero a sclice", func(t *testing.T) {
		s := make([]string, 3)
		assert.Equal(t, 3, len(s))
		s = append([]string{"benny"}, s...)
		assert.Equal(t, []string{"benny", "", "", ""}, s)
		assert.Equal(t, 4, len(s))

	})

	t.Run("it makes", func(t *testing.T) {
		s := make([]int, 3)
		s[0] = 100
		s[1] = 200
		s[2] = 300
		fmt.Printf("the slice before  is %v \n", s)
		s = append(s, 1)
		s = append(s, 2)
		s = append(s, 3)
		assert.Equal(t, 1, s[3])
		fmt.Printf("the slice  is %v \n", s)

		assert.Equal(t, 606, AddElements(s))
	})

	t.Run("It adds the elements", func(t *testing.T) {
		e := []int{100, 200, 300, 1, 2, 3}
		assert.Equal(t, 606, AddElements(e))
	})
}
