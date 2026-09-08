package goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicRace() {

	var wg sync.WaitGroup
	var ops atomic.Int64

	// Start 10 goroutines
	for range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				ops.Add(1)
			}

		}()
	}

	wg.Wait()
	fmt.Printf("the Load is %v", ops.Load())
}
