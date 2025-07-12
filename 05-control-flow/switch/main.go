package main

import (
	"fmt"
	"runtime"
	"time"
)

// This file covers switch statements in Go
// Switch statements provide a cleaner alternative to long if-else chains

func main() {
	fmt.Println("=== GO SWITCH STATEMENTS - COMPLETE GUIDE ===")

	demonstrateBasicSwitch()
	demonstrateSwitchWithExpressions()
	demonstrateSwitchWithMultipleValues()
	demonstrateSwitchWithFallthrough()
	demonstrateSwitchWithoutExpression()
	demonstrateTypeSwitch()
	demonstrateSwitchWithInitializer()
	demonstrateSwitchBestPractices()
}

func demonstrateBasicSwitch() {
	fmt.Println("\n--- BASIC SWITCH STATEMENTS ---")

	// Basic switch with integer
	var day int = 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// Switch with string
	var grade string = "B"

	switch grade {
	case "A":
		fmt.Println("Excellent!")
	case "B":
		fmt.Println("Good job!")
	case "C":
		fmt.Println("Average")
	case "D":
		fmt.Println("Below average")
	case "F":
		fmt.Println("Failed")
	default:
		fmt.Println("Invalid grade")
	}

	// Switch with character (rune)
	var letter rune = 'A'

	switch letter {
	case 'A', 'E', 'I', 'O', 'U':
		fmt.Printf("'%c' is a vowel\n", letter)
	default:
		fmt.Printf("'%c' is a consonant\n", letter)
	}

	// Switch with boolean
	var isActive bool = true

	switch isActive {
	case true:
		fmt.Println("Status: Active")
	case false:
		fmt.Println("Status: Inactive")
	}

	// Switch with floating-point numbers
	var temperature float64 = 25.5

	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature < 10:
		fmt.Println("Cold")
	case temperature < 20:
		fmt.Println("Cool")
	case temperature < 30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Same basic syntax
	// JavaScript: Requires 'break' to prevent fall-through
	// JavaScript: Fall-through is default behavior
	// Go: No fall-through by default
	// Go: Must explicitly use 'fallthrough' keyword
	// Go: More type-safe comparisons
}

func demonstrateSwitchWithExpressions() {
	fmt.Println("\n--- SWITCH WITH EXPRESSIONS ---")

	// Switch with function calls
	var currentTime time.Time = time.Now()

	switch currentTime.Weekday() {
	case time.Monday:
		fmt.Println("Monday blues")
	case time.Tuesday:
		fmt.Println("Tuesday energy")
	case time.Wednesday:
		fmt.Println("Hump day")
	case time.Thursday:
		fmt.Println("Almost there")
	case time.Friday:
		fmt.Println("TGIF!")
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend!")
	}

	// Switch with mathematical expressions
	var x int = 10
	var y int = 5

	switch x + y {
	case 15:
		fmt.Println("Sum is 15")
	case 20:
		fmt.Println("Sum is 20")
	default:
		fmt.Printf("Sum is %d\n", x+y)
	}

	// Switch with string operations
	var filename string = "document.pdf"

	switch {
	case len(filename) == 0:
		fmt.Println("Empty filename")
	case len(filename) > 50:
		fmt.Println("Filename too long")
	default:
		fmt.Println("Valid filename length")
	}

	// Switch with modulo operation
	var number int = 17

	switch number % 2 {
	case 0:
		fmt.Printf("%d is even\n", number)
	case 1:
		fmt.Printf("%d is odd\n", number)
	}

	// Switch with range checking
	var score int = 85

	switch {
	case score >= 90:
		fmt.Printf("Grade A: %d\n", score)
	case score >= 80:
		fmt.Printf("Grade B: %d\n", score)
	case score >= 70:
		fmt.Printf("Grade C: %d\n", score)
	case score >= 60:
		fmt.Printf("Grade D: %d\n", score)
	default:
		fmt.Printf("Grade F: %d\n", score)
	}

	// Switch with type conversion
	var value interface{} = 42

	switch v := value.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case float64:
		fmt.Printf("Float: %.2f\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}

	// Switch with slice operations
	var numbers []int = []int{1, 2, 3, 4, 5}

	switch len(numbers) {
	case 0:
		fmt.Println("Empty slice")
	case 1:
		fmt.Printf("Single element: %d\n", numbers[0])
	default:
		fmt.Printf("Multiple elements: %d items\n", len(numbers))
	}
}

func demonstrateSwitchWithMultipleValues() {
	fmt.Println("\n--- SWITCH WITH MULTIPLE VALUES ---")

	// Multiple values in single case
	var month int = 6

	switch month {
	case 12, 1, 2:
		fmt.Println("Winter")
	case 3, 4, 5:
		fmt.Println("Spring")
	case 6, 7, 8:
		fmt.Println("Summer")
	case 9, 10, 11:
		fmt.Println("Fall")
	default:
		fmt.Println("Invalid month")
	}

	// Multiple character values
	var char rune = 'x'

	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Printf("'%c' is a lowercase vowel\n", char)
	case 'A', 'E', 'I', 'O', 'U':
		fmt.Printf("'%c' is an uppercase vowel\n", char)
	case 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z':
		fmt.Printf("'%c' is a lowercase consonant\n", char)
	case 'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z':
		fmt.Printf("'%c' is an uppercase consonant\n", char)
	default:
		fmt.Printf("'%c' is not a letter\n", char)
	}

	// Multiple string values
	var command string = "help"

	switch command {
	case "help", "h", "?":
		fmt.Println("Showing help information")
	case "quit", "q", "exit":
		fmt.Println("Goodbye!")
	case "version", "v":
		fmt.Println("Version 1.0.0")
	case "info", "i":
		fmt.Println("System information")
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}

	// Multiple numeric ranges
	var httpStatus int = 404

	switch {
	case httpStatus >= 200 && httpStatus < 300:
		fmt.Printf("Success: %d\n", httpStatus)
	case httpStatus >= 300 && httpStatus < 400:
		fmt.Printf("Redirection: %d\n", httpStatus)
	case httpStatus == 400, httpStatus == 401, httpStatus == 403, httpStatus == 404:
		fmt.Printf("Client error: %d\n", httpStatus)
	case httpStatus >= 400 && httpStatus < 500:
		fmt.Printf("Other client error: %d\n", httpStatus)
	case httpStatus >= 500 && httpStatus < 600:
		fmt.Printf("Server error: %d\n", httpStatus)
	default:
		fmt.Printf("Unknown status: %d\n", httpStatus)
	}

	// Multiple boolean conditions
	var isLoggedIn bool = true
	var isAdmin bool = false

	switch {
	case isLoggedIn && isAdmin:
		fmt.Println("Admin access granted")
	case isLoggedIn && !isAdmin:
		fmt.Println("User access granted")
	case !isLoggedIn:
		fmt.Println("Please log in")
	}

	// Operating system detection
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Running on Windows")
	case "linux", "unix":
		fmt.Println("Running on Unix-like system")
	case "darwin":
		fmt.Println("Running on macOS")
	case "freebsd", "openbsd", "netbsd":
		fmt.Println("Running on BSD system")
	default:
		fmt.Printf("Running on %s\n", runtime.GOOS)
	}

	// File extension handling
	var filename string = "document.pdf"
	var extension string = ""

	// Extract extension
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			extension = filename[i+1:]
			break
		}
	}

	switch extension {
	case "txt", "md", "rtf":
		fmt.Println("Text document")
	case "pdf", "doc", "docx":
		fmt.Println("Document file")
	case "jpg", "jpeg", "png", "gif", "bmp":
		fmt.Println("Image file")
	case "mp3", "wav", "flac", "aac":
		fmt.Println("Audio file")
	case "mp4", "avi", "mkv", "mov":
		fmt.Println("Video file")
	case "zip", "rar", "7z", "tar", "gz":
		fmt.Println("Archive file")
	default:
		fmt.Printf("Unknown file type: %s\n", extension)
	}
}

func demonstrateSwitchWithFallthrough() {
	fmt.Println("\n--- SWITCH WITH FALLTHROUGH ---")

	// Fallthrough example
	var number int = 2

	fmt.Printf("Number %d: ", number)
	switch number {
	case 1:
		fmt.Print("One")
		fallthrough
	case 2:
		fmt.Print("Two")
		fallthrough
	case 3:
		fmt.Print("Three")
		fallthrough
	default:
		fmt.Print(" or more")
	}
	fmt.Println()

	// Fallthrough with conditions
	var score int = 92

	fmt.Printf("Score %d achievements: ", score)
	switch {
	case score >= 95:
		fmt.Print("Perfect ")
		fallthrough
	case score >= 90:
		fmt.Print("Excellent ")
		fallthrough
	case score >= 80:
		fmt.Print("Good ")
		fallthrough
	case score >= 70:
		fmt.Print("Satisfactory ")
		fallthrough
	case score >= 60:
		fmt.Print("Passing")
	default:
		fmt.Print("Needs improvement")
	}
	fmt.Println()

	// Fallthrough with multiple cases
	var dayType string = "weekday"

	switch dayType {
	case "weekend":
		fmt.Print("Weekend activities: ")
		fallthrough
	case "holiday":
		fmt.Print("Relaxing, ")
		fallthrough
	case "weekday":
		fmt.Print("Planning, ")
		fallthrough
	default:
		fmt.Print("Living life")
	}
	fmt.Println()

	// Conditional fallthrough
	var userRole string = "admin"
	var hasSpecialAccess bool = true

	switch userRole {
	case "admin":
		fmt.Print("Admin privileges: Full access")
		if hasSpecialAccess {
			fmt.Print(", Special access")
		}
	case "moderator":
		fmt.Print("Moderator privileges: Moderation tools")
		fallthrough
	case "user":
		fmt.Print(", Basic features")
	default:
		fmt.Print("No access")
	}
	fmt.Println()

	// WARNING: Fallthrough only works to the next case
	// It cannot be used with multiple labels
	fmt.Println("\nNote: fallthrough only affects the immediately following case")

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Fall-through is default (need break to prevent)
	// JavaScript: Can fall through multiple cases automatically
	// Go: No fall-through by default
	// Go: Must explicitly use fallthrough keyword
	// Go: fallthrough only goes to next case
}

func demonstrateSwitchWithoutExpression() {
	fmt.Println("\n--- SWITCH WITHOUT EXPRESSION ---")

	// Switch without expression (acts like if-else chain)
	var temperature int = 25
	var humidity int = 60

	switch {
	case temperature > 30:
		fmt.Println("Too hot")
	case temperature < 10:
		fmt.Println("Too cold")
	case humidity > 80:
		fmt.Println("Too humid")
	case humidity < 20:
		fmt.Println("Too dry")
	default:
		fmt.Println("Perfect weather")
	}

	// Complex conditions in switch
	var age int = 25
	var hasLicense bool = true
	var hasInsurance bool = true

	switch {
	case age < 16:
		fmt.Println("Too young to drive")
	case age >= 16 && age < 18 && hasLicense:
		fmt.Println("Can drive with restrictions")
	case age >= 18 && hasLicense && hasInsurance:
		fmt.Println("Can drive freely")
	case age >= 18 && hasLicense && !hasInsurance:
		fmt.Println("Need insurance to drive")
	case age >= 18 && !hasLicense:
		fmt.Println("Need to get a license")
	default:
		fmt.Println("Cannot determine driving eligibility")
	}

	// String operations in switch
	var username string = "admin"
	var password string = "secret123"

	switch {
	case len(username) < 3:
		fmt.Println("Username too short")
	case len(password) < 8:
		fmt.Println("Password too short")
	case username == "admin" && password == "secret123":
		fmt.Println("Admin login successful")
	case username == "user" && password == "user123":
		fmt.Println("User login successful")
	default:
		fmt.Println("Invalid credentials")
	}

	// Time-based switch
	var hour int = 14

	switch {
	case hour >= 5 && hour < 12:
		fmt.Println("Good morning")
	case hour >= 12 && hour < 17:
		fmt.Println("Good afternoon")
	case hour >= 17 && hour < 21:
		fmt.Println("Good evening")
	default:
		fmt.Println("Good night")
	}

	// Numeric range switch
	var number int = 42

	switch {
	case number < 0:
		fmt.Println("Negative number")
	case number == 0:
		fmt.Println("Zero")
	case number > 0 && number <= 10:
		fmt.Println("Small positive number")
	case number > 10 && number <= 100:
		fmt.Println("Medium positive number")
	case number > 100:
		fmt.Println("Large positive number")
	}

	// Error handling with switch
	var errorCode int = 404
	var errorMessage string = "Not Found"

	switch {
	case errorCode == 0:
		fmt.Println("No error")
	case errorCode >= 400 && errorCode < 500:
		fmt.Printf("Client error %d: %s\n", errorCode, errorMessage)
	case errorCode >= 500:
		fmt.Printf("Server error %d: %s\n", errorCode, errorMessage)
	default:
		fmt.Printf("Unknown error %d: %s\n", errorCode, errorMessage)
	}
}

func demonstrateTypeSwitch() {
	fmt.Println("\n--- TYPE SWITCH ---")

	// Basic type switch
	var value interface{} = 42

	switch v := value.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case float64:
		fmt.Printf("Float: %.2f\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}

	// Type switch with multiple types
	var values []interface{} = []interface{}{
		42,
		"hello",
		3.14,
		true,
		[]int{1, 2, 3},
		map[string]int{"a": 1},
	}

	for i, val := range values {
		fmt.Printf("Value %d: ", i)
		switch v := val.(type) {
		case int:
			fmt.Printf("int = %d\n", v)
		case string:
			fmt.Printf("string = %s\n", v)
		case float64:
			fmt.Printf("float64 = %.2f\n", v)
		case bool:
			fmt.Printf("bool = %t\n", v)
		case []int:
			fmt.Printf("[]int = %v\n", v)
		case map[string]int:
			fmt.Printf("map[string]int = %v\n", v)
		default:
			fmt.Printf("unknown type %T = %v\n", v, v)
		}
	}

	// Type switch with interface methods
	fmt.Println("\nType switch with interface methods:")

	var shapes []Shape = []Shape{
		Circle{radius: 5},
		Rectangle{width: 10, height: 5},
		Triangle{base: 8, height: 6},
	}

	for _, shape := range shapes {
		switch s := shape.(type) {
		case Circle:
			fmt.Printf("Circle with radius %.2f, area = %.2f\n", s.radius, s.Area())
		case Rectangle:
			fmt.Printf("Rectangle %dx%d, area = %.2f\n", int(s.width), int(s.height), s.Area())
		case Triangle:
			fmt.Printf("Triangle base=%.2f height=%.2f, area = %.2f\n", s.base, s.height, s.Area())
		default:
			fmt.Printf("Unknown shape: %T\n", s)
		}
	}

	// Type switch with nil check
	var nilValue interface{} = nil

	switch v := nilValue.(type) {
	case nil:
		fmt.Println("Value is nil")
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	default:
		fmt.Printf("Other type: %T\n", v)
	}

	// Type switch with multiple values
	processValue(42)
	processValue("hello")
	processValue(3.14)
	processValue([]int{1, 2, 3})
	processValue(nil)
}

func demonstrateSwitchWithInitializer() {
	fmt.Println("\n--- SWITCH WITH INITIALIZER ---")

	// Switch with initializer
	switch day := time.Now().Weekday(); day {
	case time.Saturday, time.Sunday:
		fmt.Printf("Weekend: %s\n", day)
	default:
		fmt.Printf("Weekday: %s\n", day)
	}

	// Switch with function call initializer
	switch result := calculateGrade(85); result {
	case "A":
		fmt.Println("Excellent performance")
	case "B":
		fmt.Println("Good performance")
	case "C":
		fmt.Println("Average performance")
	default:
		fmt.Printf("Performance grade: %s\n", result)
	}

	// Switch with multiple variable initialization
	switch x, y := 10, 20; {
	case x > y:
		fmt.Printf("%d is greater than %d\n", x, y)
	case x < y:
		fmt.Printf("%d is less than %d\n", x, y)
	default:
		fmt.Printf("%d equals %d\n", x, y)
	}

	// Switch with error handling initializer
	switch value, err := parseNumber("42"); {
	case err != nil:
		fmt.Printf("Error parsing number: %v\n", err)
	case value < 0:
		fmt.Printf("Negative number: %d\n", value)
	case value == 0:
		fmt.Printf("Zero value\n")
	default:
		fmt.Printf("Positive number: %d\n", value)
	}

	// Switch with map lookup initializer
	switch user, exists := getUserInfo("admin"); {
	case !exists:
		fmt.Println("User not found")
	case user.IsActive:
		fmt.Printf("Active user: %s\n", user.Name)
	default:
		fmt.Printf("Inactive user: %s\n", user.Name)
	}

	// Switch with slice operation initializer
	switch length := len([]int{1, 2, 3, 4, 5}); {
	case length == 0:
		fmt.Println("Empty slice")
	case length == 1:
		fmt.Println("Single element")
	case length < 5:
		fmt.Printf("Small slice: %d elements\n", length)
	default:
		fmt.Printf("Large slice: %d elements\n", length)
	}
}

func demonstrateSwitchBestPractices() {
	fmt.Println("\n--- SWITCH BEST PRACTICES ---")

	// 1. Use switch instead of long if-else chains
	fmt.Println("1. Use switch for multiple conditions:")

	// Good: using switch
	var status int = 200
	switch {
	case status >= 200 && status < 300:
		fmt.Println("Success")
	case status >= 300 && status < 400:
		fmt.Println("Redirection")
	case status >= 400 && status < 500:
		fmt.Println("Client Error")
	case status >= 500:
		fmt.Println("Server Error")
	}

	// 2. Use constants for switch cases
	fmt.Println("\n2. Use constants for switch cases:")

	const (
		StatusPending   = "pending"
		StatusApproved  = "approved"
		StatusRejected  = "rejected"
		StatusCancelled = "cancelled"
	)

	var orderStatus string = StatusApproved

	switch orderStatus {
	case StatusPending:
		fmt.Println("Order is pending")
	case StatusApproved:
		fmt.Println("Order is approved")
	case StatusRejected:
		fmt.Println("Order is rejected")
	case StatusCancelled:
		fmt.Println("Order is cancelled")
	default:
		fmt.Println("Unknown order status")
	}

	// 3. Use type switch for interface handling
	fmt.Println("\n3. Use type switch for interface handling:")

	var data interface{} = "hello world"

	switch v := data.(type) {
	case string:
		fmt.Printf("String length: %d\n", len(v))
	case int:
		fmt.Printf("Integer value: %d\n", v)
	case []int:
		fmt.Printf("Slice length: %d\n", len(v))
	default:
		fmt.Printf("Unsupported type: %T\n", v)
	}

	// 4. Handle default case appropriately
	fmt.Println("\n4. Handle default case appropriately:")

	var command string = "unknown"

	switch command {
	case "start":
		fmt.Println("Starting service")
	case "stop":
		fmt.Println("Stopping service")
	case "restart":
		fmt.Println("Restarting service")
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Available commands: start, stop, restart")
	}

	// 5. Use switch with initializer when appropriate
	fmt.Println("\n5. Use switch with initializer:")

	switch currentHour := time.Now().Hour(); {
	case currentHour < 12:
		fmt.Println("Good morning")
	case currentHour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	// 6. Avoid unnecessary fallthrough
	fmt.Println("\n6. Avoid unnecessary fallthrough:")

	// Good: clear and explicit
	var priority int = 2

	switch priority {
	case 1:
		fmt.Println("High priority")
	case 2:
		fmt.Println("Medium priority")
	case 3:
		fmt.Println("Low priority")
	default:
		fmt.Println("Invalid priority")
	}

	// 7. Use switch for validation
	fmt.Println("\n7. Use switch for validation:")

	if isValidEmail("user@example.com") {
		fmt.Println("Valid email format")
	} else {
		fmt.Println("Invalid email format")
	}

	// 8. Group related cases
	fmt.Println("\n8. Group related cases:")

	var fileExt string = "jpg"

	switch fileExt {
	case "jpg", "jpeg", "png", "gif":
		fmt.Println("Image file")
	case "mp4", "avi", "mkv":
		fmt.Println("Video file")
	case "mp3", "wav", "flac":
		fmt.Println("Audio file")
	default:
		fmt.Println("Unknown file type")
	}

	// BEST PRACTICES SUMMARY:
	fmt.Println("\nBest Practices Summary:")
	fmt.Println("1. Use switch instead of long if-else chains")
	fmt.Println("2. Use constants for switch cases")
	fmt.Println("3. Use type switch for interface handling")
	fmt.Println("4. Always handle the default case")
	fmt.Println("5. Use switch with initializer when appropriate")
	fmt.Println("6. Avoid unnecessary fallthrough")
	fmt.Println("7. Use switch for validation logic")
	fmt.Println("8. Group related cases together")
	fmt.Println("9. Keep cases simple and readable")
	fmt.Println("10. Consider using enums/constants for better maintainability")
}

// Helper types and functions

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.radius * c.radius
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Triangle struct {
	base, height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func processValue(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("Processing integer: %d\n", v)
	case string:
		fmt.Printf("Processing string: %s\n", v)
	case float64:
		fmt.Printf("Processing float: %.2f\n", v)
	case []int:
		fmt.Printf("Processing slice: %v\n", v)
	case nil:
		fmt.Println("Processing nil value")
	default:
		fmt.Printf("Cannot process type: %T\n", v)
	}
}

func calculateGrade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func parseNumber(s string) (int, error) {
	if s == "42" {
		return 42, nil
	}
	return 0, fmt.Errorf("invalid number: %s", s)
}

type User struct {
	Name     string
	IsActive bool
}

func getUserInfo(username string) (User, bool) {
	users := map[string]User{
		"admin": {Name: "Administrator", IsActive: true},
		"user":  {Name: "Regular User", IsActive: false},
	}

	user, exists := users[username]
	return user, exists
}

func isValidEmail(email string) bool {
	switch {
	case len(email) == 0:
		return false
	case len(email) > 254:
		return false
	default:
		// Simple check for @ symbol
		hasAt := false
		for _, char := range email {
			if char == '@' {
				hasAt = true
				break
			}
		}
		return hasAt
	}
}
