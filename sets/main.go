package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

var exists = struct{}{}

type set struct {
	m map[string]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *set) Add(value string) {
	s.m[value] = exists
}

func (s *set) Remove(value string) {
	delete(s.m, value)
}

func (s *set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}

func (s *set) Sort() {
	/*
		When iterating over a map with a range loop, the iteration order is not specified and is not guaranteed
		to be the same from one iteration to the next. Since Go 1 the runtime randomizes map iteration order,
		as programmers relied on the stable iteration order of the previous implementation.
		If you require a stable iteration order you must maintain a separate data structure that specifies that order.
	*/

	var keys []string
	for k := range s.m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	fmt.Println("Printing sorted set")
	for _, k := range keys {
		fmt.Println("Entry:", k)
	}

}

func main() {
	s := NewSet()

	sort := flag.Bool("s", false, "Sort set")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Please specify at least one entry to add")
		os.Exit(1)
	}

	for _, v := range flag.Args() {
		s.Add(v)
		// fmt.Println(v)
	}

	fmt.Printf("Printing set (unsorted): %v \n", s.m)
	if *sort {
		s.Sort()
	}
}
