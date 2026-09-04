package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSayHello(t *testing.T) {
	t.Run("it returns the go goroutine", func(t *testing.T) {

		// Start a goroutine
		go SayHello()

		// Start a goroutine with anonymous function
		go func() {
			fmt.Println("Hello from anonymous goroutine!")
		}()

		// Wait for goroutines to finish (bad practice - use sync.WaitGroup)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Main function done")
	})
}

func TestWorker(t *testing.T) {
	t.Run("tests the worker goroutine", func(t *testing.T) {
		var wg sync.WaitGroup

		for i := 1; i <= 5; i++ {
			wg.Add(1) // Increment the counter
			go Worker(i, &wg)
		}

		wg.Wait() // Wait for all goroutines to finish
		fmt.Println("All workers completed")
	})
}
