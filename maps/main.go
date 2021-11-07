package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

type Dictionary map[string]string

func Search(dictionary Dictionary, key string) string {
	return dictionary[key]
}
