package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// This file covers if-else statements in Go
// Control flow statements determine the execution path of your program

func main() {
	fmt.Println("=== GO IF-ELSE STATEMENTS - COMPLETE GUIDE ===")

	demonstrateBasicIfElse()
	demonstrateIfWithInitializer()
	demonstrateNestedIfElse()
	demonstrateIfElseChaining()
	demonstrateIfElseWithComplexConditions()
	demonstrateIfElseWithDifferentTypes()
	demonstrateIfElseErrorHandling()
	demonstrateIfElseBestPractices()
}

func demonstrateBasicIfElse() {
	fmt.Println("\n--- BASIC IF-ELSE STATEMENTS ---")

	// Basic if statement
	var age int = 25

	if age >= 18 {
		fmt.Printf("Age %d: You are an adult\n", age)
	}

	// If-else statement
	var score int = 85

	if score >= 90 {
		fmt.Printf("Score %d: Excellent!\n", score)
	} else {
		fmt.Printf("Score %d: Good, but can improve\n", score)
	}

	// If-else if-else chain
	var temperature int = 25

	if temperature > 30 {
		fmt.Printf("Temperature %d°C: It's hot\n", temperature)
	} else if temperature > 20 {
		fmt.Printf("Temperature %d°C: It's warm\n", temperature)
	} else if temperature > 10 {
		fmt.Printf("Temperature %d°C: It's cool\n", temperature)
	} else {
		fmt.Printf("Temperature %d°C: It's cold\n", temperature)
	}

	// Boolean variables in if statements
	var isLoggedIn bool = true
	var hasPermission bool = false

	if isLoggedIn {
		fmt.Println("User is logged in")
	} else {
		fmt.Println("User is not logged in")
	}

	if !hasPermission {
		fmt.Println("Access denied: No permission")
	}

	// Combining conditions
	if isLoggedIn && hasPermission {
		fmt.Println("Access granted")
	} else if isLoggedIn && !hasPermission {
		fmt.Println("Logged in but no permission")
	} else {
		fmt.Println("Must login first")
	}

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Same if-else syntax
	// JavaScript: Truthy/falsy values (0, "", null, undefined are falsy)
	// JavaScript: if (variable) checks for truthy
	// Go: Only boolean values in conditions
	// Go: Must explicitly check conditions (if variable != 0)
}

func demonstrateIfWithInitializer() {
	fmt.Println("\n--- IF WITH INITIALIZER ---")

	// If with initializer statement
	// Format: if init; condition { }

	// Example 1: Variable declaration and check
	if num := 42; num > 0 {
		fmt.Printf("Number %d is positive\n", num)
	}
	// Note: 'num' is only available within the if block

	// Example 2: Function call with result check
	if result := calculateSquare(5); result > 20 {
		fmt.Printf("Square result %d is greater than 20\n", result)
	} else {
		fmt.Printf("Square result %d is not greater than 20\n", result)
	}

	// Example 3: String manipulation with check
	if trimmed := strings.TrimSpace("  hello world  "); len(trimmed) > 0 {
		fmt.Printf("Trimmed string: '%s'\n", trimmed)
	}

	// Example 4: Error handling pattern
	if value, err := strconv.Atoi("123"); err == nil {
		fmt.Printf("Parsed integer: %d\n", value)
	} else {
		fmt.Printf("Error parsing: %v\n", err)
	}

	// Example 5: Multiple variable initialization
	if x, y := 10, 20; x+y > 25 {
		fmt.Printf("Sum of %d and %d is greater than 25\n", x, y)
	}

	// Example 6: Complex initialization
	if data := map[string]int{"a": 1, "b": 2}; data["a"] > 0 {
		fmt.Printf("Map contains positive value for 'a': %d\n", data["a"])
	}

	// Scope demonstration
	if localVar := "I'm local"; true {
		fmt.Printf("Inside if block: %s\n", localVar)
	}
	// localVar is not accessible here

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: No initializer syntax in if statements
	// JavaScript: Must declare variables before if statement
	// Go: Can declare and initialize within if statement
	// Go: Variables scoped to if-else block
}

func demonstrateNestedIfElse() {
	fmt.Println("\n--- NESTED IF-ELSE STATEMENTS ---")

	// Nested if-else for complex decision making
	var userType string = "admin"
	var isActive bool = true
	var loginAttempts int = 2

	if userType == "admin" {
		fmt.Println("User is admin")
		if isActive {
			fmt.Println("Admin account is active")
			if loginAttempts < 3 {
				fmt.Println("Access granted to admin")
			} else {
				fmt.Println("Too many login attempts for admin")
			}
		} else {
			fmt.Println("Admin account is suspended")
		}
	} else if userType == "user" {
		fmt.Println("User is regular user")
		if isActive {
			fmt.Println("Regular account is active")
		} else {
			fmt.Println("Regular account is suspended")
		}
	} else {
		fmt.Println("Unknown user type")
	}

	// Grade calculation with nested conditions
	var mathScore int = 85
	var englishScore int = 78
	var scienceScore int = 92

	if mathScore >= 60 && englishScore >= 60 && scienceScore >= 60 {
		fmt.Println("All subjects passed")

		var average int = (mathScore + englishScore + scienceScore) / 3
		if average >= 90 {
			fmt.Printf("Grade: A (Average: %d)\n", average)
		} else if average >= 80 {
			fmt.Printf("Grade: B (Average: %d)\n", average)
		} else if average >= 70 {
			fmt.Printf("Grade: C (Average: %d)\n", average)
		} else {
			fmt.Printf("Grade: D (Average: %d)\n", average)
		}

		// Check for honors
		if mathScore >= 90 || englishScore >= 90 || scienceScore >= 90 {
			fmt.Println("Eligible for honors (90+ in at least one subject)")
		}
	} else {
		fmt.Println("Failed one or more subjects")

		if mathScore < 60 {
			fmt.Printf("Math failed: %d\n", mathScore)
		}
		if englishScore < 60 {
			fmt.Printf("English failed: %d\n", englishScore)
		}
		if scienceScore < 60 {
			fmt.Printf("Science failed: %d\n", scienceScore)
		}
	}

	// Weather decision system
	var temperature int = 22
	var humidity int = 65
	var windSpeed int = 10

	if temperature > 15 {
		fmt.Println("Temperature is comfortable")
		if humidity < 70 {
			fmt.Println("Humidity is acceptable")
			if windSpeed < 20 {
				fmt.Println("Perfect weather for outdoor activities")
			} else {
				fmt.Println("Too windy for some activities")
			}
		} else {
			fmt.Println("Too humid")
		}
	} else {
		fmt.Println("Too cold for outdoor activities")
	}

	// File access permission system
	var fileExists bool = true
	var readPermission bool = true
	var writePermission bool = false
	var isOwner bool = false

	if fileExists {
		fmt.Println("File exists")
		if readPermission {
			fmt.Println("Can read file")
			if writePermission {
				fmt.Println("Can write to file")
			} else {
				fmt.Println("Cannot write to file")
				if isOwner {
					fmt.Println("But you are the owner - contact admin")
				}
			}
		} else {
			fmt.Println("Cannot read file - no permission")
		}
	} else {
		fmt.Println("File does not exist")
	}
}

func demonstrateIfElseChaining() {
	fmt.Println("\n--- IF-ELSE CHAINING ---")

	// Grade system with if-else chain
	var percentage int = 76

	if percentage >= 90 {
		fmt.Printf("Grade A: %d%%\n", percentage)
	} else if percentage >= 80 {
		fmt.Printf("Grade B: %d%%\n", percentage)
	} else if percentage >= 70 {
		fmt.Printf("Grade C: %d%%\n", percentage)
	} else if percentage >= 60 {
		fmt.Printf("Grade D: %d%%\n", percentage)
	} else {
		fmt.Printf("Grade F: %d%%\n", percentage)
	}

	// Day of week determination
	var dayNumber int = 3

	if dayNumber == 1 {
		fmt.Println("Monday")
	} else if dayNumber == 2 {
		fmt.Println("Tuesday")
	} else if dayNumber == 3 {
		fmt.Println("Wednesday")
	} else if dayNumber == 4 {
		fmt.Println("Thursday")
	} else if dayNumber == 5 {
		fmt.Println("Friday")
	} else if dayNumber == 6 {
		fmt.Println("Saturday")
	} else if dayNumber == 7 {
		fmt.Println("Sunday")
	} else {
		fmt.Println("Invalid day number")
	}

	// BMI calculation and categorization
	var weight float64 = 70.0 // kg
	var height float64 = 1.75 // meters
	var bmi float64 = weight / (height * height)

	if bmi < 18.5 {
		fmt.Printf("BMI %.2f: Underweight\n", bmi)
	} else if bmi < 25.0 {
		fmt.Printf("BMI %.2f: Normal weight\n", bmi)
	} else if bmi < 30.0 {
		fmt.Printf("BMI %.2f: Overweight\n", bmi)
	} else {
		fmt.Printf("BMI %.2f: Obese\n", bmi)
	}

	// HTTP status code handling
	var statusCode int = 404

	if statusCode >= 200 && statusCode < 300 {
		fmt.Printf("Success: %d\n", statusCode)
	} else if statusCode >= 300 && statusCode < 400 {
		fmt.Printf("Redirection: %d\n", statusCode)
	} else if statusCode >= 400 && statusCode < 500 {
		fmt.Printf("Client Error: %d\n", statusCode)
	} else if statusCode >= 500 && statusCode < 600 {
		fmt.Printf("Server Error: %d\n", statusCode)
	} else {
		fmt.Printf("Unknown status code: %d\n", statusCode)
	}

	// Age group classification
	var age int = 35

	if age < 0 {
		fmt.Println("Invalid age")
	} else if age < 2 {
		fmt.Printf("Infant (%d years)\n", age)
	} else if age < 13 {
		fmt.Printf("Child (%d years)\n", age)
	} else if age < 20 {
		fmt.Printf("Teenager (%d years)\n", age)
	} else if age < 65 {
		fmt.Printf("Adult (%d years)\n", age)
	} else {
		fmt.Printf("Senior (%d years)\n", age)
	}

	// Credit score rating
	var creditScore int = 720

	if creditScore >= 800 {
		fmt.Printf("Excellent credit: %d\n", creditScore)
	} else if creditScore >= 740 {
		fmt.Printf("Very good credit: %d\n", creditScore)
	} else if creditScore >= 670 {
		fmt.Printf("Good credit: %d\n", creditScore)
	} else if creditScore >= 580 {
		fmt.Printf("Fair credit: %d\n", creditScore)
	} else if creditScore >= 300 {
		fmt.Printf("Poor credit: %d\n", creditScore)
	} else {
		fmt.Printf("Invalid credit score: %d\n", creditScore)
	}
}

func demonstrateIfElseWithComplexConditions() {
	fmt.Println("\n--- IF-ELSE WITH COMPLEX CONDITIONS ---")

	// Multiple conditions with logical operators
	var salary int = 75000
	var experience int = 5
	var hasDegreee bool = true
	var hasSkills bool = true

	if salary >= 50000 && experience >= 3 && (hasDegreee || hasSkills) {
		fmt.Println("Qualified for senior position")
	} else if salary >= 30000 && experience >= 1 {
		fmt.Println("Qualified for junior position")
	} else {
		fmt.Println("Not qualified")
	}

	// Range checking
	var temperature int = 23
	var humidity int = 45

	if temperature >= 20 && temperature <= 26 && humidity >= 30 && humidity <= 60 {
		fmt.Println("Comfortable indoor conditions")
	} else if temperature < 20 || temperature > 26 {
		fmt.Println("Temperature out of comfort range")
	} else if humidity < 30 || humidity > 60 {
		fmt.Println("Humidity out of comfort range")
	}

	// String operations in conditions
	var username string = "admin"
	var password string = "secret123"
	var loginCount int = 0

	if len(username) >= 3 && len(password) >= 8 && loginCount < 3 {
		if username == "admin" && password == "secret123" {
			fmt.Println("Login successful")
		} else {
			fmt.Println("Invalid credentials")
		}
	} else {
		fmt.Println("Login requirements not met")
	}

	// Mathematical conditions
	var x float64 = 4.0
	var y float64 = 3.0

	if math.Abs(x-y) < 0.1 {
		fmt.Printf("%.2f and %.2f are approximately equal\n", x, y)
	} else if x > y {
		fmt.Printf("%.2f is greater than %.2f\n", x, y)
	} else {
		fmt.Printf("%.2f is less than %.2f\n", x, y)
	}

	// Array/slice conditions
	var numbers []int = []int{1, 2, 3, 4, 5}
	var target int = 3

	if len(numbers) > 0 && numbers[0] == target {
		fmt.Printf("First element is %d\n", target)
	} else if len(numbers) > 0 && numbers[len(numbers)-1] == target {
		fmt.Printf("Last element is %d\n", target)
	} else if len(numbers) > 2 && numbers[len(numbers)/2] == target {
		fmt.Printf("Middle element is %d\n", target)
	} else {
		fmt.Printf("Target %d not found in expected positions\n", target)
	}

	// Map operations in conditions
	var userRoles map[string]bool = map[string]bool{
		"admin":  true,
		"editor": false,
		"viewer": true,
	}

	if role, exists := userRoles["admin"]; exists && role {
		fmt.Println("User has admin role")
	} else if role, exists := userRoles["editor"]; exists && role {
		fmt.Println("User has editor role")
	} else if role, exists := userRoles["viewer"]; exists && role {
		fmt.Println("User has viewer role")
	} else {
		fmt.Println("User has no active roles")
	}

	// Time-based conditions
	var hour int = 14
	var isWeekend bool = false
	var isHoliday bool = false

	if hour >= 9 && hour <= 17 && !isWeekend && !isHoliday {
		fmt.Println("Business hours")
	} else if (hour >= 18 && hour <= 22) || (hour >= 6 && hour <= 8) {
		fmt.Println("Extended hours")
	} else if isWeekend || isHoliday {
		fmt.Println("Closed - Weekend or Holiday")
	} else {
		fmt.Println("Closed - Outside business hours")
	}
}

func demonstrateIfElseWithDifferentTypes() {
	fmt.Println("\n--- IF-ELSE WITH DIFFERENT TYPES ---")

	// String comparisons
	var message string = "Hello"

	if message == "Hello" {
		fmt.Println("Greeting detected")
	} else if strings.Contains(message, "error") {
		fmt.Println("Error message detected")
	} else if len(message) == 0 {
		fmt.Println("Empty message")
	} else {
		fmt.Printf("Message: %s\n", message)
	}

	// Floating-point comparisons
	var price float64 = 19.99
	var discount float64 = 0.1

	if price > 20.0 {
		fmt.Printf("Price $%.2f is above $20\n", price)
	} else if price >= 10.0 {
		fmt.Printf("Price $%.2f qualifies for discount\n", price)
		if discount > 0 {
			var finalPrice float64 = price * (1 - discount)
			fmt.Printf("Final price with discount: $%.2f\n", finalPrice)
		}
	} else {
		fmt.Printf("Price $%.2f is below $10\n", price)
	}

	// Byte and rune comparisons
	var byteValue byte = 65 // ASCII 'A'
	var runeValue rune = 'A'

	if byteValue >= 65 && byteValue <= 90 {
		fmt.Printf("Byte %d is uppercase letter (%c)\n", byteValue, byteValue)
	}

	if runeValue >= 'A' && runeValue <= 'Z' {
		fmt.Printf("Rune %c is uppercase letter\n", runeValue)
	}

	// Pointer comparisons
	var number int = 42
	var ptr *int = &number
	var nilPtr *int // nil by default

	if ptr != nil {
		fmt.Printf("Pointer points to value: %d\n", *ptr)
	}

	if nilPtr == nil {
		fmt.Println("Nil pointer detected")
	}

	// Interface comparisons
	var value interface{} = 42

	if intValue, ok := value.(int); ok {
		fmt.Printf("Value is integer: %d\n", intValue)
	} else if stringValue, ok := value.(string); ok {
		fmt.Printf("Value is string: %s\n", stringValue)
	} else {
		fmt.Printf("Value is unknown type: %T\n", value)
	}

	// Array comparisons
	var array1 [3]int = [3]int{1, 2, 3}
	var array2 [3]int = [3]int{1, 2, 3}
	var array3 [3]int = [3]int{1, 2, 4}

	if array1 == array2 {
		fmt.Println("Arrays are equal")
	}

	if array1 != array3 {
		fmt.Println("Arrays are different")
	}

	// Struct comparisons
	type Person struct {
		Name string
		Age  int
	}

	var person1 Person = Person{Name: "Alice", Age: 30}
	var person2 Person = Person{Name: "Alice", Age: 30}
	var person3 Person = Person{Name: "Bob", Age: 25}

	if person1 == person2 {
		fmt.Println("Persons are the same")
	}

	if person1 != person3 {
		fmt.Println("Persons are different")
	}
}

func demonstrateIfElseErrorHandling() {
	fmt.Println("\n--- IF-ELSE ERROR HANDLING ---")

	// Basic error handling pattern
	if value, err := strconv.Atoi("42"); err != nil {
		fmt.Printf("Error converting string to int: %v\n", err)
	} else {
		fmt.Printf("Successfully converted: %d\n", value)
	}

	// Multiple error checks
	if floatValue, err := strconv.ParseFloat("3.14", 64); err != nil {
		fmt.Printf("Error parsing float: %v\n", err)
	} else if floatValue < 0 {
		fmt.Printf("Negative value not allowed: %.2f\n", floatValue)
	} else if floatValue > 10 {
		fmt.Printf("Value too large: %.2f\n", floatValue)
	} else {
		fmt.Printf("Valid float value: %.2f\n", floatValue)
	}

	// Error handling with early return pattern
	fmt.Println("Processing user input...")
	processUserInput("25")
	processUserInput("invalid")
	processUserInput("-5")

	// Validation with multiple checks
	validateUser("john_doe", "password123", "john@example.com")
	validateUser("ab", "pass", "invalid-email")

	// Error handling with custom error types
	if result, err := divide(10, 2); err != nil {
		fmt.Printf("Division error: %v\n", err)
	} else {
		fmt.Printf("Division result: %.2f\n", result)
	}

	if result, err := divide(10, 0); err != nil {
		fmt.Printf("Division error: %v\n", err)
	} else {
		fmt.Printf("Division result: %.2f\n", result)
	}

	// Nested error handling
	if data, err := parseAndValidateData("123"); err != nil {
		fmt.Printf("Data processing error: %v\n", err)
	} else {
		fmt.Printf("Processed data: %d\n", data)
	}

	// Error handling with cleanup
	fmt.Println("File processing simulation...")
	processFile("valid_file.txt")
	processFile("invalid_file.txt")
}

func demonstrateIfElseBestPractices() {
	fmt.Println("\n--- IF-ELSE BEST PRACTICES ---")

	// 1. Early return pattern
	fmt.Println("1. Early return pattern:")
	checkAccessGood("admin", true)
	checkAccessGood("user", false)

	// 2. Avoid deep nesting
	fmt.Println("\n2. Avoid deep nesting:")
	var score int = 85
	var attendance float64 = 0.9
	var assignments int = 8

	// Good: flat structure with early returns
	if score < 60 {
		fmt.Printf("Failed: Low score (%d)\n", score)
		return
	}
	if attendance < 0.8 {
		fmt.Printf("Failed: Low attendance (%.1f%%)\n", attendance*100)
		return
	}
	if assignments < 7 {
		fmt.Printf("Failed: Insufficient assignments (%d)\n", assignments)
		return
	}
	fmt.Printf("Passed: Score=%d, Attendance=%.1f%%, Assignments=%d\n",
		score, attendance*100, assignments)

	// 3. Use descriptive variable names
	fmt.Println("\n3. Use descriptive variable names:")
	var userAge int = 25
	var minimumAge int = 18
	var hasValidID bool = true
	var hasConsent bool = true

	if userAge >= minimumAge && hasValidID && hasConsent {
		fmt.Println("Eligible for service")
	} else {
		fmt.Println("Not eligible for service")
	}

	// 4. Extract complex conditions
	fmt.Println("\n4. Extract complex conditions:")
	var temperature int = 22
	var humidity int = 55
	var windSpeed int = 10

	var isTemperatureComfortable bool = temperature >= 20 && temperature <= 26
	var isHumidityOk bool = humidity >= 30 && humidity <= 60
	var isWindSpeedAcceptable bool = windSpeed < 15

	if isTemperatureComfortable && isHumidityOk && isWindSpeedAcceptable {
		fmt.Println("Perfect weather conditions")
	} else {
		fmt.Println("Weather conditions not ideal")
	}

	// 5. Use constants for magic numbers
	fmt.Println("\n5. Use constants for magic numbers:")
	const (
		PassingGrade     = 60
		ExcellentGrade   = 90
		MinimumAge       = 18
		MaxLoginAttempts = 3
	)

	var studentGrade int = 85
	var studentAge int = 20
	var loginAttempts int = 1

	if studentGrade >= ExcellentGrade {
		fmt.Printf("Excellent grade: %d\n", studentGrade)
	} else if studentGrade >= PassingGrade {
		fmt.Printf("Passing grade: %d\n", studentGrade)
	} else {
		fmt.Printf("Failing grade: %d\n", studentGrade)
	}

	if studentAge >= MinimumAge && loginAttempts < MaxLoginAttempts {
		fmt.Println("Access granted")
	}

	// 6. Handle edge cases
	fmt.Println("\n6. Handle edge cases:")
	handleEdgeCases([]int{1, 2, 3, 4, 5})
	handleEdgeCases([]int{})
	handleEdgeCases(nil)

	// BEST PRACTICES SUMMARY:
	fmt.Println("\nBest Practices Summary:")
	fmt.Println("1. Use early returns to reduce nesting")
	fmt.Println("2. Avoid deep nesting - keep conditions flat")
	fmt.Println("3. Use descriptive variable names for conditions")
	fmt.Println("4. Extract complex conditions into named variables")
	fmt.Println("5. Use constants instead of magic numbers")
	fmt.Println("6. Handle edge cases and error conditions")
	fmt.Println("7. Use initializer syntax when appropriate")
	fmt.Println("8. Keep conditions simple and readable")
	fmt.Println("9. Use logical operators appropriately")
	fmt.Println("10. Consider using switch for multiple conditions")
}

// Helper functions
func calculateSquare(n int) int {
	return n * n
}

func processUserInput(input string) {
	value, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Invalid input '%s': %v\n", input, err)
		return
	}

	if value < 0 {
		fmt.Printf("Negative value not allowed: %d\n", value)
		return
	}

	if value > 100 {
		fmt.Printf("Value too large: %d\n", value)
		return
	}

	fmt.Printf("Valid input processed: %d\n", value)
}

func validateUser(username, password, email string) {
	fmt.Printf("Validating user: %s\n", username)

	if len(username) < 3 {
		fmt.Println("Username too short")
		return
	}

	if len(password) < 8 {
		fmt.Println("Password too short")
		return
	}

	if !strings.Contains(email, "@") {
		fmt.Println("Invalid email format")
		return
	}

	fmt.Println("User validation successful")
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func parseAndValidateData(input string) (int, error) {
	if value, err := strconv.Atoi(input); err != nil {
		return 0, fmt.Errorf("parsing error: %v", err)
	} else if value < 0 {
		return 0, fmt.Errorf("negative values not allowed")
	} else if value > 1000 {
		return 0, fmt.Errorf("value too large")
	} else {
		return value, nil
	}
}

func processFile(filename string) {
	fmt.Printf("Processing file: %s\n", filename)

	if !strings.HasSuffix(filename, ".txt") {
		fmt.Println("Error: Only .txt files are supported")
		return
	}

	if strings.Contains(filename, "invalid") {
		fmt.Println("Error: Invalid file detected")
		return
	}

	fmt.Println("File processed successfully")
}

func checkAccessGood(role string, isActive bool) {
	fmt.Printf("Checking access for %s (active: %t)\n", role, isActive)

	if !isActive {
		fmt.Println("Access denied: Account not active")
		return
	}

	if role != "admin" && role != "user" {
		fmt.Println("Access denied: Invalid role")
		return
	}

	if role == "admin" {
		fmt.Println("Access granted: Admin privileges")
		return
	}

	fmt.Println("Access granted: User privileges")
}

func handleEdgeCases(numbers []int) {
	fmt.Printf("Processing numbers: %v\n", numbers)

	if numbers == nil {
		fmt.Println("Error: nil slice provided")
		return
	}

	if len(numbers) == 0 {
		fmt.Println("Warning: empty slice provided")
		return
	}

	if len(numbers) == 1 {
		fmt.Printf("Single element: %d\n", numbers[0])
		return
	}

	fmt.Printf("Multiple elements: first=%d, last=%d\n", numbers[0], numbers[len(numbers)-1])
}
