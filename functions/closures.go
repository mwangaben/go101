package functions

import (
	"fmt"
	"time"
)

func Memoize(fn func(int) int) func(int) (int, bool) {
	cache := make(map[int]int)

	return func(x int) (int, bool) {
		if val, exists := cache[x]; exists {
			fmt.Printf("in cached %v \n", x)
			return val, true
		}
		fmt.Printf("miss cached %v \n", x)

		result := fn(x)
		cache[x] = result
		return result, false
	}

}
func ExpensiveSquare(x int) int {
	// Simulate expensive computation
	time.Sleep(100 * time.Millisecond)
	return x * x
}

func RateLimiter(limit int, duration time.Duration) func() bool {
	count := 0
	reset := time.Now().Add(duration)
	return func() bool {
		if time.Now().After(reset) {
			count = 0
			reset = time.Now().Add(duration)
		}
		if count < limit {
			count++
			return true
		}
		return false
	}
}

type State string

const (
	StateIdle    State = "idle"
	StateRunning State = "running"
	StateStopped State = "Stopped"
)

func StateMachine() func(State) State {
	current := StateIdle

	return func(newState State) State {
		old := current
		fmt.Printf("Transtion: %s -> %s \n", old, newState)
		current = newState
		return old
	}
}
