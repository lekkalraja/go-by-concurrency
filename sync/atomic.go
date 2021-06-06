package main

import (
	"log"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(4)

	var counter int64

	var wg sync.WaitGroup
	//var m sync.Mutex

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			/* m.Lock()
			   defer m.Unlock() */
			for j := 0; j < 10000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	log.Printf("The final Counter value : %d\n", counter)
}

/*
	If we don't use any locks/atomic variables then the counter value is inconsistant
	==================================================================================
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/atomic.go
	2021/06/05 22:14:56 The final Counter value : 138567
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/atomic.go
	2021/06/05 22:14:57 The final Counter value : 147110
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/atomic.go
	2021/06/05 22:14:58 The final Counter value : 157417
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/atomic.go
	2021/06/05 22:14:59 The final Counter value : 139581

	After doing atomic operation the counter value is correct and consistant
	=========================================================================
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/atomic.go
	2021/06/05 22:18:47 The final Counter value : 500000
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/atomic.go
	2021/06/05 22:18:48 The final Counter value : 500000
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/atomic.go
	2021/06/05 22:18:49 The final Counter value : 500000


	Even After doing locing operation's the counter value is correct and consistant
	===============================================================================
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/atomic.go
	2021/06/05 22:17:43 The final Counter value : 500000
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/atomic.go
	2021/06/05 22:17:45 The final Counter value : 500000
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/atomic.go
	2021/06/05 22:17:46 The final Counter value : 500000
*/
