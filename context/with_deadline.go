package main

import (
	"context"
	"log"
	"runtime"
	"time"
)

type data struct {
	result string
}

func main() {

	compute := func(ctx context.Context) <-chan data {
		out := make(chan data)
		go func() {
			defer close(out)
			deadline, ok := ctx.Deadline()
			if ok {
				if deadline.Sub(time.Now()) < 10*time.Millisecond {
					log.Println("Not Sufficent time... Returning back...")
					return
				}
			}
			time.Sleep(10 * time.Millisecond)
			select {
			case out <- data{"1234"}:
			case <-ctx.Done():
				return
			}

		}()
		return out
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(12*time.Millisecond))
	defer cancel() // To Release all resources

	ch := compute(ctx)

	log.Printf("[%d - Active]Got Computed Result : %+v\n", runtime.NumGoroutine(), <-ch)
}
