package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// === RETURN VALUES ===

// 1. No return value (void function)
func printMessage(message string) {
	fmt.Println(message)
}

// 2. Single return value
func square(n int) int {
	return n * n
}

// 3. Multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// 4. Named return values
func calculateStats(numbers []int) (sum int, count int, average float64) {
	count = len(numbers)
	if count == 0 {
		return // Returns zero values
	}

	for _, num := range numbers {
		sum += num
	}
	average = float64(sum) / float64(count)
	return // Returns named values
}

// 5. Named return values with explicit return
func findMinMax(numbers []int) (min, max int, found bool) {
	if len(numbers) == 0 {
		return 0, 0, false
	}

	min = numbers[0]
	max = numbers[0]

	for _, num := range numbers[1:] {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	found = true
	return min, max, found
}

// 6. Multiple return values with different types
func parseUserInput(input string) (int, string, bool, error) {
	parts := strings.Split(input, ",")
	if len(parts) != 3 {
		return 0, "", false, fmt.Errorf("invalid input format")
	}

	id, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, "", false, fmt.Errorf("invalid ID: %v", err)
	}

	name := strings.TrimSpace(parts[1])
	active, err := strconv.ParseBool(strings.TrimSpace(parts[2]))
	if err != nil {
		return 0, "", false, fmt.Errorf("invalid active status: %v", err)
	}

	return id, name, active, nil
}

// 7. Returning slices
func generateSequence(start, end int) []int {
	if start > end {
		return []int{} // Return empty slice
	}

	result := make([]int, 0, end-start+1)
	for i := start; i <= end; i++ {
		result = append(result, i)
	}
	return result
}

// 8. Returning maps
func groupByLength(words []string) map[int][]string {
	result := make(map[int][]string)
	for _, word := range words {
		length := len(word)
		result[length] = append(result[length], word)
	}
	return result
}

// 9. Returning structs
type Point struct {
	X, Y float64
}

func createPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

func createPointPtr(x, y float64) *Point {
	return &Point{X: x, Y: y}
}

// 10. Returning interfaces
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func createShape(shapeType string, params ...float64) Shape {
	switch shapeType {
	case "rectangle":
		if len(params) >= 2 {
			return Rectangle{Width: params[0], Height: params[1]}
		}
	case "circle":
		if len(params) >= 1 {
			return Circle{Radius: params[0]}
		}
	}
	return nil
}

// 11. Returning functions
func createMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func createValidator(minLength int) func(string) bool {
	return func(s string) bool {
		return len(s) >= minLength
	}
}

// 12. Returning channels
func createNumberChannel(numbers []int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, num := range numbers {
			ch <- num
		}
	}()
	return ch
}

// 13. Multiple return values with blank identifier
func processData(data []int) ([]int, []int) {
	var evens, odds []int
	for _, num := range data {
		if num%2 == 0 {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}
	}
	return evens, odds
}

// 14. Named return values with modification
func factorial(n int) (result int) {
	result = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return
}

// 15. Multiple error returns
func connectToDatabase(host string, port int) (connection string, warning string, err error) {
	if host == "" {
		return "", "", fmt.Errorf("host cannot be empty")
	}

	if port < 1 || port > 65535 {
		return "", "", fmt.Errorf("invalid port number: %d", port)
	}

	connection = fmt.Sprintf("connected to %s:%d", host, port)

	if port == 80 || port == 443 {
		warning = "using standard HTTP/HTTPS port"
	}

	return connection, warning, nil
}

// 16. Returning result and metadata
type ProcessResult struct {
	Data      []int
	Count     int
	Duration  float64
	Processed bool
}

func processNumbers(numbers []int) ProcessResult {
	result := ProcessResult{
		Data:      make([]int, 0, len(numbers)),
		Count:     len(numbers),
		Duration:  0.001, // Simulated processing time
		Processed: true,
	}

	for _, num := range numbers {
		result.Data = append(result.Data, num*2)
	}

	return result
}

// 17. Conditional returns
func safeIndex(slice []int, index int) (int, bool) {
	if index < 0 || index >= len(slice) {
		return 0, false
	}
	return slice[index], true
}

// 18. Early returns for validation
func validateUser(name string, age int, email string) error {
	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if age < 0 || age > 150 {
		return fmt.Errorf("invalid age: %d", age)
	}

	if !strings.Contains(email, "@") {
		return fmt.Errorf("invalid email format")
	}

	return nil
}

// 19. Returning pointers vs values
func createLargeStruct() *ProcessResult {
	// Return pointer to avoid copying large struct
	return &ProcessResult{
		Data:      make([]int, 1000),
		Count:     1000,
		Duration:  1.5,
		Processed: true,
	}
}

func createSmallStruct() Point {
	// Return value for small structs
	return Point{X: 1.0, Y: 2.0}
}

// 20. Chaining function returns
func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func calculate(x, y int) int {
	return multiply(add(x, y), 2)
}

func main() {
	fmt.Println("=== GO FUNCTIONS - RETURN VALUES ===")

	// === NO RETURN VALUE ===
	fmt.Println("\n--- NO RETURN VALUE ---")
	printMessage("This function returns nothing")

	// === SINGLE RETURN VALUE ===
	fmt.Println("\n--- SINGLE RETURN VALUE ---")
	result := square(5)
	fmt.Printf("Square of 5: %d\n", result)

	// Direct use in expression
	fmt.Printf("Square of 4 + 3: %d\n", square(4)+3)

	/*
		JavaScript comparison:
		function square(n) {
			return n * n;
		}

		const result = square(5);
		console.log(`Square of 5: ${result}`);
	*/

	// === MULTIPLE RETURN VALUES ===
	fmt.Println("\n--- MULTIPLE RETURN VALUES ---")
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", quotient)
	}

	// Error case
	quotient, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 0 = %.2f\n", quotient)
	}

	/*
		JavaScript comparison (using objects or arrays):
		function divide(a, b) {
			if (b === 0) {
				return { result: 0, error: "Division by zero" };
			}
			return { result: a / b, error: null };
		}

		const { result, error } = divide(10, 2);
	*/

	// === NAMED RETURN VALUES ===
	fmt.Println("\n--- NAMED RETURN VALUES ---")
	numbers := []int{1, 2, 3, 4, 5}
	sum, count, average := calculateStats(numbers)
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Sum: %d, Count: %d, Average: %.2f\n", sum, count, average)

	// Empty slice
	emptyNumbers := []int{}
	sum, count, average = calculateStats(emptyNumbers)
	fmt.Printf("Empty slice - Sum: %d, Count: %d, Average: %.2f\n", sum, count, average)

	// Find min/max
	values := []int{3, 1, 4, 1, 5, 9, 2, 6}
	min, max, found := findMinMax(values)
	if found {
		fmt.Printf("Min: %d, Max: %d\n", min, max)
	} else {
		fmt.Println("No values found")
	}

	// === MULTIPLE RETURN VALUES WITH DIFFERENT TYPES ===
	fmt.Println("\n--- MULTIPLE RETURN VALUES WITH DIFFERENT TYPES ---")
	input := "123, Alice, true"
	id, name, active, err := parseUserInput(input)
	if err != nil {
		fmt.Printf("Error parsing input: %v\n", err)
	} else {
		fmt.Printf("ID: %d, Name: %s, Active: %t\n", id, name, active)
	}

	// Invalid input
	invalidInput := "invalid,input"
	id, name, active, err = parseUserInput(invalidInput)
	if err != nil {
		fmt.Printf("Error parsing input: %v\n", err)
	}

	// === RETURNING SLICES ===
	fmt.Println("\n--- RETURNING SLICES ---")
	sequence := generateSequence(1, 5)
	fmt.Printf("Sequence 1-5: %v\n", sequence)

	emptySequence := generateSequence(5, 1)
	fmt.Printf("Invalid sequence: %v\n", emptySequence)

	// === RETURNING MAPS ===
	fmt.Println("\n--- RETURNING MAPS ---")
	words := []string{"cat", "dog", "bird", "fish", "elephant"}
	grouped := groupByLength(words)
	fmt.Println("Words grouped by length:")
	for length, wordList := range grouped {
		fmt.Printf("Length %d: %v\n", length, wordList)
	}

	// === RETURNING STRUCTS ===
	fmt.Println("\n--- RETURNING STRUCTS ---")
	point1 := createPoint(3.5, 4.2)
	fmt.Printf("Point (value): %v\n", point1)

	point2 := createPointPtr(1.5, 2.8)
	fmt.Printf("Point (pointer): %v\n", point2)

	// === RETURNING INTERFACES ===
	fmt.Println("\n--- RETURNING INTERFACES ---")
	rect := createShape("rectangle", 5, 3)
	if rect != nil {
		fmt.Printf("Rectangle area: %.2f\n", rect.Area())
	}

	circle := createShape("circle", 4)
	if circle != nil {
		fmt.Printf("Circle area: %.2f\n", circle.Area())
	}

	invalid := createShape("triangle", 1, 2, 3)
	if invalid == nil {
		fmt.Println("Invalid shape type")
	}

	// === RETURNING FUNCTIONS ===
	fmt.Println("\n--- RETURNING FUNCTIONS ---")
	double := createMultiplier(2)
	triple := createMultiplier(3)

	fmt.Printf("Double 5: %d\n", double(5))
	fmt.Printf("Triple 5: %d\n", triple(5))

	// Validator functions
	isLongEnough := createValidator(5)
	fmt.Printf("'Hello' is long enough: %t\n", isLongEnough("Hello"))
	fmt.Printf("'Hi' is long enough: %t\n", isLongEnough("Hi"))

	/*
		JavaScript comparison:
		function createMultiplier(factor) {
			return function(n) {
				return n * factor;
			};
		}

		const double = createMultiplier(2);
		const triple = createMultiplier(3);
	*/

	// === RETURNING CHANNELS ===
	fmt.Println("\n--- RETURNING CHANNELS ---")
	channelNumbers := []int{1, 2, 3, 4, 5}
	ch := createNumberChannel(channelNumbers)

	fmt.Print("Channel values: ")
	for num := range ch {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	// === USING BLANK IDENTIFIER ===
	fmt.Println("\n--- USING BLANK IDENTIFIER ---")
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens, odds := processData(data)
	fmt.Printf("Evens: %v\n", evens)
	fmt.Printf("Odds: %v\n", odds)

	// Using blank identifier to ignore one return value
	evens, _ = processData(data)
	fmt.Printf("Only evens: %v\n", evens)

	// === NAMED RETURN VALUES WITH MODIFICATION ===
	fmt.Println("\n--- NAMED RETURN VALUES WITH MODIFICATION ---")
	fact := factorial(5)
	fmt.Printf("Factorial of 5: %d\n", fact)

	// === MULTIPLE ERROR RETURNS ===
	fmt.Println("\n--- MULTIPLE ERROR RETURNS ---")
	conn, warning, err := connectToDatabase("localhost", 80)
	if err != nil {
		fmt.Printf("Connection error: %v\n", err)
	} else {
		fmt.Printf("Connection: %s\n", conn)
		if warning != "" {
			fmt.Printf("Warning: %s\n", warning)
		}
	}

	// Error case
	conn, warning, err = connectToDatabase("", 3306)
	if err != nil {
		fmt.Printf("Connection error: %v\n", err)
	}

	// === RETURNING RESULT AND METADATA ===
	fmt.Println("\n--- RETURNING RESULT AND METADATA ---")
	inputData := []int{1, 2, 3, 4, 5}
	processResult := processNumbers(inputData)
	fmt.Printf("Original: %v\n", inputData)
	fmt.Printf("Processed: %v\n", processResult.Data)
	fmt.Printf("Count: %d, Duration: %.3fs, Success: %t\n",
		processResult.Count, processResult.Duration, processResult.Processed)

	// === CONDITIONAL RETURNS ===
	fmt.Println("\n--- CONDITIONAL RETURNS ---")
	slice := []int{10, 20, 30, 40, 50}

	value, ok := safeIndex(slice, 2)
	if ok {
		fmt.Printf("Value at index 2: %d\n", value)
	}

	value, ok = safeIndex(slice, 10)
	if !ok {
		fmt.Printf("Index 10 is out of bounds\n")
	}

	// === EARLY RETURNS FOR VALIDATION ===
	fmt.Println("\n--- EARLY RETURNS FOR VALIDATION ---")
	err = validateUser("Alice", 25, "alice@example.com")
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Println("User is valid")
	}

	err = validateUser("", 25, "alice@example.com")
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
	}

	// === RETURNING POINTERS VS VALUES ===
	fmt.Println("\n--- RETURNING POINTERS VS VALUES ---")

	// Large struct - return pointer
	largeResult := createLargeStruct()
	fmt.Printf("Large struct count: %d\n", largeResult.Count)

	// Small struct - return value
	smallResult := createSmallStruct()
	fmt.Printf("Small struct: %v\n", smallResult)

	// === CHAINING FUNCTION RETURNS ===
	fmt.Println("\n--- CHAINING FUNCTION RETURNS ---")
	chainResult := calculate(3, 4)
	fmt.Printf("Calculate(3, 4): %d\n", chainResult)

	// Equivalent to: multiply(add(3, 4), 2) = multiply(7, 2) = 14
	fmt.Printf("Step by step: add(3, 4) = %d, multiply(7, 2) = %d\n",
		add(3, 4), multiply(7, 2))

	// === RETURN VALUE PATTERNS ===
	fmt.Println("\n--- RETURN VALUE PATTERNS ---")

	// 1. Error handling pattern
	fmt.Println("1. Error handling pattern:")
	if result, err := divide(10, 2); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}

	// 2. Multiple assignment
	fmt.Println("2. Multiple assignment:")
	a, b, c := 1, 2, 3
	fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)

	// 3. Ignoring specific returns
	fmt.Println("3. Ignoring specific returns:")
	sum, _, _ = calculateStats([]int{1, 2, 3, 4, 5})
	fmt.Printf("Only sum: %d\n", sum)

	// 4. Function composition
	fmt.Println("4. Function composition:")
	composed := multiply(add(2, 3), square(2))
	fmt.Printf("multiply(add(2, 3), square(2)) = %d\n", composed)

	// === BEST PRACTICES ===
	fmt.Println("\n--- RETURN VALUE BEST PRACTICES ---")
	fmt.Println("1. Use multiple return values for error handling")
	fmt.Println("2. Use named return values for complex functions")
	fmt.Println("3. Return pointers for large structs")
	fmt.Println("4. Use early returns for validation")
	fmt.Println("5. Return zero values for error cases")
	fmt.Println("6. Use blank identifier for unused returns")
	fmt.Println("7. Be consistent with error return patterns")
	fmt.Println("8. Use interfaces for flexible return types")
	fmt.Println("9. Document return value meanings")
	fmt.Println("10. Avoid returning too many values")
}
