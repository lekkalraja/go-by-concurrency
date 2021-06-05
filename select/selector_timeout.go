package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- " I am Done"
	}()

	select {
	case data := <-ch:
		log.Printf("[Channel] Got Data : %s\n", data)
	case data := <-time.After(1 * time.Second):
		log.Printf("[Timeout] Got Data : %s\n", data)
	}
}

/*
	Scenario - 1: When we run goroutine with 4s delay and timeout 2s
		outuput: raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run select/selector_timeout.go
				 2021/06/05 17:26:11 [Timeout] Got Data : 2021-06-05 17:26:11.772097327 +0800 +08 m=+2.001177328

	Scenario - 2: When we run goroutine with 1s delay and timeout 2s
		outuput: raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run select/selector_timeout.go
				 2021/06/05 17:27:14 [Channel] Got Data :  I am Done
	Scenario - 3: When we run goroutine with 1s delay and timeout 1s
		outuput: Note: When timeout and delay response are same then most of the time getting timeout response,
					   but we may get response as well if we keep on try

		raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run select/selector_timeout.go
		2021/06/05 17:28:37 [Timeout] Got Data : 2021-06-05 17:28:37.252326187 +0800 +08 m=+1.000237313
		raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run select/selector_timeout.go
		2021/06/05 17:28:39 [Timeout] Got Data : 2021-06-05 17:28:39.129811458 +0800 +08 m=+1.000272893
		raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run select/selector_timeout.go
		2021/06/05 17:28:41 [Timeout] Got Data : 2021-06-05 17:28:41.082155219 +0800 +08 m=+1.001047382
		raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run select/selector_timeout.go
		2021/06/05 17:28:42 [Timeout] Got Data : 2021-06-05 17:28:42.916490594 +0800 +08 m=+1.000251826
		raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run select/selector_timeout.go
		2021/06/05 17:28:44 [Timeout] Got Data : 2021-06-05 17:28:44.842216589 +0800 +08 m=+1.000857565
		raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run select/selector_timeout.go
		2021/06/05 17:28:46 [Timeout] Got Data : 2021-06-05 17:28:46.720404037 +0800 +08 m=+1.000232752
		raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run select/selector_timeout.go
		2021/06/05 17:28:48 [Timeout] Got Data : 2021-06-05 17:28:48.617313025 +0800 +08 m=+1.000869618
*/
