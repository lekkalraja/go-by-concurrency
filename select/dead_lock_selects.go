package main

import "log"

func main() {

	/*
		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [select (no cases)]:
	*/
	// select {}

	var ch chan string

	/*
		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan receive (nil chan)]:
	*/
	/*select {
	case v := <-ch:
		log.Printf("Received value from ch : %s \n", v)
	}*/

	/*
		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan send (nil chan)]:
	*/
	select {
	case ch <- "Data":
		log.Println("Send data to ch")
	}
}
