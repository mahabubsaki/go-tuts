package main

import (
	"fmt"
	"unsafe"
)

// This file covers ALL Go data types in detail
// Go is statically typed - every variable has a specific type
// JavaScript is dynamically typed - variables can change type

func main2() {
	fmt.Println("=== GO DATA TYPES OVERVIEW ===")

	// Go has several categories of data types:
	// 1. Basic types (bool, string, numeric)
	// 2. Aggregate types (arrays, structs)
	// 3. Reference types (pointers, slices, maps, channels, functions)
	// 4. Interface types

	demonstrateBasicTypes()
	demonstrateTypeProperties()
}

func demonstrateBasicTypes() {
	fmt.Println("\n--- Basic Types ---")

	// BOOLEAN TYPE
	var isTrue bool = true
	var isFalse bool = false
	var defaultBool bool // zero value is false

	fmt.Printf("bool: %t, %t, %t\n", isTrue, isFalse, defaultBool)

	// STRING TYPE
	var message string = "Hello, Go!"
	var emptyString string // zero value is ""

	fmt.Printf("string: %s, empty: '%s'\n", message, emptyString)

	// NUMERIC TYPES - we'll cover these in detail in separate files
	var integer int = 42
	var floatingPoint float64 = 3.14159

	fmt.Printf("int: %d, float64: %f\n", integer, floatingPoint)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: boolean, string, number, undefined, null, object, symbol, bigint
	// Go: bool, string, int/uint variants, float variants, complex variants
	//
	// JavaScript has dynamic typing:
	// let x = 42;      // number
	// x = "hello";     // now string
	// x = true;        // now boolean
	//
	// Go has static typing:
	// var x int = 42;  // always int
	// x = "hello";     // ERROR: cannot assign string to int
}

func demonstrateTypeProperties() {
	fmt.Println("\n--- Type Properties ---")

	// TYPE SIZE
	// Go types have specific sizes (unlike JavaScript's number type)
	fmt.Printf("Size of int: %d bytes\n", unsafe.Sizeof(int(0)))
	fmt.Printf("Size of float64: %d bytes\n", unsafe.Sizeof(float64(0)))
	fmt.Printf("Size of bool: %d bytes\n", unsafe.Sizeof(bool(true)))
	fmt.Printf("Size of string: %d bytes\n", unsafe.Sizeof(string("")))

	// ZERO VALUES
	// Every type in Go has a zero value (default value)
	// JavaScript: undefined for uninitialized variables
	// Go: specific zero value for each type

	var zeroBool bool     // false
	var zeroInt int       // 0
	var zeroFloat float64 // 0.0
	var zeroString string // ""

	fmt.Printf("Zero values - bool: %t, int: %d, float64: %f, string: '%s'\n",
		zeroBool, zeroInt, zeroFloat, zeroString)

	// TYPE SAFETY
	// Go prevents type-related bugs at compile time
	// JavaScript allows implicit type conversions (coercion)

	fmt.Println("\nGo prevents these JavaScript behaviors:")
	fmt.Println("JavaScript: '5' + 3 = '53' (string concatenation)")
	fmt.Println("JavaScript: '5' - 3 = 2 (numeric subtraction)")
	fmt.Println("JavaScript: true + 1 = 2 (boolean to number)")
	fmt.Println("Go: All of these would be compilation errors")
}

// TYPE CATEGORIES OVERVIEW:
//
// 1. BOOLEAN
//    - bool
//
// 2. STRING
//    - string
//
// 3. NUMERIC
//    a) Integers:
//       - Signed: int8, int16, int32, int64, int
//       - Unsigned: uint8, uint16, uint32, uint64, uint
//       - Aliases: byte (uint8), rune (int32)
//
//    b) Floating-point:
//       - float32, float64
//
//    c) Complex:
//       - complex64, complex128
//
// 4. DERIVED/COMPOSITE
//    - Arrays: [n]T
//    - Slices: []T
//    - Maps: map[K]V
//    - Structs: struct { ... }
//    - Pointers: *T
//    - Functions: func(...) (...)
//    - Channels: chan T
//    - Interfaces: interface { ... }
//
// COMPARISON WITH JAVASCRIPT:
// JavaScript has fewer primitive types but more flexible usage
// Go has more specific types but stricter rules
// This trade-off gives Go better performance and safety
// JavaScript gives more flexibility but can lead to runtime errors
