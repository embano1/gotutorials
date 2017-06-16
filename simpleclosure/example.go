// first class functions
package main

import "fmt"

func start(a int) func() int {
	i := a
	return func() int {
		i++
		fmt.Println("i is :", i)
		return i
	}
}

type myfunc func() int

func start2(a int) myfunc {
	i := a
	return func() int {
		i++
		fmt.Println("i in start2 is :", i)
		return i
	}
}

func main() {

	f := start(0)
	g := start2(3)
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(g())
	fmt.Println(g())

	fmt.Printf("f is of type %T and g is of type %T\n", f, g)
}
