package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type pwch struct {
	routine int
	number  float32
}

func printer(i int, c chan<- pwch) {
	f := rand.Float32()
	t := pwch{
		routine: i,
		number:  f,
	}
	c <- t
}

func writer(c <-chan pwch, done chan bool) {
	for {
		select {
		case v := <-c:
			fmt.Printf("Go routine %d generated number %f\n", v.routine, v.number)
		case <-time.After(time.Second * sleepmax):
			done <- true
		}
	}
}

func main() {
	pch := make(chan pwch)
	done := make(chan bool)

	c := flag.Int("c", counter, "Counter")
	flag.Parse()

	if flag.NFlag() < 1 {
		fmt.Println("Counter not specified, defaulting to", counter)
	}

	go func() {
		for i := 1; i <= *c; i++ {
			go printer(i, pch)
			sleep := time.Second * time.Duration(rand.Intn(randsleep))
			fmt.Printf("Sleeping for %v\n", sleep)
			time.Sleep(sleep)
		}
	}()

	go writer(pch, done)

	for {
		select {
		case <-done:
			fmt.Println("Writer slept too long, not really busy, giving up")
			os.Exit(0)
		case <-time.Tick(time.Second):
			fmt.Printf("Writer sleeping (max %v!), nothing to do\n", time.Second*sleepmax)

		}
	}

}
