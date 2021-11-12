package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("hello world")
}

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func Search(dictionary Dictionary, key string) (string, error) {

	definition, ok := dictionary[key]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
