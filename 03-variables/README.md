# Go Variables - Complete Guide

This folder contains comprehensive examples of Go variable handling, including declarations, constants, and scoping rules with detailed JavaScript comparisons.

## 📁 Directory Structure

### 01-declarations
- **File**: `declarations/main.go`
- **Topics**: Variable declarations (var, :=), zero values, type inference, multiple declarations, global vs local variables
- **Run**: `cd declarations && go run main.go`

### 02-constants
- **File**: `constants/main.go`
- **Topics**: Constants (const), typed vs untyped constants, iota, constant expressions, limitations, best practices
- **Run**: `cd constants && go run main.go`

### 03-scope
- **File**: `scope/main.go`
- **Topics**: Package scope, function scope, block scope, shadowing, closures, scope rules
- **Run**: `cd scope && go run main.go`

## 🎯 Learning Path

1. **Variable Declarations** - Master all ways to declare variables in Go
2. **Constants** - Understand Go's powerful constant system and iota
3. **Scope Rules** - Learn how Go handles variable visibility and lifetime

## 🔍 Key Differences from JavaScript

### Variable Declarations
- **Go**: Multiple declaration syntax (`var`, `:=`, explicit typing)
- **JavaScript**: Single declaration syntax (`var`, `let`, `const`)

### Type System
- **Go**: Static typing with compile-time type checking
- **JavaScript**: Dynamic typing with runtime type checking

### Constants
- **Go**: Compile-time constants with `iota` for enumerations
- **JavaScript**: Runtime constants with `const` keyword

### Scope Rules
- **Go**: Block-scoped variables, no hoisting
- **JavaScript**: `var` (function-scoped), `let`/`const` (block-scoped), hoisting

### Zero Values
- **Go**: Every type has a meaningful zero value
- **JavaScript**: Uninitialized variables are `undefined`

## 📚 Variable Declaration Methods

### 1. var Declaration
```go
// With type and value
var name string = "Alice"

// With type only (zero value)
var age int

// With value only (type inferred)
var salary = 50000.0

// Multiple declarations
var (
    username string = "alice"
    userID   int    = 123
    active   bool   = true
)
```

### 2. Short Declaration
```go
// Only inside functions
name := "Alice"
age := 30
x, y := 10, 20

// Mixed assignment
name, city := "Bob", "NYC"  // name reassigned, city new
```

### 3. Zero Values
```go
var s string    // ""
var i int       // 0
var b bool      // false
var f float64   // 0.0
var p *int      // nil
var slice []int // nil
var m map[string]int // nil
```

## 🔢 Constants

### Basic Constants
```go
const name = "Go"
const version = "1.21"
const pi = 3.14159
```

### Typed vs Untyped
```go
// Typed
const typedInt int = 42

// Untyped (more flexible)
const untypedInt = 42
var i32 int32 = untypedInt  // OK
var i64 int64 = untypedInt  // OK
```

### iota Enumerator
```go
const (
    StatusActive = iota  // 0
    StatusInactive      // 1
    StatusPending       // 2
)

const (
    KB = 1 << (10 * iota)  // 1024^0
    MB                     // 1024^1
    GB                     // 1024^2
    TB                     // 1024^3
)
```

## 🏗️ Scope Rules

### Package Scope
```go
var packageVar = "global"         // Private to package
var ExportedVar = "exported"      // Public to other packages
```

### Function Scope
```go
func myFunction() {
    var localVar = "function-scoped"
    // localVar accessible throughout function
}
```

### Block Scope
```go
{
    var blockVar = "block-scoped"
    // blockVar only accessible within this block
}
```

### Shadowing
```go
var name = "outer"
{
    var name = "inner"  // Shadows outer name
    fmt.Println(name)   // Prints "inner"
}
fmt.Println(name)       // Prints "outer"
```

## 🎨 Best Practices

### Variable Declarations
1. **Use short declaration (`:=`)** for local variables when type is obvious
2. **Use `var`** for zero values or when type needs to be explicit
3. **Use descriptive names** that explain the variable's purpose
4. **Keep scope as narrow as possible** for better maintainability

### Constants
1. **Use constants** for values that don't change
2. **Group related constants** together
3. **Use `iota`** for sequential enumerated values
4. **Prefer untyped constants** for flexibility
5. **Use ALL_CAPS** for exported constants (optional but common)

### Scope Management
1. **Declare variables close to their usage**
2. **Avoid shadowing** unless intentional
3. **Use package-level variables sparingly**
4. **Export only what's necessary** (capitalize only what other packages need)

## 🧪 Running Examples

```bash
# Run variable declarations examples
cd declarations && go run main.go

# Run constants examples
cd constants && go run main.go

# Run scope examples
cd scope && go run main.go
```

## 📖 Common Patterns

### Configuration Constants
```go
const (
    DatabaseURL = "localhost:5432"
    APITimeout  = 30
    MaxRetries  = 3
)
```

### Status Enumerations
```go
const (
    StatusPending = iota
    StatusApproved
    StatusRejected
)
```

### Error Handling with Zero Values
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")  // Return zero value
    }
    return a / b, nil
}
```

### Closure Pattern
```go
func createCounter(start int) func() int {
    counter := start
    return func() int {
        counter++
        return counter
    }
}
```

## ⚠️ Common Mistakes

1. **Unused variables** - Go compiler prevents unused variables
2. **Shadowing accidentally** - Can hide intended variables
3. **Modifying package-level variables** - Can cause unexpected behavior
4. **Using `:=` at package level** - Not allowed, use `var`
5. **Forgetting zero values** - All variables have default values

## 🔗 Related Topics

- **Data Types**: Understanding what types variables can hold
- **Operators**: How to manipulate variables
- **Functions**: How variables are passed and returned
- **Packages**: How variables are shared between files
- **Pointers**: How to reference variables indirectly

## 📖 Further Reading

- [Go Language Specification - Variables](https://golang.org/ref/spec#Variables)
- [Go Language Specification - Constants](https://golang.org/ref/spec#Constants)
- [Effective Go - Names](https://golang.org/doc/effective_go.html#names)
- [Go Blog - Constants](https://blog.golang.org/constants)
