package goroutines

import (
	"fmt"
	"time"
)

func NonBlockingChannels() {
	messages := make(chan string)
	signals := make(chan bool)

	go func() {
		for msg := range messages {
			fmt.Println("mesa ", msg)
		}
	}()

	select {
	case msg := <-messages:
		fmt.Println("message received ", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "Hi"
	select {
	case messages <- msg:
		fmt.Println("message received ", msg)
	default:
		fmt.Println("no message received")
	}
	// Give the goroutine time to process
	time.Sleep(300 * time.Millisecond)
	select {
	case msg2 := <-messages:
		fmt.Println("message received ", msg2)
	case sig := <-signals:
		fmt.Println("received signals ", sig)
	default:
		fmt.Println("no message received")
	}
}
