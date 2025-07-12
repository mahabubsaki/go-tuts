package main

import (
	"fmt"
	"math/rand"
	"time"
)

// === GO CHANNELS OF CHANNELS COMPREHENSIVE GUIDE ===

/*
CHANNELS OF CHANNELS PHILOSOPHY:
- Channels can contain channels as values
- Enables dynamic communication patterns
- Allows for runtime routing of messages
- Facilitates complex orchestration patterns
- Enables pipeline reconfiguration

COMPARISON WITH JAVASCRIPT:
// JavaScript - No direct equivalent, but similar concepts:
// Callback functions (functions as values)
const callbacks = [
  function(data) { console.log("Process A:", data); },
  function(data) { console.log("Process B:", data); },
];

// Promise chains
const chains = [
  Promise.resolve().then(data => processA(data)),
  Promise.resolve().then(data => processB(data)),
];

// Go - channels of channels
chan chan int  // Channel that carries channels of int
*/

// === WORKER TYPES ===

type WorkerID int

type Job struct {
	ID       int
	Data     string
	Response chan string
}

type Worker struct {
	ID       WorkerID
	JobQueue chan Job
	Quit     chan bool
}

func (w *Worker) Start() {
	go func() {
		for {
			select {
			case job := <-w.JobQueue:
				// Process job
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				result := fmt.Sprintf("Worker %d processed job %d: %s", w.ID, job.ID, job.Data)
				job.Response <- result
			case <-w.Quit:
				fmt.Printf("Worker %d stopping\n", w.ID)
				return
			}
		}
	}()
}

// === DISPATCHER STRUCT ===

type Dispatcher struct {
	Workers    []*Worker
	JobQueue   chan Job
	WorkerPool chan chan Job
}

func NewDispatcher(numWorkers int) *Dispatcher {
	return &Dispatcher{
		Workers:    make([]*Worker, numWorkers),
		JobQueue:   make(chan Job, 100),
		WorkerPool: make(chan chan Job, numWorkers),
	}
}

func (d *Dispatcher) Start() {
	// Start workers
	for i := range d.Workers {
		worker := &Worker{
			ID:       WorkerID(i),
			JobQueue: make(chan Job),
			Quit:     make(chan bool),
		}
		d.Workers[i] = worker
		worker.Start()

		// Add worker to pool
		go func(w *Worker) {
			for {
				select {
				case d.WorkerPool <- w.JobQueue:
					// Worker available
				case <-w.Quit:
					return
				}
			}
		}(worker)
	}

	// Start dispatcher
	go func() {
		for {
			select {
			case job := <-d.JobQueue:
				// Get available worker
				go func() {
					workerJobQueue := <-d.WorkerPool
					workerJobQueue <- job
				}()
			}
		}
	}()
}

func (d *Dispatcher) Stop() {
	for _, worker := range d.Workers {
		worker.Quit <- true
	}
}

// === ROUTER STRUCT ===

type Router struct {
	Routes map[string]chan string
}

func NewRouter() *Router {
	return &Router{
		Routes: make(map[string]chan string),
	}
}

func (r *Router) AddRoute(name string, handler chan string) {
	r.Routes[name] = handler
}

func (r *Router) Route(route string, message string) {
	if ch, exists := r.Routes[route]; exists {
		go func() {
			ch <- message
		}()
	}
}

// === MAIN FUNCTION ===

func main() {
	fmt.Println("=== GO CHANNELS OF CHANNELS COMPREHENSIVE GUIDE ===")

	// === BASIC CHANNEL OF CHANNELS ===
	fmt.Println("\n1. BASIC CHANNEL OF CHANNELS:")

	// Create a channel that carries channels of int
	chOfCh := make(chan chan int, 3)

	// Create some channels
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	// Send channels through the channel
	chOfCh <- ch1
	chOfCh <- ch2
	chOfCh <- ch3

	// Send data through the individual channels
	ch1 <- 100
	ch2 <- 200
	ch3 <- 300

	// Receive and use the channels
	for i := 0; i < 3; i++ {
		receivedCh := <-chOfCh
		data := <-receivedCh
		fmt.Printf("Received %d from channel %d\n", data, i+1)
	}

	// === DYNAMIC ROUTING PATTERN ===
	fmt.Println("\n2. DYNAMIC ROUTING PATTERN:")

	// Create router
	router := NewRouter()

	// Create handler channels
	logHandler := make(chan string, 10)
	emailHandler := make(chan string, 10)
	smsHandler := make(chan string, 10)

	// Add routes
	router.AddRoute("log", logHandler)
	router.AddRoute("email", emailHandler)
	router.AddRoute("sms", smsHandler)

	// Start handlers
	go func() {
		for msg := range logHandler {
			fmt.Printf("LOG: %s\n", msg)
		}
	}()

	go func() {
		for msg := range emailHandler {
			fmt.Printf("EMAIL: %s\n", msg)
		}
	}()

	go func() {
		for msg := range smsHandler {
			fmt.Printf("SMS: %s\n", msg)
		}
	}()

	// Send messages to different routes
	router.Route("log", "Application started")
	router.Route("email", "Welcome to our service")
	router.Route("sms", "Your code is 123456")
	router.Route("log", "User logged in")

	time.Sleep(100 * time.Millisecond)

	// === WORKER POOL WITH CHANNEL OF CHANNELS ===
	fmt.Println("\n3. WORKER POOL WITH CHANNEL OF CHANNELS:")

	// Create dispatcher
	dispatcher := NewDispatcher(3)
	dispatcher.Start()

	// Send jobs
	for i := 1; i <= 10; i++ {
		response := make(chan string, 1)
		job := Job{
			ID:       i,
			Data:     fmt.Sprintf("Task %d", i),
			Response: response,
		}

		dispatcher.JobQueue <- job

		// Wait for response
		go func(jobID int, resp chan string) {
			result := <-resp
			fmt.Printf("Job %d result: %s\n", jobID, result)
		}(i, response)
	}

	// Wait for processing
	time.Sleep(3 * time.Second)

	// Stop dispatcher
	dispatcher.Stop()

	// === PIPELINE RECONFIGURATION ===
	fmt.Println("\n4. PIPELINE RECONFIGURATION:")

	// Create pipeline stages
	stage1 := make(chan int, 5)
	stage2 := make(chan int, 5)
	stage3 := make(chan int, 5)

	// Pipeline configuration channel
	pipelineConfig := make(chan chan int, 3)

	// Configure pipeline
	pipelineConfig <- stage1
	pipelineConfig <- stage2
	pipelineConfig <- stage3

	// Start pipeline stages
	go func() {
		for data := range stage1 {
			result := data * 2
			fmt.Printf("Stage 1: %d -> %d\n", data, result)
			stage2 <- result
		}
	}()

	go func() {
		for data := range stage2 {
			result := data + 10
			fmt.Printf("Stage 2: %d -> %d\n", data, result)
			stage3 <- result
		}
	}()

	go func() {
		for data := range stage3 {
			result := data * 3
			fmt.Printf("Stage 3: %d -> %d (final)\n", data, result)
		}
	}()

	// Send data through pipeline
	for i := 1; i <= 5; i++ {
		stage1 <- i
	}

	time.Sleep(500 * time.Millisecond)

	// === REQUEST-RESPONSE PATTERN ===
	fmt.Println("\n5. REQUEST-RESPONSE PATTERN:")

	// Server with channel of response channels
	type Request struct {
		ID       int
		Data     string
		Response chan string
	}

	requests := make(chan Request, 10)

	// Server
	go func() {
		for req := range requests {
			// Process request
			time.Sleep(100 * time.Millisecond)
			response := fmt.Sprintf("Processed request %d: %s", req.ID, req.Data)
			req.Response <- response
		}
	}()

	// Client requests
	for i := 1; i <= 5; i++ {
		response := make(chan string, 1)
		request := Request{
			ID:       i,
			Data:     fmt.Sprintf("Request %d data", i),
			Response: response,
		}

		requests <- request

		// Wait for response
		go func(reqID int, resp chan string) {
			result := <-resp
			fmt.Printf("Request %d response: %s\n", reqID, result)
		}(i, response)
	}

	time.Sleep(1 * time.Second)

	// === LOAD BALANCER PATTERN ===
	fmt.Println("\n6. LOAD BALANCER PATTERN:")

	// Create backend servers
	backends := make([]chan string, 3)
	for i := range backends {
		backends[i] = make(chan string, 10)

		// Start backend server
		go func(id int, ch chan string) {
			for msg := range ch {
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
				fmt.Printf("Backend %d processed: %s\n", id, msg)
			}
		}(i, backends[i])
	}

	// Load balancer
	loadBalancer := make(chan chan string, len(backends))

	// Add backends to load balancer
	for _, backend := range backends {
		loadBalancer <- backend
	}

	// Send requests to load balancer
	for i := 1; i <= 10; i++ {
		// Get next backend
		backend := <-loadBalancer

		// Send request
		go func(b chan string, reqID int) {
			b <- fmt.Sprintf("Request %d", reqID)
			// Return backend to pool
			loadBalancer <- b
		}(backend, i)
	}

	time.Sleep(2 * time.Second)

	// === MULTIPLEXER PATTERN ===
	fmt.Println("\n7. MULTIPLEXER PATTERN:")

	// Input channels
	input1 := make(chan string, 5)
	input2 := make(chan string, 5)
	input3 := make(chan string, 5)

	// Multiplexer
	multiplexer := make(chan chan string, 3)
	output := make(chan string, 15)

	// Add inputs to multiplexer
	multiplexer <- input1
	multiplexer <- input2
	multiplexer <- input3

	// Start multiplexer
	go func() {
		for {
			select {
			case ch := <-multiplexer:
				go func(inputCh chan string) {
					for msg := range inputCh {
						output <- msg
					}
				}(ch)
			}
		}
	}()

	// Output handler
	go func() {
		for msg := range output {
			fmt.Printf("Multiplexed output: %s\n", msg)
		}
	}()

	// Send data to inputs
	input1 <- "Message from input 1"
	input2 <- "Message from input 2"
	input3 <- "Message from input 3"

	time.Sleep(200 * time.Millisecond)

	// === CHANNEL REGISTRY PATTERN ===
	fmt.Println("\n8. CHANNEL REGISTRY PATTERN:")

	// Channel registry
	registry := make(map[string]chan string)
	registryMutex := make(chan bool, 1)
	registryMutex <- true

	// Register channels
	register := func(name string, ch chan string) {
		<-registryMutex
		registry[name] = ch
		registryMutex <- true
	}

	// Lookup channels
	lookup := func(name string) (chan string, bool) {
		<-registryMutex
		ch, exists := registry[name]
		registryMutex <- true
		return ch, exists
	}

	// Create and register channels
	serviceA := make(chan string, 5)
	serviceB := make(chan string, 5)
	serviceC := make(chan string, 5)

	register("serviceA", serviceA)
	register("serviceB", serviceB)
	register("serviceC", serviceC)

	// Start services
	go func() {
		for msg := range serviceA {
			fmt.Printf("Service A: %s\n", msg)
		}
	}()

	go func() {
		for msg := range serviceB {
			fmt.Printf("Service B: %s\n", msg)
		}
	}()

	go func() {
		for msg := range serviceC {
			fmt.Printf("Service C: %s\n", msg)
		}
	}()

	// Send messages via registry
	if ch, exists := lookup("serviceA"); exists {
		ch <- "Task for Service A"
	}

	if ch, exists := lookup("serviceB"); exists {
		ch <- "Task for Service B"
	}

	if ch, exists := lookup("serviceC"); exists {
		ch <- "Task for Service C"
	}

	time.Sleep(200 * time.Millisecond)

	// === CHARACTERISTICS OF CHANNELS OF CHANNELS ===
	fmt.Println("\n9. CHARACTERISTICS OF CHANNELS OF CHANNELS:")

	fmt.Println("✓ Channels can contain channels as values")
	fmt.Println("✓ Enables dynamic communication patterns")
	fmt.Println("✓ Allows runtime routing of messages")
	fmt.Println("✓ Facilitates complex orchestration")
	fmt.Println("✓ Enables pipeline reconfiguration")
	fmt.Println("✓ Supports request-response patterns")
	fmt.Println("✓ Enables load balancing")
	fmt.Println("✓ Supports multiplexing")

	// === BEST PRACTICES ===
	fmt.Println("\n10. BEST PRACTICES:")

	fmt.Println("✓ Use for dynamic routing scenarios")
	fmt.Println("✓ Implement proper channel lifecycle management")
	fmt.Println("✓ Avoid channel leaks")
	fmt.Println("✓ Use buffered channels appropriately")
	fmt.Println("✓ Handle channel closing gracefully")
	fmt.Println("✓ Consider using context for cancellation")
	fmt.Println("✓ Monitor goroutine count in complex patterns")

	// === COMMON PATTERNS ===
	fmt.Println("\n11. COMMON PATTERNS:")

	fmt.Println("1. Worker Pool - Distribute work across workers")
	fmt.Println("2. Request-Response - Handle client-server communication")
	fmt.Println("3. Load Balancer - Distribute load across backends")
	fmt.Println("4. Multiplexer - Combine multiple inputs")
	fmt.Println("5. Router - Route messages based on criteria")
	fmt.Println("6. Pipeline - Reconfigurable processing stages")
	fmt.Println("7. Registry - Dynamic channel lookup")
	fmt.Println("8. Orchestrator - Complex workflow coordination")

	fmt.Println("\n=== END OF CHANNELS OF CHANNELS GUIDE ===")
}

/*
RUNNING THE PROGRAM:
go run main.go

LEARNING POINTS:

1. CHANNELS OF CHANNELS:
   - Channels can contain channels as values
   - Type: chan chan Type
   - Enables dynamic communication patterns
   - Allows runtime routing

2. PATTERNS:
   - Worker Pool: Distribute work dynamically
   - Request-Response: Handle client-server communication
   - Load Balancer: Distribute load across backends
   - Multiplexer: Combine multiple inputs
   - Router: Route messages based on criteria
   - Pipeline: Reconfigurable processing stages

3. CHARACTERISTICS:
   - Dynamic routing capabilities
   - Runtime reconfiguration
   - Complex orchestration
   - Flexible communication patterns

4. BEST PRACTICES:
   - Use for dynamic scenarios
   - Manage channel lifecycle
   - Avoid channel leaks
   - Handle closing gracefully
   - Monitor goroutine count

5. REAL-WORLD USES:
   - Microservice communication
   - Dynamic load balancing
   - Message routing systems
   - Pipeline orchestration
   - Service discovery

This demonstrates advanced channel patterns in Go
for complex concurrent communication scenarios.
*/
