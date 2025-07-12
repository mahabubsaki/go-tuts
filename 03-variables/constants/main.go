package main

import (
	"fmt"
	"math"
)

// This file covers constants in Go - a fundamental concept for immutable values
// Constants are different from variables and have special properties

// Package-level constants
const AppName = "Go Learning App"
const Version = "1.0.0"
const MaxUsers = 1000

// Grouped constants
const (
	StatusActive   = "active"
	StatusInactive = "inactive"
	StatusPending  = "pending"
)

// Typed constants
const (
	Pi     float64 = 3.14159265359
	E      float64 = 2.71828182846
	Golden float64 = 1.61803398875
)

// Untyped constants (preferred)
const (
	UntypedPi     = 3.14159265359
	UntypedE      = 2.71828182846
	UntypedInt    = 42
	UntypedString = "Hello"
)

func main() {
	fmt.Println("=== GO CONSTANTS - COMPLETE GUIDE ===")

	demonstrateBasicConstants()
	demonstrateTypedVsUntyped()
	demonstrateConstantExpressions()
	demonstrateIota()
	demonstrateConstantBestPractices()
	demonstrateConstantLimitations()
}

func demonstrateBasicConstants() {
	fmt.Println("\n--- BASIC CONSTANTS ---")

	// Local constants
	const localConst = "I'm a local constant"
	const numberConst = 42
	const boolConst = true
	const floatConst = 3.14

	fmt.Printf("localConst: %s\n", localConst)
	fmt.Printf("numberConst: %d\n", numberConst)
	fmt.Printf("boolConst: %t\n", boolConst)
	fmt.Printf("floatConst: %f\n", floatConst)

	// Constants must be compile-time computable
	const compileTime = 10 + 5                   // OK - compile-time expression
	const stringConcat = "Hello" + " " + "World" // OK - compile-time expression

	fmt.Printf("compileTime: %d\n", compileTime)
	fmt.Printf("stringConcat: %s\n", stringConcat)

	// These would cause compilation errors:
	// const runtimeValue = len("hello")  // ERROR - function call
	// const currentTime = time.Now()     // ERROR - function call
	// const userInput = os.Args[1]       // ERROR - runtime value

	// Accessing package-level constants
	fmt.Printf("AppName: %s\n", AppName)
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("MaxUsers: %d\n", MaxUsers)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: const x = 10; (block-scoped, cannot be reassigned)
	// JavaScript: const can hold runtime values
	// Go: const x = 10; (compile-time constant, must be computable at compile time)
	// Go: const values are more restrictive but more optimized
}

func demonstrateTypedVsUntyped() {
	fmt.Println("\n--- TYPED VS UNTYPED CONSTANTS ---")

	// Typed constants
	const typedInt int = 42
	const typedFloat float64 = 3.14
	const typedString string = "hello"

	fmt.Printf("typedInt: %d (type: %T)\n", typedInt, typedInt)
	fmt.Printf("typedFloat: %f (type: %T)\n", typedFloat, typedFloat)
	fmt.Printf("typedString: %s (type: %T)\n", typedString, typedString)

	// Untyped constants (more flexible)
	const untypedInt = 42
	const untypedFloat = 3.14
	const untypedString = "hello"

	// Untyped constants take the type of the context
	var int32Var int32 = untypedInt       // untypedInt becomes int32
	var float32Var float32 = untypedFloat // untypedFloat becomes float32
	var stringVar string = untypedString  // untypedString becomes string

	fmt.Printf("int32Var: %d (type: %T)\n", int32Var, int32Var)
	fmt.Printf("float32Var: %f (type: %T)\n", float32Var, float32Var)
	fmt.Printf("stringVar: %s (type: %T)\n", stringVar, stringVar)

	// Untyped constants can be used in different contexts
	var differentTypes1 int = untypedInt
	var differentTypes2 int64 = untypedInt
	var differentTypes3 float64 = untypedInt

	fmt.Printf("Same constant, different types:\n")
	fmt.Printf("  int: %d (type: %T)\n", differentTypes1, differentTypes1)
	fmt.Printf("  int64: %d (type: %T)\n", differentTypes2, differentTypes2)
	fmt.Printf("  float64: %f (type: %T)\n", differentTypes3, differentTypes3)

	// High precision untyped constants
	const hugePrecision = 1.234567890123456789012345678901234567890
	var float32FromHuge float32 = hugePrecision // Precision lost but valid
	var float64FromHuge float64 = hugePrecision // Better precision

	fmt.Printf("High precision constant:\n")
	fmt.Printf("  as float32: %f\n", float32FromHuge)
	fmt.Printf("  as float64: %f\n", float64FromHuge)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: const x = 42; (type determined at runtime)
	// JavaScript: No distinction between typed and untyped
	// Go: Untyped constants are more flexible for use in different contexts
	// Go: Typed constants are more restrictive but explicit
}

func demonstrateConstantExpressions() {
	fmt.Println("\n--- CONSTANT EXPRESSIONS ---")

	// Constants can be computed from other constants
	const a = 10
	const b = 20
	const sum = a + b
	const product = a * b
	const quotient = b / a

	fmt.Printf("a: %d, b: %d\n", a, b)
	fmt.Printf("sum: %d\n", sum)
	fmt.Printf("product: %d\n", product)
	fmt.Printf("quotient: %d\n", quotient)

	// String constants
	const firstName = "John"
	const lastName = "Doe"
	const fullName = firstName + " " + lastName

	fmt.Printf("fullName: %s\n", fullName)

	// Boolean constants
	const condition1 = true
	const condition2 = false
	const andResult = condition1 && condition2
	const orResult = condition1 || condition2

	fmt.Printf("andResult: %t\n", andResult)
	fmt.Printf("orResult: %t\n", orResult)

	// Complex expressions
	const complexExpr = (a+b)*2 - a/2
	fmt.Printf("complexExpr: %d\n", complexExpr)

	// Using built-in functions that are compile-time computable
	const stringLength = len("Hello")
	const isLetter = 'A' >= 'A' && 'A' <= 'Z'

	fmt.Printf("stringLength: %d\n", stringLength)
	fmt.Printf("isLetter: %t\n", isLetter)

	// Using math constants in expressions
	const circleArea = math.Pi * 5 * 5 // Using math.Pi constant
	fmt.Printf("circleArea: %f\n", circleArea)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: const x = 10 + 5; (computed at runtime)
	// Go: const x = 10 + 5; (computed at compile time)
	// Go: More restrictions but better performance
}

func demonstrateIota() {
	fmt.Println("\n--- IOTA (CONSTANT GENERATOR) ---")

	// iota is a special identifier used to create enumerated constants
	// It starts at 0 and increments by 1 for each constant in a group

	// Basic iota usage
	const (
		First  = iota // 0
		Second = iota // 1
		Third  = iota // 2
		Fourth = iota // 3
	)

	fmt.Printf("First: %d\n", First)
	fmt.Printf("Second: %d\n", Second)
	fmt.Printf("Third: %d\n", Third)
	fmt.Printf("Fourth: %d\n", Fourth)

	// Simplified iota usage (implicit repetition)
	const (
		Alpha = iota // 0
		Beta         // 1 (implicitly Beta = iota)
		Gamma        // 2 (implicitly Gamma = iota)
		Delta        // 3 (implicitly Delta = iota)
	)

	fmt.Printf("Alpha: %d, Beta: %d, Gamma: %d, Delta: %d\n", Alpha, Beta, Gamma, Delta)

	// iota with expressions
	const (
		KB = 1 << (10 * iota) // 1 << 0 = 1
		MB                    // 1 << 10 = 1024
		GB                    // 1 << 20 = 1048576
		TB                    // 1 << 30 = 1073741824
	)

	fmt.Printf("KB: %d\n", KB)
	fmt.Printf("MB: %d\n", MB)
	fmt.Printf("GB: %d\n", GB)
	fmt.Printf("TB: %d\n", TB)

	// iota with custom starting values
	const (
		StatusCode100 = iota + 100 // 100
		StatusCode101              // 101
		StatusCode102              // 102
	)

	fmt.Printf("StatusCode100: %d\n", StatusCode100)
	fmt.Printf("StatusCode101: %d\n", StatusCode101)
	fmt.Printf("StatusCode102: %d\n", StatusCode102)

	// Skipping values with blank identifier
	const (
		Value1 = iota // 0
		_             // 1 (skipped)
		Value3        // 2
		_             // 3 (skipped)
		Value5        // 4
	)

	fmt.Printf("Value1: %d, Value3: %d, Value5: %d\n", Value1, Value3, Value5)

	// iota resets in each const block
	const (
		ResetA = iota // 0 (starts over)
		ResetB        // 1
	)

	fmt.Printf("ResetA: %d, ResetB: %d\n", ResetA, ResetB)

	// Complex iota expressions
	const (
		PowerOfTwo   = 1 << iota // 1 << 0 = 1
		PowerOfFour  = 1 << iota // 1 << 1 = 2
		PowerOfEight = 1 << iota // 1 << 2 = 4
	)

	fmt.Printf("PowerOfTwo: %d, PowerOfFour: %d, PowerOfEight: %d\n",
		PowerOfTwo, PowerOfFour, PowerOfEight)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: No built-in enum support (use objects or const assertions)
	// JavaScript: const Status = { ACTIVE: 0, INACTIVE: 1, PENDING: 2 };
	// Go: iota provides automatic enumeration
	// Go: More convenient for sequential constant values
}

func demonstrateConstantBestPractices() {
	fmt.Println("\n--- CONSTANT BEST PRACTICES ---")

	// Naming conventions
	const (
		// Use descriptive names
		DefaultTimeout = 30
		MaxRetries     = 3
		DatabaseURL    = "localhost:5432"

		// Use ALL_CAPS for exported constants (optional but common)
		API_VERSION     = "v1"
		MAX_BUFFER_SIZE = 1024
	)

	// Grouping related constants
	const (
		// HTTP status codes
		StatusOK                  = 200
		StatusNotFound            = 404
		StatusInternalServerError = 500
	)

	// Using constants for configuration
	const (
		// Server configuration
		ServerPort    = 8080
		ServerHost    = "localhost"
		ServerTimeout = 30
	)

	// Using constants for magic numbers
	const (
		DaysInWeek      = 7
		HoursInDay      = 24
		MinutesInHour   = 60
		SecondsInMinute = 60
	)

	// Calculate derived constants
	const (
		MinutesInDay = HoursInDay * MinutesInHour
		SecondsInDay = MinutesInDay * SecondsInMinute
	)

	fmt.Printf("Configuration:\n")
	fmt.Printf("  DefaultTimeout: %d\n", DefaultTimeout)
	fmt.Printf("  MaxRetries: %d\n", MaxRetries)
	fmt.Printf("  DatabaseURL: %s\n", DatabaseURL)

	fmt.Printf("Time calculations:\n")
	fmt.Printf("  MinutesInDay: %d\n", MinutesInDay)
	fmt.Printf("  SecondsInDay: %d\n", SecondsInDay)

	// BEST PRACTICES:
	fmt.Println("\nBest Practices:")
	fmt.Println("1. Use constants for values that don't change")
	fmt.Println("2. Group related constants together")
	fmt.Println("3. Use descriptive names")
	fmt.Println("4. Prefer untyped constants for flexibility")
	fmt.Println("5. Use iota for sequential enumerated values")
	fmt.Println("6. Use constants for configuration values")
	fmt.Println("7. Avoid magic numbers - use named constants")
	fmt.Println("8. Export constants that other packages need")
}

func demonstrateConstantLimitations() {
	fmt.Println("\n--- CONSTANT LIMITATIONS ---")

	// Constants can only be basic types
	// These are allowed:
	const validString = "hello"
	const validInt = 42
	const validFloat = 3.14
	const validBool = true
	const validRune = 'A'

	// These are NOT allowed:
	// const invalidSlice = []int{1, 2, 3}        // ERROR
	// const invalidMap = map[string]int{}        // ERROR
	// const invalidStruct = struct{x int}{x: 1}  // ERROR
	// const invalidFunc = func() {}              // ERROR

	// Constants must be compile-time computable
	// These are allowed:
	const compileTimeInt = 10 + 5
	const compileTimeString = "Hello" + " World"
	const compileTimeBool = true && false

	// These are NOT allowed:
	// const runtimeLen = len([]int{1, 2, 3})     // ERROR - function call
	// const runtimeTime = time.Now()             // ERROR - function call
	// var x = 10; const fromVar = x              // ERROR - variable reference

	fmt.Printf("Valid constants:\n")
	fmt.Printf("  validString: %s\n", validString)
	fmt.Printf("  validInt: %d\n", validInt)
	fmt.Printf("  validFloat: %f\n", validFloat)
	fmt.Printf("  validBool: %t\n", validBool)
	fmt.Printf("  validRune: %c\n", validRune)

	// Constants vs variables
	fmt.Println("\nConstants vs Variables:")
	fmt.Println("Constants:")
	fmt.Println("  - Compile-time values")
	fmt.Println("  - Cannot be changed")
	fmt.Println("  - Basic types only")
	fmt.Println("  - Must be computable at compile time")
	fmt.Println("  - More efficient (inlined)")

	fmt.Println("\nVariables:")
	fmt.Println("  - Runtime values")
	fmt.Println("  - Can be changed")
	fmt.Println("  - Any type allowed")
	fmt.Println("  - Can be computed at runtime")
	fmt.Println("  - Stored in memory")

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: const x = [1, 2, 3]; (object/array constants allowed)
	// JavaScript: const can hold complex objects
	// Go: const only for basic types
	// Go: Use var for complex types that don't change
}

// Example of using constants in a real-world scenario
func demonstrateRealWorldExample() {
	fmt.Println("\n--- REAL-WORLD EXAMPLE ---")

	// Application configuration constants
	const (
		AppName    = "User Management System"
		AppVersion = "2.1.0"

		// Database configuration
		DBHost    = "localhost"
		DBPort    = 5432
		DBName    = "users"
		DBTimeout = 30

		// API configuration
		APIPrefix      = "/api/v1"
		APITimeout     = 10
		APIMaxRequests = 1000

		// User roles
		RoleAdmin = "admin"
		RoleUser  = "user"
		RoleGuest = "guest"

		// Status codes
		StatusActive    = 1
		StatusInactive  = 0
		StatusSuspended = -1
	)

	// Using constants in logic
	fmt.Printf("Application: %s v%s\n", AppName, AppVersion)
	fmt.Printf("Database: %s:%d/%s (timeout: %ds)\n", DBHost, DBPort, DBName, DBTimeout)
	fmt.Printf("API: %s (timeout: %ds, max requests: %d)\n", APIPrefix, APITimeout, APIMaxRequests)

	// Using constants for validation
	validRoles := []string{RoleAdmin, RoleUser, RoleGuest}
	fmt.Printf("Valid roles: %v\n", validRoles)

	// Using constants for status checking
	userStatus := StatusActive
	switch userStatus {
	case StatusActive:
		fmt.Println("User is active")
	case StatusInactive:
		fmt.Println("User is inactive")
	case StatusSuspended:
		fmt.Println("User is suspended")
	}
}
