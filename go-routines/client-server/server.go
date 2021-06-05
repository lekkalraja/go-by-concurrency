package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Something went wrong while creating to Listener : %v", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Something went wrong while accepting to conn : %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, "Got Response from Server \n")
		if err != nil {
			log.Printf("Something went wrong while responding to client : %v", err)
		}
		time.Sleep(time.Second)
	}
}
