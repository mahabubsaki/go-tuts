package main

import (
	"fmt"
	"os"
)

// === GO PRINTING COMPREHENSIVE GUIDE ===

/*
PRINTING PHILOSOPHY:
- Go provides multiple ways to print and format output
- fmt package is the standard for formatted I/O
- Printf family functions offer flexible formatting
- Print/Println for simple output, Printf for formatted output

COMPARISON WITH JAVASCRIPT:
// JavaScript
console.log("Hello", "World");           // Hello World
console.log("Age:", 25);                 // Age: 25
console.log(`Name: ${name}, Age: ${age}`); // Template literals

// Go
fmt.Println("Hello", "World")            // Hello World
fmt.Println("Age:", 25)                  // Age: 25
fmt.Printf("Name: %s, Age: %d\n", name, age) // Printf formatting
*/

// === DATA TYPES FOR PRINTING ===

type Person struct {
	Name  string
	Age   int
	Email string
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

func main() {
	fmt.Println("=== GO PRINTING COMPREHENSIVE GUIDE ===")

	// === BASIC PRINTING ===
	fmt.Println("\n1. BASIC PRINTING:")

	// Print - no spaces, no newline
	fmt.Print("Hello")
	fmt.Print("World")
	fmt.Print("!\n")

	// Println - spaces between args, newline at end
	fmt.Println("Hello", "World", "!")

	// Printf - formatted printing
	fmt.Printf("Hello %s!\n", "World")

	// === STRING FORMATTING ===
	fmt.Println("\n2. STRING FORMATTING:")

	name := "Alice"
	age := 30
	height := 5.6

	// Basic string formatting
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %.1f\n", height)

	// Alternative formatting
	fmt.Printf("Name: %v\n", name) // %v for any value
	fmt.Printf("Age: %v\n", age)
	fmt.Printf("Height: %v\n", height)

	// === NUMERIC FORMATTING ===
	fmt.Println("\n3. NUMERIC FORMATTING:")

	number := 42
	float := 3.14159

	// Integer formatting
	fmt.Printf("Decimal: %d\n", number)
	fmt.Printf("Binary: %b\n", number)
	fmt.Printf("Octal: %o\n", number)
	fmt.Printf("Hexadecimal: %x\n", number)
	fmt.Printf("Hexadecimal (uppercase): %X\n", number)

	// Float formatting
	fmt.Printf("Float: %f\n", float)
	fmt.Printf("Float (2 decimals): %.2f\n", float)
	fmt.Printf("Scientific notation: %e\n", float)
	fmt.Printf("Scientific notation (uppercase): %E\n", float)
	fmt.Printf("Compact format: %g\n", float)

	// === BOOLEAN AND CHARACTER FORMATTING ===
	fmt.Println("\n4. BOOLEAN AND CHARACTER FORMATTING:")

	isTrue := true
	char := 'A'

	fmt.Printf("Boolean: %t\n", isTrue)
	fmt.Printf("Character: %c\n", char)
	fmt.Printf("Character as number: %d\n", char)
	fmt.Printf("Unicode: %U\n", char)

	// === STRUCT PRINTING ===
	fmt.Println("\n5. STRUCT PRINTING:")

	person := Person{
		Name:  "John Doe",
		Age:   25,
		Email: "john@example.com",
	}

	// Different ways to print structs
	fmt.Printf("Person: %v\n", person)      // {John Doe 25 john@example.com}
	fmt.Printf("Person: %+v\n", person)     // {Name:John Doe Age:25 Email:john@example.com}
	fmt.Printf("Person: %#v\n", person)     // main.Person{Name:"John Doe", Age:25, Email:"john@example.com"}
	fmt.Printf("Person type: %T\n", person) // main.Person

	// === SLICE AND MAP PRINTING ===
	fmt.Println("\n6. SLICE AND MAP PRINTING:")

	numbers := []int{1, 2, 3, 4, 5}
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Numbers: %#v\n", numbers)
	fmt.Printf("Colors: %v\n", colors)
	fmt.Printf("Colors: %#v\n", colors)

	// === POINTER PRINTING ===
	fmt.Println("\n7. POINTER PRINTING:")

	x := 42
	ptr := &x

	fmt.Printf("Value: %v\n", x)
	fmt.Printf("Pointer: %p\n", ptr)
	fmt.Printf("Pointer value: %v\n", *ptr)

	// === WIDTH AND PRECISION ===
	fmt.Println("\n8. WIDTH AND PRECISION:")

	// Width specification
	fmt.Printf("Right-aligned: '%5d'\n", 42) // '   42'
	fmt.Printf("Left-aligned: '%-5d'\n", 42) // '42   '
	fmt.Printf("Zero-padded: '%05d'\n", 42)  // '00042'

	// String width
	fmt.Printf("Right-aligned: '%10s'\n", "hello") // '     hello'
	fmt.Printf("Left-aligned: '%-10s'\n", "hello") // 'hello     '
	fmt.Printf("Truncated: '%.3s'\n", "hello")     // 'hel'

	// Float precision
	pi := 3.14159265359
	fmt.Printf("Default: %f\n", pi)                  // 3.141593
	fmt.Printf("2 decimals: %.2f\n", pi)             // 3.14
	fmt.Printf("Width and precision: '%8.2f'\n", pi) // '    3.14'

	// === SPRINTF - STRING FORMATTING ===
	fmt.Println("\n9. SPRINTF - STRING FORMATTING:")

	formatted := fmt.Sprintf("Hello %s, you are %d years old", name, age)
	fmt.Println("Formatted string:", formatted)

	// Building complex strings
	product := Product{ID: 1, Name: "Laptop", Price: 999.99}
	description := fmt.Sprintf("Product #%d: %s - $%.2f", product.ID, product.Name, product.Price)
	fmt.Println("Product description:", description)

	// === FPRINT - WRITING TO WRITERS ===
	fmt.Println("\n10. FPRINT - WRITING TO WRITERS:")

	// Writing to stdout (same as Print)
	fmt.Fprint(os.Stdout, "Hello from Fprint\n")

	// Writing to stderr
	fmt.Fprint(os.Stderr, "Error message\n")

	// === CUSTOM PRINT FUNCTIONS ===
	fmt.Println("\n11. CUSTOM PRINT FUNCTIONS:")

	// Custom print function
	printPerson := func(p Person) {
		fmt.Printf("👤 %s (%d) - %s\n", p.Name, p.Age, p.Email)
	}

	printPerson(person)

	// Custom format function
	formatCurrency := func(amount float64) string {
		return fmt.Sprintf("$%.2f", amount)
	}

	fmt.Printf("Price: %s\n", formatCurrency(19.99))

	// === DEBUGGING PRINT ===
	fmt.Println("\n12. DEBUGGING PRINT:")

	// Quick debug print
	debugVar := map[string]interface{}{
		"user":    "admin",
		"active":  true,
		"balance": 1234.56,
	}

	fmt.Printf("DEBUG: %#v\n", debugVar)

	// === CONDITIONAL PRINTING ===
	fmt.Println("\n13. CONDITIONAL PRINTING:")

	verbose := true

	if verbose {
		fmt.Println("Verbose mode: Operation completed successfully")
	}

	// Conditional format
	status := "active"
	symbol := "✓"
	if status != "active" {
		symbol = "✗"
	}
	fmt.Printf("Status: %s %s\n", symbol, status)

	// === ESCAPE SEQUENCES ===
	fmt.Println("\n14. ESCAPE SEQUENCES:")

	fmt.Printf("Newline: Line1\\nLine2\n")
	fmt.Printf("Tab: Column1\\tColumn2\n")
	fmt.Printf("Carriage return: Hello\\rWorld\n")
	fmt.Printf("Backslash: C:\\\\Users\\\\Name\n")
	fmt.Printf("Quotes: \"Hello 'World'\"\n")

	// === PERFORMANCE CONSIDERATIONS ===
	fmt.Println("\n15. PERFORMANCE CONSIDERATIONS:")

	// For simple output, Println is faster
	fmt.Println("Simple output")

	// For formatted output, Printf is necessary
	fmt.Printf("Formatted output: %s\n", "value")

	// For building strings, Sprintf is useful
	message := fmt.Sprintf("User %s logged in at %v", "admin", "2024-01-01")
	fmt.Println(message)

	// === REAL-WORLD EXAMPLES ===
	fmt.Println("\n16. REAL-WORLD EXAMPLES:")

	// Log-like output
	fmt.Printf("[%s] %s: %s\n", "INFO", "2024-01-01 10:00:00", "Application started")

	// Table-like output
	fmt.Printf("%-10s | %-5s | %s\n", "Name", "Age", "Email")
	fmt.Printf("%-10s | %-5s | %s\n", "----------", "-----", "-----")
	fmt.Printf("%-10s | %-5d | %s\n", "Alice", 30, "alice@example.com")
	fmt.Printf("%-10s | %-5d | %s\n", "Bob", 25, "bob@example.com")

	// Progress indicator
	for i := 0; i <= 100; i += 20 {
		fmt.Printf("\rProgress: %d%% [", i)
		for j := 0; j < i/2; j++ {
			fmt.Print("█")
		}
		for j := i / 2; j < 50; j++ {
			fmt.Print("░")
		}
		fmt.Print("]")
		if i == 100 {
			fmt.Println(" Complete!")
		}
	}

	fmt.Println("\n=== END OF PRINTING GUIDE ===")
}

/*
RUNNING THE PROGRAM:
go run main.go

LEARNING POINTS:

1. PRINT FUNCTIONS:
   - Print: No spaces, no newline
   - Println: Spaces between args, newline at end
   - Printf: Formatted printing with verbs

2. FORMAT VERBS:
   - %v: Default format for any value
   - %+v: Struct with field names
   - %#v: Go-syntax representation
   - %T: Type of value
   - %s: String
   - %d: Integer
   - %f: Float
   - %t: Boolean
   - %c: Character
   - %p: Pointer

3. WIDTH AND PRECISION:
   - %5d: Right-aligned, minimum 5 characters
   - %-5d: Left-aligned, minimum 5 characters
   - %05d: Zero-padded, minimum 5 characters
   - %.2f: 2 decimal places
   - %8.2f: 8 characters wide, 2 decimal places

4. STRING BUILDING:
   - Sprintf: Format and return string
   - Fprint: Write to any io.Writer

5. DEBUGGING:
   - %#v: Shows Go syntax representation
   - %+v: Shows struct field names
   - %T: Shows type information

This demonstrates comprehensive printing and formatting in Go
for debugging, logging, and user output.
*/
