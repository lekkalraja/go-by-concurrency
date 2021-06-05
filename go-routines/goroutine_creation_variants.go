package main

import (
	"log"
	"time"
)

func main() {
	// Directly Call disply
	display("Calling Display with `main goroutine (Direct Call)`", 2)

	// go-routine function call
	go display("Calling Display with `go display()`", 2)

	// go-routine with function literal (anonymous funciton)
	go func() {
		display("Calling Display with `funciton literal`", 2)
	}()

	// go-routine with function value call

	fn := display
	go fn("Calling Display with `funciton value`", 2)

	// for time-being making main-goroutine to sleep till all child go routines completes
	log.Println("Waiting for Child thread finshing")
	time.Sleep(1 * time.Second)
}

func display(message string, times int) {
	for i := 0; i < times; i++ {
		log.Printf("Messaged sent for Display : %s \n", message)
		time.Sleep(10 * time.Millisecond)
	}
}

/*
	Note : output order varies because concurrent program doesn't preserve order of execution. it's alwasy varies
	OUTPUT:
	=======
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run go-routines/goroutine_creation_variants.go
	2021/06/04 21:19:53 Messaged sent for Display : Calling Display with `main goroutine (Direct Call)`
	2021/06/04 21:19:53 Messaged sent for Display : Calling Display with `main goroutine (Direct Call)`
	2021/06/04 21:19:53 Messaged sent for Display : Calling Display with `funciton value`
	2021/06/04 21:19:53 Messaged sent for Display : Calling Display with `go display()`
	2021/06/04 21:19:53 Messaged sent for Display : Calling Display with `funciton literal`
	2021/06/04 21:19:53 Messaged sent for Display : Calling Display with `go display()`
	2021/06/04 21:19:53 Messaged sent for Display : Calling Display with `funciton value`
	2021/06/04 21:19:53 Messaged sent for Display : Calling Display with `funciton literal`
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run go-routines/goroutine_creation_variants.go
	2021/06/04 21:19:55 Messaged sent for Display : Calling Display with `main goroutine (Direct Call)`
	2021/06/04 21:19:55 Messaged sent for Display : Calling Display with `main goroutine (Direct Call)`
	2021/06/04 21:19:55 Messaged sent for Display : Calling Display with `funciton value`
	2021/06/04 21:19:55 Messaged sent for Display : Calling Display with `funciton literal`
	2021/06/04 21:19:55 Messaged sent for Display : Calling Display with `go display()`
	2021/06/04 21:19:55 Messaged sent for Display : Calling Display with `go display()`
	2021/06/04 21:19:55 Messaged sent for Display : Calling Display with `funciton value`
	2021/06/04 21:19:55 Messaged sent for Display : Calling Display with `funciton literal`
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run go-routines/goroutine_creation_variants.go
	2021/06/04 21:20:00 Messaged sent for Display : Calling Display with `main goroutine (Direct Call)`
	2021/06/04 21:20:00 Messaged sent for Display : Calling Display with `main goroutine (Direct Call)`
	2021/06/04 21:20:00 Messaged sent for Display : Calling Display with `funciton value`
	2021/06/04 21:20:00 Messaged sent for Display : Calling Display with `go display()`
	2021/06/04 21:20:00 Messaged sent for Display : Calling Display with `funciton literal`
	2021/06/04 21:20:00 Messaged sent for Display : Calling Display with `funciton literal`
	2021/06/04 21:20:00 Messaged sent for Display : Calling Display with `funciton value`
	2021/06/04 21:20:00 Messaged sent for Display : Calling Display with `go display()`
*/
