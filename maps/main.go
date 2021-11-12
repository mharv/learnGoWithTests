package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {

	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, word string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = word
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
