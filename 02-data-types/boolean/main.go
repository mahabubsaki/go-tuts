package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

// This file covers the boolean type in Go
// Boolean logic is fundamental in programming
// Go's boolean type is similar to JavaScript's but with stricter rules

func main() {
	fmt.Println("=== GO BOOLEAN TYPE - COMPLETE GUIDE ===")

	demonstrateBooleanBasics()
	demonstrateBooleanOperations()
	demonstrateBooleanConversions()
	demonstrateBooleanInConditions()
	demonstrateBooleanComparison()
}

func demonstrateBooleanBasics() {
	fmt.Println("\n--- BOOLEAN BASICS ---")

	// Boolean type: bool
	// Only two possible values: true or false
	// Zero value is false
	// Size is typically 1 byte (8 bits) but only uses 1 bit of information

	var boolZero bool // zero value is false
	var boolTrue bool = true
	var boolFalse bool = false

	fmt.Printf("bool - Zero value: %t, Size: %d bytes\n", boolZero, unsafe.Sizeof(boolZero))
	fmt.Printf("bool - True: %t\n", boolTrue)
	fmt.Printf("bool - False: %t\n", boolFalse)

	// Boolean literals
	var literalTrue = true   // inferred as bool
	var literalFalse = false // inferred as bool

	fmt.Printf("Literal true type: %T, value: %t\n", literalTrue, literalTrue)
	fmt.Printf("Literal false type: %T, value: %t\n", literalFalse, literalFalse)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: boolean type with true/false values
	// JavaScript: Also has "truthy" and "falsy" values
	// Go: Only true and false are valid boolean values
	// Go: No implicit boolean conversion like JavaScript
}

func demonstrateBooleanOperations() {
	fmt.Println("\n--- BOOLEAN OPERATIONS ---")

	var a bool = true
	var b bool = false

	fmt.Printf("a = %t, b = %t\n", a, b)

	// Logical operations
	fmt.Printf("Logical AND: %t && %t = %t\n", a, b, a && b)
	fmt.Printf("Logical OR: %t || %t = %t\n", a, b, a || b)
	fmt.Printf("Logical NOT: !%t = %t\n", a, !a)
	fmt.Printf("Logical NOT: !%t = %t\n", b, !b)

	// Truth table for AND
	fmt.Println("\nAND Truth Table:")
	fmt.Printf("true && true = %t\n", true && true)
	fmt.Printf("true && false = %t\n", true && false)
	fmt.Printf("false && true = %t\n", false && true)
	fmt.Printf("false && false = %t\n", false && false)

	// Truth table for OR
	fmt.Println("\nOR Truth Table:")
	fmt.Printf("true || true = %t\n", true || true)
	fmt.Printf("true || false = %t\n", true || false)
	fmt.Printf("false || true = %t\n", false || true)
	fmt.Printf("false || false = %t\n", false || false)

	// Short-circuit evaluation
	fmt.Println("\nShort-circuit evaluation:")

	// Demonstrating short-circuit evaluation with safe expressions
	var x int = 10
	var y int = 0

	// In JavaScript: true || (x/y > 0) - second part not evaluated
	// In Go: true || (x > y) - second part not evaluated due to short-circuit
	fmt.Printf("true || (x > y) = %t (second expression not evaluated)\n", true || (x > y))
	fmt.Printf("false && (x > y) = %t (second expression not evaluated)\n", false && (x > y))

	// Complex boolean expressions
	var c bool = true
	var d bool = false

	fmt.Printf("Complex expression: (%t && %t) || (%t && %t) = %t\n", a, b, c, d, (a && b) || (c && d))
	fmt.Printf("De Morgan's law: !(%t && %t) = %t, (!%t || !%t) = %t\n", a, b, !(a && b), a, b, !a || !b)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Same logical operators (&&, ||, !)
	// JavaScript: Supports truthy/falsy values in logical operations
	// Go: Only works with actual boolean values
	// JavaScript: 0 && true = 0 (falsy)
	// Go: Must use explicit boolean conversion
}

func demonstrateBooleanConversions() {
	fmt.Println("\n--- BOOLEAN CONVERSIONS ---")

	// Go does not have implicit boolean conversion
	// You must explicitly convert other types to boolean

	// String to boolean conversion
	var strTrue string = "true"
	var strFalse string = "false"
	var strInvalid string = "not a boolean"

	if boolFromStr, err := strconv.ParseBool(strTrue); err == nil {
		fmt.Printf("String '%s' to bool: %t\n", strTrue, boolFromStr)
	}

	if boolFromStr, err := strconv.ParseBool(strFalse); err == nil {
		fmt.Printf("String '%s' to bool: %t\n", strFalse, boolFromStr)
	}

	if _, err := strconv.ParseBool(strInvalid); err != nil {
		fmt.Printf("String '%s' to bool: ERROR - %v\n", strInvalid, err)
	}

	// Boolean to string conversion
	var boolVal bool = true
	var boolStr string = strconv.FormatBool(boolVal)
	fmt.Printf("Bool %t to string: '%s'\n", boolVal, boolStr)

	// Integer to boolean (manual conversion)
	var intVal int = 0
	var boolFromInt bool = intVal != 0
	fmt.Printf("Int %d to bool: %t\n", intVal, boolFromInt)

	intVal = 42
	boolFromInt = intVal != 0
	fmt.Printf("Int %d to bool: %t\n", intVal, boolFromInt)

	// Boolean to integer (manual conversion)
	var boolToInt int
	if boolVal {
		boolToInt = 1
	} else {
		boolToInt = 0
	}
	fmt.Printf("Bool %t to int: %d\n", boolVal, boolToInt)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Implicit boolean conversion
	// JavaScript: if (0) {} // falsy
	// JavaScript: if ("") {} // falsy
	// JavaScript: if ([]) {} // truthy
	// Go: if condition {} // condition must be bool
	// Go: Explicit conversion required
}

func demonstrateBooleanInConditions() {
	fmt.Println("\n--- BOOLEAN IN CONDITIONS ---")

	var isReady bool = true
	var isComplete bool = false
	var count int = 10

	// If statements
	if isReady {
		fmt.Println("System is ready!")
	}

	if !isComplete {
		fmt.Println("Task is not complete yet")
	}

	// Comparison operations return boolean
	if count > 5 {
		fmt.Printf("Count (%d) is greater than 5\n", count)
	}

	if count >= 10 && count <= 20 {
		fmt.Printf("Count (%d) is between 10 and 20\n", count)
	}

	// Boolean variables in switch statements
	switch isReady {
	case true:
		fmt.Println("Ready case")
	case false:
		fmt.Println("Not ready case")
	}

	// For loops with boolean conditions
	fmt.Println("Countdown:")
	for i := 3; i > 0; i-- {
		fmt.Printf("  %d\n", i)
	}

	// While-like loop using for
	var running bool = true
	var counter int = 0

	fmt.Println("While-like loop:")
	for running {
		counter++
		fmt.Printf("  Iteration %d\n", counter)
		if counter >= 3 {
			running = false
		}
	}

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: if (value) {} // truthy check
	// JavaScript: if (0) {} // false
	// JavaScript: if ("") {} // false
	// JavaScript: if ([]) {} // true
	// Go: if condition {} // condition must be bool
	// Go: More explicit and safer
}

func demonstrateBooleanComparison() {
	fmt.Println("\n--- BOOLEAN COMPARISON ---")

	// Boolean comparison
	var bool1 bool = true
	var bool2 bool = false
	var bool3 bool = true

	fmt.Printf("bool1 = %t, bool2 = %t, bool3 = %t\n", bool1, bool2, bool3)
	fmt.Printf("bool1 == bool2: %t\n", bool1 == bool2)
	fmt.Printf("bool1 == bool3: %t\n", bool1 == bool3)
	fmt.Printf("bool1 != bool2: %t\n", bool1 != bool2)

	// You cannot use <, >, <=, >= with boolean values
	// fmt.Printf("bool1 < bool2: %t\n", bool1 < bool2) // This would cause an error

	// Boolean ordering (if needed, convert to int)
	var bool1Int int = 0
	var bool2Int int = 0

	if bool1 {
		bool1Int = 1
	}
	if bool2 {
		bool2Int = 1
	}

	fmt.Printf("bool1 as int: %d, bool2 as int: %d\n", bool1Int, bool2Int)
	fmt.Printf("bool1 < bool2 (as ints): %t\n", bool1Int < bool2Int)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Boolean comparison works the same way
	// JavaScript: true == 1 (true with type coercion)
	// JavaScript: true === 1 (false with strict comparison)
	// Go: No implicit type coercion
	// Go: Only == and != operators for boolean
}

// PRACTICAL EXAMPLES AND PATTERNS
func demonstratePracticalPatterns() {
	fmt.Println("\n--- PRACTICAL PATTERNS ---")

	// Flag pattern
	var isDebug bool = true
	var isProduction bool = false

	if isDebug && !isProduction {
		fmt.Println("Debug mode enabled")
	}

	// State machine pattern
	var isConnected bool = false
	var isAuthenticated bool = false

	switch {
	case !isConnected:
		fmt.Println("State: Disconnected")
	case isConnected && !isAuthenticated:
		fmt.Println("State: Connected but not authenticated")
	case isConnected && isAuthenticated:
		fmt.Println("State: Connected and authenticated")
	}

	// Validation pattern
	var hasValidEmail bool = true
	var hasValidPassword bool = true
	var isAgeValid bool = true

	var isFormValid bool = hasValidEmail && hasValidPassword && isAgeValid

	if isFormValid {
		fmt.Println("Form is valid, can submit")
	} else {
		fmt.Println("Form has errors:")
		if !hasValidEmail {
			fmt.Println("  - Invalid email")
		}
		if !hasValidPassword {
			fmt.Println("  - Invalid password")
		}
		if !isAgeValid {
			fmt.Println("  - Invalid age")
		}
	}

	// Boolean accumulator pattern
	var conditions = []bool{true, false, true, true}
	var allTrue bool = true
	var anyTrue bool = false

	for _, condition := range conditions {
		allTrue = allTrue && condition
		anyTrue = anyTrue || condition
	}

	fmt.Printf("All conditions true: %t\n", allTrue)
	fmt.Printf("Any condition true: %t\n", anyTrue)

	// BEST PRACTICES:
	fmt.Println("\n--- BEST PRACTICES ---")
	fmt.Println("1. Use descriptive boolean variable names (is*, has*, can*, should*)")
	fmt.Println("2. Avoid double negatives (!isNotReady)")
	fmt.Println("3. Use explicit boolean expressions instead of implicit conversions")
	fmt.Println("4. Group related boolean conditions for readability")
	fmt.Println("5. Use constants for boolean flags that don't change")
	fmt.Println("6. Consider using enums for complex state instead of multiple booleans")
}
