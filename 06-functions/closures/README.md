# Closures in Go

## Overview
Closures are functions that capture and retain access to variables from their outer scope. They're powerful for creating stateful functions, implementing design patterns, and functional programming techniques.

## Key Concepts

### 1. Variable Capture
Closures capture variables from their enclosing scope:
- **By Reference**: Variables are captured by reference, allowing modification
- **By Value**: Create new variables in each iteration to capture by value

### 2. Common Patterns

#### Counter Pattern
```go
func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

#### Factory Pattern
```go
func createAdder(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}
```

#### State Management
```go
func createBankAccount(balance float64) (func(float64) float64, func() float64) {
    deposit := func(amount float64) float64 {
        balance += amount
        return balance
    }
    
    getBalance := func() float64 {
        return balance
    }
    
    return deposit, getBalance
}
```

## JavaScript Comparison

JavaScript closures work similarly:

```javascript
// JavaScript counter
function createCounter() {
    let count = 0;
    return function() {
        return ++count;
    };
}

// JavaScript adder
function createAdder(x) {
    return function(y) {
        return x + y;
    };
}

// JavaScript bank account
function createBankAccount(initialBalance) {
    let balance = initialBalance;
    
    return {
        deposit: function(amount) {
            balance += amount;
            return balance;
        },
        getBalance: function() {
            return balance;
        }
    };
}

// JavaScript loop variable capture issue
for (var i = 0; i < 3; i++) {
    setTimeout(function() {
        console.log(i); // Prints 3, 3, 3
    }, 100);
}

// Fixed with let or IIFE
for (let i = 0; i < 3; i++) {
    setTimeout(function() {
        console.log(i); // Prints 0, 1, 2
    }, 100);
}
```

## Common Use Cases

### 1. Configuration
```go
func createLogger(prefix string) func(string) {
    return func(message string) {
        fmt.Printf("[%s] %s\n", prefix, message)
    }
}
```

### 2. Caching/Memoization
```go
func createMemoizer() func(int) int {
    cache := make(map[int]int)
    return func(n int) int {
        if result, exists := cache[n]; exists {
            return result
        }
        // Calculate and cache result
        result := expensiveCalculation(n)
        cache[n] = result
        return result
    }
}
```

### 3. Event Handling
```go
func createEventHandler() (func(string, interface{}), func()) {
    var events []string
    
    handleEvent := func(eventType string, data interface{}) {
        events = append(events, fmt.Sprintf("%s: %v", eventType, data))
    }
    
    showHistory := func() {
        for _, event := range events {
            fmt.Println(event)
        }
    }
    
    return handleEvent, showHistory
}
```

### 4. Functional Programming
```go
func createFilter(predicate func(int) bool) func([]int) []int {
    return func(slice []int) []int {
        var result []int
        for _, value := range slice {
            if predicate(value) {
                result = append(result, value)
            }
        }
        return result
    }
}
```

## Common Pitfalls

### 1. Loop Variable Capture
```go
// WRONG - captures loop variable by reference
var functions []func() int
for i := 0; i < 3; i++ {
    functions = append(functions, func() int {
        return i // All functions return 3
    })
}

// CORRECT - captures by value
var functions []func() int
for i := 0; i < 3; i++ {
    i := i // Create new variable
    functions = append(functions, func() int {
        return i
    })
}
```

### 2. Memory Leaks
Be careful with closures that capture large objects or slices, as they prevent garbage collection.

### 3. Unexpected Mutations
Variables captured by reference can be modified by the closure, which might not be intended.

## Advanced Patterns

### 1. Middleware Pattern
```go
func createMiddleware(handler func(string) string) func(string) string {
    return func(input string) string {
        // Pre-processing
        result := handler(input)
        // Post-processing
        return result
    }
}
```

### 2. Observer Pattern
```go
func createObserver() (func(func(string)), func(string)) {
    var observers []func(string)
    
    subscribe := func(observer func(string)) {
        observers = append(observers, observer)
    }
    
    notify := func(message string) {
        for _, observer := range observers {
            observer(message)
        }
    }
    
    return subscribe, notify
}
```

### 3. State Machine
```go
func createStateMachine() (func(string), func() string) {
    states := []string{"idle", "running", "paused", "stopped"}
    currentState := 0
    
    transition := func(action string) {
        // State transition logic
    }
    
    getState := func() string {
        return states[currentState]
    }
    
    return transition, getState
}
```

## Best Practices

1. **Use for Encapsulation**: Closures provide data privacy
2. **Be Aware of Capture**: Understand reference vs value capture
3. **Use for Configuration**: Great for dependency injection
4. **Leverage for FP**: Perfect for functional programming patterns
5. **Event Handling**: Excellent for callbacks and event systems
6. **Watch Loop Variables**: Be careful with loop variable capture
7. **Memory Management**: Consider memory implications
8. **Factory Pattern**: Use closures for object creation
9. **Middleware**: Perfect for decorator patterns
10. **Clean Up**: Handle resource cleanup when needed

## Performance Considerations

- Closures have slight overhead due to variable capture
- Memory usage can be higher due to retained variables
- Use profiling to identify closure-related performance issues
- Consider alternatives for performance-critical code

## Running the Code

```bash
cd go-basics/06-functions/closures
go run main.go
```

This will demonstrate all closure patterns with practical examples and comprehensive coverage of closure concepts in Go.
