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

	// if we don't get response immediatly and don't want to be block our code,
	// then have default statement, so that even if our response are not ready we can process default response.
	for i := 0; i < 3; i++ {
		select {
		case data := <-ch1:
			log.Printf("[Channel - 1] Got Data : %s\n", data)
		case data := <-ch2:
			log.Printf("[Channel - 2] Got Data : %s\n", data)
		default:
			log.Printf("[Default] Getting Data ")
		}
		time.Sleep(1 * time.Second)
	}
}

func doWork(ch chan<- string, delay time.Duration) {
	time.Sleep(delay)
	ch <- delay.String() + " Delay"
}

/*
	OUTPUT:
	=======
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run select/selector_non_blocking.go
	2021/06/05 17:33:16 [Default] Getting Data
	2021/06/05 17:33:17 [Default] Getting Data
	2021/06/05 17:33:18 [Channel - 2] Got Data : 2s Delay
*/
