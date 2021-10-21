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

func SumAll(arraysToSum ...[]int) (nums []int) {
	for _, arr := range arraysToSum {
		result := Sum(arr)
		nums = append(nums, result)
	}
	return
}
