package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

func generate(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, num := range nums {
			select {
			case out <- num:
				log.Printf("[g-goroutine] publishing number : %d \n", num)
			case <-done:
				return
			}
		}
	}()

	return out
}

func square(done <-chan struct{}, id int, inputC <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for in := range inputC {
			select {
			case out <- in * in:
				log.Printf("[s-%d] Squaring number : %d \n", id, in)
			case <-done:
				return
			}
		}
	}()

	return out
}

func merge(done <-chan struct{}, chans ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(chans))

	merger := func(ch <-chan int) {
		defer wg.Done()

		for ele := range ch {
			select {
			case out <- ele:
				log.Printf("[m-goroutine] Merging number : %d \n", ele)
			case <-done:
				return
			}
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

	done := make(chan struct{})
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	genC := generate(done, numbers...)

	square1C := square(done, 1, genC)
	square2C := square(done, 2, genC)

	mergeC := merge(done, square1C, square2C)

	log.Printf("**********************************    I need only one Element : %d \n", <-mergeC)
	close(done) // To make ALl other Go-Routines terminate
	time.Sleep(10 * time.Millisecond)
	log.Printf("**********************************    Active Go Routines : %d\n", runtime.NumGoroutine())
}

/*
Even Main GoRoutine not processing elemetns, still other goroutines are processing (to overcome this send done signal)
======================================================================================================================

raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run -race concurrency_patterns/real_pipeline.go
2021/06/06 16:20:43 [g-goroutine] publishing number : 1
2021/06/06 16:20:43 [g-goroutine] publishing number : 2
2021/06/06 16:20:43 [m-goroutine] Merging number : 1
2021/06/06 16:20:43 [s-1] Squaring number : 1
2021/06/06 16:20:43 [s-1] Squaring number : 3
2021/06/06 16:20:43 [s-2] Squaring number : 2
2021/06/06 16:20:43 **********************************    I need only one Element : 1
2021/06/06 16:20:43 [g-goroutine] publishing number : 3
2021/06/06 16:20:43 [g-goroutine] publishing number : 4
2021/06/06 16:20:43 [g-goroutine] publishing number : 5
2021/06/06 16:20:43 **********************************    Active Go Routines : 7

raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run -race concurrency_patterns/real_pipeline.go
2021/06/06 16:20:44 [g-goroutine] publishing number : 1
2021/06/06 16:20:44 [g-goroutine] publishing number : 2
2021/06/06 16:20:44 [m-goroutine] Merging number : 1
2021/06/06 16:20:44 [s-2] Squaring number : 2
2021/06/06 16:20:44 **********************************    I need only one Element : 1
2021/06/06 16:20:44 [s-1] Squaring number : 1
2021/06/06 16:20:44 [g-goroutine] publishing number : 3
2021/06/06 16:20:44 [g-goroutine] publishing number : 4
2021/06/06 16:20:44 [s-1] Squaring number : 4
2021/06/06 16:20:44 [g-goroutine] publishing number : 5
2021/06/06 16:20:44 **********************************    Active Go Routines : 7


After Sending Done Signal
=========================
2021/06/06 16:22:05 [g-goroutine] publishing number : 1
2021/06/06 16:22:05 [g-goroutine] publishing number : 2
2021/06/06 16:22:05 [m-goroutine] Merging number : 1
2021/06/06 16:22:05 **********************************    I need only one Element : 1
2021/06/06 16:22:05 [s-1] Squaring number : 1
2021/06/06 16:22:05 [s-2] Squaring number : 2
2021/06/06 16:22:05 **********************************    Active Go Routines : 1 (SHOULD BE MAIN-GOROUTINE)

raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run -race concurrency_patterns/real_pipeline.go
2021/06/06 16:22:06 [g-goroutine] publishing number : 1
2021/06/06 16:22:06 [g-goroutine] publishing number : 2
2021/06/06 16:22:06 [m-goroutine] Merging number : 4
2021/06/06 16:22:06 [s-2] Squaring number : 2
2021/06/06 16:22:06 [s-2] Squaring number : 3
2021/06/06 16:22:06 [g-goroutine] publishing number : 3
2021/06/06 16:22:06 [g-goroutine] publishing number : 4
2021/06/06 16:22:06 **********************************    I need only one Element : 4
2021/06/06 16:22:06 [s-1] Squaring number : 1
2021/06/06 16:22:06 **********************************    Active Go Routines : 1 (SHOULD BE MAIN-GOROUTINE)

raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run -race concurrency_patterns/real_pipeline.go
2021/06/06 16:22:08 [g-goroutine] publishing number : 1
2021/06/06 16:22:08 [s-1] Squaring number : 1
2021/06/06 16:22:08 [m-goroutine] Merging number : 1
2021/06/06 16:22:08 **********************************    I need only one Element : 1
2021/06/06 16:22:08 [g-goroutine] publishing number : 2
2021/06/06 16:22:08 [g-goroutine] publishing number : 3
2021/06/06 16:22:08 **********************************    Active Go Routines : 1 (SHOULD BE MAIN-GOROUTINE)
*/
