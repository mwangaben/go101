package goroutines

import (
	"fmt"
	"sync"
)

func PipeLineExample() {

	//1. generateNumbers
	numbers := generateNumbers(10)

	//2. square them using 3 workers
	squared := squareNumbers(numbers, 3)

	//3. Sum the result
	sum := sumNumbers(squared)
	fmt.Printf("Sum of squares : %d \n", sum)
}

func generateNumbers(count int) <-chan int {
	out := make(chan int)

	go func() {
		for i := 1; i <= count; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func squareNumbers(in <-chan int, workers int) <-chan int {
	out := make(chan int)

	//Fan - Out: Start workers
	var wg sync.WaitGroup

	for i := 1; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for num := range in {
				out <- num * num
			}
		}()
	}

	//Fa-In: close output when all workers done
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func sumNumbers(in <-chan int) int {
	sum := 0
	for num := range in {
		fmt.Printf(" the chan %v ", num)
		sum += num
	}
	return sum
}
