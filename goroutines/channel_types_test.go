package goroutines

import (
	"fmt"
	"testing"
	"time"
)

/**
Bidirectional channel
*/

func WorkerChannel(ch chan int) {
	ch <- 12
	<-ch
}
func TestBi_directionChannel(t *testing.T) {

	bch := make(chan int)

	bch <- 12
	res := <-bch
	fmt.Println("the result: ", res)
}

/**
Send Only channel (chan<-)
*/

func Producer(ch chan<- int) {
	ch <- 42 // Allowed ✅
	//val := <-ch  // Not Allowed ❌
}

func TestSendOnlyChannel(t *testing.T) {

	//usage:
	ch := make(chan int)
	go Producer(ch) // Pass bi-directional channel
	result := <-ch  // Main reads the result

	fmt.Println("Send Only :", result)
}

/***
Receive-only channel(<-chan)
*/

func Consumer(ch <-chan int) {
	val := <-ch // ✅ Allowed
	// ch <- 42        // ❌ COMPILE ERROR: can't send
	fmt.Println(val)
}

// Usage:
func TestReceive_Only_Channel(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 42 // Write from goroutine
	}()
	Consumer(ch) // Pass bi-directional channel
}

/**
Directional channels as function parameters
*/

// Producer: only write to the channel

func ProducerTwo(out chan<- int) {
	for i := 1; i < 5; i++ {
		out <- i
	}
	close(out)
}

// Consumer: only reads from channels
func ConsumerTwo(in <-chan int) {
	for val := range in {
		fmt.Println("received: ", val)
	}
}

// usage
func TestDirectional_As_Function_parameters(t *testing.T) {
	chTwo := make(chan int)
	go ProducerTwo(chTwo)
	ConsumerTwo(chTwo)
}

/**
Returning directional  channels
*/
//return a receive-only channel

func GetData(url string) <-chan []byte {
	ch := make(chan []byte)
	go func() {
		time.Sleep(2 * time.Second)
		data := []byte("Fetched Data from " + url)
		ch <- data
		close(ch)
	}()

	return ch
}

//Usage

func TestReturnDirectionalChannel(t *testing.T) {
	data := GetData("https://dt.umbijani.com")
	val := <-data

	fmt.Println("the returning: ", val)

}

/**
Channel to channels (advanced)
*/
//Channel that sends/receives channels

func TestChannel_Of_Channel(t *testing.T) {

	chOfCh := make(chan chan int)
	go func() {
		innerCh := make(chan int)
		chOfCh <- innerCh //Send a channel through a channel
	}()

	innerCh := <-chOfCh

	go func() {
		val := <-innerCh
		fmt.Println("the inner Channel", val)

	}()
	innerCh <- 42

}

// TODO Why use directional channels?

// 1. Documentation: Shows intent clearly
func Fetch(url string) <-chan []byte {
	ch := make(chan []byte)
	go func() {
		time.Sleep(2 * time.Second)
		data := []byte("fetched data is " + url)
		ch <- data
	}()
	return ch
}                             // Clearly returns a read-only channel
func Process(data <-chan int) {} // Clearly expects to read from channel

// 2. Safety: Prevents accidental operations
type Data struct{}

func SaveToDB(in <-chan Data) { // Can't accidentally write to input
	// in <- Data{}  // Compile error - prevents bugs!
}

// 3. Interface design: Enforces contracts
type ProducerThree interface {
	GetData() <-chan int // Returns read-only channel
}

func TestChannel_To_Channel(t *testing.T) {
	chOfCh := make(chan chan int)

	go func() {
		innerCh := make(chan int)
		chOfCh <- innerCh

		//Wait for response
		val := <-innerCh //Read value from main
		fmt.Println("Goroutine received", val)

		//Send response back
		innerCh <- val * 2

	}()

	innerCh := <-chOfCh //Receive channel

	//Send request
	innerCh <- 12

}
