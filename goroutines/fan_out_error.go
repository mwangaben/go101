package goroutines

import (
	"fmt"
	"sync"
	"time"
)

type ResultErr struct {
	Value int
	Error error
}

func FanOutWithErrors() {

	//Input Data
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//Fan-Out: Process in parallel
	results := processWithErrors(numbers, 4)

	// Collect results

	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error: %v\n ", result.Error)
		} else {
			fmt.Printf("Result: %v\n ", result.Value)
		}

	}
}

func processWithErrors(numbers []int, workers int) <-chan ResultErr {
	resultChan := make(chan ResultErr, len(numbers))
	var wg sync.WaitGroup

	// Fan-out; Start  workers
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for _, num := range numbers {
				// process with potential error
				result, err := process(num)
				resultChan <- ResultErr{Value: result, Error: err}
			}
		}(w)
	}

	//Fan-In :  Close  when  done
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	return resultChan
}

func process(num int) (int, error) {
	time.Sleep(100 * time.Millisecond)

	//Simulate  error  on odd   numbers
	if num%2 == 1 {
		return 0, fmt.Errorf("error processing %d \n", num)
	}
	return num * num, nil
}
