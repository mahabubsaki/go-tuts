package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"unsafe"
)

// This file covers complex number types in Go
// Go has built-in support for complex numbers, unlike JavaScript
// Complex numbers are useful in mathematical calculations, signal processing, etc.

func main() {
	fmt.Println("=== GO COMPLEX TYPES - COMPLETE GUIDE ===")

	demonstrateComplex64()
	demonstrateComplex128()
	demonstrateComplexOperations()
	demonstrateComplexFunctions()
	demonstrateComplexComparison()
}

func demonstrateComplex64() {
	fmt.Println("\n--- COMPLEX64 (Single Precision Complex) ---")

	// complex64: complex number with float32 real and imaginary parts
	// 64 bits total (32 bits for real, 32 bits for imaginary)
	// Less precise but uses less memory

	var complex64Zero complex64               // zero value is (0+0i)
	var complex64Real complex64 = 3.14        // real number (imaginary part is 0)
	var complex64Imaginary complex64 = 2.5i   // imaginary number (real part is 0)
	var complex64Full complex64 = 3.14 + 2.5i // full complex number

	fmt.Printf("complex64 - Zero: %v, Size: %d bytes\n", complex64Zero, unsafe.Sizeof(complex64Zero))
	fmt.Printf("complex64 - Real only: %v\n", complex64Real)
	fmt.Printf("complex64 - Imaginary only: %v\n", complex64Imaginary)
	fmt.Printf("complex64 - Full: %v\n", complex64Full)

	// Creating complex numbers using complex() function
	var realPart float32 = 1.5
	var imaginaryPart float32 = 2.5
	var complexFromParts complex64 = complex(realPart, imaginaryPart)

	fmt.Printf("complex(1.5, 2.5): %v\n", complexFromParts)

	// Extracting real and imaginary parts
	fmt.Printf("Real part: %f\n", real(complex64Full))
	fmt.Printf("Imaginary part: %f\n", imag(complex64Full))
}

func demonstrateComplex128() {
	fmt.Println("\n--- COMPLEX128 (Double Precision Complex) ---")

	// complex128: complex number with float64 real and imaginary parts
	// 128 bits total (64 bits for real, 64 bits for imaginary)
	// Higher precision, default complex type

	var complex128Zero complex128                                          // zero value is (0+0i)
	var complex128Real complex128 = 3.141592653589793                      // real number
	var complex128Imaginary complex128 = 2.718281828459045i                // imaginary number
	var complex128Full complex128 = 3.141592653589793 + 2.718281828459045i // full complex number

	fmt.Printf("complex128 - Zero: %v, Size: %d bytes\n", complex128Zero, unsafe.Sizeof(complex128Zero))
	fmt.Printf("complex128 - Real only: %v\n", complex128Real)
	fmt.Printf("complex128 - Imaginary only: %v\n", complex128Imaginary)
	fmt.Printf("complex128 - Full: %v\n", complex128Full)

	// Creating complex numbers using complex() function
	var realPart float64 = math.Pi
	var imaginaryPart float64 = math.E
	var complexFromParts complex128 = complex(realPart, imaginaryPart)

	fmt.Printf("complex(π, e): %v\n", complexFromParts)

	// Extracting real and imaginary parts
	fmt.Printf("Real part: %f\n", real(complex128Full))
	fmt.Printf("Imaginary part: %f\n", imag(complex128Full))

	// Literal types
	var inferredComplex = 3.14 + 2.5i // This is complex128 by default
	fmt.Printf("Inferred type: %T, Value: %v\n", inferredComplex, inferredComplex)
}

func demonstrateComplexOperations() {
	fmt.Println("\n--- COMPLEX OPERATIONS ---")

	var a complex128 = 3 + 4i
	var b complex128 = 1 + 2i

	fmt.Printf("a = %v, b = %v\n", a, b)

	// Arithmetic operations
	fmt.Printf("Addition: %v + %v = %v\n", a, b, a+b)
	fmt.Printf("Subtraction: %v - %v = %v\n", a, b, a-b)
	fmt.Printf("Multiplication: %v * %v = %v\n", a, b, a*b)
	fmt.Printf("Division: %v / %v = %v\n", a, b, a/b)

	// Comparison operations
	// Note: Only == and != are allowed for complex numbers
	fmt.Printf("Equal: %v == %v = %t\n", a, b, a == b)
	fmt.Printf("Not equal: %v != %v = %t\n", a, b, a != b)

	// You cannot use <, >, <=, >= with complex numbers
	// fmt.Printf("Less than: %v < %v = %t\n", a, b, a < b) // This would cause an error

	// Operations with real numbers
	var realNumber complex128 = 2.0
	fmt.Printf("Complex * Real: %v * %v = %v\n", a, realNumber, a*realNumber)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript does not have built-in complex number support
	// You would need to use a library or implement your own
	// Go has native support for complex numbers
}

func demonstrateComplexFunctions() {
	fmt.Println("\n--- COMPLEX FUNCTIONS ---")

	var z complex128 = 3 + 4i

	// Basic complex functions from cmplx package
	fmt.Printf("z = %v\n", z)
	fmt.Printf("cmplx.Abs(z) = %f (magnitude)\n", cmplx.Abs(z))
	fmt.Printf("cmplx.Phase(z) = %f (phase angle in radians)\n", cmplx.Phase(z))
	fmt.Printf("cmplx.Conj(z) = %v (complex conjugate)\n", cmplx.Conj(z))

	// Exponential and logarithmic functions
	fmt.Printf("cmplx.Exp(z) = %v\n", cmplx.Exp(z))
	fmt.Printf("cmplx.Log(z) = %v\n", cmplx.Log(z))
	fmt.Printf("cmplx.Log10(z) = %v\n", cmplx.Log10(z))

	// Power functions
	fmt.Printf("cmplx.Pow(z, 2) = %v\n", cmplx.Pow(z, 2))
	fmt.Printf("cmplx.Sqrt(z) = %v\n", cmplx.Sqrt(z))

	// Trigonometric functions
	fmt.Printf("cmplx.Sin(z) = %v\n", cmplx.Sin(z))
	fmt.Printf("cmplx.Cos(z) = %v\n", cmplx.Cos(z))
	fmt.Printf("cmplx.Tan(z) = %v\n", cmplx.Tan(z))

	// Hyperbolic functions
	fmt.Printf("cmplx.Sinh(z) = %v\n", cmplx.Sinh(z))
	fmt.Printf("cmplx.Cosh(z) = %v\n", cmplx.Cosh(z))
	fmt.Printf("cmplx.Tanh(z) = %v\n", cmplx.Tanh(z))

	// Inverse trigonometric functions
	fmt.Printf("cmplx.Asin(z) = %v\n", cmplx.Asin(z))
	fmt.Printf("cmplx.Acos(z) = %v\n", cmplx.Acos(z))
	fmt.Printf("cmplx.Atan(z) = %v\n", cmplx.Atan(z))

	// Creating complex numbers from polar coordinates
	var magnitude float64 = 5.0
	var angle float64 = math.Pi / 4 // 45 degrees
	var polar complex128 = cmplx.Rect(magnitude, angle)
	fmt.Printf("Polar form (r=%f, θ=%f): %v\n", magnitude, angle, polar)

	// Special values
	var inf complex128 = cmplx.Inf()
	var nan complex128 = cmplx.NaN()
	fmt.Printf("Complex infinity: %v\n", inf)
	fmt.Printf("Complex NaN: %v\n", nan)
	fmt.Printf("IsInf: %t\n", cmplx.IsInf(inf))
	fmt.Printf("IsNaN: %t\n", cmplx.IsNaN(nan))
}

func demonstrateComplexComparison() {
	fmt.Println("\n--- COMPLEX TYPE COMPARISON ---")

	// WHEN TO USE EACH TYPE:

	fmt.Println("complex64 (32-bit real + 32-bit imaginary):")
	fmt.Println("  - Memory-constrained applications")
	fmt.Println("  - Graphics and game programming")
	fmt.Println("  - DSP applications with limited precision needs")
	fmt.Println("  - Embedded systems")
	fmt.Println("  - Network protocols")

	fmt.Println("\ncomplex128 (64-bit real + 64-bit imaginary):")
	fmt.Println("  - Default choice for complex operations")
	fmt.Println("  - Scientific computing")
	fmt.Println("  - Mathematical calculations")
	fmt.Println("  - Signal processing")
	fmt.Println("  - Engineering applications")

	fmt.Println("\n--- JAVASCRIPT COMPARISON ---")
	fmt.Println("JavaScript:")
	fmt.Println("  - No built-in complex number support")
	fmt.Println("  - Must use libraries like math.js, ml-matrix, etc.")
	fmt.Println("  - Or implement your own complex number class")
	fmt.Println("  - Example: new Complex(3, 4) or {real: 3, imag: 4}")

	fmt.Println("\nGo:")
	fmt.Println("  - Built-in complex number support")
	fmt.Println("  - Native arithmetic operations")
	fmt.Println("  - Comprehensive math/cmplx package")
	fmt.Println("  - Two precision levels: complex64 and complex128")
	fmt.Println("  - Literal syntax: 3 + 4i")

	// Use cases
	fmt.Println("\n--- COMMON USE CASES ---")
	fmt.Println("1. Signal Processing:")
	fmt.Println("   - FFT (Fast Fourier Transform)")
	fmt.Println("   - Digital filters")
	fmt.Println("   - Audio processing")

	fmt.Println("\n2. Mathematical Calculations:")
	fmt.Println("   - Complex analysis")
	fmt.Println("   - Polynomial root finding")
	fmt.Println("   - Electrical engineering")

	fmt.Println("\n3. Computer Graphics:")
	fmt.Println("   - 2D transformations")
	fmt.Println("   - Fractals (Mandelbrot set)")
	fmt.Println("   - Rotation calculations")

	fmt.Println("\n4. Physics Simulations:")
	fmt.Println("   - Quantum mechanics")
	fmt.Println("   - Wave equations")
	fmt.Println("   - Electromagnetic fields")

	// Best practices
	fmt.Println("\n--- BEST PRACTICES ---")
	fmt.Println("1. Use complex128 unless memory is a concern")
	fmt.Println("2. Use cmplx package for complex mathematical functions")
	fmt.Println("3. Be aware that only == and != comparisons are allowed")
	fmt.Println("4. Use cmplx.Abs() for magnitude comparisons")
	fmt.Println("5. Consider real and imaginary parts separately for ordering")

	// Examples of complex number applications
	fmt.Println("\n--- PRACTICAL EXAMPLES ---")

	// Example 1: Roots of unity
	fmt.Println("Cube roots of unity:")
	for k := 0; k < 3; k++ {
		angle := 2 * math.Pi * float64(k) / 3
		root := cmplx.Exp(complex(0, angle))
		fmt.Printf("  Root %d: %v\n", k+1, root)
	}

	// Example 2: Mandelbrot set point
	fmt.Println("\nMandelbrot set test for c = -0.5 + 0.5i:")
	c := complex(-0.5, 0.5)
	z := complex(0, 0)
	for i := 0; i < 5; i++ {
		z = z*z + c
		fmt.Printf("  Iteration %d: z = %v, |z| = %f\n", i+1, z, cmplx.Abs(z))
	}
}
