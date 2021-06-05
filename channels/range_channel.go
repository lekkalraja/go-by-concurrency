package main

import "log"

func main() {

	ch := make(chan int)

	go func(elem int) {
		defer close(ch)

		for i := 0; i < elem; i++ {
			log.Printf("[Go-Routine] Sending data to channel: %d\n", i)
			ch <- i
		}

	}(5)

	for value := range ch {
		log.Printf("[main-Routine] Receiving data from channel: %d\n", value)
	}

	log.Printf("[main-Routine] Range-over channel completed")
}

/*
	Note : Even though it is un-buffered channel we are observing in output that it has buffer of capacity 1

	OUTPUT:
	=======
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run channels/range_channel.go
	2021/06/05 11:44:47 [Go-Routine] Sending data to channel: 0
	2021/06/05 11:44:47 [Go-Routine] Sending data to channel: 1
	2021/06/05 11:44:47 [main-Routine] Receiving data from channel: 0
	2021/06/05 11:44:47 [main-Routine] Receiving data from channel: 1
	2021/06/05 11:44:47 [Go-Routine] Sending data to channel: 2
	2021/06/05 11:44:47 [Go-Routine] Sending data to channel: 3
	2021/06/05 11:44:47 [main-Routine] Receiving data from channel: 2
	2021/06/05 11:44:47 [main-Routine] Receiving data from channel: 3
	2021/06/05 11:44:47 [Go-Routine] Sending data to channel: 4
	2021/06/05 11:44:47 [main-Routine] Receiving data from channel: 4
	2021/06/05 11:44:47 [main-Routine] Range-over channel completed
*/
