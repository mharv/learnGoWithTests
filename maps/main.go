package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}
