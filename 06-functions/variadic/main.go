package main

import (
	"fmt"
	"math"
)

// === VARIADIC FUNCTIONS ===

// 1. Basic variadic function
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 2. Variadic function with regular parameters
func greetAll(greeting string, names ...string) {
	for _, name := range names {
		fmt.Printf("%s, %s!\n", greeting, name)
	}
}

// 3. Variadic function with multiple return values
func findMinMax(numbers ...int) (int, int, bool) {
	if len(numbers) == 0 {
		return 0, 0, false
	}

	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers[1:] {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max, true
}

// 4. Variadic function with different types
func printAll(items ...interface{}) {
	for i, item := range items {
		fmt.Printf("Item %d: %v (%T)\n", i+1, item, item)
	}
}

// 5. Variadic function for string operations
func concatenate(separator string, strings ...string) string {
	if len(strings) == 0 {
		return ""
	}

	result := strings[0]
	for _, str := range strings[1:] {
		result += separator + str
	}
	return result
}

// 6. Variadic function with validation
func average(numbers ...float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("cannot calculate average of empty set")
	}

	total := 0.0
	for _, num := range numbers {
		total += num
	}

	return total / float64(len(numbers)), nil
}

// 7. Variadic function for mathematical operations
func multiply(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]
	for _, num := range numbers[1:] {
		result *= num
	}
	return result
}

// 8. Variadic function with slices
func processLists(lists ...[]int) []int {
	var result []int
	for _, list := range lists {
		result = append(result, list...)
	}
	return result
}

// 9. Variadic function with maps
func mergeMaps(maps ...map[string]int) map[string]int {
	result := make(map[string]int)
	for _, m := range maps {
		for key, value := range m {
			result[key] = value
		}
	}
	return result
}

// 10. Variadic function with structs
type Person struct {
	Name string
	Age  int
}

func createTeam(teamName string, members ...Person) {
	fmt.Printf("Team: %s\n", teamName)
	fmt.Printf("Members:\n")
	for i, member := range members {
		fmt.Printf("  %d. %s (Age: %d)\n", i+1, member.Name, member.Age)
	}
}

// 11. Variadic function with functions
func applyFunctions(value int, functions ...func(int) int) []int {
	results := make([]int, len(functions))
	for i, fn := range functions {
		results[i] = fn(value)
	}
	return results
}

// 12. Variadic function with error handling
func divideAll(dividend float64, divisors ...float64) ([]float64, error) {
	if len(divisors) == 0 {
		return nil, fmt.Errorf("no divisors provided")
	}

	results := make([]float64, len(divisors))
	for i, divisor := range divisors {
		if divisor == 0 {
			return nil, fmt.Errorf("division by zero at index %d", i)
		}
		results[i] = dividend / divisor
	}

	return results, nil
}

// 14. Variadic function with custom types
type Color struct {
	R, G, B int
}

func mixColors(colors ...Color) Color {
	if len(colors) == 0 {
		return Color{0, 0, 0}
	}

	totalR, totalG, totalB := 0, 0, 0
	for _, color := range colors {
		totalR += color.R
		totalG += color.G
		totalB += color.B
	}

	count := len(colors)
	return Color{
		R: totalR / count,
		G: totalG / count,
		B: totalB / count,
	}
}

// 15. Variadic function with filtering
func filterPositive(numbers ...int) []int {
	var result []int
	for _, num := range numbers {
		if num > 0 {
			result = append(result, num)
		}
	}
	return result
}

// Helper function for formatting numbers
func formatNumbers(precision int, numbers ...float64) []string {
	results := make([]string, len(numbers))
	format := fmt.Sprintf("%%.%df", precision)

	for i, num := range numbers {
		results[i] = fmt.Sprintf(format, num)
	}

	return results
}

func main() {
	fmt.Println("=== GO VARIADIC FUNCTIONS ===")

	// === BASIC VARIADIC FUNCTION ===
	fmt.Println("\n--- BASIC VARIADIC FUNCTION ---")
	fmt.Printf("Sum of 1, 2, 3: %d\n", sum(1, 2, 3))
	fmt.Printf("Sum of 1, 2, 3, 4, 5: %d\n", sum(1, 2, 3, 4, 5))
	fmt.Printf("Sum of no numbers: %d\n", sum())

	// Using slice with variadic function
	numbers := []int{10, 20, 30, 40}
	fmt.Printf("Sum of slice %v: %d\n", numbers, sum(numbers...))

	/*
		JavaScript comparison:
		function sum(...numbers) {
			return numbers.reduce((total, num) => total + num, 0);
		}

		console.log(sum(1, 2, 3));
		console.log(sum(...[10, 20, 30, 40]));
	*/

	// === VARIADIC WITH REGULAR PARAMETERS ===
	fmt.Println("\n--- VARIADIC WITH REGULAR PARAMETERS ---")
	greetAll("Hello", "Alice", "Bob", "Charlie")
	greetAll("Hi", "David")
	greetAll("Welcome") // No names provided

	// === VARIADIC WITH MULTIPLE RETURN VALUES ===
	fmt.Println("\n--- VARIADIC WITH MULTIPLE RETURN VALUES ---")
	values := []int{5, 2, 8, 1, 9, 3}
	min, max, found := findMinMax(values...)
	if found {
		fmt.Printf("Min: %d, Max: %d\n", min, max)
	}

	// Empty case
	min, max, found = findMinMax()
	if !found {
		fmt.Println("No values provided")
	}

	// === VARIADIC WITH DIFFERENT TYPES ===
	fmt.Println("\n--- VARIADIC WITH DIFFERENT TYPES ---")
	printAll(42, "hello", 3.14, true, []int{1, 2, 3})

	// === STRING OPERATIONS ===
	fmt.Println("\n--- STRING OPERATIONS ---")
	result := concatenate("-", "apple", "banana", "cherry")
	fmt.Printf("Concatenated: %s\n", result)

	result = concatenate(" | ", "first", "second", "third", "fourth")
	fmt.Printf("Pipe separated: %s\n", result)

	// === VARIADIC WITH VALIDATION ===
	fmt.Println("\n--- VARIADIC WITH VALIDATION ---")
	avg, err := average(1.5, 2.5, 3.5, 4.5)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Average: %.2f\n", avg)
	}

	// Error case
	avg, err = average()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// === MATHEMATICAL OPERATIONS ===
	fmt.Println("\n--- MATHEMATICAL OPERATIONS ---")
	product := multiply(2, 3, 4, 5)
	fmt.Printf("Product of 2, 3, 4, 5: %.2f\n", product)

	product = multiply(1.5, 2.0, 3.0)
	fmt.Printf("Product of 1.5, 2.0, 3.0: %.2f\n", product)

	// === PROCESSING SLICES ===
	fmt.Println("\n--- PROCESSING SLICES ---")
	list1 := []int{1, 2, 3}
	list2 := []int{4, 5, 6}
	list3 := []int{7, 8, 9}

	combined := processLists(list1, list2, list3)
	fmt.Printf("Combined lists: %v\n", combined)

	// === MERGING MAPS ===
	fmt.Println("\n--- MERGING MAPS ---")
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"c": 3, "d": 4}
	map3 := map[string]int{"e": 5, "f": 6}

	merged := mergeMaps(map1, map2, map3)
	fmt.Printf("Merged map: %v\n", merged)

	// === VARIADIC WITH STRUCTS ===
	fmt.Println("\n--- VARIADIC WITH STRUCTS ---")
	alice := Person{Name: "Alice", Age: 25}
	bob := Person{Name: "Bob", Age: 30}
	charlie := Person{Name: "Charlie", Age: 35}

	createTeam("Development", alice, bob, charlie)

	// === VARIADIC WITH FUNCTIONS ===
	fmt.Println("\n--- VARIADIC WITH FUNCTIONS ---")
	square := func(x int) int { return x * x }
	double := func(x int) int { return x * 2 }
	increment := func(x int) int { return x + 1 }

	results := applyFunctions(5, square, double, increment)
	fmt.Printf("Results for 5: %v\n", results)

	/*
		JavaScript comparison:
		function applyFunctions(value, ...functions) {
			return functions.map(fn => fn(value));
		}

		const square = x => x * x;
		const double = x => x * 2;
		const increment = x => x + 1;

		const results = applyFunctions(5, square, double, increment);
	*/

	// === VARIADIC WITH ERROR HANDLING ===
	fmt.Println("\n--- VARIADIC WITH ERROR HANDLING ---")
	quotients, err := divideAll(100, 2, 5, 10, 20)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Quotients: %v\n", quotients)
	}

	// Error case
	quotients, err = divideAll(100, 2, 0, 10)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// === STATISTICAL CALCULATIONS ===
	fmt.Println("\n--- STATISTICAL CALCULATIONS ---")
	data := []float64{1.5, 2.3, 3.7, 2.1, 4.2, 1.8, 3.9}
	fmt.Printf("Data: %v\n", data)

	// Simple statistics calculation
	count := len(data)
	var total float64
	for _, value := range data {
		total += value
	}
	average := total / float64(count)

	fmt.Printf("Count: %d, Sum: %.2f, Average: %.2f\n", count, total, average)

	// === CUSTOM TYPES ===
	fmt.Println("\n--- CUSTOM TYPES ---")
	red := Color{255, 0, 0}
	green := Color{0, 255, 0}
	blue := Color{0, 0, 255}

	mixedColor := mixColors(red, green, blue)
	fmt.Printf("Mixed color RGB: (%d, %d, %d)\n", mixedColor.R, mixedColor.G, mixedColor.B)

	// === FILTERING ===
	fmt.Println("\n--- FILTERING ---")
	mixedNumbers := []int{-5, 3, -2, 8, -1, 6, 0, 9}
	positive := filterPositive(mixedNumbers...)
	fmt.Printf("Original: %v\n", mixedNumbers)
	fmt.Printf("Positive only: %v\n", positive)

	// === VARIADIC PATTERNS ===
	fmt.Println("\n--- VARIADIC PATTERNS ---")

	// 1. Unpacking slices
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	totalSum := sum(append(slice1, slice2...)...)
	fmt.Printf("Sum of two slices: %d\n", totalSum)

	// 2. Combining individual values and slices
	moreNumbers := []int{7, 8, 9}
	combinedSum := sum(1, 2, 3) + sum(moreNumbers...)
	fmt.Printf("Combined sum: %d\n", combinedSum)

	// 3. Empty variadic calls
	emptySum := sum()
	fmt.Printf("Empty sum: %d\n", emptySum)

	// 4. Single value
	singleSum := sum(42)
	fmt.Printf("Single value sum: %d\n", singleSum)

	// === FORMATTING WITH VARIADIC ===
	fmt.Println("\n--- FORMATTING WITH VARIADIC ---")
	floats := []float64{3.14159, 2.71828, 1.41421}
	formatted := formatNumbers(2, floats...)
	fmt.Printf("Formatted to 2 decimal places: %v\n", formatted)

	formatted = formatNumbers(4, math.Pi, math.E, math.Sqrt2)
	fmt.Printf("Formatted to 4 decimal places: %v\n", formatted)

	// === VARIADIC LIMITATIONS ===
	fmt.Println("\n--- VARIADIC LIMITATIONS ---")
	fmt.Println("1. Variadic parameter must be the last parameter")
	fmt.Println("2. Only one variadic parameter per function")
	fmt.Println("3. Cannot have variadic parameter of variadic type")
	fmt.Println("4. Variadic parameters are converted to slices internally")

	// === BEST PRACTICES ===
	fmt.Println("\n--- VARIADIC BEST PRACTICES ---")
	fmt.Println("1. Use variadic functions for flexible argument lists")
	fmt.Println("2. Validate variadic parameters when necessary")
	fmt.Println("3. Consider performance with large variadic arguments")
	fmt.Println("4. Use meaningful parameter names")
	fmt.Println("5. Document expected parameter types and constraints")
	fmt.Println("6. Handle empty parameter lists appropriately")
	fmt.Println("7. Use type-safe variadic functions when possible")
	fmt.Println("8. Consider using slices instead for complex cases")
	fmt.Println("9. Combine regular and variadic parameters thoughtfully")
	fmt.Println("10. Test edge cases with no arguments")
}
