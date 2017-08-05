package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type address struct {
	Street string
	Zip    string
	City   string
}

type person struct {
	Name      string
	Age       int
	Addresses []address
}

func main() {

	peter := person{
		Name: "Peter",
		Age:  34,
		Addresses: []address{
			address{
				Street: "Alte Wache 4",
				Zip:    "12345",
				City:   "Munich",
			},
		},
	}

	// os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666) < check if file exists before creating; os.create will truncate existing file
	w, err := os.Create("tmp.json")
	if err != nil {
		fmt.Printf("could not create file: %v", err)
		os.Exit(1)
	}

	err = json.NewEncoder(w).Encode(peter)
	if err != nil {
		fmt.Printf("could not encode to JSON: %v", err)
		os.Exit(1)
	}

	// vscode tipp: to pretty print JSON use SHIFT-ALT-F
}
