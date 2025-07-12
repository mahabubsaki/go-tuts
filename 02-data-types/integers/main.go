package main

import (
	"fmt"
	"math"
	"unsafe"
)

// This file covers ALL integer types in Go in complete detail
// Go has many integer types with specific sizes and ranges

func main() {
	fmt.Println("=== GO INTEGER TYPES - COMPLETE GUIDE ===")

	demonstrateSignedIntegers()
	demonstrateUnsignedIntegers()
	demonstrateIntegerAliases()
	demonstrateIntegerOperations()
	demonstrateIntegerLimits()
	demonstrateIntegerComparison()
}

func demonstrateSignedIntegers() {
	fmt.Println("\n--- SIGNED INTEGERS ---")

	// SIGNED INTEGERS can hold positive and negative values
	// The first bit is used for the sign

	// int8: -128 to 127 (8 bits, 1 byte)
	var int8Min int8 = -128
	var int8Max int8 = 127
	var int8Zero int8 // zero value is 0

	fmt.Printf("int8 - Min: %d, Max: %d, Zero: %d, Size: %d bytes\n",
		int8Min, int8Max, int8Zero, unsafe.Sizeof(int8Zero))

	// int16: -32,768 to 32,767 (16 bits, 2 bytes)
	var int16Min int16 = -32768
	var int16Max int16 = 32767
	var int16Zero int16

	fmt.Printf("int16 - Min: %d, Max: %d, Zero: %d, Size: %d bytes\n",
		int16Min, int16Max, int16Zero, unsafe.Sizeof(int16Zero))

	// int32: -2,147,483,648 to 2,147,483,647 (32 bits, 4 bytes)
	var int32Min int32 = -2147483648
	var int32Max int32 = 2147483647
	var int32Zero int32

	fmt.Printf("int32 - Min: %d, Max: %d, Zero: %d, Size: %d bytes\n",
		int32Min, int32Max, int32Zero, unsafe.Sizeof(int32Zero))

	// int64: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 (64 bits, 8 bytes)
	var int64Min int64 = -9223372036854775808
	var int64Max int64 = 9223372036854775807
	var int64Zero int64

	fmt.Printf("int64 - Min: %d, Max: %d, Zero: %d, Size: %d bytes\n",
		int64Min, int64Max, int64Zero, unsafe.Sizeof(int64Zero))

	// int: size depends on the platform (32-bit or 64-bit)
	// On 64-bit systems: same as int64
	// On 32-bit systems: same as int32
	var intZero int
	fmt.Printf("int - Zero: %d, Size: %d bytes (platform dependent)\n",
		intZero, unsafe.Sizeof(intZero))
}

func demonstrateUnsignedIntegers() {
	fmt.Println("\n--- UNSIGNED INTEGERS ---")

	// UNSIGNED INTEGERS can only hold non-negative values (0 and positive)
	// All bits are used for the value (no sign bit)

	// uint8: 0 to 255 (8 bits, 1 byte)
	var uint8Min uint8 = 0
	var uint8Max uint8 = 255
	var uint8Zero uint8

	fmt.Printf("uint8 - Min: %d, Max: %d, Zero: %d, Size: %d bytes\n",
		uint8Min, uint8Max, uint8Zero, unsafe.Sizeof(uint8Zero))

	// uint16: 0 to 65,535 (16 bits, 2 bytes)
	var uint16Min uint16 = 0
	var uint16Max uint16 = 65535
	var uint16Zero uint16

	fmt.Printf("uint16 - Min: %d, Max: %d, Zero: %d, Size: %d bytes\n",
		uint16Min, uint16Max, uint16Zero, unsafe.Sizeof(uint16Zero))

	// uint32: 0 to 4,294,967,295 (32 bits, 4 bytes)
	var uint32Min uint32 = 0
	var uint32Max uint32 = 4294967295
	var uint32Zero uint32

	fmt.Printf("uint32 - Min: %d, Max: %d, Zero: %d, Size: %d bytes\n",
		uint32Min, uint32Max, uint32Zero, unsafe.Sizeof(uint32Zero))

	// uint64: 0 to 18,446,744,073,709,551,615 (64 bits, 8 bytes)
	var uint64Min uint64 = 0
	var uint64Max uint64 = 18446744073709551615
	var uint64Zero uint64

	fmt.Printf("uint64 - Min: %d, Max: %d, Zero: %d, Size: %d bytes\n",
		uint64Min, uint64Max, uint64Zero, unsafe.Sizeof(uint64Zero))

	// uint: size depends on the platform (32-bit or 64-bit)
	var uintZero uint
	fmt.Printf("uint - Zero: %d, Size: %d bytes (platform dependent)\n",
		uintZero, unsafe.Sizeof(uintZero))
}

func demonstrateIntegerAliases() {
	fmt.Println("\n--- INTEGER ALIASES ---")

	// Go provides some convenient aliases for common integer types

	// byte is an alias for uint8
	// commonly used for binary data, file I/O, network protocols
	var byteValue byte = 255
	var uint8Value uint8 = 255

	fmt.Printf("byte: %d, uint8: %d (they're the same type)\n", byteValue, uint8Value)

	// rune is an alias for int32
	// used for Unicode code points (characters)
	var runeValue rune = 'A'  // Unicode code point for 'A'
	var int32Value int32 = 65 // Same as 'A'

	fmt.Printf("rune: %d (%c), int32: %d\n", runeValue, runeValue, int32Value)

	// uintptr: unsigned integer large enough to hold a pointer
	// used for low-level programming, usually with unsafe package
	var uintptrZero uintptr
	fmt.Printf("uintptr - Zero: %d, Size: %d bytes\n", uintptrZero, unsafe.Sizeof(uintptrZero))
}

func demonstrateIntegerOperations() {
	fmt.Println("\n--- INTEGER OPERATIONS ---")

	var a int = 10
	var b int = 3

	// Arithmetic operations
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("Addition: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Subtraction: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Multiplication: %d * %d = %d\n", a, b, a*b)
	fmt.Printf("Division: %d / %d = %d (integer division)\n", a, b, a/b)
	fmt.Printf("Modulo: %d %% %d = %d\n", a, b, a%b)

	// Comparison operations
	fmt.Printf("Equal: %d == %d = %t\n", a, b, a == b)
	fmt.Printf("Not equal: %d != %d = %t\n", a, b, a != b)
	fmt.Printf("Less than: %d < %d = %t\n", a, b, a < b)
	fmt.Printf("Greater than: %d > %d = %t\n", a, b, a > b)
	fmt.Printf("Less or equal: %d <= %d = %t\n", a, b, a <= b)
	fmt.Printf("Greater or equal: %d >= %d = %t\n", a, b, a >= b)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: 10 / 3 = 3.3333... (floating-point division)
	// Go: 10 / 3 = 3 (integer division)
	// For floating-point division in Go, use float types
}

func demonstrateIntegerLimits() {
	fmt.Println("\n--- INTEGER LIMITS AND CONSTANTS ---")

	// Go provides constants for integer limits in the math package
	fmt.Printf("math.MaxInt8: %d\n", math.MaxInt8)
	fmt.Printf("math.MinInt8: %d\n", math.MinInt8)
	fmt.Printf("math.MaxInt16: %d\n", math.MaxInt16)
	fmt.Printf("math.MinInt16: %d\n", math.MinInt16)
	fmt.Printf("math.MaxInt32: %d\n", math.MaxInt32)
	fmt.Printf("math.MinInt32: %d\n", math.MinInt32)
	fmt.Printf("math.MaxInt64: %d\n", math.MaxInt64)
	fmt.Printf("math.MinInt64: %d\n", math.MinInt64)

	fmt.Printf("math.MaxUint8: %d\n", math.MaxUint8)
	fmt.Printf("math.MaxUint16: %d\n", math.MaxUint16)
	fmt.Printf("math.MaxUint32: %d\n", math.MaxUint32)
	// Use %d verb with explicit conversion for large uint64
	fmt.Printf("math.MaxUint64: %d\n", uint64(math.MaxUint64))

	// OVERFLOW BEHAVIOR:
	// Go integers wrap around on overflow (like C)
	// JavaScript numbers can become Infinity or lose precision

	var maxInt8 int8 = 127
	fmt.Printf("maxInt8: %d\n", maxInt8)
	maxInt8++ // This will wrap around to -128
	fmt.Printf("maxInt8 after increment: %d (wrapped around)\n", maxInt8)
}

func demonstrateIntegerComparison() {
	fmt.Println("\n--- INTEGER TYPE COMPARISON ---")

	// WHEN TO USE EACH TYPE:

	fmt.Println("int8 (-128 to 127):")
	fmt.Println("  - Small integers with known range")
	fmt.Println("  - Memory-constrained applications")
	fmt.Println("  - Protocol implementations")

	fmt.Println("\nint16 (-32,768 to 32,767):")
	fmt.Println("  - Small to medium integers")
	fmt.Println("  - Audio samples, coordinates")

	fmt.Println("\nint32 (-2 billion to 2 billion):")
	fmt.Println("  - Most common integer operations")
	fmt.Println("  - IDs, counts, indices")
	fmt.Println("  - Compatible with many C APIs")

	fmt.Println("\nint64 (-9 quintillion to 9 quintillion):")
	fmt.Println("  - Large numbers")
	fmt.Println("  - Timestamps, file sizes")
	fmt.Println("  - High-precision calculations")

	fmt.Println("\nint (platform dependent):")
	fmt.Println("  - Default choice for integer operations")
	fmt.Println("  - 32-bit on 32-bit systems")
	fmt.Println("  - 64-bit on 64-bit systems")

	fmt.Println("\nUnsigned variants (uint8, uint16, uint32, uint64, uint):")
	fmt.Println("  - When you only need non-negative values")
	fmt.Println("  - Bit manipulation")
	fmt.Println("  - Memory addresses, array indices")
	fmt.Println("  - Doubles the positive range")

	fmt.Println("\nbyte (alias for uint8):")
	fmt.Println("  - Binary data, file I/O")
	fmt.Println("  - Network protocols")
	fmt.Println("  - Raw memory operations")

	fmt.Println("\nrune (alias for int32):")
	fmt.Println("  - Unicode characters")
	fmt.Println("  - Text processing")
	fmt.Println("  - Internationalization")

	// JAVASCRIPT COMPARISON:
	fmt.Println("\n--- JAVASCRIPT COMPARISON ---")
	fmt.Println("JavaScript:")
	fmt.Println("  - Single 'number' type (64-bit floating-point)")
	fmt.Println("  - Can represent integers up to 2^53 - 1 safely")
	fmt.Println("  - BigInt for arbitrary precision")
	fmt.Println("  - Automatic type coercion")
	fmt.Println("  - let x = 42; (can be changed to any type)")

	fmt.Println("\nGo:")
	fmt.Println("  - Multiple integer types with specific sizes")
	fmt.Println("  - Explicit type declaration")
	fmt.Println("  - No automatic type conversion")
	fmt.Println("  - Compile-time type checking")
	fmt.Println("  - var x int = 42; (always an int)")
}
