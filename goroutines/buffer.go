package goroutines

import "fmt"

func BufferChannel() {
	// Create a buffered channel with capacity 3
	ch := make(chan string, 3)

	//Start goroutine that sends message
	go func() {
		ch <- "message 1"
		ch <- "message 2"
		ch <- "message 3"

		close(ch)

	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
