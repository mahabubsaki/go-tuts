package main

import (
	"fmt"
)

// === GO SEMICOLON HANDLING COMPREHENSIVE GUIDE ===

/*
SEMICOLON PHILOSOPHY:
- Go automatically inserts semicolons during lexical analysis
- You rarely need to write semicolons explicitly
- Understanding the rules helps avoid syntax errors
- Consistent with Go's goal of reducing syntactic noise

COMPARISON WITH JAVASCRIPT:
// JavaScript: Semicolons are optional but recommended
var x = 1;
var y = 2;
function hello() {
    console.log("Hello");
}

// Go: Semicolons are inserted automatically
var x = 1
var y = 2
func hello() {
    fmt.Println("Hello")
}
*/

// 1. AUTOMATIC SEMICOLON INSERTION RULES
// Go inserts semicolons automatically at the end of a line if:
// - The line ends with a token that can end a statement
// - The next line starts with a token that can begin a statement

func demonstrateAutomaticInsertion() {
	// These lines automatically get semicolons
	x := 1
	y := 2
	z := 3

	fmt.Printf("x = %d\n", x)
	fmt.Printf("y = %d\n", y)
	fmt.Printf("z = %d\n", z)
}

// 2. TOKENS THAT CAN END A STATEMENT
// - Identifiers
// - Basic literals (numbers, strings, etc.)
// - Keywords: break, continue, fallthrough, return
// - Operators: ++, --
// - Delimiters: ), ], }

func demonstrateEndTokens() {
	// All these get automatic semicolons
	name := "Go"       // identifier
	count := 42        // number literal
	message := "Hello" // string literal

	numbers := []int{1, 2, 3} // ] ends statement

	for i := 0; i < 3; i++ { // ++ ends statement
		fmt.Println(i)
	}

	fmt.Printf("Name: %s, Count: %d, Message: %s\n", name, count, message)
	fmt.Printf("Numbers: %v\n", numbers)
}

// 3. WHEN SEMICOLONS ARE REQUIRED
// Multiple statements on the same line
func demonstrateExplicitSemicolons() {
	// Explicit semicolons for multiple statements on one line
	x := 1
	y := 2
	z := 3

	fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)

	// For loop with multiple statements
	for i := 0; i < 3; i++ {
		fmt.Print(i)
		fmt.Print(" ")
	}
	fmt.Println()
}

// 4. COMMON SEMICOLON PITFALLS
func demonstratePitfalls() {
	// ❌ This would cause an error (if uncommented):
	// x := 1
	// + 2  // Go would insert semicolon after 1, making this invalid

	// ✅ Correct way:
	x := 1 + 2
	fmt.Printf("x = %d\n", x)

	// ❌ This would cause an error:
	// result := calculateSum(
	//     1, 2, 3
	// )  // Semicolon inserted after calculateSum(

	// ✅ Correct way:
	result := calculateSum(
		1, 2, 3) // Opening parenthesis on same line
	fmt.Printf("Result: %d\n", result)
}

func calculateSum(a, b, c int) int {
	return a + b + c
}

// 5. FUNCTION DECLARATIONS AND BRACES
func demonstrateBraceRules() {
	// ✅ Correct: Opening brace on same line
	if true {
		fmt.Println("This works")
	}

	// ❌ This would cause an error:
	// if true
	// {
	//     fmt.Println("This doesn't work")
	// }

	// ✅ Correct: Function declaration
	processData := func(data string) {
		fmt.Printf("Processing: %s\n", data)
	}

	processData("test data")
}

// 6. STRUCT LITERALS AND SLICES
func demonstrateStructLiterals() {
	// ✅ Correct: Closing brace position
	person := struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  30,
	} // Semicolon automatically inserted here

	// ✅ Correct: Slice literals
	numbers := []int{
		1,
		2,
		3,
	} // Semicolon automatically inserted here

	fmt.Printf("Person: %+v\n", person)
	fmt.Printf("Numbers: %v\n", numbers)
}

// 7. RETURN STATEMENTS
func demonstrateReturnStatements() (string, int) {
	// ✅ Correct: Same line
	return "hello", 42

	// ❌ This would cause an error:
	// return
	//     "hello", 42  // Semicolon inserted after return
}

// 8. SWITCH STATEMENTS
func demonstrateSwitchStatements() {
	value := 1

	// ✅ Correct: Opening brace on same line
	switch value {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Other")
	}

	// ✅ Correct: Switch with initialization
	switch x := value * 2; x {
	case 2:
		fmt.Println("Two")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Printf("Other: %d\n", x)
	}
}

// 9. INTERFACE DECLARATIONS
type Writer interface {
	Write([]byte) (int, error)
} // Semicolon automatically inserted

type Reader interface {
	Read([]byte) (int, error)
} // Semicolon automatically inserted

// 10. VARIABLE DECLARATIONS
var (
	globalVar1 = "value1"
	globalVar2 = "value2"
) // Semicolon automatically inserted

// 11. CONSTANT DECLARATIONS
const (
	MaxSize = 1024
	MinSize = 64
) // Semicolon automatically inserted

// 12. FOR LOOPS AND SEMICOLONS
func demonstrateForLoops() {
	// Traditional for loop with explicit semicolons
	for i := 0; i < 3; i++ {
		fmt.Printf("Traditional loop: %d\n", i)
	}

	// While-style loop (no semicolons needed)
	count := 0
	for count < 3 {
		fmt.Printf("While-style loop: %d\n", count)
		count++
	}

	// Infinite loop (no semicolons needed)
	loopCount := 0
	for {
		fmt.Printf("Infinite loop: %d\n", loopCount)
		loopCount++
		if loopCount >= 3 {
			break
		}
	}
}

// 13. DEFER STATEMENTS
func demonstrateDefer() {
	fmt.Println("Starting function")

	// Defer statements automatically get semicolons
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	defer fmt.Println("Deferred 3")

	fmt.Println("Ending function")
}

// 14. GO STATEMENTS (GOROUTINES)
func demonstrateGoroutines() {
	// Go statements automatically get semicolons
	done := make(chan bool)

	go func() {
		fmt.Println("Goroutine 1 running")
		done <- true
	}()

	go func() {
		fmt.Println("Goroutine 2 running")
		done <- true
	}()

	// Wait for both goroutines
	<-done
	<-done
}

// 15. BEST PRACTICES
func demonstrateBestPractices() {
	// 1. Don't fight the automatic insertion
	// 2. Keep opening braces on the same line
	// 3. Use explicit semicolons only when necessary
	// 4. Be aware of line break implications

	// ✅ Good: Natural Go style
	if true {
		fmt.Println("This follows Go conventions")
	}

	// ✅ Good: Multi-line expressions
	result := 1 + 2 + 3 + 4 + 5 +
		6 + 7 + 8 + 9 + 10 // Continuation on next line

	fmt.Printf("Result: %d\n", result)

	// ✅ Good: Function calls with many parameters
	complexFunction(
		"parameter1",
		"parameter2",
		"parameter3",
		"parameter4")
}

func complexFunction(a, b, c, d string) {
	fmt.Printf("Parameters: %s, %s, %s, %s\n", a, b, c, d)
}

func main() {
	fmt.Println("=== GO SEMICOLON HANDLING COMPREHENSIVE GUIDE ===")

	// === AUTOMATIC INSERTION ===
	fmt.Println("\n--- AUTOMATIC INSERTION ---")
	demonstrateAutomaticInsertion()

	// === END TOKENS ===
	fmt.Println("\n--- END TOKENS ---")
	demonstrateEndTokens()

	// === EXPLICIT SEMICOLONS ===
	fmt.Println("\n--- EXPLICIT SEMICOLONS ---")
	demonstrateExplicitSemicolons()

	// === PITFALLS ===
	fmt.Println("\n--- PITFALLS ---")
	demonstratePitfalls()

	// === BRACE RULES ===
	fmt.Println("\n--- BRACE RULES ---")
	demonstrateBraceRules()

	// === STRUCT LITERALS ===
	fmt.Println("\n--- STRUCT LITERALS ---")
	demonstrateStructLiterals()

	// === RETURN STATEMENTS ===
	fmt.Println("\n--- RETURN STATEMENTS ---")
	message, number := demonstrateReturnStatements()
	fmt.Printf("Returned: %s, %d\n", message, number)

	// === SWITCH STATEMENTS ===
	fmt.Println("\n--- SWITCH STATEMENTS ---")
	demonstrateSwitchStatements()

	// === FOR LOOPS ===
	fmt.Println("\n--- FOR LOOPS ---")
	demonstrateForLoops()

	// === DEFER STATEMENTS ===
	fmt.Println("\n--- DEFER STATEMENTS ---")
	demonstrateDefer()

	// === GOROUTINES ===
	fmt.Println("\n--- GOROUTINES ---")
	demonstrateGoroutines()

	// === BEST PRACTICES ===
	fmt.Println("\n--- BEST PRACTICES ---")
	demonstrateBestPractices()

	// === SEMICOLON RULES SUMMARY ===
	fmt.Println("\n--- SEMICOLON RULES SUMMARY ---")
	fmt.Println("1. Go automatically inserts semicolons")
	fmt.Println("2. Insertion happens at line ends with qualifying tokens")
	fmt.Println("3. Opening braces must be on the same line")
	fmt.Println("4. Use explicit semicolons for multiple statements per line")
	fmt.Println("5. Be careful with line breaks in expressions")
	fmt.Println("6. Function parameters can span multiple lines")
	fmt.Println("7. Return statements should be on the same line")
	fmt.Println("8. Don't fight the automatic insertion rules")

	// === COMMON PATTERNS ===
	fmt.Println("\n--- COMMON PATTERNS ---")
	fmt.Println("✅ if condition {")
	fmt.Println("❌ if condition")
	fmt.Println("   {")
	fmt.Println()
	fmt.Println("✅ return value")
	fmt.Println("❌ return")
	fmt.Println("   value")
	fmt.Println()
	fmt.Println("✅ x := 1 + 2 +")
	fmt.Println("   3 + 4")
	fmt.Println("❌ x := 1")
	fmt.Println("   + 2")

	fmt.Println("\nUnderstanding semicolon rules helps write idiomatic Go code!")
}
