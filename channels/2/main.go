package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/embano1/gotutorials/channels/2/generator"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {

	sigs := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.Background())
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	gen, done := generator.New(ctx)

	go func() {
		for i := range gen {
			fmt.Println("----------------------------------------")
			fmt.Println("main.go: got", i)
			sleep := rand.Intn(5)
			fmt.Println("main.go: Sleeping for", sleep, "seconds")
			time.Sleep(time.Duration(sleep) * time.Second)
		}
	}()

	sig := <-sigs
	fmt.Println("main.go: Caught", sig, "Trying graceful shutdown...")
	cancel()

	<-done
	fmt.Println("main.go: Done")
	os.Exit(0)
}
