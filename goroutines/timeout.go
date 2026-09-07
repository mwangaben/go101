package goroutines

import (
	"fmt"
	"time"
)

func SlowOperation(result chan string) {
	time.Sleep(3 * time.Second)
	result <- "Operation Completed"
}

func TimeoutChannel() {

	c1 := make(chan string, 1)

	go func() {
		fmt.Println("Processing... task one")
		time.Sleep(1 * time.Second) // simulate a process
		c1 <- "✅ chan 1"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case <-time.After(2 * time.Second):
		fmt.Println("time out")
	}

	c2 := make(chan string, 1)

	go func() {
		fmt.Println("Processing.. task two")
		time.Sleep(3 * time.Second)
		c2 <- "✅ chan 2"
	}()

	select {
	case msg2 := <-c2:
		fmt.Println(msg2)
	case <-time.After(2 * time.Second):
		fmt.Println("❌ time two task two")
	}

}
