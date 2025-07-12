package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// === CLOSURES IN GO ===

// 1. Basic closure - captures variables from outer scope
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 2. Closure with parameters
func createAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// 3. Closure with multiple captured variables
func createBankAccount(initialBalance float64) (func(float64) float64, func() float64) {
	balance := initialBalance

	deposit := func(amount float64) float64 {
		balance += amount
		return balance
	}

	getBalance := func() float64 {
		return balance
	}

	return deposit, getBalance
}

// 4. Closure capturing by reference
func createMultiplier() func(int) func(int) int {
	multiplier := 1

	return func(factor int) func(int) int {
		multiplier *= factor
		return func(value int) int {
			return value * multiplier
		}
	}
}

// 5. Closure with slice modification
func createList() (func(string), func() []string, func()) {
	var items []string

	add := func(item string) {
		items = append(items, item)
	}

	getAll := func() []string {
		return items
	}

	clear := func() {
		items = items[:0]
	}

	return add, getAll, clear
}

// 6. Closure for configuration
func createLogger(prefix string) func(string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return func(message string) {
		fmt.Printf("[%s] %s: %s\n", timestamp, prefix, message)
	}
}

// 7. Closure with error handling
func createValidator(rules ...func(string) error) func(string) []error {
	return func(input string) []error {
		var errors []error
		for _, rule := range rules {
			if err := rule(input); err != nil {
				errors = append(errors, err)
			}
		}
		return errors
	}
}

// 8. Closure for caching/memoization
func createMemoizer() func(int) int {
	cache := make(map[int]int)

	return func(n int) int {
		if result, exists := cache[n]; exists {
			fmt.Printf("Cache hit for %d\n", n)
			return result
		}

		// Expensive calculation (factorial)
		result := 1
		for i := 1; i <= n; i++ {
			result *= i
		}

		cache[n] = result
		fmt.Printf("Calculated and cached %d! = %d\n", n, result)
		return result
	}
}

// 9. Closure for event handling
func createEventHandler() (func(string, interface{}), func()) {
	var events []string

	handleEvent := func(eventType string, data interface{}) {
		event := fmt.Sprintf("%s: %v", eventType, data)
		events = append(events, event)
		fmt.Printf("Event handled: %s\n", event)
	}

	showHistory := func() {
		fmt.Println("Event history:")
		for i, event := range events {
			fmt.Printf("%d: %s\n", i+1, event)
		}
	}

	return handleEvent, showHistory
}

// 10. Closure for functional programming
func createFilter(predicate func(int) bool) func([]int) []int {
	return func(slice []int) []int {
		var result []int
		for _, value := range slice {
			if predicate(value) {
				result = append(result, value)
			}
		}
		return result
	}
}

func createMapper(transform func(int) int) func([]int) []int {
	return func(slice []int) []int {
		result := make([]int, len(slice))
		for i, value := range slice {
			result[i] = transform(value)
		}
		return result
	}
}

// 11. Closure for partial application
func createPartialApplication(fn func(int, int, int) int, a, b int) func(int) int {
	return func(c int) int {
		return fn(a, b, c)
	}
}

// 12. Closure with cleanup
func createResource(name string) (func() string, func()) {
	fmt.Printf("Creating resource: %s\n", name)

	use := func() string {
		return fmt.Sprintf("Using resource: %s", name)
	}

	cleanup := func() {
		fmt.Printf("Cleaning up resource: %s\n", name)
	}

	return use, cleanup
}

// 13. Closure for state machine
func createStateMachine() (func(string), func() string) {
	states := []string{"idle", "running", "paused", "stopped"}
	currentState := 0

	transition := func(action string) {
		switch action {
		case "start":
			if currentState == 0 {
				currentState = 1
			}
		case "pause":
			if currentState == 1 {
				currentState = 2
			}
		case "resume":
			if currentState == 2 {
				currentState = 1
			}
		case "stop":
			if currentState == 1 || currentState == 2 {
				currentState = 3
			}
		case "reset":
			currentState = 0
		}
	}

	getState := func() string {
		return states[currentState]
	}

	return transition, getState
}

// 14. Closure for rate limiting
func createRateLimiter(maxCalls int, duration time.Duration) func() bool {
	var calls []time.Time

	return func() bool {
		now := time.Now()
		// Remove old calls
		for len(calls) > 0 && now.Sub(calls[0]) > duration {
			calls = calls[1:]
		}

		if len(calls) >= maxCalls {
			return false
		}

		calls = append(calls, now)
		return true
	}
}

// 15. Closure for middleware pattern
func createMiddleware(handler func(string) string) func(string) string {
	return func(input string) string {
		fmt.Printf("Middleware: Before processing '%s'\n", input)
		result := handler(input)
		fmt.Printf("Middleware: After processing, result: '%s'\n", result)
		return result
	}
}

// 16. Advanced closure with nested functions
func createCalculator() map[string]func(float64, float64) float64 {
	// Private helper function
	safeDiv := func(a, b float64) float64 {
		if b == 0 {
			fmt.Println("Warning: Division by zero")
			return math.Inf(1)
		}
		return a / b
	}

	return map[string]func(float64, float64) float64{
		"add": func(a, b float64) float64 { return a + b },
		"sub": func(a, b float64) float64 { return a - b },
		"mul": func(a, b float64) float64 { return a * b },
		"div": safeDiv,
		"pow": func(a, b float64) float64 { return math.Pow(a, b) },
		"mod": func(a, b float64) float64 { return math.Mod(a, b) },
	}
}

// 17. Closure for dependency injection
func createService(config map[string]interface{}) func(string) interface{} {
	return func(key string) interface{} {
		if value, exists := config[key]; exists {
			return value
		}
		return nil
	}
}

// 18. Closure loop variable capture (common pitfall)
func demonstrateLoopCapture() {
	fmt.Println("\n--- LOOP VARIABLE CAPTURE ---")

	// WRONG way - captures loop variable by reference
	var wrongFunctions []func() int
	for i := 0; i < 3; i++ {
		wrongFunctions = append(wrongFunctions, func() int {
			return i // All functions will return 3
		})
	}

	fmt.Println("Wrong way (captures by reference):")
	for j, fn := range wrongFunctions {
		fmt.Printf("Function %d returns: %d\n", j, fn())
	}

	// CORRECT way - captures loop variable by value
	var correctFunctions []func() int
	for i := 0; i < 3; i++ {
		i := i // Create new variable in each iteration
		correctFunctions = append(correctFunctions, func() int {
			return i
		})
	}

	fmt.Println("Correct way (captures by value):")
	for j, fn := range correctFunctions {
		fmt.Printf("Function %d returns: %d\n", j, fn())
	}
}

// 19. Closure for builder pattern
func createQueryBuilder() func(string) func() string {
	var parts []string

	return func(part string) func() string {
		parts = append(parts, part)
		return func() string {
			return strings.Join(parts, " ")
		}
	}
}

// 20. Closure for observer pattern
func createObserver() (func(func(string)), func(string)) {
	var observers []func(string)

	subscribe := func(observer func(string)) {
		observers = append(observers, observer)
	}

	notify := func(message string) {
		for _, observer := range observers {
			observer(message)
		}
	}

	return subscribe, notify
}

func main() {
	fmt.Println("=== GO CLOSURES COMPREHENSIVE GUIDE ===")

	// === BASIC COUNTER ===
	fmt.Println("\n--- BASIC COUNTER ---")
	counter := createCounter()
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())

	// Multiple counters are independent
	counter2 := createCounter()
	fmt.Printf("Counter2: %d\n", counter2())
	fmt.Printf("Counter: %d\n", counter())

	/*
		JavaScript comparison:
		function createCounter() {
			let count = 0;
			return function() {
				return ++count;
			};
		}

		const counter = createCounter();
		console.log(counter()); // 1
		console.log(counter()); // 2
	*/

	// === ADDER CLOSURE ===
	fmt.Println("\n--- ADDER CLOSURE ---")
	add5 := createAdder(5)
	add10 := createAdder(10)

	fmt.Printf("add5(3) = %d\n", add5(3))
	fmt.Printf("add10(3) = %d\n", add10(3))

	// === BANK ACCOUNT ===
	fmt.Println("\n--- BANK ACCOUNT ---")
	deposit, getBalance := createBankAccount(100.0)

	fmt.Printf("Initial balance: $%.2f\n", getBalance())
	fmt.Printf("After deposit $50: $%.2f\n", deposit(50.0))
	fmt.Printf("After deposit $25: $%.2f\n", deposit(25.0))
	fmt.Printf("Current balance: $%.2f\n", getBalance())

	// === MULTIPLIER CLOSURE ===
	fmt.Println("\n--- MULTIPLIER CLOSURE ---")
	multiplier := createMultiplier()
	double := multiplier(2)
	triple := multiplier(3) // Now multiplier is 6

	fmt.Printf("double(5) = %d\n", double(5))
	fmt.Printf("triple(5) = %d\n", triple(5))

	// === LIST CLOSURE ===
	fmt.Println("\n--- LIST CLOSURE ---")
	add, getAll, clear := createList()

	add("apple")
	add("banana")
	add("cherry")
	fmt.Printf("Items: %v\n", getAll())

	clear()
	fmt.Printf("After clear: %v\n", getAll())

	// === LOGGER CLOSURE ===
	fmt.Println("\n--- LOGGER CLOSURE ---")
	errorLogger := createLogger("ERROR")
	infoLogger := createLogger("INFO")

	errorLogger("Something went wrong")
	infoLogger("Application started")

	// === VALIDATOR CLOSURE ===
	fmt.Println("\n--- VALIDATOR CLOSURE ---")
	validator := createValidator(
		func(s string) error {
			if len(s) < 3 {
				return fmt.Errorf("string too short")
			}
			return nil
		},
		func(s string) error {
			if !strings.Contains(s, "@") {
				return fmt.Errorf("must contain @")
			}
			return nil
		},
	)

	errors := validator("ab")
	fmt.Printf("Validation errors for 'ab': %v\n", errors)

	errors = validator("test@example.com")
	fmt.Printf("Validation errors for 'test@example.com': %v\n", errors)

	// === MEMOIZATION ===
	fmt.Println("\n--- MEMOIZATION ---")
	factorial := createMemoizer()

	fmt.Printf("factorial(5) = %d\n", factorial(5))
	fmt.Printf("factorial(5) = %d\n", factorial(5)) // Should hit cache
	fmt.Printf("factorial(6) = %d\n", factorial(6))

	// === EVENT HANDLER ===
	fmt.Println("\n--- EVENT HANDLER ---")
	handleEvent, showHistory := createEventHandler()

	handleEvent("click", "button1")
	handleEvent("keypress", "Enter")
	handleEvent("scroll", 100)
	showHistory()

	// === FUNCTIONAL PROGRAMMING ===
	fmt.Println("\n--- FUNCTIONAL PROGRAMMING ---")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evenFilter := createFilter(func(n int) bool { return n%2 == 0 })
	squareMapper := createMapper(func(n int) int { return n * n })

	evens := evenFilter(numbers)
	squares := squareMapper(evens)

	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Evens: %v\n", evens)
	fmt.Printf("Squares of evens: %v\n", squares)

	// === PARTIAL APPLICATION ===
	fmt.Println("\n--- PARTIAL APPLICATION ---")
	multiply := func(a, b, c int) int { return a * b * c }
	multiplyBy2And3 := createPartialApplication(multiply, 2, 3)

	fmt.Printf("multiplyBy2And3(4) = %d\n", multiplyBy2And3(4))

	// === RESOURCE MANAGEMENT ===
	fmt.Println("\n--- RESOURCE MANAGEMENT ---")
	useResource, cleanup := createResource("Database Connection")
	fmt.Println(useResource())
	defer cleanup()

	// === STATE MACHINE ===
	fmt.Println("\n--- STATE MACHINE ---")
	transition, getState := createStateMachine()

	fmt.Printf("Initial state: %s\n", getState())
	transition("start")
	fmt.Printf("After start: %s\n", getState())
	transition("pause")
	fmt.Printf("After pause: %s\n", getState())
	transition("resume")
	fmt.Printf("After resume: %s\n", getState())
	transition("stop")
	fmt.Printf("After stop: %s\n", getState())

	// === RATE LIMITING ===
	fmt.Println("\n--- RATE LIMITING ---")
	rateLimiter := createRateLimiter(3, 2*time.Second)

	for i := 0; i < 5; i++ {
		if rateLimiter() {
			fmt.Printf("Request %d: Allowed\n", i+1)
		} else {
			fmt.Printf("Request %d: Rate limited\n", i+1)
		}
	}

	// === MIDDLEWARE PATTERN ===
	fmt.Println("\n--- MIDDLEWARE PATTERN ---")
	handler := func(input string) string {
		return strings.ToUpper(input)
	}

	middleware := createMiddleware(handler)
	result := middleware("hello world")
	fmt.Printf("Final result: %s\n", result)

	// === ADVANCED CALCULATOR ===
	fmt.Println("\n--- ADVANCED CALCULATOR ---")
	calc := createCalculator()

	fmt.Printf("10 + 5 = %.2f\n", calc["add"](10, 5))
	fmt.Printf("10 - 5 = %.2f\n", calc["sub"](10, 5))
	fmt.Printf("10 * 5 = %.2f\n", calc["mul"](10, 5))
	fmt.Printf("10 / 5 = %.2f\n", calc["div"](10, 5))
	fmt.Printf("10 / 0 = %.2f\n", calc["div"](10, 0))
	fmt.Printf("2 ^ 3 = %.2f\n", calc["pow"](2, 3))

	// === DEPENDENCY INJECTION ===
	fmt.Println("\n--- DEPENDENCY INJECTION ---")
	service := createService(map[string]interface{}{
		"database": "postgresql://localhost:5432/mydb",
		"cache":    "redis://localhost:6379",
		"timeout":  30,
	})

	fmt.Printf("Database: %v\n", service("database"))
	fmt.Printf("Cache: %v\n", service("cache"))
	fmt.Printf("Timeout: %v\n", service("timeout"))

	// === LOOP VARIABLE CAPTURE ===
	demonstrateLoopCapture()

	// === QUERY BUILDER ===
	fmt.Println("\n--- QUERY BUILDER ---")
	builder := createQueryBuilder()

	builder("SELECT *")
	builder("FROM users")
	builder("WHERE age > 18")
	finalQuery := builder("ORDER BY name")

	fmt.Printf("Final query: %s\n", finalQuery())

	// === OBSERVER PATTERN ===
	fmt.Println("\n--- OBSERVER PATTERN ---")
	subscribe, notify := createObserver()

	subscribe(func(msg string) {
		fmt.Printf("Observer 1 received: %s\n", msg)
	})

	subscribe(func(msg string) {
		fmt.Printf("Observer 2 received: %s\n", msg)
	})

	notify("Hello World!")
	notify("System shutting down")

	// === CLOSURE BEST PRACTICES ===
	fmt.Println("\n--- CLOSURE BEST PRACTICES ---")
	fmt.Println("1. Use closures for encapsulation and data privacy")
	fmt.Println("2. Be aware of variable capture by reference vs value")
	fmt.Println("3. Use closures for configuration and dependency injection")
	fmt.Println("4. Leverage closures for functional programming patterns")
	fmt.Println("5. Use closures for event handling and callbacks")
	fmt.Println("6. Be careful with loop variable capture")
	fmt.Println("7. Use closures for caching and memoization")
	fmt.Println("8. Closures are great for factory patterns")
	fmt.Println("9. Use closures for middleware and decorators")
	fmt.Println("10. Clean up resources in closures when needed")
}
