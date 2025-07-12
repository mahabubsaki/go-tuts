package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// === GO FORMATTING COMPREHENSIVE GUIDE ===

/*
FORMATTING PHILOSOPHY:
- Go has a single, consistent formatting style
- gofmt eliminates debates about formatting
- Consistency improves readability and maintainability
- Automatic formatting reduces cognitive load

COMPARISON WITH JAVASCRIPT:
// JavaScript: Multiple valid formatting styles
var x=1,y=2,z=3;
var x = 1, y = 2, z = 3;
var x = 1,
    y = 2,
    z = 3;

// Go: Single consistent style enforced by gofmt
var (
    x = 1
    y = 2
    z = 3
)
*/

// 1. Package and import formatting
// Imports are automatically organized by goimports
// Standard library imports first, then third-party, then local

// 2. Function formatting
func demonstrateFormatting() {
	// Variable declarations
	var name string = "Go"
	var version float64 = 1.21
	var isStable bool = true

	// Short variable declarations
	language := "Go"
	year := 2009

	// Multiple variable declarations
	var (
		author     = "Google"
		opensource = true
		compiled   = true
	)

	// Print formatted output
	fmt.Printf("Language: %s\n", language)
	fmt.Printf("Version: %.2f\n", version)
	fmt.Printf("Year: %d\n", year)
	fmt.Printf("Author: %s\n", author)
	fmt.Printf("Stable: %t\n", isStable)
	fmt.Printf("Open Source: %t\n", opensource)
	fmt.Printf("Compiled: %t\n", compiled)

	// Demonstrate variable naming conventions
	fmt.Printf("Name: %s\n", name)
}

// 3. Struct formatting
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
	Address   Address
}

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

// 4. Interface formatting
type Writer interface {
	Write([]byte) (int, error)
}

type ReadWriter interface {
	Writer
	Read([]byte) (int, error)
}

// 5. Method formatting
func (p *Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) UpdateEmail(email string) {
	p.Email = email
}

// 6. Function with multiple return values
func parsePersonData(data string) (Person, error) {
	parts := strings.Split(data, ",")
	if len(parts) < 4 {
		return Person{}, fmt.Errorf("invalid data format")
	}

	return Person{
		FirstName: strings.TrimSpace(parts[0]),
		LastName:  strings.TrimSpace(parts[1]),
		Email:     strings.TrimSpace(parts[2]),
		Address: Address{
			Street: strings.TrimSpace(parts[3]),
		},
	}, nil
}

// 7. Control structure formatting
func demonstrateControlStructures() {
	// If statement formatting
	age := 25
	if age >= 18 {
		fmt.Println("Adult")
	} else if age >= 13 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Child")
	}

	// For loop formatting
	for i := 0; i < 5; i++ {
		fmt.Printf("Count: %d\n", i)
	}

	// Range loop formatting
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Switch statement formatting
	switch age {
	case 18:
		fmt.Println("Just became an adult")
	case 21:
		fmt.Println("Can drink in the US")
	default:
		fmt.Println("Another age")
	}
}

// 8. Array and slice formatting
func demonstrateCollections() {
	// Array formatting
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	// Slice formatting
	fruits := []string{
		"apple",
		"banana",
		"orange",
		"grape",
	}

	// Map formatting
	ages := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}

	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Fruits: %v\n", fruits)
	fmt.Printf("Ages: %v\n", ages)
}

// 9. Comment formatting
func demonstrateComments() {
	// Single-line comments should have a space after //
	fmt.Println("Single-line comment example")

	/*
		Multi-line comments are formatted consistently
		Each line should align properly
		Use for longer explanations
	*/

	// TODO: This is a properly formatted TODO comment
	// FIXME: This is a properly formatted FIXME comment
	// NOTE: This is a properly formatted NOTE comment
}

// 10. Error handling formatting
func demonstrateErrorHandling() {
	file, err := os.Open("nonexistent.txt")
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()

	// Multiple error checks
	data := "John,Doe,john@example.com,123 Main St"
	person, err := parsePersonData(data)
	if err != nil {
		log.Printf("Error parsing person data: %v", err)
		return
	}

	fmt.Printf("Parsed person: %+v\n", person)
}

// 11. Constant formatting
const (
	MaxRetries = 3
	Timeout    = 30
	BufferSize = 1024
)

// 12. Enum-like constant formatting
type Status int

const (
	StatusPending Status = iota
	StatusApproved
	StatusRejected
	StatusCancelled
)

// 13. Long function parameter formatting
func processUserData(
	firstName string,
	lastName string,
	email string,
	age int,
	address Address,
	preferences map[string]interface{},
) (*Person, error) {
	// Function implementation
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Age:       age,
		Address:   address,
	}, nil
}

// 14. Channel formatting
func demonstrateChannels() {
	// Channel declarations
	ch := make(chan int)
	done := make(chan bool)

	// Goroutine with proper formatting
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
		done <- true
	}()

	// Channel reading
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}

	<-done
}

// 15. Function literal formatting
func demonstrateFunctionLiterals() {
	// Anonymous function formatting
	calculate := func(a, b int, operation string) int {
		switch operation {
		case "add":
			return a + b
		case "subtract":
			return a - b
		case "multiply":
			return a * b
		case "divide":
			if b != 0 {
				return a / b
			}
			return 0
		default:
			return 0
		}
	}

	fmt.Printf("Addition: %d\n", calculate(10, 5, "add"))
	fmt.Printf("Subtraction: %d\n", calculate(10, 5, "subtract"))
}

func main() {
	fmt.Println("=== GO FORMATTING COMPREHENSIVE GUIDE ===")

	// === BASIC FORMATTING ===
	fmt.Println("\n--- BASIC FORMATTING ---")
	demonstrateFormatting()

	// === CONTROL STRUCTURES ===
	fmt.Println("\n--- CONTROL STRUCTURES ---")
	demonstrateControlStructures()

	// === COLLECTIONS ===
	fmt.Println("\n--- COLLECTIONS ---")
	demonstrateCollections()

	// === COMMENTS ===
	fmt.Println("\n--- COMMENTS ---")
	demonstrateComments()

	// === ERROR HANDLING ===
	fmt.Println("\n--- ERROR HANDLING ---")
	demonstrateErrorHandling()

	// === CHANNELS ===
	fmt.Println("\n--- CHANNELS ---")
	demonstrateChannels()

	// === FUNCTION LITERALS ===
	fmt.Println("\n--- FUNCTION LITERALS ---")
	demonstrateFunctionLiterals()

	// === FORMATTING TOOLS ===
	fmt.Println("\n--- FORMATTING TOOLS ---")
	fmt.Println("1. gofmt: Formats Go source code")
	fmt.Println("2. goimports: Manages imports and formats code")
	fmt.Println("3. gofumpt: Stricter formatting rules")
	fmt.Println("4. golangci-lint: Comprehensive linting")

	// === BEST PRACTICES ===
	fmt.Println("\n--- FORMATTING BEST PRACTICES ---")
	fmt.Println("1. Always use gofmt/goimports")
	fmt.Println("2. Configure your editor to format on save")
	fmt.Println("3. Use consistent naming conventions")
	fmt.Println("4. Add spaces around operators")
	fmt.Println("5. Use meaningful variable names")
	fmt.Println("6. Group related declarations")
	fmt.Println("7. Keep functions reasonably sized")
	fmt.Println("8. Use proper indentation")
	fmt.Println("9. Add comments for complex logic")
	fmt.Println("10. Follow Go idioms and conventions")

	fmt.Println("\nConsistent formatting makes Go code readable and maintainable!")
}
