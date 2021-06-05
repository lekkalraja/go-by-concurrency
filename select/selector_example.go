package main

import (
	"log"
	"time"
)

func main() {
	ch1 := make(chan string)

	ch2 := make(chan string)

	go doWork(ch1, 3*time.Second)
	go doWork(ch2, 2*time.Second)

	// should loop with resp to number of go-routines,
	// so that whatever comes we can process data without missing.
	// Note: If we don't have loop, then we only process the fastest response
	for i := 0; i < 2; i++ {
		select {
		case data := <-ch1:
			log.Printf("[Channel - 1] Got Data : %s\n", data)
		case data := <-ch2:
			log.Printf("[Channel - 2] Got Data : %s\n", data)
		}
	}
}

func doWork(ch chan<- string, delay time.Duration) {
	time.Sleep(delay)
	ch <- delay.String() + " Delay"
}
