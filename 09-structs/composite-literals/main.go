package main

import (
	"fmt"
	"time"
)

// === GO COMPOSITE LITERALS COMPREHENSIVE GUIDE ===

/*
COMPOSITE LITERALS PHILOSOPHY:
- Composite literals create values for structs, arrays, slices, and maps
- Provide a concise way to initialize complex data structures
- Can be used for both literal values and pointer allocation
- Support both keyed and non-keyed initialization

COMPARISON WITH JAVASCRIPT:
// JavaScript - Object literals
const person = {
  name: "Alice",
  age: 30,
  address: {
    street: "123 Main St",
    city: "Boston"
  }
};

// JavaScript - Array literals
const numbers = [1, 2, 3, 4, 5];
const matrix = [[1, 2], [3, 4]];

// Go - Composite literals
person := Person{
  Name: "Alice",
  Age: 30,
  Address: Address{
    Street: "123 Main St",
    City: "Boston",
  },
}

numbers := []int{1, 2, 3, 4, 5}
matrix := [][]int{{1, 2}, {3, 4}}
*/

// === STRUCT DEFINITIONS ===

type Person struct {
	Name     string
	Age      int
	Email    string
	Address  Address
	Hobbies  []string
	Settings map[string]interface{}
}

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
	Country string
}

type Employee struct {
	Person
	ID       int
	Position string
	Salary   float64
	HireDate time.Time
}

type Point struct {
	X, Y float64
}

type Rectangle struct {
	TopLeft     Point
	BottomRight Point
}

type Company struct {
	Name      string
	Employees []Employee
	Locations []Address
}

// Shape interface
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	width := r.BottomRight.X - r.TopLeft.X
	height := r.BottomRight.Y - r.TopLeft.Y
	return width * height
}

func main() {
	fmt.Println("=== GO COMPOSITE LITERALS COMPREHENSIVE GUIDE ===")

	// === STRUCT LITERALS ===
	fmt.Println("\n1. STRUCT LITERALS:")

	// Basic struct literal (positional)
	p1 := Person{"Alice", 30, "alice@example.com", Address{}, []string{}, map[string]interface{}{}}
	fmt.Printf("Person (positional): %+v\n", p1)

	// Struct literal with field names (recommended)
	p2 := Person{
		Name:     "Bob",
		Age:      25,
		Email:    "bob@example.com",
		Address:  Address{},
		Hobbies:  []string{},
		Settings: map[string]interface{}{},
	}
	fmt.Printf("Person (named fields): %+v\n", p2)

	// Partial initialization (zero values for omitted fields)
	p3 := Person{
		Name:  "Charlie",
		Age:   35,
		Email: "charlie@example.com",
	}
	fmt.Printf("Person (partial): %+v\n", p3)

	// === NESTED STRUCT LITERALS ===
	fmt.Println("\n2. NESTED STRUCT LITERALS:")

	// Nested struct initialization
	p4 := Person{
		Name:  "Diana",
		Age:   28,
		Email: "diana@example.com",
		Address: Address{
			Street:  "456 Oak Ave",
			City:    "Seattle",
			State:   "WA",
			ZipCode: "98101",
			Country: "USA",
		},
		Hobbies: []string{"reading", "hiking", "photography"},
		Settings: map[string]interface{}{
			"theme":         "dark",
			"notifications": true,
			"language":      "en",
		},
	}
	fmt.Printf("Person (nested): %+v\n", p4)

	// Multi-level nesting
	company := Company{
		Name: "Tech Corp",
		Employees: []Employee{
			{
				Person: Person{
					Name:  "John",
					Age:   32,
					Email: "john@techcorp.com",
					Address: Address{
						Street:  "789 Tech Blvd",
						City:    "San Francisco",
						State:   "CA",
						ZipCode: "94105",
						Country: "USA",
					},
				},
				ID:       1001,
				Position: "Senior Developer",
				Salary:   120000.0,
				HireDate: time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC),
			},
			{
				Person: Person{
					Name:  "Sarah",
					Age:   29,
					Email: "sarah@techcorp.com",
				},
				ID:       1002,
				Position: "Product Manager",
				Salary:   115000.0,
				HireDate: time.Date(2021, 3, 10, 0, 0, 0, 0, time.UTC),
			},
		},
		Locations: []Address{
			{
				Street:  "123 Innovation Way",
				City:    "San Francisco",
				State:   "CA",
				ZipCode: "94105",
				Country: "USA",
			},
			{
				Street:  "456 Remote Ave",
				City:    "Austin",
				State:   "TX",
				ZipCode: "78701",
				Country: "USA",
			},
		},
	}
	fmt.Printf("Company: %+v\n", company)

	// === ARRAY LITERALS ===
	fmt.Println("\n3. ARRAY LITERALS:")

	// Basic array literal
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array: %v\n", numbers)

	// Array with partial initialization
	partial := [5]int{1, 2} // Rest are zero values
	fmt.Printf("Partial array: %v\n", partial)

	// Array with index-based initialization
	indexed := [5]int{2: 100, 4: 200} // Index 2 = 100, Index 4 = 200
	fmt.Printf("Indexed array: %v\n", indexed)

	// Array with ellipsis (compiler counts elements)
	ellipsis := [...]int{10, 20, 30, 40, 50}
	fmt.Printf("Ellipsis array: %v, length: %d\n", ellipsis, len(ellipsis))

	// Multi-dimensional array
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("2D array: %v\n", matrix)

	// === SLICE LITERALS ===
	fmt.Println("\n4. SLICE LITERALS:")

	// Basic slice literal
	fruits := []string{"apple", "banana", "orange"}
	fmt.Printf("Slice: %v, len: %d, cap: %d\n", fruits, len(fruits), cap(fruits))

	// Empty slice
	empty := []int{}
	fmt.Printf("Empty slice: %v, len: %d, cap: %d\n", empty, len(empty), cap(empty))

	// Slice with different types
	mixed := []interface{}{1, "hello", 3.14, true}
	fmt.Printf("Mixed slice: %v\n", mixed)

	// Slice of structs
	points := []Point{
		{X: 0, Y: 0},
		{X: 1, Y: 1},
		{X: 2, Y: 4},
		{X: 3, Y: 9},
	}
	fmt.Printf("Points slice: %v\n", points)

	// 2D slice
	matrix2D := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("2D slice: %v\n", matrix2D)

	// === MAP LITERALS ===
	fmt.Println("\n5. MAP LITERALS:")

	// Basic map literal
	ages := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}
	fmt.Printf("Ages map: %v\n", ages)

	// Empty map
	emptyMap := map[string]int{}
	fmt.Printf("Empty map: %v\n", emptyMap)

	// Map with struct values
	people := map[string]Person{
		"alice": {
			Name:  "Alice Johnson",
			Age:   30,
			Email: "alice@example.com",
		},
		"bob": {
			Name:  "Bob Smith",
			Age:   25,
			Email: "bob@example.com",
		},
	}
	fmt.Printf("People map: %v\n", people)

	// Map with slice values
	groups := map[string][]string{
		"developers": {"Alice", "Bob", "Charlie"},
		"designers":  {"Diana", "Eve"},
		"managers":   {"Frank", "Grace"},
	}
	fmt.Printf("Groups map: %v\n", groups)

	// Nested map
	config := map[string]map[string]interface{}{
		"database": {
			"host":     "localhost",
			"port":     5432,
			"username": "admin",
			"ssl":      true,
		},
		"cache": {
			"type":    "redis",
			"ttl":     300,
			"enabled": true,
		},
	}
	fmt.Printf("Config map: %v\n", config)

	// === POINTER TO COMPOSITE LITERALS ===
	fmt.Println("\n6. POINTER TO COMPOSITE LITERALS:")

	// Pointer to struct literal
	personPtr := &Person{
		Name:  "Helen",
		Age:   27,
		Email: "helen@example.com",
	}
	fmt.Printf("Person pointer: %p, value: %+v\n", personPtr, *personPtr)

	// Pointer to array literal
	arrayPtr := &[3]int{1, 2, 3}
	fmt.Printf("Array pointer: %p, value: %v\n", arrayPtr, *arrayPtr)

	// Pointer to slice literal
	slicePtr := &[]string{"go", "rust", "python"}
	fmt.Printf("Slice pointer: %p, value: %v\n", slicePtr, *slicePtr)

	// Pointer to map literal
	mapPtr := &map[string]int{"a": 1, "b": 2}
	fmt.Printf("Map pointer: %p, value: %v\n", mapPtr, *mapPtr)

	// === FUNCTION PARAMETERS ===
	fmt.Println("\n7. FUNCTION PARAMETERS:")

	// Function taking struct literal
	printPerson := func(p Person) {
		fmt.Printf("Person: %s, Age: %d\n", p.Name, p.Age)
	}

	// Pass struct literal directly
	printPerson(Person{Name: "Ivan", Age: 33})

	// Function taking slice literal
	sum := func(numbers []int) int {
		total := 0
		for _, n := range numbers {
			total += n
		}
		return total
	}

	// Pass slice literal directly
	result := sum([]int{1, 2, 3, 4, 5})
	fmt.Printf("Sum result: %d\n", result)

	// Function taking map literal
	printMap := func(m map[string]int) {
		for k, v := range m {
			fmt.Printf("%s: %d\n", k, v)
		}
	}

	// Pass map literal directly
	printMap(map[string]int{"x": 10, "y": 20})

	// === RETURN VALUES ===
	fmt.Println("\n8. RETURN VALUES:")

	// Function returning struct literal
	createPerson := func(name string, age int) Person {
		return Person{
			Name:  name,
			Age:   age,
			Email: fmt.Sprintf("%s@example.com", name),
		}
	}

	newPerson := createPerson("Jack", 40)
	fmt.Printf("Created person: %+v\n", newPerson)

	// Function returning pointer to struct literal
	createPersonPtr := func(name string, age int) *Person {
		return &Person{
			Name:  name,
			Age:   age,
			Email: fmt.Sprintf("%s@example.com", name),
		}
	}

	newPersonPtr := createPersonPtr("Kate", 35)
	fmt.Printf("Created person pointer: %+v\n", *newPersonPtr)

	// Function returning slice literal
	createNumbers := func(count int) []int {
		return []int{1, 2, 3, 4, 5}[:count]
	}

	nums := createNumbers(3)
	fmt.Printf("Created numbers: %v\n", nums)

	// === INTERFACE LITERALS ===
	fmt.Println("\n9. INTERFACE LITERALS:")

	// Slice of interface literals
	shapes := []Shape{
		Circle{Radius: 5.0},
		Rectangle{
			TopLeft:     Point{X: 0, Y: 0},
			BottomRight: Point{X: 10, Y: 5},
		},
	}

	for i, shape := range shapes {
		fmt.Printf("Shape %d area: %.2f\n", i, shape.Area())
	}

	// === COMPLEX NESTED STRUCTURES ===
	fmt.Println("\n10. COMPLEX NESTED STRUCTURES:")

	// Complex data structure
	data := map[string]interface{}{
		"users": []map[string]interface{}{
			{
				"id":   1,
				"name": "Alice",
				"profile": map[string]interface{}{
					"age":    30,
					"skills": []string{"Go", "Python", "JavaScript"},
				},
			},
			{
				"id":   2,
				"name": "Bob",
				"profile": map[string]interface{}{
					"age":    25,
					"skills": []string{"Java", "C++", "Rust"},
				},
			},
		},
		"settings": map[string]interface{}{
			"theme":    "dark",
			"language": "en",
			"features": []string{"notifications", "analytics"},
		},
	}

	fmt.Printf("Complex data: %v\n", data)

	// === TYPE INFERENCE ===
	fmt.Println("\n11. TYPE INFERENCE:")

	// Type inference in composite literals
	var point Point = Point{1.5, 2.5} // Explicit type
	point2 := Point{3.0, 4.0}         // Inferred type

	fmt.Printf("Point: %v\n", point)
	fmt.Printf("Point2: %v\n", point2)

	// Type inference in slices
	var numbers1 []int = []int{1, 2, 3} // Explicit type
	numbers2 := []int{4, 5, 6}          // Inferred type

	fmt.Printf("Numbers1: %v\n", numbers1)
	fmt.Printf("Numbers2: %v\n", numbers2)

	// === ZERO VALUES AND OMITTED FIELDS ===
	fmt.Println("\n12. ZERO VALUES AND OMITTED FIELDS:")

	// Struct with omitted fields
	partialPerson := Person{
		Name: "Linda",
		Age:  28,
		// Email, Address, Hobbies, Settings will be zero values
	}

	fmt.Printf("Partial person: %+v\n", partialPerson)
	fmt.Printf("Email is empty: %t\n", partialPerson.Email == "")
	fmt.Printf("Hobbies is nil: %t\n", partialPerson.Hobbies == nil)
	fmt.Printf("Settings is nil: %t\n", partialPerson.Settings == nil)

	// Initialize with zero values explicitly
	explicitZero := Person{
		Name:     "Mike",
		Age:      0,
		Email:    "",
		Address:  Address{},
		Hobbies:  []string{},
		Settings: map[string]interface{}{},
	}

	fmt.Printf("Explicit zero: %+v\n", explicitZero)

	// === BEST PRACTICES ===
	fmt.Println("\n13. BEST PRACTICES:")

	fmt.Println("✓ Use field names for struct literals (readability)")
	fmt.Println("✓ Use composite literals for initialization")
	fmt.Println("✓ Prefer slice literals over append for known values")
	fmt.Println("✓ Use map literals for initial key-value pairs")
	fmt.Println("✓ Use pointer to composite literals to avoid copying")
	fmt.Println("✗ Don't mix positional and named fields")
	fmt.Println("✗ Don't rely on field order for positional initialization")

	// === REAL-WORLD EXAMPLES ===
	fmt.Println("\n14. REAL-WORLD EXAMPLES:")

	// Configuration structure
	appConfig := struct {
		Server struct {
			Host string
			Port int
		}
		Database struct {
			URL      string
			MaxConns int
		}
		Features []string
	}{
		Server: struct {
			Host string
			Port int
		}{
			Host: "localhost",
			Port: 8080,
		},
		Database: struct {
			URL      string
			MaxConns int
		}{
			URL:      "postgres://localhost:5432/mydb",
			MaxConns: 10,
		},
		Features: []string{"auth", "logging", "metrics"},
	}

	fmt.Printf("App config: %+v\n", appConfig)

	// API response structure
	apiResponse := struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data"`
		Error  string      `json:"error,omitempty"`
	}{
		Status: 200,
		Data: map[string]interface{}{
			"users": []map[string]interface{}{
				{"id": 1, "name": "Alice"},
				{"id": 2, "name": "Bob"},
			},
		},
	}

	fmt.Printf("API response: %+v\n", apiResponse)

	// Test data
	testCases := []struct {
		input    string
		expected int
	}{
		{"hello", 5},
		{"world", 5},
		{"go", 2},
		{"", 0},
	}

	for _, tc := range testCases {
		actual := len(tc.input)
		fmt.Printf("Test: %s -> %d (expected: %d)\n", tc.input, actual, tc.expected)
	}

	fmt.Println("\n=== END OF COMPOSITE LITERALS GUIDE ===")
}

/*
RUNNING THE PROGRAM:
go run main.go

LEARNING POINTS:

1. COMPOSITE LITERALS:
   - Concise way to create and initialize complex data structures
   - Support structs, arrays, slices, and maps
   - Can be used for both values and pointers

2. STRUCT LITERALS:
   - Positional: Person{"Alice", 30, "alice@example.com"}
   - Named fields: Person{Name: "Alice", Age: 30}
   - Partial initialization uses zero values

3. ARRAY LITERALS:
   - Fixed size: [5]int{1, 2, 3, 4, 5}
   - Partial: [5]int{1, 2} (rest are zero)
   - Indexed: [5]int{2: 100, 4: 200}
   - Ellipsis: [...]int{1, 2, 3}

4. SLICE LITERALS:
   - Dynamic: []int{1, 2, 3, 4, 5}
   - Multi-dimensional: [][]int{{1, 2}, {3, 4}}
   - Empty: []int{}

5. MAP LITERALS:
   - Basic: map[string]int{"a": 1, "b": 2}
   - Empty: map[string]int{}
   - Complex values: map[string][]string{"key": {"val1", "val2"}}

6. POINTERS TO LITERALS:
   - &Person{Name: "Alice"} creates pointer to struct
   - Useful for avoiding copies of large structures

7. BEST PRACTICES:
   - Always use field names for struct literals
   - Use composite literals for initialization
   - Consider pointer literals for large structures
   - Prefer literals over multiple assignments

This demonstrates comprehensive composite literal usage in Go
for efficient data structure initialization and manipulation.
*/
