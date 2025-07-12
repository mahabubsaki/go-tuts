package main

import "fmt"

// === FUNCTION DECLARATION ===

// 1. Basic function with no parameters and no return value
func greet() {
	fmt.Println("Hello, World!")
}

// 2. Function with parameters
func greetPerson(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 3. Function with parameters and return value
func add(a, b int) int {
	return a + b
}

// 4. Function with multiple parameters of different types
func introduce(name string, age int, isStudent bool) {
	fmt.Printf("My name is %s, I am %d years old", name, age)
	if isStudent {
		fmt.Println(" and I am a student.")
	} else {
		fmt.Println(" and I am not a student.")
	}
}

// 5. Function with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// 6. Function with named return values
func calculateStats(numbers []int) (sum int, count int, average float64) {
	count = len(numbers)
	if count == 0 {
		return // Returns zero values: sum=0, count=0, average=0.0
	}

	for _, num := range numbers {
		sum += num
	}
	average = float64(sum) / float64(count)
	return // Returns the named values
}

// 7. Function with same type parameters (shorthand)
func multiply(x, y, z int) int {
	return x * y * z
}

// 8. Function that returns a function
func makeMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// 9. Function with slice parameter
func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 10. Function with map parameter
func printMap(m map[string]int) {
	for key, value := range m {
		fmt.Printf("%s: %d\n", key, value)
	}
}

// 11. Function with pointer parameter
func increment(n *int) {
	*n++
}

// 12. Function with struct parameter
type Person struct {
	Name string
	Age  int
}

func describePerson(p Person) {
	fmt.Printf("Person: %s, Age: %d\n", p.Name, p.Age)
}

// 13. Function with struct pointer parameter
func updateAge(p *Person, newAge int) {
	p.Age = newAge
}

// 14. Function with interface parameter
type Greeter interface {
	Greet() string
}

type English struct{}

func (e English) Greet() string { return "Hello!" }

type Spanish struct{}

func (s Spanish) Greet() string { return "¡Hola!" }

func sayHello(g Greeter) {
	fmt.Println(g.Greet())
}

// 15. Exported vs unexported functions
// Exported function (starts with capital letter)
func PublicFunction() {
	fmt.Println("This is an exported function")
}

// Unexported function (starts with lowercase letter)
func privateFunction() {
	fmt.Println("This is an unexported function")
}

func main() {
	fmt.Println("=== GO FUNCTIONS - BASIC DECLARATION ===")

	// === BASIC FUNCTION CALLS ===
	fmt.Println("\n--- BASIC FUNCTION CALLS ---")
	greet()
	greetPerson("Alice")
	greetPerson("Bob")

	/*
		JavaScript comparison:
		function greet() {
			console.log("Hello, World!");
		}

		function greetPerson(name) {
			console.log(`Hello, ${name}!`);
		}
	*/

	// === FUNCTIONS WITH RETURN VALUES ===
	fmt.Println("\n--- FUNCTIONS WITH RETURN VALUES ---")
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	product := multiply(2, 3, 4)
	fmt.Printf("2 × 3 × 4 = %d\n", product)

	/*
		JavaScript comparison:
		function add(a, b) {
			return a + b;
		}

		const result = add(5, 3);
		console.log(`5 + 3 = ${result}`);
	*/

	// === FUNCTIONS WITH MULTIPLE PARAMETERS ===
	fmt.Println("\n--- FUNCTIONS WITH MULTIPLE PARAMETERS ---")
	introduce("Charlie", 25, true)
	introduce("Diana", 30, false)

	// === FUNCTIONS WITH MULTIPLE RETURN VALUES ===
	fmt.Println("\n--- FUNCTIONS WITH MULTIPLE RETURN VALUES ---")
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", quotient)
	}

	quotient, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 0 = %.2f\n", quotient)
	}

	/*
		JavaScript comparison (using objects or arrays):
		function divide(a, b) {
			if (b === 0) {
				return { result: 0, error: "Cannot divide by zero" };
			}
			return { result: a / b, error: null };
		}

		const { result, error } = divide(10, 2);
	*/

	// === NAMED RETURN VALUES ===
	fmt.Println("\n--- NAMED RETURN VALUES ---")
	numbers := []int{1, 2, 3, 4, 5}
	s, c, avg := calculateStats(numbers)
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Sum: %d, Count: %d, Average: %.2f\n", s, c, avg)

	// Empty slice
	emptyNumbers := []int{}
	s, c, avg = calculateStats(emptyNumbers)
	fmt.Printf("Empty slice - Sum: %d, Count: %d, Average: %.2f\n", s, c, avg)

	// === FUNCTIONS RETURNING FUNCTIONS ===
	fmt.Println("\n--- FUNCTIONS RETURNING FUNCTIONS ---")
	double := makeMultiplier(2)
	triple := makeMultiplier(3)

	fmt.Printf("Double of 5: %d\n", double(5))
	fmt.Printf("Triple of 5: %d\n", triple(5))

	/*
		JavaScript comparison:
		function makeMultiplier(factor) {
			return function(n) {
				return n * factor;
			};
		}

		const double = makeMultiplier(2);
		const triple = makeMultiplier(3);
	*/

	// === FUNCTIONS WITH DIFFERENT PARAMETER TYPES ===
	fmt.Println("\n--- FUNCTIONS WITH DIFFERENT PARAMETER TYPES ---")

	// Slice parameter
	nums := []int{1, 2, 3, 4, 5}
	total := sum(nums)
	fmt.Printf("Sum of %v = %d\n", nums, total)

	// Map parameter
	scores := map[string]int{
		"Alice": 95,
		"Bob":   87,
		"Carol": 92,
	}
	fmt.Println("Scores:")
	printMap(scores)

	// Pointer parameter
	value := 10
	fmt.Printf("Before increment: %d\n", value)
	increment(&value)
	fmt.Printf("After increment: %d\n", value)

	// Struct parameter
	person := Person{Name: "Eve", Age: 28}
	describePerson(person)

	// Struct pointer parameter
	fmt.Printf("Before update: %s is %d years old\n", person.Name, person.Age)
	updateAge(&person, 29)
	fmt.Printf("After update: %s is %d years old\n", person.Name, person.Age)

	// Interface parameter
	fmt.Println("Greeting in different languages:")
	sayHello(English{})
	sayHello(Spanish{})

	// === EXPORTED VS UNEXPORTED FUNCTIONS ===
	fmt.Println("\n--- EXPORTED VS UNEXPORTED FUNCTIONS ---")
	PublicFunction()
	privateFunction()

	// === FUNCTION SIGNATURES ===
	fmt.Println("\n--- FUNCTION SIGNATURES ---")
	fmt.Println("Function signatures in Go:")
	fmt.Println("func functionName(param1 type1, param2 type2) returnType")
	fmt.Println("func functionName(param1, param2 type) returnType")
	fmt.Println("func functionName(param type) (return1 type1, return2 type2)")
	fmt.Println("func functionName(param type) (namedReturn1 type1, namedReturn2 type2)")

	// === FUNCTION CALLING PATTERNS ===
	fmt.Println("\n--- FUNCTION CALLING PATTERNS ---")

	// Direct function call
	fmt.Println("Direct call:", add(3, 4))

	// Function result in expression
	fmt.Println("In expression:", add(1, 2)+add(3, 4))

	// Function result as parameter
	fmt.Printf("As parameter: %d\n", multiply(add(1, 2), add(3, 4), 2))

	// Chaining function calls
	fmt.Printf("Chained calls: %d\n", makeMultiplier(2)(5))

	// Multiple assignment from function
	a, b := 5, 3
	sum, product := add(a, b), multiply(a, b, 1)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)

	// === ANONYMOUS FUNCTIONS ===
	fmt.Println("\n--- ANONYMOUS FUNCTIONS ---")

	// Anonymous function assigned to variable
	square := func(n int) int {
		return n * n
	}
	fmt.Printf("Square of 4: %d\n", square(4))

	// Immediately invoked function expression (IIFE)
	result = func(x, y int) int {
		return x*x + y*y
	}(3, 4)
	fmt.Printf("3² + 4² = %d\n", result)

	/*
		JavaScript comparison:
		// Anonymous function
		const square = function(n) {
			return n * n;
		};

		// IIFE
		const result = (function(x, y) {
			return x*x + y*y;
		})(3, 4);
	*/

	// === FUNCTION BEST PRACTICES ===
	fmt.Println("\n--- FUNCTION BEST PRACTICES ---")
	fmt.Println("1. Use descriptive function names")
	fmt.Println("2. Keep functions small and focused")
	fmt.Println("3. Use multiple return values for error handling")
	fmt.Println("4. Use named return values for complex functions")
	fmt.Println("5. Prefer composition over deep nesting")
	fmt.Println("6. Export functions only when necessary")
	fmt.Println("7. Use interfaces for flexible function parameters")
	fmt.Println("8. Document exported functions")
	fmt.Println("9. Handle errors appropriately")
	fmt.Println("10. Use consistent parameter ordering")
}
