package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	input := "a"
	repeated := Repeat(input, 5)
	expected := strings.Repeat(input, 5)

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeatedString := Repeat("You", 3)
	fmt.Println(repeatedString)
	// Output: YouYouYou
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
