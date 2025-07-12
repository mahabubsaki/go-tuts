package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// === ERROR HANDLING IN GO ===

// Custom error types
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error in field '%s': %s", e.Field, e.Message)
}

type NetworkError struct {
	Operation string
	URL       string
	Err       error
}

func (e NetworkError) Error() string {
	return fmt.Sprintf("network error during %s to %s: %v", e.Operation, e.URL, e.Err)
}

func (e NetworkError) Unwrap() error {
	return e.Err
}

// Multiple error type
type MultiError struct {
	Errors []error
}

func (e MultiError) Error() string {
	var msgs []string
	for _, err := range e.Errors {
		msgs = append(msgs, err.Error())
	}
	return fmt.Sprintf("multiple errors: %s", strings.Join(msgs, "; "))
}

// Service error with code
type ServiceError struct {
	Code    int
	Message string
	Cause   error
}

func (e ServiceError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("service error [%d]: %s (caused by: %v)", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("service error [%d]: %s", e.Code, e.Message)
}

func (e ServiceError) Unwrap() error {
	return e.Cause
}

// User struct for demonstration
type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

// Validate user data
func (u User) Validate() error {
	if u.Name == "" {
		return ValidationError{Field: "Name", Message: "cannot be empty"}
	}
	if u.Email == "" {
		return ValidationError{Field: "Email", Message: "cannot be empty"}
	}
	if !strings.Contains(u.Email, "@") {
		return ValidationError{Field: "Email", Message: "must contain @"}
	}
	if u.Age < 0 {
		return ValidationError{Field: "Age", Message: "cannot be negative"}
	}
	if u.Age > 150 {
		return ValidationError{Field: "Age", Message: "cannot be greater than 150"}
	}
	return nil
}

// Database simulation
func saveUser(user User) error {
	// Simulate validation
	if err := user.Validate(); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	// Simulate database error
	if user.ID == 0 {
		return ServiceError{
			Code:    500,
			Message: "database connection failed",
			Cause:   errors.New("connection timeout"),
		}
	}

	// Simulate success
	return nil
}

// File operations
func readConfig(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read config file %s: %w", filename, err)
	}
	return string(data), nil
}

func writeConfig(filename, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file %s: %w", filename, err)
	}
	return nil
}

// Network simulation
func fetchData(url string) (string, error) {
	// Simulate network errors
	if url == "" {
		return "", NetworkError{
			Operation: "fetch",
			URL:       url,
			Err:       errors.New("empty URL"),
		}
	}

	if strings.Contains(url, "timeout") {
		return "", NetworkError{
			Operation: "fetch",
			URL:       url,
			Err:       errors.New("connection timeout"),
		}
	}

	if strings.Contains(url, "notfound") {
		return "", NetworkError{
			Operation: "fetch",
			URL:       url,
			Err:       errors.New("404 not found"),
		}
	}

	return "data from " + url, nil
}

// Parse integer with context
func parseIntWithContext(s string, context string) (int, error) {
	value, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %s as integer: %w", context, err)
	}
	return value, nil
}

// Division with error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Safe division that returns zero on error
func safeDivide(a, b float64) float64 {
	result, err := divide(a, b)
	if err != nil {
		return 0
	}
	return result
}

// Batch operations
func processUsers(users []User) error {
	var errs []error

	for i, user := range users {
		if err := saveUser(user); err != nil {
			errs = append(errs, fmt.Errorf("user %d: %w", i, err))
		}
	}

	if len(errs) > 0 {
		return MultiError{Errors: errs}
	}

	return nil
}

// Resource cleanup example
func processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close() // Ensure file is closed even if error occurs

	// Simulate file processing
	buffer := make([]byte, 1024)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		return fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	return nil
}

// Panic and recover example
func riskyOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	// This will panic
	panic("something went wrong!")
}

// Safe operation wrapper
func safeOperation(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()

	f()
	return nil
}

// Error with stack trace simulation
type StackError struct {
	Message string
	Stack   []string
}

func (e StackError) Error() string {
	return fmt.Sprintf("%s\nStack: %v", e.Message, e.Stack)
}

// Error checking helper
func checkError(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// Error logging helper
func logError(operation string, err error) {
	if err != nil {
		fmt.Printf("[ERROR] %s failed: %v\n", operation, err)
	}
}

func main() {
	fmt.Println("=== GO ERROR HANDLING COMPREHENSIVE GUIDE ===")

	// === BASIC ERROR HANDLING ===
	fmt.Println("\n--- BASIC ERROR HANDLING ---")

	// Simple error checking
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	// Division by zero
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", result)
	}

	/*
		JavaScript comparison:
		// JavaScript uses try-catch for error handling
		try {
			const result = 10 / 0; // Returns Infinity, doesn't throw
			console.log(result);
		} catch (error) {
			console.log("Error:", error.message);
		}

		// JavaScript error throwing
		function divide(a, b) {
			if (b === 0) {
				throw new Error("Division by zero");
			}
			return a / b;
		}

		try {
			const result = divide(10, 0);
		} catch (error) {
			console.log("Error:", error.message);
		}
	*/

	// === ERROR CREATION ===
	fmt.Println("\n--- ERROR CREATION ---")

	// Using errors.New
	err1 := errors.New("simple error message")
	fmt.Printf("Error 1: %v\n", err1)

	// Using fmt.Errorf
	err2 := fmt.Errorf("formatted error: %s", "something went wrong")
	fmt.Printf("Error 2: %v\n", err2)

	// Error with context
	username := "john"
	err3 := fmt.Errorf("user %s not found", username)
	fmt.Printf("Error 3: %v\n", err3)

	// === CUSTOM ERROR TYPES ===
	fmt.Println("\n--- CUSTOM ERROR TYPES ---")

	// Validation error
	user := User{
		Name:  "", // Invalid: empty name
		Email: "invalid-email",
		Age:   -5,
	}

	err = user.Validate()
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)

		// Type assertion to get specific error details
		if ve, ok := err.(ValidationError); ok {
			fmt.Printf("Field: %s, Message: %s\n", ve.Field, ve.Message)
		}
	}

	// === ERROR WRAPPING ===
	fmt.Println("\n--- ERROR WRAPPING ---")

	// Wrap errors for context
	baseErr := errors.New("connection failed")
	wrappedErr := fmt.Errorf("failed to connect to database: %w", baseErr)
	doubleWrappedErr := fmt.Errorf("user service unavailable: %w", wrappedErr)

	fmt.Printf("Original error: %v\n", baseErr)
	fmt.Printf("Wrapped error: %v\n", wrappedErr)
	fmt.Printf("Double wrapped: %v\n", doubleWrappedErr)

	// Unwrap errors
	fmt.Printf("Unwrapped: %v\n", errors.Unwrap(wrappedErr))

	// Check if error is of specific type
	if errors.Is(doubleWrappedErr, baseErr) {
		fmt.Println("Double wrapped error contains base error")
	}

	// === ERROR CHECKING PATTERNS ===
	fmt.Println("\n--- ERROR CHECKING PATTERNS ---")

	// Early return pattern
	func() {
		value, err := parseIntWithContext("abc", "user input")
		if err != nil {
			fmt.Printf("Parse error: %v\n", err)
			return
		}
		fmt.Printf("Parsed value: %d\n", value)
	}()

	// Multiple error checks
	func() {
		val1, err := parseIntWithContext("10", "first number")
		if err != nil {
			fmt.Printf("Error parsing first number: %v\n", err)
			return
		}

		val2, err := parseIntWithContext("20", "second number")
		if err != nil {
			fmt.Printf("Error parsing second number: %v\n", err)
			return
		}

		fmt.Printf("Sum: %d\n", val1+val2)
	}()

	// === ERROR HANDLING WITH DEFER ===
	fmt.Println("\n--- ERROR HANDLING WITH DEFER ---")

	// File operations with defer
	func() {
		fmt.Println("Attempting to process non-existent file...")
		err := processFile("nonexistent.txt")
		if err != nil {
			fmt.Printf("File processing error: %v\n", err)
		}
	}()

	// === NETWORK ERROR HANDLING ===
	fmt.Println("\n--- NETWORK ERROR HANDLING ---")

	urls := []string{
		"http://example.com",
		"",
		"http://timeout.com",
		"http://notfound.com",
	}

	for _, url := range urls {
		data, err := fetchData(url)
		if err != nil {
			fmt.Printf("Failed to fetch %s: %v\n", url, err)

			// Check specific error type
			var netErr NetworkError
			if errors.As(err, &netErr) {
				fmt.Printf("  Network operation: %s\n", netErr.Operation)
				fmt.Printf("  URL: %s\n", netErr.URL)
			}
		} else {
			fmt.Printf("Successfully fetched: %s\n", data)
		}
	}

	// === MULTIPLE ERROR HANDLING ===
	fmt.Println("\n--- MULTIPLE ERROR HANDLING ---")

	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 30},
		{ID: 0, Name: "Bob", Email: "bob@example.com", Age: 25},  // Will fail (ID=0)
		{ID: 2, Name: "", Email: "charlie@example.com", Age: 35}, // Will fail (empty name)
		{ID: 3, Name: "David", Email: "david@example.com", Age: 40},
	}

	err = processUsers(users)
	if err != nil {
		fmt.Printf("Batch processing errors: %v\n", err)
	}

	// === PANIC AND RECOVER ===
	fmt.Println("\n--- PANIC AND RECOVER ---")

	// Demonstrate panic recovery
	fmt.Println("Calling risky operation...")
	riskyOperation()
	fmt.Println("Continued after panic recovery")

	// Safe operation wrapper
	err = safeOperation(func() {
		panic("controlled panic")
	})
	if err != nil {
		fmt.Printf("Safe operation error: %v\n", err)
	}

	// === SENTINEL ERRORS ===
	fmt.Println("\n--- SENTINEL ERRORS ---")

	// Define sentinel errors
	var (
		ErrUserNotFound = errors.New("user not found")
		ErrInvalidInput = errors.New("invalid input")
	)

	// Function that returns sentinel errors
	findUser := func(id int) (*User, error) {
		if id <= 0 {
			return nil, ErrInvalidInput
		}
		if id == 999 {
			return nil, ErrUserNotFound
		}
		return &User{ID: id, Name: "User" + fmt.Sprint(id)}, nil
	}

	// Check for specific sentinel errors
	foundUser, err := findUser(999)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			fmt.Println("User not found - expected error")
		} else if errors.Is(err, ErrInvalidInput) {
			fmt.Println("Invalid input provided")
		} else {
			fmt.Printf("Unexpected error: %v\n", err)
		}
	} else {
		fmt.Printf("Found user: %+v\n", foundUser)
	}

	// === ERROR HANDLING PATTERNS ===
	fmt.Println("\n--- ERROR HANDLING PATTERNS ---")

	// 1. Default value pattern
	safeValue := safeDivide(10, 0)
	fmt.Printf("Safe division result: %.2f\n", safeValue)

	// 2. Retry pattern
	retryOperation := func(maxRetries int) error {
		for i := 0; i < maxRetries; i++ {
			_, err := fetchData("http://timeout.com")
			if err == nil {
				return nil
			}
			fmt.Printf("Attempt %d failed: %v\n", i+1, err)
		}
		return fmt.Errorf("operation failed after %d retries", maxRetries)
	}

	err = retryOperation(3)
	if err != nil {
		fmt.Printf("Retry pattern error: %v\n", err)
	}

	// 3. Circuit breaker pattern (simplified)
	type CircuitBreaker struct {
		failures  int
		threshold int
		open      bool
	}

	cb := &CircuitBreaker{threshold: 3}

	callWithCircuitBreaker := func(cb *CircuitBreaker) error {
		if cb.open {
			return errors.New("circuit breaker is open")
		}

		// Simulate operation
		_, err := fetchData("http://timeout.com")
		if err != nil {
			cb.failures++
			if cb.failures >= cb.threshold {
				cb.open = true
				fmt.Println("Circuit breaker opened due to failures")
			}
			return err
		}

		cb.failures = 0
		return nil
	}

	for i := 0; i < 5; i++ {
		err := callWithCircuitBreaker(cb)
		if err != nil {
			fmt.Printf("Circuit breaker call %d failed: %v\n", i+1, err)
		}
	}

	// === ERROR LOGGING ===
	fmt.Println("\n--- ERROR LOGGING ---")

	// Log errors with context
	logError("user creation", saveUser(User{ID: 0, Name: "Test"}))
	logError("division", func() error {
		_, err := divide(10, 0)
		return err
	}())

	// === ERROR HANDLING HELPERS ===
	fmt.Println("\n--- ERROR HANDLING HELPERS ---")

	// Check and handle errors
	checkError(errors.New("sample error"))
	checkError(nil)

	// Must functions (panic on error)
	mustParse := func(s string) int {
		value, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("failed to parse %s: %v", s, err))
		}
		return value
	}

	// Use must function with recovery
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Must function panicked: %v\n", r)
			}
		}()

		value := mustParse("abc") // Will panic
		fmt.Printf("Parsed value: %d\n", value)
	}()

	// === CONTEXT WITH ERRORS ===
	fmt.Println("\n--- CONTEXT WITH ERRORS ---")

	createStackError := func() error {
		return StackError{
			Message: "operation failed",
			Stack:   []string{"main.go:123", "handler.go:45", "service.go:78"},
		}
	}

	stackErr := createStackError()
	fmt.Printf("Stack error: %v\n", stackErr)

	// === BEST PRACTICES ===
	fmt.Println("\n--- ERROR HANDLING BEST PRACTICES ---")
	fmt.Println("1. Always handle errors explicitly")
	fmt.Println("2. Use descriptive error messages")
	fmt.Println("3. Wrap errors to add context")
	fmt.Println("4. Use custom error types for specific errors")
	fmt.Println("5. Use sentinel errors for expected conditions")
	fmt.Println("6. Don't ignore errors (use _ if intentional)")
	fmt.Println("7. Use defer for cleanup")
	fmt.Println("8. Panic only for truly exceptional cases")
	fmt.Println("9. Use errors.Is and errors.As for error checking")
	fmt.Println("10. Log errors with sufficient context")

	// === COMMON ANTI-PATTERNS ===
	fmt.Println("\n--- COMMON ANTI-PATTERNS TO AVOID ---")
	fmt.Println("❌ Ignoring errors: _, err := someOperation()")
	fmt.Println("❌ Panic for recoverable errors")
	fmt.Println("❌ Using panic for control flow")
	fmt.Println("❌ Not wrapping errors with context")
	fmt.Println("❌ Using string comparison for error checking")
	fmt.Println("❌ Not cleaning up resources on error")
	fmt.Println("❌ Generic error messages without context")
	fmt.Println("❌ Swallowing errors without logging")

	fmt.Println("\nProper error handling makes Go programs robust and maintainable!")
}
