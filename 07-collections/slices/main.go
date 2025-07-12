package main

import (
	"fmt"
	"sort"
	"strings"
)

// === SLICES IN GO ===

func main() {
	fmt.Println("=== GO SLICES COMPREHENSIVE GUIDE ===")

	// === BASIC SLICE DECLARATION ===
	fmt.Println("\n--- BASIC SLICE DECLARATION ---")

	// Different ways to declare slices
	var numbers1 []int
	fmt.Printf("Nil slice: %v (len: %d, cap: %d)\n", numbers1, len(numbers1), cap(numbers1))

	numbers2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice literal: %v (len: %d, cap: %d)\n", numbers2, len(numbers2), cap(numbers2))

	// Using make
	numbers3 := make([]int, 5)
	fmt.Printf("Made slice: %v (len: %d, cap: %d)\n", numbers3, len(numbers3), cap(numbers3))

	numbers4 := make([]int, 3, 10)
	fmt.Printf("Made slice with capacity: %v (len: %d, cap: %d)\n", numbers4, len(numbers4), cap(numbers4))

	/*
		JavaScript comparison:
		// JavaScript arrays are dynamic like Go slices
		const numbers1 = [];
		const numbers2 = [1, 2, 3, 4, 5];
		const numbers3 = new Array(5).fill(0);
		const numbers4 = new Array(10).fill(0).slice(0, 3);

		// JavaScript arrays have length property
		console.log(numbers2.length); // 5
	*/

	// === SLICE FROM ARRAY ===
	fmt.Println("\n--- SLICE FROM ARRAY ---")

	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original array: %v\n", array)

	// Different slicing operations
	slice1 := array[2:7]
	fmt.Printf("array[2:7]: %v (len: %d, cap: %d)\n", slice1, len(slice1), cap(slice1))

	slice2 := array[:5]
	fmt.Printf("array[:5]: %v (len: %d, cap: %d)\n", slice2, len(slice2), cap(slice2))

	slice3 := array[3:]
	fmt.Printf("array[3:]: %v (len: %d, cap: %d)\n", slice3, len(slice3), cap(slice3))

	slice4 := array[:]
	fmt.Printf("array[:]: %v (len: %d, cap: %d)\n", slice4, len(slice4), cap(slice4))

	// === SLICE OPERATIONS ===
	fmt.Println("\n--- SLICE OPERATIONS ---")

	// Append operations
	var fruits []string
	fmt.Printf("Empty slice: %v\n", fruits)

	fruits = append(fruits, "apple")
	fmt.Printf("After append 'apple': %v\n", fruits)

	fruits = append(fruits, "banana", "cherry")
	fmt.Printf("After append multiple: %v\n", fruits)

	moreFruits := []string{"date", "elderberry"}
	fruits = append(fruits, moreFruits...)
	fmt.Printf("After append slice: %v\n", fruits)

	// === SLICE CAPACITY AND GROWTH ===
	fmt.Println("\n--- SLICE CAPACITY AND GROWTH ---")

	var powers []int
	fmt.Printf("Initial: len=%d cap=%d %v\n", len(powers), cap(powers), powers)

	for i := 0; i < 10; i++ {
		powers = append(powers, i*i)
		fmt.Printf("After append %d²: len=%d cap=%d %v\n", i, len(powers), cap(powers), powers)
	}

	// === SLICE COPYING ===
	fmt.Println("\n--- SLICE COPYING ---")

	original := []int{1, 2, 3, 4, 5}

	// Shallow copy (same underlying array)
	shallow := original
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Shallow copy: %v\n", shallow)

	shallow[0] = 999
	fmt.Printf("After modifying shallow[0]:\n")
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Shallow copy: %v\n", shallow)

	// Deep copy using copy function
	original = []int{1, 2, 3, 4, 5} // Reset
	deep := make([]int, len(original))
	copy(deep, original)

	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Deep copy: %v\n", deep)

	deep[0] = 888
	fmt.Printf("After modifying deep[0]:\n")
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Deep copy: %v\n", deep)

	// === SLICE ITERATION ===
	fmt.Println("\n--- SLICE ITERATION ---")

	colors := []string{"red", "green", "blue", "yellow", "purple"}

	// Traditional for loop
	fmt.Println("Using traditional for loop:")
	for i := 0; i < len(colors); i++ {
		fmt.Printf("  colors[%d] = %s\n", i, colors[i])
	}

	// Range loop
	fmt.Println("Using range:")
	for index, color := range colors {
		fmt.Printf("  colors[%d] = %s\n", index, color)
	}

	// Range with index only
	fmt.Println("Range (index only):")
	for index := range colors {
		fmt.Printf("  Index %d\n", index)
	}

	// Range with value only
	fmt.Println("Range (value only):")
	for _, color := range colors {
		fmt.Printf("  Color: %s\n", color)
	}

	// === SLICE MODIFICATION ===
	fmt.Println("\n--- SLICE MODIFICATION ---")

	numbers := []int{10, 20, 30, 40, 50}
	fmt.Printf("Original: %v\n", numbers)

	// Remove element at index 2
	index := 2
	numbers = append(numbers[:index], numbers[index+1:]...)
	fmt.Printf("After removing index 2: %v\n", numbers)

	// Insert element at index 1
	index = 1
	value := 15
	numbers = append(numbers[:index], append([]int{value}, numbers[index:]...)...)
	fmt.Printf("After inserting 15 at index 1: %v\n", numbers)

	// Remove last element
	if len(numbers) > 0 {
		numbers = numbers[:len(numbers)-1]
		fmt.Printf("After removing last element: %v\n", numbers)
	}

	// Remove first element
	if len(numbers) > 0 {
		numbers = numbers[1:]
		fmt.Printf("After removing first element: %v\n", numbers)
	}

	// === SLICE FUNCTIONS ===
	fmt.Println("\n--- SLICE FUNCTIONS ---")

	// Function that modifies slice
	modifySlice := func(s []int) {
		for i := range s {
			s[i] *= 2
		}
	}

	// Function that returns new slice
	doubleSlice := func(s []int) []int {
		result := make([]int, len(s))
		for i, v := range s {
			result[i] = v * 2
		}
		return result
	}

	testSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Before modifySlice: %v\n", testSlice)
	modifySlice(testSlice)
	fmt.Printf("After modifySlice: %v\n", testSlice)

	newSlice := doubleSlice([]int{1, 2, 3, 4, 5})
	fmt.Printf("doubleSlice result: %v\n", newSlice)

	// === SLICE ALGORITHMS ===
	fmt.Println("\n--- SLICE ALGORITHMS ---")

	// Find sum
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for _, v := range values {
		sum += v
	}
	fmt.Printf("Sum of %v: %d\n", values, sum)

	// Find min and max
	if len(values) > 0 {
		min, max := values[0], values[0]
		for _, v := range values {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
		fmt.Printf("Min: %d, Max: %d\n", min, max)
	}

	// Check if slice contains value
	contains := func(slice []int, value int) bool {
		for _, v := range slice {
			if v == value {
				return true
			}
		}
		return false
	}

	fmt.Printf("Contains 5: %t\n", contains(values, 5))
	fmt.Printf("Contains 15: %t\n", contains(values, 15))

	// === SLICE SORTING ===
	fmt.Println("\n--- SLICE SORTING ---")

	unsorted := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Unsorted: %v\n", unsorted)

	sort.Ints(unsorted)
	fmt.Printf("Sorted: %v\n", unsorted)

	// Sort strings
	words := []string{"banana", "apple", "cherry", "date"}
	fmt.Printf("Unsorted words: %v\n", words)
	sort.Strings(words)
	fmt.Printf("Sorted words: %v\n", words)

	// Custom sorting
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	fmt.Printf("Before sorting: %v\n", people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Printf("After sorting by age: %v\n", people)

	// === SLICE FILTERING ===
	fmt.Println("\n--- SLICE FILTERING ---")

	numbers10 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Filter even numbers
	var evens []int
	for _, n := range numbers10 {
		if n%2 == 0 {
			evens = append(evens, n)
		}
	}
	fmt.Printf("Even numbers: %v\n", evens)

	// Filter using function
	filter := func(slice []int, predicate func(int) bool) []int {
		var result []int
		for _, v := range slice {
			if predicate(v) {
				result = append(result, v)
			}
		}
		return result
	}

	odds := filter(numbers10, func(n int) bool { return n%2 != 0 })
	fmt.Printf("Odd numbers: %v\n", odds)

	// === SLICE MAPPING ===
	fmt.Println("\n--- SLICE MAPPING ---")

	// Map function
	mapFunc := func(slice []int, transform func(int) int) []int {
		result := make([]int, len(slice))
		for i, v := range slice {
			result[i] = transform(v)
		}
		return result
	}

	squares := mapFunc([]int{1, 2, 3, 4, 5}, func(n int) int { return n * n })
	fmt.Printf("Squares: %v\n", squares)

	// === SLICE REDUCE ===
	fmt.Println("\n--- SLICE REDUCE ---")

	// Reduce function
	reduce := func(slice []int, initial int, reducer func(acc, curr int) int) int {
		result := initial
		for _, v := range slice {
			result = reducer(result, v)
		}
		return result
	}

	sumResult := reduce([]int{1, 2, 3, 4, 5}, 0, func(acc, curr int) int { return acc + curr })
	fmt.Printf("Sum using reduce: %d\n", sumResult)

	product := reduce([]int{1, 2, 3, 4, 5}, 1, func(acc, curr int) int { return acc * curr })
	fmt.Printf("Product using reduce: %d\n", product)

	// === MULTIDIMENSIONAL SLICES ===
	fmt.Println("\n--- MULTIDIMENSIONAL SLICES ---")

	// 2D slice
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("2D slice:")
	for i, row := range matrix {
		for j, val := range row {
			fmt.Printf("matrix[%d][%d] = %d  ", i, j, val)
		}
		fmt.Println()
	}

	// Creating 2D slice dynamically
	rows, cols := 3, 4
	dynamic2D := make([][]int, rows)
	for i := range dynamic2D {
		dynamic2D[i] = make([]int, cols)
		for j := range dynamic2D[i] {
			dynamic2D[i][j] = i*cols + j + 1
		}
	}

	fmt.Println("Dynamic 2D slice:")
	for _, row := range dynamic2D {
		fmt.Printf("%v\n", row)
	}

	// === SLICE PERFORMANCE ===
	fmt.Println("\n--- SLICE PERFORMANCE ---")

	// Pre-allocate when size is known
	n := 1000

	// Inefficient - multiple allocations
	var inefficient []int
	for i := 0; i < n; i++ {
		inefficient = append(inefficient, i)
	}

	// Efficient - single allocation
	efficient := make([]int, n)
	for i := 0; i < n; i++ {
		efficient[i] = i
	}

	fmt.Printf("Both slices created with %d elements\n", n)
	fmt.Printf("Pre-allocation is more efficient for known sizes\n")

	// === SLICE GOTCHAS ===
	fmt.Println("\n--- SLICE GOTCHAS ---")

	// Gotcha 1: Slice sharing underlying array
	arr := [5]int{1, 2, 3, 4, 5}
	sliceA := arr[1:3]
	sliceB := arr[2:4]

	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("sliceA: %v\n", sliceA)
	fmt.Printf("sliceB: %v\n", sliceB)

	sliceA[1] = 999 // This affects both slices!
	fmt.Printf("After sliceA[1] = 999:\n")
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("sliceA: %v\n", sliceA)
	fmt.Printf("sliceB: %v\n", sliceB)

	// Gotcha 2: Slice growth can change underlying array
	fmt.Println("\nGotcha 2: Slice growth")
	base := []int{1, 2, 3, 4, 5}
	sub := base[1:3]

	fmt.Printf("Base: %v (len: %d, cap: %d)\n", base, len(base), cap(base))
	fmt.Printf("Sub: %v (len: %d, cap: %d)\n", sub, len(sub), cap(sub))

	sub = append(sub, 999)
	fmt.Printf("After append to sub:\n")
	fmt.Printf("Base: %v\n", base)
	fmt.Printf("Sub: %v\n", sub)

	// === STRING SLICING ===
	fmt.Println("\n--- STRING SLICING ---")

	text := "Hello, World!"
	fmt.Printf("Original string: %s\n", text)

	// String slicing
	fmt.Printf("text[0:5]: %s\n", text[0:5])
	fmt.Printf("text[7:]: %s\n", text[7:])
	fmt.Printf("text[:5]: %s\n", text[:5])

	// Convert string to slice of bytes
	bytes := []byte(text)
	fmt.Printf("As bytes: %v\n", bytes)

	// Convert string to slice of runes (for Unicode)
	runes := []rune(text)
	fmt.Printf("As runes: %v\n", runes)

	// Work with Unicode
	unicode := "Hello, 世界!"
	fmt.Printf("Unicode string: %s\n", unicode)
	fmt.Printf("Byte length: %d\n", len(unicode))
	fmt.Printf("Rune length: %d\n", len([]rune(unicode)))

	// === SLICE UTILITIES ===
	fmt.Println("\n--- SLICE UTILITIES ---")

	// Join slice of strings
	words2 := []string{"Go", "is", "awesome"}
	sentence := strings.Join(words2, " ")
	fmt.Printf("Joined: %s\n", sentence)

	// Split string to slice
	parts := strings.Split(sentence, " ")
	fmt.Printf("Split: %v\n", parts)

	// Reverse slice
	reverse := func(slice []int) {
		for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	toReverse := []int{1, 2, 3, 4, 5}
	fmt.Printf("Before reverse: %v\n", toReverse)
	reverse(toReverse)
	fmt.Printf("After reverse: %v\n", toReverse)

	// === SLICE BEST PRACTICES ===
	fmt.Println("\n--- SLICE BEST PRACTICES ---")
	fmt.Println("1. Use slices instead of arrays for most cases")
	fmt.Println("2. Pre-allocate when size is known")
	fmt.Println("3. Be aware of slice/array sharing")
	fmt.Println("4. Use copy() for deep copying")
	fmt.Println("5. Check for nil before using")
	fmt.Println("6. Use range for iteration")
	fmt.Println("7. Consider capacity when appending")
	fmt.Println("8. Use built-in functions (sort, copy, etc.)")
	fmt.Println("9. Be careful with slice mutations")
	fmt.Println("10. Use string slicing for substring operations")
}
