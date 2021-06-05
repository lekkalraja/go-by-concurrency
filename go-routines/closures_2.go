package main

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for index := range []int{1, 2, 3} {
		wg.Add(1)

		go func(counter int) {
			wg.Done()
			log.Printf("[Go-Routine] The Value of index : %d \n", counter)
		}(index)
	}

	wg.Wait()
	log.Printf("[main-routine] Exiting............")
}
