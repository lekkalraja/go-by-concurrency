package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var pool = sync.Pool{
	New: func() interface{} {
		log.Println("allocation new bytes.Buffer")
		return new(bytes.Buffer)
	},
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i < 20; i++ {
		wg.Add(1)
		customLog(os.Stdout, "debug-string-"+strconv.Itoa(i), &wg)
	}
	wg.Wait()
}

func customLog(w io.Writer, debug string, wg *sync.WaitGroup) {
	// var b bytes.Buffer
	defer wg.Done()
	b := pool.Get().(*bytes.Buffer) // Getting from pool
	b.Reset()                       // reset buffer to erase old call content

	// log.Printf("The Object Reference of buffer : %p\n", &b)

	b.WriteString(time.Now().Format("15:04:03"))
	b.WriteString(" : ")
	b.WriteString(debug)
	b.WriteString("\n")

	w.Write(b.Bytes())

	pool.Put(b) // putting back to pool
}

/*
	WITHOUT Having Pool : (creating buffer object every time)
	====================
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e000
	12:32:12 : debug-string-1
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e060
	12:32:12 : debug-string-2
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e090
	12:32:12 : debug-string-3
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e0c0
	12:32:12 : debug-string-4
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e0f0
	12:32:12 : debug-string-5
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e120
	12:32:12 : debug-string-6
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e150
	12:32:12 : debug-string-7
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e180
	12:32:12 : debug-string-8
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e1b0
	12:32:12 : debug-string-9
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e1e0
	12:32:12 : debug-string-10
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e210
	12:32:12 : debug-string-11
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e240
	12:32:12 : debug-string-12
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e270
	12:32:12 : debug-string-13
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e2a0
	12:32:12 : debug-string-14
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e2d0
	12:32:12 : debug-string-15
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e300
	12:32:12 : debug-string-16
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e330
	12:32:12 : debug-string-17
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e360
	12:32:12 : debug-string-18
	2021/06/06 12:32:49 The Object Reference of buffer : 0xc00010e390
	12:32:12 : debug-string-19

	WITH POOL:
	==========
	raja@raja-Latitude-3460:~/Documents/coding/golang/go-by-concurrency$ go run sync/pool.go
	2021/06/06 12:38:33 allocation new bytes.Buffer
	12:38:12 : debug-string-1
	12:38:12 : debug-string-2
	12:38:12 : debug-string-3
	12:38:12 : debug-string-4
	12:38:12 : debug-string-5
	12:38:12 : debug-string-6
	12:38:12 : debug-string-7
	12:38:12 : debug-string-8
	12:38:12 : debug-string-9
	12:38:12 : debug-string-10
	12:38:12 : debug-string-11
	12:38:12 : debug-string-12
	12:38:12 : debug-string-13
	12:38:12 : debug-string-14
	12:38:12 : debug-string-15
	12:38:12 : debug-string-16
	12:38:12 : debug-string-17
	12:38:12 : debug-string-18
	12:38:12 : debug-string-19
*/
