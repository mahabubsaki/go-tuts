package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"
)

// === ADVANCED STRUCT CONCEPTS ===

// 1. Struct with JSON tags and validation
type User struct {
	ID        int       `json:"id" validate:"required,min=1"`
	Username  string    `json:"username" validate:"required,min=3,max=20"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"-" validate:"required,min=8"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Profile   *Profile  `json:"profile,omitempty"`
}

type Profile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"bio"`
	Avatar    string `json:"avatar"`
}

// 2. Struct with custom methods
func (u *User) String() string {
	return fmt.Sprintf("User{ID: %d, Username: %s, Email: %s}", u.ID, u.Username, u.Email)
}

func (u *User) FullName() string {
	if u.Profile != nil {
		return fmt.Sprintf("%s %s", u.Profile.FirstName, u.Profile.LastName)
	}
	return u.Username
}

func (u *User) IsValid() bool {
	return u.ID > 0 && u.Username != "" && u.Email != ""
}

func (u *User) UpdateProfile(firstName, lastName, bio string) {
	if u.Profile == nil {
		u.Profile = &Profile{}
	}
	u.Profile.FirstName = firstName
	u.Profile.LastName = lastName
	u.Profile.Bio = bio
	u.UpdatedAt = time.Now()
}

// 3. Struct embedding with method promotion
type Entity struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e *Entity) Touch() {
	e.UpdatedAt = time.Now()
}

func (e *Entity) IsNew() bool {
	return e.ID == 0
}

type Product struct {
	Entity
	Name        string
	Price       float64
	Description string
	Category    string
	InStock     bool
}

func (p *Product) ApplyDiscount(percentage float64) {
	p.Price = p.Price * (1 - percentage/100)
	p.Touch() // Promoted method from Entity
}

// 4. Struct with interface implementation
type Serializable interface {
	ToJSON() ([]byte, error)
	FromJSON([]byte) error
}

type Config struct {
	AppName     string          `json:"app_name"`
	Version     string          `json:"version"`
	Debug       bool            `json:"debug"`
	Database    DatabaseConfig  `json:"database"`
	Server      ServerConfig    `json:"server"`
	Features    map[string]bool `json:"features"`
	Limits      map[string]int  `json:"limits"`
	Environment string          `json:"environment"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type ServerConfig struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Timeout int    `json:"timeout"`
}

func (c *Config) ToJSON() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}

func (c *Config) FromJSON(data []byte) error {
	return json.Unmarshal(data, c)
}

// 5. Struct with generics (Go 1.18+)
type Container[T any] struct {
	Value T
	Meta  map[string]interface{}
}

func (c *Container[T]) SetMeta(key string, value interface{}) {
	if c.Meta == nil {
		c.Meta = make(map[string]interface{})
	}
	c.Meta[key] = value
}

func (c *Container[T]) GetMeta(key string) (interface{}, bool) {
	if c.Meta == nil {
		return nil, false
	}
	value, exists := c.Meta[key]
	return value, exists
}

// 6. Struct with validation
type Validator interface {
	Validate() error
}

type ValidationError struct {
	Field   string
	Message string
}

func (ve ValidationError) Error() string {
	return fmt.Sprintf("validation error in field '%s': %s", ve.Field, ve.Message)
}

type Account struct {
	ID       int     `json:"id"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
	Owner    string  `json:"owner"`
}

func (a *Account) Validate() error {
	if a.Balance < 0 {
		return ValidationError{Field: "Balance", Message: "cannot be negative"}
	}
	if a.Currency == "" {
		return ValidationError{Field: "Currency", Message: "is required"}
	}
	if len(a.Currency) != 3 {
		return ValidationError{Field: "Currency", Message: "must be 3 characters"}
	}
	if a.Owner == "" {
		return ValidationError{Field: "Owner", Message: "is required"}
	}
	return nil
}

// 7. Struct factory pattern
type ConnectionType int

const (
	HTTP ConnectionType = iota
	HTTPS
	WebSocket
)

type Connection struct {
	Type     ConnectionType
	URL      string
	Timeout  time.Duration
	Headers  map[string]string
	Metadata map[string]interface{}
}

type ConnectionBuilder struct {
	connection Connection
}

func NewConnectionBuilder() *ConnectionBuilder {
	return &ConnectionBuilder{
		connection: Connection{
			Type:     HTTP,
			Timeout:  30 * time.Second,
			Headers:  make(map[string]string),
			Metadata: make(map[string]interface{}),
		},
	}
}

func (cb *ConnectionBuilder) SetType(connType ConnectionType) *ConnectionBuilder {
	cb.connection.Type = connType
	return cb
}

func (cb *ConnectionBuilder) SetURL(url string) *ConnectionBuilder {
	cb.connection.URL = url
	return cb
}

func (cb *ConnectionBuilder) SetTimeout(timeout time.Duration) *ConnectionBuilder {
	cb.connection.Timeout = timeout
	return cb
}

func (cb *ConnectionBuilder) AddHeader(key, value string) *ConnectionBuilder {
	cb.connection.Headers[key] = value
	return cb
}

func (cb *ConnectionBuilder) AddMetadata(key string, value interface{}) *ConnectionBuilder {
	cb.connection.Metadata[key] = value
	return cb
}

func (cb *ConnectionBuilder) Build() Connection {
	return cb.connection
}

// 8. Struct with sync methods
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (sc *SafeCounter) Increment() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.count++
}

func (sc *SafeCounter) Get() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.count
}

func (sc *SafeCounter) Reset() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.count = 0
}

// 9. Struct with reflection capabilities
type Model struct {
	tableName string
	fields    map[string]reflect.Value
}

func NewModel(tableName string, data interface{}) *Model {
	m := &Model{
		tableName: tableName,
		fields:    make(map[string]reflect.Value),
	}

	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if field.IsExported() {
			m.fields[field.Name] = v.Field(i)
		}
	}

	return m
}

func (m *Model) GetField(name string) (interface{}, bool) {
	if field, exists := m.fields[name]; exists {
		return field.Interface(), true
	}
	return nil, false
}

func (m *Model) SetField(name string, value interface{}) bool {
	if field, exists := m.fields[name]; exists && field.CanSet() {
		field.Set(reflect.ValueOf(value))
		return true
	}
	return false
}

func (m *Model) GetTableName() string {
	return m.tableName
}

// 10. Struct with custom marshaling
type DateTime struct {
	time.Time
}

func (dt *DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(dt.Time.Format("2006-01-02 15:04:05"))
}

func (dt *DateTime) UnmarshalJSON(data []byte) error {
	var timeStr string
	if err := json.Unmarshal(data, &timeStr); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return err
	}

	dt.Time = t
	return nil
}

type Event struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	StartTime   DateTime `json:"start_time"`
	EndTime     DateTime `json:"end_time"`
	Location    string   `json:"location"`
}

func main() {
	fmt.Println("=== ADVANCED STRUCT CONCEPTS ===")

	// === JSON SERIALIZATION ===
	fmt.Println("\n--- JSON SERIALIZATION ---")

	user := User{
		ID:        1,
		Username:  "john_doe",
		Email:     "john@example.com",
		Password:  "secret123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Profile: &Profile{
			FirstName: "John",
			LastName:  "Doe",
			Bio:       "Software Developer",
			Avatar:    "avatar.jpg",
		},
	}

	// Serialize to JSON
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
	} else {
		fmt.Printf("User JSON:\n%s\n", string(jsonData))
	}

	// Deserialize from JSON
	var newUser User
	err = json.Unmarshal(jsonData, &newUser)
	if err != nil {
		fmt.Printf("Error unmarshaling: %v\n", err)
	} else {
		fmt.Printf("Deserialized user: %+v\n", newUser)
	}

	/*
		JavaScript comparison:
		// JavaScript object serialization
		const user = {
			id: 1,
			username: "john_doe",
			email: "john@example.com",
			createdAt: new Date(),
			profile: {
				firstName: "John",
				lastName: "Doe",
				bio: "Software Developer"
			}
		};

		const jsonData = JSON.stringify(user, null, 2);
		const newUser = JSON.parse(jsonData);
	*/

	// === STRUCT METHODS ===
	fmt.Println("\n--- STRUCT METHODS ---")

	fmt.Printf("User string: %s\n", user.String())
	fmt.Printf("Full name: %s\n", user.FullName())
	fmt.Printf("Is valid: %t\n", user.IsValid())

	user.UpdateProfile("Jane", "Smith", "Senior Developer")
	fmt.Printf("Updated full name: %s\n", user.FullName())

	// === STRUCT EMBEDDING ===
	fmt.Println("\n--- STRUCT EMBEDDING ---")

	product := Product{
		Entity: Entity{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:        "Laptop",
		Price:       1200.00,
		Description: "High-performance laptop",
		Category:    "Electronics",
		InStock:     true,
	}

	fmt.Printf("Product: %+v\n", product)
	fmt.Printf("Is new: %t\n", product.IsNew()) // Promoted method

	product.ApplyDiscount(10)
	fmt.Printf("Price after discount: %.2f\n", product.Price)

	// === INTERFACE IMPLEMENTATION ===
	fmt.Println("\n--- INTERFACE IMPLEMENTATION ---")

	config := Config{
		AppName:     "MyApp",
		Version:     "1.0.0",
		Debug:       true,
		Environment: "development",
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			Database: "myapp_db",
			Username: "admin",
			Password: "secret",
		},
		Server: ServerConfig{
			Host:    "0.0.0.0",
			Port:    8080,
			Timeout: 30,
		},
		Features: map[string]bool{
			"authentication": true,
			"logging":        true,
			"caching":        false,
		},
		Limits: map[string]int{
			"max_connections": 100,
			"rate_limit":      1000,
		},
	}

	configJSON, err := config.ToJSON()
	if err != nil {
		fmt.Printf("Error serializing config: %v\n", err)
	} else {
		fmt.Printf("Config JSON:\n%s\n", string(configJSON))
	}

	// === GENERICS ===
	fmt.Println("\n--- GENERICS ---")

	// String container
	stringContainer := Container[string]{
		Value: "Hello, World!",
	}
	stringContainer.SetMeta("type", "greeting")
	stringContainer.SetMeta("length", len(stringContainer.Value))

	fmt.Printf("String container: %+v\n", stringContainer)

	if metaType, exists := stringContainer.GetMeta("type"); exists {
		fmt.Printf("Meta type: %v\n", metaType)
	}

	// Integer container
	intContainer := Container[int]{
		Value: 42,
	}
	intContainer.SetMeta("type", "answer")
	intContainer.SetMeta("prime", false)

	fmt.Printf("Int container: %+v\n", intContainer)

	// === VALIDATION ===
	fmt.Println("\n--- VALIDATION ---")

	validAccount := Account{
		ID:       1,
		Balance:  1000.50,
		Currency: "USD",
		Owner:    "John Doe",
	}

	invalidAccount := Account{
		ID:       2,
		Balance:  -100.00,
		Currency: "US",
		Owner:    "",
	}

	fmt.Printf("Valid account: %+v\n", validAccount)
	if err := validAccount.Validate(); err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Println("Account is valid!")
	}

	fmt.Printf("Invalid account: %+v\n", invalidAccount)
	if err := invalidAccount.Validate(); err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Println("Account is valid!")
	}

	// === BUILDER PATTERN ===
	fmt.Println("\n--- BUILDER PATTERN ---")

	connection := NewConnectionBuilder().
		SetType(HTTPS).
		SetURL("https://api.example.com").
		SetTimeout(60*time.Second).
		AddHeader("Authorization", "Bearer token123").
		AddHeader("Content-Type", "application/json").
		AddMetadata("version", "1.0").
		AddMetadata("retry", true).
		Build()

	fmt.Printf("Connection: %+v\n", connection)

	// === THREAD-SAFE STRUCT ===
	fmt.Println("\n--- THREAD-SAFE STRUCT ---")

	counter := &SafeCounter{}

	// Simulate concurrent access
	for i := 0; i < 5; i++ {
		go func(n int) {
			for j := 0; j < 10; j++ {
				counter.Increment()
			}
		}(i)
	}

	// Wait a bit for goroutines to complete
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Counter value: %d\n", counter.Get())

	// === REFLECTION ===
	fmt.Println("\n--- REFLECTION ---")

	type Person struct {
		Name string
		Age  int
		City string
	}

	person := Person{
		Name: "Alice",
		Age:  30,
		City: "New York",
	}

	model := NewModel("people", person)
	fmt.Printf("Table name: %s\n", model.GetTableName())

	if name, exists := model.GetField("Name"); exists {
		fmt.Printf("Name field: %v\n", name)
	}

	model.SetField("Age", 31)
	if age, exists := model.GetField("Age"); exists {
		fmt.Printf("Updated age: %v\n", age)
	}

	// === CUSTOM MARSHALING ===
	fmt.Println("\n--- CUSTOM MARSHALING ---")

	event := Event{
		ID:          1,
		Name:        "Go Conference",
		Description: "Annual Go programming conference",
		StartTime:   DateTime{time.Date(2024, 6, 15, 9, 0, 0, 0, time.UTC)},
		EndTime:     DateTime{time.Date(2024, 6, 15, 17, 0, 0, 0, time.UTC)},
		Location:    "San Francisco",
	}

	eventJSON, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling event: %v\n", err)
	} else {
		fmt.Printf("Event JSON:\n%s\n", string(eventJSON))
	}

	// === STRUCT TAGS REFLECTION ===
	fmt.Println("\n--- STRUCT TAGS REFLECTION ---")

	userType := reflect.TypeOf(User{})
	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		jsonTag := field.Tag.Get("json")
		validateTag := field.Tag.Get("validate")

		fmt.Printf("Field: %s, JSON: %s, Validate: %s\n",
			field.Name, jsonTag, validateTag)
	}

	// === ADVANCED PATTERNS ===
	fmt.Println("\n--- ADVANCED PATTERNS ---")

	context := &Context{}
	context.SetState(ConcreteStateA{})
	fmt.Printf("Context state: %s\n", context.Request())

	context.SetState(ConcreteStateB{})
	fmt.Printf("Context state: %s\n", context.Request())

	// === PERFORMANCE CONSIDERATIONS ===
	fmt.Println("\n--- PERFORMANCE CONSIDERATIONS ---")

	// Small struct - pass by value
	type SmallStruct struct {
		A int
		B int
	}

	// Large struct - pass by pointer
	type LargeStruct struct {
		Data [1000]int
		Name string
		Meta map[string]interface{}
	}

	small := SmallStruct{A: 1, B: 2}
	large := LargeStruct{Name: "large"}

	fmt.Printf("Small struct size: %d bytes\n", reflect.TypeOf(small).Size())
	fmt.Printf("Large struct size: %d bytes\n", reflect.TypeOf(large).Size())

	// === BEST PRACTICES ===
	fmt.Println("\n--- ADVANCED STRUCT BEST PRACTICES ---")
	fmt.Println("1. Use struct embedding for composition")
	fmt.Println("2. Implement interfaces for polymorphism")
	fmt.Println("3. Use builder pattern for complex construction")
	fmt.Println("4. Add validation methods to structs")
	fmt.Println("5. Use JSON tags for serialization")
	fmt.Println("6. Consider thread safety for shared structs")
	fmt.Println("7. Use reflection carefully (performance impact)")
	fmt.Println("8. Custom marshaling for special formatting")
	fmt.Println("9. Use generics for type-safe containers")
	fmt.Println("10. Pass large structs by pointer")
}

// Visitor pattern types
type Visitor interface {
	Visit(interface{}) interface{}
}

type StringVisitor struct{}

func (sv StringVisitor) Visit(v interface{}) interface{} {
	if s, ok := v.(string); ok {
		return strings.ToUpper(s)
	}
	return v
}

// State pattern types
type State interface {
	Handle() string
}

type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() string {
	return c.state.Handle()
}

type ConcreteStateA struct{}

func (cs ConcreteStateA) Handle() string { return "State A" }

type ConcreteStateB struct{}

func (cs ConcreteStateB) Handle() string { return "State B" }
