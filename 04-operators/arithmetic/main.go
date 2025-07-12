package main

import (
	"fmt"
	"math"
)

// This file covers all arithmetic operators in Go
// Arithmetic operators are used for mathematical calculations

func main() {
	fmt.Println("=== GO ARITHMETIC OPERATORS - COMPLETE GUIDE ===")

	demonstrateBasicArithmetic()
	demonstrateIntegerArithmetic()
	demonstrateFloatingPointArithmetic()
	demonstrateIncrementDecrement()
	demonstrateAssignmentOperators()
	demonstrateOverflowBehavior()
	demonstrateArithmeticBestPractices()
}

func demonstrateBasicArithmetic() {
	fmt.Println("\n--- BASIC ARITHMETIC OPERATORS ---")

	// Basic arithmetic operators: +, -, *, /, %
	var a int = 15
	var b int = 4

	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("Addition: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Subtraction: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Multiplication: %d * %d = %d\n", a, b, a*b)
	fmt.Printf("Division: %d / %d = %d\n", a, b, a/b)
	fmt.Printf("Modulo: %d %% %d = %d\n", a, b, a%b)

	// Unary operators: +, -
	var positive int = +10
	var negative int = -10

	fmt.Printf("Unary plus: +10 = %d\n", positive)
	fmt.Printf("Unary minus: -10 = %d\n", negative)
	fmt.Printf("Negation: -(%d) = %d\n", positive, -positive)

	// Operator precedence demonstration
	var result int = 2 + 3*4
	fmt.Printf("Precedence: 2 + 3 * 4 = %d (multiplication first)\n", result)

	var resultParens int = (2 + 3) * 4
	fmt.Printf("With parentheses: (2 + 3) * 4 = %d\n", resultParens)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Same arithmetic operators (+, -, *, /, %)
	// JavaScript: Automatic type coercion (5 + "3" = "53")
	// Go: No automatic type coercion (compile-time error)
	// JavaScript: Division always returns float (10/3 = 3.333...)
	// Go: Integer division returns integer (10/3 = 3)
}

func demonstrateIntegerArithmetic() {
	fmt.Println("\n--- INTEGER ARITHMETIC ---")

	// Integer division (truncation)
	var dividend int = 17
	var divisor int = 5

	fmt.Printf("Integer division: %d / %d = %d (truncated)\n", dividend, divisor, dividend/divisor)
	fmt.Printf("Remainder: %d %% %d = %d\n", dividend, divisor, dividend%divisor)

	// Verify division formula: dividend = divisor * quotient + remainder
	var quotient int = dividend / divisor
	var remainder int = dividend % divisor
	var verification int = divisor*quotient + remainder

	fmt.Printf("Verification: %d * %d + %d = %d\n", divisor, quotient, remainder, verification)

	// Negative numbers in division and modulo
	fmt.Printf("Negative dividend: %d / %d = %d\n", -dividend, divisor, -dividend/divisor)
	fmt.Printf("Negative divisor: %d / %d = %d\n", dividend, -divisor, dividend/(-divisor))
	fmt.Printf("Both negative: %d / %d = %d\n", -dividend, -divisor, (-dividend)/(-divisor))

	fmt.Printf("Modulo with negative: %d %% %d = %d\n", -dividend, divisor, -dividend%divisor)
	fmt.Printf("Modulo with negative divisor: %d %% %d = %d\n", dividend, -divisor, dividend%(-divisor))

	// Different integer types
	var int8Val int8 = 100
	var int16Val int16 = 1000

	fmt.Printf("int8 arithmetic: %d + 27 = %d\n", int8Val, int8Val+27)
	fmt.Printf("int16 arithmetic: %d * 2 = %d\n", int16Val, int16Val*2)

	// Type mixing requires explicit conversion
	// var mixed = int8Val + int16Val // This would cause an error
	var mixed = int(int8Val) + int(int16Val)
	fmt.Printf("Mixed types (converted): %d + %d = %d\n", int8Val, int16Val, mixed)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: 17 / 5 = 3.4 (floating-point division)
	// JavaScript: Math.floor(17 / 5) = 3 (to get integer division)
	// Go: 17 / 5 = 3 (integer division by default)
	// Go: More explicit type handling
}

func demonstrateFloatingPointArithmetic() {
	fmt.Println("\n--- FLOATING-POINT ARITHMETIC ---")

	// Floating-point arithmetic
	var x float64 = 17.5
	var y float64 = 5.2

	fmt.Printf("x = %f, y = %f\n", x, y)
	fmt.Printf("Addition: %f + %f = %f\n", x, y, x+y)
	fmt.Printf("Subtraction: %f - %f = %f\n", x, y, x-y)
	fmt.Printf("Multiplication: %f * %f = %f\n", x, y, x*y)
	fmt.Printf("Division: %f / %f = %f\n", x, y, x/y)

	// Floating-point precision issues
	var a float64 = 0.1
	var b float64 = 0.2
	var sum float64 = a + b

	fmt.Printf("Precision issue: %f + %f = %.17f\n", a, b, sum)
	fmt.Printf("Is sum == 0.3? %t\n", sum == 0.3)

	// Comparing floating-point numbers
	var epsilon float64 = 1e-9
	var isEqual bool = math.Abs(sum-0.3) < epsilon

	fmt.Printf("Using epsilon comparison: %t\n", isEqual)

	// float32 vs float64 precision
	var float32Val float32 = 1.23456789
	var float64Val float64 = 1.23456789

	fmt.Printf("float32 precision: %.10f\n", float32Val)
	fmt.Printf("float64 precision: %.10f\n", float64Val)

	// Mathematical functions
	var angle float64 = math.Pi / 4

	fmt.Printf("sin(π/4) = %f\n", math.Sin(angle))
	fmt.Printf("cos(π/4) = %f\n", math.Cos(angle))
	fmt.Printf("sqrt(2) = %f\n", math.Sqrt(2))
	fmt.Printf("pow(2, 3) = %f\n", math.Pow(2, 3))

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: All numbers are 64-bit floating-point
	// JavaScript: Same precision issues (0.1 + 0.2 !== 0.3)
	// Go: Explicit choice between float32 and float64
	// Go: More predictable floating-point behavior
}

func demonstrateIncrementDecrement() {
	fmt.Println("\n--- INCREMENT AND DECREMENT OPERATORS ---")

	// Post-increment and post-decrement
	var counter int = 5

	fmt.Printf("Initial counter: %d\n", counter)

	counter++
	fmt.Printf("After counter++: %d\n", counter)

	counter--
	fmt.Printf("After counter--: %d\n", counter)

	// Important: Go only has post-increment/decrement
	// No pre-increment/decrement like C/C++/Java
	// ++counter and --counter are NOT valid in Go

	// Increment/decrement are statements, not expressions
	// var x = counter++  // This would cause an error
	// var y = ++counter  // This would cause an error

	// Correct way to increment and assign
	counter++
	var x int = counter
	fmt.Printf("Increment then assign: %d\n", x)

	// Increment with different integer types
	var int8Counter int8 = 100
	var int64Counter int64 = 1000000000

	fmt.Printf("int8 before: %d\n", int8Counter)
	int8Counter++
	fmt.Printf("int8 after: %d\n", int8Counter)

	fmt.Printf("int64 before: %d\n", int64Counter)
	int64Counter++
	fmt.Printf("int64 after: %d\n", int64Counter)

	// Floating-point increment/decrement
	var floatCounter float64 = 3.14

	fmt.Printf("float before: %f\n", floatCounter)
	floatCounter++
	fmt.Printf("float after: %f\n", floatCounter)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: ++x (pre-increment) and x++ (post-increment)
	// JavaScript: Both are expressions that return values
	// Go: Only x++ and x-- (post-increment/decrement)
	// Go: These are statements, not expressions
	// Go: Simpler and less error-prone
}

func demonstrateAssignmentOperators() {
	fmt.Println("\n--- ASSIGNMENT OPERATORS ---")

	// Basic assignment
	var value int = 10
	fmt.Printf("Initial value: %d\n", value)

	// Compound assignment operators
	value += 5 // equivalent to: value = value + 5
	fmt.Printf("After += 5: %d\n", value)

	value -= 3 // equivalent to: value = value - 3
	fmt.Printf("After -= 3: %d\n", value)

	value *= 2 // equivalent to: value = value * 2
	fmt.Printf("After *= 2: %d\n", value)

	value /= 4 // equivalent to: value = value / 4
	fmt.Printf("After /= 4: %d\n", value)

	value %= 3 // equivalent to: value = value % 3
	fmt.Printf("After %%= 3: %d\n", value)

	// Bitwise assignment operators
	var bits int = 12 // 1100 in binary
	fmt.Printf("Initial bits: %d (binary: %b)\n", bits, bits)

	bits &= 10 // bitwise AND assignment (1100 & 1010 = 1000)
	fmt.Printf("After &= 10: %d (binary: %b)\n", bits, bits)

	bits |= 5 // bitwise OR assignment (1000 | 0101 = 1101)
	fmt.Printf("After |= 5: %d (binary: %b)\n", bits, bits)

	bits ^= 3 // bitwise XOR assignment (1101 ^ 0011 = 1110)
	fmt.Printf("After ^= 3: %d (binary: %b)\n", bits, bits)

	bits <<= 1 // left shift assignment (1110 << 1 = 11100)
	fmt.Printf("After <<= 1: %d (binary: %b)\n", bits, bits)

	bits >>= 2 // right shift assignment (11100 >> 2 = 111)
	fmt.Printf("After >>= 2: %d (binary: %b)\n", bits, bits)

	// Assignment with different types
	var floatVal float64 = 10.5
	floatVal += 2.5
	fmt.Printf("Float assignment: %f\n", floatVal)

	var stringVal string = "Hello"
	stringVal += " World"
	fmt.Printf("String assignment: %s\n", stringVal)

	// Multiple assignment
	var a, b int = 1, 2
	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("After swap: a = %d, b = %d\n", a, b)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Same compound assignment operators
	// JavaScript: += works for both numbers and strings
	// Go: += works for numbers and strings
	// Go: More type-safe operations
}

func demonstrateOverflowBehavior() {
	fmt.Println("\n--- OVERFLOW BEHAVIOR ---")

	// Integer overflow wraps around
	var maxInt8 int8 = 127
	fmt.Printf("Max int8: %d\n", maxInt8)

	maxInt8++
	fmt.Printf("After increment (overflow): %d\n", maxInt8)

	// Underflow
	var minInt8 int8 = -128
	fmt.Printf("Min int8: %d\n", minInt8)

	minInt8--
	fmt.Printf("After decrement (underflow): %d\n", minInt8)

	// Unsigned integer overflow
	var maxUint8 uint8 = 255
	fmt.Printf("Max uint8: %d\n", maxUint8)

	maxUint8++
	fmt.Printf("After increment (overflow): %d\n", maxUint8)

	// Floating-point overflow
	var largeFloat float64 = math.MaxFloat64
	fmt.Printf("Max float64: %e\n", largeFloat)

	largeFloat *= 2
	fmt.Printf("After multiplication (overflow): %e\n", largeFloat)

	// Division by zero
	var zeroFloat float64 = 0.0

	// Integer division by zero causes panic
	// var result = 10 / 0  // This would cause a runtime panic

	// Floating-point division by zero produces infinity
	var infinityResult float64 = 10.0 / zeroFloat
	fmt.Printf("10.0 / 0.0 = %f\n", infinityResult)

	// NaN (Not a Number)
	var nanResult float64 = 0.0 / zeroFloat
	fmt.Printf("0.0 / 0.0 = %f\n", nanResult)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Numbers can become Infinity or lose precision
	// JavaScript: Division by zero produces Infinity or NaN
	// Go: Integer overflow wraps around silently
	// Go: Floating-point overflow produces ±Inf
	// Go: Integer division by zero causes panic
}

func demonstrateArithmeticBestPractices() {
	fmt.Println("\n--- ARITHMETIC BEST PRACTICES ---")

	// Type consistency
	var a int = 10
	var b int = 3
	var c int = a + b // Same types

	fmt.Printf("Same types: %d + %d = %d\n", a, b, c)

	// Explicit type conversion when needed
	var floatA float64 = 10.5
	var intB int = 3
	var result float64 = floatA + float64(intB)

	fmt.Printf("Mixed types (converted): %f + %d = %f\n", floatA, intB, result)

	// Checking for overflow before operations
	var x int8 = 100
	var y int8 = 50

	// Check if addition would overflow
	if x > 127-y {
		fmt.Printf("Addition %d + %d would overflow int8\n", x, y)
	} else {
		fmt.Printf("Safe addition: %d + %d = %d\n", x, y, x+y)
	}

	// Using appropriate numeric types
	var smallNumber int8 = 10             // Use int8 for small numbers
	var bigNumber int64 = 1000000         // Use int64 for large numbers
	var precision float64 = 3.14159265359 // Use float64 for precision

	fmt.Printf("Appropriate types: int8=%d, int64=%d, float64=%f\n",
		smallNumber, bigNumber, precision)

	// Floating-point comparison
	var val1 float64 = 0.1 + 0.2
	var val2 float64 = 0.3

	// Wrong way
	fmt.Printf("Direct comparison: %f == %f is %t\n", val1, val2, val1 == val2)

	// Right way
	var epsilon float64 = 1e-9
	var isEqual bool = math.Abs(val1-val2) < epsilon
	fmt.Printf("Epsilon comparison: %t\n", isEqual)

	// Avoiding division by zero
	var divisor int = 0
	if divisor != 0 {
		fmt.Printf("Safe division: %d / %d = %d\n", 10, divisor, 10/divisor)
	} else {
		fmt.Printf("Cannot divide by zero\n")
	}

	// Using math functions for complex operations
	var angle float64 = 45.0 * math.Pi / 180.0 // Convert degrees to radians
	fmt.Printf("sin(45°) = %f\n", math.Sin(angle))

	// BEST PRACTICES SUMMARY:
	fmt.Println("\nBest Practices Summary:")
	fmt.Println("1. Use consistent types in arithmetic operations")
	fmt.Println("2. Explicitly convert types when mixing different numeric types")
	fmt.Println("3. Be aware of integer overflow/underflow behavior")
	fmt.Println("4. Use epsilon comparison for floating-point equality")
	fmt.Println("5. Check for division by zero before dividing")
	fmt.Println("6. Choose appropriate numeric types for your data range")
	fmt.Println("7. Use math package functions for complex calculations")
	fmt.Println("8. Understand the difference between integer and floating-point division")
	fmt.Println("9. Use parentheses to make operator precedence clear")
	fmt.Println("10. Consider using constants for mathematical values")
}
