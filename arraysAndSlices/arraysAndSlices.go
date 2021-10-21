package main

func main() {

}

func Sum(numbers [5]int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}
