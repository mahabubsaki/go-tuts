package main

import (
	"fmt"
	"math"
	"strings"
)

// This file covers all comparison operators in Go
// Comparison operators return boolean values (true/false)

func main() {
	fmt.Println("=== GO COMPARISON OPERATORS - COMPLETE GUIDE ===")

	demonstrateBasicComparison()
	demonstrateNumericComparison()
	demonstrateStringComparison()
	demonstrateFloatingPointComparison()
	demonstrateBooleanComparison()
	demonstrateTypeSpecificComparison()
	demonstrateComparisonChaining()
	demonstrateComparisonBestPractices()
}

func demonstrateBasicComparison() {
	fmt.Println("\n--- BASIC COMPARISON OPERATORS ---")

	// Basic comparison operators: ==, !=, <, >, <=, >=
	var a int = 10
	var b int = 20
	var c int = 10

	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	// Equality operators
	fmt.Printf("a == b: %t\n", a == b)
	fmt.Printf("a == c: %t\n", a == c)
	fmt.Printf("a != b: %t\n", a != b)
	fmt.Printf("a != c: %t\n", a != c)

	// Relational operators
	fmt.Printf("a < b: %t\n", a < b)
	fmt.Printf("a > b: %t\n", a > b)
	fmt.Printf("a <= b: %t\n", a <= b)
	fmt.Printf("a >= b: %t\n", a >= b)
	fmt.Printf("a <= c: %t\n", a <= c)
	fmt.Printf("a >= c: %t\n", a >= c)

	// Comparison precedence
	var result bool = a+5 > b-5
	fmt.Printf("a + 5 > b - 5: %d > %d = %t\n", a+5, b-5, result)

	// Comparison with literals
	fmt.Printf("a == 10: %t\n", a == 10)
	fmt.Printf("b > 15: %t\n", b > 15)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: == (loose equality) vs === (strict equality)
	// JavaScript: Type coercion with == (1 == "1" is true)
	// Go: Only == (always strict, no type coercion)
	// Go: Compile-time type checking prevents many errors
}

func demonstrateNumericComparison() {
	fmt.Println("\n--- NUMERIC COMPARISON ---")

	// Integer comparisons
	var int8Val int8 = 100
	var int16Val int16 = 100
	var int32Val int32 = 100
	var int64Val int64 = 100

	fmt.Printf("Different integer types (all value 100):\n")
	fmt.Printf("int8: %d, int16: %d, int32: %d, int64: %d\n",
		int8Val, int16Val, int32Val, int64Val)

	// Type conversion needed for comparison
	// fmt.Printf("int8 == int16: %t\n", int8Val == int16Val)  // Error!
	fmt.Printf("int8 == int16 (converted): %t\n", int(int8Val) == int(int16Val))

	// Signed vs unsigned comparison
	var signedInt int = -5
	var unsignedInt uint = 5

	fmt.Printf("Signed: %d, Unsigned: %d\n", signedInt, unsignedInt)
	// fmt.Printf("signed == unsigned: %t\n", signedInt == unsignedInt)  // Error!

	// Safe comparison after checking sign
	if signedInt >= 0 {
		fmt.Printf("Safe comparison: %d == %d: %t\n", signedInt, unsignedInt,
			uint(signedInt) == unsignedInt)
	} else {
		fmt.Printf("Cannot compare negative signed with unsigned\n")
	}

	// Floating-point comparisons
	var float32Val float32 = 3.14
	var float64Val float64 = 3.14

	fmt.Printf("float32: %f, float64: %f\n", float32Val, float64Val)
	fmt.Printf("float32 == float64 (converted): %t\n", float64(float32Val) == float64Val)

	// Range comparisons
	var value int = 15
	var min int = 10
	var max int = 20

	fmt.Printf("Range check: %d <= %d <= %d: %t\n", min, value, max,
		min <= value && value <= max)

	// Zero comparisons
	var zero int = 0
	var positive int = 5
	var negative int = -5

	fmt.Printf("Zero comparisons:\n")
	fmt.Printf("%d == 0: %t\n", zero, zero == 0)
	fmt.Printf("%d > 0: %t\n", positive, positive > 0)
	fmt.Printf("%d < 0: %t\n", negative, negative < 0)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Automatic type conversion (1 == "1")
	// JavaScript: Can compare different numeric types
	// Go: Explicit type conversion required
	// Go: Type safety prevents many bugs
}

func demonstrateStringComparison() {
	fmt.Println("\n--- STRING COMPARISON ---")

	// String equality
	var str1 string = "Hello"
	var str2 string = "Hello"
	var str3 string = "hello"
	var str4 string = "World"

	fmt.Printf("str1 = '%s', str2 = '%s', str3 = '%s', str4 = '%s'\n",
		str1, str2, str3, str4)

	fmt.Printf("str1 == str2: %t\n", str1 == str2)
	fmt.Printf("str1 == str3: %t\n", str1 == str3) // Case-sensitive
	fmt.Printf("str1 != str4: %t\n", str1 != str4)

	// Lexicographic comparison
	fmt.Printf("str1 < str4: %t\n", str1 < str4)
	fmt.Printf("str1 > str3: %t\n", str1 > str3) // Capital letters come before lowercase

	// Case-insensitive comparison
	fmt.Printf("Case-insensitive comparison using strings.EqualFold:\n")
	fmt.Printf("EqualFold('%s', '%s'): %t\n", str1, str3, strings.EqualFold(str1, str3))

	// String comparison with different cases
	var upperCase string = "HELLO"
	var lowerCase string = "hello"
	var mixedCase string = "Hello"

	fmt.Printf("String case comparisons:\n")
	fmt.Printf("'%s' == '%s': %t\n", upperCase, lowerCase, upperCase == lowerCase)
	fmt.Printf("'%s' == '%s': %t\n", upperCase, mixedCase, upperCase == mixedCase)
	fmt.Printf("'%s' == '%s': %t\n", lowerCase, mixedCase, lowerCase == mixedCase)

	// UTF-8 and Unicode comparison
	var asciiStr string = "Hello"
	var unicodeStr string = "Hello"
	var emojiStr string = "Hello👋"

	fmt.Printf("Unicode comparisons:\n")
	fmt.Printf("'%s' == '%s': %t\n", asciiStr, unicodeStr, asciiStr == unicodeStr)
	fmt.Printf("'%s' == '%s': %t\n", asciiStr, emojiStr, asciiStr == emojiStr)

	// Empty string comparisons
	var emptyStr string = ""
	var spaceStr string = " "

	fmt.Printf("Empty string comparisons:\n")
	fmt.Printf("'' == '': %t\n", emptyStr == "")
	fmt.Printf("'' == ' ': %t\n", emptyStr == spaceStr)
	fmt.Printf("len('') == 0: %t\n", len(emptyStr) == 0)

	// String length comparison
	var shortStr string = "Hi"
	var longStr string = "Hello World"

	fmt.Printf("Length comparisons:\n")
	fmt.Printf("len('%s') = %d\n", shortStr, len(shortStr))
	fmt.Printf("len('%s') = %d\n", longStr, len(longStr))
	fmt.Printf("len('%s') < len('%s'): %t\n", shortStr, longStr, len(shortStr) < len(longStr))

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: String comparison works similarly
	// JavaScript: localeCompare() for locale-aware comparison
	// Go: Byte-by-byte comparison (UTF-8)
	// Go: strings package for advanced string operations
}

func demonstrateFloatingPointComparison() {
	fmt.Println("\n--- FLOATING-POINT COMPARISON ---")

	// Floating-point precision issues
	var a float64 = 0.1 + 0.2
	var b float64 = 0.3

	fmt.Printf("a = 0.1 + 0.2 = %.17f\n", a)
	fmt.Printf("b = 0.3 = %.17f\n", b)
	fmt.Printf("a == b: %t\n", a == b)

	// Proper floating-point comparison
	var epsilon float64 = 1e-9
	var isEqual bool = math.Abs(a-b) < epsilon

	fmt.Printf("Using epsilon (%.e): %t\n", epsilon, isEqual)

	// Different epsilon values
	var largeEpsilon float64 = 1e-6
	var smallEpsilon float64 = 1e-15

	fmt.Printf("Large epsilon (%.e): %t\n", largeEpsilon, math.Abs(a-b) < largeEpsilon)
	fmt.Printf("Small epsilon (%.e): %t\n", smallEpsilon, math.Abs(a-b) < smallEpsilon)

	// Special floating-point values
	var positiveInf float64 = math.Inf(1)
	var negativeInf float64 = math.Inf(-1)
	var notANumber float64 = math.NaN()

	fmt.Printf("Special values:\n")
	fmt.Printf("Positive infinity: %f\n", positiveInf)
	fmt.Printf("Negative infinity: %f\n", negativeInf)
	fmt.Printf("NaN: %f\n", notANumber)

	// Comparing special values
	fmt.Printf("Infinity comparisons:\n")
	fmt.Printf("+Inf == +Inf: %t\n", positiveInf == positiveInf)
	fmt.Printf("+Inf > 1000000: %t\n", positiveInf > 1000000)
	fmt.Printf("-Inf < -1000000: %t\n", negativeInf < -1000000)

	// NaN comparisons (NaN != NaN)
	fmt.Printf("NaN comparisons:\n")
	fmt.Printf("NaN == NaN: %t\n", notANumber == notANumber)
	fmt.Printf("NaN != NaN: %t\n", notANumber != notANumber)
	fmt.Printf("IsNaN(NaN): %t\n", math.IsNaN(notANumber))

	// float32 vs float64 precision
	var float32Val float32 = 1.0000001
	var float64Val float64 = 1.0000001

	fmt.Printf("Precision differences:\n")
	fmt.Printf("float32: %.10f\n", float32Val)
	fmt.Printf("float64: %.10f\n", float64Val)
	fmt.Printf("Equal after conversion: %t\n", float64(float32Val) == float64Val)

	// Safe floating-point comparison function
	fmt.Printf("Safe comparison function results:\n")
	fmt.Printf("floatEquals(0.1+0.2, 0.3): %t\n", floatEquals(0.1+0.2, 0.3, 1e-9))
	fmt.Printf("floatEquals(1.0, 1.0000001): %t\n", floatEquals(1.0, 1.0000001, 1e-6))

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Same floating-point precision issues
	// JavaScript: Number.EPSILON for comparison
	// Go: math.Abs() for absolute difference
	// Go: More explicit handling of special values
}

func demonstrateBooleanComparison() {
	fmt.Println("\n--- BOOLEAN COMPARISON ---")

	// Boolean values
	var trueVal bool = true
	var falseVal bool = false

	fmt.Printf("Boolean values: true = %t, false = %t\n", trueVal, falseVal)

	// Boolean equality
	fmt.Printf("true == true: %t\n", trueVal == true)
	fmt.Printf("false == false: %t\n", falseVal == false)
	fmt.Printf("true == false: %t\n", trueVal == falseVal)
	fmt.Printf("true != false: %t\n", trueVal != falseVal)

	// Boolean comparison with expressions
	var x int = 5
	var y int = 10

	var expr1 bool = x < y
	var expr2 bool = x > y

	fmt.Printf("x < y: %t\n", expr1)
	fmt.Printf("x > y: %t\n", expr2)
	fmt.Printf("(x < y) == (x > y): %t\n", expr1 == expr2)
	fmt.Printf("(x < y) != (x > y): %t\n", expr1 != expr2)

	// Boolean values in conditions
	if trueVal {
		fmt.Printf("true evaluates to true in if condition\n")
	}

	if !falseVal {
		fmt.Printf("!false evaluates to true in if condition\n")
	}

	// Comparing boolean functions
	var isPositive bool = x > 0
	var isEven bool = x%2 == 0

	fmt.Printf("x = %d\n", x)
	fmt.Printf("isPositive: %t\n", isPositive)
	fmt.Printf("isEven: %t\n", isEven)
	fmt.Printf("isPositive == isEven: %t\n", isPositive == isEven)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Truthy/falsy values (0, "", null, undefined)
	// JavaScript: == converts to boolean (1 == true)
	// Go: Only true/false boolean values
	// Go: No automatic conversion to boolean
}

func demonstrateTypeSpecificComparison() {
	fmt.Println("\n--- TYPE-SPECIFIC COMPARISON ---")

	// Rune comparison (Unicode code points)
	var char1 rune = 'A'
	var char2 rune = 'B'
	var char3 rune = 'A'

	fmt.Printf("Rune comparison:\n")
	fmt.Printf("'%c' == '%c': %t\n", char1, char2, char1 == char2)
	fmt.Printf("'%c' == '%c': %t\n", char1, char3, char1 == char3)
	fmt.Printf("'%c' < '%c': %t\n", char1, char2, char1 < char2)
	fmt.Printf("Unicode values: '%c'=%d, '%c'=%d\n", char1, char1, char2, char2)

	// Byte comparison
	var byte1 byte = 65 // ASCII 'A'
	var byte2 byte = 66 // ASCII 'B'

	fmt.Printf("Byte comparison:\n")
	fmt.Printf("%d == %d: %t\n", byte1, byte2, byte1 == byte2)
	fmt.Printf("%d < %d: %t\n", byte1, byte2, byte1 < byte2)
	fmt.Printf("byte vs rune: %d == %d: %t\n", byte1, char1, byte1 == byte(char1))

	// Pointer comparison
	var int1 int = 10
	var int2 int = 10
	var ptr1 *int = &int1
	var ptr2 *int = &int2
	var ptr3 *int = &int1

	fmt.Printf("Pointer comparison:\n")
	fmt.Printf("Values: *ptr1=%d, *ptr2=%d\n", *ptr1, *ptr2)
	fmt.Printf("ptr1 == ptr2: %t (different addresses)\n", ptr1 == ptr2)
	fmt.Printf("ptr1 == ptr3: %t (same address)\n", ptr1 == ptr3)
	fmt.Printf("*ptr1 == *ptr2: %t (same values)\n", *ptr1 == *ptr2)

	// Nil pointer comparison
	var nilPtr *int = nil
	fmt.Printf("ptr1 == nil: %t\n", ptr1 == nil)
	fmt.Printf("nilPtr == nil: %t\n", nilPtr == nil)

	// Array comparison
	var arr1 [3]int = [3]int{1, 2, 3}
	var arr2 [3]int = [3]int{1, 2, 3}
	var arr3 [3]int = [3]int{1, 2, 4}

	fmt.Printf("Array comparison:\n")
	fmt.Printf("arr1 == arr2: %t\n", arr1 == arr2)
	fmt.Printf("arr1 == arr3: %t\n", arr1 == arr3)

	// Struct comparison
	type Person struct {
		Name string
		Age  int
	}

	var person1 Person = Person{Name: "Alice", Age: 30}
	var person2 Person = Person{Name: "Alice", Age: 30}
	var person3 Person = Person{Name: "Bob", Age: 30}

	fmt.Printf("Struct comparison:\n")
	fmt.Printf("person1 == person2: %t\n", person1 == person2)
	fmt.Printf("person1 == person3: %t\n", person1 == person3)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Objects compared by reference
	// JavaScript: Arrays compared by reference
	// Go: Arrays compared by value
	// Go: Structs compared by value
	// Go: Pointers compared by address
}

func demonstrateComparisonChaining() {
	fmt.Println("\n--- COMPARISON CHAINING ---")

	// Go doesn't support chained comparisons like Python
	// 1 < x < 10 is NOT valid in Go

	var x int = 5

	// Wrong way (doesn't work in Go)
	// if 1 < x < 10 {  // Error!

	// Right way - use logical operators
	if 1 < x && x < 10 {
		fmt.Printf("%d is between 1 and 10\n", x)
	}

	// Multiple comparisons
	var a, b, c int = 3, 5, 7

	// Check if values are in ascending order
	if a < b && b < c {
		fmt.Printf("%d < %d < %d: ascending order\n", a, b, c)
	}

	// Check if all values are equal
	if a == b && b == c {
		fmt.Printf("All values are equal\n")
	} else {
		fmt.Printf("Values are not all equal\n")
	}

	// Complex comparison expressions
	var min, max int = 1, 10
	var values []int = []int{0, 5, 15, 8, 12}

	fmt.Printf("Values in range [%d, %d]:\n", min, max)
	for _, val := range values {
		if min <= val && val <= max {
			fmt.Printf("%d is in range\n", val)
		} else {
			fmt.Printf("%d is out of range\n", val)
		}
	}

	// Comparison with multiple conditions
	var score int = 85

	if score >= 90 {
		fmt.Printf("Score %d: Grade A\n", score)
	} else if score >= 80 && score < 90 {
		fmt.Printf("Score %d: Grade B\n", score)
	} else if score >= 70 && score < 80 {
		fmt.Printf("Score %d: Grade C\n", score)
	} else {
		fmt.Printf("Score %d: Grade F\n", score)
	}

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Same logical operator approach
	// JavaScript: No chained comparisons either
	// Go: Clear and explicit comparison logic
}

func demonstrateComparisonBestPractices() {
	fmt.Println("\n--- COMPARISON BEST PRACTICES ---")

	// Use explicit type conversion
	var int32Val int32 = 100
	var int64Val int64 = 100

	// Good: explicit conversion
	if int64(int32Val) == int64Val {
		fmt.Printf("Values are equal after conversion\n")
	}

	// Floating-point comparison with epsilon
	var floatVal1 float64 = 0.1 + 0.2
	var floatVal2 float64 = 0.3

	if floatEquals(floatVal1, floatVal2, 1e-9) {
		fmt.Printf("Floating-point values are equal within epsilon\n")
	}

	// String comparison considerations
	var userInput string = "  Hello  "
	var expected string = "Hello"

	// Good: trim whitespace before comparison
	if strings.TrimSpace(userInput) == expected {
		fmt.Printf("Input matches expected value\n")
	}

	// Case-insensitive comparison
	var str1 string = "Hello"
	var str2 string = "HELLO"

	if strings.EqualFold(str1, str2) {
		fmt.Printf("Strings are equal (case-insensitive)\n")
	}

	// Nil checking
	var ptr *int = nil

	if ptr != nil {
		fmt.Printf("Pointer value: %d\n", *ptr)
	} else {
		fmt.Printf("Pointer is nil\n")
	}

	// Range validation
	var age int = 25
	var minAge, maxAge int = 18, 65

	if age >= minAge && age <= maxAge {
		fmt.Printf("Age %d is valid\n", age)
	}

	// Zero value checking
	var count int = 0
	var name string = ""
	var isActive bool = false

	if count == 0 {
		fmt.Printf("Count is zero\n")
	}

	if name == "" {
		fmt.Printf("Name is empty\n")
	}

	if !isActive {
		fmt.Printf("Not active\n")
	}

	// BEST PRACTICES SUMMARY:
	fmt.Println("\nBest Practices Summary:")
	fmt.Println("1. Use explicit type conversion when comparing different types")
	fmt.Println("2. Use epsilon comparison for floating-point numbers")
	fmt.Println("3. Consider case sensitivity in string comparisons")
	fmt.Println("4. Trim whitespace before string comparisons if needed")
	fmt.Println("5. Always check for nil pointers before dereferencing")
	fmt.Println("6. Use logical operators for range checking")
	fmt.Println("7. Be aware of special floating-point values (NaN, Inf)")
	fmt.Println("8. Compare zero values explicitly")
	fmt.Println("9. Use appropriate comparison functions for complex types")
	fmt.Println("10. Make comparison logic clear and readable")
}

// Helper function for safe floating-point comparison
func floatEquals(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}
