package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/embano1/gotutorials/channels/2/generator"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	gen, done := generator.New(ctx)

	go func() {
		sig := <-sigs
		fmt.Println("main.go: Caught", sig, "Trying graceful shutdown...")
		cancel()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
	LOOP:
		for {
			select {
			case i := <-gen:
				fmt.Println("----------------------------------------")
				fmt.Println("main.go: got", i)
				sleep := rand.Intn(5)
				fmt.Println("main.go: Sleeping for", sleep, "seconds")
				time.Sleep(time.Duration(sleep) * time.Second)
			case <-ctx.Done():
				fmt.Println("main.go: We ware told to stop")
				wg.Done()
				break LOOP
			}
		}
	}()

	<-done
	wg.Wait()
	fmt.Println("main.go: Done")
}
