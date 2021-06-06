package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	var t *time.Timer

	t = time.AfterFunc(randomDuration(), func() {
		log.Printf("Time Remaining : %v\n", time.Now().Sub(start))
		t.Reset(randomDuration())
	})
	time.Sleep(5 * time.Second)

}

// returns random duration between 0-1 seconds
func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

/*
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run race-detector/having_race.go
	2021/06/06 12:49:53 Time Remaining : 948.052153ms
	2021/06/06 12:49:53 Time Remaining : 1.030775662s
	2021/06/06 12:49:54 Time Remaining : 1.697741067s
	2021/06/06 12:49:54 Time Remaining : 1.933409849s
	2021/06/06 12:49:54 Time Remaining : 2.220854364s
	2021/06/06 12:49:55 Time Remaining : 2.770633512s
	2021/06/06 12:49:56 Time Remaining : 3.404561401s
	2021/06/06 12:49:56 Time Remaining : 3.737406715s
	2021/06/06 12:49:56 Time Remaining : 3.920847751s
	2021/06/06 12:49:57 Time Remaining : 4.401442628s

	Run With Race-Enabled
	=====================
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run -race race-detector/having_race.go
	2021/06/06 12:50:36 Time Remaining : 948.710583ms
	==================
	WARNING: DATA RACE
	Read at 0x00c0000c4018 by goroutine 8:
	main.main.func1()
		/home/raja/Documents/coding/golang/go-by-concurrency/race-detector/having_race.go:15 +0x115

	Previous write at 0x00c0000c4018 by main goroutine:
	main.main()
		/home/raja/Documents/coding/golang/go-by-concurrency/race-detector/having_race.go:13 +0x194

	Goroutine 8 (running) created at:
	time.goFunc()
		/usr/local/go/src/time/sleep.go:180 +0x51
	==================
	2021/06/06 12:50:36 Time Remaining : 1.036978009s
	2021/06/06 12:50:37 Time Remaining : 1.704193628s
	2021/06/06 12:50:37 Time Remaining : 1.94011791s
	2021/06/06 12:50:37 Time Remaining : 2.227826217s
	2021/06/06 12:50:38 Time Remaining : 2.777850692s
	2021/06/06 12:50:39 Time Remaining : 3.412244947s
	2021/06/06 12:50:39 Time Remaining : 3.744878602s
	2021/06/06 12:50:39 Time Remaining : 3.928391771s
	2021/06/06 12:50:40 Time Remaining : 4.409232613s
	Found 1 data race(s)
	exit status 66
*/
