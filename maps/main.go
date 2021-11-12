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
var ErrWordExists = errors.New("cannot add word because it already exists")

func (d Dictionary) Search(key string) (string, error) {

	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, word string) error {
	if d[key] == "" {
		d[key] = word
	} else {
		return ErrWordExists
	}
	return nil
}
