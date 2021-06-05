package main

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var index int

	wg.Add(1)

	go func() {
		defer wg.Done()
		// Clousure
		index++
		log.Printf("Incremented index value : %d \n", index)
	}()

	wg.Wait()
	if index == 0 {
		log.Printf("The value of index doesn't change i.e. : %d \n", index)
	}
}

/*
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run go-routines/wait_group.go
	2021/06/04 21:56:15 Incremented index value : 1
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run go-routines/wait_group.go
	2021/06/04 21:56:16 Incremented index value : 1
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run go-routines/wait_group.go
	2021/06/04 21:56:17 Incremented index value : 1
*/
