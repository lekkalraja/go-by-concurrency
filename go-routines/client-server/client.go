package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Something went wrong while dial-ing : %v", err)
	}
	defer conn.Close()
	io.Copy(os.Stdout, conn)
}
