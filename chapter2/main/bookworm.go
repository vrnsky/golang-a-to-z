package main

import (
	"encoding/json"
	"os"
)

var bookworms []Bookworm


type Bookworm struct {
	Name string `json:"name"`
	Book []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title string `json:"title"`
}

func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Declare the variable in which the file will be decoded.
	var bookworms []Bookworm

	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

