package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	var t *time.Timer

	ch := make(chan bool)
	t = time.AfterFunc(randomDuration(), func() {
		log.Printf("Time Remaining : %v\n", time.Now().Sub(start))
		ch <- true
	})
	for time.Since(start) < 5*time.Second {
		<-ch
		t.Reset(randomDuration())
	}

}

// returns random duration between 0-1 seconds
func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

/*

Resolved race condition by avoiding the mutation of timer by multiple go routines:
==================================================================================
raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run -race race-detector/remove_race.go
2021/06/06 12:54:22 Time Remaining : 948.876341ms
2021/06/06 12:54:22 Time Remaining : 1.033102639s
2021/06/06 12:54:23 Time Remaining : 1.70043857s
2021/06/06 12:54:23 Time Remaining : 1.936129042s
2021/06/06 12:54:24 Time Remaining : 2.223814169s
2021/06/06 12:54:24 Time Remaining : 2.773912677s
2021/06/06 12:54:25 Time Remaining : 3.408273563s
2021/06/06 12:54:25 Time Remaining : 3.741363587s
2021/06/06 12:54:25 Time Remaining : 3.924967755s
2021/06/06 12:54:26 Time Remaining : 4.405907768s
2021/06/06 12:54:27 Time Remaining : 5.166842802s

*/
