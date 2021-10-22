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

func SumAllTails(arraysToSum ...[]int) []int {

	var sums []int

	for _, arr := range arraysToSum {
		if len(arr) >= 2 {
			tail := arr[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}

	}

	return sums
}
