package main

import (
	"log"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)

	var balance float64
	var wg sync.WaitGroup
	var m sync.Mutex

	deposit := func(amount float64) {
		m.Lock()
		defer m.Unlock()
		balance += amount
	}

	withdraw := func(amount float64) {
		m.Lock()
		defer m.Unlock()
		balance -= amount
	}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			withdraw(1)
		}()
	}

	wg.Wait()

	log.Printf("The Remaining Balnce after all txns : %.2f \n", balance)

}

/*
	If we don't use any mutex the output is invalid and it is inconsistant like below
	=================================================================================
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/mutex.go
	2021/06/05 22:04:43 The Remaining Balnce after all txns : 238.00
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/mutex.go
	2021/06/05 22:04:46 The Remaining Balnce after all txns : -164.00
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/mutex.go
	2021/06/05 22:04:48 The Remaining Balnce after all txns : -118.00
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/mutex.go
	2021/06/05 22:04:50 The Remaining Balnce after all txns : -125.00
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$

	After Making routine-safe the out is correct and consistant
	===========================================================
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/mutex.go
	2021/06/05 22:07:32 The Remaining Balnce after all txns : 0.00
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/mutex.go
	2021/06/05 22:07:33 The Remaining Balnce after all txns : 0.00
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/mutex.go
	2021/06/05 22:07:34 The Remaining Balnce after all txns : 0.00
*/
