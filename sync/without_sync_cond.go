package main

import (
	"log"
	"sync"
	"time"
)

var cache = make(map[string]string)

func main() {

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1)
	go func() {
		mu.Lock()
		defer wg.Done()
		for len(cache) == 0 {
			mu.Unlock()
			log.Printf("[go-routine] waiting for Elements from cache : %v\n", cache)
			time.Sleep(1 * time.Second)
			mu.Lock()
		}
		log.Printf("[go-routine] Got Elements from cache : %v\n", cache)
		mu.Unlock()
	}()

	time.Sleep(1 * time.Second)
	log.Println("[main-routine] preparing resources to populate cache")
	time.Sleep(1 * time.Second)
	log.Println("[main-routine] Populating cache")
	mu.Lock()
	cache["key"] = "ABC12345"
	mu.Unlock()
	log.Println("[main-routine] Populated cache")
	wg.Wait()
	log.Println("[main-routine] Done!")
}

/*
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/without_sync_cond.go
	2021/06/06 11:50:40 [go-routine] waiting for Elements from cache : map[]
	2021/06/06 11:50:41 [main-routine] preparing resources to populate cache
	2021/06/06 11:50:41 [go-routine] waiting for Elements from cache : map[]
	2021/06/06 11:50:42 [main-routine] Populating cache
	2021/06/06 11:50:42 [main-routine] Populated cache
	2021/06/06 11:50:42 [go-routine] Got Elements from cache : map[key:ABC12345]
	2021/06/06 11:50:42 [main-routine] Done!
*/
