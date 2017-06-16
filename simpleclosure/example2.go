package main

import (
	"fmt"
	"math/rand"
	"time"
)

type operator func([]int) string

func calc(o operator, x ...int) {
	fmt.Println(o(x))
}

func main() {
	var add = func(x []int) string {
		var a int
		for _, v := range x {
			a += v

		}
		return fmt.Sprintf("Addition gave %d", a)
	}

	var multi = func(x []int) string {
		a := 1
		for _, v := range x {
			a *= v

		}
		return fmt.Sprintf("Multiplaction gave %d", a)
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	i := r1.Intn(100)
	j := r2.Intn(100)

	fmt.Println("Working with: ", i, j)
	calc(add, i, j)
	calc(multi, i, j)

}
