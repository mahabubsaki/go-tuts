// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message
package main

import (
	"fmt"
	"math"
	"strconv"
)

func summaryRanges(nums []int) []string {
	result := []string{}
	if len(nums) == 0 {

		return result
	} else {

	}
	numTup := [2]int{math.MaxInt32, math.MaxInt32}
	length := len(nums)
	contigous := false
	for index, value := range nums {
		if contigous || index == 0 {
			numTup[0] = value
			contigous = false
		}
		if index == length {
			continue
		}

		if index+1 < length && value+1 == nums[index+1] {
			if numTup[0] != value {
				numTup[1] = value
			}

		} else {
			if numTup[0] != value {
				numTup[1] = value
			}
			if numTup[1] == math.MaxInt32 {
				result = append(result, strconv.Itoa(numTup[0]))
				contigous = true
			} else {
				formated := fmt.Sprintf("%d->%d", numTup[0], numTup[1])
				result = append(result, formated)
				contigous = true
			}
			numTup[0] = math.MaxInt32
			numTup[1] = math.MaxInt32
		}
	}

	return result
}

func main() {
	fmt.Println("Try programiz.pro")
	numbers := []int{-10, -3, -2, 0, 2, 3, 4, 6, 8, 9, 10, 11, 15}
	str := summaryRanges(numbers)
	fmt.Println(str)
}

// loop=>
//   add the current number to a tuple
//   check currents next number and next number equal or not
//   if not equal,
//      if tuple have one number=>tuples first number to string array
//      if tuple have two number=> add formated number like start->end
//   if equal
//      add current number to tuple if not same
// if index zero or last index skip that step
