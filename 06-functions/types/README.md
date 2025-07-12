# Function Types in Go

## Overview
Function types in Go allow you to treat functions as first-class values that can be:
- Assigned to variables
- Passed as parameters
- Returned from functions
- Stored in data structures

## Key Concepts

### 1. Function Type Declaration
```go
type BinaryOperation func(int, int) int
type UnaryOperation func(int) int
type Predicate func(int) bool
```

### 2. Function Types as Parameters
Functions can accept other functions as parameters, enabling higher-order functions.

### 3. Function Types as Return Values
Functions can return other functions, enabling function factories and closures.

### 4. Format Specifiers
Go's `fmt` package provides comprehensive format specifiers for different data types:

#### Integer Formats
- `%d` - Decimal
- `%b` - Binary
- `%o` - Octal
- `%x` - Hexadecimal (lowercase)
- `%X` - Hexadecimal (uppercase)
- `%c` - Character
- `%q` - Quoted character
- `%U` - Unicode

#### Float Formats
- `%f` - Decimal
- `%e` - Scientific notation (lowercase)
- `%E` - Scientific notation (uppercase)
- `%g` - Compact format
- `%G` - Compact format (uppercase)

#### String Formats
- `%s` - String
- `%q` - Quoted string
- `%x` - Hex dump
- `%X` - Hex dump (uppercase)

#### Universal Formats
- `%v` - Default format
- `%+v` - Struct with field names
- `%#v` - Go-syntax representation
- `%T` - Type of value
- `%p` - Pointer address

#### Width and Precision
- `%5d` - Width 5
- `%-5d` - Left-aligned
- `%05d` - Zero-padded
- `%10.2f` - Width 10, 2 decimal places
- `%*d` - Dynamic width
- `%.*f` - Dynamic precision

## JavaScript Comparison

JavaScript doesn't have explicit function types, but TypeScript provides similar functionality:

```typescript
// TypeScript function types
type BinaryOperation = (a: number, b: number) => number;
type UnaryOperation = (x: number) => number;

const add: BinaryOperation = (a, b) => a + b;
const square: UnaryOperation = (x) => x * x;

// Higher-order functions
function calculate(a: number, b: number, op: BinaryOperation): number {
    return op(a, b);
}

// Function factories
function createAdder(x: number): UnaryOperation {
    return (y: number) => x + y;
}

// Closures
function createCounter(): () => number {
    let count = 0;
    return () => ++count;
}
```

## Best Practices

1. **Use Meaningful Names**: Give function types descriptive names
2. **Group Related Types**: Keep related function types together
3. **Leverage Closures**: Use closures for state management
4. **Function Composition**: Prefer composition over inheritance
5. **Keep Signatures Simple**: Avoid overly complex function signatures
6. **Document Contracts**: Clearly document function type contracts
7. **Use Interfaces**: When function types get complex, consider interfaces

## Common Patterns

### 1. Callback Pattern
```go
func processData(data []int, callback func(int) int) []int {
    result := make([]int, len(data))
    for i, v := range data {
        result[i] = callback(v)
    }
    return result
}
```

### 2. Factory Pattern
```go
func createValidator(min, max int) func(int) bool {
    return func(value int) bool {
        return value >= min && value <= max
    }
}
```

### 3. Decorator Pattern
```go
func logDecorator(fn func(int) int) func(int) int {
    return func(x int) int {
        fmt.Printf("Calling function with %d\n", x)
        result := fn(x)
        fmt.Printf("Function returned %d\n", result)
        return result
    }
}
```

## Running the Code

```bash
cd go-basics/06-functions/types
go run main.go
```

This will demonstrate all function type concepts with practical examples and comprehensive format specifier usage.
