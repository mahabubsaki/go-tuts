package main

import (
	"fmt"
	"math"
	"unsafe"
)

// This file covers ALL floating-point types in Go in complete detail
// Go has two floating-point types: float32 and float64

func main() {
	fmt.Println("=== GO FLOATING-POINT TYPES - COMPLETE GUIDE ===")

	demonstrateFloat32()
	demonstrateFloat64()
	demonstrateFloatingPointOperations()
	demonstrateFloatingPointLimits()
	demonstrateFloatingPointPrecision()
	demonstrateSpecialValues()
	demonstrateFloatingPointComparison()
}

func demonstrateFloat32() {
	fmt.Println("\n--- FLOAT32 (Single Precision) ---")

	// float32: 32-bit floating-point number
	// IEEE 754 single precision
	// 1 sign bit, 8 exponent bits, 23 mantissa bits
	// Range: approximately ±1.4e-45 to ±3.4e38
	// About 7 decimal digits of precision

	var float32Zero float32 // zero value is 0.0
	var float32Positive float32 = 3.14159
	var float32Negative float32 = -2.71828
	var float32Scientific float32 = 1.23e-4
	var float32Large float32 = 1.5e20

	fmt.Printf("float32 - Zero: %f, Size: %d bytes\n", float32Zero, unsafe.Sizeof(float32Zero))
	fmt.Printf("float32 - Positive: %f\n", float32Positive)
	fmt.Printf("float32 - Negative: %f\n", float32Negative)
	fmt.Printf("float32 - Scientific: %f (%e)\n", float32Scientific, float32Scientific)
	fmt.Printf("float32 - Large: %f (%e)\n", float32Large, float32Large)

	// Different formatting options
	fmt.Printf("Formatting options for %f:\n", float32Positive)
	fmt.Printf("  %%f: %f\n", float32Positive)
	fmt.Printf("  %%e: %e\n", float32Positive)
	fmt.Printf("  %%g: %g\n", float32Positive)
	fmt.Printf("  %%.2f: %.2f\n", float32Positive)
	fmt.Printf("  %%.4f: %.4f\n", float32Positive)
}

func demonstrateFloat64() {
	fmt.Println("\n--- FLOAT64 (Double Precision) ---")

	// float64: 64-bit floating-point number
	// IEEE 754 double precision
	// 1 sign bit, 11 exponent bits, 52 mantissa bits
	// Range: approximately ±5.0e-324 to ±1.8e308
	// About 15-17 decimal digits of precision
	// This is the default floating-point type in Go

	var float64Zero float64 // zero value is 0.0
	var float64Positive float64 = 3.141592653589793
	var float64Negative float64 = -2.718281828459045
	var float64Scientific float64 = 1.23456789e-10
	var float64Large float64 = 1.5e100

	fmt.Printf("float64 - Zero: %f, Size: %d bytes\n", float64Zero, unsafe.Sizeof(float64Zero))
	fmt.Printf("float64 - Positive: %f\n", float64Positive)
	fmt.Printf("float64 - Negative: %f\n", float64Negative)
	fmt.Printf("float64 - Scientific: %f (%e)\n", float64Scientific, float64Scientific)
	fmt.Printf("float64 - Large: %f (%e)\n", float64Large, float64Large)

	// Higher precision compared to float32
	fmt.Printf("Precision comparison:\n")
	fmt.Printf("  float32: %.10f\n", float32(float64Positive))
	fmt.Printf("  float64: %.10f\n", float64Positive)

	// Literal types
	var inferredFloat = 3.14 // This is float64 by default
	fmt.Printf("Inferred type: %T, Value: %f\n", inferredFloat, inferredFloat)
}

func demonstrateFloatingPointOperations() {
	fmt.Println("\n--- FLOATING-POINT OPERATIONS ---")

	var a float64 = 10.5
	var b float64 = 3.2

	// Arithmetic operations
	fmt.Printf("a = %f, b = %f\n", a, b)
	fmt.Printf("Addition: %f + %f = %f\n", a, b, a+b)
	fmt.Printf("Subtraction: %f - %f = %f\n", a, b, a-b)
	fmt.Printf("Multiplication: %f * %f = %f\n", a, b, a*b)
	fmt.Printf("Division: %f / %f = %f\n", a, b, a/b)

	// Math functions
	fmt.Printf("math.Sqrt(%f) = %f\n", a, math.Sqrt(a))
	fmt.Printf("math.Pow(%f, 2) = %f\n", a, math.Pow(a, 2))
	fmt.Printf("math.Sin(%f) = %f\n", a, math.Sin(a))
	fmt.Printf("math.Cos(%f) = %f\n", a, math.Cos(a))
	fmt.Printf("math.Log(%f) = %f\n", a, math.Log(a))
	fmt.Printf("math.Exp(%f) = %f\n", a, math.Exp(a))

	// Comparison operations
	fmt.Printf("Equal: %f == %f = %t\n", a, b, a == b)
	fmt.Printf("Not equal: %f != %f = %t\n", a, b, a != b)
	fmt.Printf("Less than: %f < %f = %t\n", a, b, a < b)
	fmt.Printf("Greater than: %f > %f = %t\n", a, b, a > b)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Same arithmetic operations
	// JavaScript: Math.sqrt(), Math.pow(), Math.sin(), etc.
	// Go: math.Sqrt(), math.Pow(), math.Sin(), etc.
	// JavaScript: Numbers are always 64-bit floating-point
	// Go: You choose between float32 and float64
}

func demonstrateFloatingPointLimits() {
	fmt.Println("\n--- FLOATING-POINT LIMITS ---")

	// Constants from math package
	fmt.Printf("math.MaxFloat32: %e\n", math.MaxFloat32)
	fmt.Printf("math.SmallestNonzeroFloat32: %e\n", math.SmallestNonzeroFloat32)
	fmt.Printf("math.MaxFloat64: %e\n", math.MaxFloat64)
	fmt.Printf("math.SmallestNonzeroFloat64: %e\n", math.SmallestNonzeroFloat64)

	// Mathematical constants
	fmt.Printf("math.Pi: %.15f\n", math.Pi)
	fmt.Printf("math.E: %.15f\n", math.E)
	fmt.Printf("math.Phi: %.15f\n", math.Phi)
	fmt.Printf("math.Sqrt2: %.15f\n", math.Sqrt2)
	fmt.Printf("math.SqrtE: %.15f\n", math.SqrtE)
	fmt.Printf("math.SqrtPi: %.15f\n", math.SqrtPi)
	fmt.Printf("math.SqrtPhi: %.15f\n", math.SqrtPhi)
	fmt.Printf("math.Ln2: %.15f\n", math.Ln2)
	fmt.Printf("math.Log2E: %.15f\n", math.Log2E)
	fmt.Printf("math.Ln10: %.15f\n", math.Ln10)
	fmt.Printf("math.Log10E: %.15f\n", math.Log10E)
}

func demonstrateFloatingPointPrecision() {
	fmt.Println("\n--- FLOATING-POINT PRECISION ---")

	// Precision differences between float32 and float64
	var precise float64 = 1.23456789012345
	var lessprecise float32 = 1.23456789012345

	fmt.Printf("Original: 1.23456789012345\n")
	fmt.Printf("float64:  %.15f\n", precise)
	fmt.Printf("float32:  %.15f\n", lessprecise)

	// Precision loss in calculations
	var large float64 = 1e15
	var small float64 = 1.0

	fmt.Printf("Large number: %.1f\n", large)
	fmt.Printf("Small number: %.1f\n", small)
	fmt.Printf("Large + Small: %.1f\n", large+small)
	fmt.Printf("(Large + Small) - Large: %.1f\n", (large+small)-large)

	// Demonstrating floating-point arithmetic issues
	var result float64 = 0.1 + 0.2
	fmt.Printf("0.1 + 0.2 = %.17f\n", result)
	fmt.Printf("0.1 + 0.2 == 0.3: %t\n", result == 0.3)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript has the same floating-point precision issues
	// console.log(0.1 + 0.2); // 0.30000000000000004
	// console.log(0.1 + 0.2 === 0.3); // false
	// Both languages use IEEE 754 floating-point arithmetic
}

func demonstrateSpecialValues() {
	fmt.Println("\n--- SPECIAL VALUES ---")

	// Special floating-point values
	var positiveInfinity float64 = math.Inf(1)
	var negativeInfinity float64 = math.Inf(-1)
	var notANumber float64 = math.NaN()

	fmt.Printf("Positive infinity: %f\n", positiveInfinity)
	fmt.Printf("Negative infinity: %f\n", negativeInfinity)
	fmt.Printf("Not a number: %f\n", notANumber)

	// Testing special values
	fmt.Printf("math.IsInf(+Inf, 1): %t\n", math.IsInf(positiveInfinity, 1))
	fmt.Printf("math.IsInf(-Inf, -1): %t\n", math.IsInf(negativeInfinity, -1))
	fmt.Printf("math.IsInf(+Inf, 0): %t\n", math.IsInf(positiveInfinity, 0))
	fmt.Printf("math.IsNaN(NaN): %t\n", math.IsNaN(notANumber))

	// Operations that produce special values
	var zero float64 = 0.0
	var one float64 = 1.0
	var negativeOne float64 = -1.0

	fmt.Printf("1.0 / 0.0 = %f\n", one/zero)
	fmt.Printf("-1.0 / 0.0 = %f\n", negativeOne/zero)
	fmt.Printf("0.0 / 0.0 = %f\n", zero/zero)
	fmt.Printf("math.Sqrt(-1) = %f\n", math.Sqrt(-1))

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Infinity, -Infinity, NaN
	// JavaScript: isFinite(), isNaN(), Number.isInfinite()
	// Go: math.Inf(), math.NaN(), math.IsInf(), math.IsNaN()
	// Both handle special values similarly
}

func demonstrateFloatingPointComparison() {
	fmt.Println("\n--- FLOATING-POINT TYPE COMPARISON ---")

	// WHEN TO USE EACH TYPE:

	fmt.Println("float32 (32-bit, ~7 decimal digits):")
	fmt.Println("  - Memory-constrained applications")
	fmt.Println("  - Graphics programming")
	fmt.Println("  - Scientific computing with limited precision needs")
	fmt.Println("  - Embedded systems")
	fmt.Println("  - Network protocols requiring specific precision")

	fmt.Println("\nfloat64 (64-bit, ~15-17 decimal digits):")
	fmt.Println("  - Default choice for floating-point operations")
	fmt.Println("  - High-precision calculations")
	fmt.Println("  - Financial calculations")
	fmt.Println("  - Scientific computing")
	fmt.Println("  - Most mathematical operations")

	fmt.Println("\n--- JAVASCRIPT COMPARISON ---")
	fmt.Println("JavaScript:")
	fmt.Println("  - Single 'number' type (64-bit floating-point)")
	fmt.Println("  - Always double precision")
	fmt.Println("  - No choice of precision")
	fmt.Println("  - let x = 3.14; (always 64-bit)")

	fmt.Println("\nGo:")
	fmt.Println("  - Two floating-point types with different precision")
	fmt.Println("  - Explicit type declaration")
	fmt.Println("  - Choose precision based on needs")
	fmt.Println("  - var x float64 = 3.14; (explicit precision)")

	// Performance comparison
	fmt.Println("\n--- PERFORMANCE CONSIDERATIONS ---")
	fmt.Println("float32:")
	fmt.Println("  - Uses half the memory of float64")
	fmt.Println("  - Faster on some architectures")
	fmt.Println("  - Better cache performance")
	fmt.Println("  - Limited precision may cause errors")

	fmt.Println("\nfloat64:")
	fmt.Println("  - Uses more memory")
	fmt.Println("  - Standard precision for most calculations")
	fmt.Println("  - Better accuracy for complex calculations")
	fmt.Println("  - Recommended for most use cases")

	// Best practices
	fmt.Println("\n--- BEST PRACTICES ---")
	fmt.Println("1. Use float64 unless you have specific reasons to use float32")
	fmt.Println("2. Be aware of floating-point precision limitations")
	fmt.Println("3. Use math.Abs() for floating-point comparisons")
	fmt.Println("4. Consider using decimal packages for financial calculations")
	fmt.Println("5. Use appropriate formatting for display")
}
