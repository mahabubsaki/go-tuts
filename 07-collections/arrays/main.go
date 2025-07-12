package main

import (
	"fmt"
	"sort"
)

// === ARRAYS IN GO ===

func main() {
	fmt.Println("=== GO ARRAYS COMPREHENSIVE GUIDE ===")

	// === BASIC ARRAY DECLARATION ===
	fmt.Println("\n--- BASIC ARRAY DECLARATION ---")

	// Different ways to declare arrays
	var numbers1 [5]int
	fmt.Printf("Zero-value array: %v\n", numbers1)

	var numbers2 = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Initialized array: %v\n", numbers2)

	numbers3 := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("Short declaration: %v\n", numbers3)

	// Array with inferred length
	numbers4 := [...]int{100, 200, 300}
	fmt.Printf("Inferred length array: %v (length: %d)\n", numbers4, len(numbers4))

	/*
		JavaScript comparison:
		// JavaScript arrays are actually objects and are dynamic
		const numbers1 = new Array(5).fill(0);  // [0, 0, 0, 0, 0]
		const numbers2 = [1, 2, 3, 4, 5];
		const numbers3 = [10, 20, 30, 40, 50];
		const numbers4 = [100, 200, 300];

		// JavaScript arrays can hold different types
		const mixed = [1, "hello", true, null];
	*/

	// === ARRAY INITIALIZATION PATTERNS ===
	fmt.Println("\n--- ARRAY INITIALIZATION PATTERNS ---")

	// Partial initialization
	partial := [5]int{1, 2} // Remaining elements are zero
	fmt.Printf("Partial initialization: %v\n", partial)

	// Specific index initialization
	indexed := [5]int{0: 10, 2: 30, 4: 50}
	fmt.Printf("Indexed initialization: %v\n", indexed)

	// String arrays
	fruits := [3]string{"apple", "banana", "cherry"}
	fmt.Printf("String array: %v\n", fruits)

	// Boolean arrays
	flags := [4]bool{true, false, true, false}
	fmt.Printf("Boolean array: %v\n", flags)

	// === ARRAY ACCESS AND MODIFICATION ===
	fmt.Println("\n--- ARRAY ACCESS AND MODIFICATION ---")

	scores := [5]int{85, 92, 78, 95, 88}
	fmt.Printf("Original scores: %v\n", scores)

	// Access elements
	fmt.Printf("First score: %d\n", scores[0])
	fmt.Printf("Last score: %d\n", scores[len(scores)-1])

	// Modify elements
	scores[0] = 90
	scores[4] = 93
	fmt.Printf("Modified scores: %v\n", scores)

	// === ARRAY PROPERTIES ===
	fmt.Println("\n--- ARRAY PROPERTIES ---")

	temperatures := [7]float64{23.5, 25.0, 22.8, 26.1, 24.3, 21.9, 25.5}
	fmt.Printf("Temperatures: %v\n", temperatures)
	fmt.Printf("Array length: %d\n", len(temperatures))
	fmt.Printf("Array capacity: %d\n", cap(temperatures))

	// === ITERATING OVER ARRAYS ===
	fmt.Println("\n--- ITERATING OVER ARRAYS ---")

	colors := [4]string{"red", "green", "blue", "yellow"}

	// Using traditional for loop
	fmt.Println("Using traditional for loop:")
	for i := 0; i < len(colors); i++ {
		fmt.Printf("  colors[%d] = %s\n", i, colors[i])
	}

	// Using range
	fmt.Println("Using range:")
	for index, color := range colors {
		fmt.Printf("  colors[%d] = %s\n", index, color)
	}

	// Range with index only
	fmt.Println("Using range (index only):")
	for index := range colors {
		fmt.Printf("  Index %d: %s\n", index, colors[index])
	}

	// Range with value only
	fmt.Println("Using range (value only):")
	for _, color := range colors {
		fmt.Printf("  Color: %s\n", color)
	}

	// === ARRAY OPERATIONS ===
	fmt.Println("\n--- ARRAY OPERATIONS ---")

	values := [6]int{3, 1, 4, 1, 5, 9}
	fmt.Printf("Original values: %v\n", values)

	// Find sum
	sum := 0
	for _, value := range values {
		sum += value
	}
	fmt.Printf("Sum: %d\n", sum)

	// Find average
	average := float64(sum) / float64(len(values))
	fmt.Printf("Average: %.2f\n", average)

	// Find min and max
	min, max := values[0], values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	fmt.Printf("Min: %d, Max: %d\n", min, max)

	// === ARRAY COMPARISON ===
	fmt.Println("\n--- ARRAY COMPARISON ---")

	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 4}

	fmt.Printf("arr1: %v\n", arr1)
	fmt.Printf("arr2: %v\n", arr2)
	fmt.Printf("arr3: %v\n", arr3)
	fmt.Printf("arr1 == arr2: %t\n", arr1 == arr2)
	fmt.Printf("arr1 == arr3: %t\n", arr1 == arr3)

	// === ARRAY COPYING ===
	fmt.Println("\n--- ARRAY COPYING ---")

	original := [5]int{10, 20, 30, 40, 50}
	var copy1 [5]int
	copy1 = original // Arrays are value types, this creates a copy

	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Copy: %v\n", copy1)

	// Modify original
	original[0] = 999
	fmt.Printf("After modifying original[0] to 999:\n")
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Copy: %v\n", copy1) // Copy is unchanged

	// === MULTIDIMENSIONAL ARRAYS ===
	fmt.Println("\n--- MULTIDIMENSIONAL ARRAYS ---")

	// 2D array
	matrix := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	fmt.Println("2D Array (matrix):")
	for i, row := range matrix {
		for j, value := range row {
			fmt.Printf("matrix[%d][%d] = %2d  ", i, j, value)
		}
		fmt.Println()
	}

	// 3D array
	cube := [2][2][2]int{
		{
			{1, 2},
			{3, 4},
		},
		{
			{5, 6},
			{7, 8},
		},
	}

	fmt.Println("3D Array (cube):")
	for i, layer := range cube {
		fmt.Printf("Layer %d:\n", i)
		for j, row := range layer {
			for k, value := range row {
				fmt.Printf("  cube[%d][%d][%d] = %d\n", i, j, k, value)
			}
		}
	}

	// === ARRAY FUNCTIONS ===
	fmt.Println("\n--- ARRAY FUNCTIONS ---")

	// Function that takes array by value
	printArray := func(arr [5]int) {
		fmt.Printf("Array in function: %v\n", arr)
		arr[0] = 999 // This won't affect the original
	}

	// Function that takes array by reference
	modifyArray := func(arr *[5]int) {
		fmt.Printf("Array pointer in function: %v\n", *arr)
		(*arr)[0] = 999 // This will affect the original
	}

	testArray := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Before function call: %v\n", testArray)

	printArray(testArray)
	fmt.Printf("After printArray (by value): %v\n", testArray)

	modifyArray(&testArray)
	fmt.Printf("After modifyArray (by pointer): %v\n", testArray)

	// === ARRAY SORTING ===
	fmt.Println("\n--- ARRAY SORTING ---")

	unsorted := [6]int{64, 34, 25, 12, 22, 11}
	fmt.Printf("Unsorted: %v\n", unsorted)

	// Convert to slice for sorting
	slice := unsorted[:]
	sort.Ints(slice)
	fmt.Printf("Sorted: %v\n", unsorted)

	// === ARRAY SEARCHING ===
	fmt.Println("\n--- ARRAY SEARCHING ---")

	searchArray := [8]int{2, 4, 6, 8, 10, 12, 14, 16}
	target := 10

	// Linear search
	found := false
	foundIndex := -1
	for i, value := range searchArray {
		if value == target {
			found = true
			foundIndex = i
			break
		}
	}

	fmt.Printf("Searching for %d in %v\n", target, searchArray)
	if found {
		fmt.Printf("Found at index %d\n", foundIndex)
	} else {
		fmt.Printf("Not found\n")
	}

	// Binary search (array must be sorted)
	binarySearch := func(arr [8]int, target int) int {
		left, right := 0, len(arr)-1

		for left <= right {
			mid := (left + right) / 2
			if arr[mid] == target {
				return mid
			} else if arr[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return -1
	}

	index := binarySearch(searchArray, target)
	fmt.Printf("Binary search result: index %d\n", index)

	// === ARRAY LIMITATIONS ===
	fmt.Println("\n--- ARRAY LIMITATIONS ---")

	// Arrays have fixed size
	fmt.Println("Array limitations:")
	fmt.Println("1. Fixed size - cannot resize after declaration")
	fmt.Println("2. Size must be known at compile time")
	fmt.Println("3. Size is part of the type - [5]int != [6]int")
	fmt.Println("4. Passed by value - copying can be expensive")
	fmt.Println("5. No built-in methods like append, remove")

	// === ARRAY VS SLICE PREVIEW ===
	fmt.Println("\n--- ARRAY VS SLICE PREVIEW ---")

	// Array - fixed size
	fixedArray := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array: %v (type: %T)\n", fixedArray, fixedArray)

	// Slice - dynamic size
	dynamicSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice: %v (type: %T)\n", dynamicSlice, dynamicSlice)

	// Arrays are used when:
	fmt.Println("\nUse arrays when:")
	fmt.Println("1. Size is known and fixed")
	fmt.Println("2. You need value semantics")
	fmt.Println("3. Performance is critical and size is small")
	fmt.Println("4. You want to prevent accidental resizing")

	// === ARRAY BEST PRACTICES ===
	fmt.Println("\n--- ARRAY BEST PRACTICES ---")
	fmt.Println("1. Use slices instead of arrays for most cases")
	fmt.Println("2. Arrays are useful for fixed-size collections")
	fmt.Println("3. Use arrays for small, known-size data")
	fmt.Println("4. Consider memory usage with large arrays")
	fmt.Println("5. Use pointers for large arrays in functions")
	fmt.Println("6. Initialize arrays explicitly when needed")
	fmt.Println("7. Use range for iteration")
	fmt.Println("8. Remember arrays are value types")
	fmt.Println("9. Use [...] for compile-time size inference")
	fmt.Println("10. Consider multidimensional arrays for matrices")
}
