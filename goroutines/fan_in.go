package goroutines

import (
	"fmt"
	"sync"
	"time"
)

// Fan-In:Collect result from multiple workers
func FanInExample() {
	//Create channels  for results
	results := make(chan int)

	// Fan-Out: Start 3 workers
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerWithResult(i, results, &wg)
	}

	// Fan in: collectd all the result
	go func() {
		wg.Wait()
		close(results) // close when all workers are done
	}()

	//receives all results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}

func workerWithResult(id int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	//Simulate work
	time.Sleep(time.Duration(id*100) * time.Millisecond)

	//send result
	results <- id * 10

}
