package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

var sockFile string

func main() {
	flag.StringVar(&sockFile, "sock", "/tmp/gounix.sock", "Path to socket file")
	flag.Parse()

	addr := &net.UnixAddr{
		Name: sockFile,
		Net:  "unixgram",
	}

	uc, err := net.DialUnix("unixgram", nil, addr)
	if err != nil {
		log.Fatalf("could not connect: %v\n", err)
	}

	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter message: ")

	for s.Scan() {
		n, err := uc.Write(s.Bytes())
		if err != nil {
			log.Fatalf("could not send: %v\n", err)
		}
		log.Printf("bytes sent: %d", n)
		fmt.Print("Enter message: ")
	}

}
