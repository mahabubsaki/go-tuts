package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// === GO INIT FUNCTION COMPREHENSIVE GUIDE ===

/*
INIT FUNCTION PHILOSOPHY:
- init() functions are called automatically before main()
- Used for package initialization and setup
- Can have multiple init() functions in same package
- Execute in order of appearance and import dependency
- Cannot be called explicitly

COMPARISON WITH JAVASCRIPT:
// JavaScript - No direct equivalent, but similar concepts:
// IIFE (Immediately Invoked Function Expression)
(function() {
  console.log("Initialization code");
  // Setup code here
})();

// Module initialization
const config = initializeConfig();

// Go - init function
func init() {
  fmt.Println("Initialization code")
  // Setup code here
}
*/

// === PACKAGE-LEVEL VARIABLES ===

var (
	// These will be initialized before init() functions
	startTime = time.Now()
	config    Configuration
	database  *Database
	logger    *Logger
)

// === CONFIGURATION STRUCT ===

type Configuration struct {
	AppName     string
	Version     string
	Environment string
	Port        int
	Debug       bool
}

// === DATABASE STRUCT ===

type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

func (db *Database) Connect() error {
	fmt.Printf("Connecting to database: %s:%d/%s\n", db.Host, db.Port, db.Name)
	return nil
}

func (db *Database) Close() error {
	fmt.Printf("Closing database connection\n")
	return nil
}

// === LOGGER STRUCT ===

type Logger struct {
	Level  string
	Output string
}

func (l *Logger) Log(message string) {
	fmt.Printf("[%s] %s: %s\n", l.Level, time.Now().Format("15:04:05"), message)
}

// === FIRST INIT FUNCTION ===

func init() {
	fmt.Println("=== FIRST INIT FUNCTION ===")
	fmt.Println("This is the first init function")

	// Initialize configuration
	config = Configuration{
		AppName:     "MyGoApp",
		Version:     "1.0.0",
		Environment: "development",
		Port:        8080,
		Debug:       true,
	}

	fmt.Printf("Configuration initialized: %+v\n", config)
}

// === SECOND INIT FUNCTION ===

func init() {
	fmt.Println("=== SECOND INIT FUNCTION ===")
	fmt.Println("This is the second init function")

	// Initialize database
	database = &Database{
		Host:     "localhost",
		Port:     5432,
		Username: "admin",
		Password: "password",
		Name:     "myapp_db",
	}

	fmt.Printf("Database configuration initialized: %+v\n", database)
}

// === THIRD INIT FUNCTION ===

func init() {
	fmt.Println("=== THIRD INIT FUNCTION ===")
	fmt.Println("This is the third init function")

	// Initialize logger
	logger = &Logger{
		Level:  "INFO",
		Output: "stdout",
	}

	fmt.Printf("Logger initialized: %+v\n", logger)
	logger.Log("Logger is ready")
}

// === FOURTH INIT FUNCTION - ENVIRONMENT SETUP ===

func init() {
	fmt.Println("=== FOURTH INIT FUNCTION - ENVIRONMENT SETUP ===")

	// Check environment variables
	if env := os.Getenv("APP_ENV"); env != "" {
		config.Environment = env
		fmt.Printf("Environment set from ENV: %s\n", env)
	}

	if port := os.Getenv("APP_PORT"); port != "" {
		// In real app, you'd parse this properly
		fmt.Printf("Port override from ENV: %s\n", port)
	}

	// Set debug mode based on environment
	if config.Environment == "production" {
		config.Debug = false
		logger.Level = "ERROR"
	} else {
		config.Debug = true
		logger.Level = "DEBUG"
	}

	fmt.Printf("Final configuration: %+v\n", config)
}

// === FIFTH INIT FUNCTION - RESOURCE INITIALIZATION ===

func init() {
	fmt.Println("=== FIFTH INIT FUNCTION - RESOURCE INITIALIZATION ===")

	// Initialize resources that depend on configuration
	if config.Debug {
		logger.Log("Debug mode enabled")
	}

	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	logger.Log("All resources initialized successfully")
}

// === SIXTH INIT FUNCTION - VALIDATION ===

func init() {
	fmt.Println("=== SIXTH INIT FUNCTION - VALIDATION ===")

	// Validate configuration
	if config.AppName == "" {
		log.Fatal("App name is required")
	}

	if config.Port <= 0 || config.Port > 65535 {
		log.Fatal("Invalid port number")
	}

	// Validate database connection
	if database == nil {
		log.Fatal("Database not initialized")
	}

	// Validate logger
	if logger == nil {
		log.Fatal("Logger not initialized")
	}

	logger.Log("All validations passed")
}

// === MAIN FUNCTION ===

func main() {
	fmt.Println("=== GO INIT FUNCTION COMPREHENSIVE GUIDE ===")

	// Calculate initialization time
	initDuration := time.Since(startTime)
	fmt.Printf("Initialization took: %v\n", initDuration)

	// === DEMONSTRATING INITIALIZATION RESULTS ===
	fmt.Println("\n1. INITIALIZATION RESULTS:")

	logger.Log("Application starting")
	fmt.Printf("App Name: %s\n", config.AppName)
	fmt.Printf("Version: %s\n", config.Version)
	fmt.Printf("Environment: %s\n", config.Environment)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Debug Mode: %t\n", config.Debug)

	// === USING INITIALIZED RESOURCES ===
	fmt.Println("\n2. USING INITIALIZED RESOURCES:")

	// Use logger
	logger.Log("Processing user request")
	logger.Log("Database query executed")
	logger.Log("Response sent to client")

	// Simulate application work
	fmt.Println("\n3. SIMULATING APPLICATION WORK:")

	// Simulate some work
	for i := 1; i <= 5; i++ {
		logger.Log(fmt.Sprintf("Processing item %d", i))
		time.Sleep(100 * time.Millisecond)
	}

	// === INIT FUNCTION CHARACTERISTICS ===
	fmt.Println("\n4. INIT FUNCTION CHARACTERISTICS:")

	fmt.Println("✓ init() functions run before main()")
	fmt.Println("✓ Multiple init() functions can exist")
	fmt.Println("✓ They execute in order of appearance")
	fmt.Println("✓ Cannot be called explicitly")
	fmt.Println("✓ Cannot have parameters or return values")
	fmt.Println("✓ Used for package initialization")

	// === COMMON INIT PATTERNS ===
	fmt.Println("\n5. COMMON INIT PATTERNS:")

	// Pattern 1: Configuration loading
	fmt.Println("Pattern 1: Configuration loading")
	fmt.Printf("  - Config loaded: %t\n", config.AppName != "")

	// Pattern 2: Database connection
	fmt.Println("Pattern 2: Database connection")
	fmt.Printf("  - Database connected: %t\n", database != nil)

	// Pattern 3: Logger setup
	fmt.Println("Pattern 3: Logger setup")
	fmt.Printf("  - Logger ready: %t\n", logger != nil)

	// Pattern 4: Environment variables
	fmt.Println("Pattern 4: Environment variables")
	fmt.Printf("  - Environment: %s\n", config.Environment)

	// Pattern 5: Resource validation
	fmt.Println("Pattern 5: Resource validation")
	fmt.Println("  - All resources validated")

	// === INIT FUNCTION BEST PRACTICES ===
	fmt.Println("\n6. INIT FUNCTION BEST PRACTICES:")

	fmt.Println("✓ Keep init() functions simple and focused")
	fmt.Println("✓ Handle errors appropriately (log.Fatal for critical errors)")
	fmt.Println("✓ Initialize package-level variables")
	fmt.Println("✓ Set up configuration and connections")
	fmt.Println("✓ Validate critical settings")
	fmt.Println("✗ Don't perform heavy computations")
	fmt.Println("✗ Don't start goroutines unless necessary")
	fmt.Println("✗ Don't depend on init() execution order between packages")

	// === INIT VS MAIN COMPARISON ===
	fmt.Println("\n7. INIT VS MAIN COMPARISON:")

	fmt.Println("init() functions:")
	fmt.Println("  - Run automatically before main()")
	fmt.Println("  - Cannot be called explicitly")
	fmt.Println("  - Used for package initialization")
	fmt.Println("  - Can have multiple per package")
	fmt.Println("  - No parameters or return values")

	fmt.Println("main() function:")
	fmt.Println("  - Entry point of the program")
	fmt.Println("  - Runs after all init() functions")
	fmt.Println("  - Contains application logic")
	fmt.Println("  - Only one per package")
	fmt.Println("  - No parameters or return values")

	// === PACKAGE INITIALIZATION ORDER ===
	fmt.Println("\n8. PACKAGE INITIALIZATION ORDER:")

	fmt.Println("1. Package-level variables (in dependency order)")
	fmt.Println("2. init() functions (in order of appearance)")
	fmt.Println("3. main() function (if in main package)")
	fmt.Println("4. Dependencies are initialized first")
	fmt.Println("5. Each package is initialized only once")

	// === REAL-WORLD EXAMPLES ===
	fmt.Println("\n9. REAL-WORLD EXAMPLES:")

	// Example 1: Database connection pool
	fmt.Println("Example 1: Database connection pool")
	fmt.Println("  - Connection pool initialized in init()")
	fmt.Println("  - Ready for use in main()")

	// Example 2: Template parsing
	fmt.Println("Example 2: Template parsing")
	fmt.Println("  - HTML templates parsed in init()")
	fmt.Println("  - Ready for rendering in handlers")

	// Example 3: Flag registration
	fmt.Println("Example 3: Flag registration")
	fmt.Println("  - Command line flags registered in init()")
	fmt.Println("  - Parsed before main() execution")

	// Example 4: HTTP middleware setup
	fmt.Println("Example 4: HTTP middleware setup")
	fmt.Println("  - Middleware chain configured in init()")
	fmt.Println("  - Ready for HTTP server startup")

	// === DEBUGGING INIT FUNCTIONS ===
	fmt.Println("\n10. DEBUGGING INIT FUNCTIONS:")

	fmt.Println("Debugging techniques:")
	fmt.Println("  - Add fmt.Println() statements")
	fmt.Println("  - Use log.Printf() for detailed logging")
	fmt.Println("  - Check initialization order")
	fmt.Println("  - Validate initialization results")

	// === TESTING WITH INIT FUNCTIONS ===
	fmt.Println("\n11. TESTING WITH INIT FUNCTIONS:")

	fmt.Println("Testing considerations:")
	fmt.Println("  - init() functions run before tests")
	fmt.Println("  - May need to reset state for tests")
	fmt.Println("  - Consider using build tags for test init()")
	fmt.Println("  - Test initialization separately if needed")

	// === PERFORMANCE CONSIDERATIONS ===
	fmt.Println("\n12. PERFORMANCE CONSIDERATIONS:")

	fmt.Println("Performance tips:")
	fmt.Println("  - Keep init() functions lightweight")
	fmt.Println("  - Avoid heavy I/O operations")
	fmt.Println("  - Use lazy initialization for expensive resources")
	fmt.Println("  - Consider concurrent initialization if safe")

	// === CLEANUP ===
	fmt.Println("\n13. CLEANUP:")

	// Cleanup resources
	if database != nil {
		database.Close()
	}

	logger.Log("Application shutting down")

	// Calculate total runtime
	totalRuntime := time.Since(startTime)
	fmt.Printf("Total runtime: %v\n", totalRuntime)

	fmt.Println("\n=== END OF INIT FUNCTION GUIDE ===")
}

/*
RUNNING THE PROGRAM:
go run main.go

LEARNING POINTS:

1. INIT FUNCTIONS:
   - Called automatically before main()
   - Used for package initialization
   - Can have multiple per package
   - Execute in order of appearance

2. CHARACTERISTICS:
   - No parameters or return values
   - Cannot be called explicitly
   - Run only once per package
   - Execute after variable initialization

3. COMMON USES:
   - Configuration loading
   - Database connections
   - Logger setup
   - Environment variable processing
   - Resource validation

4. EXECUTION ORDER:
   - Package-level variables first
   - init() functions in order
   - Dependencies initialized first
   - main() function last

5. BEST PRACTICES:
   - Keep simple and focused
   - Handle errors appropriately
   - Initialize package-level variables
   - Validate critical settings
   - Avoid heavy computations

6. DEBUGGING:
   - Add logging statements
   - Check initialization order
   - Validate results
   - Use build tags for test init()

7. PERFORMANCE:
   - Keep lightweight
   - Avoid heavy I/O
   - Use lazy initialization
   - Consider concurrent init if safe

8. REAL-WORLD PATTERNS:
   - Database connection pools
   - Template parsing
   - Flag registration
   - Middleware setup
   - Resource pre-loading

This demonstrates comprehensive init() function usage in Go
for effective package initialization and setup.
*/
