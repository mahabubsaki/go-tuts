package main

import (
	"fmt"
	"unsafe"
)

// === GO ALLOCATION WITH NEW COMPREHENSIVE GUIDE ===

/*
NEW FUNCTION PHILOSOPHY:
- new() allocates memory for a type and returns a pointer
- Always returns a pointer to zero-valued memory
- Different from make() which is for slices, maps, and channels
- Memory is zero-initialized (zero value of the type)

COMPARISON WITH JAVASCRIPT:
// JavaScript - Object creation
const obj = new Object();        // Constructor function
const obj2 = {};                 // Object literal
const arr = new Array(5);        // Array constructor

// Go - Memory allocation
ptr := new(int)                  // Allocates int, returns *int
var i int                        // Zero value on stack
ptr2 := &i                       // Address of existing variable
*/

// === CUSTOM TYPES FOR EXAMPLES ===

type Person struct {
	Name    string
	Age     int
	Email   string
	Address *Address
}

type Address struct {
	Street  string
	City    string
	Country string
}

type Node struct {
	Value int
	Next  *Node
}

// Counter for demonstration
type Counter struct {
	value int
}

func (c *Counter) Reset() {
	c.value = 0
}

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func main() {
	fmt.Println("=== GO ALLOCATION WITH NEW COMPREHENSIVE GUIDE ===")

	// === BASIC NEW ALLOCATION ===
	fmt.Println("\n1. BASIC NEW ALLOCATION:")

	// Allocate basic types
	intPtr := new(int)
	floatPtr := new(float64)
	stringPtr := new(string)
	boolPtr := new(bool)

	fmt.Printf("int pointer: %p, value: %d\n", intPtr, *intPtr)
	fmt.Printf("float pointer: %p, value: %f\n", floatPtr, *floatPtr)
	fmt.Printf("string pointer: %p, value: %q\n", stringPtr, *stringPtr)
	fmt.Printf("bool pointer: %p, value: %t\n", boolPtr, *boolPtr)

	// Modify values through pointers
	*intPtr = 42
	*floatPtr = 3.14159
	*stringPtr = "Hello, World!"
	*boolPtr = true

	fmt.Println("\nAfter modification:")
	fmt.Printf("int value: %d\n", *intPtr)
	fmt.Printf("float value: %f\n", *floatPtr)
	fmt.Printf("string value: %q\n", *stringPtr)
	fmt.Printf("bool value: %t\n", *boolPtr)

	// === NEW VS VARIABLE DECLARATION ===
	fmt.Println("\n2. NEW VS VARIABLE DECLARATION:")

	// Using new()
	newInt := new(int)
	fmt.Printf("new(int): %p, value: %d\n", newInt, *newInt)

	// Using var declaration
	var varInt int
	fmt.Printf("var int: %p, value: %d\n", &varInt, varInt)

	// Using var with pointer
	var ptrInt *int = &varInt
	fmt.Printf("var pointer: %p, value: %d\n", ptrInt, *ptrInt)

	// Using short declaration
	shortInt := 0
	fmt.Printf("short declaration: %p, value: %d\n", &shortInt, shortInt)

	// === STRUCT ALLOCATION ===
	fmt.Println("\n3. STRUCT ALLOCATION:")

	// Allocate struct with new()
	personPtr := new(Person)
	fmt.Printf("Person pointer: %p\n", personPtr)
	fmt.Printf("Person zero value: %+v\n", *personPtr)

	// Initialize struct fields
	personPtr.Name = "Alice"
	personPtr.Age = 30
	personPtr.Email = "alice@example.com"

	fmt.Printf("Person after init: %+v\n", *personPtr)

	// Allocate nested struct
	addressPtr := new(Address)
	addressPtr.Street = "123 Main St"
	addressPtr.City = "New York"
	addressPtr.Country = "USA"

	personPtr.Address = addressPtr
	fmt.Printf("Person with address: %+v\n", *personPtr)

	// === ARRAY ALLOCATION ===
	fmt.Println("\n4. ARRAY ALLOCATION:")

	// Allocate array with new()
	arrayPtr := new([5]int)
	fmt.Printf("Array pointer: %p\n", arrayPtr)
	fmt.Printf("Array zero value: %v\n", *arrayPtr)

	// Initialize array elements
	for i := 0; i < 5; i++ {
		arrayPtr[i] = i * i
	}

	fmt.Printf("Array after init: %v\n", *arrayPtr)

	// Allocate 2D array
	matrix := new([3][3]int)
	fmt.Printf("Matrix pointer: %p\n", matrix)
	fmt.Printf("Matrix zero value: %v\n", *matrix)

	// Initialize matrix
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = i*3 + j + 1
		}
	}

	fmt.Printf("Matrix after init: %v\n", *matrix)

	// === SLICE ALLOCATION (INCORRECT USAGE) ===
	fmt.Println("\n5. SLICE ALLOCATION (INCORRECT USAGE):")

	// new() with slice - creates pointer to nil slice
	slicePtr := new([]int)
	fmt.Printf("Slice pointer: %p\n", slicePtr)
	fmt.Printf("Slice value: %v\n", *slicePtr)
	fmt.Printf("Slice is nil: %t\n", *slicePtr == nil)

	// This would panic: (*slicePtr)[0] = 1
	// Need to allocate the slice first
	*slicePtr = make([]int, 5)
	fmt.Printf("After make: %v\n", *slicePtr)

	// Initialize slice
	for i := 0; i < 5; i++ {
		(*slicePtr)[i] = i + 1
	}
	fmt.Printf("Slice after init: %v\n", *slicePtr)

	// === COMPARISON: NEW VS MAKE ===
	fmt.Println("\n6. COMPARISON: NEW VS MAKE:")

	// new() - allocates memory, returns pointer to zero value
	mapPtr := new(map[string]int)
	fmt.Printf("new(map): %p, value: %v, is nil: %t\n", mapPtr, *mapPtr, *mapPtr == nil)

	// make() - allocates and initializes, returns the type itself
	mapVal := make(map[string]int)
	fmt.Printf("make(map): %p, value: %v, is nil: %t\n", &mapVal, mapVal, mapVal == nil)

	// Using the map
	mapVal["key1"] = 10
	mapVal["key2"] = 20
	fmt.Printf("make(map) after use: %v\n", mapVal)

	// To use new() allocated map, need to initialize
	*mapPtr = make(map[string]int)
	(*mapPtr)["key1"] = 100
	(*mapPtr)["key2"] = 200
	fmt.Printf("new(map) after init: %v\n", *mapPtr)

	// === LINKED LIST EXAMPLE ===
	fmt.Println("\n7. LINKED LIST EXAMPLE:")

	// Create linked list using new()
	head := new(Node)
	head.Value = 1

	current := head
	for i := 2; i <= 5; i++ {
		newNode := new(Node)
		newNode.Value = i
		current.Next = newNode
		current = newNode
	}

	// Print linked list
	fmt.Print("Linked list: ")
	current = head
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
	fmt.Println()

	// === MEMORY LAYOUT AND SIZE ===
	fmt.Println("\n8. MEMORY LAYOUT AND SIZE:")

	// Size of different types
	fmt.Printf("Size of int: %d bytes\n", unsafe.Sizeof(int(0)))
	fmt.Printf("Size of *int: %d bytes\n", unsafe.Sizeof(new(int)))
	fmt.Printf("Size of Person: %d bytes\n", unsafe.Sizeof(Person{}))
	fmt.Printf("Size of *Person: %d bytes\n", unsafe.Sizeof(new(Person)))
	fmt.Printf("Size of [5]int: %d bytes\n", unsafe.Sizeof([5]int{}))
	fmt.Printf("Size of *[5]int: %d bytes\n", unsafe.Sizeof(new([5]int)))

	// Demonstrate memory addresses
	p1 := new(int)
	p2 := new(int)
	p3 := new(int)

	fmt.Printf("Pointer 1: %p\n", p1)
	fmt.Printf("Pointer 2: %p\n", p2)
	fmt.Printf("Pointer 3: %p\n", p3)

	// === ZERO VALUES ===
	fmt.Println("\n9. ZERO VALUES:")

	// Different types and their zero values
	intZero := new(int)
	floatZero := new(float64)
	stringZero := new(string)
	boolZero := new(bool)
	sliceZero := new([]int)
	mapZero := new(map[string]int)
	chanZero := new(chan int)
	personZero := new(Person)

	fmt.Printf("int zero: %v\n", *intZero)
	fmt.Printf("float64 zero: %v\n", *floatZero)
	fmt.Printf("string zero: %q\n", *stringZero)
	fmt.Printf("bool zero: %v\n", *boolZero)
	fmt.Printf("[]int zero: %v (nil: %t)\n", *sliceZero, *sliceZero == nil)
	fmt.Printf("map zero: %v (nil: %t)\n", *mapZero, *mapZero == nil)
	fmt.Printf("chan zero: %v (nil: %t)\n", *chanZero, *chanZero == nil)
	fmt.Printf("Person zero: %+v\n", *personZero)

	// === PERFORMANCE CONSIDERATIONS ===
	fmt.Println("\n10. PERFORMANCE CONSIDERATIONS:")

	// Stack allocation (usually faster)
	stackVar := 42
	fmt.Printf("Stack variable: %d (address: %p)\n", stackVar, &stackVar)

	// Heap allocation (with new)
	heapVar := new(int)
	*heapVar = 42
	fmt.Printf("Heap variable: %d (address: %p)\n", *heapVar, heapVar)

	// Bulk allocation
	const size = 1000
	numbers := new([size]int)
	for i := 0; i < size; i++ {
		numbers[i] = i
	}
	fmt.Printf("Bulk allocated array with %d elements\n", size)

	// === COMMON PATTERNS ===
	fmt.Println("\n11. COMMON PATTERNS:")

	// Factory function
	newPerson := func(name string, age int) *Person {
		p := new(Person)
		p.Name = name
		p.Age = age
		return p
	}

	alice := newPerson("Alice", 30)
	bob := newPerson("Bob", 25)

	fmt.Printf("Alice: %+v\n", *alice)
	fmt.Printf("Bob: %+v\n", *bob)

	// Optional initialization
	newPersonWithAddress := func(name string, age int, address *Address) *Person {
		p := new(Person)
		p.Name = name
		p.Age = age
		if address != nil {
			p.Address = address
		}
		return p
	}

	addr := &Address{Street: "123 Oak St", City: "Boston", Country: "USA"}
	charlie := newPersonWithAddress("Charlie", 35, addr)
	fmt.Printf("Charlie: %+v\n", *charlie)

	// === BEST PRACTICES ===
	fmt.Println("\n12. BEST PRACTICES:")

	fmt.Println("✓ Use new() for allocating pointers to zero values")
	fmt.Println("✓ Use make() for slices, maps, and channels")
	fmt.Println("✓ Consider stack allocation for small, short-lived variables")
	fmt.Println("✓ Use new() in factory functions and constructors")
	fmt.Println("✓ Always check for nil before using pointers to reference types")
	fmt.Println("✗ Don't use new() for slices, maps, or channels directly")
	fmt.Println("✗ Don't forget to initialize reference types after new()")

	// === WHEN TO USE NEW() ===
	fmt.Println("\n13. WHEN TO USE NEW():")

	// Good use cases:
	// 1. When you need a pointer to a zero value
	counter := new(int)
	increment := func() {
		*counter++
	}

	increment()
	increment()
	fmt.Printf("Counter: %d\n", *counter)

	// 2. When implementing optional fields
	type Config struct {
		Name    string
		Timeout *int  // Optional field
		Debug   *bool // Optional field
	}

	config := Config{
		Name:    "MyApp",
		Timeout: new(int),
		Debug:   new(bool),
	}

	*config.Timeout = 30
	*config.Debug = true

	fmt.Printf("Config: %+v\n", config)

	// 3. When you need consistent interface
	type Resettable interface {
		Reset()
	}

	// Factory function
	newCounter := func() *Counter {
		return new(Counter) // Zero value is what we want
	}

	myCounter := newCounter()
	myCounter.Increment()
	myCounter.Increment()
	fmt.Printf("Counter value: %d\n", myCounter.Value())

	myCounter.Reset()
	fmt.Printf("Counter after reset: %d\n", myCounter.Value())

	fmt.Println("\n=== END OF NEW ALLOCATION GUIDE ===")
}

/*
RUNNING THE PROGRAM:
go run main.go

LEARNING POINTS:

1. NEW() FUNCTION:
   - Allocates memory for a type
   - Returns pointer to zero-initialized memory
   - Different from make() which is for specific types

2. ZERO VALUES:
   - new() always returns pointer to zero value
   - int: 0, float64: 0.0, string: "", bool: false
   - Reference types (slice, map, chan): nil
   - Structs: all fields set to their zero values

3. MEMORY ALLOCATION:
   - new() allocates on heap
   - Returns pointer for consistency
   - Garbage collected when no longer referenced

4. COMMON PATTERNS:
   - Factory functions
   - Optional fields in structs
   - Consistent interfaces
   - Pointer initialization

5. BEST PRACTICES:
   - Use new() for zero-value initialization
   - Use make() for slices, maps, channels
   - Consider stack allocation for performance
   - Always check nil for reference types

6. WHEN NOT TO USE:
   - Don't use new() for slices/maps/channels directly
   - Consider if stack allocation is better
   - Don't use if you need non-zero initialization

This demonstrates comprehensive new() allocation patterns in Go
for memory management and pointer initialization.
*/
