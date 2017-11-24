package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/trace"

	"github.com/embano1/gotutorials/concprime"
)

func main() {

	n := flag.Int("n", 10, "Find first n primes")
	t := flag.Bool("t", false, "Enable tracing to STDOUT (default false)")
	tf := flag.String("f", "trace.out", "Name of trace output file")
	print := flag.Bool("p", false, "Print primes (default false)")
	flag.Parse()

	if *t {
		f, err := os.Create(*tf)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		err = trace.Start(f)
		if err != nil {
			panic(err)
		}
		defer trace.Stop()
	}

	fmt.Printf("Finding the first %d primes\n", *n)
	concprime.Find(*n, *print)
}
