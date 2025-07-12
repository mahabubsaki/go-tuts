package main

import (
	"fmt"
	"unsafe"
)

// === POINTERS IN GO ===

// Struct for demonstration
type Person struct {
	Name string
	Age  int
}

// Method with pointer receiver
func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge
}

// Method with value receiver
func (p Person) GetInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

// Function that modifies value through pointer
func modifyValue(ptr *int) {
	*ptr = 100
}

// Function that doesn't modify original (value copy)
func modifyValueCopy(val int) {
	val = 200
}

// Function that returns pointer
func createPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

// Function that swaps two values using pointers
func swap(a, b *int) {
	*a, *b = *b, *a
}

// Linked list node
type Node struct {
	Value int
	Next  *Node
}

// Linked list operations
func (n *Node) Append(value int) {
	current := n
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node{Value: value}
}

func (n *Node) Print() {
	current := n
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
	fmt.Println()
}

// Binary tree node
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Insert(value int) {
	if value < t.Value {
		if t.Left == nil {
			t.Left = &TreeNode{Value: value}
		} else {
			t.Left.Insert(value)
		}
	} else {
		if t.Right == nil {
			t.Right = &TreeNode{Value: value}
		} else {
			t.Right.Insert(value)
		}
	}
}

func (t *TreeNode) InorderTraversal() {
	if t != nil {
		t.Left.InorderTraversal()
		fmt.Printf("%d ", t.Value)
		t.Right.InorderTraversal()
	}
}

// Pointer to function
type Operation func(int, int) int

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func calculate(op Operation, a, b int) int {
	return op(a, b)
}

func main() {
	fmt.Println("=== GO POINTERS COMPREHENSIVE GUIDE ===")
	
	// === BASIC POINTER CONCEPTS ===
	fmt.Println("\n--- BASIC POINTER CONCEPTS ---")
	
	// Variable declaration and pointer
	var x int = 42
	var ptr *int = &x
	
	fmt.Printf("Variable x: %d\n", x)
	fmt.Printf("Address of x: %p\n", &x)
	fmt.Printf("Pointer ptr: %p\n", ptr)
	fmt.Printf("Value at ptr: %d\n", *ptr)
	
	// Modify value through pointer
	*ptr = 100
	fmt.Printf("After modifying through pointer:\n")
	fmt.Printf("Variable x: %d\n", x)
	fmt.Printf("Value at ptr: %d\n", *ptr)
	
	/*
	JavaScript comparison:
	// JavaScript doesn't have pointers, but has references
	let obj = {value: 42};
	let ref = obj; // Reference to the same object
	
	console.log(obj.value); // 42
	ref.value = 100;
	console.log(obj.value); // 100 (modified through reference)
	
	// For primitives, JavaScript passes by value
	let x = 42;
	let y = x; // Copy of value
	y = 100;
	console.log(x); // 42 (unchanged)
	*/
	
	// === POINTER DECLARATION AND INITIALIZATION ===
	fmt.Println("\n--- POINTER DECLARATION AND INITIALIZATION ---")
	
	// Different ways to declare pointers
	var p1 *int                    // nil pointer
	var p2 *int = &x              // pointer to existing variable
	p3 := &x                      // short declaration
	p4 := new(int)                // pointer to zero value
	*p4 = 50
	
	fmt.Printf("p1 (nil): %v\n", p1)
	fmt.Printf("p2: %p, value: %d\n", p2, *p2)
	fmt.Printf("p3: %p, value: %d\n", p3, *p3)
	fmt.Printf("p4: %p, value: %d\n", p4, *p4)
	
	// Check for nil pointer
	if p1 == nil {
		fmt.Println("p1 is nil")
	}
	
	// === POINTER ARITHMETIC (LIMITED IN GO) ===
	fmt.Println("\n--- POINTER ARITHMETIC ---")
	
	// Go doesn't support pointer arithmetic like C/C++
	// But we can demonstrate memory addresses
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array: %v\n", arr)
	
	for i := range arr {
		fmt.Printf("arr[%d] = %d, address: %p\n", i, arr[i], &arr[i])
	}
	
	// Calculate address differences
	fmt.Printf("Address difference between arr[1] and arr[0]: %d bytes\n", 
		uintptr(unsafe.Pointer(&arr[1])) - uintptr(unsafe.Pointer(&arr[0])))
	
	// === POINTERS AND FUNCTIONS ===
	fmt.Println("\n--- POINTERS AND FUNCTIONS ---")
	
	value := 50
	fmt.Printf("Original value: %d\n", value)
	
	// Pass by value (doesn't modify original)
	modifyValueCopy(value)
	fmt.Printf("After modifyValueCopy: %d\n", value)
	
	// Pass by pointer (modifies original)
	modifyValue(&value)
	fmt.Printf("After modifyValue: %d\n", value)
	
	// Swap function
	a, b := 10, 20
	fmt.Printf("Before swap: a=%d, b=%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swap: a=%d, b=%d\n", a, b)
	
	// === POINTERS AND STRUCTS ===
	fmt.Println("\n--- POINTERS AND STRUCTS ---")
	
	// Create struct
	person1 := Person{Name: "Alice", Age: 30}
	fmt.Printf("Person1: %+v\n", person1)
	
	// Pointer to struct
	personPtr := &person1
	fmt.Printf("PersonPtr: %+v\n", personPtr)
	
	// Access fields through pointer (automatic dereferencing)
	fmt.Printf("Name through pointer: %s\n", personPtr.Name)
	personPtr.Age = 31
	fmt.Printf("Modified age: %d\n", person1.Age)
	
	// Method calls with pointer receiver
	person1.UpdateAge(32)
	fmt.Printf("After UpdateAge: %+v\n", person1)
	
	// Create struct using pointer-returning function
	person2 := createPerson("Bob", 25)
	fmt.Printf("Person2: %+v\n", person2)
	
	// === POINTER TO POINTER ===
	fmt.Println("\n--- POINTER TO POINTER ---")
	
	var val int = 42
	var ptr1 *int = &val
	var ptr2 **int = &ptr1
	
	fmt.Printf("Value: %d\n", val)
	fmt.Printf("Pointer to value: %p\n", ptr1)
	fmt.Printf("Pointer to pointer: %p\n", ptr2)
	fmt.Printf("Value through ptr1: %d\n", *ptr1)
	fmt.Printf("Value through ptr2: %d\n", **ptr2)
	
	// Modify through pointer to pointer
	**ptr2 = 100
	fmt.Printf("After modifying through ptr2: %d\n", val)
	
	// === SLICES AND POINTERS ===
	fmt.Println("\n--- SLICES AND POINTERS ---")
	
	// Slice elements are accessible by pointer
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original slice: %v\n", slice)
	
	// Get pointer to slice element
	elementPtr := &slice[2]
	fmt.Printf("Element at index 2: %d\n", *elementPtr)
	
	// Modify through pointer
	*elementPtr = 100
	fmt.Printf("After modifying through pointer: %v\n", slice)
	
	// Slice of pointers
	var ptrSlice []*int
	for i := range slice {
		ptrSlice = append(ptrSlice, &slice[i])
	}
	
	fmt.Printf("Pointer slice addresses: ")
	for _, ptr := range ptrSlice {
		fmt.Printf("%p ", ptr)
	}
	fmt.Println()
	
	// === MAPS AND POINTERS ===
	fmt.Println("\n--- MAPS AND POINTERS ---")
	
	// Map of pointers to structs
	people := map[string]*Person{
		"alice": {Name: "Alice", Age: 30},
		"bob":   {Name: "Bob", Age: 25},
	}
	
	fmt.Printf("People map:\n")
	for key, person := range people {
		fmt.Printf("  %s: %+v\n", key, person)
	}
	
	// Modify through map pointer
	people["alice"].Age = 31
	fmt.Printf("After modifying Alice's age: %+v\n", people["alice"])
	
	// === LINKED LIST IMPLEMENTATION ===
	fmt.Println("\n--- LINKED LIST IMPLEMENTATION ---")
	
	// Create linked list
	head := &Node{Value: 1}
	head.Append(2)
	head.Append(3)
	head.Append(4)
	
	fmt.Print("Linked list: ")
	head.Print()
	
	// Traverse manually
	fmt.Print("Manual traversal: ")
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
	fmt.Println()
	
	// === BINARY TREE IMPLEMENTATION ===
	fmt.Println("\n--- BINARY TREE IMPLEMENTATION ---")
	
	// Create binary tree
	root := &TreeNode{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)
	
	fmt.Print("Binary tree (inorder): ")
	root.InorderTraversal()
	fmt.Println()
	
	// === FUNCTION POINTERS ===
	fmt.Println("\n--- FUNCTION POINTERS ---")
	
	// Function variables
	var op Operation
	
	op = add
	result := calculate(op, 10, 5)
	fmt.Printf("Add operation: %d\n", result)
	
	op = multiply
	result = calculate(op, 10, 5)
	fmt.Printf("Multiply operation: %d\n", result)
	
	// Anonymous function
	op = func(a, b int) int {
		return a - b
	}
	result = calculate(op, 10, 5)
	fmt.Printf("Subtract operation: %d\n", result)
	
	// === POINTER RECEIVERS VS VALUE RECEIVERS ===
	fmt.Println("\n--- POINTER VS VALUE RECEIVERS ---")
	
	person := Person{Name: "Charlie", Age: 40}
	
	// Value receiver - doesn't modify original
	info := person.GetInfo()
	fmt.Printf("Info: %s\n", info)
	
	// Pointer receiver - modifies original
	person.UpdateAge(41)
	fmt.Printf("After UpdateAge: %+v\n", person)
	
	// Method calls work with both values and pointers
	personPtr = &person
	personPtr.UpdateAge(42)
	fmt.Printf("After UpdateAge via pointer: %+v\n", person)
	
	// === MEMORY MANAGEMENT ===
	fmt.Println("\n--- MEMORY MANAGEMENT ---")
	
	// Go has garbage collection, but understanding pointers helps
	// Local variables
	func() {
		localVar := 100
		localPtr := &localVar
		fmt.Printf("Local variable: %d, address: %p\n", localVar, localPtr)
	}()
	
	// Dynamic allocation
	dynamicPtr := new(int)
	*dynamicPtr = 200
	fmt.Printf("Dynamic variable: %d, address: %p\n", *dynamicPtr, dynamicPtr)
	
	// === COMMON POINTER PATTERNS ===
	fmt.Println("\n--- COMMON POINTER PATTERNS ---")
	
	// 1. Optional values (like nullable types)
	var optionalInt *int
	if optionalInt == nil {
		fmt.Println("Optional int is nil")
	}
	
	value = 42
	optionalInt = &value
	if optionalInt != nil {
		fmt.Printf("Optional int has value: %d\n", *optionalInt)
	}
	
	// 2. Self-referential structures
	type ListNode struct {
		Value int
		Next  *ListNode
	}
	
	node1 := &ListNode{Value: 1}
	node2 := &ListNode{Value: 2}
	node1.Next = node2
	
	fmt.Printf("Node1 -> Node2: %d -> %d\n", node1.Value, node1.Next.Value)
	
	// 3. Circular references (be careful with these)
	type CircularNode struct {
		Value int
		Next  *CircularNode
	}
	
	c1 := &CircularNode{Value: 1}
	c2 := &CircularNode{Value: 2}
	c1.Next = c2
	c2.Next = c1 // Creates cycle
	
	fmt.Printf("Circular: %d -> %d -> %d\n", c1.Value, c1.Next.Value, c1.Next.Next.Value)
	
	// === POINTER SAFETY ===
	fmt.Println("\n--- POINTER SAFETY ---")
	
	// Go prevents many common pointer errors
	// 1. No pointer arithmetic beyond array bounds
	// 2. No manual memory management
	// 3. Garbage collection prevents memory leaks
	// 4. No dangling pointers in safe code
	
	// Safe pointer usage
	safePtr := &Person{Name: "Safe", Age: 25}
	if safePtr != nil {
		fmt.Printf("Safe pointer: %+v\n", safePtr)
	}
	
	// === PERFORMANCE CONSIDERATIONS ===
	fmt.Println("\n--- PERFORMANCE CONSIDERATIONS ---")
	
	// Small structs - value is often better
	type SmallStruct struct {
		A, B int
	}
	
	// Large structs - pointer is often better
	type LargeStruct struct {
		Data [1000]int
	}
	
	small := SmallStruct{A: 1, B: 2}
	large := LargeStruct{}
	
	fmt.Printf("Small struct size: %d bytes\n", unsafe.Sizeof(small))
	fmt.Printf("Large struct size: %d bytes\n", unsafe.Sizeof(large))
	fmt.Printf("Pointer size: %d bytes\n", unsafe.Sizeof(&large))
	
	// === POINTER BEST PRACTICES ===
	fmt.Println("\n--- POINTER BEST PRACTICES ---")
	fmt.Println("1. Always check for nil before dereferencing")
	fmt.Println("2. Use pointers for large structs to avoid copying")
	fmt.Println("3. Use pointer receivers for methods that modify the receiver")
	fmt.Println("4. Use value receivers for methods that don't modify the receiver")
	fmt.Println("5. Prefer values over pointers for small data types")
	fmt.Println("6. Use pointers for optional values (nil-able)")
	fmt.Println("7. Be careful with circular references")
	fmt.Println("8. Use new() for zero-value initialization")
	fmt.Println("9. Use &T{} for composite literal initialization")
	fmt.Println("10. Let the garbage collector handle memory management")
	
	// === COMMON MISTAKES ===
	fmt.Println("\n--- COMMON MISTAKES TO AVOID ---")
	
	// 1. Dereferencing nil pointer (would panic)
	// var nilPtr *int
	// fmt.Println(*nilPtr) // This would panic!
	
	// 2. Taking address of map element (not allowed)
	// m := map[string]int{"key": 42}
	// ptr := &m["key"] // This would cause compile error!
	
	// 3. Correct way to work with map elements
	m := map[string]*int{"key": &value}
	if ptr, exists := m["key"]; exists {
		fmt.Printf("Map element through pointer: %d\n", *ptr)
	}
	
	fmt.Println("\nPointers provide powerful memory management capabilities in Go!")
}
