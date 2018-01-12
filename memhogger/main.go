// Inspired by https://github.com/kubernetes-up-and-running/kuard
package main

// TODO: add runtime.mem
import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

// Grow in 100MB chunks
const growby = 100

type memhogger struct {
	leaks [][]byte
}

func main() {

	iter := flag.Int("i", 5, "How many 100MB chunks to allocate")
	delay := flag.Duration("d", 1*time.Second, "Delay between allocations")
	burn := flag.Bool("b", false, "Also burn CPU (default: off)")
	flag.Parse()

	var m memhogger

	// Burn a CPU
	if *burn {
		go func() {
			for {
				// Yield to gosched
				time.Sleep(time.Microsecond)
			}
		}()
	}

	fmt.Printf("Tuning go runtime for %d GOMAXPROCS\n", runtime.GOMAXPROCS(0))
	fmt.Printf("Eating %d MB of your tasty memory (chunk delay: %v)...\n", growby**(iter), *delay)
	fmt.Printf("CPU Burning enabled: %v\n", *burn)

	// Eat memory
	go func() {
		for i := 0; i < *iter; i++ {
			leak := make([]byte, growby*1024*1024, growby*1024*1024)
			for i := 0; i < len(leak); i++ {
				leak[i] = 'x'
			}

			m.leaks = append(m.leaks, leak)
			time.Sleep(*delay)
		}
	}()

	// TODO: implement http.mux with API for
	// - runtime.mem pprof
	// - reset leaks:
	//   m.leaks = nil
	//   runtime.GC()
	//   debug.FreeOSMemory()
	// https://github.com/kubernetes-up-and-running/kuard/blob/1fe8f0528424f7aaaebeff93213089e6e1c5ca57/pkg/memory/api.go#L59:21

	// TODO: remove with http.listenandserve()
	select {}
}
