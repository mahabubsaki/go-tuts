package main

import (
	"fmt"
	"strings"
	"time"
)

// === GO NAMING CONVENTIONS COMPREHENSIVE GUIDE ===

/*
NAMING PHILOSOPHY:
- Names should be descriptive but not verbose
- Use MixedCaps (camelCase) for multi-word names
- Short names for short scopes
- Longer names for longer scopes
- Avoid abbreviations unless universally understood

COMPARISON WITH JAVASCRIPT:
// JavaScript: Multiple naming conventions
var userName = "john";        // camelCase
var user_name = "john";       // snake_case
var UserName = "john";        // PascalCase
var username = "john";        // lowercase

// Go: Consistent MixedCaps
var userName = "john"         // unexported (private)
var UserName = "john"         // exported (public)
*/

// 1. PACKAGE NAMES
// Package names should be short, lowercase, single words
// Examples: fmt, http, json, strings, time
// Avoid: myPackage, my_package, MyPackage

// 2. CONSTANTS
// Use MixedCaps for constants, not ALL_CAPS
const (
	MaxRetries     = 3
	DefaultTimeout = 30 * time.Second
	BufferSize     = 1024

	// Avoid: MAX_RETRIES, DEFAULT_TIMEOUT, BUFFER_SIZE
)

// 3. VARIABLES
// Use short names for short scopes, descriptive names for longer scopes
func demonstrateVariableNaming() {
	// Short scope - short names
	for i := 0; i < 10; i++ {
		fmt.Printf("Count: %d\n", i)
	}

	// Medium scope - descriptive names
	userCount := 100
	maxUsers := 1000

	// Long scope - very descriptive names
	authenticatedUserCount := 0
	concurrentConnectionLimit := 100

	fmt.Printf("User count: %d\n", userCount)
	fmt.Printf("Max users: %d\n", maxUsers)
	fmt.Printf("Authenticated users: %d\n", authenticatedUserCount)
	fmt.Printf("Connection limit: %d\n", concurrentConnectionLimit)
}

// 4. FUNCTION NAMES
// Use MixedCaps, start with lowercase for unexported functions
func calculateTotal(items []float64) float64 {
	var total float64
	for _, item := range items {
		total += item
	}
	return total
}

// Exported function - starts with uppercase
func CalculateAverage(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	return calculateTotal(numbers) / float64(len(numbers))
}

// 5. TYPE NAMES
// Use MixedCaps, typically nouns
type User struct {
	ID       int
	Name     string
	Email    string
	Password string // unexported in real code
}

type DatabaseConnection struct {
	Host     string
	Port     int
	Database string
	Username string
	password string // unexported
}

// 6. INTERFACE NAMES
// Single-method interfaces often end with "-er"
type Writer interface {
	Write([]byte) (int, error)
}

type Reader interface {
	Read([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// Multi-method interfaces use descriptive names
type UserRepository interface {
	GetUser(id int) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id int) error
}

// 7. GETTERS AND SETTERS
// Don't use "Get" prefix for getters
type Person struct {
	name string
	age  int
}

// Good: Name(), not GetName()
func (p *Person) Name() string {
	return p.name
}

// Good: SetName(), setters can have "Set" prefix
func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) Age() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	p.age = age
}

// 8. METHOD NAMES
// Use MixedCaps, descriptive action verbs
func (u *User) ValidateEmail() bool {
	return strings.Contains(u.Email, "@") && strings.Contains(u.Email, ".")
}

func (u *User) UpdatePassword(newPassword string) {
	u.Password = newPassword
}

func (u *User) IsAdult() bool {
	// Assuming age is calculated elsewhere
	return true
}

// 9. RECEIVER NAMES
// Use short names, typically first letter or first few letters of type
func (u *User) String() string {
	return fmt.Sprintf("User{ID: %d, Name: %s}", u.ID, u.Name)
}

func (dc *DatabaseConnection) Connect() error {
	// Connection logic
	return nil
}

func (dc *DatabaseConnection) Disconnect() error {
	// Disconnection logic
	return nil
}

// 10. ENUM-LIKE CONSTANTS
// Use the type name as prefix
type Status int

const (
	StatusPending Status = iota
	StatusApproved
	StatusRejected
	StatusCancelled
)

type Priority int

const (
	PriorityLow Priority = iota
	PriorityMedium
	PriorityHigh
	PriorityCritical
)

// 11. ERROR VARIABLES
// Prefix with "Err"
var (
	ErrUserNotFound   = fmt.Errorf("user not found")
	ErrInvalidInput   = fmt.Errorf("invalid input")
	ErrUnauthorized   = fmt.Errorf("unauthorized access")
	ErrInternalServer = fmt.Errorf("internal server error")
)

// 12. ACRONYMS AND INITIALISMS
// Keep them as they are commonly written
type HTTPClient struct {
	URL      string
	APIKey   string
	XMLData  string
	JSONData string
}

func (c *HTTPClient) SendHTTPRequest() error {
	return nil
}

func (c *HTTPClient) ParseXMLResponse() error {
	return nil
}

func (c *HTTPClient) ParseJSONResponse() error {
	return nil
}

// 13. BOOLEAN VARIABLES AND FUNCTIONS
// Use clear, positive names
type Config struct {
	isEnabled    bool
	isProduction bool
	hasFeature   bool
}

func (c *Config) IsEnabled() bool {
	return c.isEnabled
}

func (c *Config) IsProduction() bool {
	return c.isProduction
}

func (c *Config) HasFeature() bool {
	return c.hasFeature
}

// 14. CHANNEL NAMES
// Use descriptive names indicating purpose
func demonstrateChannelNaming() {
	// Good names
	done := make(chan bool)
	results := make(chan string)
	errors := make(chan error)
	userUpdates := make(chan User)

	// Usage example with userUpdates
	go func() {
		userUpdates <- User{ID: 1, Name: "Test User"}
	}()

	select {
	case user := <-userUpdates:
		fmt.Printf("User update: %v\n", user)
	default:
		fmt.Println("No user updates")
	}

	// Usage example
	go func() {
		// Some work
		results <- "task completed"
		done <- true
	}()

	select {
	case result := <-results:
		fmt.Printf("Result: %s\n", result)
	case err := <-errors:
		fmt.Printf("Error: %v\n", err)
	case <-done:
		fmt.Println("Work completed")
	}
}

// 15. NAMING ANTI-PATTERNS TO AVOID
func demonstrateNamingAntiPatterns() {
	// ❌ Bad: Abbreviations
	// var usr User
	// var cfg Config
	// var db Database

	// ✅ Good: Full words
	var user User
	var config Config
	var database DatabaseConnection

	// ❌ Bad: Hungarian notation
	// var strName string
	// var intAge int
	// var boolActive bool

	// ✅ Good: Descriptive names
	var name string
	var age int
	var active bool

	// ❌ Bad: Redundant names
	// var userUser User
	// var stringString string

	// ✅ Good: Clear, concise names
	var currentUser User
	var message string

	fmt.Printf("User: %v\n", user)
	fmt.Printf("Config: %v\n", config)
	fmt.Printf("Database: %v\n", database)
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Active: %t\n", active)
	fmt.Printf("Current user: %v\n", currentUser)
	fmt.Printf("Message: %s\n", message)
}

func main() {
	fmt.Println("=== GO NAMING CONVENTIONS COMPREHENSIVE GUIDE ===")

	// === VARIABLE NAMING ===
	fmt.Println("\n--- VARIABLE NAMING ---")
	demonstrateVariableNaming()

	// === GETTERS AND SETTERS ===
	fmt.Println("\n--- GETTERS AND SETTERS ---")
	person := &Person{}
	person.SetName("John Doe")
	person.SetAge(30)
	fmt.Printf("Person name: %s\n", person.Name())
	fmt.Printf("Person age: %d\n", person.Age())

	// === METHODS ===
	fmt.Println("\n--- METHODS ---")
	user := &User{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
	}
	fmt.Printf("User: %s\n", user.String())
	fmt.Printf("Valid email: %t\n", user.ValidateEmail())

	// === ENUMS ===
	fmt.Println("\n--- ENUMS ---")
	status := StatusPending
	priority := PriorityHigh
	fmt.Printf("Status: %v\n", status)
	fmt.Printf("Priority: %v\n", priority)

	// === CHANNELS ===
	fmt.Println("\n--- CHANNELS ---")
	demonstrateChannelNaming()

	// === ANTI-PATTERNS ===
	fmt.Println("\n--- NAMING ANTI-PATTERNS ---")
	demonstrateNamingAntiPatterns()

	// === NAMING RULES SUMMARY ===
	fmt.Println("\n--- NAMING RULES SUMMARY ---")
	fmt.Println("1. Use MixedCaps (camelCase) for multi-word names")
	fmt.Println("2. Capitalize first letter to export (make public)")
	fmt.Println("3. Use short names for short scopes")
	fmt.Println("4. Use descriptive names for longer scopes")
	fmt.Println("5. Don't use underscores or hyphens")
	fmt.Println("6. Avoid abbreviations unless universally known")
	fmt.Println("7. Use consistent naming patterns")
	fmt.Println("8. Interface names often end with -er")
	fmt.Println("9. Don't use 'Get' prefix for getters")
	fmt.Println("10. Use positive names for booleans")

	// === PACKAGE NAMING ===
	fmt.Println("\n--- PACKAGE NAMING BEST PRACTICES ---")
	fmt.Println("✅ Good: fmt, http, json, strings, time")
	fmt.Println("❌ Bad: myPackage, my_package, MyPackage")
	fmt.Println("✅ Good: user, config, database")
	fmt.Println("❌ Bad: userManager, user_utils, UserStuff")

	fmt.Println("\nGood naming makes Go code self-documenting and maintainable!")
}
