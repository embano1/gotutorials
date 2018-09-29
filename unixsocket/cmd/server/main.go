package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/embano1/gotutorials/unixsocket/internal/byteblob"
)

var socket string
var bufSize int
var removeSock bool
var wg sync.WaitGroup

func main() {
	flag.StringVar(&socket, "socket", "/tmp/go.sock", "Specify Unix Domain Socket path")
	flag.IntVar(&bufSize, "size", 4096, "Size of buffer for receiving data")
	flag.BoolVar(&removeSock, "f", false, "Force cleanup of specified socket (default: false)")
	flag.Parse()

	if removeSock {
		os.Remove(socket)
	}

	log.Println("Starting echo server")
	ln, err := net.Listen("unix", socket)
	if err != nil {
		log.Fatal("Listen error: ", err)
	}
	defer ln.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func(c chan os.Signal, wg *sync.WaitGroup) {
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		cancel()
		wg.Wait()
		os.Exit(0)
	}(sigc, &wg)

	for {
		fd, err := ln.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}

		log.Println("got connection")
		wg.Add(1)
		go byteblob.Receive(ctx, &wg, fd, bufSize)
	}

}
