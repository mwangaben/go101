package goroutines

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func AddWithMutex() {
	c := Container{
		mu:       sync.Mutex{},
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		doIncrement("a", 1000)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		doIncrement("a", 10000)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		doIncrement("b", 1000)
	}()

	wg.Wait()

	fmt.Printf("the counters are : %v ", c.counters)
}
