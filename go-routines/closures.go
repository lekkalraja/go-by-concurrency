package main

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incr := func(wg *sync.WaitGroup) {
		var index int
		wg.Add(1)
		go func() {
			defer wg.Done()
			index++
			log.Printf("[Child-Rouinte] Value of index : %d \n", index)
		}()
		log.Printf("[Main-Rouinte] Returning from the incr function : %d \n", index)
	}
	incr(&wg)
	wg.Wait()
	log.Println("[Main-Rouinte] Exiting....")
}
