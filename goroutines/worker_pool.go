package goroutines

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID     int
	Number int
}

type Result struct {
	JobID int
	Value int
}

func WorkerPool(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker  %d processing job %d\n", id, job.ID)
		time.Sleep(time.Millisecond * 500) //Simulate work
		results <- Result{
			JobID: job.ID,
			Value: job.Number,
		}
	}
}
