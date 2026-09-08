package goroutines

import (
	"fmt"
	"time"
)

func BurstRateLimiter() {

	burstyLimiter := make(chan time.Time, 3)

	for range 3 {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(800 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for br := 1; br <= 5; br++ {
		burstyRequests <- br
	}

	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request ", req, time.Now())
	}
}

func SimpleRateLimiter() {
	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for request := range requests {
		<-limiter
		fmt.Println("request", request, time.Now())
	}
}
