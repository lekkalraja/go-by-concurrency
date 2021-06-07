package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	req, _ := http.NewRequest("GET", "http://api.plos.org/search?q=title:DNA", nil)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)

	defer cancel()

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}

/*
	WITH TIMEOUT AS : 10 ms
	=======================
	panic: Get "http://api.plos.org/search?q=title:DNA": context deadline exceeded

	goroutine 1 [running]:
	main.main()
			/home/raja/Documents/coding/golang/go-by-concurrency/context/with_timeout.go:24 +0x425
	exit status 2
*/
