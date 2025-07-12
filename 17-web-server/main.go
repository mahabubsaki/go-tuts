package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// === GO WEB SERVER COMPREHENSIVE GUIDE ===

/*
WEB SERVER PHILOSOPHY:
- Go's net/http package provides a powerful, production-ready HTTP server
- Built-in support for HTTP/2, TLS, and WebSocket
- Concurrent request handling with goroutines
- Simple yet flexible API for building web applications

COMPARISON WITH JAVASCRIPT:
// JavaScript (Node.js/Express)
const express = require('express');
const app = express();

app.get('/', (req, res) => {
  res.json({ message: 'Hello World!' });
});

app.listen(3000, () => {
  console.log('Server running on port 3000');
});

// Go: Built-in HTTP server
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"message": "Hello World!"})
})
log.Fatal(http.ListenAndServe(":8080", nil))
*/

// === DATA MODELS ===

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Product represents a product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	InStock     bool    `json:"in_stock"`
}

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// === IN-MEMORY DATA STORE ===

var (
	users = map[int]*User{
		1: {ID: 1, Name: "Alice Johnson", Email: "alice@example.com"},
		2: {ID: 2, Name: "Bob Smith", Email: "bob@example.com"},
		3: {ID: 3, Name: "Charlie Brown", Email: "charlie@example.com"},
	}

	products = map[int]*Product{
		1: {ID: 1, Name: "Laptop", Description: "High-performance laptop", Price: 1299.99, InStock: true},
		2: {ID: 2, Name: "Mouse", Description: "Wireless mouse", Price: 29.99, InStock: true},
		3: {ID: 3, Name: "Keyboard", Description: "Mechanical keyboard", Price: 149.99, InStock: false},
	}
)

// === BASIC HTTP HANDLERS ===

// 1. Simple Hello World handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! Time: %s\n", time.Now().Format(time.RFC3339))
}

// 2. JSON response handler
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	response := APIResponse{
		Success: true,
		Data: map[string]interface{}{
			"message":   "Hello from Go web server!",
			"timestamp": time.Now().Format(time.RFC3339),
			"method":    r.Method,
			"path":      r.URL.Path,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// 3. Request information handler
func requestInfoHandler(w http.ResponseWriter, r *http.Request) {
	info := map[string]interface{}{
		"method":      r.Method,
		"url":         r.URL.String(),
		"headers":     r.Header,
		"remote_addr": r.RemoteAddr,
		"user_agent":  r.UserAgent(),
		"host":        r.Host,
	}

	response := APIResponse{
		Success: true,
		Data:    info,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// === RESTful API HANDLERS ===

// 4. Users API handlers
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsersHandler(w, r)
	case http.MethodPost:
		createUserHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Convert map to slice
	userList := make([]*User, 0, len(users))
	for _, user := range users {
		userList = append(userList, user)
	}

	response := APIResponse{
		Success: true,
		Data:    userList,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User

	// Decode JSON body
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		response := APIResponse{
			Success: false,
			Error:   "Invalid JSON body",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Generate new ID
	newUser.ID = len(users) + 1
	users[newUser.ID] = &newUser

	response := APIResponse{
		Success: true,
		Data:    newUser,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// 5. Individual user handler
func userHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/users/")
	userID, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, exists := users[userID]
	if !exists {
		response := APIResponse{
			Success: false,
			Error:   "User not found",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := APIResponse{
		Success: true,
		Data:    user,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// 6. Products API handlers
func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getProductsHandler(w, r)
	case http.MethodPost:
		createProductHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Check for query parameters
	inStockOnly := r.URL.Query().Get("in_stock") == "true"

	productList := make([]*Product, 0, len(products))
	for _, product := range products {
		if !inStockOnly || product.InStock {
			productList = append(productList, product)
		}
	}

	response := APIResponse{
		Success: true,
		Data:    productList,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct Product

	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		response := APIResponse{
			Success: false,
			Error:   "Invalid JSON body",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Generate new ID
	newProduct.ID = len(products) + 1
	products[newProduct.ID] = &newProduct

	response := APIResponse{
		Success: true,
		Data:    newProduct,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// === MIDDLEWARE ===

// 7. Logging middleware
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler
		next(w, r)

		// Log request details
		log.Printf("%s %s %s %v",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			time.Since(start))
	}
}

// 8. CORS middleware
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// 9. Authentication middleware (simple example)
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for API key in header
		apiKey := r.Header.Get("X-API-Key")
		if apiKey != "secret-key" {
			response := APIResponse{
				Success: false,
				Error:   "Invalid API key",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		next(w, r)
	}
}

// === TEMPLATE RENDERING ===

// 10. HTML template handler
func templateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>Go Web Server</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .container { max-width: 800px; margin: 0 auto; }
        .user { padding: 10px; border: 1px solid #ccc; margin: 10px 0; }
        .product { padding: 10px; border: 1px solid #ddd; margin: 10px 0; }
        .in-stock { color: green; }
        .out-of-stock { color: red; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Go Web Server Dashboard</h1>
        
        <h2>Users</h2>
        {{range .Users}}
        <div class="user">
            <h3>{{.Name}}</h3>
            <p>Email: {{.Email}}</p>
            <p>ID: {{.ID}}</p>
        </div>
        {{end}}
        
        <h2>Products</h2>
        {{range .Products}}
        <div class="product">
            <h3>{{.Name}}</h3>
            <p>{{.Description}}</p>
            <p>Price: ${{printf "%.2f" .Price}}</p>
            <p class="{{if .InStock}}in-stock{{else}}out-of-stock{{end}}">
                {{if .InStock}}In Stock{{else}}Out of Stock{{end}}
            </p>
        </div>
        {{end}}
        
        <h2>Server Information</h2>
        <p>Current Time: {{.CurrentTime}}</p>
        <p>Request Method: {{.Method}}</p>
        <p>Request Path: {{.Path}}</p>
    </div>
</body>
</html>
`

	t, err := template.New("dashboard").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert maps to slices for template
	userList := make([]*User, 0, len(users))
	for _, user := range users {
		userList = append(userList, user)
	}

	productList := make([]*Product, 0, len(products))
	for _, product := range products {
		productList = append(productList, product)
	}

	data := struct {
		Users       []*User
		Products    []*Product
		CurrentTime string
		Method      string
		Path        string
	}{
		Users:       userList,
		Products:    productList,
		CurrentTime: time.Now().Format(time.RFC3339),
		Method:      r.Method,
		Path:        r.URL.Path,
	}

	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, data)
}

// === STATIC FILE SERVING ===

// 11. Static file handler
func staticHandler(w http.ResponseWriter, r *http.Request) {
	// Simple static content
	css := `
body {
    font-family: 'Segoe UI', Arial, sans-serif;
    line-height: 1.6;
    color: #333;
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
}

h1 {
    color: #2c3e50;
    border-bottom: 2px solid #3498db;
    padding-bottom: 10px;
}

.endpoint {
    background: #f8f9fa;
    border: 1px solid #dee2e6;
    border-radius: 4px;
    padding: 15px;
    margin: 10px 0;
}

.method {
    font-weight: bold;
    color: #495057;
}

.url {
    color: #007bff;
    font-family: monospace;
}

.description {
    color: #6c757d;
    font-style: italic;
}
`

	if strings.HasSuffix(r.URL.Path, ".css") {
		w.Header().Set("Content-Type", "text/css")
		w.Write([]byte(css))
		return
	}

	http.NotFound(w, r)
}

// === HEALTH CHECK ===

// 12. Health check handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	health := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"uptime":    "unknown", // In a real app, track uptime
		"version":   "1.0.0",
		"services": map[string]string{
			"database": "healthy",
			"cache":    "healthy",
			"queue":    "healthy",
		},
	}

	response := APIResponse{
		Success: true,
		Data:    health,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// === ERROR HANDLING ===

// 13. Not found handler
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	response := APIResponse{
		Success: false,
		Error:   "Endpoint not found",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(response)
}

// === MAIN SERVER ===

func main() {
	fmt.Println("=== GO WEB SERVER COMPREHENSIVE GUIDE ===")

	// === BASIC ROUTES ===
	http.HandleFunc("/", loggingMiddleware(corsMiddleware(helloHandler)))
	http.HandleFunc("/json", loggingMiddleware(corsMiddleware(jsonHandler)))
	http.HandleFunc("/request-info", loggingMiddleware(corsMiddleware(requestInfoHandler)))

	// === API ROUTES ===
	http.HandleFunc("/api/users", loggingMiddleware(corsMiddleware(usersHandler)))
	http.HandleFunc("/api/users/", loggingMiddleware(corsMiddleware(userHandler)))
	http.HandleFunc("/api/products", loggingMiddleware(corsMiddleware(productsHandler)))

	// === PROTECTED ROUTES ===
	http.HandleFunc("/api/admin/users", loggingMiddleware(corsMiddleware(authMiddleware(getUsersHandler))))

	// === TEMPLATE ROUTES ===
	http.HandleFunc("/dashboard", loggingMiddleware(templateHandler))

	// === STATIC FILES ===
	http.HandleFunc("/static/", staticHandler)

	// === HEALTH CHECK ===
	http.HandleFunc("/health", loggingMiddleware(corsMiddleware(healthHandler)))

	// === 404 HANDLER ===
	// Note: This won't work as expected with http.HandleFunc
	// For proper 404 handling, use a custom ServeMux or router

	// === SERVER CONFIGURATION ===
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// === START SERVER ===
	fmt.Println("\n=== SERVER STARTING ===")
	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("\nAvailable endpoints:")
	fmt.Println("GET    /                     - Hello World")
	fmt.Println("GET    /json                 - JSON response")
	fmt.Println("GET    /request-info         - Request information")
	fmt.Println("GET    /api/users            - List users")
	fmt.Println("POST   /api/users            - Create user")
	fmt.Println("GET    /api/users/{id}       - Get user by ID")
	fmt.Println("GET    /api/products         - List products")
	fmt.Println("GET    /api/products?in_stock=true - In-stock products")
	fmt.Println("POST   /api/products         - Create product")
	fmt.Println("GET    /api/admin/users      - Protected users endpoint (X-API-Key: secret-key)")
	fmt.Println("GET    /dashboard            - HTML dashboard")
	fmt.Println("GET    /static/styles.css    - CSS file")
	fmt.Println("GET    /health               - Health check")
	fmt.Println()
	fmt.Println("Example requests:")
	fmt.Println("curl http://localhost:8080/")
	fmt.Println("curl http://localhost:8080/api/users")
	fmt.Println("curl -X POST http://localhost:8080/api/users -H 'Content-Type: application/json' -d '{\"name\":\"John\",\"email\":\"john@example.com\"}'")
	fmt.Println("curl http://localhost:8080/api/admin/users -H 'X-API-Key: secret-key'")
	fmt.Println("curl http://localhost:8080/health")

	// Start server
	log.Fatal(server.ListenAndServe())
}

/*
RUNNING THE WEB SERVER:

1. Run the server:
   go run main.go

2. Test endpoints:
   curl http://localhost:8080/
   curl http://localhost:8080/api/users
   curl -X POST http://localhost:8080/api/users -H "Content-Type: application/json" -d '{"name":"John","email":"john@example.com"}'
   curl http://localhost:8080/api/admin/users -H "X-API-Key: secret-key"

3. Visit in browser:
   http://localhost:8080/dashboard

LEARNING POINTS:

1. HTTP SERVER BASICS:
   - http.HandleFunc for route registration
   - http.ResponseWriter for sending responses
   - http.Request for accessing request data

2. RESTful API DESIGN:
   - HTTP methods (GET, POST, PUT, DELETE)
   - JSON request/response handling
   - Status codes and error responses

3. MIDDLEWARE PATTERNS:
   - Function wrapping for cross-cutting concerns
   - Logging, CORS, authentication
   - Composable middleware chain

4. TEMPLATE RENDERING:
   - HTML template parsing and execution
   - Data binding and template functions
   - Dynamic content generation

5. STATIC FILE SERVING:
   - Serving CSS, JavaScript, images
   - Content-Type headers
   - File path handling

6. ERROR HANDLING:
   - HTTP status codes
   - JSON error responses
   - Graceful error recovery

7. SERVER CONFIGURATION:
   - Timeouts and limits
   - Security considerations
   - Production-ready settings

8. CONCURRENCY:
   - Goroutines handle each request
   - Thread-safe data access
   - Concurrent request processing

This demonstrates how to build production-ready web servers in Go
using the standard library and best practices.
*/
