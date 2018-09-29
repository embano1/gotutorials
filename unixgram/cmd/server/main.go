package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"
)

var sockFile string
var sleep time.Duration

func main() {
	flag.StringVar(&sockFile, "sock", "/tmp/gounix.sock", "Path to socket file")
	flag.DurationVar(&sleep, "sleep", 5*time.Second, "Poll interval")
	flag.Parse()

	uc, err := net.ListenUnixgram("unixgram", &net.UnixAddr{
		Name: sockFile,
		Net:  "unixgram",
	})
	if err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)

	log.Printf("starting message reader... (poll interval: %v)\n", sleep)

	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			default:
				buff := make([]byte, 1024)
				oob := make([]byte, 1024)
				n, _, _, _, err := uc.ReadMsgUnix(buff, oob)
				if err != nil {
					log.Printf("error reading message: %v\n", err)
					continue
				}
				log.Printf("received message: %s (bytes: %d)\n", string(buff), n)
				time.Sleep(sleep)
			}
		}
	}()

	s := <-sigCh
	log.Printf("received signal: %v", s)
	cancel()
	wg.Wait()

	uc.Close()
	err = os.Remove(sockFile)
	if err != nil {
		log.Fatalf("could not remove socket file: %v", err)
	}
}
