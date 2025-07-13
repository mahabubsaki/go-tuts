package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	result := make(map[[3]int]bool)
	hashmap := make(map[int]int)
	for index, val := range nums {

		for innerIndex := index + 1; innerIndex < len(nums); innerIndex++ {
			targetOffset := val + nums[innerIndex]
			offsetAvailable := hashmap[targetOffset*-1]
			if (nums[offsetAvailable]+nums[index]+nums[innerIndex]) == 0 && offsetAvailable != index && index != innerIndex && offsetAvailable != innerIndex {
				candidate := [3]int{nums[offsetAvailable], nums[index], nums[innerIndex]}
				sort.Ints(candidate[:])
				result[candidate] = true
			}
			hashmap[nums[innerIndex]] = innerIndex

		}
	}
	keys := [][]int{}
	for k := range result {
		keys = append(keys, []int{k[0], k[1], k[2]})
	}
	return keys
}
func main() {
	result := threeSum([]int{1, 2, -2, -1})
	fmt.Println(result)
}
