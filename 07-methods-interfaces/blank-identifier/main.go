package main

import (
	"fmt"
	"os"
	"strconv"
)

// === GO BLANK IDENTIFIER COMPREHENSIVE GUIDE ===

/*
BLANK IDENTIFIER PHILOSOPHY:
- The blank identifier (_) is a special identifier that discards values
- Used when you must provide a variable but don't need its value
- Common in Go for ignoring return values, unused imports, and interface checking
- Helps avoid "unused variable" compilation errors

COMPARISON WITH JAVASCRIPT:
// JavaScript - No direct equivalent, but similar concepts:
const [first, , third] = [1, 2, 3];  // Skip second element
const {name, ...rest} = obj;          // Destructuring with rest

// Go - Blank identifier
first, _, third := 1, 2, 3           // Skip second value
value, _ := strconv.Atoi("123")      // Ignore error
*/

// === INTERFACE DEFINITIONS ===

type Writer interface {
	Write([]byte) (int, error)
}

type Reader interface {
	Read([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// Custom types for examples
type MyWriter struct{}

func (w MyWriter) Write(data []byte) (int, error) {
	fmt.Printf("Writing: %s\n", string(data))
	return len(data), nil
}

type MyReader struct {
	data []byte
	pos  int
}

func (r *MyReader) Read(buf []byte) (int, error) {
	if r.pos >= len(r.data) {
		return 0, fmt.Errorf("EOF")
	}
	n := copy(buf, r.data[r.pos:])
	r.pos += n
	return n, nil
}

type MyCloser struct{}

func (c MyCloser) Close() error {
	fmt.Println("Closing resource")
	return nil
}

// Multiple return values function
func divide(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("division by zero")
	}
	return a / b, a % b, nil
}

// Function with unused parameters
func processData(data []int, _ string, _ bool) int {
	sum := 0
	for _, v := range data {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("=== GO BLANK IDENTIFIER COMPREHENSIVE GUIDE ===")

	// === IGNORING RETURN VALUES ===
	fmt.Println("\n1. IGNORING RETURN VALUES:")

	// Ignore error return value
	value, _ := strconv.Atoi("123")
	fmt.Printf("Parsed value: %d\n", value)

	// Ignore both error and other return values
	_, _, _ = divide(10, 3)
	fmt.Println("Division performed, results ignored")

	// Ignore specific return values
	quotient, _, _ := divide(10, 3)
	fmt.Printf("Quotient: %d (remainder and error ignored)\n", quotient)

	_, remainder, _ := divide(10, 3)
	fmt.Printf("Remainder: %d (quotient and error ignored)\n", remainder)

	// === RANGE LOOPS ===
	fmt.Println("\n2. RANGE LOOPS:")

	numbers := []int{10, 20, 30, 40, 50}

	// Ignore index, use only value
	fmt.Print("Values: ")
	for _, value := range numbers {
		fmt.Printf("%d ", value)
	}
	fmt.Println()

	// Ignore value, use only index
	fmt.Print("Indices: ")
	for index, _ := range numbers {
		fmt.Printf("%d ", index)
	}
	fmt.Println()

	// Ignore both (just iterate)
	count := 0
	for _, _ = range numbers {
		count++
	}
	fmt.Printf("Count: %d\n", count)

	// With maps
	data := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}

	// Ignore key, use only value
	fmt.Print("Map values: ")
	for _, value := range data {
		fmt.Printf("%d ", value)
	}
	fmt.Println()

	// Ignore value, use only key
	fmt.Print("Map keys: ")
	for key, _ := range data {
		fmt.Printf("%s ", key)
	}
	fmt.Println()

	// === MULTIPLE ASSIGNMENT ===
	fmt.Println("\n3. MULTIPLE ASSIGNMENT:")

	// Function returns multiple values
	first, second, third := 1, 2, 3
	fmt.Printf("All values: %d, %d, %d\n", first, second, third)

	// Ignore some values
	a, _, c := 1, 2, 3
	fmt.Printf("First and third: %d, %d\n", a, c)

	// Ignore middle values
	start, _, _, end := 1, 2, 3, 4
	fmt.Printf("Start and end: %d, %d\n", start, end)

	// === INTERFACE TYPE ASSERTIONS ===
	fmt.Println("\n4. INTERFACE TYPE ASSERTIONS:")

	var iface interface{} = "hello"

	// Type assertion with blank identifier (ignore ok value)
	if str, ok := iface.(string); ok {
		fmt.Printf("String value: %s\n", str)
	}

	// Check type without using the value
	if _, ok := iface.(int); ok {
		fmt.Println("It's an int")
	} else {
		fmt.Println("It's not an int")
	}

	// Type switch with blank identifier
	switch v := iface.(type) {
	case string:
		fmt.Printf("String: %s\n", v)
	case int:
		fmt.Printf("Integer: %d\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}

	// === INTERFACE COMPLIANCE CHECK ===
	fmt.Println("\n5. INTERFACE COMPLIANCE CHECK:")

	// Compile-time check that MyWriter implements Writer
	var _ Writer = MyWriter{}
	fmt.Println("MyWriter implements Writer interface")

	// Compile-time check that MyReader implements Reader
	var _ Reader = &MyReader{}
	fmt.Println("MyReader implements Reader interface")

	// Compile-time check that MyCloser implements Closer
	var _ Closer = MyCloser{}
	fmt.Println("MyCloser implements Closer interface")

	// Multiple interface checks
	var _ Writer = MyWriter{}
	var _ Closer = MyCloser{}

	// === UNUSED VARIABLES ===
	fmt.Println("\n6. UNUSED VARIABLES:")

	// Function with unused parameters
	result := processData([]int{1, 2, 3, 4, 5}, "unused", true)
	fmt.Printf("Process result: %d\n", result)

	// Unused variables in function
	unusedFunction := func() {
		used := 42
		_ = used // Prevent unused variable error
		fmt.Println("Function with unused variable")
	}

	unusedFunction()

	// === IMPORTING PACKAGES FOR SIDE EFFECTS ===
	fmt.Println("\n7. IMPORTING PACKAGES FOR SIDE EFFECTS:")

	// Note: This would be done at import level
	// import _ "net/http/pprof"  // Import for side effects only
	fmt.Println("Packages can be imported with blank identifier for side effects")

	// === STRUCT FIELD IGNORING ===
	fmt.Println("\n8. STRUCT FIELD IGNORING:")

	type Point struct {
		X, Y, Z float64
	}

	point := Point{X: 1.0, Y: 2.0, Z: 3.0}

	// Ignore some fields in assignment
	x, _, z := point.X, point.Y, point.Z
	fmt.Printf("X: %.1f, Z: %.1f (Y ignored)\n", x, z)

	// === CHANNEL OPERATIONS ===
	fmt.Println("\n9. CHANNEL OPERATIONS:")

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	// Ignore the received value
	_ = <-ch
	fmt.Println("Received value from channel (ignored)")

	// Ignore the ok value in channel receive
	var value2 int
	value2, _ = <-ch
	fmt.Printf("Received value: %d (ok ignored)\n", value2)

	// Ignore both value and ok
	_, _ = <-ch
	fmt.Println("Received from channel (both value and ok ignored)")

	close(ch)

	// === ERROR HANDLING ===
	fmt.Println("\n10. ERROR HANDLING:")

	// Sometimes you want to ignore errors (not recommended generally)
	file, _ := os.Open("nonexistent.txt")
	if file != nil {
		defer file.Close()
		fmt.Println("File opened successfully")
	} else {
		fmt.Println("File not found (error ignored)")
	}

	// Better approach: handle the error
	file2, err := os.Open("nonexistent.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	} else {
		defer file2.Close()
		fmt.Println("File opened successfully")
	}

	// === SLICE OPERATIONS ===
	fmt.Println("\n11. SLICE OPERATIONS:")

	slice := []int{1, 2, 3, 4, 5}

	// Ignore some elements when iterating
	for i, v := range slice {
		if i%2 == 0 {
			fmt.Printf("Even index %d: %d\n", i, v)
		} else {
			fmt.Printf("Odd index %d: _ (value ignored)\n", i)
		}
	}

	// === FUNCTION CALLS ===
	fmt.Println("\n12. FUNCTION CALLS:")

	// Function that returns multiple values
	multiReturn := func() (int, string, bool) {
		return 42, "hello", true
	}

	// Use only first return value
	number, _, _ := multiReturn()
	fmt.Printf("First return value: %d\n", number)

	// Use only middle return value
	_, text, _ := multiReturn()
	fmt.Printf("Middle return value: %s\n", text)

	// Use only last return value
	_, _, flag := multiReturn()
	fmt.Printf("Last return value: %t\n", flag)

	// === DEFER STATEMENTS ===
	fmt.Println("\n13. DEFER STATEMENTS:")

	// Defer function call ignoring return values
	defer func() {
		_, _ = fmt.Println("Deferred function executed")
	}()

	// === TESTING SCENARIOS ===
	fmt.Println("\n14. TESTING SCENARIOS:")

	// Test function that returns error
	testFunc := func() error {
		return fmt.Errorf("test error")
	}

	// In tests, sometimes you ignore errors for specific test cases
	_ = testFunc()
	fmt.Println("Test function called (error ignored)")

	// === INTERFACE EMBEDDING ===
	fmt.Println("\n15. INTERFACE EMBEDDING:")

	// Interface that embeds other interfaces
	type ReadWriteCloser interface {
		Reader
		Writer
		Closer
	}

	// Check if a type implements the embedded interface
	type CombinedType struct {
		*MyReader
		MyWriter
		MyCloser
	}

	var _ ReadWriteCloser = CombinedType{
		MyReader: &MyReader{},
		MyWriter: MyWriter{},
		MyCloser: MyCloser{},
	}

	fmt.Println("Composite type implements ReadWriteCloser")

	// === BENCHMARK SCENARIOS ===
	fmt.Println("\n16. BENCHMARK SCENARIOS:")

	// Benchmark function (conceptual)
	benchmarkFunc := func() {
		// Simulate some work
		for i := 0; i < 1000; i++ {
			_ = i * i // Ignore result
		}
	}

	benchmarkFunc()
	fmt.Println("Benchmark function executed")

	// === BEST PRACTICES ===
	fmt.Println("\n17. BEST PRACTICES:")

	fmt.Println("✓ Use _ to ignore unused return values")
	fmt.Println("✓ Use _ in range loops when you don't need index or value")
	fmt.Println("✓ Use _ for interface compliance checks")
	fmt.Println("✓ Use _ to prevent unused variable errors")
	fmt.Println("✓ Use _ for side-effect imports")
	fmt.Println("✗ Don't overuse _ - sometimes you should handle errors")
	fmt.Println("✗ Don't use _ when you actually need the value")

	// === COMMON PATTERNS ===
	fmt.Println("\n18. COMMON PATTERNS:")

	// Pattern 1: JSON unmarshaling with ignored fields
	jsonData := `{"name": "John", "age": 30, "city": "NYC"}`
	_ = jsonData // Simulate usage
	fmt.Println("JSON parsing pattern (fields can be ignored)")

	// Pattern 2: Database query with ignored columns
	queryResult := func() (int, string, bool, error) {
		return 1, "John", true, nil
	}

	id, name, _, err := queryResult()
	if err != nil {
		fmt.Printf("Query error: %v\n", err)
	} else {
		fmt.Printf("Query result: ID=%d, Name=%s\n", id, name)
	}

	// Pattern 3: HTTP response handling
	httpResponse := func() ([]byte, int, error) {
		return []byte("response"), 200, nil
	}

	body, _, err := httpResponse()
	if err != nil {
		fmt.Printf("HTTP error: %v\n", err)
	} else {
		fmt.Printf("HTTP response body: %s\n", string(body))
	}

	// Pattern 4: Concurrent operations
	done := make(chan bool)
	go func() {
		// Simulate work
		for i := 0; i < 5; i++ {
			_ = i // Ignore loop variable
		}
		done <- true
	}()

	_ = <-done // Wait for completion, ignore value
	fmt.Println("Concurrent operation completed")

	fmt.Println("\n=== END OF BLANK IDENTIFIER GUIDE ===")
}

/*
RUNNING THE PROGRAM:
go run main.go

LEARNING POINTS:

1. BLANK IDENTIFIER (_):
   - Special identifier that discards values
   - Used to ignore unused variables, return values, and imports
   - Helps avoid compilation errors

2. COMMON USES:
   - Ignore return values: value, _ := function()
   - Ignore loop variables: for _, v := range slice
   - Ignore error values: result, _ := risky_operation()
   - Interface compliance: var _ Interface = Type{}

3. RANGE LOOPS:
   - for _, value := range slice (ignore index)
   - for index, _ := range slice (ignore value)
   - for _, _ = range slice (ignore both)

4. MULTIPLE ASSIGNMENT:
   - a, _, c := 1, 2, 3 (ignore middle value)
   - first, _, _, last := multi_return() (ignore middle values)

5. INTERFACE PATTERNS:
   - Type assertions: if _, ok := value.(Type); ok
   - Compliance checks: var _ Interface = Type{}
   - Type switches: switch v := value.(type)

6. ERROR HANDLING:
   - Sometimes ignore errors: result, _ := operation()
   - Generally better to handle errors explicitly
   - Use blank identifier sparingly for errors

7. BEST PRACTICES:
   - Use _ when you truly don't need the value
   - Don't ignore errors unless you have a good reason
   - Use _ for interface compliance checks
   - Use _ to prevent unused variable errors

8. COMMON PATTERNS:
   - JSON unmarshaling with ignored fields
   - Database queries with ignored columns
   - HTTP responses with ignored status codes
   - Concurrent operations with ignored return values

This demonstrates comprehensive blank identifier usage in Go
for cleaner code and avoiding unused variable errors.
*/
