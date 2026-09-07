package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestSlowOperation(t *testing.T) {

	t.Run("time delayed", func(t *testing.T) {
		ch := make(chan string)

		go SlowOperation(ch)

		select {
		case res := <-ch:
			fmt.Println("Success: ", res)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout: operation took too long")
		}

	})

	t.Run("ticker", func(t *testing.T) {
		// 2. Periodic ticker
		ticker := time.NewTicker(1 * time.Second)
		for {
			<-ticker.C // Read from the ticker's channel
			fmt.Println("Tick!")
		}
	})

	t.Run("Run after function ", func(t *testing.T) {
		// 3. Delay without blocking
		time.AfterFunc(3*time.Second, func() {
			fmt.Println("This runs after 3 seconds")
		})
	})

	t.Run("it timeout the task two ", func(t *testing.T) {
		TimeoutChannel()
	})

	t.Run("it test non blocking channel", func(t *testing.T) {
		NonBlockingChannels()
	})

}
