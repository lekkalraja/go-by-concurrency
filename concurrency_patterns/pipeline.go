package main

import "log"

// Build a pipeline : generator -> square -> print

// generator - converts a list of integers to a channel
func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, num := range nums {
			log.Printf("[Generator-goroutine] Publishing %d on channel\n", num)
			out <- num
		}
	}()

	return out
}

// receive on inbound channel, square the number and then output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for num := range in {
			log.Printf("[Square] Squaring %d, and pushing on to channel\n", num)
			out <- num * num
		}
	}()

	return out
}

// set up the pipeline
// run the last stage of pipeline, receive the values from square stage
// print each on , until channel is closed
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	/*genC := generator(numbers...)

	squareC := square(genC)

	for out := range squareC {
		log.Printf("[Main] Got Number : %d \n", out)
	}*/

	for out := range square(generator(numbers...)) {
		log.Printf("[Main] Got Number : %d \n", out)
	}
}

/*
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run -race concurrency_patterns/pipeline.go
	2021/06/06 15:10:49 [Generator-goroutine] Publishing 1 on channel
	2021/06/06 15:10:49 [Generator-goroutine] Publishing 2 on channel
	2021/06/06 15:10:49 [Square] Squaring 1, and pushing on to channel
	2021/06/06 15:10:49 [Square] Squaring 2, and pushing on to channel
	2021/06/06 15:10:49 [Generator-goroutine] Publishing 3 on channel
	2021/06/06 15:10:49 [Main] Got Number : 1
	2021/06/06 15:10:49 [Main] Got Number : 4
	2021/06/06 15:10:49 [Square] Squaring 3, and pushing on to channel
	2021/06/06 15:10:49 [Main] Got Number : 9
	2021/06/06 15:10:49 [Generator-goroutine] Publishing 4 on channel
	2021/06/06 15:10:49 [Generator-goroutine] Publishing 5 on channel
	2021/06/06 15:10:49 [Square] Squaring 4, and pushing on to channel
	2021/06/06 15:10:49 [Square] Squaring 5, and pushing on to channel
	2021/06/06 15:10:49 [Generator-goroutine] Publishing 6 on channel
	2021/06/06 15:10:49 [Main] Got Number : 16
	2021/06/06 15:10:49 [Main] Got Number : 25
	2021/06/06 15:10:49 [Square] Squaring 6, and pushing on to channel
	2021/06/06 15:10:49 [Main] Got Number : 36
	2021/06/06 15:10:49 [Generator-goroutine] Publishing 7 on channel
	2021/06/06 15:10:49 [Generator-goroutine] Publishing 8 on channel
	2021/06/06 15:10:49 [Square] Squaring 7, and pushing on to channel
	2021/06/06 15:10:49 [Square] Squaring 8, and pushing on to channel
	2021/06/06 15:10:49 [Main] Got Number : 49
	2021/06/06 15:10:49 [Main] Got Number : 64
*/
