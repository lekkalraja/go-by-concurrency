package main

import "log"

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)
	go write(ch1)
	log.Printf("[main-routine] Received data from ch1 : %s \n", read(ch1))
	go write(ch1)
	go swap(ch1, ch2)
	log.Printf("[main-routine] Received data from ch2 : %s \n", read(ch2))
	log.Println("[main-routine] Finished Channel Directions")
}

func write(ch chan<- string) {
	log.Println("[go-routine] Writing data to channel")
	ch <- "Writing Data"
	/* value := <-ch // ERROR: Invalid Operation: cannot receivefor send-only channel
	   fmt.Println(value) */
}

func read(ch <-chan string) string {
	value, ok := <-ch
	// ch <- "Invalid Operation" // Error : Invalid operation: cannot send to receive-only type <-chan string
	if ok {
		return value
	}
	return "Channel is Closed"
}

func swap(readChan <-chan string, writeChan chan<- string) {
	log.Println("[go-routine] Swapping data")
	value := read(readChan)
	writeChan <- value
}

/*
	OUTPUT:
	=======
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run channels/channel_directions.go
	2021/06/05 12:03:19 [go-routine] Writing data to channel
	2021/06/05 12:03:19 [main-routine] Received data from ch1 : Writing Data
	2021/06/05 12:03:19 [go-routine] Swapping data
	2021/06/05 12:03:19 [go-routine] Writing data to channel
	2021/06/05 12:03:19 [main-routine] Received data from ch2 : Writing Data
	2021/06/05 12:03:19 [main-routine] Finished Channel Directions
*/
