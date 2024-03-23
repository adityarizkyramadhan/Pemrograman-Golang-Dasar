package main

import (
	"fmt"
)

func FindMin(nums ...int) int {
	var min int = nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func FindMax(nums ...int) int {
	var max int = nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func SumMinMax(nums ...int) int {
	return FindMin(nums...) + FindMax(nums...)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
