package main

import (
	"log"
	"sync"
	"time"
)

var cache = make(map[string]string)

func main() {

	var wg sync.WaitGroup

	c := sync.NewCond(&sync.Mutex{})

	wg.Add(1)
	go func() {
		c.L.Lock()
		defer wg.Done()
		for len(cache) == 0 {
			log.Printf("[go-routine] waiting for Elements from cache : %v\n", cache)
			c.Wait()
		}
		log.Printf("[go-routine] Got Elements from cache : %v\n", cache)
		c.L.Unlock()
	}()

	time.Sleep(1 * time.Second)
	log.Println("[main-routine] preparing resources to populate cache")
	time.Sleep(1 * time.Second)
	log.Println("[main-routine] Populating cache")
	c.L.Lock()
	cache["key"] = "ABC12345"
	log.Println("[main-routine] Populated cache")
	log.Println("[main-routine] Giving Signal for waiting go-routine")
	c.Signal() // can call by holding lock
	c.L.Unlock()
	// c.Signal() // can call without having lock
	wg.Wait()
	log.Println("[main-routine] Done!")
}

/*
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/with_cond_by_signal.go
	2021/06/06 11:55:14 [go-routine] waiting for Elements from cache : map[]
	2021/06/06 11:55:15 [main-routine] preparing resources to populate cache
	2021/06/06 11:55:16 [main-routine] Populating cache
	2021/06/06 11:55:16 [main-routine] Populated cache
	2021/06/06 11:55:16 [main-routine] Giving Signal for waiting go-routine
	2021/06/06 11:55:16 [go-routine] Got Elements from cache : map[key:ABC12345]
	2021/06/06 11:55:16 [main-routine] Done!
*/
