package main

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	/*var flag bool = true
	var mu sync.Mutex */

	var once sync.Once

	load := func() {
		log.Println("Initializaiton tasks are loading!!!!")
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			/*mu.Lock()
			defer mu.Unlock()
			if flag {
				flag = false
				load()
			}*/
			once.Do(load)
		}()
	}

	wg.Done()
	log.Println("sync.Once Done!!!")
}
