package main

import (
	"fmt"
	"math"
	"strings"
)

// === METHODS IN GO ===

// 1. Basic struct for method examples
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
}

// 2. Value receiver methods
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d, Email: %s}",
		p.FullName(), p.Age, p.Email)
}

// 3. Pointer receiver methods
func (p *Person) SetAge(age int) {
	p.Age = age
}

func (p *Person) SetEmail(email string) {
	p.Email = email
}

func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("%s is now %d years old!\n", p.FullName(), p.Age)
}

func (p *Person) UpdateName(firstName, lastName string) {
	p.FirstName = firstName
	p.LastName = lastName
}

// 4. Rectangle struct with geometric methods
type Rectangle struct {
	Width  float64
	Height float64
}

// Value receiver methods for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) IsSquare() bool {
	return r.Width == r.Height
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle{Width: %.2f, Height: %.2f}", r.Width, r.Height)
}

// Pointer receiver methods for Rectangle
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func (r *Rectangle) Resize(width, height float64) {
	r.Width = width
	r.Height = height
}

// 5. Circle struct
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle{Radius: %.2f}", c.Radius)
}

func (c *Circle) SetRadius(radius float64) {
	c.Radius = radius
}

// 6. Bank Account with method chaining
type BankAccount struct {
	AccountNumber string
	Balance       float64
	Owner         string
}

func (b *BankAccount) Deposit(amount float64) *BankAccount {
	if amount > 0 {
		b.Balance += amount
		fmt.Printf("Deposited $%.2f. New balance: $%.2f\n", amount, b.Balance)
	}
	return b
}

func (b *BankAccount) Withdraw(amount float64) *BankAccount {
	if amount > 0 && amount <= b.Balance {
		b.Balance -= amount
		fmt.Printf("Withdrew $%.2f. New balance: $%.2f\n", amount, b.Balance)
	} else {
		fmt.Printf("Insufficient funds or invalid amount\n")
	}
	return b
}

func (b *BankAccount) GetBalance() float64 {
	return b.Balance
}

func (b *BankAccount) String() string {
	return fmt.Sprintf("BankAccount{Number: %s, Owner: %s, Balance: $%.2f}",
		b.AccountNumber, b.Owner, b.Balance)
}

// 7. Methods on custom types
type Temperature float64

func (t Temperature) Celsius() float64 {
	return float64(t)
}

func (t Temperature) Fahrenheit() float64 {
	return float64(t)*9/5 + 32
}

func (t Temperature) Kelvin() float64 {
	return float64(t) + 273.15
}

func (t Temperature) String() string {
	return fmt.Sprintf("%.2f°C", float64(t))
}

// 8. Methods on slice types
type IntSlice []int

func (s IntSlice) Sum() int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}

func (s IntSlice) Average() float64 {
	if len(s) == 0 {
		return 0
	}
	return float64(s.Sum()) / float64(len(s))
}

func (s IntSlice) Max() int {
	if len(s) == 0 {
		return 0
	}
	max := s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}

func (s IntSlice) Min() int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func (s IntSlice) Contains(value int) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

func (s IntSlice) String() string {
	return fmt.Sprintf("IntSlice%v", []int(s))
}

// 9. Methods on map types
type StringCounter map[string]int

func (sc StringCounter) Add(key string) {
	sc[key]++
}

func (sc StringCounter) Get(key string) int {
	return sc[key]
}

func (sc StringCounter) Total() int {
	total := 0
	for _, count := range sc {
		total += count
	}
	return total
}

func (sc StringCounter) MostCommon() (string, int) {
	var maxKey string
	var maxCount int
	for key, count := range sc {
		if count > maxCount {
			maxKey = key
			maxCount = count
		}
	}
	return maxKey, maxCount
}

func (sc StringCounter) String() string {
	return fmt.Sprintf("StringCounter%v", map[string]int(sc))
}

// 10. Complex struct with nested methods
type Car struct {
	Make    string
	Model   string
	Year    int
	Mileage int
	Engine  Engine
}

type Engine struct {
	Type       string
	Horsepower int
	Cylinders  int
}

func (c Car) String() string {
	return fmt.Sprintf("%d %s %s", c.Year, c.Make, c.Model)
}

func (c Car) Age() int {
	return 2024 - c.Year
}

func (c Car) IsVintage() bool {
	return c.Age() >= 25
}

func (c *Car) Drive(miles int) {
	c.Mileage += miles
	fmt.Printf("Drove %d miles. Total mileage: %d\n", miles, c.Mileage)
}

func (c *Car) Service() {
	fmt.Printf("Car %s has been serviced at %d miles\n", c.String(), c.Mileage)
}

func (e Engine) String() string {
	return fmt.Sprintf("%s Engine (%d HP, %d cylinders)", e.Type, e.Horsepower, e.Cylinders)
}

func (e Engine) PowerToWeightRatio(weight float64) float64 {
	return float64(e.Horsepower) / weight
}

// 11. Method sets and interface compatibility
type Writer interface {
	Write([]byte) (int, error)
}

type Buffer struct {
	data []byte
}

func (b *Buffer) Write(p []byte) (int, error) {
	b.data = append(b.data, p...)
	return len(p), nil
}

func (b *Buffer) String() string {
	return string(b.data)
}

func (b *Buffer) Len() int {
	return len(b.data)
}

func (b *Buffer) Reset() {
	b.data = b.data[:0]
}

// 12. Methods with different receiver types
type Counter struct {
	value int
}

func (c Counter) Value() int {
	return c.value
}

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) Decrement() {
	c.value--
}

func (c *Counter) Add(n int) {
	c.value += n
}

func (c *Counter) Reset() {
	c.value = 0
}

func (c Counter) String() string {
	return fmt.Sprintf("Counter{value: %d}", c.value)
}

// 13. Methods on function types
type Operation func(int, int) int

func (op Operation) Apply(a, b int) int {
	return op(a, b)
}

func (op Operation) ApplyToSlice(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	result := make([]int, len(slice)-1)
	for i := 0; i < len(slice)-1; i++ {
		result[i] = op(slice[i], slice[i+1])
	}
	return result
}

// 14. Methods for validation
type Email string

func (e Email) IsValid() bool {
	return strings.Contains(string(e), "@") && strings.Contains(string(e), ".")
}

func (e Email) Domain() string {
	parts := strings.Split(string(e), "@")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

func (e Email) Username() string {
	parts := strings.Split(string(e), "@")
	if len(parts) == 2 {
		return parts[0]
	}
	return ""
}

func (e Email) String() string {
	return string(e)
}

// 15. Methods with error handling
type SafeDivider struct {
	numerator   float64
	denominator float64
}

func (sd SafeDivider) Divide() (float64, error) {
	if sd.denominator == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return sd.numerator / sd.denominator, nil
}

func (sd SafeDivider) String() string {
	return fmt.Sprintf("SafeDivider{%f / %f}", sd.numerator, sd.denominator)
}

func (sd *SafeDivider) SetNumerator(n float64) {
	sd.numerator = n
}

func (sd *SafeDivider) SetDenominator(d float64) {
	sd.denominator = d
}

func main() {
	fmt.Println("=== GO METHODS COMPREHENSIVE GUIDE ===")

	// === BASIC METHODS ===
	fmt.Println("\n--- BASIC METHODS ---")
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
		Email:     "john.doe@example.com",
	}

	fmt.Printf("Person: %s\n", person)
	fmt.Printf("Full name: %s\n", person.FullName())
	fmt.Printf("Is adult: %t\n", person.IsAdult())

	/*
		JavaScript comparison:
		class Person {
			constructor(firstName, lastName, age, email) {
				this.firstName = firstName;
				this.lastName = lastName;
				this.age = age;
				this.email = email;
			}

			fullName() {
				return this.firstName + " " + this.lastName;
			}

			isAdult() {
				return this.age >= 18;
			}

			toString() {
				return `Person{Name: ${this.fullName()}, Age: ${this.age}, Email: ${this.email}}`;
			}
		}
	*/

	// === POINTER RECEIVER METHODS ===
	fmt.Println("\n--- POINTER RECEIVER METHODS ---")
	person.SetAge(30)
	person.SetEmail("john.doe.new@example.com")
	person.HaveBirthday()
	fmt.Printf("Updated person: %s\n", person)

	// === GEOMETRIC METHODS ===
	fmt.Println("\n--- GEOMETRIC METHODS ---")
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle: %s\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())
	fmt.Printf("Is square: %t\n", rect.IsSquare())

	rect.Scale(2.0)
	fmt.Printf("After scaling: %s\n", rect)
	fmt.Printf("New area: %.2f\n", rect.Area())

	circle := Circle{Radius: 3}
	fmt.Printf("Circle: %s\n", circle)
	fmt.Printf("Area: %.2f\n", circle.Area())
	fmt.Printf("Perimeter: %.2f\n", circle.Perimeter())
	fmt.Printf("Diameter: %.2f\n", circle.Diameter())

	// === METHOD CHAINING ===
	fmt.Println("\n--- METHOD CHAINING ---")
	account := &BankAccount{
		AccountNumber: "12345",
		Balance:       1000.00,
		Owner:         "Alice Johnson",
	}

	account.Deposit(500).Withdraw(200).Deposit(100)
	fmt.Printf("Final account: %s\n", account)

	// === METHODS ON CUSTOM TYPES ===
	fmt.Println("\n--- METHODS ON CUSTOM TYPES ---")
	temp := Temperature(25.0)
	fmt.Printf("Temperature: %s\n", temp)
	fmt.Printf("Fahrenheit: %.2f°F\n", temp.Fahrenheit())
	fmt.Printf("Kelvin: %.2f K\n", temp.Kelvin())

	// === METHODS ON SLICE TYPES ===
	fmt.Println("\n--- METHODS ON SLICE TYPES ---")
	numbers := IntSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Numbers: %s\n", numbers)
	fmt.Printf("Sum: %d\n", numbers.Sum())
	fmt.Printf("Average: %.2f\n", numbers.Average())
	fmt.Printf("Max: %d\n", numbers.Max())
	fmt.Printf("Min: %d\n", numbers.Min())
	fmt.Printf("Contains 5: %t\n", numbers.Contains(5))
	fmt.Printf("Contains 15: %t\n", numbers.Contains(15))

	// === METHODS ON MAP TYPES ===
	fmt.Println("\n--- METHODS ON MAP TYPES ---")
	counter := make(StringCounter)
	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}

	for _, word := range words {
		counter.Add(word)
	}

	fmt.Printf("Counter: %s\n", counter)
	fmt.Printf("Total count: %d\n", counter.Total())
	mostCommon, count := counter.MostCommon()
	fmt.Printf("Most common: %s (%d times)\n", mostCommon, count)

	// === COMPLEX STRUCT METHODS ===
	fmt.Println("\n--- COMPLEX STRUCT METHODS ---")
	car := Car{
		Make:    "Toyota",
		Model:   "Camry",
		Year:    2020,
		Mileage: 15000,
		Engine: Engine{
			Type:       "V6",
			Horsepower: 301,
			Cylinders:  6,
		},
	}

	fmt.Printf("Car: %s\n", car)
	fmt.Printf("Age: %d years\n", car.Age())
	fmt.Printf("Is vintage: %t\n", car.IsVintage())
	fmt.Printf("Engine: %s\n", car.Engine)
	fmt.Printf("Power-to-weight ratio: %.2f HP/lb\n", car.Engine.PowerToWeightRatio(3500))

	car.Drive(500)
	car.Service()

	// === BUFFER METHODS ===
	fmt.Println("\n--- BUFFER METHODS ---")
	buffer := &Buffer{}
	buffer.Write([]byte("Hello, "))
	buffer.Write([]byte("World!"))
	fmt.Printf("Buffer content: %s\n", buffer.String())
	fmt.Printf("Buffer length: %d\n", buffer.Len())

	buffer.Reset()
	fmt.Printf("After reset: %s\n", buffer.String())

	// === COUNTER METHODS ===
	fmt.Println("\n--- COUNTER METHODS ---")
	counter2 := &Counter{value: 0}
	fmt.Printf("Initial: %s\n", counter2)

	counter2.Increment()
	counter2.Increment()
	counter2.Add(5)
	fmt.Printf("After operations: %s\n", counter2)

	counter2.Decrement()
	fmt.Printf("After decrement: %s\n", counter2)

	// === METHODS ON FUNCTION TYPES ===
	fmt.Println("\n--- METHODS ON FUNCTION TYPES ---")
	add := Operation(func(a, b int) int { return a + b })
	multiply := Operation(func(a, b int) int { return a * b })

	fmt.Printf("add.Apply(5, 3) = %d\n", add.Apply(5, 3))
	fmt.Printf("multiply.Apply(5, 3) = %d\n", multiply.Apply(5, 3))

	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original: %v\n", nums)
	fmt.Printf("Add adjacent: %v\n", add.ApplyToSlice(nums))
	fmt.Printf("Multiply adjacent: %v\n", multiply.ApplyToSlice(nums))

	// === EMAIL VALIDATION ===
	fmt.Println("\n--- EMAIL VALIDATION ---")
	email1 := Email("user@example.com")
	email2 := Email("invalid-email")

	fmt.Printf("Email 1: %s\n", email1)
	fmt.Printf("Is valid: %t\n", email1.IsValid())
	fmt.Printf("Username: %s\n", email1.Username())
	fmt.Printf("Domain: %s\n", email1.Domain())

	fmt.Printf("Email 2: %s\n", email2)
	fmt.Printf("Is valid: %t\n", email2.IsValid())

	// === SAFE DIVIDER ===
	fmt.Println("\n--- SAFE DIVIDER ---")
	divider := SafeDivider{numerator: 10, denominator: 2}

	if result, err := divider.Divide(); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}

	divider.SetDenominator(0)
	if result, err := divider.Divide(); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}

	// === METHOD BEST PRACTICES ===
	fmt.Println("\n--- METHOD BEST PRACTICES ---")
	fmt.Println("1. Use pointer receivers for methods that modify the receiver")
	fmt.Println("2. Use value receivers for methods that don't modify the receiver")
	fmt.Println("3. Be consistent with receiver types within a type")
	fmt.Println("4. Use meaningful method names that describe actions")
	fmt.Println("5. Keep methods focused on a single responsibility")
	fmt.Println("6. Use String() method for custom string representation")
	fmt.Println("7. Consider method chaining for fluent interfaces")
	fmt.Println("8. Handle errors appropriately in methods")
	fmt.Println("9. Use methods to encapsulate behavior with data")
	fmt.Println("10. Consider interface compatibility when designing methods")
}
