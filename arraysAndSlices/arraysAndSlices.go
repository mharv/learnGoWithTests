package main

func main() {

}

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(arraysToSum ...[]int) []int {

	var sums []int

	for _, arr := range arraysToSum {
		sums = append(sums, Sum(arr))
	}

	return sums
}
