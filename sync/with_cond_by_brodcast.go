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

	wg.Add(2)
	go func() {
		c.L.Lock()
		defer wg.Done()
		for len(cache) < 1 {
			log.Printf("[go-routine-1] waiting for Elements from cache : %v\n", cache)
			c.Wait()
		}
		log.Printf("[go-routine-1] Got Elements from cache : %v\n", cache)
		c.L.Unlock()
	}()

	go func() {
		c.L.Lock()
		defer wg.Done()
		for len(cache) < 2 {
			log.Printf("[go-routine-2] waiting for Elements from cache : %v\n", cache)
			c.Wait()
		}
		log.Printf("[go-routine-2] Got more than one Element from cache : %v\n", cache)
		c.L.Unlock()
	}()

	time.Sleep(1 * time.Second)
	log.Println("[main-routine] preparing resources to populate cache")
	time.Sleep(1 * time.Second)
	log.Println("[main-routine] Populating cache")
	c.L.Lock()
	cache["username"] = "Raja"
	cache["password"] = "Business"
	log.Println("[main-routine] Populated cache")
	log.Println("[main-routine] Giving Broadcast Signal for waiting go-routine")
	c.Broadcast() // can call by holding lock
	c.L.Unlock()
	// c.Signal() // can call without having lock
	wg.Wait()
	log.Println("[main-routine] Done!")
}

/*
raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/with_cond_by_brodcast.go
2021/06/06 11:59:57 [go-routine-2] waiting for Elements from cache : map[]
2021/06/06 11:59:57 [go-routine-1] waiting for Elements from cache : map[]
2021/06/06 11:59:58 [main-routine] preparing resources to populate cache
2021/06/06 11:59:59 [main-routine] Populating cache
2021/06/06 11:59:59 [main-routine] Populated cache
2021/06/06 11:59:59 [main-routine] Giving Broadcast Signal for waiting go-routine
2021/06/06 11:59:59 [go-routine-2] Got more than one Element from cache : map[password:Business username:Raja]
2021/06/06 11:59:59 [go-routine-1] Got Elements from cache : map[password:Business username:Raja]
2021/06/06 11:59:59 [main-routine] Done!
*/
