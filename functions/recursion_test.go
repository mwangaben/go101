package functions

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestFact(t *testing.T) {
	t.Run("it tests the recursion", func(t *testing.T) {
		fmt.Printf("The fact %v \n", Fact(1))
		fmt.Printf("The fact %v \n", Fact(2))
		fmt.Printf("The fact %v \n", Fact(3))
		fmt.Printf("The fact %v \n", Fact(4))
		fmt.Printf("The fact %v \n", Fact(10))
	})

	t.Run("it add slices", func(t *testing.T) {
		sums := []int{1, 2, 3, 4, 5}
		adds := []int{100, 200, 300}

		assert.Equal(t, 15, SumSlice(sums))
		assert.Equal(t, 600, SumSlice(adds))
	})

	t.Run("it show the Power on a given number", func(t *testing.T) {
		assert.Equal(t, 2, Power(2, 1))

		assert.Equal(t, 4, Power(2, 2))
		assert.Equal(t, 8, Power(2, 3))
		assert.Equal(t, 25, Power(5, 2))
	})

}
