# Go Data Types - Complete Guide

This folder contains comprehensive examples of all Go data types with detailed explanations and JavaScript comparisons.

## 📁 Directory Structure

### 01-overview
- **File**: `01-overview.go`
- **Topics**: Basic introduction to Go's type system, categories of types, zero values, type safety
- **Run**: `go run 01-overview.go`

### 02-integers
- **File**: `integers/main.go`
- **Topics**: All integer types (int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint), aliases (byte, rune, uintptr), operations, limits, overflow behavior
- **Run**: `cd integers && go run main.go`

### 03-floating-point
- **File**: `floating-point/main.go`
- **Topics**: float32, float64, IEEE 754, precision, special values (Inf, NaN), mathematical functions
- **Run**: `cd floating-point && go run main.go`

### 04-complex
- **File**: `complex/main.go`
- **Topics**: complex64, complex128, complex number operations, mathematical functions, practical applications
- **Run**: `cd complex && go run main.go`

### 05-boolean
- **File**: `boolean/main.go`
- **Topics**: bool type, logical operations, truth tables, comparisons, practical patterns
- **Run**: `cd boolean && go run main.go`

### 06-string
- **File**: `string/main.go`
- **Topics**: string type, literals, UTF-8 encoding, Unicode, runes, string operations, strings package
- **Run**: `cd string && go run main.go`

## 🎯 Learning Path

1. **Start with overview** - Understand Go's type system philosophy
2. **Learn integers** - Master all integer types and their use cases
3. **Understand floating-point** - Learn about precision and mathematical operations
4. **Explore complex numbers** - Discover Go's built-in complex number support
5. **Master booleans** - Understand logical operations and patterns
6. **Deep dive into strings** - Learn about Unicode, UTF-8, and string manipulation

## 🔍 Key Differences from JavaScript

### Type System
- **Go**: Static typing with explicit type declarations
- **JavaScript**: Dynamic typing with implicit conversions

### Number Types
- **Go**: Multiple integer types (int8, int16, int32, int64, etc.) and floating-point types (float32, float64)
- **JavaScript**: Single `number` type (64-bit floating-point) + `BigInt`

### Strings
- **Go**: UTF-8 encoded, byte vs rune distinction
- **JavaScript**: UTF-16 encoded strings

### Booleans
- **Go**: Only `true` and `false` are valid boolean values
- **JavaScript**: Truthy/falsy values in boolean contexts

### Complex Numbers
- **Go**: Built-in complex64 and complex128 types
- **JavaScript**: No built-in complex number support

## 📚 Best Practices

1. **Use appropriate integer types** - Choose based on value range and memory constraints
2. **Default to float64** - Unless you have specific memory or precision requirements
3. **Use complex128** - For complex number calculations unless memory is critical
4. **Be explicit with types** - Go's type safety prevents many runtime errors
5. **Handle string encoding properly** - Use runes for Unicode character operations
6. **Understand zero values** - Every type has a meaningful zero value

## 🧪 Running Examples

```bash
# Run overview
go run 01-overview.go

# Run specific type examples
cd integers && go run main.go
cd floating-point && go run main.go
cd complex && go run main.go
cd boolean && go run main.go
cd string && go run main.go
```

## 📖 Further Reading

- [Go Language Specification - Types](https://golang.org/ref/spec#Types)
- [Effective Go - Data](https://golang.org/doc/effective_go.html#data)
- [Go Blog - Strings, bytes, runes and characters](https://blog.golang.org/strings)
- [Go Blog - Constants](https://blog.golang.org/constants)
