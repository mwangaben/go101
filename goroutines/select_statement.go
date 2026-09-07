package goroutines

import (
	"fmt"
	"time"
)

func AsyncTask(id int, delay time.Duration, result chan string) {
	time.Sleep(delay)

	result <- fmt.Sprintf("Task %d completed after %v", id, delay)
}

func CallAsyncTask() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	//Start multiple tasks with difference delays
	go AsyncTask(1, 2*time.Second, ch1)
	go AsyncTask(2, 1*time.Second, ch2)
	go AsyncTask(3, 3*time.Second, ch3)

	//	 Wait for all result using select

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case msg3 := <-ch3:
			fmt.Println(msg3)
		}
	}

}

func SelectForRange() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()

	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)

		}

	}

}
