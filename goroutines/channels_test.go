package goroutines

import (
	"fmt"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("it test the Sum", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		result := make(chan int)

		go Sum(numbers[:len(numbers)/2], result)

		sum1 := <-result
		assert2.Equal(t, 15, sum1)

		go Sum(numbers[len(numbers)/2:], result)

		fmt.Printf("the numbers : %v \n", numbers[len(numbers)/2:])
		sum2 := <-result

		assert2.Equal(t, 40, sum2)
	})
}
