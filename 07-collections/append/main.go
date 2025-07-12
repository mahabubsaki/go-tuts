package main

import (
	"fmt"
	"strings"
)

// === GO APPEND FUNCTION COMPREHENSIVE GUIDE ===

/*
APPEND PHILOSOPHY:
- append() is a built-in function for adding elements to slices
- It returns a new slice with the elements added
- May allocate new underlying array if capacity is exceeded
- Essential for dynamic data structures in Go

COMPARISON WITH JAVASCRIPT:
// JavaScript: Array methods
let arr = [1, 2, 3];
arr.push(4);                    // Modifies original array
arr.push(5, 6, 7);             // Add multiple elements
arr = arr.concat([8, 9]);      // Creates new array

// Go: append function
slice := []int{1, 2, 3}
slice = append(slice, 4)                    // Returns new slice
slice = append(slice, 5, 6, 7)             // Add multiple elements
slice = append(slice, []int{8, 9}...)      // Concatenate slices
*/

// 1. BASIC APPEND USAGE
func demonstrateBasicAppend() {
	fmt.Println("=== BASIC APPEND ===")

	// Start with empty slice
	var numbers []int
	fmt.Printf("Initial: %v, len=%d, cap=%d\n", numbers, len(numbers), cap(numbers))

	// Add single element
	numbers = append(numbers, 1)
	fmt.Printf("After append(1): %v, len=%d, cap=%d\n", numbers, len(numbers), cap(numbers))

	// Add more elements
	numbers = append(numbers, 2, 3, 4)
	fmt.Printf("After append(2,3,4): %v, len=%d, cap=%d\n", numbers, len(numbers), cap(numbers))
}

// 2. APPEND AND CAPACITY
func demonstrateAppendCapacity() {
	fmt.Println("\n=== APPEND AND CAPACITY ===")

	// Start with slice of capacity 2
	slice := make([]int, 0, 2)
	fmt.Printf("Initial: %v, len=%d, cap=%d\n", slice, len(slice), cap(slice))

	for i := 1; i <= 8; i++ {
		slice = append(slice, i)
		fmt.Printf("After append(%d): %v, len=%d, cap=%d\n", i, slice, len(slice), cap(slice))
	}
}

// 3. APPEND TO NIL SLICE
func demonstrateAppendToNil() {
	fmt.Println("\n=== APPEND TO NIL SLICE ===")

	var slice []string
	fmt.Printf("Nil slice: %v, len=%d, cap=%d\n", slice, len(slice), cap(slice))

	// Append to nil slice works perfectly
	slice = append(slice, "hello")
	fmt.Printf("After append: %v, len=%d, cap=%d\n", slice, len(slice), cap(slice))

	slice = append(slice, "world", "!")
	fmt.Printf("After more appends: %v, len=%d, cap=%d\n", slice, len(slice), cap(slice))
}

// 4. APPEND SLICE TO SLICE
func demonstrateAppendSliceToSlice() {
	fmt.Println("\n=== APPEND SLICE TO SLICE ===")

	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)

	// Use ... to unpack slice2
	result := append(slice1, slice2...)
	fmt.Printf("append(slice1, slice2...): %v\n", result)

	// Original slices unchanged
	fmt.Printf("slice1 after append: %v\n", slice1)
	fmt.Printf("slice2 after append: %v\n", slice2)
}

// 5. APPEND WITH DIFFERENT TYPES
func demonstrateAppendWithTypes() {
	fmt.Println("\n=== APPEND WITH DIFFERENT TYPES ===")

	// String slice
	words := []string{"Go", "is"}
	words = append(words, "awesome", "!")
	fmt.Printf("Words: %v\n", words)

	// Boolean slice
	flags := []bool{true, false}
	flags = append(flags, true, true, false)
	fmt.Printf("Flags: %v\n", flags)

	// Struct slice
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
	}

	people = append(people, Person{Name: "Charlie", Age: 35})
	fmt.Printf("People: %v\n", people)
}

// 6. APPEND AND MEMORY REALLOCATION
func demonstrateMemoryReallocation() {
	fmt.Println("\n=== MEMORY REALLOCATION ===")

	slice := make([]int, 0, 2)
	fmt.Printf("Initial slice pointer: %p\n", slice)

	// Fill up to capacity
	slice = append(slice, 1, 2)
	fmt.Printf("After filling to capacity: %p, %v\n", slice, slice)

	// This will cause reallocation
	slice = append(slice, 3)
	fmt.Printf("After exceeding capacity: %p, %v\n", slice, slice)
}

// 7. APPEND IN FUNCTIONS
func appendToSlice(slice []int, value int) []int {
	return append(slice, value)
}

func demonstrateAppendInFunctions() {
	fmt.Println("\n=== APPEND IN FUNCTIONS ===")

	original := []int{1, 2, 3}
	fmt.Printf("Original: %v\n", original)

	// Function returns new slice
	modified := appendToSlice(original, 4)
	fmt.Printf("Modified: %v\n", modified)
	fmt.Printf("Original unchanged: %v\n", original)
}

// 8. APPEND WITH PRE-ALLOCATION
func demonstratePreAllocation() {
	fmt.Println("\n=== APPEND WITH PRE-ALLOCATION ===")

	// Without pre-allocation
	var slice1 []int
	for i := 0; i < 1000; i++ {
		slice1 = append(slice1, i)
	}
	fmt.Printf("Without pre-allocation: len=%d, cap=%d\n", len(slice1), cap(slice1))

	// With pre-allocation
	slice2 := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		slice2 = append(slice2, i)
	}
	fmt.Printf("With pre-allocation: len=%d, cap=%d\n", len(slice2), cap(slice2))
}

// 9. APPEND PATTERNS
func demonstrateAppendPatterns() {
	fmt.Println("\n=== APPEND PATTERNS ===")

	// Pattern 1: Building strings
	var parts []string
	parts = append(parts, "Hello")
	parts = append(parts, "beautiful")
	parts = append(parts, "world")
	result := strings.Join(parts, " ")
	fmt.Printf("Built string: %s\n", result)

	// Pattern 2: Filtering
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evens []int
	for _, num := range numbers {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	fmt.Printf("Even numbers: %v\n", evens)

	// Pattern 3: Transforming
	var squares []int
	for _, num := range numbers {
		squares = append(squares, num*num)
	}
	fmt.Printf("Squares: %v\n", squares)
}

// 10. APPEND WITH INTERFACES
func demonstrateAppendWithInterfaces() {
	fmt.Println("\n=== APPEND WITH INTERFACES ===")

	// Empty interface slice can hold any type
	var items []interface{}
	items = append(items, 42)
	items = append(items, "hello")
	items = append(items, 3.14)
	items = append(items, true)

	fmt.Printf("Mixed types: %v\n", items)

	// Type assertion to get original values
	for i, item := range items {
		fmt.Printf("Item %d: %v (type: %T)\n", i, item, item)
	}
}

// 11. APPEND ERROR HANDLING
func demonstrateAppendErrorHandling() {
	fmt.Println("\n=== APPEND ERROR HANDLING ===")

	// Function that might fail
	parseAndAppend := func(slice []int, str string) ([]int, error) {
		var value int
		if str == "invalid" {
			return slice, fmt.Errorf("invalid input: %s", str)
		}

		// Simulate parsing
		switch str {
		case "one":
			value = 1
		case "two":
			value = 2
		case "three":
			value = 3
		default:
			return slice, fmt.Errorf("unknown value: %s", str)
		}

		return append(slice, value), nil
	}

	var numbers []int
	inputs := []string{"one", "two", "invalid", "three"}

	for _, input := range inputs {
		var err error
		numbers, err = parseAndAppend(numbers, input)
		if err != nil {
			fmt.Printf("Error with '%s': %v\n", input, err)
		} else {
			fmt.Printf("Successfully appended from '%s': %v\n", input, numbers)
		}
	}
}

// 12. APPEND WITH CHANNELS
func demonstrateAppendWithChannels() {
	fmt.Println("\n=== APPEND WITH CHANNELS ===")

	ch := make(chan int, 5)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	var results []int
	for value := range ch {
		results = append(results, value)
	}

	fmt.Printf("Results from channel: %v\n", results)
}

// 13. APPEND PERFORMANCE CONSIDERATIONS
func demonstrateAppendPerformance() {
	fmt.Println("\n=== APPEND PERFORMANCE ===")

	// Inefficient: repeated small appends
	var slice1 []int
	for i := 0; i < 1000; i++ {
		slice1 = append(slice1, i)
	}
	fmt.Printf("Inefficient append: final cap=%d\n", cap(slice1))

	// Efficient: pre-allocate capacity
	slice2 := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		slice2 = append(slice2, i)
	}
	fmt.Printf("Efficient append: final cap=%d\n", cap(slice2))

	// Efficient: batch append
	batch := make([]int, 1000)
	for i := range batch {
		batch[i] = i
	}
	slice3 := append([]int{}, batch...)
	fmt.Printf("Batch append: final cap=%d\n", cap(slice3))
}

// 14. APPEND BEST PRACTICES
func demonstrateAppendBestPractices() {
	fmt.Println("\n=== APPEND BEST PRACTICES ===")

	// 1. Always assign the result of append
	slice := []int{1, 2, 3}
	slice = append(slice, 4) // ✅ Good
	// append(slice, 5)       // ❌ Bad: result ignored

	// 2. Pre-allocate when size is known
	knownSize := 100
	efficientSlice := make([]int, 0, knownSize)
	for i := 0; i < knownSize; i++ {
		efficientSlice = append(efficientSlice, i)
	}

	// 3. Use ... operator for slice concatenation
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	combined := append(slice1, slice2...) // ✅ Good

	// 4. Be careful with slice references
	originalSlice := []int{1, 2, 3}
	copySlice := append([]int{}, originalSlice...) // Create independent copy

	fmt.Printf("Original: %v\n", originalSlice)
	fmt.Printf("Copy: %v\n", copySlice)
	fmt.Printf("Combined: %v\n", combined)
	fmt.Printf("Efficient slice length: %d\n", len(efficientSlice))
}

// 15. APPEND GOTCHAS
func demonstrateAppendGotchas() {
	fmt.Println("\n=== APPEND GOTCHAS ===")

	// Gotcha 1: Forgetting to assign result
	slice := []int{1, 2, 3}
	fmt.Printf("Before append: %v\n", slice)
	_ = append(slice, 4)                            // Result ignored!
	fmt.Printf("After ignored append: %v\n", slice) // Still [1, 2, 3]

	// Gotcha 2: Shared underlying array
	original := []int{1, 2, 3}
	slice1 := original[:2] // [1, 2]
	slice2 := original[1:] // [2, 3]

	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Slice1: %v\n", slice1)
	fmt.Printf("Slice2: %v\n", slice2)

	// Appending to slice1 might affect original if capacity allows
	slice1 = append(slice1, 99)
	fmt.Printf("After append to slice1: %v\n", slice1)
	fmt.Printf("Original after append: %v\n", original)

	// Gotcha 3: Nil slice vs empty slice
	var nilSlice []int
	emptySlice := []int{}

	fmt.Printf("nil slice: %v, len=%d, cap=%d\n", nilSlice, len(nilSlice), cap(nilSlice))
	fmt.Printf("empty slice: %v, len=%d, cap=%d\n", emptySlice, len(emptySlice), cap(emptySlice))

	// Both work the same with append
	nilSlice = append(nilSlice, 1)
	emptySlice = append(emptySlice, 1)

	fmt.Printf("nil slice after append: %v\n", nilSlice)
	fmt.Printf("empty slice after append: %v\n", emptySlice)
}

func main() {
	fmt.Println("=== GO APPEND FUNCTION COMPREHENSIVE GUIDE ===")

	// === BASIC APPEND ===
	demonstrateBasicAppend()

	// === APPEND AND CAPACITY ===
	demonstrateAppendCapacity()

	// === APPEND TO NIL SLICE ===
	demonstrateAppendToNil()

	// === APPEND SLICE TO SLICE ===
	demonstrateAppendSliceToSlice()

	// === APPEND WITH DIFFERENT TYPES ===
	demonstrateAppendWithTypes()

	// === MEMORY REALLOCATION ===
	demonstrateMemoryReallocation()

	// === APPEND IN FUNCTIONS ===
	demonstrateAppendInFunctions()

	// === PRE-ALLOCATION ===
	demonstratePreAllocation()

	// === APPEND PATTERNS ===
	demonstrateAppendPatterns()

	// === APPEND WITH INTERFACES ===
	demonstrateAppendWithInterfaces()

	// === APPEND ERROR HANDLING ===
	demonstrateAppendErrorHandling()

	// === APPEND WITH CHANNELS ===
	demonstrateAppendWithChannels()

	// === APPEND PERFORMANCE ===
	demonstrateAppendPerformance()

	// === APPEND BEST PRACTICES ===
	demonstrateAppendBestPractices()

	// === APPEND GOTCHAS ===
	demonstrateAppendGotchas()

	// === APPEND SUMMARY ===
	fmt.Println("\n=== APPEND SUMMARY ===")
	fmt.Println("1. append() returns a new slice")
	fmt.Println("2. Always assign the result of append")
	fmt.Println("3. Capacity doubles when exceeded")
	fmt.Println("4. Use ... to append slices")
	fmt.Println("5. Pre-allocate capacity when size is known")
	fmt.Println("6. append() works with nil slices")
	fmt.Println("7. Be careful with shared underlying arrays")
	fmt.Println("8. Consider memory allocation patterns")
	fmt.Println("9. Use copy() for independent slices")
	fmt.Println("10. append() is the primary way to grow slices")

	fmt.Println("\nThe append function is essential for dynamic data structures in Go!")
}
