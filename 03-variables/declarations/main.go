package main

import "fmt"

// This file covers all variable declaration methods in Go
// Go has several ways to declare variables, each with different use cases

// Package-level variables (global scope)
var globalVar string = "I'm a global variable"
var globalInt int = 42

// Multiple variable declarations
var (
	name    string = "Go"
	version string = "1.21"
	stable  bool   = true
)

// Type inference at package level
var inferredGlobal = "This type is inferred"

func main() {
	fmt.Println("=== GO VARIABLE DECLARATIONS - COMPLETE GUIDE ===")

	demonstrateVarDeclaration()
	demonstrateShortDeclaration()
	demonstrateZeroValues()
	demonstrateMultipleDeclarations()
	demonstrateTypeInference()
	demonstrateScope()
	demonstrateGlobalVariables()
}

func demonstrateVarDeclaration() {
	fmt.Println("\n--- VAR DECLARATION ---")

	// Basic var declaration with type and value
	var message string = "Hello, Go!"
	var count int = 10
	var isActive bool = true
	var price float64 = 99.99

	fmt.Printf("message: %s\n", message)
	fmt.Printf("count: %d\n", count)
	fmt.Printf("isActive: %t\n", isActive)
	fmt.Printf("price: %.2f\n", price)

	// var declaration with type only (zero value)
	var emptyString string
	var emptyInt int
	var emptyBool bool
	var emptyFloat float64

	fmt.Printf("emptyString: '%s'\n", emptyString)
	fmt.Printf("emptyInt: %d\n", emptyInt)
	fmt.Printf("emptyBool: %t\n", emptyBool)
	fmt.Printf("emptyFloat: %f\n", emptyFloat)

	// var declaration with value only (type inferred)
	var inferredString = "Type inferred as string"
	var inferredInt = 42
	var inferredFloat = 3.14
	var inferredBool = true

	fmt.Printf("inferredString: %s (type: %T)\n", inferredString, inferredString)
	fmt.Printf("inferredInt: %d (type: %T)\n", inferredInt, inferredInt)
	fmt.Printf("inferredFloat: %f (type: %T)\n", inferredFloat, inferredFloat)
	fmt.Printf("inferredBool: %t (type: %T)\n", inferredBool, inferredBool)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: var x = 10; (function-scoped, can be redeclared)
	// JavaScript: let x = 10; (block-scoped, cannot be redeclared)
	// JavaScript: const x = 10; (block-scoped, cannot be reassigned)
	// Go: var x int = 10; (typed, explicitly scoped)
}

func demonstrateShortDeclaration() {
	fmt.Println("\n--- SHORT DECLARATION (:=) ---")

	// Short variable declaration (only inside functions)
	name := "Alice"
	age := 30
	salary := 75000.50
	married := true

	fmt.Printf("name: %s (type: %T)\n", name, name)
	fmt.Printf("age: %d (type: %T)\n", age, age)
	fmt.Printf("salary: %.2f (type: %T)\n", salary, salary)
	fmt.Printf("married: %t (type: %T)\n", married, married)

	// Short declaration with multiple variables
	x, y := 10, 20
	fmt.Printf("x: %d, y: %d\n", x, y)

	// Mixed assignment (at least one new variable)
	name, city := "Bob", "New York" // name is reassigned, city is new
	fmt.Printf("name: %s, city: %s\n", name, city)

	// Short declaration with function return
	result, ok := divideInts(10, 2)
	fmt.Printf("division result: %d, ok: %t\n", result, ok)

	// Short declaration limitations
	// 1. Only works inside functions
	// 2. Cannot specify type explicitly
	// 3. At least one variable must be new

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: let x = 10; (type inferred, can be reassigned)
	// JavaScript: const x = 10; (type inferred, cannot be reassigned)
	// Go: x := 10; (type inferred, can be reassigned, function-scoped only)
}

func demonstrateZeroValues() {
	fmt.Println("\n--- ZERO VALUES ---")

	// Go initializes all variables to their zero values
	// This is different from JavaScript's undefined

	var boolVal bool
	var intVal int
	var floatVal float64
	var stringVal string
	var pointerVal *int
	var sliceVal []int
	var mapVal map[string]int
	var funcVal func()
	var interfaceVal interface{}
	var channelVal chan int

	fmt.Printf("bool zero value: %t\n", boolVal)
	fmt.Printf("int zero value: %d\n", intVal)
	fmt.Printf("float64 zero value: %f\n", floatVal)
	fmt.Printf("string zero value: '%s'\n", stringVal)
	fmt.Printf("pointer zero value: %v\n", pointerVal)
	fmt.Printf("slice zero value: %v\n", sliceVal)
	fmt.Printf("map zero value: %v\n", mapVal)
	fmt.Printf("function zero value: %v\n", funcVal == nil)
	fmt.Printf("interface zero value: %v\n", interfaceVal)
	fmt.Printf("channel zero value: %v\n", channelVal)

	// Checking for zero values
	if stringVal == "" {
		fmt.Println("String is empty (zero value)")
	}

	if sliceVal == nil {
		fmt.Println("Slice is nil (zero value)")
	}

	if mapVal == nil {
		fmt.Println("Map is nil (zero value)")
	}

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: var x; (x is undefined)
	// JavaScript: let x; (x is undefined)
	// Go: var x int; (x is 0, the zero value for int)
	// Go has no undefined - every type has a meaningful zero value
}

func demonstrateMultipleDeclarations() {
	fmt.Println("\n--- MULTIPLE DECLARATIONS ---")

	// Multiple var declarations (traditional)
	var a, b, c int = 1, 2, 3
	fmt.Printf("a: %d, b: %d, c: %d\n", a, b, c)

	// Multiple var declarations with different types
	var (
		username string  = "alice"
		userID   int     = 123
		active   bool    = true
		balance  float64 = 1500.75
	)

	fmt.Printf("username: %s, userID: %d, active: %t, balance: %.2f\n",
		username, userID, active, balance)

	// Multiple short declarations
	name, age, city := "Bob", 25, "Boston"
	fmt.Printf("name: %s, age: %d, city: %s\n", name, age, city)

	// Multiple declarations with type inference
	var x, y = 10, 20.5 // x is int, y is float64
	fmt.Printf("x: %d (type: %T), y: %f (type: %T)\n", x, x, y, y)

	// Variable swapping
	x, y = int(y), float64(x) // Swapping with type conversion
	fmt.Printf("After swap - x: %d, y: %f\n", x, y)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: let x = 1, y = 2, z = 3;
	// JavaScript: let [x, y, z] = [1, 2, 3]; (destructuring)
	// Go: var x, y, z = 1, 2, 3
	// Go: x, y, z := 1, 2, 3
}

func demonstrateTypeInference() {
	fmt.Println("\n--- TYPE INFERENCE ---")

	// Type inference with literals
	var intLiteral = 42         // int
	var floatLiteral = 3.14     // float64
	var stringLiteral = "hello" // string
	var boolLiteral = true      // bool
	var runeLiteral = 'A'       // rune (int32)

	fmt.Printf("intLiteral: %d (type: %T)\n", intLiteral, intLiteral)
	fmt.Printf("floatLiteral: %f (type: %T)\n", floatLiteral, floatLiteral)
	fmt.Printf("stringLiteral: %s (type: %T)\n", stringLiteral, stringLiteral)
	fmt.Printf("boolLiteral: %t (type: %T)\n", boolLiteral, boolLiteral)
	fmt.Printf("runeLiteral: %c (type: %T)\n", runeLiteral, runeLiteral)

	// Type inference with expressions
	var sum = 10 + 5                // int
	var product = 3.14 * 2          // float64
	var concat = "Hello" + " World" // string
	var comparison = 10 > 5         // bool

	fmt.Printf("sum: %d (type: %T)\n", sum, sum)
	fmt.Printf("product: %f (type: %T)\n", product, product)
	fmt.Printf("concat: %s (type: %T)\n", concat, concat)
	fmt.Printf("comparison: %t (type: %T)\n", comparison, comparison)

	// Type inference with function calls
	var length = len("Hello") // int
	var sqrt = 9.0            // float64

	fmt.Printf("length: %d (type: %T)\n", length, length)
	fmt.Printf("sqrt: %f (type: %T)\n", sqrt, sqrt)

	// Explicit type when needed
	var explicitFloat32 float32 = 3.14
	var explicitInt64 int64 = 42

	fmt.Printf("explicitFloat32: %f (type: %T)\n", explicitFloat32, explicitFloat32)
	fmt.Printf("explicitInt64: %d (type: %T)\n", explicitInt64, explicitInt64)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Dynamic typing - type determined at runtime
	// JavaScript: let x = 42; x = "hello"; (type can change)
	// Go: Static typing - type determined at compile time
	// Go: var x = 42; x = "hello"; (compilation error)
}

func demonstrateScope() {
	fmt.Println("\n--- VARIABLE SCOPE ---")

	// Function scope
	var functionVar = "Function scope"
	fmt.Printf("functionVar: %s\n", functionVar)

	// Block scope
	{
		var blockVar = "Block scope"
		fmt.Printf("blockVar: %s\n", blockVar)

		// Can access function scope from block
		fmt.Printf("Accessing functionVar from block: %s\n", functionVar)
	}

	// blockVar is not accessible here
	// fmt.Printf("blockVar: %s\n", blockVar) // This would cause an error

	// if statement scope
	if condition := true; condition {
		var ifVar = "If statement scope"
		fmt.Printf("ifVar: %s\n", ifVar)
		fmt.Printf("condition: %t\n", condition)
	}

	// condition and ifVar are not accessible here

	// for loop scope
	for i := 0; i < 3; i++ {
		var loopVar = "Loop scope"
		fmt.Printf("Loop %d: %s\n", i, loopVar)
	}

	// i and loopVar are not accessible here

	// Shadowing (variable hiding)
	var shadowVar = "Outer scope"
	fmt.Printf("shadowVar (outer): %s\n", shadowVar)

	{
		var shadowVar = "Inner scope" // This shadows the outer variable
		fmt.Printf("shadowVar (inner): %s\n", shadowVar)
	}

	fmt.Printf("shadowVar (outer again): %s\n", shadowVar)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: var (function-scoped), let/const (block-scoped)
	// JavaScript: Hoisting behavior
	// Go: All variables are block-scoped
	// Go: No hoisting - variables must be declared before use
}

func demonstrateGlobalVariables() {
	fmt.Println("\n--- GLOBAL VARIABLES ---")

	// Accessing package-level variables
	fmt.Printf("globalVar: %s\n", globalVar)
	fmt.Printf("globalInt: %d\n", globalInt)

	// Accessing grouped declarations
	fmt.Printf("name: %s\n", name)
	fmt.Printf("version: %s\n", version)
	fmt.Printf("stable: %t\n", stable)

	// Accessing inferred global
	fmt.Printf("inferredGlobal: %s (type: %T)\n", inferredGlobal, inferredGlobal)

	// Modifying global variables
	globalVar = "Modified global variable"
	globalInt = 100

	fmt.Printf("Modified globalVar: %s\n", globalVar)
	fmt.Printf("Modified globalInt: %d\n", globalInt)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: var x = 10; (global scope if not in function)
	// JavaScript: window.x or global.x in Node.js
	// Go: Variables declared at package level are package-scoped
	// Go: Capitalized names are exported (public across packages)
}

// Helper function for demonstrations
func divideInts(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// VARIABLE DECLARATION SUMMARY:
//
// 1. var name type = value    (explicit type, with value)
// 2. var name type           (explicit type, zero value)
// 3. var name = value        (inferred type, with value)
// 4. name := value           (short declaration, inferred type, function-only)
//
// SCOPE RULES:
// - Package level: Accessible throughout the package
// - Function level: Accessible throughout the function
// - Block level: Accessible within the block
// - Variables must be declared before use
// - No hoisting like JavaScript
//
// BEST PRACTICES:
// - Use short declaration (:=) for local variables when type is obvious
// - Use var for zero values or when type needs to be explicit
// - Use descriptive names
// - Keep scope as narrow as possible
// - Use const for values that don't change
