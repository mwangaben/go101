package goroutines

func Sum(numbers []int, result chan int) {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	result <- sum //Send the result yo the channel
}
