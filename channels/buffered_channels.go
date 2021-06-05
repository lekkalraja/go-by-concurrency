package main

import "log"

func main() {

	ch := make(chan int, 5)

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
	OUTPUT:
	=======
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run channels/buffered_channels.go
	2021/06/05 11:47:49 [Go-Routine] Sending data to channel: 0
	2021/06/05 11:47:49 [Go-Routine] Sending data to channel: 1
	2021/06/05 11:47:49 [Go-Routine] Sending data to channel: 2
	2021/06/05 11:47:49 [Go-Routine] Sending data to channel: 3
	2021/06/05 11:47:49 [Go-Routine] Sending data to channel: 4
	2021/06/05 11:47:49 [main-Routine] Receiving data from channel: 0
	2021/06/05 11:47:49 [main-Routine] Receiving data from channel: 1
	2021/06/05 11:47:49 [main-Routine] Receiving data from channel: 2
	2021/06/05 11:47:49 [main-Routine] Receiving data from channel: 3
	2021/06/05 11:47:49 [main-Routine] Receiving data from channel: 4
	2021/06/05 11:47:49 [main-Routine] Range-over channel completed
*/
