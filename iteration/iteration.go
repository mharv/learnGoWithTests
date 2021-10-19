package main

import "fmt"

func Repeat(input string, repeats int) string {
	var output string
	for i := 0; i < repeats; i++ {
		output += input
	}
	return output
}

func main() {
	fmt.Println(Repeat("a", 5))
}
