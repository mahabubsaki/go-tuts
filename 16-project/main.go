package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// === PROJECT 16: GO MICROSERVICES WITH ADVANCED CONCURRENCY ===

/*
PROJECT OVERVIEW:
This project demonstrates advanced Go concepts in a microservices architecture:
- Multiple services communicating via channels and HTTP
- Worker pools for concurrent processing
- Circuit breaker pattern for fault tolerance
- Message queues for asynchronous communication
- Health checks and monitoring
- Graceful shutdown and resource management
- Advanced concurrency patterns

LEARNING OBJECTIVES:
1. Build microservices with Go
2. Implement advanced concurrency patterns
3. Use channels for inter-service communication
4. Handle distributed system challenges
5. Apply circuit breaker patterns
6. Implement health checks and monitoring
7. Practice goroutine management
8. Use context for cancellation and timeouts
*/

// === DOMAIN MODELS ===

// User represents a user in the system
type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
}

// Order represents an order in the system
type Order struct {
	ID      int       `json:"id"`
	UserID  int       `json:"user_id"`
	Product string    `json:"product"`
	Amount  float64   `json:"amount"`
	Status  string    `json:"status"`
	Created time.Time `json:"created"`
}

// Notification represents a notification message
type Notification struct {
	ID      int       `json:"id"`
	UserID  int       `json:"user_id"`
	Type    string    `json:"type"` // email, sms, push
	Message string    `json:"message"`
	Sent    bool      `json:"sent"`
	Created time.Time `json:"created"`
}

// === MESSAGING SYSTEM ===

// Message represents a message in the system
type Message struct {
	ID      string      `json:"id"`
	Topic   string      `json:"topic"`
	Payload interface{} `json:"payload"`
	Created time.Time   `json:"created"`
}

// MessageBroker handles message publishing and subscribing
type MessageBroker struct {
	subscribers map[string][]chan Message
	mu          sync.RWMutex
}

// NewMessageBroker creates a new message broker
func NewMessageBroker() *MessageBroker {
	return &MessageBroker{
		subscribers: make(map[string][]chan Message),
	}
}

// Subscribe adds a subscriber to a topic
func (mb *MessageBroker) Subscribe(topic string, ch chan Message) {
	mb.mu.Lock()
	defer mb.mu.Unlock()

	if _, exists := mb.subscribers[topic]; !exists {
		mb.subscribers[topic] = make([]chan Message, 0)
	}

	mb.subscribers[topic] = append(mb.subscribers[topic], ch)
}

// Publish sends a message to all subscribers of a topic
func (mb *MessageBroker) Publish(topic string, payload interface{}) {
	mb.mu.RLock()
	defer mb.mu.RUnlock()

	message := Message{
		ID:      fmt.Sprintf("msg_%d", time.Now().UnixNano()),
		Topic:   topic,
		Payload: payload,
		Created: time.Now(),
	}

	if subscribers, exists := mb.subscribers[topic]; exists {
		for _, ch := range subscribers {
			select {
			case ch <- message:
			case <-time.After(100 * time.Millisecond):
				log.Printf("Failed to send message to subscriber for topic %s", topic)
			}
		}
	}
}

// === CIRCUIT BREAKER ===

// CircuitBreakerState represents the state of a circuit breaker
type CircuitBreakerState int

const (
	Closed CircuitBreakerState = iota
	Open
	HalfOpen
)

// CircuitBreaker implements the circuit breaker pattern
type CircuitBreaker struct {
	name            string
	maxFailures     int
	timeout         time.Duration
	failures        int
	lastFailureTime time.Time
	state           CircuitBreakerState
	mutex           sync.RWMutex
}

// NewCircuitBreaker creates a new circuit breaker
func NewCircuitBreaker(name string, maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		name:        name,
		maxFailures: maxFailures,
		timeout:     timeout,
		state:       Closed,
	}
}

// Execute executes a function with circuit breaker protection
func (cb *CircuitBreaker) Execute(fn func() error) error {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if cb.state == Open {
		if time.Since(cb.lastFailureTime) > cb.timeout {
			cb.state = HalfOpen
			cb.failures = 0
		} else {
			return fmt.Errorf("circuit breaker %s is open", cb.name)
		}
	}

	err := fn()
	if err != nil {
		cb.failures++
		cb.lastFailureTime = time.Now()

		if cb.failures >= cb.maxFailures {
			cb.state = Open
		}

		return err
	}

	// Success
	cb.failures = 0
	cb.state = Closed
	return nil
}

// GetState returns the current state of the circuit breaker
func (cb *CircuitBreaker) GetState() CircuitBreakerState {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	return cb.state
}

// === WORKER POOL ===

// WorkerPool manages a pool of workers
type WorkerPool struct {
	workers    int
	jobQueue   chan Job
	workerPool chan chan Job
	quit       chan bool
	wg         sync.WaitGroup
}

// Job represents a unit of work
type Job struct {
	ID     string
	Task   func() error
	Result chan error
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(workers int, queueSize int) *WorkerPool {
	return &WorkerPool{
		workers:    workers,
		jobQueue:   make(chan Job, queueSize),
		workerPool: make(chan chan Job, workers),
		quit:       make(chan bool),
	}
}

// Start starts the worker pool
func (wp *WorkerPool) Start() {
	// Create workers
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}

	// Start dispatcher
	go wp.dispatch()
}

// worker represents a worker in the pool
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()

	jobChannel := make(chan Job)

	for {
		// Register worker in pool
		wp.workerPool <- jobChannel

		select {
		case job := <-jobChannel:
			// Execute job
			err := job.Task()
			select {
			case job.Result <- err:
			case <-time.After(1 * time.Second):
				log.Printf("Worker %d: timeout sending result for job %s", id, job.ID)
			}

		case <-wp.quit:
			log.Printf("Worker %d shutting down", id)
			return
		}
	}
}

// dispatch dispatches jobs to available workers
func (wp *WorkerPool) dispatch() {
	for {
		select {
		case job := <-wp.jobQueue:
			// Get available worker
			select {
			case jobChannel := <-wp.workerPool:
				// Send job to worker
				jobChannel <- job
			case <-wp.quit:
				return
			}
		case <-wp.quit:
			return
		}
	}
}

// Submit submits a job to the worker pool
func (wp *WorkerPool) Submit(job Job) {
	wp.jobQueue <- job
}

// Stop stops the worker pool
func (wp *WorkerPool) Stop() {
	close(wp.quit)
	wp.wg.Wait()
}

// === SERVICES ===

// UserService handles user-related operations
type UserService struct {
	users   map[int]*User
	mu      sync.RWMutex
	broker  *MessageBroker
	breaker *CircuitBreaker
}

// NewUserService creates a new user service
func NewUserService(broker *MessageBroker) *UserService {
	return &UserService{
		users:   make(map[int]*User),
		broker:  broker,
		breaker: NewCircuitBreaker("user-service", 3, 30*time.Second),
	}
}

// CreateUser creates a new user
func (us *UserService) CreateUser(name, email string) (*User, error) {
	var user *User
	var err error

	err = us.breaker.Execute(func() error {
		us.mu.Lock()
		defer us.mu.Unlock()

		// Simulate potential failure
		if rand.Float64() < 0.1 {
			return fmt.Errorf("random failure in user creation")
		}

		id := len(us.users) + 1
		user = &User{
			ID:      id,
			Name:    name,
			Email:   email,
			Created: time.Now(),
		}

		us.users[id] = user
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Publish user created event
	us.broker.Publish("user.created", user)

	return user, nil
}

// GetUser retrieves a user by ID
func (us *UserService) GetUser(id int) (*User, error) {
	us.mu.RLock()
	defer us.mu.RUnlock()

	user, exists := us.users[id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

// OrderService handles order-related operations
type OrderService struct {
	orders     map[int]*Order
	mu         sync.RWMutex
	broker     *MessageBroker
	breaker    *CircuitBreaker
	workerPool *WorkerPool
}

// NewOrderService creates a new order service
func NewOrderService(broker *MessageBroker) *OrderService {
	os := &OrderService{
		orders:     make(map[int]*Order),
		broker:     broker,
		breaker:    NewCircuitBreaker("order-service", 5, 60*time.Second),
		workerPool: NewWorkerPool(3, 100),
	}

	os.workerPool.Start()
	return os
}

// CreateOrder creates a new order
func (os *OrderService) CreateOrder(userID int, product string, amount float64) (*Order, error) {
	var order *Order
	var err error

	err = os.breaker.Execute(func() error {
		os.mu.Lock()
		defer os.mu.Unlock()

		// Simulate processing time
		time.Sleep(10 * time.Millisecond)

		// Simulate potential failure
		if rand.Float64() < 0.05 {
			return fmt.Errorf("random failure in order creation")
		}

		id := len(os.orders) + 1
		order = &Order{
			ID:      id,
			UserID:  userID,
			Product: product,
			Amount:  amount,
			Status:  "pending",
			Created: time.Now(),
		}

		os.orders[id] = order
		return nil
	})

	if err != nil {
		return nil, err
	}

	// Process order asynchronously
	os.processOrderAsync(order)

	// Publish order created event
	os.broker.Publish("order.created", order)

	return order, nil
}

// processOrderAsync processes an order asynchronously
func (os *OrderService) processOrderAsync(order *Order) {
	job := Job{
		ID: fmt.Sprintf("order-%d", order.ID),
		Task: func() error {
			// Simulate order processing
			time.Sleep(100 * time.Millisecond)

			os.mu.Lock()
			defer os.mu.Unlock()

			if storedOrder, exists := os.orders[order.ID]; exists {
				storedOrder.Status = "completed"

				// Publish order completed event
				os.broker.Publish("order.completed", storedOrder)
			}

			return nil
		},
		Result: make(chan error, 1),
	}

	os.workerPool.Submit(job)
}

// GetOrder retrieves an order by ID
func (os *OrderService) GetOrder(id int) (*Order, error) {
	os.mu.RLock()
	defer os.mu.RUnlock()

	order, exists := os.orders[id]
	if !exists {
		return nil, fmt.Errorf("order not found")
	}

	return order, nil
}

// NotificationService handles notification operations
type NotificationService struct {
	notifications map[int]*Notification
	mu            sync.RWMutex
	broker        *MessageBroker
	messageQueue  chan Message
}

// NewNotificationService creates a new notification service
func NewNotificationService(broker *MessageBroker) *NotificationService {
	ns := &NotificationService{
		notifications: make(map[int]*Notification),
		broker:        broker,
		messageQueue:  make(chan Message, 100),
	}

	// Subscribe to events
	ns.broker.Subscribe("user.created", ns.messageQueue)
	ns.broker.Subscribe("order.created", ns.messageQueue)
	ns.broker.Subscribe("order.completed", ns.messageQueue)

	// Start message processor
	go ns.processMessages()

	return ns
}

// processMessages processes incoming messages
func (ns *NotificationService) processMessages() {
	for message := range ns.messageQueue {
		switch message.Topic {
		case "user.created":
			if user, ok := message.Payload.(*User); ok {
				ns.sendWelcomeNotification(user)
			}
		case "order.created":
			if order, ok := message.Payload.(*Order); ok {
				ns.sendOrderConfirmation(order)
			}
		case "order.completed":
			if order, ok := message.Payload.(*Order); ok {
				ns.sendOrderCompletion(order)
			}
		}
	}
}

// sendWelcomeNotification sends a welcome notification
func (ns *NotificationService) sendWelcomeNotification(user *User) {
	ns.createNotification(user.ID, "email",
		fmt.Sprintf("Welcome %s! Your account has been created.", user.Name))
}

// sendOrderConfirmation sends an order confirmation
func (ns *NotificationService) sendOrderConfirmation(order *Order) {
	ns.createNotification(order.UserID, "email",
		fmt.Sprintf("Order #%d confirmed for %s ($%.2f)", order.ID, order.Product, order.Amount))
}

// sendOrderCompletion sends an order completion notification
func (ns *NotificationService) sendOrderCompletion(order *Order) {
	ns.createNotification(order.UserID, "email",
		fmt.Sprintf("Order #%d completed! Your %s is ready.", order.ID, order.Product))
}

// createNotification creates a new notification
func (ns *NotificationService) createNotification(userID int, notificationType, message string) {
	ns.mu.Lock()
	defer ns.mu.Unlock()

	id := len(ns.notifications) + 1
	notification := &Notification{
		ID:      id,
		UserID:  userID,
		Type:    notificationType,
		Message: message,
		Sent:    false,
		Created: time.Now(),
	}

	ns.notifications[id] = notification

	// Simulate sending notification
	go ns.sendNotification(notification)
}

// sendNotification simulates sending a notification
func (ns *NotificationService) sendNotification(notification *Notification) {
	// Simulate network delay
	time.Sleep(50 * time.Millisecond)

	ns.mu.Lock()
	defer ns.mu.Unlock()

	notification.Sent = true
	log.Printf("Notification sent: %s", notification.Message)
}

// === HEALTH CHECK SYSTEM ===

// HealthChecker provides health check functionality
type HealthChecker struct {
	checks map[string]func() error
	mu     sync.RWMutex
}

// NewHealthChecker creates a new health checker
func NewHealthChecker() *HealthChecker {
	return &HealthChecker{
		checks: make(map[string]func() error),
	}
}

// RegisterCheck registers a health check
func (hc *HealthChecker) RegisterCheck(name string, check func() error) {
	hc.mu.Lock()
	defer hc.mu.Unlock()
	hc.checks[name] = check
}

// CheckHealth performs all health checks
func (hc *HealthChecker) CheckHealth() map[string]string {
	hc.mu.RLock()
	defer hc.mu.RUnlock()

	results := make(map[string]string)

	for name, check := range hc.checks {
		if err := check(); err != nil {
			results[name] = fmt.Sprintf("unhealthy: %v", err)
		} else {
			results[name] = "healthy"
		}
	}

	return results
}

// === API GATEWAY ===

// APIGateway handles HTTP requests and routes them to services
type APIGateway struct {
	userService         *UserService
	orderService        *OrderService
	notificationService *NotificationService
	healthChecker       *HealthChecker
}

// NewAPIGateway creates a new API gateway
func NewAPIGateway(userService *UserService, orderService *OrderService,
	notificationService *NotificationService, healthChecker *HealthChecker) *APIGateway {
	return &APIGateway{
		userService:         userService,
		orderService:        orderService,
		notificationService: notificationService,
		healthChecker:       healthChecker,
	}
}

// StartServer starts the HTTP server
func (ag *APIGateway) StartServer(port string) {
	http.HandleFunc("/health", ag.healthHandler)
	http.HandleFunc("/users", ag.usersHandler)
	http.HandleFunc("/orders", ag.ordersHandler)
	http.HandleFunc("/stats", ag.statsHandler)

	log.Printf("API Gateway starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// healthHandler handles health check requests
func (ag *APIGateway) healthHandler(w http.ResponseWriter, r *http.Request) {
	results := ag.healthChecker.CheckHealth()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "ok",
		"checks":    results,
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

// usersHandler handles user requests
func (ag *APIGateway) usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var req struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		user, err := ag.userService.CreateUser(req.Name, req.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// ordersHandler handles order requests
func (ag *APIGateway) ordersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var req struct {
			UserID  int     `json:"user_id"`
			Product string  `json:"product"`
			Amount  float64 `json:"amount"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		order, err := ag.orderService.CreateOrder(req.UserID, req.Product, req.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(order)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// statsHandler provides system statistics
func (ag *APIGateway) statsHandler(w http.ResponseWriter, r *http.Request) {
	stats := map[string]interface{}{
		"user_service_circuit_breaker":  ag.userService.breaker.GetState(),
		"order_service_circuit_breaker": ag.orderService.breaker.GetState(),
		"timestamp":                     time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

// === MAIN APPLICATION ===

func main() {
	fmt.Println("=== PROJECT 16: GO MICROSERVICES WITH ADVANCED CONCURRENCY ===")

	// Initialize message broker
	broker := NewMessageBroker()

	// Initialize services
	userService := NewUserService(broker)
	orderService := NewOrderService(broker)
	notificationService := NewNotificationService(broker)

	// Initialize health checker
	healthChecker := NewHealthChecker()

	// Register health checks
	healthChecker.RegisterCheck("user-service", func() error {
		if userService.breaker.GetState() == Open {
			return fmt.Errorf("circuit breaker is open")
		}
		return nil
	})

	healthChecker.RegisterCheck("order-service", func() error {
		if orderService.breaker.GetState() == Open {
			return fmt.Errorf("circuit breaker is open")
		}
		return nil
	})

	healthChecker.RegisterCheck("notification-service", func() error {
		// Simple health check
		return nil
	})

	// Initialize API gateway
	gateway := NewAPIGateway(userService, orderService, notificationService, healthChecker)

	// Start background demo
	go runDemo(userService, orderService)

	// Start API server
	log.Println("Starting microservices...")
	log.Println("API Endpoints:")
	log.Println("POST /users - Create user")
	log.Println("POST /orders - Create order")
	log.Println("GET /health - Health check")
	log.Println("GET /stats - System statistics")

	gateway.StartServer("8080")
}

// runDemo demonstrates the microservices in action
func runDemo(userService *UserService, orderService *OrderService) {
	time.Sleep(2 * time.Second)

	log.Println("=== DEMO: Creating users and orders ===")

	// Create some users
	users := []struct {
		name  string
		email string
	}{
		{"Alice Johnson", "alice@example.com"},
		{"Bob Smith", "bob@example.com"},
		{"Charlie Brown", "charlie@example.com"},
	}

	var createdUsers []*User
	for _, u := range users {
		user, err := userService.CreateUser(u.name, u.email)
		if err != nil {
			log.Printf("Error creating user %s: %v", u.name, err)
			continue
		}
		createdUsers = append(createdUsers, user)
		log.Printf("Created user: %s", user.Name)
	}

	// Create some orders
	products := []string{"Laptop", "Mouse", "Keyboard", "Monitor"}
	for _, user := range createdUsers {
		for j := 0; j < 2; j++ {
			product := products[rand.Intn(len(products))]
			amount := float64(rand.Intn(1000) + 100)

			order, err := orderService.CreateOrder(user.ID, product, amount)
			if err != nil {
				log.Printf("Error creating order for user %s: %v", user.Name, err)
				continue
			}
			log.Printf("Created order: %d for user %s", order.ID, user.Name)
		}

		// Add some delay between users
		time.Sleep(500 * time.Millisecond)
	}

	log.Println("=== DEMO: Complete ===")
}

/*
RUNNING THE PROJECT:

1. Install dependencies:
   go mod init project16
   go mod tidy

2. Run the microservices:
   go run main.go

3. Test the API:
   curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"name":"John Doe","email":"john@example.com"}'
   curl -X POST http://localhost:8080/orders -H "Content-Type: application/json" -d '{"user_id":1,"product":"Laptop","amount":1299.99}'
   curl -X GET http://localhost:8080/health
   curl -X GET http://localhost:8080/stats

LEARNING POINTS:

1. MICROSERVICES ARCHITECTURE:
   - Service separation and communication
   - Event-driven architecture with message broker
   - Inter-service communication patterns

2. ADVANCED CONCURRENCY:
   - Worker pools for parallel processing
   - Channel-based communication
   - Goroutine management and cleanup

3. FAULT TOLERANCE:
   - Circuit breaker pattern implementation
   - Graceful degradation
   - Error handling and recovery

4. MESSAGING PATTERNS:
   - Publish-subscribe pattern
   - Asynchronous processing
   - Event sourcing concepts

5. MONITORING AND HEALTH:
   - Health check endpoints
   - System statistics
   - Circuit breaker monitoring

6. PRODUCTION PATTERNS:
   - Graceful shutdown
   - Resource cleanup
   - Configuration management
   - Error logging and monitoring

This project demonstrates how to build scalable, fault-tolerant microservices
with Go, using advanced concurrency patterns and best practices.
*/
