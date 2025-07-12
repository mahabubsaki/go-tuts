package main

import (
	"fmt"
	"time"
)

// === GO STRUCT EMBEDDING COMPREHENSIVE GUIDE ===

/*
STRUCT EMBEDDING PHILOSOPHY:
- Go supports composition through struct embedding
- Embedded structs provide their fields and methods to the embedding struct
- No inheritance like OOP languages, but similar functionality
- Promotes composition over inheritance principle

COMPARISON WITH JAVASCRIPT:
// JavaScript - Inheritance/Composition
class Animal {
  constructor(name) {
    this.name = name;
  }
  speak() {
    console.log(`${this.name} makes a sound`);
  }
}

class Dog extends Animal {
  constructor(name, breed) {
    super(name);
    this.breed = breed;
  }

  speak() {
    console.log(`${this.name} barks`);
  }
}

// Go - Struct Embedding
type Animal struct {
  Name string
}

func (a Animal) Speak() {
  fmt.Printf("%s makes a sound\n", a.Name)
}

type Dog struct {
  Animal  // Embedded struct
  Breed string
}

func (d Dog) Speak() {
  fmt.Printf("%s barks\n", d.Name)  // Access embedded field
}
*/

// === BASIC EMBEDDING ===

type Animal struct {
	Name    string
	Age     int
	Species string
}

func (a Animal) Speak() {
	fmt.Printf("%s makes a sound\n", a.Name)
}

func (a Animal) Info() {
	fmt.Printf("%s is a %d-year-old %s\n", a.Name, a.Age, a.Species)
}

type Dog struct {
	Animal // Embedded struct
	Breed  string
}

func (d Dog) Speak() {
	fmt.Printf("%s barks\n", d.Name) // Access embedded field directly
}

func (d Dog) Fetch() {
	fmt.Printf("%s is fetching the ball\n", d.Name)
}

type Cat struct {
	Animal // Embedded struct
	Indoor bool
}

func (c Cat) Speak() {
	fmt.Printf("%s meows\n", c.Name)
}

func (c Cat) Purr() {
	fmt.Printf("%s is purring\n", c.Name)
}

// === MULTIPLE EMBEDDING ===

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p Person) Greet() {
	fmt.Printf("Hello, I'm %s\n", p.FullName())
}

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
	Country string
}

func (a Address) String() string {
	return fmt.Sprintf("%s, %s, %s %s, %s", a.Street, a.City, a.State, a.ZipCode, a.Country)
}

type ContactInfo struct {
	Email string
	Phone string
}

func (c ContactInfo) String() string {
	return fmt.Sprintf("Email: %s, Phone: %s", c.Email, c.Phone)
}

type Employee struct {
	Person      // Embedded struct
	Address     // Embedded struct
	ContactInfo // Embedded struct
	ID          int
	Position    string
	Salary      float64
	HireDate    time.Time
}

func (e Employee) String() string {
	return fmt.Sprintf("Employee %d: %s (%s) - %s", e.ID, e.FullName(), e.Position, e.ContactInfo.String())
}

// === EMBEDDED INTERFACES ===

type Speaker interface {
	Speak()
}

type Mover interface {
	Move()
}

type Robot struct {
	Name    string
	Model   string
	Battery int
}

func (r Robot) Speak() {
	fmt.Printf("Robot %s says: Hello\n", r.Name)
}

func (r Robot) Move() {
	fmt.Printf("Robot %s is moving\n", r.Name)
}

type AdvancedRobot struct {
	Robot // Embedded struct
	AI    bool
}

func (ar AdvancedRobot) Think() {
	if ar.AI {
		fmt.Printf("Robot %s is thinking\n", ar.Name)
	}
}

// === EMBEDDING WITH POINTERS ===

type Engine struct {
	Type  string
	Power int
}

func (e Engine) Start() {
	fmt.Printf("%s engine starting (%d HP)\n", e.Type, e.Power)
}

func (e Engine) Stop() {
	fmt.Printf("%s engine stopping\n", e.Type)
}

type Car struct {
	*Engine // Embedded pointer
	Brand   string
	Model   string
	Year    int
}

func (c Car) String() string {
	return fmt.Sprintf("%d %s %s", c.Year, c.Brand, c.Model)
}

// === FIELD PROMOTION ===

type Base struct {
	ID   int
	Name string
}

type Derived struct {
	Base  // Embedded struct
	Value string
}

// === METHOD PROMOTION ===

type Timer struct {
	Start time.Time
	End   time.Time
}

func (t Timer) Duration() time.Duration {
	return t.End.Sub(t.Start)
}

func (t Timer) String() string {
	return fmt.Sprintf("Timer: %v", t.Duration())
}

type Task struct {
	Timer // Embedded struct
	Name  string
	Done  bool
}

func (t Task) Complete() {
	t.Done = true
	t.End = time.Now()
}

// === EMBEDDING CONFLICTS ===

type A struct {
	Value int
}

func (a A) Method() {
	fmt.Println("Method from A")
}

type B struct {
	Value int
}

func (b B) Method() {
	fmt.Println("Method from B")
}

type C struct {
	A     // Embedded struct
	B     // Embedded struct
	Local int
}

func (c C) Method() {
	fmt.Println("Method from C")
}

// Timestamped mixin
type Timestamped struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Timestamped) Touch() {
	t.UpdatedAt = time.Now()
}

// Serializable mixin
type Serializable struct{}

func (s Serializable) ToJSON() string {
	return "JSON representation"
}

// Database model base
type Model struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Model) Save() {
	m.UpdatedAt = time.Now()
	fmt.Printf("Model %d saved at %v\n", m.ID, m.UpdatedAt)
}

// Specific models
type Article struct {
	Model    // Embedded base model
	Title    string
	Content  string
	AuthorID int
}

type Comment struct {
	Model     // Embedded base model
	ArticleID int
	Content   string
	AuthorID  int
}

func main() {
	fmt.Println("=== GO STRUCT EMBEDDING COMPREHENSIVE GUIDE ===")

	// === BASIC EMBEDDING ===
	fmt.Println("\n1. BASIC EMBEDDING:")

	// Create embedded structs
	dog := Dog{
		Animal: Animal{
			Name:    "Rex",
			Age:     3,
			Species: "Canine",
		},
		Breed: "Golden Retriever",
	}

	cat := Cat{
		Animal: Animal{
			Name:    "Whiskers",
			Age:     2,
			Species: "Feline",
		},
		Indoor: true,
	}

	// Access embedded fields directly
	fmt.Printf("Dog name: %s\n", dog.Name)   // Direct access to embedded field
	fmt.Printf("Dog age: %d\n", dog.Age)     // Direct access to embedded field
	fmt.Printf("Dog breed: %s\n", dog.Breed) // Own field

	// Call embedded methods
	dog.Info()  // Method from embedded Animal
	dog.Speak() // Overridden method
	dog.Fetch() // Own method

	fmt.Println()
	cat.Info()  // Method from embedded Animal
	cat.Speak() // Overridden method
	cat.Purr()  // Own method

	// === MULTIPLE EMBEDDING ===
	fmt.Println("\n2. MULTIPLE EMBEDDING:")

	employee := Employee{
		Person: Person{
			FirstName: "John",
			LastName:  "Doe",
			Age:       30,
		},
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			State:   "NY",
			ZipCode: "10001",
			Country: "USA",
		},
		ContactInfo: ContactInfo{
			Email: "john.doe@company.com",
			Phone: "555-1234",
		},
		ID:       1001,
		Position: "Software Engineer",
		Salary:   85000.0,
		HireDate: time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC),
	}

	// Access fields from multiple embedded structs
	fmt.Printf("Employee name: %s\n", employee.FullName())              // Method from Person
	fmt.Printf("Employee address: %s\n", employee.Address.String())     // Method from Address
	fmt.Printf("Employee contact: %s\n", employee.ContactInfo.String()) // Method from ContactInfo
	fmt.Printf("Employee info: %s\n", employee.String())                // Own method

	// Call embedded methods
	employee.Greet() // Method from Person

	// === EMBEDDED INTERFACES ===
	fmt.Println("\n3. EMBEDDED INTERFACES:")

	robot := Robot{
		Name:    "R2D2",
		Model:   "Astromech",
		Battery: 85,
	}

	advancedRobot := AdvancedRobot{
		Robot: robot,
		AI:    true,
	}

	// Use embedded methods
	advancedRobot.Speak() // Through embedded Robot
	advancedRobot.Move()  // Through embedded Robot
	advancedRobot.Think() // Own method

	// === EMBEDDING WITH POINTERS ===
	fmt.Println("\n4. EMBEDDING WITH POINTERS:")

	engine := &Engine{
		Type:  "V8",
		Power: 450,
	}

	car := Car{
		Engine: engine,
		Brand:  "Ford",
		Model:  "Mustang",
		Year:   2023,
	}

	fmt.Printf("Car: %s\n", car.String())
	car.Start() // Method from embedded pointer
	car.Stop()  // Method from embedded pointer

	// Modify embedded pointer
	car.Engine.Power = 500
	fmt.Printf("Updated engine power: %d HP\n", car.Engine.Power)

	// === FIELD PROMOTION ===
	fmt.Println("\n5. FIELD PROMOTION:")

	derived := Derived{
		Base: Base{
			ID:   1,
			Name: "Base Name",
		},
		Value: "Derived Value",
	}

	// Access promoted fields
	fmt.Printf("Derived ID: %d\n", derived.ID)       // Promoted field
	fmt.Printf("Derived Name: %s\n", derived.Name)   // Promoted field
	fmt.Printf("Derived Value: %s\n", derived.Value) // Own field

	// Access through embedded struct
	fmt.Printf("Base ID: %d\n", derived.Base.ID)
	fmt.Printf("Base Name: %s\n", derived.Base.Name)

	// === METHOD PROMOTION ===
	fmt.Println("\n6. METHOD PROMOTION:")

	task := Task{
		Timer: Timer{
			Start: time.Now(),
		},
		Name: "Complete project",
		Done: false,
	}

	// Simulate work
	time.Sleep(1 * time.Millisecond)
	task.Complete()

	// Use promoted methods
	fmt.Printf("Task duration: %v\n", task.Duration()) // Promoted method
	fmt.Printf("Task info: %s\n", task.String())       // Promoted method

	// === EMBEDDING CONFLICTS ===
	fmt.Println("\n7. EMBEDDING CONFLICTS:")

	c := C{
		A:     A{Value: 10},
		B:     B{Value: 20},
		Local: 30,
	}

	// Accessing conflicting fields requires qualification
	fmt.Printf("A.Value: %d\n", c.A.Value)
	fmt.Printf("B.Value: %d\n", c.B.Value)
	fmt.Printf("Local: %d\n", c.Local)

	// Method resolution
	c.Method()   // Calls C's method
	c.A.Method() // Calls A's method
	c.B.Method() // Calls B's method

	// === ANONYMOUS FIELDS ===
	fmt.Println("\n8. ANONYMOUS FIELDS:")

	// Anonymous field of basic type
	type Container struct {
		string // Anonymous field
		int    // Anonymous field
		Name   string
	}

	container := Container{
		string: "Hello",
		int:    42,
		Name:   "Container",
	}

	fmt.Printf("Container string: %s\n", container.string)
	fmt.Printf("Container int: %d\n", container.int)
	fmt.Printf("Container name: %s\n", container.Name)

	// === COMPOSITION PATTERNS ===
	fmt.Println("\n9. COMPOSITION PATTERNS:")

	// Mixin pattern
	type User struct {
		Timestamped // Mixin
		ID          int
		Username    string
		Email       string
	}

	user := User{
		Timestamped: Timestamped{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ID:       1,
		Username: "alice",
		Email:    "alice@example.com",
	}

	fmt.Printf("User created: %v\n", user.CreatedAt)
	user.Touch() // Mixin method
	fmt.Printf("User updated: %v\n", user.UpdatedAt)

	// === INTERFACE SATISFACTION ===
	fmt.Println("\n10. INTERFACE SATISFACTION:")

	// Interface
	type Stringer interface {
		String() string
	}

	// Function that accepts the interface
	printStringer := func(s Stringer) {
		fmt.Printf("String: %s\n", s.String())
	}

	// Embedded struct satisfies interface
	printStringer(employee) // Employee embeds Address which has String()
	printStringer(car)      // Car has String() method

	// === EMBEDDING BEST PRACTICES ===
	fmt.Println("\n11. EMBEDDING BEST PRACTICES:")

	// Good: Embedding for "is-a" relationship
	type Vehicle struct {
		Brand string
		Model string
	}

	type Truck struct {
		Vehicle   // Truck "is-a" Vehicle
		Capacity  int
		FourWheel bool
	}

	truck := Truck{
		Vehicle: Vehicle{
			Brand: "Ford",
			Model: "F-150",
		},
		Capacity:  1000,
		FourWheel: true,
	}

	fmt.Printf("Truck: %s %s (Capacity: %d kg)\n", truck.Brand, truck.Model, truck.Capacity)

	// Good: Embedding for mixins
	type Document struct {
		Serializable // Mixin functionality
		Title        string
		Content      string
	}

	doc := Document{
		Title:   "My Document",
		Content: "Document content",
	}

	fmt.Printf("Document JSON: %s\n", doc.ToJSON())

	// === REAL-WORLD EXAMPLE ===
	fmt.Println("\n12. REAL-WORLD EXAMPLE:")

	article := Article{
		Model: Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Title:    "Go Programming",
		Content:  "Learn Go programming language",
		AuthorID: 1,
	}

	comment := Comment{
		Model: Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ArticleID: 1,
		Content:   "Great article!",
		AuthorID:  2,
	}

	article.Save() // Embedded method
	comment.Save() // Embedded method

	fmt.Printf("Article: %s (ID: %d)\n", article.Title, article.ID)
	fmt.Printf("Comment: %s (ID: %d)\n", comment.Content, comment.ID)

	fmt.Println("\n=== END OF STRUCT EMBEDDING GUIDE ===")
}

/*
RUNNING THE PROGRAM:
go run main.go

LEARNING POINTS:

1. STRUCT EMBEDDING:
   - Compose structs by embedding other structs
   - Embedded fields and methods are promoted
   - Access embedded fields directly or through embedded struct

2. FIELD PROMOTION:
   - Fields from embedded structs are promoted to embedding struct
   - Can access directly: obj.field or through embedded: obj.Embedded.field
   - Conflicts require qualification

3. METHOD PROMOTION:
   - Methods from embedded structs are promoted
   - Can be overridden in embedding struct
   - Original method still accessible through embedded struct

4. MULTIPLE EMBEDDING:
   - Can embed multiple structs
   - Name conflicts must be resolved with qualification
   - Useful for mixins and composition

5. EMBEDDED INTERFACES:
   - Can embed interfaces in structs
   - Struct must satisfy the embedded interface
   - Provides interface methods to the struct

6. POINTER EMBEDDING:
   - Can embed pointers to structs
   - Allows sharing of embedded struct between multiple embedders
   - Methods work the same way

7. COMPOSITION PATTERNS:
   - Mixin pattern for shared functionality
   - Base model pattern for common fields
   - Interface satisfaction through embedding

8. BEST PRACTICES:
   - Use embedding for "is-a" relationships
   - Use embedding for mixins and shared functionality
   - Avoid deep embedding hierarchies
   - Prefer composition over inheritance

This demonstrates comprehensive struct embedding in Go
for effective composition and code reuse patterns.
*/
