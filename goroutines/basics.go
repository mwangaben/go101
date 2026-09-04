package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func SayHello() {
	fmt.Println("Hello from goroutine!")
}

func Worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() //Signal that this goroutine is done
	fmt.Printf("worker %d starting..... \n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d finished \n", id)
}
