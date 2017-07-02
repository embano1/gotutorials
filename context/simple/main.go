package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func counter(ctx context.Context, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Println("Int Generator starting")
	i := 0

	for {
		select {
		case <-ctx.Done():
			fmt.Println("[Int Generator] Got signal to stop, shutting down")
			time.Sleep(time.Second)
			return
		default:
			fmt.Println("i is ", i)
			time.Sleep(time.Second)
			i++
		}
	}

}

func main() {

	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go counter(ctx, &wg)

	time.Sleep(5 * time.Second)
	fmt.Println("Shutting down.")
	cancel()
	wg.Wait()

}
