package main

import "fmt"

// === FUNCTION PARAMETERS ===

// 1. Single parameter
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 2. Multiple parameters with different types
func personalInfo(name string, age int, height float64, isStudent bool) {
	fmt.Printf("Name: %s, Age: %d, Height: %.1f, Student: %t\n", name, age, height, isStudent)
}

// 3. Multiple parameters with same type (shorthand)
func addThree(a, b, c int) int {
	return a + b + c
}

// 4. Multiple parameters with mixed types
func processData(id int, name string, scores []int, active bool) {
	fmt.Printf("ID: %d, Name: %s, Active: %t\n", id, name, active)
	fmt.Printf("Scores: %v\n", scores)
}

// 5. Function with slice parameter
func sumSlice(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 6. Function with map parameter
func printGrades(grades map[string]int) {
	for student, grade := range grades {
		fmt.Printf("%s: %d\n", student, grade)
	}
}

// 7. Function with array parameter (copies the array)
func sumArray(numbers [5]int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 8. Function with pointer parameter
func updateValue(ptr *int, newValue int) {
	*ptr = newValue
}

// 9. Function with struct parameter (pass by value)
type Person struct {
	Name string
	Age  int
}

func describePerson(p Person) {
	fmt.Printf("Person: %s, Age: %d\n", p.Name, p.Age)
}

// 10. Function with struct pointer parameter (pass by reference)
func updatePerson(p *Person, newName string, newAge int) {
	p.Name = newName
	p.Age = newAge
}

// 11. Function with interface parameter
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
	return 3.14159 * c.Radius * c.Radius
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

// 12. Function with function parameter
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// 13. Function with channel parameter
func sendData(ch chan<- int, data []int) {
	for _, value := range data {
		ch <- value
	}
	close(ch)
}

func receiveData(ch <-chan int) []int {
	var result []int
	for value := range ch {
		result = append(result, value)
	}
	return result
}

// 14. Function with variadic parameters
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 15. Function with regular and variadic parameters
func formatMessage(prefix string, messages ...string) string {
	result := prefix + ": "
	for i, msg := range messages {
		if i > 0 {
			result += ", "
		}
		result += msg
	}
	return result
}

// 16. Function with empty interface parameter (accepts any type)
func printAnything(value interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

// 17. Function with multiple interface parameters
func compare(a, b interface{}) {
	fmt.Printf("Comparing %v (%T) with %v (%T)\n", a, a, b, b)
}

// 18. Function with optional-like parameters using structs
type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	SSL      bool
}

func connect(config Config) {
	fmt.Printf("Connecting to %s:%d as %s (SSL: %t)\n", config.Host, config.Port, config.Username, config.SSL)
}

// 19. Function with default values simulation
func connectWithDefaults(host string, port int) {
	if host == "" {
		host = "localhost"
	}
	if port == 0 {
		port = 8080
	}
	fmt.Printf("Connecting to %s:%d\n", host, port)
}

// 20. Function with parameter validation
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("=== GO FUNCTIONS - PARAMETERS ===")

	// === SINGLE PARAMETER ===
	fmt.Println("\n--- SINGLE PARAMETER ---")
	greet("Alice")
	greet("Bob")

	/*
		JavaScript comparison:
		function greet(name) {
			console.log(`Hello, ${name}!`);
		}
	*/

	// === MULTIPLE PARAMETERS ===
	fmt.Println("\n--- MULTIPLE PARAMETERS ---")
	personalInfo("Charlie", 25, 175.5, true)
	personalInfo("Diana", 30, 168.0, false)

	// Same type parameters
	result := addThree(1, 2, 3)
	fmt.Printf("1 + 2 + 3 = %d\n", result)

	// Mixed type parameters
	scores := []int{85, 90, 78, 92}
	processData(1001, "Eve", scores, true)

	/*
		JavaScript comparison:
		function personalInfo(name, age, height, isStudent) {
			console.log(`Name: ${name}, Age: ${age}, Height: ${height}, Student: ${isStudent}`);
		}
	*/

	// === SLICE PARAMETERS ===
	fmt.Println("\n--- SLICE PARAMETERS ---")
	numbers := []int{1, 2, 3, 4, 5}
	total := sumSlice(numbers)
	fmt.Printf("Sum of %v = %d\n", numbers, total)

	// Modifying slice doesn't affect original (slice header is copied)
	fmt.Println("Original slice:", numbers)

	// === MAP PARAMETERS ===
	fmt.Println("\n--- MAP PARAMETERS ---")
	grades := map[string]int{
		"Alice": 95,
		"Bob":   87,
		"Carol": 92,
	}
	fmt.Println("Student grades:")
	printGrades(grades)

	// === ARRAY PARAMETERS ===
	fmt.Println("\n--- ARRAY PARAMETERS ---")
	arr := [5]int{1, 2, 3, 4, 5}
	arraySum := sumArray(arr)
	fmt.Printf("Sum of array %v = %d\n", arr, arraySum)

	// === POINTER PARAMETERS ===
	fmt.Println("\n--- POINTER PARAMETERS ---")
	value := 10
	fmt.Printf("Before update: %d\n", value)
	updateValue(&value, 20)
	fmt.Printf("After update: %d\n", value)

	/*
		JavaScript comparison (objects are passed by reference):
		function updateValue(obj, newValue) {
			obj.value = newValue;
		}

		const obj = { value: 10 };
		updateValue(obj, 20);
		console.log(obj.value); // 20
	*/

	// === STRUCT PARAMETERS ===
	fmt.Println("\n--- STRUCT PARAMETERS ---")
	person := Person{Name: "Frank", Age: 28}
	fmt.Println("Original person:", person)

	// Pass by value - original struct is not modified
	describePerson(person)

	// Pass by reference - original struct is modified
	fmt.Println("Before update:", person)
	updatePerson(&person, "Franklin", 29)
	fmt.Println("After update:", person)

	// === INTERFACE PARAMETERS ===
	fmt.Println("\n--- INTERFACE PARAMETERS ---")
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	fmt.Printf("Rectangle area: %.2f\n", calculateArea(rect))
	fmt.Printf("Circle area: %.2f\n", calculateArea(circle))

	// === FUNCTION PARAMETERS ===
	fmt.Println("\n--- FUNCTION PARAMETERS ---")
	add := func(a, b int) int { return a + b }
	multiply := func(a, b int) int { return a * b }

	fmt.Printf("5 + 3 = %d\n", applyOperation(5, 3, add))
	fmt.Printf("5 × 3 = %d\n", applyOperation(5, 3, multiply))

	// Anonymous function as parameter
	fmt.Printf("5 - 3 = %d\n", applyOperation(5, 3, func(a, b int) int {
		return a - b
	}))

	/*
		JavaScript comparison:
		function applyOperation(a, b, operation) {
			return operation(a, b);
		}

		const add = (a, b) => a + b;
		const multiply = (a, b) => a * b;

		console.log(applyOperation(5, 3, add));
		console.log(applyOperation(5, 3, multiply));
	*/

	// === CHANNEL PARAMETERS ===
	fmt.Println("\n--- CHANNEL PARAMETERS ---")
	ch := make(chan int)
	data := []int{1, 2, 3, 4, 5}

	go sendData(ch, data)
	received := receiveData(ch)
	fmt.Printf("Sent: %v, Received: %v\n", data, received)

	// === VARIADIC PARAMETERS ===
	fmt.Println("\n--- VARIADIC PARAMETERS ---")
	fmt.Printf("Sum of 1, 2, 3: %d\n", sum(1, 2, 3))
	fmt.Printf("Sum of 1, 2, 3, 4, 5: %d\n", sum(1, 2, 3, 4, 5))

	// Unpacking slice
	nums := []int{10, 20, 30}
	fmt.Printf("Sum of slice %v: %d\n", nums, sum(nums...))

	// Regular and variadic parameters
	message := formatMessage("Error", "File not found", "Permission denied", "Invalid input")
	fmt.Println(message)

	/*
		JavaScript comparison:
		function sum(...numbers) {
			return numbers.reduce((total, num) => total + num, 0);
		}

		console.log(sum(1, 2, 3));
		console.log(sum(1, 2, 3, 4, 5));
	*/

	// === EMPTY INTERFACE PARAMETERS ===
	fmt.Println("\n--- EMPTY INTERFACE PARAMETERS ---")
	printAnything(42)
	printAnything("Hello")
	printAnything([]int{1, 2, 3})
	printAnything(person)
	printAnything(true)

	// Multiple interface parameters
	compare(42, "Hello")
	compare([]int{1, 2}, map[string]int{"a": 1})

	// === STRUCT-BASED CONFIGURATION ===
	fmt.Println("\n--- STRUCT-BASED CONFIGURATION ---")
	config := Config{
		Host:     "api.example.com",
		Port:     443,
		Username: "admin",
		Password: "secret",
		SSL:      true,
	}
	connect(config)

	// Partial configuration
	basicConfig := Config{
		Host: "localhost",
		Port: 8080,
	}
	connect(basicConfig)

	// === SIMULATED DEFAULT VALUES ===
	fmt.Println("\n--- SIMULATED DEFAULT VALUES ---")
	connectWithDefaults("", 0)            // Uses defaults
	connectWithDefaults("example.com", 0) // Uses default port
	connectWithDefaults("", 9000)         // Uses default host
	connectWithDefaults("api.com", 8080)  // Uses provided values

	// === PARAMETER VALIDATION ===
	fmt.Println("\n--- PARAMETER VALIDATION ---")
	result1, err1 := divide(10, 2)
	if err1 != nil {
		fmt.Printf("Error: %v\n", err1)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", result1)
	}

	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Printf("Error: %v\n", err2)
	} else {
		fmt.Printf("10 ÷ 0 = %.2f\n", result2)
	}

	// === PARAMETER PATTERNS ===
	fmt.Println("\n--- PARAMETER PATTERNS ---")

	// 1. Pass by value (copies the data)
	fmt.Println("Pass by value:")
	originalInt := 100
	fmt.Printf("Original: %d\n", originalInt)
	// Function receives a copy, original is unchanged

	// 2. Pass by reference (passes memory address)
	fmt.Println("Pass by reference:")
	fmt.Printf("Original: %d\n", value)
	updateValue(&value, 200)
	fmt.Printf("After update: %d\n", value)

	// 3. Slice/Map/Channel parameters (reference types)
	fmt.Println("Reference types:")
	slice := []int{1, 2, 3}
	fmt.Printf("Original slice: %v\n", slice)
	// Modifying slice contents affects original

	// === BEST PRACTICES ===
	fmt.Println("\n--- PARAMETER BEST PRACTICES ---")
	fmt.Println("1. Use descriptive parameter names")
	fmt.Println("2. Group parameters by type when possible")
	fmt.Println("3. Use pointers for large structs to avoid copying")
	fmt.Println("4. Use interfaces for flexible parameter types")
	fmt.Println("5. Validate parameters when necessary")
	fmt.Println("6. Use variadic parameters for variable argument lists")
	fmt.Println("7. Use struct configuration for many optional parameters")
	fmt.Println("8. Keep parameter lists short and focused")
	fmt.Println("9. Use channels for concurrent communication")
	fmt.Println("10. Document parameter requirements clearly")
}
