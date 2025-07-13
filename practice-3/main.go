package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for right > left {
		candidate := numbers[right] + numbers[left]
		if candidate == target {
			return []int{left + 1, right + 1}
		} else if candidate > target {
			right = right - 1
		} else {
			left = left + 1
		}
	}
	return nil
}
func main() {
	result := twoSum([]int{5, 25, 75}, 100)
	fmt.Println(result)
}
