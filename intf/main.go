package main

import (
	"fmt"
	"os"

	"github.com/embano1/gotutorials/intf/pkg/greeter"
)

func main() {
	var s string
	if len(os.Args) < 2 {
		s = ""
	} else {
		s = os.Args[1]
	}

	g := greeter.New(s)

	fmt.Println(g.Greet())

	// Uncomment to see pointer and type of *greeter.lobby (unexported)
	// fmt.Printf("%p\n", g)
	// fmt.Printf("%T\n", g)
}
