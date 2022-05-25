package main

import "fmt"

// Potentially expensive function
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	// Channels are used to send data between goroutines
	answer := make(chan int)

	// You spawn a goroutine with the `go` keyword
	go func() {
		// Place the return value into the channel with an arrow
		answer <- fib(40)
	}()

	// The main function can continue processing while
	fmt.Println("Calculating fib(40)...")

	// Take the answer out of the channel, when ready.
	// will block on empty/unbuffered channel
	fmt.Println(<-answer)

	// make a buffered channel of strings that can buffer up to two string
	queue := make(chan string, 2)

	// we send the two values to the channel
	queue <- "one"
	queue <- "two"

	// close the channel
	// if not closing the channel, the range will loop forever
	close(queue)

	// loop on channel until all data is received
	// range is an iterator mechanism, think JS forof forin
	for elem := range queue {
		fmt.Println(elem)
	}
}
