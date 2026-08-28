package functions

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	assert2 "github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func SeqInt() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func multiplier(factor int) func(i int) int {
	return func(i int) int {
		return i * factor
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func makeFibonacci(fn func() int) []int {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, fn())
	}
	return numbers
}

func accumulator() func(n int) int {
	sum := 0
	return func(n int) int {
		sum += n
		return sum
	}
}
func calculator() (func(n int) int, func() int) {
	sum := 0
	add := func(n int) int {
		sum += n
		return sum
	}
	getTotal := func() int {
		return sum
	}

	return add, getTotal

}

func maxOut() func() (int, error) {
	counter := 0
	maxs := 3
	return func() (int, error) {
		if counter <= maxs {
			counter++
			return counter, nil
		} else {
			return counter, fmt.Errorf("sorry you have maxout %v", counter)
		}
	}
}

func createFile(filename string) func(string) error {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return func(text string) error {
		_, err := file.WriteString(text)
		return err
	}
}

func TestClosures(t *testing.T) {

	t.Run("It display the sequence", func(t *testing.T) {
		nextInt := SeqInt()
		more := SeqInt()
		assert.Equal(t, 1, nextInt())
		assert.Equal(t, 2, nextInt())
		assert.Equal(t, 3, nextInt())

		assert.Equal(t, 1, more())
	})

	t.Run("it capture the values", func(t *testing.T) {
		x := 12
		captureX := func() int {
			return x
		}
		assert.Equal(t, 12, captureX())
	})

	t.Run("Make a counter", func(t *testing.T) {
		x := 10
		counter := func() int {
			x++
			return x
		}
		assert.Equal(t, 11, counter())
	})

	t.Run("it make multiplier", func(t *testing.T) {
		double := multiplier(2)
		triple := multiplier(3)
		tenFolds := multiplier(10)

		assert.Equal(t, 10, double(5))
		assert.Equal(t, 15, triple(5))
		assert.Equal(t, 100, tenFolds(10))
	})

	t.Run("it create fibonacci Generator", func(t *testing.T) {
		expected := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
		assert.Equal(t, expected, makeFibonacci(fibonacci()))
	})

	t.Run("it accumulates addition", func(t *testing.T) {

		add := accumulator()
		assert.Equal(t, 10, add(10))
		assert.Equal(t, 21, add(11))
		assert.Equal(t, 50, add(29))
	})

	t.Run("it calculate and get total", func(t *testing.T) {
		add, getTotal := calculator()

		assert.Equal(t, 10, add(10))
		assert.Equal(t, 30, add(20))
		assert.Equal(t, 30, getTotal())

		assert.Equal(t, 25, add(-5))
		assert.Equal(t, 25, getTotal())
	})

	t.Run("it check the maxOut closure", func(t *testing.T) {
		m := maxOut()
		c, err := m()
		c2, err := m()
		_, err = m()

		assert.Equal(t, 1, c)
		assert.Equal(t, nil, err)

		assert.Equal(t, 2, c2)
		assert.Equal(t, nil, err)

		_, err = m()
		_, err = m()

		assert.IsEqual("sorry you have maxout 4", err)

	})

	t.Run("It create text files ", func(t *testing.T) {
		writer := createFile("logs.text")
		writer("Hello")
		writer(" world")

		content, err := os.ReadFile("logs.text")
		assert.Equal(t, nil, err)

		fmt.Printf("content is : %v", string(content))

		assert.Equal(t, "Hello world", string(content))
		assert2.FileExists(t, "logs.text")
	})

	t.Run("it tests Cache closure", func(t *testing.T) {
		cached := Memoize(ExpensiveSquare)

		val, fromCached := cached(5)
		assert.Equal(t, 25, val)
		assert.Equal(t, false, fromCached)

		val2, fromCached2 := cached(5)
		assert.Equal(t, 25, val2)
		assert.Equal(t, true, fromCached2)

	})

	t.Run("it tests the rate limiter", func(t *testing.T) {
		limiter := RateLimiter(3, time.Second)

		assert.Equal(t, true, limiter())
		assert.Equal(t, true, limiter())
		assert.Equal(t, true, limiter())
		assert.Equal(t, false, limiter())
	})

	t.Run("it tests the StateMAchine", func(t *testing.T) {
		transition := StateMachine()

		assert.Equal(t, StateIdle, transition(StateRunning))
		assert.Equal(t, StateRunning, transition(StateStopped))
		assert.Equal(t, StateStopped, transition(StateRunning))
	})

}
