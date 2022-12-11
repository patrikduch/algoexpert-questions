package main

import "fmt"

func TwoNumberSum(array []int, target int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if (array[i] + array[j]) == target {
				return []int{array[i], array[j]}
			}
		}
	}
	return make([]int, 0)
}

func TwoNumberSumBetter(array []int, target int) []int {

	nums := make(map[int]bool)

	for i := 0; i < len(array); i++ {

		if nums[target-array[i]] {
			return []int{target - array[i], array[i]}
		}

		nums[array[i]] = true
	}

	return make([]int, 0)
}

func main() {
	testArray := TwoNumberSumBetter([]int{1, 4}, 5)
	fmt.Println(testArray)
}
