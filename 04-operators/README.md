package main

import "fmt"

// This file demonstrates all Go operators with a comprehensive summary

func main() {
	fmt.Println("=== GO OPERATORS - COMPREHENSIVE SUMMARY ===")
	
	demonstrateOperatorCategories()
	demonstrateOperatorPrecedence()
	demonstrateOperatorComparisons()
}

func demonstrateOperatorCategories() {
	fmt.Println("\n--- OPERATOR CATEGORIES ---")
	
	// 1. ARITHMETIC OPERATORS
	fmt.Println("1. ARITHMETIC OPERATORS:")
	fmt.Printf("   + (addition): 5 + 3 = %d\n", 5+3)
	fmt.Printf("   - (subtraction): 5 - 3 = %d\n", 5-3)
	fmt.Printf("   * (multiplication): 5 * 3 = %d\n", 5*3)
	fmt.Printf("   / (division): 5 / 3 = %d\n", 5/3)
	fmt.Printf("   %% (modulo): 5 %% 3 = %d\n", 5%3)
	fmt.Printf("   ++ (increment): variable++\n")
	fmt.Printf("   -- (decrement): variable--\n")
	fmt.Printf("   + (unary plus): +5 = %d\n", +5)
	fmt.Printf("   - (unary minus): -5 = %d\n", -5)
	
	// 2. COMPARISON OPERATORS
	fmt.Println("\n2. COMPARISON OPERATORS:")
	fmt.Printf("   == (equal): 5 == 3 is %t\n", 5 == 3)
	fmt.Printf("   != (not equal): 5 != 3 is %t\n", 5 != 3)
	fmt.Printf("   < (less than): 5 < 3 is %t\n", 5 < 3)
	fmt.Printf("   > (greater than): 5 > 3 is %t\n", 5 > 3)
	fmt.Printf("   <= (less than or equal): 5 <= 3 is %t\n", 5 <= 3)
	fmt.Printf("   >= (greater than or equal): 5 >= 3 is %t\n", 5 >= 3)
	
	// 3. LOGICAL OPERATORS
	fmt.Println("\n3. LOGICAL OPERATORS:")
	fmt.Printf("   && (logical AND): true && false is %t\n", true && false)
	fmt.Printf("   || (logical OR): true || false is %t\n", true || false)
	fmt.Printf("   ! (logical NOT): !true is %t\n", !true)
	
	// 4. BITWISE OPERATORS
	fmt.Println("\n4. BITWISE OPERATORS:")
	fmt.Printf("   & (bitwise AND): 5 & 3 = %d\n", 5&3)
	fmt.Printf("   | (bitwise OR): 5 | 3 = %d\n", 5|3)
	fmt.Printf("   ^ (bitwise XOR): 5 ^ 3 = %d\n", 5^3)
	fmt.Printf("   &^ (bit clear): 5 &^ 3 = %d\n", 5&^3)
	fmt.Printf("   << (left shift): 5 << 1 = %d\n", 5<<1)
	fmt.Printf("   >> (right shift): 5 >> 1 = %d\n", 5>>1)
	fmt.Printf("   ^ (bitwise NOT): ^5 = %d\n", ^5)
	
	// 5. ASSIGNMENT OPERATORS
	fmt.Println("\n5. ASSIGNMENT OPERATORS:")
	fmt.Printf("   = (simple assignment)\n")
	fmt.Printf("   += (add and assign)\n")
	fmt.Printf("   -= (subtract and assign)\n")
	fmt.Printf("   *= (multiply and assign)\n")
	fmt.Printf("   /= (divide and assign)\n")
	fmt.Printf("   %%= (modulo and assign)\n")
	fmt.Printf("   &= (bitwise AND and assign)\n")
	fmt.Printf("   |= (bitwise OR and assign)\n")
	fmt.Printf("   ^= (bitwise XOR and assign)\n")
	fmt.Printf("   <<= (left shift and assign)\n")
	fmt.Printf("   >>= (right shift and assign)\n")
	fmt.Printf("   &^= (bit clear and assign)\n")
	
	// 6. POINTER OPERATORS
	fmt.Println("\n6. POINTER OPERATORS:")
	fmt.Printf("   & (address of operator)\n")
	fmt.Printf("   * (dereference operator)\n")
	
	// 7. CHANNEL OPERATORS
	fmt.Println("\n7. CHANNEL OPERATORS:")
	fmt.Printf("   <- (channel send/receive)\n")
}

func demonstrateOperatorPrecedence() {
	fmt.Println("\n--- OPERATOR PRECEDENCE (High to Low) ---")
	
	fmt.Println("1. Primary expressions: () [] -> .")
	fmt.Println("2. Unary operators: + - ! ^ * & <- (receive)")
	fmt.Println("3. Multiplicative: * / % << >> & &^")
	fmt.Println("4. Additive: + - | ^")
	fmt.Println("5. Comparison: == != < <= > >=")
	fmt.Println("6. Logical AND: &&")
	fmt.Println("7. Logical OR: ||")
	fmt.Println("8. Assignment: = += -= *= /= %= &= |= ^= <<= >>= &^=")
	
	// Demonstrate precedence with examples
	fmt.Println("\nPrecedence Examples:")
	fmt.Printf("2 + 3 * 4 = %d (multiplication first)\n", 2+3*4)
	fmt.Printf("(2 + 3) * 4 = %d (parentheses override)\n", (2+3)*4)
	fmt.Printf("5 > 3 && 2 < 4 = %t (comparison before logical)\n", 5 > 3 && 2 < 4)
	fmt.Printf("1 | 2 & 3 = %d (& before |)\n", 1|2&3)
	fmt.Printf("(1 | 2) & 3 = %d (parentheses change order)\n", (1|2)&3)
}

func demonstrateOperatorComparisons() {
	fmt.Println("\n--- COMPARISON WITH JAVASCRIPT ---")
	
	fmt.Println("SIMILARITIES:")
	fmt.Println("- Same arithmetic operators: +, -, *, /, %")
	fmt.Println("- Same comparison operators: ==, !=, <, >, <=, >=")
	fmt.Println("- Same logical operators: &&, ||, !")
	fmt.Println("- Same bitwise operators: &, |, ^, <<, >>")
	fmt.Println("- Same assignment operators: =, +=, -=, *=, /=, %=")
	fmt.Println("- Similar increment/decrement: ++, --")
	
	fmt.Println("\nDIFFERENCES:")
	fmt.Println("JavaScript vs Go:")
	fmt.Println("- JS: == (loose) vs === (strict) | Go: only == (always strict)")
	fmt.Println("- JS: Automatic type coercion | Go: Explicit type conversion")
	fmt.Println("- JS: Truthy/falsy values | Go: Only true/false")
	fmt.Println("- JS: ~ for bitwise NOT | Go: ^ for bitwise NOT")
	fmt.Println("- JS: >>> for unsigned right shift | Go: >> behavior depends on type")
	fmt.Println("- JS: ++var and var++ are expressions | Go: only var++ as statement")
	fmt.Println("- JS: Division always returns float | Go: Integer division truncates")
	fmt.Println("- JS: Logical operators can return non-boolean | Go: Always boolean")
	fmt.Println("- Go: &^ bit clear operator (not in JavaScript)")
	fmt.Println("- Go: Multiple assignment (a, b = 1, 2)")
	fmt.Println("- Go: Channel operators (<-)")
	fmt.Println("- Go: Pointer operators (&, *)")
	
	fmt.Println("\nGO ADVANTAGES:")
	fmt.Println("- Type safety prevents many runtime errors")
	fmt.Println("- Explicit type conversions make code clearer")
	fmt.Println("- No automatic type coercion reduces bugs")
	fmt.Println("- Consistent behavior across different data types")
	fmt.Println("- Better performance due to static typing")
	fmt.Println("- More predictable bitwise operations")
}

/*
OPERATORS LEARNING SUMMARY:

1. ARITHMETIC OPERATORS - Complete ✓
   - Basic operations: +, -, *, /, %
   - Unary operators: +, -
   - Increment/decrement: ++, --
   - Assignment operators: +=, -=, *=, /=, %=
   - Integer vs floating-point arithmetic
   - Overflow/underflow behavior
   - Type conversion requirements

2. COMPARISON OPERATORS - Complete ✓
   - Equality: ==, !=
   - Relational: <, >, <=, >=
   - Type-specific comparisons
   - Floating-point precision issues
   - String comparison (lexicographic)
   - Boolean comparison
   - Pointer comparison

3. LOGICAL OPERATORS - Complete ✓
   - AND: &&
   - OR: ||
   - NOT: !
   - Short-circuit evaluation
   - Truth tables
   - Complex logical expressions
   - Boolean best practices

4. BITWISE OPERATORS - Complete ✓
   - AND: &
   - OR: |
   - XOR: ^
   - NOT: ^
   - Bit clear: &^
   - Left shift: <<
   - Right shift: >>
   - Bit manipulation techniques
   - Practical applications

5. ASSIGNMENT OPERATORS - Complete ✓
   - Simple assignment: =
   - Compound assignment: +=, -=, *=, /=, %=
   - Bitwise assignment: &=, |=, ^=, <<=, >>=, &^=
   - Multiple assignment: a, b = 1, 2

6. OPERATOR PRECEDENCE - Complete ✓
   - Precedence rules
   - Associativity
   - Parentheses for clarity
   - Best practices

KEY TAKEAWAYS:
- Go is strongly typed - no automatic type conversion
- Explicit type conversions required for mixed-type operations
- Integer division truncates (no automatic float conversion)
- Logical operators only work with boolean values
- Bitwise operators work with integer types
- Short-circuit evaluation in logical operators
- Use parentheses for clarity in complex expressions
- Go's type system prevents many runtime errors
- Consistent and predictable behavior across operations
*/
