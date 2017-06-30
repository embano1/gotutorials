// Warning: intentionally incorrect code with comments, for testing purposes only to show data races in multi-threaded execution mode
// Execution modes: inconsistent (1), pseudo-serialized (2), mutex guarded (3)
// (1) Inconsistent: go run main.go
// (2) Pseudo serialization (force single threaded): GOMAXPROCS=1 go run main.go
// (3) Mutex guarded: uncomment count.Lock() and count.Unlock()

package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
)

type counter struct {
	Result int
	sync.Mutex
}

const MAXLOOP = 100

func usage() {
	fmt.Fprintln(os.Stderr, `
Warning: intentionally incorrect code
For testing purposes only to show data races in multi-threaded execution mode
Execution modes: inconsistent (1), pseudo-serialized (2), mutex guarded (3)
(1) Inconsistent: just execute binary
(2) Pseudo serialization (force single threaded): set flag "-c=1"
(3) Mutex guarded: set flag "-m"
	`)
	fmt.Fprintf(os.Stderr, "usage: %s \n", os.Args[0])
	flag.PrintDefaults()
}

func main() {

	verbose := flag.Bool("v", false, "Verbose output")
	numcpus := flag.Int("c", runtime.NumCPU(), "Numbers of CPUs to use")
	mutex := flag.Bool("m", false, "Use mutex")
	flag.Usage = usage
	flag.Parse()

	runtime.GOMAXPROCS(*numcpus)
	count := &counter{}
	var wg sync.WaitGroup

	for i := 1; i <= MAXLOOP; i++ {
		wg.Add(1)
		go func(i int) {
			if *mutex {
				count.Lock()
				if *verbose {
					fmt.Printf("Routine %d where counter value = %d, increasing by 1...\n", i, count.Result)
				}
				count.Result++
				count.Unlock()
			} else {
				if *verbose {
					fmt.Printf("Routine %d where counter value = %d, increasing by 1...\n", i, count.Result)
				}
				count.Result++
			}
			wg.Done()
		}(i)

	}
	wg.Wait()

	fmt.Println("Result:", count.Result)
}
