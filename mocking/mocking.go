package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("test")
}

func Countdown(out *bytes.Buffer) {
	fmt.Fprint(out, "3")
}
