package main

import (
	"fmt"
	"time"
	"unsafe"
)

// === STRUCTS IN GO ===

// 1. Basic struct definition
type Person struct {
	Name string
	Age  int
}

// 2. Struct with different field types
type Student struct {
	ID         int
	Name       string
	Age        int
	Grade      float64
	IsActive   bool
	Subjects   []string
	Address    Address
	EnrollDate time.Time
}

// 3. Nested struct
type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
	Country string
}

// 4. Struct with methods
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// 5. Struct with embedded fields (composition)
type Animal struct {
	Name    string
	Species string
	Age     int
}

type Dog struct {
	Animal
	Breed string
}

type Cat struct {
	Animal
	IndoorOnly bool
}

// 6. Struct with tags
type User struct {
	ID       int    `json:"id" db:"user_id"`
	Name     string `json:"name" db:"full_name"`
	Email    string `json:"email" db:"email_address"`
	Password string `json:"-" db:"password_hash"`
}

// 7. Anonymous struct
func demonstrateAnonymousStruct() {
	// Anonymous struct declaration and initialization
	person := struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  30,
	}

	fmt.Printf("Anonymous struct: %+v\n", person)
}

// 8. Struct with function fields
type Calculator struct {
	Name      string
	Operation func(float64, float64) float64
}

// 9. Struct with pointer fields
type Node struct {
	Value int
	Next  *Node
}

// 10. Struct with interface fields
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

type Container struct {
	Name  string
	Shape Shape
}

func main() {
	fmt.Println("=== GO STRUCTS COMPREHENSIVE GUIDE ===")

	// === BASIC STRUCT USAGE ===
	fmt.Println("\n--- BASIC STRUCT USAGE ---")

	// Different ways to create structs
	var p1 Person
	fmt.Printf("Zero-value struct: %+v\n", p1)

	p2 := Person{"Alice", 30}
	fmt.Printf("Positional initialization: %+v\n", p2)

	p3 := Person{Name: "Bob", Age: 25}
	fmt.Printf("Named initialization: %+v\n", p3)

	p4 := Person{Name: "Charlie"} // Age gets zero value
	fmt.Printf("Partial initialization: %+v\n", p4)

	/*
		JavaScript comparison:
		// JavaScript objects
		const p1 = {}; // Empty object
		const p2 = {name: "Alice", age: 30};
		const p3 = {name: "Bob", age: 25};
		const p4 = {name: "Charlie"}; // age is undefined

		// JavaScript classes
		class Person {
			constructor(name, age) {
				this.name = name;
				this.age = age;
			}
		}

		const person = new Person("Alice", 30);
	*/

	// === STRUCT FIELD ACCESS ===
	fmt.Println("\n--- STRUCT FIELD ACCESS ---")

	person := Person{Name: "David", Age: 35}
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)

	// Modify fields
	person.Age = 36
	fmt.Printf("Updated age: %d\n", person.Age)

	// Struct pointer
	personPtr := &person
	fmt.Printf("Via pointer: %s, %d\n", personPtr.Name, personPtr.Age)

	// Go automatically dereferences
	personPtr.Age = 37
	fmt.Printf("After pointer modification: %d\n", person.Age)

	// === COMPLEX STRUCT ===
	fmt.Println("\n--- COMPLEX STRUCT ---")

	student := Student{
		ID:       12345,
		Name:     "Emma Watson",
		Age:      20,
		Grade:    3.8,
		IsActive: true,
		Subjects: []string{"Math", "Physics", "Computer Science"},
		Address: Address{
			Street:  "123 Main St",
			City:    "Boston",
			State:   "MA",
			ZipCode: "02101",
			Country: "USA",
		},
		EnrollDate: time.Date(2020, 9, 1, 0, 0, 0, 0, time.UTC),
	}

	fmt.Printf("Student: %+v\n", student)
	fmt.Printf("Address: %+v\n", student.Address)
	fmt.Printf("Subjects: %v\n", student.Subjects)
	fmt.Printf("Enroll Date: %s\n", student.EnrollDate.Format("2006-01-02"))

	// === STRUCT METHODS ===
	fmt.Println("\n--- STRUCT METHODS ---")

	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())

	rect.Scale(2.0)
	fmt.Printf("After scaling: %+v\n", rect)
	fmt.Printf("New area: %.2f\n", rect.Area())

	// === STRUCT COMPOSITION ===
	fmt.Println("\n--- STRUCT COMPOSITION ---")

	dog := Dog{
		Animal: Animal{
			Name:    "Buddy",
			Species: "Canis lupus",
			Age:     3,
		},
		Breed: "Golden Retriever",
	}

	cat := Cat{
		Animal: Animal{
			Name:    "Whiskers",
			Species: "Felis catus",
			Age:     5,
		},
		IndoorOnly: true,
	}

	fmt.Printf("Dog: %+v\n", dog)
	fmt.Printf("Cat: %+v\n", cat)

	// Access embedded fields directly
	fmt.Printf("Dog's name: %s\n", dog.Name)
	fmt.Printf("Cat's species: %s\n", cat.Species)

	// Access embedded struct
	fmt.Printf("Dog's animal: %+v\n", dog.Animal)

	// === STRUCT COPYING ===
	fmt.Println("\n--- STRUCT COPYING ---")

	original := Person{Name: "John", Age: 30}
	copy := original // Structs are value types

	fmt.Printf("Original: %+v\n", original)
	fmt.Printf("Copy: %+v\n", copy)

	copy.Age = 31
	fmt.Printf("After modifying copy:\n")
	fmt.Printf("Original: %+v\n", original)
	fmt.Printf("Copy: %+v\n", copy)

	// === STRUCT COMPARISON ===
	fmt.Println("\n--- STRUCT COMPARISON ---")

	p1 = Person{Name: "Alice", Age: 25}
	p2 = Person{Name: "Alice", Age: 25}
	p3 = Person{Name: "Bob", Age: 25}

	fmt.Printf("p1 == p2: %t\n", p1 == p2)
	fmt.Printf("p1 == p3: %t\n", p1 == p3)

	// Structs with slices/maps/functions are not comparable
	// student1 := Student{Name: "Alice", Subjects: []string{"Math"}}
	// student2 := Student{Name: "Alice", Subjects: []string{"Math"}}
	// fmt.Printf("student1 == student2: %t\n", student1 == student2) // Error!

	// === ANONYMOUS STRUCT ===
	fmt.Println("\n--- ANONYMOUS STRUCT ---")

	demonstrateAnonymousStruct()

	// Anonymous struct in slice
	configs := []struct {
		Name  string
		Value int
	}{
		{"MaxConnections", 100},
		{"Timeout", 30},
		{"RetryCount", 3},
	}

	fmt.Println("Config values:")
	for _, config := range configs {
		fmt.Printf("  %s: %d\n", config.Name, config.Value)
	}

	// === STRUCT WITH FUNCTION FIELDS ===
	fmt.Println("\n--- STRUCT WITH FUNCTION FIELDS ---")

	add := Calculator{
		Name: "Adder",
		Operation: func(a, b float64) float64 {
			return a + b
		},
	}

	multiply := Calculator{
		Name: "Multiplier",
		Operation: func(a, b float64) float64 {
			return a * b
		},
	}

	fmt.Printf("%s: 5 + 3 = %.2f\n", add.Name, add.Operation(5, 3))
	fmt.Printf("%s: 5 * 3 = %.2f\n", multiply.Name, multiply.Operation(5, 3))

	// === STRUCT WITH POINTERS ===
	fmt.Println("\n--- STRUCT WITH POINTERS ---")

	// Create a linked list
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}

	node1.Next = node2
	node2.Next = node3

	// Traverse the list
	fmt.Print("Linked list: ")
	current := node1
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
	fmt.Println()

	// === STRUCT WITH INTERFACE ===
	fmt.Println("\n--- STRUCT WITH INTERFACE ---")

	circle := Circle{Radius: 5}
	container := Container{
		Name:  "Circle Container",
		Shape: circle,
	}

	fmt.Printf("Container: %s\n", container.Name)
	fmt.Printf("Shape area: %.2f\n", container.Shape.Area())

	// === STRUCT TAGS ===
	fmt.Println("\n--- STRUCT TAGS ---")

	user := User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "secret123",
	}

	fmt.Printf("User struct: %+v\n", user)
	// Note: In real applications, you'd use reflection to read tags
	// for JSON marshaling, database operations, etc.

	// === STRUCT INITIALIZATION PATTERNS ===
	fmt.Println("\n--- STRUCT INITIALIZATION PATTERNS ---")

	// Constructor function pattern
	NewPerson := func(name string, age int) Person {
		return Person{Name: name, Age: age}
	}

	person1 := NewPerson("Alice", 30)
	fmt.Printf("Constructor pattern: %+v\n", person1)

	// Builder pattern
	type PersonBuilder struct {
		person Person
	}

	NewPersonBuilder := func() *PersonBuilder {
		return &PersonBuilder{}
	}

	SetName := func(pb *PersonBuilder, name string) *PersonBuilder {
		pb.person.Name = name
		return pb
	}

	SetAge := func(pb *PersonBuilder, age int) *PersonBuilder {
		pb.person.Age = age
		return pb
	}

	Build := func(pb *PersonBuilder) Person {
		return pb.person
	}

	person2 := Build(SetAge(SetName(NewPersonBuilder(), "Bob"), 25))
	fmt.Printf("Builder pattern: %+v\n", person2)

	// === STRUCT SLICES ===
	fmt.Println("\n--- STRUCT SLICES ---")

	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	fmt.Println("People:")
	for i, person := range people {
		fmt.Printf("  %d: %+v\n", i, person)
	}

	// Modify struct in slice
	people[0].Age = 31
	fmt.Printf("After modification: %+v\n", people[0])

	// === STRUCT MAPS ===
	fmt.Println("\n--- STRUCT MAPS ---")

	personMap := map[string]Person{
		"alice":   {Name: "Alice", Age: 30},
		"bob":     {Name: "Bob", Age: 25},
		"charlie": {Name: "Charlie", Age: 35},
	}

	fmt.Println("Person map:")
	for key, person := range personMap {
		fmt.Printf("  %s: %+v\n", key, person)
	}

	// === STRUCT MEMORY LAYOUT ===
	fmt.Println("\n--- STRUCT MEMORY LAYOUT ---")

	type SmallStruct struct {
		A int8
		B int32
		C int8
	}

	small := SmallStruct{A: 1, B: 2, C: 3}
	fmt.Printf("SmallStruct: %+v\n", small)
	fmt.Printf("Size of SmallStruct: %d bytes\n", unsafe.Sizeof(small))

	// Field alignment affects struct size
	type AlignedStruct struct {
		A int8
		C int8
		B int32
	}

	aligned := AlignedStruct{A: 1, C: 3, B: 2}
	fmt.Printf("AlignedStruct: %+v\n", aligned)
	fmt.Printf("Size of AlignedStruct: %d bytes\n", unsafe.Sizeof(aligned))

	// === STRUCT BEST PRACTICES ===
	fmt.Println("\n--- STRUCT BEST PRACTICES ---")
	fmt.Println("1. Use meaningful struct and field names")
	fmt.Println("2. Group related fields together")
	fmt.Println("3. Use composition over inheritance")
	fmt.Println("4. Consider field alignment for memory efficiency")
	fmt.Println("5. Use struct tags for metadata")
	fmt.Println("6. Implement String() method for custom formatting")
	fmt.Println("7. Use constructor functions for complex initialization")
	fmt.Println("8. Make zero value useful when possible")
	fmt.Println("9. Use embedded structs for composition")
	fmt.Println("10. Consider using pointers for large structs")
}
