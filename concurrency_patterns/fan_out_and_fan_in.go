package main

import (
	"log"
	"sync"
)

func generate(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, num := range nums {
			log.Printf("[g-goroutine] publishing number : %d \n", num)
			out <- num
		}
	}()

	return out
}

func square(id int, inputC <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for in := range inputC {
			log.Printf("[s-%d] Squaring number : %d \n", id, in)
			out <- in * in
		}
	}()

	return out
}

func merge(chans ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(chans))

	merger := func(ch <-chan int) {
		defer wg.Done()

		for ele := range ch {
			log.Printf("[m-goroutine] Merging number : %d \n", ele)
			out <- ele
		}
	}

	go func() {
		for _, c := range chans {
			go merger(c)
		}
	}()

	// Check for all go-routines once finishes close the channel
	go func() {
		wg.Wait()
		defer close(out)
	}()

	return out
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	genC := generate(numbers...)

	square1C := square(1, genC)
	square2C := square(2, genC)

	mergeC := merge(square1C, square2C)

	for out := range mergeC {
		log.Printf("[m-goroutine] Printing Response : %d\n", out)
	}
}

/*
raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run -race concurrency_patterns/fan_out_and_fan_in.go
2021/06/06 15:40:02 [g-goroutine] publishing number : 1
2021/06/06 15:40:02 [g-goroutine] publishing number : 2
2021/06/06 15:40:02 [g-goroutine] publishing number : 3
2021/06/06 15:40:02 [s-1] Squaring number : 2
2021/06/06 15:40:02 [s-1] Squaring number : 3
2021/06/06 15:40:02 [g-goroutine] publishing number : 4
2021/06/06 15:40:02 [m-goroutine] Merging number : 4
2021/06/06 15:40:02 [m-goroutine] Merging number : 9
2021/06/06 15:40:02 [s-2] Squaring number : 1
2021/06/06 15:40:02 [m-goroutine] Merging number : 1
2021/06/06 15:40:02 [m-goroutine] Printing Response : 4
2021/06/06 15:40:02 [m-goroutine] Printing Response : 9
2021/06/06 15:40:02 [m-goroutine] Printing Response : 1
2021/06/06 15:40:02 [g-goroutine] publishing number : 5
2021/06/06 15:40:02 [g-goroutine] publishing number : 6
2021/06/06 15:40:02 [s-2] Squaring number : 5
2021/06/06 15:40:02 [s-2] Squaring number : 6
2021/06/06 15:40:02 [g-goroutine] publishing number : 7
2021/06/06 15:40:02 [s-1] Squaring number : 4
2021/06/06 15:40:02 [s-1] Squaring number : 7
2021/06/06 15:40:02 [g-goroutine] publishing number : 8
2021/06/06 15:40:02 [m-goroutine] Merging number : 25
2021/06/06 15:40:02 [m-goroutine] Merging number : 36
2021/06/06 15:40:02 [m-goroutine] Printing Response : 25
2021/06/06 15:40:02 [m-goroutine] Printing Response : 36
2021/06/06 15:40:02 [m-goroutine] Merging number : 16
2021/06/06 15:40:02 [m-goroutine] Merging number : 49
2021/06/06 15:40:02 [s-2] Squaring number : 8
2021/06/06 15:40:02 [m-goroutine] Merging number : 64
2021/06/06 15:40:02 [m-goroutine] Printing Response : 16
2021/06/06 15:40:02 [m-goroutine] Printing Response : 49
2021/06/06 15:40:02 [m-goroutine] Printing Response : 64
*/
