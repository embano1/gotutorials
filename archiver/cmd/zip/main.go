package main

import (
	"fmt"
	"os"

	"github.com/mholt/archiver"
	"github.com/pkg/errors"
)

func usage() string {
	return fmt.Sprintf("usage: %s <input> <output>.zip", os.Args[0])
}

func run(out string, in ...string) error {
	f, err := os.Create(out)
	if err != nil {
		return errors.Wrap(err, "could not open file")
	}

	err = archiver.Zip.Write(f, in)
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

	err := run(out, in)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
