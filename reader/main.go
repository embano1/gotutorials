package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

type Reader struct {
	s string
	i int64
}

var ErrNotInitialized = errors.New("could not read into []byte: slice not initialized")

func (r *Reader) Read(b []byte) (n int, err error) {
	if b == nil {
		return 0, ErrNotInitialized
	}

	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	// fmt.Println(r.i)
	// time.Sleep(1 * time.Second)
	return n, nil
}

func main() {
	r := Reader{s: "Hello Go World"}

	var b []byte
	// b := make([]byte, 1)
	buff := bytes.NewBuffer(b)

	// needs pointer otherwise "Does not implement io.Reader interface"
	buff.ReadFrom(&r)
	buff.WriteString(fmt.Sprintf("\nSomething I appended after reading from %T\n", r))
	buff.WriteTo(os.Stdout)
}
