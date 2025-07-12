package main

import (
	"fmt"
	"io"
	"math"
	"sort"
	"strings"
)

// === INTERFACES IN GO ===

// 1. Basic interface definition
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Drawable interface {
	Draw()
}

// 2. Types implementing interfaces
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Draw() {
	fmt.Printf("Drawing rectangle: %.2f x %.2f\n", r.Width, r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Draw() {
	fmt.Printf("Drawing circle with radius: %.2f\n", c.Radius)
}

// 3. Interface composition
type DrawableShape interface {
	Shape
	Drawable
}

// 4. Empty interface
func PrintAny(value interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

// 5. Interface with methods
type Writer interface {
	Write([]byte) (int, error)
}

type Reader interface {
	Read([]byte) (int, error)
}

type ReadWriter interface {
	Reader
	Writer
}

// 6. Custom implementations
type StringWriter struct {
	data strings.Builder
}

func (sw *StringWriter) Write(p []byte) (int, error) {
	return sw.data.Write(p)
}

func (sw *StringWriter) String() string {
	return sw.data.String()
}

type Buffer struct {
	data []byte
}

func (b *Buffer) Write(p []byte) (int, error) {
	b.data = append(b.data, p...)
	return len(p), nil
}

func (b *Buffer) Read(p []byte) (int, error) {
	if len(b.data) == 0 {
		return 0, io.EOF
	}

	n := copy(p, b.data)
	b.data = b.data[n:]
	return n, nil
}

func (b *Buffer) String() string {
	return string(b.data)
}

// 7. Type assertions
func ProcessValue(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("String: %s (length: %d)\n", v, len(v))
	case int:
		fmt.Printf("Integer: %d\n", v)
	case float64:
		fmt.Printf("Float: %.2f\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	case []int:
		fmt.Printf("Integer slice: %v\n", v)
	case Shape:
		fmt.Printf("Shape with area: %.2f\n", v.Area())
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

// 8. Interface polymorphism
func CalculateArea(shapes []Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func DrawShapes(shapes []DrawableShape) {
	for i, shape := range shapes {
		fmt.Printf("Shape %d: ", i+1)
		shape.Draw()
		fmt.Printf("  Area: %.2f, Perimeter: %.2f\n", shape.Area(), shape.Perimeter())
	}
}

// 9. Interface for sorting
type Person struct {
	Name string
	Age  int
}

type People []Person

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// 10. Custom interface for business logic
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
	GetTransactionFee() float64
}

type CreditCardProcessor struct {
	CardNumber string
	FeeRate    float64
}

func (ccp CreditCardProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("Processing credit card payment: $%.2f\n", amount)
	return nil
}

func (ccp CreditCardProcessor) GetTransactionFee() float64 {
	return ccp.FeeRate
}

type PayPalProcessor struct {
	Email   string
	FeeRate float64
}

func (pp PayPalProcessor) ProcessPayment(amount float64) error {
	fmt.Printf("Processing PayPal payment: $%.2f\n", amount)
	return nil
}

func (pp PayPalProcessor) GetTransactionFee() float64 {
	return pp.FeeRate
}

// 11. Interface for database operations
type Repository interface {
	Save(entity interface{}) error
	FindByID(id string) (interface{}, error)
	FindAll() ([]interface{}, error)
	Delete(id string) error
}

type MemoryRepository struct {
	data map[string]interface{}
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		data: make(map[string]interface{}),
	}
}

func (mr *MemoryRepository) Save(entity interface{}) error {
	// Simplified: using string representation as ID
	id := fmt.Sprintf("%v", entity)
	mr.data[id] = entity
	fmt.Printf("Saved entity: %v\n", entity)
	return nil
}

func (mr *MemoryRepository) FindByID(id string) (interface{}, error) {
	if entity, exists := mr.data[id]; exists {
		return entity, nil
	}
	return nil, fmt.Errorf("entity with ID %s not found", id)
}

func (mr *MemoryRepository) FindAll() ([]interface{}, error) {
	var entities []interface{}
	for _, entity := range mr.data {
		entities = append(entities, entity)
	}
	return entities, nil
}

func (mr *MemoryRepository) Delete(id string) error {
	if _, exists := mr.data[id]; exists {
		delete(mr.data, id)
		fmt.Printf("Deleted entity with ID: %s\n", id)
		return nil
	}
	return fmt.Errorf("entity with ID %s not found", id)
}

// 12. Interface for observers
type Observer interface {
	Update(message string)
}

type Subject interface {
	Subscribe(observer Observer)
	Unsubscribe(observer Observer)
	Notify(message string)
}

type EmailObserver struct {
	Email string
}

func (eo EmailObserver) Update(message string) {
	fmt.Printf("Email notification to %s: %s\n", eo.Email, message)
}

type SMSObserver struct {
	Phone string
}

func (so SMSObserver) Update(message string) {
	fmt.Printf("SMS notification to %s: %s\n", so.Phone, message)
}

type NewsPublisher struct {
	observers []Observer
}

func (np *NewsPublisher) Subscribe(observer Observer) {
	np.observers = append(np.observers, observer)
}

func (np *NewsPublisher) Unsubscribe(observer Observer) {
	for i, obs := range np.observers {
		if obs == observer {
			np.observers = append(np.observers[:i], np.observers[i+1:]...)
			break
		}
	}
}

func (np *NewsPublisher) Notify(message string) {
	for _, observer := range np.observers {
		observer.Update(message)
	}
}

// 13. Interface for strategy pattern
type DiscountStrategy interface {
	CalculateDiscount(amount float64) float64
}

type PercentageDiscount struct {
	Percentage float64
}

func (pd PercentageDiscount) CalculateDiscount(amount float64) float64 {
	return amount * pd.Percentage / 100
}

type FixedDiscount struct {
	Amount float64
}

func (fd FixedDiscount) CalculateDiscount(amount float64) float64 {
	if amount > fd.Amount {
		return fd.Amount
	}
	return amount
}

type ShoppingCart struct {
	items    []string
	total    float64
	discount DiscountStrategy
}

func (sc *ShoppingCart) SetDiscountStrategy(discount DiscountStrategy) {
	sc.discount = discount
}

func (sc *ShoppingCart) CalculateTotal() float64 {
	if sc.discount != nil {
		return sc.total - sc.discount.CalculateDiscount(sc.total)
	}
	return sc.total
}

// 14. Interface for command pattern
type Command interface {
	Execute() error
	Undo() error
}

type LightOnCommand struct {
	location string
}

func (loc LightOnCommand) Execute() error {
	fmt.Printf("Turning on light in %s\n", loc.location)
	return nil
}

func (loc LightOnCommand) Undo() error {
	fmt.Printf("Turning off light in %s\n", loc.location)
	return nil
}

type LightOffCommand struct {
	location string
}

func (loc LightOffCommand) Execute() error {
	fmt.Printf("Turning off light in %s\n", loc.location)
	return nil
}

func (loc LightOffCommand) Undo() error {
	fmt.Printf("Turning on light in %s\n", loc.location)
	return nil
}

type RemoteControl struct {
	commands []Command
	history  []Command
}

func (rc *RemoteControl) SetCommand(slot int, command Command) {
	if slot >= len(rc.commands) {
		// Extend slice if needed
		for len(rc.commands) <= slot {
			rc.commands = append(rc.commands, nil)
		}
	}
	rc.commands[slot] = command
}

func (rc *RemoteControl) PressButton(slot int) error {
	if slot < len(rc.commands) && rc.commands[slot] != nil {
		err := rc.commands[slot].Execute()
		if err == nil {
			rc.history = append(rc.history, rc.commands[slot])
		}
		return err
	}
	return fmt.Errorf("no command set for slot %d", slot)
}

func (rc *RemoteControl) PressUndo() error {
	if len(rc.history) > 0 {
		lastCommand := rc.history[len(rc.history)-1]
		rc.history = rc.history[:len(rc.history)-1]
		return lastCommand.Undo()
	}
	return fmt.Errorf("no command to undo")
}

// 15. Interface for factory pattern
type Animal interface {
	Speak() string
	Move() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Move() string {
	return "Running"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Move() string {
	return "Sneaking"
}

type AnimalFactory interface {
	CreateAnimal(name string) Animal
}

type DogFactory struct{}

func (df DogFactory) CreateAnimal(name string) Animal {
	return Dog{Name: name}
}

type CatFactory struct{}

func (cf CatFactory) CreateAnimal(name string) Animal {
	return Cat{Name: name}
}

// 16. Interface for validation
type Validator interface {
	Validate() []string
}

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) Validate() []string {
	var errors []string

	if len(u.Name) < 2 {
		errors = append(errors, "name must be at least 2 characters")
	}

	if !strings.Contains(u.Email, "@") {
		errors = append(errors, "email must contain @ symbol")
	}

	if u.Age < 0 || u.Age > 150 {
		errors = append(errors, "age must be between 0 and 150")
	}

	return errors
}

// 17. Interface for configuration
type ConfigProvider interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
}

type MapConfig struct {
	data map[string]interface{}
}

func (mc MapConfig) GetString(key string) string {
	if value, exists := mc.data[key]; exists {
		if str, ok := value.(string); ok {
			return str
		}
	}
	return ""
}

func (mc MapConfig) GetInt(key string) int {
	if value, exists := mc.data[key]; exists {
		if i, ok := value.(int); ok {
			return i
		}
	}
	return 0
}

func (mc MapConfig) GetBool(key string) bool {
	if value, exists := mc.data[key]; exists {
		if b, ok := value.(bool); ok {
			return b
		}
	}
	return false
}

func main() {
	fmt.Println("=== GO INTERFACES COMPREHENSIVE GUIDE ===")

	// === BASIC INTERFACES ===
	fmt.Println("\n--- BASIC INTERFACES ---")
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 2}

	shapes := []Shape{rect, circle}

	fmt.Printf("Total area: %.2f\n", CalculateArea(shapes))

	// Interface polymorphism
	var shape Shape = rect
	fmt.Printf("Rectangle area: %.2f\n", shape.Area())

	shape = circle
	fmt.Printf("Circle area: %.2f\n", shape.Area())

	/*
		JavaScript comparison:
		// JavaScript doesn't have interfaces, but uses duck typing
		// and protocols. TypeScript provides interface support:

		interface Shape {
			area(): number;
			perimeter(): number;
		}

		class Rectangle implements Shape {
			constructor(private width: number, private height: number) {}

			area(): number {
				return this.width * this.height;
			}

			perimeter(): number {
				return 2 * (this.width + this.height);
			}
		}

		// Duck typing in JavaScript:
		function calculateArea(shape) {
			if (shape.area && typeof shape.area === 'function') {
				return shape.area();
			}
			throw new Error('Object must have an area method');
		}
	*/

	// === INTERFACE COMPOSITION ===
	fmt.Println("\n--- INTERFACE COMPOSITION ---")
	drawableShapes := []DrawableShape{rect, circle}
	DrawShapes(drawableShapes)

	// === EMPTY INTERFACE ===
	fmt.Println("\n--- EMPTY INTERFACE ---")
	PrintAny(42)
	PrintAny("Hello, World!")
	PrintAny(3.14159)
	PrintAny(true)
	PrintAny([]int{1, 2, 3})

	// === TYPE ASSERTIONS ===
	fmt.Println("\n--- TYPE ASSERTIONS ---")
	values := []interface{}{
		"Hello",
		42,
		3.14159,
		true,
		[]int{1, 2, 3, 4, 5},
		rect,
		"Another string",
	}

	for _, value := range values {
		ProcessValue(value)
	}

	// === WRITER INTERFACE ===
	fmt.Println("\n--- WRITER INTERFACE ---")
	var writer Writer = &StringWriter{}
	writer.Write([]byte("Hello, "))
	writer.Write([]byte("World!"))

	if sw, ok := writer.(*StringWriter); ok {
		fmt.Printf("StringWriter content: %s\n", sw.String())
	}

	// === READWRITER INTERFACE ===
	fmt.Println("\n--- READWRITER INTERFACE ---")
	buffer := &Buffer{}
	buffer.Write([]byte("Hello, Buffer!"))

	readData := make([]byte, 6)
	n, err := buffer.Read(readData)
	if err != nil {
		fmt.Printf("Error reading: %v\n", err)
	} else {
		fmt.Printf("Read %d bytes: %s\n", n, string(readData[:n]))
	}

	// === SORTING INTERFACE ===
	fmt.Println("\n--- SORTING INTERFACE ---")
	people := People{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"Diana", 28},
	}

	fmt.Printf("Before sorting: %v\n", people)
	sort.Sort(people)
	fmt.Printf("After sorting by age: %v\n", people)

	// === PAYMENT PROCESSOR ===
	fmt.Println("\n--- PAYMENT PROCESSOR ---")
	processors := []PaymentProcessor{
		CreditCardProcessor{CardNumber: "1234-5678-9012-3456", FeeRate: 2.5},
		PayPalProcessor{Email: "user@example.com", FeeRate: 3.0},
	}

	amount := 100.0
	for _, processor := range processors {
		processor.ProcessPayment(amount)
		fee := processor.GetTransactionFee()
		fmt.Printf("Transaction fee: %.2f%%\n", fee)
		fmt.Printf("Total cost: $%.2f\n", amount+amount*fee/100)
		fmt.Println()
	}

	// === REPOSITORY PATTERN ===
	fmt.Println("\n--- REPOSITORY PATTERN ---")
	repo := NewMemoryRepository()

	// Save entities
	repo.Save(User{Name: "John", Email: "john@example.com", Age: 30})
	repo.Save(User{Name: "Jane", Email: "jane@example.com", Age: 25})

	// Find all entities
	entities, _ := repo.FindAll()
	fmt.Printf("All entities: %v\n", entities)

	// === OBSERVER PATTERN ===
	fmt.Println("\n--- OBSERVER PATTERN ---")
	publisher := &NewsPublisher{}

	emailObs := EmailObserver{Email: "user@example.com"}
	smsObs := SMSObserver{Phone: "+1234567890"}

	publisher.Subscribe(emailObs)
	publisher.Subscribe(smsObs)

	publisher.Notify("Breaking news: Go interfaces are awesome!")

	// === STRATEGY PATTERN ===
	fmt.Println("\n--- STRATEGY PATTERN ---")
	cart := &ShoppingCart{
		items: []string{"item1", "item2", "item3"},
		total: 100.0,
	}

	// No discount
	fmt.Printf("Total without discount: $%.2f\n", cart.CalculateTotal())

	// Percentage discount
	cart.SetDiscountStrategy(PercentageDiscount{Percentage: 10})
	fmt.Printf("Total with 10%% discount: $%.2f\n", cart.CalculateTotal())

	// Fixed discount
	cart.SetDiscountStrategy(FixedDiscount{Amount: 15})
	fmt.Printf("Total with $15 fixed discount: $%.2f\n", cart.CalculateTotal())

	// === COMMAND PATTERN ===
	fmt.Println("\n--- COMMAND PATTERN ---")
	remote := &RemoteControl{}

	livingRoomOn := LightOnCommand{location: "Living Room"}
	kitchenOff := LightOffCommand{location: "Kitchen"}

	remote.SetCommand(0, livingRoomOn)
	remote.SetCommand(1, kitchenOff)

	remote.PressButton(0)
	remote.PressButton(1)
	remote.PressUndo()
	remote.PressUndo()

	// === FACTORY PATTERN ===
	fmt.Println("\n--- FACTORY PATTERN ---")
	factories := map[string]AnimalFactory{
		"dog": DogFactory{},
		"cat": CatFactory{},
	}

	animals := []Animal{
		factories["dog"].CreateAnimal("Buddy"),
		factories["cat"].CreateAnimal("Whiskers"),
	}

	for _, animal := range animals {
		fmt.Printf("Animal says: %s and is %s\n", animal.Speak(), animal.Move())
	}

	// === VALIDATION ===
	fmt.Println("\n--- VALIDATION ---")
	users := []User{
		{Name: "John Doe", Email: "john@example.com", Age: 30},
		{Name: "J", Email: "invalid-email", Age: -5},
		{Name: "Jane Smith", Email: "jane@example.com", Age: 25},
	}

	for _, user := range users {
		errors := user.Validate()
		if len(errors) == 0 {
			fmt.Printf("User %s is valid\n", user.Name)
		} else {
			fmt.Printf("User %s has errors: %v\n", user.Name, errors)
		}
	}

	// === CONFIGURATION ===
	fmt.Println("\n--- CONFIGURATION ---")
	config := MapConfig{
		data: map[string]interface{}{
			"app_name":        "My App",
			"port":            8080,
			"debug":           true,
			"max_connections": 100,
		},
	}

	fmt.Printf("App name: %s\n", config.GetString("app_name"))
	fmt.Printf("Port: %d\n", config.GetInt("port"))
	fmt.Printf("Debug mode: %t\n", config.GetBool("debug"))
	fmt.Printf("Max connections: %d\n", config.GetInt("max_connections"))

	// === INTERFACE BEST PRACTICES ===
	fmt.Println("\n--- INTERFACE BEST PRACTICES ---")
	fmt.Println("1. Keep interfaces small and focused")
	fmt.Println("2. Define interfaces at the point of use")
	fmt.Println("3. Use composition over inheritance")
	fmt.Println("4. Accept interfaces, return concrete types")
	fmt.Println("5. Use empty interface{} sparingly")
	fmt.Println("6. Use type assertions carefully")
	fmt.Println("7. Implement interfaces implicitly")
	fmt.Println("8. Use interfaces for testing and mocking")
	fmt.Println("9. Consider interface segregation")
	fmt.Println("10. Document interface contracts clearly")
}
