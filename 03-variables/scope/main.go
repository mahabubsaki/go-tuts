package main

import "fmt"

// This file demonstrates Go's scoping rules in detail
// Understanding scope is crucial for writing maintainable Go code

// Package-level scope (global to this package)
var packageVar = "I'm accessible throughout the package"
var PackageExported = "I'm exported (accessible from other packages)"

// Package-level constants
const packageConst = "Package-level constant"

// Function declaration at package level
func packageFunction() {
	fmt.Println("Package-level function")
}

func main() {
	fmt.Println("=== GO VARIABLE SCOPE - COMPLETE GUIDE ===")

	demonstratePackageScope()
	demonstrateFunctionScope()
	demonstrateBlockScope()
	demonstrateLoopScope()
	demonstrateConditionalScope()
	demonstrateShadowing()
	demonstrateScopeRules()
	demonstrateClosures()
}

func demonstratePackageScope() {
	fmt.Println("\n--- PACKAGE SCOPE ---")

	// Accessing package-level variables
	fmt.Printf("packageVar: %s\n", packageVar)
	fmt.Printf("PackageExported: %s\n", PackageExported)
	fmt.Printf("packageConst: %s\n", packageConst)

	// Calling package-level function
	packageFunction()

	// Modifying package-level variables
	originalValue := packageVar
	packageVar = "Modified from main function"
	fmt.Printf("Modified packageVar: %s\n", packageVar)

	// Restore original value
	packageVar = originalValue

	// Package scope rules:
	fmt.Println("\nPackage Scope Rules:")
	fmt.Println("1. Variables declared at package level are accessible throughout the package")
	fmt.Println("2. Capitalized names are exported (public to other packages)")
	fmt.Println("3. Lowercase names are private to the package")
	fmt.Println("4. Cannot use short declaration (:=) at package level")
	fmt.Println("5. Must use var or const keywords")

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: var x = 10; (global scope if not in function)
	// JavaScript: Global variables become properties of global object
	// Go: Package-level variables are scoped to the package
	// Go: Explicit export control with capitalization
}

func demonstrateFunctionScope() {
	fmt.Println("\n--- FUNCTION SCOPE ---")

	// Function-level variables
	var functionVar = "Function-scoped variable"
	functionShort := "Short declaration variable"

	fmt.Printf("functionVar: %s\n", functionVar)
	fmt.Printf("functionShort: %s\n", functionShort)

	// Function parameters are also function-scoped
	helperFunction := func(param string) {
		fmt.Printf("Parameter: %s\n", param)

		// Can access function-level variables from outer function
		fmt.Printf("Accessing functionVar from inner function: %s\n", functionVar)

		// Local variable in inner function
		var innerVar = "Inner function variable"
		fmt.Printf("innerVar: %s\n", innerVar)
	}

	helperFunction("Hello from parameter")

	// innerVar is not accessible here
	// fmt.Printf("innerVar: %s\n", innerVar) // This would cause an error

	// Function scope rules:
	fmt.Println("\nFunction Scope Rules:")
	fmt.Println("1. Variables declared in a function are accessible throughout the function")
	fmt.Println("2. Function parameters are function-scoped")
	fmt.Println("3. Variables declared in inner functions are not accessible in outer functions")
	fmt.Println("4. Inner functions can access outer function variables (closure)")
	fmt.Println("5. Can use both var and short declaration (:=)")

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: var (function-scoped), let/const (block-scoped)
	// JavaScript: Functions create their own scope
	// Go: All variables are block-scoped (including function scope)
	// Go: More consistent scoping rules
}

func demonstrateBlockScope() {
	fmt.Println("\n--- BLOCK SCOPE ---")

	var outerVar = "Outer block variable"
	fmt.Printf("outerVar: %s\n", outerVar)

	// Block scope with braces
	{
		var blockVar = "Block-scoped variable"
		innerShort := "Short declaration in block"

		fmt.Printf("blockVar: %s\n", blockVar)
		fmt.Printf("innerShort: %s\n", innerShort)

		// Can access outer variables
		fmt.Printf("Accessing outerVar from block: %s\n", outerVar)

		// Nested blocks
		{
			var nestedVar = "Nested block variable"
			fmt.Printf("nestedVar: %s\n", nestedVar)

			// Can access all outer variables
			fmt.Printf("Accessing blockVar from nested block: %s\n", blockVar)
			fmt.Printf("Accessing outerVar from nested block: %s\n", outerVar)
		}

		// nestedVar is not accessible here
		// fmt.Printf("nestedVar: %s\n", nestedVar) // This would cause an error
	}

	// blockVar and innerShort are not accessible here
	// fmt.Printf("blockVar: %s\n", blockVar) // This would cause an error

	// Block scope rules:
	fmt.Println("\nBlock Scope Rules:")
	fmt.Println("1. Variables are scoped to the block where they're declared")
	fmt.Println("2. Blocks are defined by curly braces {}")
	fmt.Println("3. Inner blocks can access outer block variables")
	fmt.Println("4. Outer blocks cannot access inner block variables")
	fmt.Println("5. Variables must be declared before use")

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: let/const are block-scoped (ES6+)
	// JavaScript: var is function-scoped (can be confusing)
	// Go: All variables are block-scoped
	// Go: More consistent and predictable scoping
}

func demonstrateLoopScope() {
	fmt.Println("\n--- LOOP SCOPE ---")

	// for loop with initialization
	fmt.Println("for loop with initialization:")
	for i := 0; i < 3; i++ {
		var loopVar = fmt.Sprintf("Loop iteration %d", i)
		fmt.Printf("  i: %d, loopVar: %s\n", i, loopVar)
	}

	// i and loopVar are not accessible here
	// fmt.Printf("i: %d\n", i) // This would cause an error

	// range loop
	fmt.Println("\nrange loop:")
	numbers := []int{10, 20, 30}
	for index, value := range numbers {
		var rangeVar = fmt.Sprintf("Index %d has value %d", index, value)
		fmt.Printf("  %s\n", rangeVar)
	}

	// index, value, and rangeVar are not accessible here

	// while-like loop
	fmt.Println("\nwhile-like loop:")
	counter := 0
	for counter < 3 {
		var whileVar = fmt.Sprintf("While iteration %d", counter)
		fmt.Printf("  %s\n", whileVar)
		counter++
	}

	// whileVar is not accessible here, but counter is (declared outside loop)
	fmt.Printf("Final counter value: %d\n", counter)

	// Loop scope rules:
	fmt.Println("\nLoop Scope Rules:")
	fmt.Println("1. Loop variables (i, index, value) are scoped to the loop")
	fmt.Println("2. Variables declared inside loop body are scoped to each iteration")
	fmt.Println("3. Variables declared outside loop are accessible inside loop")
	fmt.Println("4. Loop creates a new scope for each iteration")

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: for (let i = 0; i < 3; i++) - let is block-scoped
	// JavaScript: for (var i = 0; i < 3; i++) - var is function-scoped (can cause issues)
	// Go: Loop variables are always block-scoped
	// Go: Consistent behavior across all loop types
}

func demonstrateConditionalScope() {
	fmt.Println("\n--- CONDITIONAL SCOPE ---")

	// if statement with initialization
	if condition := true; condition {
		var ifVar = "Inside if block"
		fmt.Printf("condition: %t, ifVar: %s\n", condition, ifVar)
	}

	// condition and ifVar are not accessible here
	// fmt.Printf("condition: %t\n", condition) // This would cause an error

	// if-else with different scopes
	x := 10
	if x > 5 {
		var positiveVar = "x is greater than 5"
		fmt.Printf("positiveVar: %s\n", positiveVar)
	} else {
		var negativeVar = "x is not greater than 5"
		fmt.Printf("negativeVar: %s\n", negativeVar)
	}

	// positiveVar and negativeVar are not accessible here

	// switch statement scope
	switch value := "hello"; value {
	case "hello":
		var caseVar = "Found hello"
		fmt.Printf("value: %s, caseVar: %s\n", value, caseVar)
	case "world":
		var caseVar = "Found world" // Different caseVar in different case
		fmt.Printf("value: %s, caseVar: %s\n", value, caseVar)
	default:
		var caseVar = "Default case"
		fmt.Printf("value: %s, caseVar: %s\n", value, caseVar)
	}

	// value and caseVar are not accessible here

	// Conditional scope rules:
	fmt.Println("\nConditional Scope Rules:")
	fmt.Println("1. Variables declared in if initialization are scoped to the if statement")
	fmt.Println("2. Variables declared in if/else blocks are scoped to those blocks")
	fmt.Println("3. Variables declared in switch initialization are scoped to the switch")
	fmt.Println("4. Variables declared in case blocks are scoped to those cases")
	fmt.Println("5. Each case can have its own variables with the same name")

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: if (let condition = true) - not allowed
	// JavaScript: if blocks create scope for let/const
	// Go: if (condition := true; condition) - allowed and useful
	// Go: More flexible conditional scoping
}

func demonstrateShadowing() {
	fmt.Println("\n--- VARIABLE SHADOWING ---")

	// Outer variable
	var name = "Outer Alice"
	fmt.Printf("Outer name: %s\n", name)

	// Shadowing in block scope
	{
		var name = "Inner Alice" // This shadows the outer variable
		fmt.Printf("Inner name: %s\n", name)

		// Even more nested shadowing
		{
			var name = "Nested Alice"
			fmt.Printf("Nested name: %s\n", name)
		}

		fmt.Printf("Back to inner name: %s\n", name)
	}

	fmt.Printf("Back to outer name: %s\n", name)

	// Shadowing with different types
	var value = 42
	fmt.Printf("Original value: %d (type: %T)\n", value, value)

	{
		var value = "I'm a string now"
		fmt.Printf("Shadowed value: %s (type: %T)\n", value, value)
	}

	fmt.Printf("Original value restored: %d (type: %T)\n", value, value)

	// Shadowing package-level variables
	var packageVar = "Local packageVar" // Shadows the package-level variable
	fmt.Printf("Shadowed packageVar: %s\n", packageVar)

	// Shadowing in function parameters
	shadowingFunction := func(name string) {
		fmt.Printf("Parameter name: %s\n", name)

		// Can shadow the parameter
		{
			var name = "Block name"
			fmt.Printf("Shadowed parameter name: %s\n", name)
		}

		fmt.Printf("Parameter name restored: %s\n", name)
	}

	shadowingFunction("Parameter Alice")

	// Shadowing rules:
	fmt.Println("\nShadowing Rules:")
	fmt.Println("1. Inner scope variables can shadow outer scope variables")
	fmt.Println("2. Shadowed variables are temporarily hidden, not destroyed")
	fmt.Println("3. Original variables are restored when inner scope ends")
	fmt.Println("4. Shadowing variables can have different types")
	fmt.Println("5. Function parameters can be shadowed")
	fmt.Println("6. Package-level variables can be shadowed by local variables")

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Similar shadowing behavior with let/const
	// JavaScript: var has function scope, can cause confusion
	// Go: Consistent shadowing rules across all scopes
	// Go: More predictable behavior
}

func demonstrateScopeRules() {
	fmt.Println("\n--- SCOPE RULES SUMMARY ---")

	// Demonstration of all scope rules
	var level1 = "Level 1"
	fmt.Printf("Level 1: %s\n", level1)

	if true {
		var level2 = "Level 2"
		fmt.Printf("Level 2: %s, can access level 1: %s\n", level2, level1)

		for i := 0; i < 1; i++ {
			var level3 = "Level 3"
			fmt.Printf("Level 3: %s, can access level 2: %s, level 1: %s\n",
				level3, level2, level1)

			{
				var level4 = "Level 4"
				fmt.Printf("Level 4: %s, can access all: %s, %s, %s\n",
					level4, level3, level2, level1)
			}
		}
	}

	// Scope hierarchy visualization
	fmt.Println("\nScope Hierarchy:")
	fmt.Println("Package Scope (global to package)")
	fmt.Println("├── Function Scope")
	fmt.Println("    ├── Block Scope")
	fmt.Println("    │   ├── Nested Block Scope")
	fmt.Println("    │   └── Loop/Conditional Scope")
	fmt.Println("    └── Another Block Scope")

	// Key rules
	fmt.Println("\nKey Scope Rules:")
	fmt.Println("1. Variables are accessible within their scope and all nested scopes")
	fmt.Println("2. Variables are not accessible outside their scope")
	fmt.Println("3. Variables must be declared before use")
	fmt.Println("4. Inner scopes can shadow outer scope variables")
	fmt.Println("5. Package scope variables can be exported with capitalization")
	fmt.Println("6. No hoisting - variables exist only after declaration")

	// COMPARISON WITH JAVASCRIPT:
	fmt.Println("\nComparison with JavaScript:")
	fmt.Println("JavaScript (var): Function-scoped, hoisted, can be confusing")
	fmt.Println("JavaScript (let/const): Block-scoped, not hoisted, ES6+")
	fmt.Println("Go: Always block-scoped, never hoisted, consistent")
	fmt.Println("Go: More predictable and safer scoping rules")
}

func demonstrateClosures() {
	fmt.Println("\n--- CLOSURES AND SCOPE ---")

	// Closure example
	var outerVar = "I'm from outer function"

	innerFunction := func() {
		fmt.Printf("Closure accessing: %s\n", outerVar)

		// Modify outer variable
		outerVar = "Modified by closure"
	}

	fmt.Printf("Before closure: %s\n", outerVar)
	innerFunction()
	fmt.Printf("After closure: %s\n", outerVar)

	// Closure with parameters
	createCounter := func(start int) func() int {
		counter := start
		return func() int {
			counter++
			return counter
		}
	}

	counter1 := createCounter(0)
	counter2 := createCounter(10)

	fmt.Printf("Counter1: %d, %d, %d\n", counter1(), counter1(), counter1())
	fmt.Printf("Counter2: %d, %d, %d\n", counter2(), counter2(), counter2())

	// Closure rules:
	fmt.Println("\nClosure Rules:")
	fmt.Println("1. Functions can access variables from outer scopes")
	fmt.Println("2. Functions can modify variables from outer scopes")
	fmt.Println("3. Functions keep references to outer scope variables")
	fmt.Println("4. Each closure has its own copy of outer scope variables")
	fmt.Println("5. Closures extend the lifetime of outer scope variables")

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Similar closure behavior
	// JavaScript: Functions remember their lexical scope
	// Go: Similar closure behavior
	// Go: More explicit and predictable scoping
}

// Helper function to demonstrate scope
func helperFunction() {
	fmt.Println("Helper function can access package-level variables:")
	fmt.Printf("  packageVar: %s\n", packageVar)
	fmt.Printf("  PackageExported: %s\n", PackageExported)

	// But cannot access main function's local variables
	// fmt.Printf("  functionVar: %s\n", functionVar) // This would cause an error
}
