package main

import (
	"context"
	"log"
)

func main() {

	// generator generates integers in a separate goroutine
	// and sends them to the returned channel.
	// The callers of gen need to cancel the goroutine once
	// they consume 5 integers so that internal goroutine started by gen is not leaked
	generator := func(ctx context.Context) <-chan int {
		out := make(chan int)

		go func() {
			defer close(out)

			n := 1
			for {
				select {
				case out <- n:
					n++
				case <-ctx.Done():
					return
				}
			}
		}()

		return out
	}

	ctx, cancel := context.WithCancel(context.Background())

	ch := generator(ctx)

	for n := range ch {
		log.Printf("Received Number : %d\n", n)
		if n == 10 {
			cancel()
		}
	}
}

/*
raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run context/with_cancel.go
2021/06/07 17:06:40 Received Number : 1
2021/06/07 17:06:40 Received Number : 2
2021/06/07 17:06:40 Received Number : 3
2021/06/07 17:06:40 Received Number : 4
2021/06/07 17:06:40 Received Number : 5
2021/06/07 17:06:40 Received Number : 6
2021/06/07 17:06:40 Received Number : 7
2021/06/07 17:06:40 Received Number : 8
2021/06/07 17:06:40 Received Number : 9
2021/06/07 17:06:40 Received Number : 10
*/
