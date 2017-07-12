package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	heap := flag.String("heap", "512m", "Size of heap")
	stack := flag.String("stack", "512m", "Size of stack")

	flag.Parse()
	heapenv := os.Getenv("HEAP")
	stackenv := os.Getenv("STACK")

	fmt.Printf(`Heap from ENV: %s
Stack from ENV: %s
Heap from CLI: %s
Stack from CLI: %s
`,
		heapenv, stackenv, *heap, *stack)

	// Sleep so the pod does not finish and we can log inspect output
	time.Sleep(time.Hour)
}
