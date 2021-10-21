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

	lengthOfArrays := len(arraysToSum)
	sums := make([]int, lengthOfArrays)

	for i, arr := range arraysToSum {
		sums[i] = Sum(arr)
	}

	return sums
}
