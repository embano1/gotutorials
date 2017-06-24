package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type stream struct {
	Name string
	Age  int
}

func main() {

	input := []byte(`[
		{"name": "John","age": 30},
		{"name": "Marry", "age": 34}
	]`)

	var s []stream
	err := json.Unmarshal(input, &s)
	if err != nil {
		fmt.Println("Error decoding input :", err)
		os.Exit(1)
	}

	for _, v := range s {
		fmt.Printf("%v is %d years old.\n", v.Name, v.Age)
	}
}
