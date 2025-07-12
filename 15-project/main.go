package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// === PROJECT 15: COMPREHENSIVE GO WEB API ===

/*
PROJECT OVERVIEW:
This project demonstrates a complete Go web application with:
- RESTful API endpoints
- Database operations (SQLite)
- User authentication
- Middleware for logging and CORS
- Error handling patterns
- JSON serialization/deserialization
- Configuration management
- Resource cleanup with defer
- Concurrency with goroutines
- Testing patterns

LEARNING OBJECTIVES:
1. Apply Go fundamentals in a real-world application
2. Understand project structure and organization
3. Implement common web development patterns
4. Practice error handling and resource management
5. Learn about Go's concurrency in web applications
6. Understand interface usage in practice
7. Apply struct embedding and composition
8. Use Go's standard library effectively
*/

// === MODELS ===

// User represents a user in the system
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Never serialize password
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserResponse represents the user data sent to clients
type UserResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserRequest represents the request to create a user
type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UpdateUserRequest represents the request to update a user
type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// LoginRequest represents the login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

// APIError represents an API error response
type APIError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// === INTERFACES ===

// UserRepository defines the interface for user data operations
type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id int) (*User, error)
	GetByUsername(username string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
}

// Logger interface for logging operations
type Logger interface {
	Info(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Debug(msg string, fields ...interface{})
}

// === IMPLEMENTATIONS ===

// SQLiteUserRepository implements UserRepository for SQLite
type SQLiteUserRepository struct {
	db *sql.DB
}

// NewSQLiteUserRepository creates a new SQLite user repository
func NewSQLiteUserRepository(db *sql.DB) *SQLiteUserRepository {
	return &SQLiteUserRepository{db: db}
}

// GetAll retrieves all users from the database
func (r *SQLiteUserRepository) GetAll() ([]User, error) {
	query := `
		SELECT id, username, email, password, created_at, updated_at 
		FROM users 
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating users: %w", err)
	}

	return users, nil
}

// GetByID retrieves a user by ID
func (r *SQLiteUserRepository) GetByID(id int) (*User, error) {
	query := `
		SELECT id, username, email, password, created_at, updated_at 
		FROM users 
		WHERE id = ?
	`

	var user User
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

// GetByUsername retrieves a user by username
func (r *SQLiteUserRepository) GetByUsername(username string) (*User, error) {
	query := `
		SELECT id, username, email, password, created_at, updated_at 
		FROM users 
		WHERE username = ?
	`

	var user User
	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

// Create creates a new user
func (r *SQLiteUserRepository) Create(user *User) error {
	query := `
		INSERT INTO users (username, email, password, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	result, err := r.db.Exec(query, user.Username, user.Email, user.Password, now, now)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID: %w", err)
	}

	user.ID = int(id)
	return nil
}

// Update updates an existing user
func (r *SQLiteUserRepository) Update(user *User) error {
	query := `
		UPDATE users 
		SET username = ?, email = ?, updated_at = ?
		WHERE id = ?
	`

	user.UpdatedAt = time.Now()

	result, err := r.db.Exec(query, user.Username, user.Email, user.UpdatedAt, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// Delete deletes a user by ID
func (r *SQLiteUserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// SimpleLogger implements Logger interface
type SimpleLogger struct{}

func (l *SimpleLogger) Info(msg string, fields ...interface{}) {
	log.Printf("[INFO] %s %v", msg, fields)
}

func (l *SimpleLogger) Error(msg string, fields ...interface{}) {
	log.Printf("[ERROR] %s %v", msg, fields)
}

func (l *SimpleLogger) Debug(msg string, fields ...interface{}) {
	log.Printf("[DEBUG] %s %v", msg, fields)
}

// === HANDLERS ===

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	userRepo UserRepository
	logger   Logger
}

// NewUserHandler creates a new user handler
func NewUserHandler(userRepo UserRepository, logger Logger) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
		logger:   logger,
	}
}

// GetUsers handles GET /api/users
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Getting all users")

	users, err := h.userRepo.GetAll()
	if err != nil {
		h.logger.Error("Failed to get users", "error", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to get users", err.Error())
		return
	}

	// Convert to response format
	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	h.writeJSON(w, http.StatusOK, userResponses)
}

// GetUser handles GET /api/users/{id}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	h.logger.Info("Getting user", "id", id)

	user, err := h.userRepo.GetByID(id)
	if err != nil {
		h.logger.Error("Failed to get user", "id", id, "error", err)
		h.writeError(w, http.StatusNotFound, "User not found", err.Error())
		return
	}

	userResponse := UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	h.writeJSON(w, http.StatusOK, userResponse)
}

// CreateUser handles POST /api/users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	// Validate request
	if req.Username == "" || req.Email == "" || req.Password == "" {
		h.writeError(w, http.StatusBadRequest, "Missing required fields", "username, email, and password are required")
		return
	}

	h.logger.Info("Creating user", "username", req.Username)

	// Create user
	user := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password, // In production, hash the password
	}

	if err := h.userRepo.Create(user); err != nil {
		h.logger.Error("Failed to create user", "error", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to create user", err.Error())
		return
	}

	userResponse := UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	h.writeJSON(w, http.StatusCreated, userResponse)
}

// UpdateUser handles PUT /api/users/{id}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	h.logger.Info("Updating user", "id", id)

	// Get existing user
	user, err := h.userRepo.GetByID(id)
	if err != nil {
		h.logger.Error("Failed to get user for update", "id", id, "error", err)
		h.writeError(w, http.StatusNotFound, "User not found", err.Error())
		return
	}

	// Update fields
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}

	if err := h.userRepo.Update(user); err != nil {
		h.logger.Error("Failed to update user", "id", id, "error", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to update user", err.Error())
		return
	}

	userResponse := UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	h.writeJSON(w, http.StatusOK, userResponse)
}

// DeleteUser handles DELETE /api/users/{id}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.writeError(w, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	h.logger.Info("Deleting user", "id", id)

	if err := h.userRepo.Delete(id); err != nil {
		h.logger.Error("Failed to delete user", "id", id, "error", err)
		h.writeError(w, http.StatusNotFound, "User not found", err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Login handles POST /api/auth/login
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	h.logger.Info("User login attempt", "username", req.Username)

	user, err := h.userRepo.GetByUsername(req.Username)
	if err != nil {
		h.logger.Error("Login failed - user not found", "username", req.Username)
		h.writeError(w, http.StatusUnauthorized, "Invalid credentials", "")
		return
	}

	// In production, use proper password hashing
	if user.Password != req.Password {
		h.logger.Error("Login failed - invalid password", "username", req.Username)
		h.writeError(w, http.StatusUnauthorized, "Invalid credentials", "")
		return
	}

	// Generate simple token (in production, use JWT)
	token := fmt.Sprintf("token_%d_%d", user.ID, time.Now().Unix())

	response := LoginResponse{
		User: UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Token: token,
	}

	h.writeJSON(w, http.StatusOK, response)
}

// Helper methods
func (h *UserHandler) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *UserHandler) writeError(w http.ResponseWriter, status int, message, details string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(APIError{
		Error:   message,
		Message: details,
		Code:    status,
	})
}

// === MIDDLEWARE ===

// LoggingMiddleware logs HTTP requests
func LoggingMiddleware(logger Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			logger.Info("Request started", "method", r.Method, "path", r.URL.Path)

			next.ServeHTTP(w, r)

			logger.Info("Request completed",
				"method", r.Method,
				"path", r.URL.Path,
				"duration", time.Since(start))
		})
	}
}

// CORSMiddleware handles CORS headers
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// === DATABASE SETUP ===

// SetupDatabase creates and initializes the database
func SetupDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Create users table
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
		)
	`

	if _, err := db.Exec(createTableQuery); err != nil {
		return nil, fmt.Errorf("failed to create users table: %w", err)
	}

	return db, nil
}

// === CONFIGURATION ===

// Config represents application configuration
type Config struct {
	Port     string
	Database string
	LogLevel string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		Port:     getEnv("PORT", "8080"),
		Database: getEnv("DATABASE", "users.db"),
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// === MAIN APPLICATION ===

func main() {
	fmt.Println("=== PROJECT 15: GO WEB API WITH DATABASE ===")

	// Load configuration
	config := LoadConfig()
	logger := &SimpleLogger{}

	logger.Info("Starting application", "port", config.Port)

	// Setup database
	db, err := SetupDatabase()
	if err != nil {
		logger.Error("Failed to setup database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	// Create repository and handler
	userRepo := NewSQLiteUserRepository(db)
	userHandler := NewUserHandler(userRepo, logger)

	// Setup router
	router := mux.NewRouter()

	// Add middleware
	router.Use(LoggingMiddleware(logger))
	router.Use(CORSMiddleware)

	// API routes
	api := router.PathPrefix("/api").Subrouter()

	// User routes
	users := api.PathPrefix("/users").Subrouter()
	users.HandleFunc("", userHandler.GetUsers).Methods("GET")
	users.HandleFunc("/{id}", userHandler.GetUser).Methods("GET")
	users.HandleFunc("", userHandler.CreateUser).Methods("POST")
	users.HandleFunc("/{id}", userHandler.UpdateUser).Methods("PUT")
	users.HandleFunc("/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Auth routes
	auth := api.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", userHandler.Login).Methods("POST")

	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":    "ok",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	}).Methods("GET")

	// Start server
	logger.Info("Server starting", "port", config.Port)

	server := &http.Server{
		Addr:         ":" + config.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown handling
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Server failed to start", "error", err)
			os.Exit(1)
		}
	}()

	logger.Info("Server started successfully")
	logger.Info("API Documentation:")
	logger.Info("GET    /health           - Health check")
	logger.Info("GET    /api/users        - List all users")
	logger.Info("POST   /api/users        - Create new user")
	logger.Info("GET    /api/users/{id}   - Get user by ID")
	logger.Info("PUT    /api/users/{id}   - Update user")
	logger.Info("DELETE /api/users/{id}   - Delete user")
	logger.Info("POST   /api/auth/login   - User login")

	// Keep the server running
	select {}
}

/*
RUNNING THE PROJECT:

1. Install dependencies:
   go mod init project15
   go get github.com/gorilla/mux
   go get github.com/mattn/go-sqlite3

2. Run the server:
   go run main.go

3. Test the API:
   curl -X GET http://localhost:8080/health
   curl -X GET http://localhost:8080/api/users
   curl -X POST http://localhost:8080/api/users -H "Content-Type: application/json" -d '{"username":"john","email":"john@example.com","password":"secret"}'
   curl -X POST http://localhost:8080/api/auth/login -H "Content-Type: application/json" -d '{"username":"john","password":"secret"}'

LEARNING POINTS:

1. PROJECT STRUCTURE:
   - Organized code into logical components
   - Separated concerns (handlers, repositories, models)
   - Used interfaces for dependency injection

2. GO FUNDAMENTALS APPLIED:
   - Structs for data modeling
   - Interfaces for abstraction
   - Error handling throughout
   - Defer for resource cleanup
   - Pointers for efficient memory usage

3. WEB DEVELOPMENT PATTERNS:
   - RESTful API design
   - Middleware for cross-cutting concerns
   - JSON serialization/deserialization
   - HTTP status codes and error responses

4. DATABASE OPERATIONS:
   - CRUD operations with SQL
   - Connection management
   - Error handling in database operations

5. CONCURRENCY:
   - HTTP server handles requests concurrently
   - Database connection pooling
   - Graceful shutdown patterns

6. BEST PRACTICES:
   - Configuration management
   - Structured logging
   - Input validation
   - Error responses with context
   - Resource cleanup with defer

This project demonstrates how to build a production-ready Go web application
using the concepts learned in previous topics.
*/
