package main

import (
	"fmt"
	"os"

	"github.com/mholt/archiver"
	"github.com/pkg/errors"
)

func usage() string {
	return fmt.Sprintf("usage: %s <input>.rar <destination>", os.Args[0])
}

func run(in, out string) error {
	f, err := os.Open(in)
	if err != nil {
		return errors.Wrap(err, "could not open file")
	}

	err = archiver.Rar.Read(f, out)
	if err != nil {
		return errors.Wrap(err, "could not read archive")
	}
	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println(usage())
		os.Exit(1)
	}

	in := os.Args[1]
	out := os.Args[2]

	err := run(in, out)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
