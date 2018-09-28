package main

import (
	"embano1/bytes/internal/byteblob"
	"flag"
	"log"
	"net"
)

var size int
var socket string

func main() {
	flag.IntVar(&size, "size", 4096, "Number of bytes to transfer")
	flag.StringVar(&socket, "socket", "/tmp/go.sock", "Specify Unix Domain Socket path")
	flag.Parse()

	b, err := byteblob.New(size)
	if err != nil {
		log.Fatalf("could not get byte blob: %v\n", err)
	}

	conn, err := net.Dial("unix", socket)
	if err != nil {
		log.Fatalf("could not open connection: %v\n", err)
	}

	n, err := conn.Write(b)
	if err != nil {
		log.Fatalf("could not send byte blob: %v\n", err)
	}

	log.Printf("bytes sent: %d\n", n)

}
