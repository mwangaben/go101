package functions

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func Add(a, b, c int) int {
	return a + b + c
}

func SimpleAdd(a int, b int) int {
	return a + b
}

func SumAndMultiply(a, b int) (int, int) {
	sum := a + b
	products := a * b
	return sum, products
}

func SumAll(numbers ...int) int {
	var total int
	for _, n := range numbers {
		total += n
	}
	return total
}

func TestFunctions(t *testing.T) {

	t.Run("It test simple function with single return", func(t *testing.T) {
		assert.Equal(t, 600, Add(100, 200, 300))
	})

	t.Run("It add simple numbers", func(t *testing.T) {
		assert.Equal(t, 14, SimpleAdd(12, 2))
	})

	t.Run("It add and multiply numbers", func(t *testing.T) {
		sumResult, productResult := SumAndMultiply(3, 6)

		assert.Equal(t, 9, sumResult)
		assert.Equal(t, 18, productResult)
	})

	t.Run("Sum All numbers", func(t *testing.T) {

		assert.Equal(t, 6, SumAll(1, 2, 3))
		assert.Equal(t, 40, SumAll(10, 20, 10))
	})

	t.Run("it accepts slices", func(t *testing.T) {
		numbers := []int{1, 2, 10}
		assert.Equal(t, 13, SumAll(numbers...))
	})

}
