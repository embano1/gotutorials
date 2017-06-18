package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	w1, w2 := "worker1", "worker2"
	wg := &sync.WaitGroup{}

	gen := func() <-chan int {
		c := make(chan int)
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			close(c)
		}()
		return c
	}

	ch := gen()
	wg.Add(1)
	go func() {

		for v := range ch {
			fmt.Println(w1, ":", v)
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {

		for v := range ch {
			fmt.Println(w2, ":", v)
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		}
		wg.Done()
	}()

	wg.Wait()

	// not really neccessary, for demo purposes only
	if _, ok := <-ch; !ok {
		fmt.Println("Channel closed, exiting")
	}

}

