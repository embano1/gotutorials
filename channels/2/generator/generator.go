package generator

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// New returns a read-only channel of type int, used to receive generated numbers.
// If receives from this channel block, the result is discarded and generator throttles (1s).
// It also returns a read-only channel of type struct{} to signal that it´s safe to exit the program when a receive from this channel succeeds (i.e. it gets closed).
func New(ctx context.Context) (<-chan int, <-chan struct{}) {
	gen := make(chan int)
	done := make(chan struct{})
	go generate(ctx, gen, done)
	return gen, done
}

func generate(ctx context.Context, c chan<- int, done chan<- struct{}) {
	defer close(done)

LOOP:
	for {
		i := rand.Intn(100)

		select {
		case c <- i:
			fmt.Println("Sent", i)
		case <-ctx.Done():
			fmt.Println("generator.go: We were told to stop")
			break LOOP
		default:
			fmt.Println("generator.go: Blocked, throwing away", i, "and throttling")
			time.Sleep(time.Second)
		}
	}
}
