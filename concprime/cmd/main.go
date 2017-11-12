package main

import (
	"flag"
	"fmt"

	"github.com/embano1/gotutorials/concprime"
)

func main() {

	n := flag.Int("n", 10, "Find first n primes")
	print := flag.Bool("p", false, "Print primes (default false)")
	flag.Parse()
	fmt.Printf("Finding the first %d primes\n", *n)
	concprime.Find(*n, *print)
}
