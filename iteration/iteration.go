package main

func Repeat(input string) string {
	var output string
	for i := 0; i < 5; i++ {
		output = output + input
	}
	return output
}
