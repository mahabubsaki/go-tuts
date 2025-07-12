package main

import (
	"fmt"
)

func calculateDifference(bigSize []int, smallSize []int, diff int, name string) []int {
	for _, bigVal := range bigSize {
		for _, smallVal := range smallSize {
			if bigVal-smallVal == diff {
				if name == "alice" {
					return []int{bigVal, smallVal}
				} else {
					return []int{smallVal, bigVal}
				}
			}
		}
	}
	return []int{}
}

// Helper function to sum an array
func sumArray(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	aliceTotal := sumArray(aliceSizes)
	bobTotal := sumArray(bobSizes)

	targetDifference := (bobTotal - aliceTotal) / 2

	// Put Bob's boxes in a hash set for O(1) lookup
	bobSet := make(map[int]bool)
	for _, bobBox := range bobSizes {
		bobSet[bobBox] = true
	}

	// For each Alice box, check if the required Bob box exists
	for _, aliceBox := range aliceSizes {
		requiredBobBox := aliceBox + targetDifference
		if bobSet[requiredBobBox] {
			return []int{aliceBox, requiredBobBox}
		}
	}

	return []int{}

}

func main() {
	fmt.Println("Try programiz.pro")
	numbers1 := []int{1, 2, 5}
	numbers2 := []int{2, 4}
	str := fairCandySwap(numbers1, numbers2)
	fmt.Println(str)
}
