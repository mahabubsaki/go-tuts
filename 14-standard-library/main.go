package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"io"
	"log"
	"math"
	"math/big"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// === GO STANDARD LIBRARY COMPREHENSIVE GUIDE ===

// Data structures for demonstrations
type Person struct {
	Name    string    `json:"name" xml:"name"`
	Age     int       `json:"age" xml:"age"`
	Email   string    `json:"email" xml:"email"`
	Created time.Time `json:"created" xml:"created"`
}

type Company struct {
	Name      string   `json:"name" xml:"name"`
	Employees []Person `json:"employees" xml:"employees>person"`
}

func main() {
	fmt.Println("=== GO STANDARD LIBRARY COMPREHENSIVE GUIDE ===")

	// === FMT PACKAGE ===
	fmt.Println("\n--- FMT PACKAGE ---")

	name := "Go"
	version := 1.21
	active := true

	// Print functions
	fmt.Print("Hello, ")
	fmt.Print("World!\n")

	fmt.Println("This is a line")
	fmt.Printf("Language: %s, Version: %.1f, Active: %t\n", name, version, active)

	// Format specifiers
	fmt.Printf("String: %s\n", "hello")
	fmt.Printf("Integer: %d\n", 42)
	fmt.Printf("Float: %.2f\n", 3.14159)
	fmt.Printf("Boolean: %t\n", true)
	fmt.Printf("Pointer: %p\n", &name)
	fmt.Printf("Type: %T\n", name)
	fmt.Printf("Value: %v\n", name)
	fmt.Printf("Detailed: %+v\n", Person{Name: "John", Age: 30})
	fmt.Printf("Go syntax: %#v\n", []int{1, 2, 3})

	// String formatting
	formatted := fmt.Sprintf("Formatted: %s %d", "test", 123)
	fmt.Println(formatted)

	// Scan functions
	var input string
	fmt.Print("Enter something (or press Enter to skip): ")
	// fmt.Scanln(&input) // Commented to avoid blocking
	input = "sample input"
	fmt.Printf("You entered: %s\n", input)

	/*
		JavaScript comparison:
		// JavaScript string interpolation
		const name = "Go";
		const version = 1.21;
		const active = true;

		console.log(`Language: ${name}, Version: ${version}, Active: ${active}`);
		console.log("String:", "hello");
		console.log("Number:", 42);
		console.log("Boolean:", true);
		console.log("Type:", typeof name);
		console.log("Value:", name);
	*/

	// === STRINGS PACKAGE ===
	fmt.Println("\n--- STRINGS PACKAGE ---")

	text := "Hello, World! Go is awesome!"

	// String operations
	fmt.Printf("Contains 'Go': %t\n", strings.Contains(text, "Go"))
	fmt.Printf("Starts with 'Hello': %t\n", strings.HasPrefix(text, "Hello"))
	fmt.Printf("Ends with 'some!': %t\n", strings.HasSuffix(text, "some!"))
	fmt.Printf("Index of 'Go': %d\n", strings.Index(text, "Go"))
	fmt.Printf("Count of 'o': %d\n", strings.Count(text, "o"))

	// String manipulation
	fmt.Printf("Uppercase: %s\n", strings.ToUpper(text))
	fmt.Printf("Lowercase: %s\n", strings.ToLower(text))
	fmt.Printf("Replace 'Go' with 'JavaScript': %s\n", strings.Replace(text, "Go", "JavaScript", 1))
	fmt.Printf("Replace all 'o' with '0': %s\n", strings.ReplaceAll(text, "o", "0"))

	// String splitting and joining
	parts := strings.Split(text, " ")
	fmt.Printf("Split by space: %v\n", parts)
	joined := strings.Join(parts, "-")
	fmt.Printf("Joined with '-': %s\n", joined)

	// String trimming
	whitespace := "  \t\n  Hello, World!  \n\t  "
	fmt.Printf("Trimmed: '%s'\n", strings.TrimSpace(whitespace))
	fmt.Printf("Trim prefix: '%s'\n", strings.TrimPrefix(text, "Hello, "))
	fmt.Printf("Trim suffix: '%s'\n", strings.TrimSuffix(text, "!"))

	// String building
	var builder strings.Builder
	builder.WriteString("Building ")
	builder.WriteString("a ")
	builder.WriteString("string")
	fmt.Printf("Built string: %s\n", builder.String())

	// === STRCONV PACKAGE ===
	fmt.Println("\n--- STRCONV PACKAGE ---")

	// String to number conversion
	str := "42"
	if num, err := strconv.Atoi(str); err == nil {
		fmt.Printf("String to int: %d\n", num)
	}

	if float, err := strconv.ParseFloat("3.14159", 64); err == nil {
		fmt.Printf("String to float: %.2f\n", float)
	}

	if boolean, err := strconv.ParseBool("true"); err == nil {
		fmt.Printf("String to bool: %t\n", boolean)
	}

	// Number to string conversion
	fmt.Printf("Int to string: %s\n", strconv.Itoa(42))
	fmt.Printf("Float to string: %s\n", strconv.FormatFloat(3.14159, 'f', 2, 64))
	fmt.Printf("Bool to string: %s\n", strconv.FormatBool(true))

	// Base conversions
	fmt.Printf("Binary: %s\n", strconv.FormatInt(42, 2))
	fmt.Printf("Octal: %s\n", strconv.FormatInt(42, 8))
	fmt.Printf("Hexadecimal: %s\n", strconv.FormatInt(42, 16))

	// === MATH PACKAGE ===
	fmt.Println("\n--- MATH PACKAGE ---")

	// Basic math functions
	fmt.Printf("Absolute: %.2f\n", math.Abs(-3.14))
	fmt.Printf("Ceiling: %.2f\n", math.Ceil(3.14))
	fmt.Printf("Floor: %.2f\n", math.Floor(3.14))
	fmt.Printf("Round: %.2f\n", math.Round(3.14))
	fmt.Printf("Max: %.2f\n", math.Max(3.14, 2.71))
	fmt.Printf("Min: %.2f\n", math.Min(3.14, 2.71))

	// Power and logarithms
	fmt.Printf("Power: %.2f\n", math.Pow(2, 3))
	fmt.Printf("Square root: %.2f\n", math.Sqrt(16))
	fmt.Printf("Logarithm: %.2f\n", math.Log(math.E))
	fmt.Printf("Log10: %.2f\n", math.Log10(100))

	// Trigonometric functions
	fmt.Printf("Sin(π/2): %.2f\n", math.Sin(math.Pi/2))
	fmt.Printf("Cos(π): %.2f\n", math.Cos(math.Pi))
	fmt.Printf("Tan(π/4): %.2f\n", math.Tan(math.Pi/4))

	// Constants
	fmt.Printf("π: %.6f\n", math.Pi)
	fmt.Printf("e: %.6f\n", math.E)

	// Random numbers (using crypto/rand for security)
	randomBytes := make([]byte, 4)
	rand.Read(randomBytes)
	fmt.Printf("Random bytes: %v\n", randomBytes)

	// Big numbers
	big1 := new(big.Int)
	big1.SetString("12345678901234567890", 10)
	big2 := new(big.Int)
	big2.SetString("98765432109876543210", 10)
	result := new(big.Int).Add(big1, big2)
	fmt.Printf("Big number addition: %s\n", result.String())

	// === TIME PACKAGE ===
	fmt.Println("\n--- TIME PACKAGE ---")

	// Current time
	now := time.Now()
	fmt.Printf("Current time: %s\n", now.Format("2006-01-02 15:04:05"))

	// Time creation
	specific := time.Date(2024, time.June, 15, 14, 30, 0, 0, time.UTC)
	fmt.Printf("Specific time: %s\n", specific.Format("2006-01-02 15:04:05 MST"))

	// Time formatting
	fmt.Printf("RFC3339: %s\n", now.Format(time.RFC3339))
	fmt.Printf("Kitchen: %s\n", now.Format(time.Kitchen))
	fmt.Printf("Custom: %s\n", now.Format("January 2, 2006 at 3:04 PM"))

	// Time parsing
	if parsed, err := time.Parse("2006-01-02", "2024-06-15"); err == nil {
		fmt.Printf("Parsed time: %s\n", parsed.Format("January 2, 2006"))
	}

	// Time calculations
	future := now.Add(24 * time.Hour)
	fmt.Printf("24 hours later: %s\n", future.Format("2006-01-02 15:04:05"))

	duration := future.Sub(now)
	fmt.Printf("Duration: %s\n", duration)

	// Time comparison
	fmt.Printf("Before: %t\n", now.Before(future))
	fmt.Printf("After: %t\n", now.After(future))
	fmt.Printf("Equal: %t\n", now.Equal(future))

	// Unix timestamp
	fmt.Printf("Unix timestamp: %d\n", now.Unix())
	fromUnix := time.Unix(now.Unix(), 0)
	fmt.Printf("From Unix: %s\n", fromUnix.Format("2006-01-02 15:04:05"))

	// === OS PACKAGE ===
	fmt.Println("\n--- OS PACKAGE ---")

	// Environment variables
	fmt.Printf("PATH: %s\n", os.Getenv("PATH")[:50]+"...")
	os.Setenv("MY_VAR", "test_value")
	fmt.Printf("MY_VAR: %s\n", os.Getenv("MY_VAR"))

	// Working directory
	if wd, err := os.Getwd(); err == nil {
		fmt.Printf("Working directory: %s\n", wd)
	}

	// Command line arguments
	fmt.Printf("Command line args: %v\n", os.Args)

	// File operations
	fileName := "test.txt"
	content := "Hello, World!"

	// Write file
	if err := os.WriteFile(fileName, []byte(content), 0644); err == nil {
		fmt.Printf("File written successfully\n")

		// Read file
		if data, err := os.ReadFile(fileName); err == nil {
			fmt.Printf("File content: %s\n", string(data))
		}

		// File info
		if info, err := os.Stat(fileName); err == nil {
			fmt.Printf("File size: %d bytes\n", info.Size())
			fmt.Printf("File mode: %s\n", info.Mode())
			fmt.Printf("Modified: %s\n", info.ModTime().Format("2006-01-02 15:04:05"))
		}

		// Remove file
		os.Remove(fileName)
		fmt.Printf("File removed\n")
	}

	// === FILEPATH PACKAGE ===
	fmt.Println("\n--- FILEPATH PACKAGE ---")

	path := "/home/user/documents/file.txt"

	fmt.Printf("Directory: %s\n", filepath.Dir(path))
	fmt.Printf("Base name: %s\n", filepath.Base(path))
	fmt.Printf("Extension: %s\n", filepath.Ext(path))

	// Path joining
	joinedPath := filepath.Join("home", "user", "documents", "file.txt")
	fmt.Printf("Joined path: %s\n", joinedPath)

	// Path cleaning
	messy := "home//user/../user/./documents/file.txt"
	fmt.Printf("Cleaned path: %s\n", filepath.Clean(messy))

	// === IO PACKAGE ===
	fmt.Println("\n--- IO PACKAGE ---")

	// String reader
	reader := strings.NewReader("Hello, World!")
	readBuffer := make([]byte, 5)

	for {
		n, err := reader.Read(readBuffer)
		if err == io.EOF {
			break
		}
		fmt.Printf("Read %d bytes: %s\n", n, string(readBuffer[:n]))
	}

	// Bytes buffer
	var buf bytes.Buffer
	buf.WriteString("Hello, ")
	buf.WriteString("World!")
	fmt.Printf("Buffer content: %s\n", buf.String())

	// Copy
	src := strings.NewReader("Copy this text")
	dst := &bytes.Buffer{}
	io.Copy(dst, src)
	fmt.Printf("Copied: %s\n", dst.String())

	// === BUFIO PACKAGE ===
	fmt.Println("\n--- BUFIO PACKAGE ---")

	// Scanner
	text = "line1\nline2\nline3"
	scanner := bufio.NewScanner(strings.NewReader(text))

	fmt.Println("Scanning lines:")
	for scanner.Scan() {
		fmt.Printf("  %s\n", scanner.Text())
	}

	// Writer
	var writeBuffer bytes.Buffer
	writer := bufio.NewWriter(&writeBuffer)
	writer.WriteString("Buffered ")
	writer.WriteString("writing")
	writer.Flush()
	fmt.Printf("Buffered content: %s\n", writeBuffer.String())

	// === REGEXP PACKAGE ===
	fmt.Println("\n--- REGEXP PACKAGE ---")

	// Compile regex
	re := regexp.MustCompile(`\d+`)
	text = "I have 123 apples and 456 oranges"

	// Find matches
	matches := re.FindAllString(text, -1)
	fmt.Printf("Number matches: %v\n", matches)

	// Replace
	replaced := re.ReplaceAllString(text, "X")
	fmt.Printf("Replaced: %s\n", replaced)

	// Email validation
	emailRe := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	emails := []string{"test@example.com", "invalid.email", "user@domain.co.uk"}

	for _, email := range emails {
		fmt.Printf("'%s' is valid: %t\n", email, emailRe.MatchString(email))
	}

	// === SORT PACKAGE ===
	fmt.Println("\n--- SORT PACKAGE ---")

	// Sort integers
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
	sort.Ints(numbers)
	fmt.Printf("Sorted numbers: %v\n", numbers)

	// Sort strings
	words := []string{"banana", "apple", "cherry", "date"}
	sort.Strings(words)
	fmt.Printf("Sorted words: %v\n", words)

	// Custom sort
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Printf("Sorted by age: %v\n", people)

	// Search
	index := sort.SearchInts(numbers, 5)
	fmt.Printf("Index of 5: %d\n", index)

	// === JSON PACKAGE ===
	fmt.Println("\n--- JSON PACKAGE ---")

	// Struct to JSON
	person := Person{
		Name:    "John Doe",
		Age:     30,
		Email:   "john@example.com",
		Created: time.Now(),
	}

	jsonData, err := json.MarshalIndent(person, "", "  ")
	if err == nil {
		fmt.Printf("JSON:\n%s\n", string(jsonData))
	}

	// JSON to struct
	var parsedPerson Person
	err = json.Unmarshal(jsonData, &parsedPerson)
	if err == nil {
		fmt.Printf("Parsed: %+v\n", parsedPerson)
	}

	// Map to JSON
	data := map[string]interface{}{
		"name":     "Go",
		"version":  1.21,
		"active":   true,
		"features": []string{"fast", "simple", "reliable"},
	}

	if mapJSON, err := json.Marshal(data); err == nil {
		fmt.Printf("Map JSON: %s\n", string(mapJSON))
	}

	// === XML PACKAGE ===
	fmt.Println("\n--- XML PACKAGE ---")

	// Struct to XML
	company := Company{
		Name: "Tech Corp",
		Employees: []Person{
			{Name: "Alice", Age: 30, Email: "alice@techcorp.com"},
			{Name: "Bob", Age: 25, Email: "bob@techcorp.com"},
		},
	}

	xmlData, err := xml.MarshalIndent(company, "", "  ")
	if err == nil {
		fmt.Printf("XML:\n%s\n", string(xmlData))
	}

	// XML to struct
	var parsedCompany Company
	err = xml.Unmarshal(xmlData, &parsedCompany)
	if err == nil {
		fmt.Printf("Parsed XML: %+v\n", parsedCompany)
	}

	// === CSV PACKAGE ===
	fmt.Println("\n--- CSV PACKAGE ---")

	// Write CSV
	var csvBuffer bytes.Buffer
	csvWriter := csv.NewWriter(&csvBuffer)

	records := [][]string{
		{"Name", "Age", "Email"},
		{"Alice", "30", "alice@example.com"},
		{"Bob", "25", "bob@example.com"},
		{"Charlie", "35", "charlie@example.com"},
	}

	for _, record := range records {
		csvWriter.Write(record)
	}
	csvWriter.Flush()

	fmt.Printf("CSV:\n%s\n", csvBuffer.String())

	// Read CSV
	csvReader := csv.NewReader(strings.NewReader(csvBuffer.String()))
	csvRecords, err := csvReader.ReadAll()
	if err == nil {
		fmt.Printf("Parsed CSV: %v\n", csvRecords)
	}

	// === TEMPLATE PACKAGE ===
	fmt.Println("\n--- TEMPLATE PACKAGE ---")

	// Text template
	tmplText := "Hello, {{.Name}}! You are {{.Age}} years old."
	tmpl := template.Must(template.New("greeting").Parse(tmplText))

	var textResult bytes.Buffer
	tmpl.Execute(&textResult, person)
	fmt.Printf("Text template: %s\n", textResult.String())

	// HTML template
	htmlTmpl := `<h1>{{.Name}}</h1><p>Age: {{.Age}}</p><p>Email: {{.Email}}</p>`
	htmlTemplate := template.Must(template.New("profile").Parse(htmlTmpl))

	var htmlResult bytes.Buffer
	htmlTemplate.Execute(&htmlResult, person)
	fmt.Printf("HTML template: %s\n", htmlResult.String())

	// === NET/HTTP PACKAGE ===
	fmt.Println("\n--- NET/HTTP PACKAGE ---")

	// HTTP client (commented to avoid external dependency)
	/*
		client := &http.Client{
			Timeout: 5 * time.Second,
		}

		resp, err := client.Get("https://api.github.com/users/octocat")
		if err == nil {
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("Response: %s\n", string(body)[:100]+"...")
		}
	*/
	fmt.Println("HTTP client example (commented out)")

	// URL parsing
	urlStr := "https://example.com/path?param1=value1&param2=value2"
	parsedURL, err := url.Parse(urlStr)
	if err == nil {
		fmt.Printf("Scheme: %s\n", parsedURL.Scheme)
		fmt.Printf("Host: %s\n", parsedURL.Host)
		fmt.Printf("Path: %s\n", parsedURL.Path)
		fmt.Printf("Query: %s\n", parsedURL.RawQuery)
		fmt.Printf("Param1: %s\n", parsedURL.Query().Get("param1"))
	}

	// === CRYPTO PACKAGE ===
	fmt.Println("\n--- CRYPTO PACKAGE ---")

	cryptoData := []byte("Hello, World!")

	// MD5 hash
	md5Hash := md5.Sum(cryptoData)
	fmt.Printf("MD5: %x\n", md5Hash)

	// SHA256 hash
	sha256Hash := sha256.Sum256(cryptoData)
	fmt.Printf("SHA256: %x\n", sha256Hash)

	// Base64 encoding
	encoded := base64.StdEncoding.EncodeToString(cryptoData)
	fmt.Printf("Base64: %s\n", encoded)

	// Base64 decoding
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err == nil {
		fmt.Printf("Decoded: %s\n", string(decoded))
	}

	// === REFLECT PACKAGE ===
	fmt.Println("\n--- REFLECT PACKAGE ---")

	// Type reflection
	value := 42
	reflectValue := reflect.ValueOf(value)
	reflectType := reflect.TypeOf(value)

	fmt.Printf("Value: %v\n", reflectValue)
	fmt.Printf("Type: %v\n", reflectType)
	fmt.Printf("Kind: %v\n", reflectType.Kind())

	// Struct reflection
	personValue := reflect.ValueOf(person)
	personType := reflect.TypeOf(person)

	fmt.Printf("Struct name: %s\n", personType.Name())
	fmt.Printf("Number of fields: %d\n", personType.NumField())

	for i := 0; i < personType.NumField(); i++ {
		field := personType.Field(i)
		value := personValue.Field(i)
		fmt.Printf("Field %d: %s (%s) = %v\n", i, field.Name, field.Type, value)
	}

	// === SYNC PACKAGE ===
	fmt.Println("\n--- SYNC PACKAGE ---")

	// Mutex
	var mu sync.Mutex
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("Counter: %d\n", counter)

	// Once
	var once sync.Once

	for i := 0; i < 3; i++ {
		once.Do(func() {
			fmt.Println("Initialized once")
		})
	}

	// Pool
	pool := &sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}

	poolBuf := pool.Get().([]byte)
	pool.Put(poolBuf)
	fmt.Printf("Pool buffer length: %d\n", len(poolBuf))

	// === CONTEXT PACKAGE ===
	fmt.Println("\n--- CONTEXT PACKAGE ---")

	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Operation completed within timeout")
	case <-ctx.Done():
		fmt.Println("Operation timed out")
	}

	// Context with value
	ctx = context.WithValue(context.Background(), "userID", "12345")
	if userID := ctx.Value("userID"); userID != nil {
		fmt.Printf("User ID from context: %s\n", userID)
	}

	// === LOG PACKAGE ===
	fmt.Println("\n--- LOG PACKAGE ---")

	// Default logger
	log.Printf("This is a log message: %s", "test")

	// Custom logger
	var logBuffer bytes.Buffer
	logger := log.New(&logBuffer, "CUSTOM: ", log.LstdFlags|log.Lshortfile)
	logger.Printf("Custom log message: %d", 42)
	fmt.Printf("Log output: %s", logBuffer.String())

	// === STANDARD LIBRARY BEST PRACTICES ===
	fmt.Println("\n--- STANDARD LIBRARY BEST PRACTICES ---")
	fmt.Println("1. Use fmt for formatted output")
	fmt.Println("2. Use strings package for string manipulation")
	fmt.Println("3. Use strconv for type conversions")
	fmt.Println("4. Use time package for time operations")
	fmt.Println("5. Use os package for system operations")
	fmt.Println("6. Use io package for input/output operations")
	fmt.Println("7. Use regexp for pattern matching")
	fmt.Println("8. Use sort for sorting operations")
	fmt.Println("9. Use json/xml for data serialization")
	fmt.Println("10. Use net/http for HTTP operations")
	fmt.Println("11. Use crypto for cryptographic operations")
	fmt.Println("12. Use reflect for runtime type inspection")
	fmt.Println("13. Use sync for synchronization")
	fmt.Println("14. Use context for cancellation and values")
	fmt.Println("15. Use log for logging")

	// === COMMON PATTERNS ===
	fmt.Println("\n--- COMMON PATTERNS ---")
	fmt.Println("1. Error handling with multiple return values")
	fmt.Println("2. Defer for resource cleanup")
	fmt.Println("3. Interface{} for generic values")
	fmt.Println("4. Struct tags for metadata")
	fmt.Println("5. Context for request-scoped data")
	fmt.Println("6. Channels for communication")
	fmt.Println("7. Goroutines for concurrency")
	fmt.Println("8. Mutexes for synchronization")
	fmt.Println("9. Closures for callbacks")
	fmt.Println("10. Type assertions for interface handling")

	fmt.Println("\nThe Go standard library provides comprehensive tools for modern software development!")
}
