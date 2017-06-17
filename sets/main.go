package main

import (
	"fmt"
	"os"
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

func main() {
	s := NewSet()

	if len(os.Args[1:]) < 1 {
		fmt.Println("Please specify at least one entry to add")
		os.Exit(1)
	}

	for _, v := range os.Args[1:] {
		s.Add(v)
		fmt.Println(v)
	}

	fmt.Printf("Printing set: %+v", s.m)

}
