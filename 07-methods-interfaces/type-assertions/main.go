package main

import (
	"fmt"
	"reflect"
)

// === GO TYPE ASSERTIONS COMPREHENSIVE GUIDE ===

/*
TYPE ASSERTIONS PHILOSOPHY:
- Type assertions extract the underlying concrete value from an interface
- Used to access the dynamic type and value stored in an interface
- Can be used to check if an interface holds a specific type
- Essential for working with interface{} and other interfaces

COMPARISON WITH JAVASCRIPT:
// JavaScript - Type checking
if (typeof value === 'string') {
  console.log(value.toUpperCase());
}

if (value instanceof Array) {
  console.log(value.length);
}

// Go - Type assertions
if str, ok := value.(string); ok {
  fmt.Println(strings.ToUpper(str))
}

if arr, ok := value.([]int); ok {
  fmt.Println(len(arr))
}
*/

// === INTERFACE DEFINITIONS ===

type Shape interface {
	Area() float64
}

type Drawable interface {
	Draw()
}

type ColoredShape interface {
	Shape
	Color() string
}

// === CONCRETE TYPES ===

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Draw() {
	fmt.Printf("Drawing circle with radius %.2f\n", c.Radius)
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Draw() {
	fmt.Printf("Drawing rectangle %.2f x %.2f\n", r.Width, r.Height)
}

type ColoredCircle struct {
	Circle
	ColorName string
}

func (cc ColoredCircle) Color() string {
	return cc.ColorName
}

func (cc ColoredCircle) Draw() {
	fmt.Printf("Drawing %s circle with radius %.2f\n", cc.ColorName, cc.Radius)
}

// === CUSTOM TYPES ===

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (age %d)", p.Name, p.Age)
}

type Employee struct {
	Person
	ID       int
	Position string
}

func (e Employee) String() string {
	return fmt.Sprintf("Employee %d: %s - %s", e.ID, e.Person.String(), e.Position)
}

type basicDrawable struct{}

func (basicDrawable) Draw() {
	fmt.Println("Drawing basic shape")
}

func main() {
	fmt.Println("=== GO TYPE ASSERTIONS COMPREHENSIVE GUIDE ===")

	// === BASIC TYPE ASSERTIONS ===
	fmt.Println("\n1. BASIC TYPE ASSERTIONS:")

	var value interface{} = "hello, world"

	// Basic type assertion (panics if wrong type)
	str := value.(string)
	fmt.Printf("String value: %s\n", str)

	// Safe type assertion with ok value
	if str2, ok := value.(string); ok {
		fmt.Printf("Safe string assertion: %s\n", str2)
	} else {
		fmt.Println("Not a string")
	}

	// Type assertion with wrong type
	if num, ok := value.(int); ok {
		fmt.Printf("Integer value: %d\n", num)
	} else {
		fmt.Println("Not an integer")
	}

	// === INTERFACE{} EXAMPLES ===
	fmt.Println("\n2. INTERFACE{} EXAMPLES:")

	values := []interface{}{
		42,
		"hello",
		3.14,
		true,
		[]int{1, 2, 3},
		map[string]int{"a": 1},
		Person{Name: "Alice", Age: 30},
	}

	for i, v := range values {
		fmt.Printf("Value %d: ", i)

		// Type assertion for different types
		switch val := v.(type) {
		case int:
			fmt.Printf("Integer: %d\n", val)
		case string:
			fmt.Printf("String: %s\n", val)
		case float64:
			fmt.Printf("Float: %.2f\n", val)
		case bool:
			fmt.Printf("Boolean: %t\n", val)
		case []int:
			fmt.Printf("Int slice: %v\n", val)
		case map[string]int:
			fmt.Printf("String-int map: %v\n", val)
		case Person:
			fmt.Printf("Person: %s\n", val.String())
		default:
			fmt.Printf("Unknown type: %T\n", val)
		}
	}

	// === INTERFACE TYPE ASSERTIONS ===
	fmt.Println("\n3. INTERFACE TYPE ASSERTIONS:")

	shapes := []Shape{
		Circle{Radius: 5.0},
		Rectangle{Width: 10.0, Height: 5.0},
		ColoredCircle{
			Circle:    Circle{Radius: 3.0},
			ColorName: "red",
		},
	}

	for i, shape := range shapes {
		fmt.Printf("Shape %d - Area: %.2f\n", i, shape.Area())

		// Type assertion to concrete types
		switch s := shape.(type) {
		case Circle:
			fmt.Printf("  Circle with radius: %.2f\n", s.Radius)
		case Rectangle:
			fmt.Printf("  Rectangle: %.2f x %.2f\n", s.Width, s.Height)
		case ColoredCircle:
			fmt.Printf("  Colored circle: %s, radius: %.2f\n", s.ColorName, s.Radius)
		}

		// Type assertion to interface
		if drawable, ok := shape.(Drawable); ok {
			drawable.Draw()
		}

		// Type assertion to embedded interface
		if colored, ok := shape.(ColoredShape); ok {
			fmt.Printf("  Color: %s\n", colored.Color())
		}
	}

	// === MULTIPLE INTERFACE ASSERTIONS ===
	fmt.Println("\n4. MULTIPLE INTERFACE ASSERTIONS:")

	var obj interface{} = ColoredCircle{
		Circle:    Circle{Radius: 4.0},
		ColorName: "blue",
	}

	// Check multiple interfaces
	if shape, ok := obj.(Shape); ok {
		fmt.Printf("Object is a Shape with area: %.2f\n", shape.Area())
	}

	if drawable, ok := obj.(Drawable); ok {
		fmt.Print("Object is Drawable: ")
		drawable.Draw()
	}

	if colored, ok := obj.(ColoredShape); ok {
		fmt.Printf("Object is ColoredShape with color: %s\n", colored.Color())
	}

	// === TYPE SWITCHES ===
	fmt.Println("\n5. TYPE SWITCHES:")

	processValue := func(value interface{}) {
		switch v := value.(type) {
		case nil:
			fmt.Println("Value is nil")
		case int:
			fmt.Printf("Integer: %d (doubled: %d)\n", v, v*2)
		case string:
			fmt.Printf("String: %s (length: %d)\n", v, len(v))
		case []int:
			fmt.Printf("Int slice: %v (sum: %d)\n", v, sum(v))
		case Person:
			fmt.Printf("Person: %s\n", v.String())
		case Employee:
			fmt.Printf("Employee: %s\n", v.String())
		default:
			fmt.Printf("Unknown type: %T\n", v)
		}
	}

	// Test type switches
	testValues := []interface{}{
		nil,
		42,
		"hello",
		[]int{1, 2, 3, 4, 5},
		Person{Name: "Bob", Age: 25},
		Employee{Person: Person{Name: "Charlie", Age: 35}, ID: 123, Position: "Manager"},
		3.14,
	}

	for _, val := range testValues {
		processValue(val)
	}

	// === ERROR HANDLING WITH TYPE ASSERTIONS ===
	fmt.Println("\n6. ERROR HANDLING WITH TYPE ASSERTIONS:")

	// Function that returns interface{} or error
	getValue := func(key string) (interface{}, error) {
		data := map[string]interface{}{
			"name":   "Alice",
			"age":    30,
			"scores": []int{85, 90, 78},
			"active": true,
		}

		if val, exists := data[key]; exists {
			return val, nil
		}
		return nil, fmt.Errorf("key not found: %s", key)
	}

	// Safe type assertion with error handling
	if val, err := getValue("name"); err == nil {
		if name, ok := val.(string); ok {
			fmt.Printf("Name: %s\n", name)
		} else {
			fmt.Printf("Name is not a string: %T\n", val)
		}
	}

	if val, err := getValue("age"); err == nil {
		if age, ok := val.(int); ok {
			fmt.Printf("Age: %d\n", age)
		} else {
			fmt.Printf("Age is not an int: %T\n", val)
		}
	}

	// === REFLECTION VS TYPE ASSERTIONS ===
	fmt.Println("\n7. REFLECTION VS TYPE ASSERTIONS:")

	var data interface{} = Person{Name: "David", Age: 28}

	// Using type assertion
	if person, ok := data.(Person); ok {
		fmt.Printf("Type assertion - Person: %s, Age: %d\n", person.Name, person.Age)
	}

	// Using reflection
	rv := reflect.ValueOf(data)
	rt := reflect.TypeOf(data)

	if rt.Kind() == reflect.Struct {
		fmt.Printf("Reflection - Type: %s\n", rt.Name())
		for i := 0; i < rv.NumField(); i++ {
			field := rt.Field(i)
			value := rv.Field(i)
			fmt.Printf("  Field %s: %v\n", field.Name, value.Interface())
		}
	}

	// === PERFORMANCE CONSIDERATIONS ===
	fmt.Println("\n8. PERFORMANCE CONSIDERATIONS:")

	// Type assertions are generally faster than reflection
	benchmarkData := []interface{}{
		"string1", "string2", "string3",
		42, 43, 44,
		3.14, 2.71, 1.41,
	}

	stringCount := 0
	intCount := 0
	floatCount := 0

	for _, item := range benchmarkData {
		switch item.(type) {
		case string:
			stringCount++
		case int:
			intCount++
		case float64:
			floatCount++
		}
	}

	fmt.Printf("Type counts - Strings: %d, Ints: %d, Floats: %d\n",
		stringCount, intCount, floatCount)

	// === COMMON PATTERNS ===
	fmt.Println("\n9. COMMON PATTERNS:")

	// Pattern 1: Optional fields in structs
	type Config struct {
		Name    string
		Timeout interface{} // Can be int (seconds) or string (duration)
	}

	config := Config{
		Name:    "MyApp",
		Timeout: 30, // Could also be "30s"
	}

	// Handle different timeout types
	switch timeout := config.Timeout.(type) {
	case int:
		fmt.Printf("Config timeout: %d seconds\n", timeout)
	case string:
		fmt.Printf("Config timeout: %s\n", timeout)
	default:
		fmt.Printf("Unknown timeout type: %T\n", timeout)
	}

	// Pattern 2: JSON unmarshaling
	jsonData := map[string]interface{}{
		"name":   "Product A",
		"price":  99.99,
		"active": true,
		"tags":   []interface{}{"electronics", "gadget"},
	}

	// Extract and convert values
	if name, ok := jsonData["name"].(string); ok {
		fmt.Printf("Product name: %s\n", name)
	}

	if price, ok := jsonData["price"].(float64); ok {
		fmt.Printf("Product price: $%.2f\n", price)
	}

	if tags, ok := jsonData["tags"].([]interface{}); ok {
		fmt.Print("Product tags: ")
		for _, tag := range tags {
			if tagStr, ok := tag.(string); ok {
				fmt.Printf("%s ", tagStr)
			}
		}
		fmt.Println()
	}

	// === EMPTY INTERFACE HANDLING ===
	fmt.Println("\n10. EMPTY INTERFACE HANDLING:")

	// Function that accepts any type
	printValue := func(value interface{}) {
		switch v := value.(type) {
		case nil:
			fmt.Println("Value: <nil>")
		case bool:
			fmt.Printf("Boolean: %t\n", v)
		case int, int8, int16, int32, int64:
			fmt.Printf("Integer: %d\n", v)
		case uint, uint8, uint16, uint32, uint64:
			fmt.Printf("Unsigned Integer: %d\n", v)
		case float32, float64:
			fmt.Printf("Float: %f\n", v)
		case string:
			fmt.Printf("String: %q\n", v)
		case []byte:
			fmt.Printf("Bytes: %q\n", string(v))
		default:
			fmt.Printf("Other: %T = %v\n", v, v)
		}
	}

	// Test various types
	printValue(nil)
	printValue(true)
	printValue(42)
	printValue(3.14)
	printValue("hello")
	printValue([]byte("world"))
	printValue(Person{Name: "Eve", Age: 32})

	// === BEST PRACTICES ===
	fmt.Println("\n11. BEST PRACTICES:")

	fmt.Println("✓ Always use the comma ok idiom: val, ok := iface.(Type)")
	fmt.Println("✓ Use type switches for multiple type checks")
	fmt.Println("✓ Prefer specific interfaces over interface{}")
	fmt.Println("✓ Handle the case where type assertion fails")
	fmt.Println("✓ Use type assertions for performance-critical code")
	fmt.Println("✗ Don't use type assertions without checking ok")
	fmt.Println("✗ Don't overuse interface{} - prefer typed interfaces")
	fmt.Println("✗ Don't ignore the possibility of nil interfaces")

	// === ADVANCED PATTERNS ===
	fmt.Println("\n12. ADVANCED PATTERNS:")

	// Pattern: Interface upgrade
	upgradeToDrawable := func(shape Shape) Drawable {
		if drawable, ok := shape.(Drawable); ok {
			return drawable
		}
		// Return a wrapper that makes any shape drawable
		return struct {
			Shape
			Drawable
		}{shape, basicDrawable{}}
	}

	circle := Circle{Radius: 2.0}
	drawable := upgradeToDrawable(circle)
	drawable.Draw()

	// Pattern: Type-safe casting
	toString := func(value interface{}) (string, bool) {
		switch v := value.(type) {
		case string:
			return v, true
		case fmt.Stringer:
			return v.String(), true
		case int:
			return fmt.Sprintf("%d", v), true
		case float64:
			return fmt.Sprintf("%.2f", v), true
		case bool:
			return fmt.Sprintf("%t", v), true
		default:
			return "", false
		}
	}

	testValues2 := []interface{}{
		"hello",
		42,
		3.14,
		true,
		Person{Name: "Frank", Age: 40},
		[]int{1, 2, 3},
	}

	for _, val := range testValues2 {
		if str, ok := toString(val); ok {
			fmt.Printf("Converted to string: %s\n", str)
		} else {
			fmt.Printf("Cannot convert %T to string\n", val)
		}
	}

	fmt.Println("\n=== END OF TYPE ASSERTIONS GUIDE ===")
}

// Helper function
func sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

/*
RUNNING THE PROGRAM:
go run main.go

LEARNING POINTS:

1. TYPE ASSERTIONS:
   - Extract concrete value from interface: value.(Type)
   - Safe form: value, ok := iface.(Type)
   - Panics if wrong type without ok check

2. SYNTAX:
   - Basic: value := iface.(Type)
   - Safe: value, ok := iface.(Type)
   - Type switch: switch v := iface.(type)

3. COMMON PATTERNS:
   - Check if interface holds specific type
   - Extract value from interface{}
   - Convert between interface types
   - Handle multiple possible types

4. TYPE SWITCHES:
   - switch v := value.(type) { case Type1: ... }
   - More efficient than multiple type assertions
   - Can handle multiple types in one switch

5. INTERFACE COMPATIBILITY:
   - Assert to more specific interfaces
   - Check if type implements interface
   - Extract methods from interface values

6. ERROR HANDLING:
   - Always check ok value in production code
   - Handle failed type assertions gracefully
   - Use type switches for multiple possibilities

7. PERFORMANCE:
   - Type assertions are fast
   - Faster than reflection
   - Type switches are efficient

8. BEST PRACTICES:
   - Use comma ok idiom
   - Prefer typed interfaces over interface{}
   - Handle nil interfaces
   - Use type switches for multiple types

This demonstrates comprehensive type assertion usage in Go
for safe and efficient interface value extraction.
*/
