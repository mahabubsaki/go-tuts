package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This file covers all types of loops in Go
// Go has only one loop construct: the for loop (but it's very versatile)

func main() {
	fmt.Println("=== GO LOOPS - COMPLETE GUIDE ===")

	demonstrateBasicForLoop()
	demonstrateWhileStyleLoop()
	demonstrateInfiniteLoop()
	demonstrateForRangeLoop()
	demonstrateNestedLoops()
	demonstrateLoopControl()
	demonstrateLoopWithLabels()
	demonstrateLoopBestPractices()
}

func demonstrateBasicForLoop() {
	fmt.Println("\n--- BASIC FOR LOOP ---")

	// Classic for loop with initialization, condition, and increment
	fmt.Println("Counting from 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// For loop counting backwards
	fmt.Println("Counting backwards from 5 to 1:")
	for i := 5; i >= 1; i-- {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// For loop with step size
	fmt.Println("Even numbers from 2 to 10:")
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// For loop with different increment
	fmt.Println("Multiples of 3 from 3 to 15:")
	for i := 3; i <= 15; i += 3 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// For loop with variables
	var start int = 10
	var end int = 20
	var step int = 2

	fmt.Printf("From %d to %d with step %d:\n", start, end, step)
	for i := start; i <= end; i += step {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// For loop with complex condition
	fmt.Println("Powers of 2 less than 100:")
	for i := 1; i < 100; i *= 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// For loop with multiple variables
	fmt.Println("Fibonacci sequence (first 10 numbers):")
	for i, a, b := 0, 0, 1; i < 10; i, a, b = i+1, b, a+b {
		fmt.Printf("%d ", a)
	}
	fmt.Println()

	// For loop with string iteration
	fmt.Println("Character indices in 'Hello':")
	text := "Hello"
	for i := 0; i < len(text); i++ {
		fmt.Printf("Index %d: '%c'\n", i, text[i])
	}

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: for (let i = 0; i < 10; i++) { }
	// JavaScript: Multiple loop types (for, while, do-while, for...in, for...of)
	// Go: Only for loop, but very flexible
	// Go: Multiple variables in for loop initialization
	// Go: := operator for variable declaration in loop
}

func demonstrateWhileStyleLoop() {
	fmt.Println("\n--- WHILE-STYLE LOOP ---")

	// While-style loop (condition only)
	fmt.Println("While-style loop counting to 5:")
	i := 1
	for i <= 5 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()

	// While loop with complex condition
	fmt.Println("Finding first power of 2 greater than 50:")
	power := 1
	for power <= 50 {
		power *= 2
	}
	fmt.Printf("Result: %d\n", power)

	// While loop with user input simulation
	fmt.Println("Simulating user input processing:")
	attempts := 0
	maxAttempts := 3
	success := false

	for attempts < maxAttempts && !success {
		attempts++
		fmt.Printf("Attempt %d: ", attempts)

		// Simulate random success/failure
		if rand.Intn(2) == 1 {
			fmt.Println("Success!")
			success = true
		} else {
			fmt.Println("Failed")
		}
	}

	if success {
		fmt.Printf("Succeeded after %d attempts\n", attempts)
	} else {
		fmt.Printf("Failed after %d attempts\n", attempts)
	}

	// While loop with countdown
	fmt.Println("Countdown:")
	countdown := 5
	for countdown > 0 {
		fmt.Printf("%d... ", countdown)
		countdown--
		time.Sleep(100 * time.Millisecond) // Small delay for effect
	}
	fmt.Println("Go!")

	// While loop with string processing
	fmt.Println("Processing string character by character:")
	message := "Go is awesome!"
	index := 0
	for index < len(message) {
		if message[index] == ' ' {
			fmt.Print("_")
		} else {
			fmt.Printf("%c", message[index])
		}
		index++
	}
	fmt.Println()

	// While loop with accumulator
	fmt.Println("Sum of numbers from 1 to 10:")
	sum := 0
	num := 1
	for num <= 10 {
		sum += num
		num++
	}
	fmt.Printf("Sum: %d\n", sum)

	// While loop with early termination
	fmt.Println("Finding first even number in sequence:")
	sequence := []int{1, 3, 5, 7, 8, 9, 10}
	idx := 0
	for idx < len(sequence) {
		if sequence[idx]%2 == 0 {
			fmt.Printf("Found first even number: %d at index %d\n", sequence[idx], idx)
			break
		}
		idx++
	}

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: while (condition) { }
	// JavaScript: do { } while (condition)
	// Go: for condition { } (no separate while keyword)
	// Go: No do-while equivalent
}

func demonstrateInfiniteLoop() {
	fmt.Println("\n--- INFINITE LOOP ---")

	// Infinite loop with break
	fmt.Println("Infinite loop with break condition:")
	counter := 0
	for {
		counter++
		fmt.Printf("Iteration %d\n", counter)

		if counter >= 5 {
			fmt.Println("Breaking out of infinite loop")
			break
		}
	}

	// Infinite loop with multiple break conditions
	fmt.Println("Infinite loop with multiple break conditions:")
	value := 0
	for {
		value += rand.Intn(10) + 1
		fmt.Printf("Current value: %d\n", value)

		if value > 20 {
			fmt.Println("Value exceeded 20, breaking")
			break
		}

		if value == 15 {
			fmt.Println("Hit exactly 15, breaking")
			break
		}
	}

	// Infinite loop with continue
	fmt.Println("Processing numbers (skipping negative):")
	numbers := []int{1, -2, 3, -4, 5, -6, 7, 8, -9, 10}
	i := 0
	for {
		if i >= len(numbers) {
			break
		}

		if numbers[i] < 0 {
			fmt.Printf("Skipping negative number: %d\n", numbers[i])
			i++
			continue
		}

		fmt.Printf("Processing positive number: %d\n", numbers[i])
		i++
	}

	// Infinite loop with timeout simulation
	fmt.Println("Simulating timeout mechanism:")
	startTime := time.Now()
	timeout := 200 * time.Millisecond

	for {
		// Simulate some work
		time.Sleep(50 * time.Millisecond)

		if time.Since(startTime) > timeout {
			fmt.Println("Timeout reached, breaking")
			break
		}

		fmt.Println("Working...")
	}

	// Server-style infinite loop simulation
	fmt.Println("Server-style loop (running 3 iterations):")
	requests := 0
	for {
		requests++
		fmt.Printf("Handling request #%d\n", requests)

		// Simulate request processing
		time.Sleep(100 * time.Millisecond)

		// Stop after 3 requests for demo
		if requests >= 3 {
			fmt.Println("Demo complete, stopping server")
			break
		}
	}

	// WARNING: Be careful with infinite loops!
	fmt.Println("Note: Always ensure infinite loops have a way to exit!")

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: for (;;) { } or while (true) { }
	// JavaScript: Same break/continue behavior
	// Go: for { } (cleanest syntax)
	// Go: Same break/continue behavior
}

func demonstrateForRangeLoop() {
	fmt.Println("\n--- FOR RANGE LOOP ---")

	// Range over slice
	fmt.Println("Ranging over slice:")
	numbers := []int{10, 20, 30, 40, 50}

	for index, value := range numbers {
		fmt.Printf("Index %d: %d\n", index, value)
	}

	// Range over slice (value only)
	fmt.Println("Ranging over slice (values only):")
	for _, value := range numbers {
		fmt.Printf("%d ", value)
	}
	fmt.Println()

	// Range over slice (index only)
	fmt.Println("Ranging over slice (indices only):")
	for index := range numbers {
		fmt.Printf("Index %d ", index)
	}
	fmt.Println()

	// Range over array
	fmt.Println("Ranging over array:")
	colors := [4]string{"red", "green", "blue", "yellow"}

	for i, color := range colors {
		fmt.Printf("Color %d: %s\n", i, color)
	}

	// Range over string
	fmt.Println("Ranging over string:")
	text := "Hello, 世界"

	for index, runeValue := range text {
		fmt.Printf("Index %d: '%c' (Unicode: %d)\n", index, runeValue, runeValue)
	}

	// Range over map
	fmt.Println("Ranging over map:")
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 35,
	}

	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// Range over map (keys only)
	fmt.Println("Ranging over map (keys only):")
	for name := range ages {
		fmt.Printf("Name: %s\n", name)
	}

	// Range over map (values only)
	fmt.Println("Ranging over map (values only):")
	for _, age := range ages {
		fmt.Printf("Age: %d\n", age)
	}

	// Range over channel
	fmt.Println("Ranging over channel:")
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}

	// Range with modification
	fmt.Println("Modifying slice during range:")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original: %v\n", nums)

	for i, value := range nums {
		nums[i] = value * 2
	}
	fmt.Printf("Modified: %v\n", nums)

	// Range with filtering
	fmt.Println("Filtering with range:")
	mixedNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evenNumbers []int

	for _, num := range mixedNumbers {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	fmt.Printf("Even numbers: %v\n", evenNumbers)

	// Range with struct slice
	fmt.Println("Ranging over struct slice:")
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Carol", Age: 35},
	}

	for i, person := range people {
		fmt.Printf("Person %d: %s (%d years)\n", i+1, person.Name, person.Age)
	}

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: for...of for values, for...in for keys
	// JavaScript: forEach method for arrays
	// Go: for...range works with arrays, slices, maps, strings, channels
	// Go: Returns index/key and value
	// Go: Can ignore index with _
}

func demonstrateNestedLoops() {
	fmt.Println("\n--- NESTED LOOPS ---")

	// Simple nested loop
	fmt.Println("Multiplication table (5x5):")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("%2d ", i*j)
		}
		fmt.Println()
	}

	// Nested loop with 2D array
	fmt.Println("2D array traversal:")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for i, row := range matrix {
		for j, value := range row {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, value)
		}
	}

	// Nested loop with pattern printing
	fmt.Println("Star pattern:")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// Nested loop with break
	fmt.Println("Finding target in 2D array:")
	target := 5
	found := false

	for i := 0; i < len(matrix) && !found; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				fmt.Printf("Found %d at position [%d][%d]\n", target, i, j)
				found = true
				break
			}
		}
	}

	// Nested loop with continue
	fmt.Println("Skipping diagonal elements:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == j {
				continue // Skip diagonal elements
			}
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
		}
	}

	// Nested range loops
	fmt.Println("Nested range over maps:")
	departments := map[string][]string{
		"Engineering": {"Alice", "Bob", "Charlie"},
		"Marketing":   {"David", "Eve"},
		"Sales":       {"Frank", "Grace", "Henry"},
	}

	for dept, employees := range departments {
		fmt.Printf("%s Department:\n", dept)
		for i, employee := range employees {
			fmt.Printf("  %d. %s\n", i+1, employee)
		}
	}

	// Complex nested loop example
	fmt.Println("Finding all pairs that sum to 10:")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == 10 {
				fmt.Printf("Pair: %d + %d = 10\n", nums[i], nums[j])
			}
		}
	}

	// Nested loop with different data types
	fmt.Println("Processing key-value pairs:")
	data := map[string][]int{
		"group1": {1, 2, 3},
		"group2": {4, 5, 6},
		"group3": {7, 8, 9},
	}

	for groupName, values := range data {
		fmt.Printf("Processing %s:\n", groupName)
		for _, value := range values {
			fmt.Printf("  Value: %d, Square: %d\n", value, value*value)
		}
	}
}

func demonstrateLoopControl() {
	fmt.Println("\n--- LOOP CONTROL (BREAK, CONTINUE) ---")

	// Break statement
	fmt.Println("Using break to exit loop:")
	for i := 1; i <= 10; i++ {
		if i == 6 {
			fmt.Printf("Breaking at %d\n", i)
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Continue statement
	fmt.Println("Using continue to skip iterations:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Break in range loop
	fmt.Println("Break in range loop:")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, num := range numbers {
		if num > 5 {
			fmt.Printf("Stopping at index %d, value %d\n", i, num)
			break
		}
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	// Continue in range loop
	fmt.Println("Continue in range loop (skip negative):")
	mixedNumbers := []int{1, -2, 3, -4, 5, -6, 7, -8, 9, -10}

	for _, num := range mixedNumbers {
		if num < 0 {
			continue
		}
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	// Multiple continue conditions
	fmt.Println("Multiple continue conditions:")
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		if i%3 == 0 {
			continue // Skip multiples of 3
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Break with search
	fmt.Println("Search with break:")
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	searchName := "Charlie"

	for i, name := range names {
		if name == searchName {
			fmt.Printf("Found %s at index %d\n", searchName, i)
			break
		}
	}

	// Continue with validation
	fmt.Println("Validation with continue:")
	inputs := []string{"123", "abc", "456", "def", "789"}

	for _, input := range inputs {
		if len(input) != 3 {
			fmt.Printf("Skipping invalid input: %s\n", input)
			continue
		}

		// Check if all characters are digits
		isValid := true
		for _, char := range input {
			if char < '0' || char > '9' {
				isValid = false
				break
			}
		}

		if !isValid {
			fmt.Printf("Skipping non-numeric input: %s\n", input)
			continue
		}

		fmt.Printf("Valid input: %s\n", input)
	}

	// Break from nested loop (wrong way)
	fmt.Println("Break from nested loop (breaks only inner loop):")
	for i := 1; i <= 3; i++ {
		fmt.Printf("Outer loop: %d\n", i)
		for j := 1; j <= 3; j++ {
			if j == 2 {
				fmt.Println("  Breaking inner loop")
				break
			}
			fmt.Printf("  Inner loop: %d\n", j)
		}
	}

	// Using flag for nested loop control
	fmt.Println("Using flag for nested loop control:")
	found := false
	for i := 1; i <= 3 && !found; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				fmt.Printf("Found target at [%d][%d]\n", i, j)
				found = true
				break
			}
		}
	}

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Same break and continue behavior
	// JavaScript: break/continue work with labels
	// Go: Same break and continue behavior
	// Go: break/continue work with labels (covered in next section)
}

func demonstrateLoopWithLabels() {
	fmt.Println("\n--- LOOP WITH LABELS ---")

	// Basic label usage
	fmt.Println("Breaking from nested loop with labels:")

OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				fmt.Printf("Breaking from outer loop at [%d][%d]\n", i, j)
				break OuterLoop
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	// Continue with label
	fmt.Println("Continue with label:")

OuterContinue:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				fmt.Printf("Continuing outer loop at [%d][%d]\n", i, j)
				continue OuterContinue
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	// Complex label example
	fmt.Println("Complex label example - finding prime numbers:")

FindPrimes:
	for num := 2; num <= 20; num++ {
		for divisor := 2; divisor < num; divisor++ {
			if num%divisor == 0 {
				continue FindPrimes // Not prime, check next number
			}
		}
		fmt.Printf("%d is prime\n", num)
	}

	// Label with range
	fmt.Println("Label with range loop:")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

SearchMatrix:
	for i, row := range matrix {
		for j, value := range row {
			if value == 5 {
				fmt.Printf("Found 5 at [%d][%d]\n", i, j)
				break SearchMatrix
			}
		}
	}

	// Multiple labels
	fmt.Println("Multiple labels example:")

Level1:
	for i := 1; i <= 2; i++ {
	Level2:
		for j := 1; j <= 2; j++ {
			for k := 1; k <= 2; k++ {
				if i == 1 && j == 2 && k == 1 {
					fmt.Printf("Breaking to Level1 at [%d][%d][%d]\n", i, j, k)
					break Level1
				}
				if i == 2 && j == 1 && k == 2 {
					fmt.Printf("Continuing Level2 at [%d][%d][%d]\n", i, j, k)
					continue Level2
				}
				fmt.Printf("i=%d, j=%d, k=%d\n", i, j, k)
			}
		}
	}

	// Practical label example
	fmt.Println("Practical example - processing data with error handling:")

	data := [][]string{
		{"valid", "data", "here"},
		{"more", "valid", "data"},
		{"error", "invalid", "data"},
		{"never", "processed", "data"},
	}

ProcessData:
	for i, row := range data {
		for j, item := range row {
			if item == "error" {
				fmt.Printf("Error found at [%d][%d], stopping all processing\n", i, j)
				break ProcessData
			}
			fmt.Printf("Processing item [%d][%d]: %s\n", i, j, item)
		}
	}

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Labels work similarly (label: statement)
	// JavaScript: Can use labels with break/continue
	// Go: Labels work similarly
	// Go: Labels must be followed by loop statements
	// Go: More commonly used than in JavaScript
}

func demonstrateLoopBestPractices() {
	fmt.Println("\n--- LOOP BEST PRACTICES ---")

	// 1. Use range when possible
	fmt.Println("1. Use range when possible:")

	// Good: using range
	numbers := []int{1, 2, 3, 4, 5}
	for i, num := range numbers {
		fmt.Printf("Index %d: %d\n", i, num)
	}

	// 2. Use descriptive variable names
	fmt.Println("2. Use descriptive variable names:")

	// Good: descriptive names
	students := []string{"Alice", "Bob", "Charlie"}
	for studentIndex, studentName := range students {
		fmt.Printf("Student %d: %s\n", studentIndex+1, studentName)
	}

	// 3. Avoid infinite loops without exit conditions
	fmt.Println("3. Avoid infinite loops without exit conditions:")

	// Good: clear exit condition
	attempts := 0
	maxAttempts := 5
	for {
		attempts++
		if attempts > maxAttempts {
			fmt.Println("Max attempts reached")
			break
		}
		fmt.Printf("Attempt %d\n", attempts)
		// Simulate some condition that might succeed
		if attempts == 3 {
			fmt.Println("Success!")
			break
		}
	}

	// 4. Use labels for complex nested loops
	fmt.Println("4. Use labels for complex nested loops:")

	// Good: clear label usage
ProcessRows:
	for rowIndex := 0; rowIndex < 3; rowIndex++ {
		for colIndex := 0; colIndex < 3; colIndex++ {
			if rowIndex == 1 && colIndex == 1 {
				fmt.Println("Found center position, stopping all processing")
				break ProcessRows
			}
			fmt.Printf("Processing [%d][%d]\n", rowIndex, colIndex)
		}
	}

	// 5. Pre-allocate slices when size is known
	fmt.Println("5. Pre-allocate slices when size is known:")

	// Good: pre-allocated slice
	sourceData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := make([]int, 0, len(sourceData)/2) // Pre-allocate capacity

	for _, num := range sourceData {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	fmt.Printf("Even numbers: %v\n", evenNumbers)

	// 6. Use continue to reduce nesting
	fmt.Println("6. Use continue to reduce nesting:")

	// Good: using continue
	for _, num := range sourceData {
		if num%2 != 0 {
			continue // Skip odd numbers
		}
		if num < 5 {
			continue // Skip small numbers
		}
		fmt.Printf("Processing large even number: %d\n", num)
	}

	// 7. Extract complex loop bodies into functions
	fmt.Println("7. Extract complex loop bodies into functions:")

	users := []User{
		{Name: "Alice", Age: 25, Active: true},
		{Name: "Bob", Age: 30, Active: false},
		{Name: "Charlie", Age: 35, Active: true},
	}

	for _, user := range users {
		processUser(user)
	}

	// 8. Use appropriate loop type
	fmt.Println("8. Use appropriate loop type:")

	// Good: for-range for slices
	colors := []string{"red", "green", "blue"}
	for _, color := range colors {
		fmt.Printf("Color: %s\n", color)
	}

	// Good: traditional for loop for numeric ranges
	for i := 0; i < 5; i++ {
		fmt.Printf("Count: %d\n", i)
	}

	// 9. Handle empty collections
	fmt.Println("9. Handle empty collections:")

	var emptySlice []int
	if len(emptySlice) == 0 {
		fmt.Println("Empty slice, nothing to process")
	} else {
		for _, value := range emptySlice {
			fmt.Printf("Value: %d\n", value)
		}
	}

	// 10. Use defer for cleanup in loops
	fmt.Println("10. Use defer for cleanup in loops:")

	fmt.Println("Simulating file processing with cleanup:")
	files := []string{"file1.txt", "file2.txt", "file3.txt"}

	for _, filename := range files {
		processFile(filename)
	}

	// BEST PRACTICES SUMMARY:
	fmt.Println("\nBest Practices Summary:")
	fmt.Println("1. Use range when iterating over collections")
	fmt.Println("2. Use descriptive variable names")
	fmt.Println("3. Avoid infinite loops without clear exit conditions")
	fmt.Println("4. Use labels for complex nested loop control")
	fmt.Println("5. Pre-allocate slices when size is known")
	fmt.Println("6. Use continue to reduce nesting")
	fmt.Println("7. Extract complex loop bodies into functions")
	fmt.Println("8. Choose appropriate loop type for the task")
	fmt.Println("9. Handle empty collections gracefully")
	fmt.Println("10. Use defer for cleanup when needed")
}

// Helper types and functions

type User struct {
	Name   string
	Age    int
	Active bool
}

func processUser(user User) {
	fmt.Printf("Processing user: %s", user.Name)
	if !user.Active {
		fmt.Print(" (inactive)")
	}
	fmt.Printf(" - Age: %d\n", user.Age)
}

func processFile(filename string) {
	fmt.Printf("Opening file: %s\n", filename)

	// Simulate file processing
	defer func() {
		fmt.Printf("Closing file: %s\n", filename)
	}()

	// Simulate some processing
	fmt.Printf("Processing file: %s\n", filename)
}
