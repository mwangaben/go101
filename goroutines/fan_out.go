package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func FanOutExample() {

	//create channels for tasks
	tasks := make(chan int, 10)

	//load tasks
	for i := 1; i <= 10; i++ {
		tasks <- i
	}

	close(tasks)

	//Fan out : starts three workers
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, tasks, &wg)
	}

	wg.Wait()
	fmt.Println("All tasks completed")

}

func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("Worker %d  processing task %d \n", id, task)
		time.Sleep(100 * time.Millisecond)
	}
}
