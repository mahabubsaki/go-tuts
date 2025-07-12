package main

import (
	"fmt"
	"os"
	"time"
)

// === GO DEFER STATEMENT COMPREHENSIVE GUIDE ===

/*
DEFER PHILOSOPHY:
- defer schedules function calls to run after the surrounding function returns
- Deferred calls are executed in LIFO (Last In, First Out) order
- Perfect for cleanup operations and resource management
- Arguments are evaluated when defer is executed, not when deferred function runs

COMPARISON WITH JAVASCRIPT:
// JavaScript: Manual cleanup with try/finally
try {
    const file = fs.openSync('file.txt', 'r');
    // work with file
} finally {
    fs.closeSync(file);
}

// Go: Automatic cleanup with defer
file, err := os.Open("file.txt")
if err != nil {
    return err
}
defer file.Close()
// work with file
*/

// 1. BASIC DEFER USAGE
func demonstrateBasicDefer() {
	fmt.Println("Function start")

	// Deferred calls execute in reverse order
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	defer fmt.Println("Deferred 3")

	fmt.Println("Function middle")
	fmt.Println("Function end")
	// Output order: Function start, Function middle, Function end, Deferred 3, Deferred 2, Deferred 1
}

// 2. DEFER WITH ARGUMENTS
func demonstrateDeferArguments() {
	x := 10

	// Arguments are evaluated when defer is executed
	defer fmt.Printf("x was %d when defer was called\n", x)

	x = 20
	fmt.Printf("x is now %d\n", x)

	// The deferred function will print 10, not 20
}

// 3. DEFER WITH ANONYMOUS FUNCTIONS
func demonstrateDeferWithClosures() {
	x := 10

	// Using closure to capture variables by reference
	defer func() {
		fmt.Printf("x is %d in closure\n", x)
	}()

	x = 20
	fmt.Printf("x is now %d\n", x)

	// The deferred closure will print 20
}

// 4. RESOURCE CLEANUP
func demonstrateResourceCleanup() error {
	// File handling with defer
	file, err := os.Create("temp.txt")
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close() // Ensures file is closed even if function returns early

	// Write to file
	_, err = file.WriteString("Hello, World!")
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	fmt.Println("File operations completed")
	return nil
}

// 5. DEFER WITH LOCKS
type Counter struct {
	value int
}

func (c *Counter) Increment() {
	// In real code, you would use sync.Mutex
	fmt.Println("Acquiring lock")
	defer fmt.Println("Releasing lock")

	c.value++
	fmt.Printf("Counter value: %d\n", c.value)
}

// 6. DEFER FOR TIMING
func demonstrateTimingWithDefer() {
	start := time.Now()
	defer func() {
		fmt.Printf("Function took %v to execute\n", time.Since(start))
	}()

	// Simulate some work
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Work completed")
}

// 7. DEFER WITH RECOVERY
func demonstrateDeferWithRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	fmt.Println("Before panic")
	panic("Something went wrong!")
	fmt.Println("This won't be executed")
}

// 8. DEFER IN LOOPS
func demonstrateDeferInLoops() {
	fmt.Println("--- DEFER IN LOOPS ---")

	// ❌ Problem: All defers execute at function end
	for i := 0; i < 3; i++ {
		defer fmt.Printf("Deferred from loop iteration %d\n", i)
	}

	fmt.Println("Loop completed")
	// Output: Loop completed, Deferred from loop iteration 2, 1, 0
}

// 9. DEFER WITH FUNCTION SCOPE
func demonstrateDeferScope() {
	fmt.Println("--- DEFER SCOPE ---")

	// ✅ Solution: Use function literals for proper scoping
	for i := 0; i < 3; i++ {
		func(iteration int) {
			defer fmt.Printf("Properly scoped defer from iteration %d\n", iteration)
			fmt.Printf("Processing iteration %d\n", iteration)
		}(i)
	}
}

// 10. DEFER WITH RETURN VALUES
func demonstrateDeferWithReturnValues() (result int) {
	defer func() {
		fmt.Printf("Function returning %d\n", result)
		result = result * 2 // Can modify named return values
	}()

	result = 10
	return result
}

// 11. DEFER WITH MULTIPLE RETURN VALUES
func demonstrateDeferWithMultipleReturns() (string, int, error) {
	defer func() {
		fmt.Println("Cleanup operations")
	}()

	// Multiple return paths - defer ensures cleanup
	if time.Now().Second()%2 == 0 {
		return "even", 2, nil
	}

	return "odd", 1, nil
}

// 12. DEFER WITH ERROR HANDLING
func demonstrateDeferWithErrorHandling() error {
	fmt.Println("--- DEFER WITH ERROR HANDLING ---")

	// Setup
	resource := "database connection"
	fmt.Printf("Opening %s\n", resource)

	defer func() {
		fmt.Printf("Closing %s\n", resource)
	}()

	// Simulate error condition
	if time.Now().Second()%2 == 0 {
		return fmt.Errorf("connection failed")
	}

	fmt.Println("Operations completed successfully")
	return nil
}

// 13. DEFER WITH STACK UNWINDING
func demonstrateStackUnwinding() {
	fmt.Println("--- STACK UNWINDING ---")

	defer fmt.Println("Level 1 cleanup")
	level2()
}

func level2() {
	defer fmt.Println("Level 2 cleanup")
	level3()
}

func level3() {
	defer fmt.Println("Level 3 cleanup")
	fmt.Println("Deep in the call stack")
}

// 14. DEFER WITH CHANNELS
func demonstrateDeferWithChannels() {
	fmt.Println("--- DEFER WITH CHANNELS ---")

	ch := make(chan int)

	go func() {
		defer close(ch) // Ensures channel is closed

		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()

	// Read from channel
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
}

// 15. DEFER BEST PRACTICES
func demonstrateDeferBestPractices() {
	fmt.Println("--- DEFER BEST PRACTICES ---")

	// 1. Use defer immediately after resource acquisition
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close() // Defer immediately after successful creation

	// 2. Use defer for cleanup operations
	fmt.Println("Starting operations")
	defer fmt.Println("Operations completed")

	// 3. Use defer with named return values for post-processing
	processData := func() (result string) {
		defer func() {
			result = fmt.Sprintf("Processed: %s", result)
		}()

		return "raw data"
	}

	fmt.Println(processData())

	// 4. Use defer for timing and profiling
	start := time.Now()
	defer func() {
		fmt.Printf("Operation took: %v\n", time.Since(start))
	}()

	// Simulate work
	time.Sleep(50 * time.Millisecond)
}

// 16. DEFER GOTCHAS AND PITFALLS
func demonstrateDeferGotchas() {
	fmt.Println("--- DEFER GOTCHAS ---")

	// Gotcha 1: Arguments evaluated at defer time
	x := 1
	defer fmt.Printf("x = %d (evaluated at defer time)\n", x)
	x = 2

	// Gotcha 2: Defer in loops
	fmt.Println("Defer in loops:")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("Loop defer %d\n", i)
	}

	// Gotcha 3: Defer with pointers
	data := &struct{ value int }{value: 1}
	defer fmt.Printf("Data value: %d (pointer dereference)\n", data.value)
	data.value = 2

	fmt.Println("Main function continues...")
}

func main() {
	fmt.Println("=== GO DEFER STATEMENT COMPREHENSIVE GUIDE ===")

	// === BASIC DEFER ===
	fmt.Println("\n--- BASIC DEFER ---")
	demonstrateBasicDefer()

	// === DEFER ARGUMENTS ===
	fmt.Println("\n--- DEFER ARGUMENTS ---")
	demonstrateDeferArguments()

	// === DEFER WITH CLOSURES ===
	fmt.Println("\n--- DEFER WITH CLOSURES ---")
	demonstrateDeferWithClosures()

	// === RESOURCE CLEANUP ===
	fmt.Println("\n--- RESOURCE CLEANUP ---")
	if err := demonstrateResourceCleanup(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// === DEFER WITH LOCKS ===
	fmt.Println("\n--- DEFER WITH LOCKS ---")
	counter := &Counter{}
	counter.Increment()

	// === TIMING WITH DEFER ===
	fmt.Println("\n--- TIMING WITH DEFER ---")
	demonstrateTimingWithDefer()

	// === DEFER WITH RECOVERY ===
	fmt.Println("\n--- DEFER WITH RECOVERY ---")
	demonstrateDeferWithRecover()

	// === DEFER IN LOOPS ===
	fmt.Println("\n--- DEFER IN LOOPS ---")
	demonstrateDeferInLoops()

	// === DEFER SCOPE ===
	fmt.Println("\n--- DEFER SCOPE ---")
	demonstrateDeferScope()

	// === DEFER WITH RETURN VALUES ===
	fmt.Println("\n--- DEFER WITH RETURN VALUES ---")
	result := demonstrateDeferWithReturnValues()
	fmt.Printf("Final result: %d\n", result)

	// === DEFER WITH MULTIPLE RETURNS ===
	fmt.Println("\n--- DEFER WITH MULTIPLE RETURNS ---")
	str, num, err := demonstrateDeferWithMultipleReturns()
	fmt.Printf("Results: %s, %d, %v\n", str, num, err)

	// === DEFER WITH ERROR HANDLING ===
	fmt.Println("\n--- DEFER WITH ERROR HANDLING ---")
	if err := demonstrateDeferWithErrorHandling(); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
	}

	// === STACK UNWINDING ===
	fmt.Println("\n--- STACK UNWINDING ---")
	demonstrateStackUnwinding()

	// === DEFER WITH CHANNELS ===
	fmt.Println("\n--- DEFER WITH CHANNELS ---")
	demonstrateDeferWithChannels()

	// === DEFER BEST PRACTICES ===
	fmt.Println("\n--- DEFER BEST PRACTICES ---")
	demonstrateDeferBestPractices()

	// === DEFER GOTCHAS ===
	fmt.Println("\n--- DEFER GOTCHAS ---")
	demonstrateDeferGotchas()

	// === DEFER RULES SUMMARY ===
	fmt.Println("\n--- DEFER RULES SUMMARY ---")
	fmt.Println("1. Deferred functions execute in LIFO order")
	fmt.Println("2. Arguments are evaluated when defer is executed")
	fmt.Println("3. Deferred functions can access and modify named return values")
	fmt.Println("4. Use defer for cleanup operations")
	fmt.Println("5. Place defer statements immediately after resource acquisition")
	fmt.Println("6. Be careful with defer in loops")
	fmt.Println("7. Use closures to capture variables by reference")
	fmt.Println("8. Defer is essential for panic recovery")
	fmt.Println("9. Defer ensures cleanup even on early returns")
	fmt.Println("10. Use defer for timing and profiling")

	// === COMMON PATTERNS ===
	fmt.Println("\n--- COMMON PATTERNS ---")
	fmt.Println("✅ resource, err := acquire()")
	fmt.Println("   if err != nil { return err }")
	fmt.Println("   defer resource.Close()")
	fmt.Println()
	fmt.Println("✅ defer func() {")
	fmt.Println("     if r := recover(); r != nil {")
	fmt.Println("         // handle panic")
	fmt.Println("     }")
	fmt.Println("   }()")
	fmt.Println()
	fmt.Println("✅ start := time.Now()")
	fmt.Println("   defer func() {")
	fmt.Println("     fmt.Printf(\"Duration: %v\\n\", time.Since(start))")
	fmt.Println("   }()")

	// Clean up temp file
	os.Remove("temp.txt")
	os.Remove("example.txt")

	fmt.Println("\nDefer makes Go programs more robust and maintainable!")
}
