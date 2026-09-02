package interators

import (
	"fmt"
	"testing"
)

func TestBasicIterator(t *testing.T) {

	t.Run("it test the basic iterator", func(t *testing.T) {

		nums := []int{1, 2, 3, 4, 5}

		for i, v := range nums {
			fmt.Printf("index %v : value: %v \n", i, v)
		}
	})

	t.Run("it test yield", func(t *testing.T) {
		for v := range Count(1, 5) {
			fmt.Printf("the count: %v \n", v)
		}
	})

	t.Run("it test yield with index", func(t *testing.T) {
		for idx, val := range CountWithIndex(1, 5) {
			fmt.Printf("the index %v : value : %v \n", idx, val)
		}
	})
}
