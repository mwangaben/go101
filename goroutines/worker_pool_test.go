package goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestWorkerPool(t *testing.T) {
	t.Run("worker pool", func(t *testing.T) {
		const numJobs = 10
		const numWorkers = 3

		jobs := make(chan Job, numJobs)
		results := make(chan Result, numWorkers)

		var wg sync.WaitGroup

		//Start Workers
		for i := 1; i <= numWorkers; i++ {
			wg.Add(1)
			go WorkerPool(i, jobs, results, &wg)
		}

		//Sends Jobs
		for j := 1; j <= numJobs; j++ {
			jobs <- Job{
				ID:     j,
				Number: j,
			}
		}
		close(jobs)

		//Close result when all worker are done
		go func() {
			wg.Wait()
			close(results)
		}()

		for result := range results {
			fmt.Printf("result: Job %d -> %d\n ", result.JobID, result.Value)
		}
	})
}
