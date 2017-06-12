package main

import (
	"errors"
	"fmt"
)

type testerror struct {
	s string
}

func (t testerror) Error() string {
	return t.s
}

func NewErr() error {
	return testerror{s: "default handling"}
}

func main() {
	err := errors.New("test")
	fmt.Printf("%T %q\n", err, err)

	test := NewErr()
	fmt.Printf("%T %q\n", test, test)

}
