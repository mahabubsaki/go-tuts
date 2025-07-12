package main

import (
	"fmt"
	"math/bits"
)

// This file covers all bitwise operators in Go
// Bitwise operators work on individual bits of integers

func main() {
	fmt.Println("=== GO BITWISE OPERATORS - COMPLETE GUIDE ===")

	demonstrateBasicBitwiseOperators()
	demonstrateBitwiseShiftOperators()
	demonstrateBitwiseOperationsWithDifferentTypes()
	demonstrateBitwiseTruthTables()
	demonstrateBitwiseManipulationTechniques()
	demonstratePracticalBitwiseExamples()
	demonstrateBitwiseOperatorPrecedence()
	demonstrateBitwiseBestPractices()
}

func demonstrateBasicBitwiseOperators() {
	fmt.Println("\n--- BASIC BITWISE OPERATORS ---")

	// Basic bitwise operators: &, |, ^, &^, ~(using ^)
	var a uint8 = 12 // 00001100 in binary
	var b uint8 = 10 // 00001010 in binary

	fmt.Printf("a = %d (binary: %08b)\n", a, a)
	fmt.Printf("b = %d (binary: %08b)\n", b, b)

	// Bitwise AND (&)
	var andResult uint8 = a & b
	fmt.Printf("a & b = %d (binary: %08b)\n", andResult, andResult)

	// Bitwise OR (|)
	var orResult uint8 = a | b
	fmt.Printf("a | b = %d (binary: %08b)\n", orResult, orResult)

	// Bitwise XOR (^)
	var xorResult uint8 = a ^ b
	fmt.Printf("a ^ b = %d (binary: %08b)\n", xorResult, xorResult)

	// Bitwise NOT (^) - unary operator
	var notA uint8 = ^a
	var notB uint8 = ^b
	fmt.Printf("^a = %d (binary: %08b)\n", notA, notA)
	fmt.Printf("^b = %d (binary: %08b)\n", notB, notB)

	// Bitwise AND NOT (&^) - bit clear
	var andNotResult uint8 = a &^ b
	fmt.Printf("a &^ b = %d (binary: %08b)\n", andNotResult, andNotResult)

	// Demonstrating bit positions
	fmt.Printf("\nBit position analysis:\n")
	fmt.Printf("Position: 76543210\n")
	fmt.Printf("a:        %08b\n", a)
	fmt.Printf("b:        %08b\n", b)
	fmt.Printf("a & b:    %08b\n", a&b)
	fmt.Printf("a | b:    %08b\n", a|b)
	fmt.Printf("a ^ b:    %08b\n", a^b)
	fmt.Printf("a &^ b:   %08b\n", a&^b)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Same bitwise operators (&, |, ^, ~)
	// JavaScript: Numbers are 64-bit floating-point, bitwise ops use 32-bit
	// JavaScript: ~ is bitwise NOT, Go uses ^ for bitwise NOT
	// Go: &^ is bit clear operator (not in JavaScript)
	// Go: More explicit integer types for bitwise operations
}

func demonstrateBitwiseShiftOperators() {
	fmt.Println("\n--- BITWISE SHIFT OPERATORS ---")

	// Left shift (<<) and right shift (>>)
	var value uint8 = 5 // 00000101 in binary

	fmt.Printf("Original value: %d (binary: %08b)\n", value, value)

	// Left shift - multiply by 2^n
	for i := 1; i <= 4; i++ {
		var shifted uint8 = value << i
		fmt.Printf("value << %d = %d (binary: %08b) [multiply by %d]\n",
			i, shifted, shifted, 1<<i)
	}

	// Right shift - divide by 2^n
	var largeValue uint8 = 80 // 01010000 in binary
	fmt.Printf("\nOriginal value: %d (binary: %08b)\n", largeValue, largeValue)

	for i := 1; i <= 4; i++ {
		var shifted uint8 = largeValue >> i
		fmt.Printf("value >> %d = %d (binary: %08b) [divide by %d]\n",
			i, shifted, shifted, 1<<i)
	}

	// Signed vs unsigned right shift
	var signedValue int8 = -8 // 11111000 in binary (two's complement)
	fmt.Printf("\nSigned right shift:\n")
	fmt.Printf("Original: %d (binary: %08b)\n", signedValue, uint8(signedValue))

	for i := 1; i <= 3; i++ {
		var shifted int8 = signedValue >> i
		fmt.Printf("value >> %d = %d (binary: %08b) [arithmetic shift]\n",
			i, shifted, uint8(shifted))
	}

	// Unsigned right shift
	var unsignedValue uint8 = 248 // 11111000 in binary
	fmt.Printf("\nUnsigned right shift:\n")
	fmt.Printf("Original: %d (binary: %08b)\n", unsignedValue, unsignedValue)

	for i := 1; i <= 3; i++ {
		var shifted uint8 = unsignedValue >> i
		fmt.Printf("value >> %d = %d (binary: %08b) [logical shift]\n",
			i, shifted, shifted)
	}

	// Shift with different data types
	var int16Val int16 = 1000
	var int32Val int32 = 100000
	var int64Val int64 = 1000000

	fmt.Printf("\nShift with different types:\n")
	fmt.Printf("int16: %d << 2 = %d\n", int16Val, int16Val<<2)
	fmt.Printf("int32: %d << 3 = %d\n", int32Val, int32Val<<3)
	fmt.Printf("int64: %d << 4 = %d\n", int64Val, int64Val<<4)

	// Overflow with shifts
	var overflowValue uint8 = 200
	fmt.Printf("\nOverflow with shifts:\n")
	fmt.Printf("Original: %d (binary: %08b)\n", overflowValue, overflowValue)

	var overflowResult uint8 = overflowValue << 1
	fmt.Printf("After << 1: %d (binary: %08b) [overflow!]\n", overflowResult, overflowResult)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Same shift operators (<<, >>)
	// JavaScript: >>> for unsigned right shift
	// Go: Right shift behavior depends on signed/unsigned type
	// Go: No separate unsigned right shift operator
}

func demonstrateBitwiseOperationsWithDifferentTypes() {
	fmt.Println("\n--- BITWISE OPERATIONS WITH DIFFERENT TYPES ---")

	// Different integer types
	var int8Val int8 = 15        // 00001111
	var int16Val int16 = 255     // 0000000011111111
	var int32Val int32 = 65535   // 00000000000000001111111111111111
	var int64Val int64 = 1048575 // many zeros and ones

	fmt.Printf("int8:  %d (binary: %08b)\n", int8Val, int8Val)
	fmt.Printf("int16: %d (binary: %016b)\n", int16Val, int16Val)
	fmt.Printf("int32: %d (binary: %032b)\n", int32Val, int32Val)
	fmt.Printf("int64: %d (binary: %064b)\n", int64Val, int64Val)

	// Type conversion needed for bitwise operations
	fmt.Printf("\nBitwise operations with type conversion:\n")
	fmt.Printf("int8 & int16 (converted): %d\n", int8Val&int8(int16Val))
	fmt.Printf("int16 & int32 (converted): %d\n", int16Val&int16(int32Val))

	// Unsigned types
	var uint8Val uint8 = 255
	var uint16Val uint16 = 65535
	var uint32Val uint32 = 4294967295
	var uint64Val uint64 = 18446744073709551615

	fmt.Printf("\nUnsigned types:\n")
	fmt.Printf("uint8:  %d (binary: %08b)\n", uint8Val, uint8Val)
	fmt.Printf("uint16: %d (binary: %016b)\n", uint16Val, uint16Val)
	fmt.Printf("uint32: %d (binary: %032b)\n", uint32Val, uint32Val)
	fmt.Printf("uint64: %d (max value)\n", uint64Val)

	// Bitwise operations on maximum values
	fmt.Printf("\nBitwise NOT on maximum values:\n")
	fmt.Printf("^uint8(255) = %d\n", ^uint8Val)
	fmt.Printf("^uint16(65535) = %d\n", ^uint16Val)

	// Mixing signed and unsigned (requires conversion)
	var signed int8 = -1     // 11111111 in two's complement
	var unsigned uint8 = 255 // 11111111

	fmt.Printf("\nSigned vs unsigned:\n")
	fmt.Printf("signed int8(-1): %08b\n", signed)
	fmt.Printf("unsigned uint8(255): %08b\n", unsigned)
	fmt.Printf("Same bit pattern: %t\n", uint8(signed) == unsigned)

	// Byte operations
	var byteVal byte = 170 // byte is alias for uint8
	fmt.Printf("\nByte operations:\n")
	fmt.Printf("byte value: %d (binary: %08b)\n", byteVal, byteVal)
	fmt.Printf("byte << 1: %d (binary: %08b)\n", byteVal<<1, byteVal<<1)
	fmt.Printf("byte >> 1: %d (binary: %08b)\n", byteVal>>1, byteVal>>1)

	// Rune operations
	var runeVal rune = 'A' // rune is alias for int32
	fmt.Printf("\nRune operations:\n")
	fmt.Printf("rune 'A': %d (binary: %032b)\n", runeVal, runeVal)
	fmt.Printf("rune | 32: %d ('%c') [convert to lowercase]\n", runeVal|32, runeVal|32)
}

func demonstrateBitwiseTruthTables() {
	fmt.Println("\n--- BITWISE TRUTH TABLES ---")

	// AND truth table
	fmt.Printf("AND Truth Table (&):\n")
	fmt.Printf("A | B | A & B\n")
	fmt.Printf("--|---|------\n")
	fmt.Printf("0 | 0 |  %d\n", 0&0)
	fmt.Printf("0 | 1 |  %d\n", 0&1)
	fmt.Printf("1 | 0 |  %d\n", 1&0)
	fmt.Printf("1 | 1 |  %d\n", 1&1)

	// OR truth table
	fmt.Printf("\nOR Truth Table (|):\n")
	fmt.Printf("A | B | A | B\n")
	fmt.Printf("--|---|------\n")
	fmt.Printf("0 | 0 |  %d\n", 0|0)
	fmt.Printf("0 | 1 |  %d\n", 0|1)
	fmt.Printf("1 | 0 |  %d\n", 1|0)
	fmt.Printf("1 | 1 |  %d\n", 1|1)

	// XOR truth table
	fmt.Printf("\nXOR Truth Table (^):\n")
	fmt.Printf("A | B | A ^ B\n")
	fmt.Printf("--|---|------\n")
	fmt.Printf("0 | 0 |  %d\n", 0^0)
	fmt.Printf("0 | 1 |  %d\n", 0^1)
	fmt.Printf("1 | 0 |  %d\n", 1^0)
	fmt.Printf("1 | 1 |  %d\n", 1^1)

	// NOT truth table
	fmt.Printf("\nNOT Truth Table (^):\n")
	fmt.Printf("A | ^A (8-bit)\n")
	fmt.Printf("--|----------\n")
	fmt.Printf("0 | %d\n", ^uint8(0))
	fmt.Printf("1 | %d\n", ^uint8(1))

	// AND NOT truth table
	fmt.Printf("\nAND NOT Truth Table (&^):\n")
	fmt.Printf("A | B | A &^ B\n")
	fmt.Printf("--|---|-------\n")
	fmt.Printf("0 | 0 |   %d\n", 0&^0)
	fmt.Printf("0 | 1 |   %d\n", 0&^1)
	fmt.Printf("1 | 0 |   %d\n", 1&^0)
	fmt.Printf("1 | 1 |   %d\n", 1&^1)

	// Demonstrating with actual byte values
	fmt.Printf("\nByte-level operations:\n")
	var a uint8 = 0xAA // 10101010
	var b uint8 = 0xCC // 11001100

	fmt.Printf("a = 0x%02X (binary: %08b)\n", a, a)
	fmt.Printf("b = 0x%02X (binary: %08b)\n", b, b)
	fmt.Printf("a & b = 0x%02X (binary: %08b)\n", a&b, a&b)
	fmt.Printf("a | b = 0x%02X (binary: %08b)\n", a|b, a|b)
	fmt.Printf("a ^ b = 0x%02X (binary: %08b)\n", a^b, a^b)
	fmt.Printf("a &^ b = 0x%02X (binary: %08b)\n", a&^b, a&^b)
}

func demonstrateBitwiseManipulationTechniques() {
	fmt.Println("\n--- BITWISE MANIPULATION TECHNIQUES ---")

	// Setting bits
	var value uint8 = 0 // 00000000
	fmt.Printf("Original value: %d (binary: %08b)\n", value, value)

	// Set bit at position 2
	var setBit uint8 = value | (1 << 2)
	fmt.Printf("Set bit 2: %d (binary: %08b)\n", setBit, setBit)

	// Set multiple bits
	var setMultiple uint8 = value | (1<<1 | 1<<3 | 1<<5)
	fmt.Printf("Set bits 1,3,5: %d (binary: %08b)\n", setMultiple, setMultiple)

	// Clearing bits
	value = 255 // 11111111
	fmt.Printf("\nClearing bits from: %d (binary: %08b)\n", value, value)

	// Clear bit at position 2
	var clearBit uint8 = value &^ (1 << 2)
	fmt.Printf("Clear bit 2: %d (binary: %08b)\n", clearBit, clearBit)

	// Clear multiple bits
	var clearMultiple uint8 = value &^ (1<<1 | 1<<3 | 1<<5)
	fmt.Printf("Clear bits 1,3,5: %d (binary: %08b)\n", clearMultiple, clearMultiple)

	// Toggling bits
	value = 170 // 10101010
	fmt.Printf("\nToggling bits in: %d (binary: %08b)\n", value, value)

	// Toggle bit at position 0
	var toggleBit uint8 = value ^ (1 << 0)
	fmt.Printf("Toggle bit 0: %d (binary: %08b)\n", toggleBit, toggleBit)

	// Toggle multiple bits
	var toggleMultiple uint8 = value ^ (1<<0 | 1<<2 | 1<<4)
	fmt.Printf("Toggle bits 0,2,4: %d (binary: %08b)\n", toggleMultiple, toggleMultiple)

	// Checking bits
	value = 42 // 00101010
	fmt.Printf("\nChecking bits in: %d (binary: %08b)\n", value, value)

	for i := 0; i < 8; i++ {
		var isSet bool = (value & (1 << i)) != 0
		fmt.Printf("Bit %d is set: %t\n", i, isSet)
	}

	// Counting set bits
	fmt.Printf("\nCounting set bits:\n")
	var testValues []uint8 = []uint8{0, 1, 3, 7, 15, 31, 63, 127, 255}

	for _, val := range testValues {
		var count int = bits.OnesCount8(val)
		fmt.Printf("Value %d (binary: %08b) has %d set bits\n", val, val, count)
	}

	// Finding the rightmost set bit
	fmt.Printf("\nFinding rightmost set bit:\n")
	for _, val := range []uint8{8, 12, 20, 24} {
		var rightmostBit uint8 = val & (^val + 1) // or val & (-val) for signed
		fmt.Printf("Value %d (binary: %08b), rightmost bit: %d (binary: %08b)\n",
			val, val, rightmostBit, rightmostBit)
	}

	// Checking if power of 2
	fmt.Printf("\nChecking if power of 2:\n")
	for _, val := range []uint8{1, 2, 3, 4, 5, 8, 16, 32} {
		var isPowerOf2 bool = val != 0 && (val&(val-1)) == 0
		fmt.Printf("Value %d is power of 2: %t\n", val, isPowerOf2)
	}

	// Swapping values using XOR
	fmt.Printf("\nSwapping values using XOR:\n")
	var x, y uint8 = 25, 40
	fmt.Printf("Before swap: x=%d, y=%d\n", x, y)

	x = x ^ y
	y = x ^ y
	x = x ^ y

	fmt.Printf("After swap: x=%d, y=%d\n", x, y)
}

func demonstratePracticalBitwiseExamples() {
	fmt.Println("\n--- PRACTICAL BITWISE EXAMPLES ---")

	// Flags and permissions
	const (
		READ    uint8 = 1 << 0 // 00000001
		WRITE   uint8 = 1 << 1 // 00000010
		EXECUTE uint8 = 1 << 2 // 00000100
		DELETE  uint8 = 1 << 3 // 00001000
	)

	fmt.Printf("Permission flags:\n")
	fmt.Printf("READ:    %d (binary: %08b)\n", READ, READ)
	fmt.Printf("WRITE:   %d (binary: %08b)\n", WRITE, WRITE)
	fmt.Printf("EXECUTE: %d (binary: %08b)\n", EXECUTE, EXECUTE)
	fmt.Printf("DELETE:  %d (binary: %08b)\n", DELETE, DELETE)

	// Setting permissions
	var permissions uint8 = READ | WRITE | EXECUTE
	fmt.Printf("\nUser permissions: %d (binary: %08b)\n", permissions, permissions)

	// Checking permissions
	fmt.Printf("Has READ permission: %t\n", (permissions&READ) != 0)
	fmt.Printf("Has WRITE permission: %t\n", (permissions&WRITE) != 0)
	fmt.Printf("Has DELETE permission: %t\n", (permissions&DELETE) != 0)

	// Adding permission
	permissions = permissions | DELETE
	fmt.Printf("After adding DELETE: %d (binary: %08b)\n", permissions, permissions)

	// Removing permission
	permissions = permissions &^ WRITE
	fmt.Printf("After removing WRITE: %d (binary: %08b)\n", permissions, permissions)

	// RGB color manipulation
	fmt.Printf("\nRGB color manipulation:\n")
	var color uint32 = 0xFF6A33 // Orange color

	// Extract RGB components
	var red uint8 = uint8((color >> 16) & 0xFF)
	var green uint8 = uint8((color >> 8) & 0xFF)
	var blue uint8 = uint8(color & 0xFF)

	fmt.Printf("Color: 0x%06X\n", color)
	fmt.Printf("Red: %d, Green: %d, Blue: %d\n", red, green, blue)

	// Modify components
	var newRed uint8 = 128
	var newColor uint32 = (color & 0x00FFFF) | (uint32(newRed) << 16)

	fmt.Printf("New color: 0x%06X\n", newColor)

	// Hash table sizing (power of 2)
	fmt.Printf("\nHash table sizing:\n")
	for size := 1; size <= 64; size *= 2 {
		var mask uint32 = uint32(size - 1)
		var hash uint32 = 123456789
		var bucket uint32 = hash & mask

		fmt.Printf("Size: %d, Mask: 0x%08X, Hash: %d, Bucket: %d\n",
			size, mask, hash, bucket)
	}

	// Bit manipulation for optimization
	fmt.Printf("\nOptimization examples:\n")

	// Fast multiplication/division by powers of 2
	var num uint32 = 100
	fmt.Printf("Original: %d\n", num)
	fmt.Printf("Multiply by 4: %d (shift left 2)\n", num<<2)
	fmt.Printf("Multiply by 8: %d (shift left 3)\n", num<<3)
	fmt.Printf("Divide by 4: %d (shift right 2)\n", num>>2)
	fmt.Printf("Divide by 8: %d (shift right 3)\n", num>>3)

	// Check if number is even/odd
	var numbers []uint32 = []uint32{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("\nEven/Odd check using bitwise AND:\n")
	for _, n := range numbers {
		var isEven bool = (n & 1) == 0
		fmt.Printf("Number %d is even: %t\n", n, isEven)
	}

	// Bit reversal
	fmt.Printf("\nBit reversal:\n")
	var original uint8 = 0b10110010
	var reversed uint8 = reverseBits(original)
	fmt.Printf("Original: %08b, Reversed: %08b\n", original, reversed)
}

func demonstrateBitwiseOperatorPrecedence() {
	fmt.Println("\n--- BITWISE OPERATOR PRECEDENCE ---")

	// Precedence order: ^ (unary), <<, >>, &, ^, |
	var a, b, c uint8 = 12, 8, 4

	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	// Examples of precedence
	var result1 uint8 = a | b&c
	var result2 uint8 = (a | b) & c
	var result3 uint8 = a | (b & c)

	fmt.Printf("a | b & c: %d (equivalent to a | (b & c))\n", result1)
	fmt.Printf("(a | b) & c: %d\n", result2)
	fmt.Printf("a | (b & c): %d\n", result3)

	// Shift operators precedence
	var shift1 uint8 = a + b<<1
	var shift2 uint8 = (a + b) << 1
	var shift3 uint8 = a + (b << 1)

	fmt.Printf("a + b << 1: %d (equivalent to a + (b << 1))\n", shift1)
	fmt.Printf("(a + b) << 1: %d\n", shift2)
	fmt.Printf("a + (b << 1): %d\n", shift3)

	// XOR precedence
	var xor1 uint8 = a ^ b&c
	var xor2 uint8 = (a ^ b) & c
	var xor3 uint8 = a ^ (b & c)

	fmt.Printf("a ^ b & c: %d (equivalent to a ^ (b & c))\n", xor1)
	fmt.Printf("(a ^ b) & c: %d\n", xor2)
	fmt.Printf("a ^ (b & c): %d\n", xor3)

	// Best practice: use parentheses for clarity
	fmt.Printf("\nRecommended: Use parentheses for clarity:\n")
	fmt.Printf("Original: a | b & c\n")
	fmt.Printf("Clear: a | (b & c)\n")
	fmt.Printf("Original: a + b << 1\n")
	fmt.Printf("Clear: a + (b << 1)\n")
}

func demonstrateBitwiseBestPractices() {
	fmt.Println("\n--- BITWISE BEST PRACTICES ---")

	// Use constants for bit flags
	const (
		FLAG_A uint8 = 1 << 0
		FLAG_B uint8 = 1 << 1
		FLAG_C uint8 = 1 << 2
		FLAG_D uint8 = 1 << 3
	)

	var flags uint8 = FLAG_A | FLAG_C

	fmt.Printf("Using named constants for flags:\n")
	fmt.Printf("flags = FLAG_A | FLAG_C = %d (binary: %08b)\n", flags, flags)

	// Use appropriate integer types
	var smallFlags uint8 = 0xFF        // 8 bits
	var mediumFlags uint16 = 0xFFFF    // 16 bits
	var largeFlags uint32 = 0xFFFFFFFF // 32 bits

	fmt.Printf("\nAppropriate types for different flag counts:\n")
	fmt.Printf("uint8 for up to 8 flags: %d\n", smallFlags)
	fmt.Printf("uint16 for up to 16 flags: %d\n", mediumFlags)
	fmt.Printf("uint32 for up to 32 flags: %d\n", largeFlags)

	// Use helper functions for complex bit operations
	fmt.Printf("\nUsing helper functions:\n")
	var value uint8 = 0b10101010

	fmt.Printf("Original: %08b\n", value)
	fmt.Printf("Set bit 0: %08b\n", setBit(value, 0))
	fmt.Printf("Clear bit 1: %08b\n", clearBit(value, 1))
	fmt.Printf("Toggle bit 2: %08b\n", toggleBit(value, 2))
	fmt.Printf("Is bit 3 set: %t\n", isBitSet(value, 3))

	// Document bit layouts
	fmt.Printf("\nDocumented bit layout:\n")
	fmt.Printf("Register layout (8-bit):\n")
	fmt.Printf("Bit 7: Reserved\n")
	fmt.Printf("Bit 6: Error flag\n")
	fmt.Printf("Bit 5: Ready flag\n")
	fmt.Printf("Bit 4: Enable flag\n")
	fmt.Printf("Bits 3-0: Counter value\n")

	var register uint8 = 0b00110101
	fmt.Printf("Register value: %08b\n", register)
	fmt.Printf("Error flag: %t\n", isBitSet(register, 6))
	fmt.Printf("Ready flag: %t\n", isBitSet(register, 5))
	fmt.Printf("Enable flag: %t\n", isBitSet(register, 4))
	fmt.Printf("Counter value: %d\n", register&0x0F)

	// Use bit manipulation for performance
	fmt.Printf("\nPerformance optimizations:\n")
	var numbers []uint32 = []uint32{10, 20, 30, 40, 50}

	// Fast even/odd check
	fmt.Printf("Even/odd check (using & 1):\n")
	for _, n := range numbers {
		if n&1 == 0 {
			fmt.Printf("%d is even\n", n)
		} else {
			fmt.Printf("%d is odd\n", n)
		}
	}

	// Fast power of 2 multiplication
	fmt.Printf("Power of 2 multiplication:\n")
	for _, n := range numbers {
		fmt.Printf("%d * 4 = %d (using << 2)\n", n, n<<2)
	}

	// BEST PRACTICES SUMMARY:
	fmt.Println("\nBest Practices Summary:")
	fmt.Println("1. Use named constants for bit flags")
	fmt.Println("2. Choose appropriate integer types for bit count")
	fmt.Println("3. Use helper functions for complex bit operations")
	fmt.Println("4. Document bit layouts clearly")
	fmt.Println("5. Use parentheses to clarify precedence")
	fmt.Println("6. Leverage bit manipulation for performance")
	fmt.Println("7. Be aware of signed vs unsigned behavior")
	fmt.Println("8. Use bit manipulation for flags and permissions")
	fmt.Println("9. Consider endianness for multi-byte operations")
	fmt.Println("10. Test bit operations thoroughly")
}

// Helper functions
func setBit(value uint8, position uint) uint8 {
	return value | (1 << position)
}

func clearBit(value uint8, position uint) uint8 {
	return value &^ (1 << position)
}

func toggleBit(value uint8, position uint) uint8 {
	return value ^ (1 << position)
}

func isBitSet(value uint8, position uint) bool {
	return (value & (1 << position)) != 0
}

func reverseBits(value uint8) uint8 {
	var result uint8 = 0
	for i := 0; i < 8; i++ {
		result = (result << 1) | (value & 1)
		value >>= 1
	}
	return result
}
