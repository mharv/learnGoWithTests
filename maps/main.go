package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("hello world")
}

type Dictionary map[string]string

func Search(dictionary Dictionary, key string) (string, error) {

	definition, ok := dictionary[key]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}

	return definition, nil
}
