package main

import (
	"fmt"
	"io"
)

func main() {
	fmt.Println("test")
}

func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}
