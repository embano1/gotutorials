package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

const j = `
			{"name": "John","age": 30}
			{"name": "Marry", "age": 34}
`

type stream struct {
	Name string
	Age  int
}

func main() {

	var s stream
	dec := json.NewDecoder(strings.NewReader(j))

	for {

		if err := dec.Decode(&s); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding input: ", err)
			os.Exit(1)
		}

		fmt.Println("Name: ", s.Name, "Age: ", s.Age)

	}
}
