package main

import "log"

func main() {

	// Owner of a channel
	getChan := func() <-chan string {
		ch := make(chan string) // Creating a channel

		go func() {
			defer close(ch) // closing a channel
			log.Printf("[go-routine] Writing data to Chan")
			ch <- "Getting data from owner" // writing data to channel
		}()

		log.Printf("[main-routine] Returning from the Get Chan")
		return ch // giving read-only channel
	}

	consumer := func(ch <-chan string) {
		value, ok := <-ch

		if ok {
			log.Printf("[main-routine] Getting data from chan : %v\n", value)
		}
	}

	c := getChan()

	consumer(c)
}

/*
	OUTPUT:
	=======
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run channels/ownership_channel.go
	2021/06/05 12:23:28 [main-routine] Returning from the Get Chan
	2021/06/05 12:23:28 [go-routine] Writing data to Chan
	2021/06/05 12:23:28 [main-routine] Getting data from chan : Getting data from owner
*/
