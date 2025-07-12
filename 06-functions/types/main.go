package main

import "fmt"

// === FUNCTION TYPES ===

// 1. Function type declaration
type BinaryOperation func(int, int) int
type UnaryOperation func(int) int
type Predicate func(int) bool
type StringProcessor func(string) string

// 2. Function type as parameter
func calculate(a, b int, op BinaryOperation) int {
	return op(a, b)
}

// 3. Function type as return value
func getOperation(operationType string) BinaryOperation {
	switch operationType {
	case "add":
		return func(a, b int) int { return a + b }
	case "multiply":
		return func(a, b int) int { return a * b }
	case "subtract":
		return func(a, b int) int { return a - b }
	default:
		return func(a, b int) int { return 0 }
	}
}

// 4. Function variables
var (
	add      BinaryOperation = func(a, b int) int { return a + b }
	multiply BinaryOperation = func(a, b int) int { return a * b }
	square   UnaryOperation  = func(x int) int { return x * x }
	isEven   Predicate       = func(x int) bool { return x%2 == 0 }
)

// 5. Function literals and anonymous functions
func demonstrateFunctionLiterals() {
	// Function literal assigned to variable
	double := func(x int) int {
		return x * 2
	}

	// Immediately invoked function expression (IIFE)
	result := func(x, y int) int {
		return x*x + y*y
	}(3, 4)

	fmt.Printf("Double of 5: %d\n", double(5))
	fmt.Printf("3² + 4² = %d\n", result)
}

// 6. Higher-order functions
func applyUnary(value int, op UnaryOperation) int {
	return op(value)
}

func applyBinary(a, b int, op BinaryOperation) int {
	return op(a, b)
}

func filter(slice []int, predicate Predicate) []int {
	var result []int
	for _, value := range slice {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

func mapValues(slice []int, op UnaryOperation) []int {
	result := make([]int, len(slice))
	for i, value := range slice {
		result[i] = op(value)
	}
	return result
}

// 7. Function composition
func compose(f, g UnaryOperation) UnaryOperation {
	return func(x int) int {
		return f(g(x))
	}
}

// 8. Currying functions
func addCurried(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func multiplyCurried(a int) func(int) int {
	return func(b int) int {
		return a * b
	}
}

// 9. Function with closure
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func createAccumulator(initial int) func(int) int {
	sum := initial
	return func(value int) int {
		sum += value
		return sum
	}
}

// 10. Complex function types
type EventHandler func(string, interface{})
type Validator func(interface{}) error
type Transformer func(interface{}) interface{}

func processEvent(event string, data interface{}, handler EventHandler) {
	handler(event, data)
}

// 11. Function type with methods
type Calculator struct {
	operation BinaryOperation
}

func (c *Calculator) SetOperation(op BinaryOperation) {
	c.operation = op
}

func (c *Calculator) Calculate(a, b int) int {
	if c.operation == nil {
		return 0
	}
	return c.operation(a, b)
}

// 12. Function slice
type OperationList []BinaryOperation

func (ol OperationList) ApplyAll(a, b int) []int {
	results := make([]int, len(ol))
	for i, op := range ol {
		results[i] = op(a, b)
	}
	return results
}

// 13. Function with interface
type Processor interface {
	Process(int) int
}

type FunctionProcessor struct {
	fn UnaryOperation
}

func (fp FunctionProcessor) Process(x int) int {
	return fp.fn(x)
}

// 14. Function factories
func createValidator(min, max int) Predicate {
	return func(value int) bool {
		return value >= min && value <= max
	}
}

func createFormatter(prefix, suffix string) StringProcessor {
	return func(s string) string {
		return prefix + s + suffix
	}
}

// 15. FORMAT SPECIFIERS EXPLANATION
func demonstrateFormatSpecifiers() {
	fmt.Println("=== FORMAT SPECIFIERS COMPREHENSIVE GUIDE ===")

	// Integer format specifiers
	num := 42
	fmt.Println("\n--- INTEGER FORMAT SPECIFIERS ---")
	fmt.Printf("%%d (decimal): %d\n", num)
	fmt.Printf("%%b (binary): %b\n", num)
	fmt.Printf("%%o (octal): %o\n", num)
	fmt.Printf("%%x (hex lowercase): %x\n", num)
	fmt.Printf("%%X (hex uppercase): %X\n", num)
	fmt.Printf("%%c (character): %c\n", 65) // ASCII 'A'
	fmt.Printf("%%q (quoted character): %q\n", 65)
	fmt.Printf("%%U (Unicode): %U\n", 65)

	// Float format specifiers
	pi := 3.14159265359
	fmt.Println("\n--- FLOAT FORMAT SPECIFIERS ---")
	fmt.Printf("%%f (decimal): %f\n", pi)
	fmt.Printf("%%.2f (2 decimal places): %.2f\n", pi)
	fmt.Printf("%%e (scientific lowercase): %e\n", pi)
	fmt.Printf("%%E (scientific uppercase): %E\n", pi)
	fmt.Printf("%%g (compact): %g\n", pi)
	fmt.Printf("%%G (compact uppercase): %G\n", pi)

	// String format specifiers
	str := "Hello, World!"
	fmt.Println("\n--- STRING FORMAT SPECIFIERS ---")
	fmt.Printf("%%s (string): %s\n", str)
	fmt.Printf("%%q (quoted string): %q\n", str)
	fmt.Printf("%%x (hex dump): %x\n", str)
	fmt.Printf("%%X (hex dump uppercase): %X\n", str)

	// Boolean format specifiers
	flag := true
	fmt.Println("\n--- BOOLEAN FORMAT SPECIFIERS ---")
	fmt.Printf("%%t (boolean): %t\n", flag)
	fmt.Printf("%%t (boolean false): %t\n", false)

	// Pointer format specifiers
	ptr := &num
	fmt.Println("\n--- POINTER FORMAT SPECIFIERS ---")
	fmt.Printf("%%p (pointer): %p\n", ptr)
	fmt.Printf("%%v (value): %v\n", ptr)
	fmt.Printf("%%#v (Go-syntax): %#v\n", ptr)

	// Type format specifiers
	fmt.Println("\n--- TYPE FORMAT SPECIFIERS ---")
	fmt.Printf("%%T (type): %T\n", num)
	fmt.Printf("%%T (type): %T\n", pi)
	fmt.Printf("%%T (type): %T\n", str)
	fmt.Printf("%%T (type): %T\n", flag)

	// Width and precision
	fmt.Println("\n--- WIDTH AND PRECISION ---")
	fmt.Printf("%%5d (width 5): '%5d'\n", num)
	fmt.Printf("%%-5d (left align): '%-5d'\n", num)
	fmt.Printf("%%05d (zero pad): '%05d'\n", num)
	fmt.Printf("%%10.2f (width 10, 2 decimals): '%10.2f'\n", pi)
	fmt.Printf("%%-10.2f (left align): '%-10.2f'\n", pi)

	// Plus and space flags
	fmt.Println("\n--- FLAGS ---")
	fmt.Printf("%%+d (always sign): %+d\n", num)
	fmt.Printf("%%+d (always sign): %+d\n", -num)
	fmt.Printf("%% d (space for positive): % d\n", num)
	fmt.Printf("%% d (space for positive): % d\n", -num)

	// Hash flag
	fmt.Println("\n--- HASH FLAG ---")
	fmt.Printf("%%#x (with 0x prefix): %#x\n", num)
	fmt.Printf("%%#o (with 0 prefix): %#o\n", num)
	fmt.Printf("%%#v (Go-syntax): %#v\n", str)

	// Complex types
	slice := []int{1, 2, 3}
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println("\n--- COMPLEX TYPES ---")
	fmt.Printf("%%v (slice): %v\n", slice)
	fmt.Printf("%%#v (slice Go-syntax): %#v\n", slice)
	fmt.Printf("%%v (map): %v\n", m)
	fmt.Printf("%%#v (map Go-syntax): %#v\n", m)

	// Struct formatting
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Alice", Age: 30}
	fmt.Println("\n--- STRUCT FORMATTING ---")
	fmt.Printf("%%v (struct): %v\n", person)
	fmt.Printf("%%+v (struct with fields): %+v\n", person)
	fmt.Printf("%%#v (struct Go-syntax): %#v\n", person)

	// Error handling in formatting
	fmt.Println("\n--- ERROR HANDLING ---")
	fmt.Printf("%%v (nil): %v\n", nil)
	fmt.Printf("%%T (nil): %T\n", nil)

	// Positional arguments
	fmt.Println("\n--- POSITIONAL ARGUMENTS ---")
	fmt.Printf("%%[1]d %%[2]d %%[1]d: %[1]d %[2]d %[1]d\n", 10, 20)

	// Width and precision with *
	fmt.Println("\n--- DYNAMIC WIDTH/PRECISION ---")
	fmt.Printf("%%*d (dynamic width): '%*d'\n", 8, num)
	fmt.Printf("%%.*f (dynamic precision): '%.*f'\n", 3, pi)
	fmt.Printf("%%*.*f (dynamic both): '%*.*f'\n", 10, 2, pi)
}

func main() {
	fmt.Println("=== GO FUNCTION TYPES ===")

	// === FUNCTION TYPE DECLARATIONS ===
	fmt.Println("\n--- FUNCTION TYPE DECLARATIONS ---")
	fmt.Printf("add(5, 3) = %d\n", add(5, 3))
	fmt.Printf("multiply(4, 6) = %d\n", multiply(4, 6))
	fmt.Printf("square(7) = %d\n", square(7))
	fmt.Printf("isEven(8) = %t\n", isEven(8))
	fmt.Printf("isEven(9) = %t\n", isEven(9))

	/*
		JavaScript comparison:
		// JavaScript doesn't have explicit function types,
		// but you can use TypeScript for type safety
		type BinaryOperation = (a: number, b: number) => number;
		type UnaryOperation = (x: number) => number;

		const add: BinaryOperation = (a, b) => a + b;
		const square: UnaryOperation = (x) => x * x;
	*/

	// === FUNCTION TYPES AS PARAMETERS ===
	fmt.Println("\n--- FUNCTION TYPES AS PARAMETERS ---")
	result := calculate(10, 5, add)
	fmt.Printf("calculate(10, 5, add) = %d\n", result)

	result = calculate(10, 5, multiply)
	fmt.Printf("calculate(10, 5, multiply) = %d\n", result)

	// === FUNCTION TYPES AS RETURN VALUES ===
	fmt.Println("\n--- FUNCTION TYPES AS RETURN VALUES ---")
	addOp := getOperation("add")
	multiplyOp := getOperation("multiply")

	fmt.Printf("addOp(8, 3) = %d\n", addOp(8, 3))
	fmt.Printf("multiplyOp(8, 3) = %d\n", multiplyOp(8, 3))

	// === FUNCTION LITERALS ===
	fmt.Println("\n--- FUNCTION LITERALS ---")
	demonstrateFunctionLiterals()

	// === HIGHER-ORDER FUNCTIONS ===
	fmt.Println("\n--- HIGHER-ORDER FUNCTIONS ---")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Filter even numbers
	evens := filter(numbers, isEven)
	fmt.Printf("Even numbers: %v\n", evens)

	// Map square function
	squares := mapValues(evens, square)
	fmt.Printf("Squares of evens: %v\n", squares)

	// Custom predicates
	isPositive := func(x int) bool { return x > 0 }
	positives := filter([]int{-3, -1, 0, 1, 2, 3}, isPositive)
	fmt.Printf("Positive numbers: %v\n", positives)

	// === FUNCTION COMPOSITION ===
	fmt.Println("\n--- FUNCTION COMPOSITION ---")
	double := func(x int) int { return x * 2 }
	increment := func(x int) int { return x + 1 }

	// Compose double and increment
	doubleAndIncrement := compose(increment, double)
	fmt.Printf("doubleAndIncrement(5) = %d\n", doubleAndIncrement(5))

	// === CURRYING ===
	fmt.Println("\n--- CURRYING ---")
	add5 := addCurried(5)
	multiply3 := multiplyCurried(3)

	fmt.Printf("add5(10) = %d\n", add5(10))
	fmt.Printf("multiply3(7) = %d\n", multiply3(7))

	// === CLOSURES ===
	fmt.Println("\n--- CLOSURES ---")
	counter := createCounter()
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())

	accumulator := createAccumulator(100)
	fmt.Printf("Accumulator: %d\n", accumulator(10))
	fmt.Printf("Accumulator: %d\n", accumulator(20))
	fmt.Printf("Accumulator: %d\n", accumulator(5))

	// === FUNCTION WITH METHODS ===
	fmt.Println("\n--- FUNCTION WITH METHODS ---")
	calc := &Calculator{}
	calc.SetOperation(add)
	fmt.Printf("Calculator add: %d\n", calc.Calculate(15, 25))

	calc.SetOperation(multiply)
	fmt.Printf("Calculator multiply: %d\n", calc.Calculate(6, 7))

	// === FUNCTION SLICE ===
	fmt.Println("\n--- FUNCTION SLICE ---")
	operations := OperationList{add, multiply, func(a, b int) int { return a - b }}
	results := operations.ApplyAll(12, 4)
	fmt.Printf("All operations on 12, 4: %v\n", results)

	// === FUNCTION WITH INTERFACE ===
	fmt.Println("\n--- FUNCTION WITH INTERFACE ---")
	processor := FunctionProcessor{fn: square}
	fmt.Printf("Process 6: %d\n", processor.Process(6))

	// === FUNCTION FACTORIES ===
	fmt.Println("\n--- FUNCTION FACTORIES ---")
	validator := createValidator(1, 10)
	fmt.Printf("Validator(5): %t\n", validator(5))
	fmt.Printf("Validator(15): %t\n", validator(15))

	formatter := createFormatter("[", "]")
	fmt.Printf("Formatter: %s\n", formatter("Hello"))

	// === EVENT HANDLING ===
	fmt.Println("\n--- EVENT HANDLING ---")
	eventHandler := func(event string, data interface{}) {
		fmt.Printf("Event: %s, Data: %v\n", event, data)
	}

	processEvent("click", map[string]int{"x": 100, "y": 200}, eventHandler)
	processEvent("keypress", "Enter", eventHandler)

	// === FORMAT SPECIFIERS DEMONSTRATION ===
	demonstrateFormatSpecifiers()

	// === FUNCTION TYPE BEST PRACTICES ===
	fmt.Println("\n--- FUNCTION TYPE BEST PRACTICES ---")
	fmt.Println("1. Use meaningful function type names")
	fmt.Println("2. Group related function types together")
	fmt.Println("3. Use function types for callback patterns")
	fmt.Println("4. Leverage closures for state management")
	fmt.Println("5. Use higher-order functions for flexibility")
	fmt.Println("6. Prefer composition over inheritance")
	fmt.Println("7. Use function factories for configuration")
	fmt.Println("8. Keep function signatures simple")
	fmt.Println("9. Document function type contracts")
	fmt.Println("10. Use interfaces when function types get complex")
}
