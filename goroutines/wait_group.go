package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func workerWait(id int) {

	fmt.Printf("star processing job: %v \n", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("job %v  finished \n", id)
}

func ProcessJob(count int) {
	var wg sync.WaitGroup

	for i := 1; i <= count; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			workerWait(workerID)
		}(i)
	}

	wg.Wait()

}
