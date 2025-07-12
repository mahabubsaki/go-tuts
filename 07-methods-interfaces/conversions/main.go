package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// === GO TYPE CONVERSIONS COMPREHENSIVE GUIDE ===

/*
TYPE CONVERSIONS PHILOSOPHY:
- Go supports explicit type conversions between compatible types
- No implicit conversions (except for untyped constants)
- Type conversions are compile-time checked for safety
- Different from type assertions which work on interfaces

COMPARISON WITH JAVASCRIPT:
// JavaScript - Implicit and explicit conversions
let num = 42;
let str = String(num);        // Explicit conversion
let auto = num + "";          // Implicit conversion

// Go - Explicit conversions only
var num int = 42
var str string = strconv.Itoa(num)  // Must use conversion functions
var float64Val float64 = float64(num)  // Explicit type conversion
*/

// === CUSTOM TYPES ===

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", float64(c))
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", float64(f))
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%.2f K", float64(k))
}

// Conversion functions
func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(float64(c)*9/5 + 32)
}

func FahrenheitToCelsius(f Fahrenheit) Celsius {
	return Celsius((float64(f) - 32) * 5 / 9)
}

func CelsiusToKelvin(c Celsius) Kelvin {
	return Kelvin(float64(c) + 273.15)
}

type UserID int
type ProductID int
type OrderID int

func (u UserID) String() string {
	return fmt.Sprintf("User-%d", int(u))
}

func (p ProductID) String() string {
	return fmt.Sprintf("Product-%d", int(p))
}

func (o OrderID) String() string {
	return fmt.Sprintf("Order-%d", int(o))
}

// === BYTE TYPES ===

type MyString string
type MyBytes []byte

func (ms MyString) ToBytes() MyBytes {
	return MyBytes([]byte(ms))
}

func (mb MyBytes) ToString() MyString {
	return MyString(string(mb))
}

func main() {
	fmt.Println("=== GO TYPE CONVERSIONS COMPREHENSIVE GUIDE ===")

	// === BASIC NUMERIC CONVERSIONS ===
	fmt.Println("\n1. BASIC NUMERIC CONVERSIONS:")

	// Integer conversions
	var i int = 42
	var i8 int8 = int8(i)
	var i16 int16 = int16(i)
	var i32 int32 = int32(i)
	var i64 int64 = int64(i)

	fmt.Printf("int: %d\n", i)
	fmt.Printf("int8: %d\n", i8)
	fmt.Printf("int16: %d\n", i16)
	fmt.Printf("int32: %d\n", i32)
	fmt.Printf("int64: %d\n", i64)

	// Unsigned integer conversions
	var ui uint = uint(i)
	var ui8 uint8 = uint8(i)
	var ui16 uint16 = uint16(i)
	var ui32 uint32 = uint32(i)
	var ui64 uint64 = uint64(i)

	fmt.Printf("uint: %d\n", ui)
	fmt.Printf("uint8: %d\n", ui8)
	fmt.Printf("uint16: %d\n", ui16)
	fmt.Printf("uint32: %d\n", ui32)
	fmt.Printf("uint64: %d\n", ui64)

	// Float conversions
	var f32 float32 = float32(i)
	var f64 float64 = float64(i)

	fmt.Printf("float32: %f\n", f32)
	fmt.Printf("float64: %f\n", f64)

	// === FLOAT TO INTEGER CONVERSIONS ===
	fmt.Println("\n2. FLOAT TO INTEGER CONVERSIONS:")

	pi := 3.14159
	fmt.Printf("Original float: %f\n", pi)
	fmt.Printf("Converted to int: %d (truncated)\n", int(pi))
	fmt.Printf("Converted to int64: %d (truncated)\n", int64(pi))

	// Negative float
	negFloat := -2.9
	fmt.Printf("Negative float: %f\n", negFloat)
	fmt.Printf("Converted to int: %d (truncated)\n", int(negFloat))

	// === STRING CONVERSIONS ===
	fmt.Println("\n3. STRING CONVERSIONS:")

	// Numeric to string
	num := 123
	str := strconv.Itoa(num)
	fmt.Printf("int to string: %d -> %s\n", num, str)

	// String to numeric
	str = "456"
	if num, err := strconv.Atoi(str); err == nil {
		fmt.Printf("string to int: %s -> %d\n", str, num)
	}

	// Float to string
	floatVal := 3.14159
	floatStr := strconv.FormatFloat(floatVal, 'f', 2, 64)
	fmt.Printf("float to string: %f -> %s\n", floatVal, floatStr)

	// String to float
	floatStr = "2.71828"
	if parsedFloat, err := strconv.ParseFloat(floatStr, 64); err == nil {
		fmt.Printf("string to float: %s -> %f\n", floatStr, parsedFloat)
	}

	// Boolean conversions
	boolVal := true
	boolStr := strconv.FormatBool(boolVal)
	fmt.Printf("bool to string: %t -> %s\n", boolVal, boolStr)

	boolStr = "false"
	if parsedBool, err := strconv.ParseBool(boolStr); err == nil {
		fmt.Printf("string to bool: %s -> %t\n", boolStr, parsedBool)
	}

	// === BYTE AND STRING CONVERSIONS ===
	fmt.Println("\n4. BYTE AND STRING CONVERSIONS:")

	// String to byte slice
	text := "Hello, 世界"
	bytes := []byte(text)
	fmt.Printf("string to bytes: %s -> %v\n", text, bytes)
	fmt.Printf("byte representation: %q\n", bytes)

	// Byte slice to string
	backToString := string(bytes)
	fmt.Printf("bytes to string: %v -> %s\n", bytes, backToString)

	// Rune conversions
	runes := []rune(text)
	fmt.Printf("string to runes: %s -> %v\n", text, runes)
	fmt.Printf("rune representation: %q\n", runes)

	runeToString := string(runes)
	fmt.Printf("runes to string: %v -> %s\n", runes, runeToString)

	// === CUSTOM TYPE CONVERSIONS ===
	fmt.Println("\n5. CUSTOM TYPE CONVERSIONS:")

	// Temperature conversions
	celsius := Celsius(25.0)
	fahrenheit := CelsiusToFahrenheit(celsius)
	kelvin := CelsiusToKelvin(celsius)

	fmt.Printf("Temperature conversions:\n")
	fmt.Printf("  %s\n", celsius)
	fmt.Printf("  %s\n", fahrenheit)
	fmt.Printf("  %s\n", kelvin)

	// Back conversion
	backToCelsius := FahrenheitToCelsius(fahrenheit)
	fmt.Printf("  Back to Celsius: %s\n", backToCelsius)

	// ID type conversions
	userID := UserID(123)
	productID := ProductID(456)
	orderID := OrderID(789)

	fmt.Printf("ID conversions:\n")
	fmt.Printf("  %s\n", userID)
	fmt.Printf("  %s\n", productID)
	fmt.Printf("  %s\n", orderID)

	// Convert between ID types (be careful!)
	convertedID := UserID(int(productID))
	fmt.Printf("  Converted ProductID to UserID: %s\n", convertedID)

	// === UNSAFE CONVERSIONS ===
	fmt.Println("\n6. POTENTIALLY UNSAFE CONVERSIONS:")

	// Large number to smaller type (may overflow)
	largeNum := int64(1000000)
	smallNum := int8(largeNum) // Potential overflow
	fmt.Printf("int64 to int8: %d -> %d (may overflow)\n", largeNum, smallNum)

	// Negative to unsigned (wraps around)
	negNum := -10
	unsignedNum := uint(negNum)
	fmt.Printf("negative int to uint: %d -> %d (wraps around)\n", negNum, unsignedNum)

	// === POINTER CONVERSIONS ===
	fmt.Println("\n7. POINTER CONVERSIONS:")

	// Cannot convert between pointer types directly
	// Must use unsafe package or convert values

	x := 42
	ptr := &x
	fmt.Printf("Pointer to int: %p -> %d\n", ptr, *ptr)

	// Convert pointer value to different type
	floatPtr := (*float64)(nil)
	if ptr != nil {
		floatVal := float64(*ptr)
		floatPtr = &floatVal
	}
	if floatPtr != nil {
		fmt.Printf("Converted value through pointer: %f\n", *floatPtr)
	}

	// === INTERFACE CONVERSIONS ===
	fmt.Println("\n8. INTERFACE CONVERSIONS:")

	// Any type can be converted to interface{}
	var empty interface{}

	empty = 42
	fmt.Printf("int to interface{}: %v (type: %T)\n", empty, empty)

	empty = "hello"
	fmt.Printf("string to interface{}: %v (type: %T)\n", empty, empty)

	empty = []int{1, 2, 3}
	fmt.Printf("slice to interface{}: %v (type: %T)\n", empty, empty)

	// Convert back using type assertion
	if str, ok := empty.([]int); ok {
		fmt.Printf("interface{} back to slice: %v\n", str)
	}

	// === REFLECTION-BASED CONVERSIONS ===
	fmt.Println("\n9. REFLECTION-BASED CONVERSIONS:")

	// Using reflection to convert between types
	convertWithReflection := func(value interface{}, targetType reflect.Type) interface{} {
		v := reflect.ValueOf(value)
		if v.Type().ConvertibleTo(targetType) {
			return v.Convert(targetType).Interface()
		}
		return nil
	}

	// Example conversions
	intVal := 42

	// Convert int to float64
	if converted := convertWithReflection(intVal, reflect.TypeOf(float64(0))); converted != nil {
		fmt.Printf("Reflection conversion int to float64: %d -> %f\n", intVal, converted)
	}

	// Convert int to string (this won't work directly)
	if converted := convertWithReflection(intVal, reflect.TypeOf("")); converted != nil {
		fmt.Printf("Reflection conversion int to string: %d -> %s\n", intVal, converted)
	} else {
		fmt.Printf("Cannot convert int to string via reflection\n")
	}

	// === SLICE CONVERSIONS ===
	fmt.Println("\n10. SLICE CONVERSIONS:")

	// Convert between slice types
	intSlice := []int{1, 2, 3, 4, 5}

	// Convert to float64 slice
	floatSlice := make([]float64, len(intSlice))
	for i, v := range intSlice {
		floatSlice[i] = float64(v)
	}
	fmt.Printf("int slice to float64 slice: %v -> %v\n", intSlice, floatSlice)

	// Convert to string slice
	stringSlice := make([]string, len(intSlice))
	for i, v := range intSlice {
		stringSlice[i] = strconv.Itoa(v)
	}
	fmt.Printf("int slice to string slice: %v -> %v\n", intSlice, stringSlice)

	// === CUSTOM STRING CONVERSIONS ===
	fmt.Println("\n11. CUSTOM STRING CONVERSIONS:")

	myStr := MyString("Hello, World!")
	myBytes := myStr.ToBytes()
	fmt.Printf("MyString to MyBytes: %s -> %v\n", myStr, myBytes)

	backToMyString := myBytes.ToString()
	fmt.Printf("MyBytes to MyString: %v -> %s\n", myBytes, backToMyString)

	// === CONSTANT CONVERSIONS ===
	fmt.Println("\n12. CONSTANT CONVERSIONS:")

	// Untyped constants can be converted implicitly
	const untypedInt = 42
	const untypedFloat = 3.14

	var i1 int = untypedInt       // Implicit conversion
	var i2 int64 = untypedInt     // Implicit conversion
	var f1 float32 = untypedFloat // Implicit conversion
	var f2 float64 = untypedFloat // Implicit conversion

	fmt.Printf("Untyped constant conversions:\n")
	fmt.Printf("  int: %d\n", i1)
	fmt.Printf("  int64: %d\n", i2)
	fmt.Printf("  float32: %f\n", f1)
	fmt.Printf("  float64: %f\n", f2)

	// === CONVERSION FUNCTIONS ===
	fmt.Println("\n13. CONVERSION FUNCTIONS:")

	// Generic conversion function
	convertNumber := func(value interface{}) map[string]interface{} {
		result := make(map[string]interface{})

		switch v := value.(type) {
		case int:
			result["int"] = v
			result["int64"] = int64(v)
			result["float64"] = float64(v)
			result["string"] = strconv.Itoa(v)
		case float64:
			result["float64"] = v
			result["int"] = int(v)
			result["string"] = strconv.FormatFloat(v, 'f', 2, 64)
		case string:
			if intVal, err := strconv.Atoi(v); err == nil {
				result["int"] = intVal
				result["int64"] = int64(intVal)
				result["float64"] = float64(intVal)
			}
			result["string"] = v
		}

		return result
	}

	// Test conversion function
	testValues := []interface{}{42, 3.14, "123"}

	for _, val := range testValues {
		conversions := convertNumber(val)
		fmt.Printf("Conversions for %v (%T):\n", val, val)
		for k, v := range conversions {
			fmt.Printf("  %s: %v\n", k, v)
		}
	}

	// === BEST PRACTICES ===
	fmt.Println("\n14. BEST PRACTICES:")

	fmt.Println("✓ Always use explicit conversions")
	fmt.Println("✓ Check for potential overflow when converting to smaller types")
	fmt.Println("✓ Use strconv package for string conversions")
	fmt.Println("✓ Handle errors when parsing strings")
	fmt.Println("✓ Be careful with signed/unsigned conversions")
	fmt.Println("✓ Use custom types for type safety")
	fmt.Println("✗ Don't ignore potential data loss")
	fmt.Println("✗ Don't convert between incompatible types")

	// === REAL-WORLD EXAMPLES ===
	fmt.Println("\n15. REAL-WORLD EXAMPLES:")

	// Example 1: Configuration parsing
	_ = "timeout=30"                   // Simulated config value
	parts := []string{"timeout", "30"} // Simulated parsing
	if len(parts) == 2 {
		if timeoutVal, err := strconv.Atoi(parts[1]); err == nil {
			timeout := time.Duration(timeoutVal) * time.Second
			fmt.Printf("Config: %s = %v\n", parts[0], timeout)
		}
	}

	// Example 2: API response handling
	apiResponse := map[string]interface{}{
		"id":     123,
		"name":   "Product A",
		"price":  99.99,
		"active": true,
	}

	// Convert API response to typed struct
	type Product struct {
		ID     int
		Name   string
		Price  float64
		Active bool
	}

	var product Product
	if id, ok := apiResponse["id"].(int); ok {
		product.ID = id
	}
	if name, ok := apiResponse["name"].(string); ok {
		product.Name = name
	}
	if price, ok := apiResponse["price"].(float64); ok {
		product.Price = price
	}
	if active, ok := apiResponse["active"].(bool); ok {
		product.Active = active
	}

	fmt.Printf("API Response converted to struct: %+v\n", product)

	// Example 3: Database value conversion
	dbValue := "123.45"
	if price, err := strconv.ParseFloat(dbValue, 64); err == nil {
		priceInCents := int(price * 100)
		fmt.Printf("Database price conversion: %s -> $%.2f -> %d cents\n",
			dbValue, price, priceInCents)
	}

	fmt.Println("\n=== END OF TYPE CONVERSIONS GUIDE ===")
}

/*
RUNNING THE PROGRAM:
go run main.go

LEARNING POINTS:

1. TYPE CONVERSIONS:
   - Go requires explicit conversions between types
   - Syntax: Type(value)
   - No implicit conversions except for untyped constants

2. NUMERIC CONVERSIONS:
   - Between integer types: int32(value)
   - Between float types: float64(value)
   - Integer to float: float64(intValue)
   - Float to integer: int(floatValue) - truncates

3. STRING CONVERSIONS:
   - Use strconv package for string/number conversions
   - strconv.Itoa() for int to string
   - strconv.Atoi() for string to int
   - strconv.ParseFloat() for string to float

4. BYTE CONVERSIONS:
   - String to bytes: []byte(string)
   - Bytes to string: string([]byte)
   - String to runes: []rune(string)
   - Runes to string: string([]rune)

5. CUSTOM TYPE CONVERSIONS:
   - Convert between custom types with same underlying type
   - Type(value) where Type is custom type
   - Use conversion functions for complex conversions

6. INTERFACE CONVERSIONS:
   - Any type converts to interface{}
   - Use type assertions to convert back
   - Type switches for multiple possibilities

7. UNSAFE CONVERSIONS:
   - May cause overflow or data loss
   - Negative to unsigned wraps around
   - Large to small integer may overflow

8. BEST PRACTICES:
   - Always explicit conversions
   - Check for overflow potential
   - Handle conversion errors
   - Use custom types for type safety
   - Test edge cases

This demonstrates comprehensive type conversion patterns in Go
for safe and effective type manipulation.
*/
