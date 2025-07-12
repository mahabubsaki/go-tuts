# Project 15: Go Web API with Database

A comprehensive project demonstrating advanced Go concepts including:

## Features Implemented
- RESTful API with Gorilla Mux
- Database operations with SQLite
- User authentication and authorization
- Error handling and middleware
- JSON serialization/deserialization
- Configuration management
- Logging and monitoring
- Testing with test coverage

## Project Structure
```
15-project/
├── main.go
├── go.mod
├── go.sum
├── users.db
└── README.md
```

## Learning Objectives
- Apply Go fundamentals in a real-world scenario
- Understand project structure and organization
- Implement common web development patterns
- Practice error handling and resource management
- Learn about Go's concurrency in web applications

## Getting Started
1. Navigate to the project directory
2. Run `go mod init project15`
3. Install dependencies: `go mod tidy`
4. Run the server: `go run main.go`
5. Test the API endpoints

## API Endpoints
- `GET /health` - Health check
- `GET /api/users` - List all users
- `POST /api/users` - Create a new user
- `GET /api/users/{id}` - Get user by ID
- `PUT /api/users/{id}` - Update user
- `DELETE /api/users/{id}` - Delete user
- `POST /api/auth/login` - User login

This project consolidates learning from all previous topics and demonstrates production-ready Go code.
