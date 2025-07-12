package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// === CONCURRENCY IN GO ===

// Worker function
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

// Producer function
func producer(jobs chan<- int, numJobs int) {
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)
}

// Consumer function
func consumer(results <-chan int, numJobs int) {
	for i := 1; i <= numJobs; i++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
}

// Ping-pong example
func pingPong(pings chan<- string, pongs <-chan string) {
	msg := <-pongs
	fmt.Printf("Received: %s\n", msg)
	pings <- "pong"
}

// Rate limiter example
func rateLimiter(requests chan string, limiter <-chan time.Time) {
	for req := range requests {
		<-limiter
		fmt.Printf("Processing request: %s at %v\n", req, time.Now())
	}
}

// Fan-out fan-in pattern
func fanOut(input <-chan int, workers int) []<-chan int {
	outputs := make([]<-chan int, workers)
	for i := 0; i < workers; i++ {
		output := make(chan int)
		outputs[i] = output
		go func(out chan<- int) {
			defer close(out)
			for n := range input {
				out <- n * n
			}
		}(output)
	}
	return outputs
}

func fanIn(inputs ...<-chan int) <-chan int {
	output := make(chan int)
	var wg sync.WaitGroup

	for _, input := range inputs {
		wg.Add(1)
		go func(in <-chan int) {
			defer wg.Done()
			for n := range in {
				output <- n
			}
		}(input)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

// Pipeline example
func pipeline() {
	// Stage 1: Generate numbers
	numbers := make(chan int)
	go func() {
		defer close(numbers)
		for i := 1; i <= 10; i++ {
			numbers <- i
		}
	}()

	// Stage 2: Square numbers
	squares := make(chan int)
	go func() {
		defer close(squares)
		for n := range numbers {
			squares <- n * n
		}
	}()

	// Stage 3: Sum squares
	sum := 0
	for s := range squares {
		sum += s
	}

	fmt.Printf("Sum of squares: %d\n", sum)
}

// Context example
func contextExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	done := make(chan bool)

	go func() {
		time.Sleep(3 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Operation completed")
	case <-ctx.Done():
		fmt.Println("Operation timed out:", ctx.Err())
	}
}

// Mutex example
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// RWMutex example
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	return val, ok
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Once example
var once sync.Once
var config map[string]string

func loadConfig() {
	once.Do(func() {
		fmt.Println("Loading configuration...")
		config = map[string]string{
			"database_url": "localhost:5432",
			"api_key":      "secret123",
		}
	})
}

// WaitGroup example
func waitGroupExample() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d starting\n", id)
			time.Sleep(time.Second)
			fmt.Printf("Goroutine %d finished\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines completed")
}

// Pool example
var pool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1024)
	},
}

func poolExample() {
	// Get from pool
	buf := pool.Get().([]byte)
	defer pool.Put(buf)

	// Use buffer
	copy(buf, "Hello, World!")
	fmt.Printf("Buffer content: %s\n", string(buf[:13]))
}

// Atomic operations example
func atomicExample() {
	var counter int64
	var wg sync.WaitGroup

	// This would be unsafe without synchronization
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// atomic.AddInt64(&counter, 1) // Safe atomic operation
			counter++ // Unsafe operation for demonstration
		}()
	}

	wg.Wait()
	fmt.Printf("Counter value: %d\n", counter)
}

// Error handling with goroutines
func errorHandlingExample() {
	type result struct {
		value int
		err   error
	}

	results := make(chan result, 3)

	for i := 1; i <= 3; i++ {
		go func(n int) {
			defer func() {
				if r := recover(); r != nil {
					results <- result{err: fmt.Errorf("panic: %v", r)}
				}
			}()

			if n == 2 {
				panic("simulated error")
			}

			results <- result{value: n * n}
		}(i)
	}

	for i := 0; i < 3; i++ {
		r := <-results
		if r.err != nil {
			fmt.Printf("Error: %v\n", r.err)
		} else {
			fmt.Printf("Result: %d\n", r.value)
		}
	}
}

func main() {
	fmt.Println("=== GO CONCURRENCY COMPREHENSIVE GUIDE ===")

	// === GOROUTINES BASICS ===
	fmt.Println("\n--- GOROUTINES BASICS ---")

	// Basic goroutine
	go func() {
		fmt.Println("This is a goroutine")
	}()

	// Multiple goroutines
	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d is running\n", id)
		}(i)
	}

	// Give goroutines time to execute
	time.Sleep(100 * time.Millisecond)

	/*
		JavaScript comparison:
		// JavaScript uses async/await and Promises
		async function asyncFunction() {
			console.log("This is async");
		}

		// Multiple async operations
		for (let i = 1; i <= 3; i++) {
			setTimeout(() => {
				console.log(`Async operation ${i}`);
			}, 0);
		}

		// Promise-based concurrency
		const promises = [];
		for (let i = 1; i <= 3; i++) {
			promises.push(new Promise(resolve => {
				setTimeout(() => {
					console.log(`Promise ${i}`);
					resolve(i);
				}, 100);
			}));
		}

		Promise.all(promises).then(results => {
			console.log("All promises completed:", results);
		});
	*/

	// === CHANNELS BASICS ===
	fmt.Println("\n--- CHANNELS BASICS ---")

	// Unbuffered channel
	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine"
	}()

	msg := <-ch
	fmt.Println("Received:", msg)

	// Buffered channel
	buffered := make(chan int, 2)
	buffered <- 1
	buffered <- 2

	fmt.Println("Buffered channel value 1:", <-buffered)
	fmt.Println("Buffered channel value 2:", <-buffered)

	// === CHANNEL DIRECTIONS ===
	fmt.Println("\n--- CHANNEL DIRECTIONS ---")

	// Send-only channel
	sendOnly := make(chan<- int)
	go func(ch chan<- int) {
		ch <- 42
	}(sendOnly)

	// Receive-only channel
	receiveOnly := make(<-chan int, 1)
	go func() {
		val := <-receiveOnly
		fmt.Println("Received from receive-only:", val)
	}()

	// Wait for goroutines
	time.Sleep(100 * time.Millisecond)

	// === CHANNEL CLOSING ===
	fmt.Println("\n--- CHANNEL CLOSING ---")

	jobs := make(chan int, 3)
	jobs <- 1
	jobs <- 2
	jobs <- 3
	close(jobs)

	// Receive from closed channel
	for job := range jobs {
		fmt.Printf("Job: %d\n", job)
	}

	// Check if channel is closed
	value, ok := <-jobs
	fmt.Printf("Channel closed: value=%d, ok=%t\n", value, ok)

	// === SELECT STATEMENT ===
	fmt.Println("\n--- SELECT STATEMENT ---")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Channel 1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Channel 2"
	}()

	// Select from multiple channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from ch2:", msg2)
		}
	}

	// === SELECT WITH TIMEOUT ===
	fmt.Println("\n--- SELECT WITH TIMEOUT ---")

	slow := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		slow <- "slow result"
	}()

	select {
	case result := <-slow:
		fmt.Println("Got result:", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Operation timed out")
	}

	// === SELECT WITH DEFAULT ===
	fmt.Println("\n--- SELECT WITH DEFAULT ---")

	nonBlocking := make(chan string)

	select {
	case msg := <-nonBlocking:
		fmt.Println("Got message:", msg)
	default:
		fmt.Println("No message available")
	}

	// === WORKER POOL PATTERN ===
	fmt.Println("\n--- WORKER POOL PATTERN ---")

	const numJobs = 5
	const numWorkers = 3

	jobsChan := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobsChan, results)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobsChan <- j
	}
	close(jobsChan)

	// Collect results
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("Job result: %d\n", result)
	}

	// === PRODUCER-CONSUMER PATTERN ===
	fmt.Println("\n--- PRODUCER-CONSUMER PATTERN ---")

	producerJobs := make(chan int, 3)
	consumerResults := make(chan int, 3)

	go producer(producerJobs, 3)
	go consumer(consumerResults, 3)

	// Process jobs
	for job := range producerJobs {
		consumerResults <- job * 10
	}

	// === PING-PONG PATTERN ===
	fmt.Println("\n--- PING-PONG PATTERN ---")

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	pings <- "ping"

	go func() {
		msg := <-pings
		fmt.Printf("Received: %s\n", msg)
		pongs <- "pong"
	}()

	go pingPong(pings, pongs)

	time.Sleep(100 * time.Millisecond)

	// === RATE LIMITING ===
	fmt.Println("\n--- RATE LIMITING ---")

	requests := make(chan string, 5)
	limiter := time.Tick(200 * time.Millisecond)

	// Send requests
	for i := 1; i <= 5; i++ {
		requests <- fmt.Sprintf("request-%d", i)
	}
	close(requests)

	go rateLimiter(requests, limiter)

	time.Sleep(2 * time.Second)

	// === PIPELINE PATTERN ===
	fmt.Println("\n--- PIPELINE PATTERN ---")

	pipeline()

	// === FAN-OUT FAN-IN PATTERN ===
	fmt.Println("\n--- FAN-OUT FAN-IN PATTERN ---")

	input := make(chan int, 10)

	// Send input
	go func() {
		defer close(input)
		for i := 1; i <= 5; i++ {
			input <- i
		}
	}()

	// Fan-out to multiple workers
	outputs := fanOut(input, 3)

	// Fan-in results
	result := fanIn(outputs...)

	// Collect results
	for r := range result {
		fmt.Printf("Fan-in result: %d\n", r)
	}

	// === CONTEXT USAGE ===
	fmt.Println("\n--- CONTEXT USAGE ---")

	contextExample()

	// Context with cancellation
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context cancelled, stopping work")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(500 * time.Millisecond)
	cancel()
	time.Sleep(100 * time.Millisecond)

	// === MUTEX SYNCHRONIZATION ===
	fmt.Println("\n--- MUTEX SYNCHRONIZATION ---")

	counter := &Counter{}
	var wg sync.WaitGroup

	// Start multiple goroutines
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Counter final value: %d\n", counter.Value())

	// === RWMUTEX EXAMPLE ===
	fmt.Println("\n--- RWMUTEX EXAMPLE ---")

	cache := NewCache()

	// Multiple readers
	for i := 0; i < 3; i++ {
		go func(id int) {
			cache.Set(fmt.Sprintf("key-%d", id), fmt.Sprintf("value-%d", id))
			fmt.Printf("Writer %d finished\n", id)
		}(i)
	}

	// Multiple writers
	for i := 0; i < 5; i++ {
		go func(id int) {
			if val, ok := cache.Get(fmt.Sprintf("key-%d", id%3)); ok {
				fmt.Printf("Reader %d got: %s\n", id, val)
			}
		}(i)
	}

	time.Sleep(100 * time.Millisecond)

	// === SYNC.ONCE EXAMPLE ===
	fmt.Println("\n--- SYNC.ONCE EXAMPLE ---")

	// Multiple goroutines trying to load config
	for i := 0; i < 3; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d calling loadConfig\n", id)
			loadConfig()
			fmt.Printf("Goroutine %d: config loaded\n", id)
		}(i)
	}

	time.Sleep(100 * time.Millisecond)

	// === WAITGROUP EXAMPLE ===
	fmt.Println("\n--- WAITGROUP EXAMPLE ---")

	waitGroupExample()

	// === SYNC.POOL EXAMPLE ===
	fmt.Println("\n--- SYNC.POOL EXAMPLE ---")

	poolExample()

	// === ATOMIC OPERATIONS ===
	fmt.Println("\n--- ATOMIC OPERATIONS ---")

	atomicExample()

	// === ERROR HANDLING WITH GOROUTINES ===
	fmt.Println("\n--- ERROR HANDLING WITH GOROUTINES ---")

	errorHandlingExample()

	// === COMMON PATTERNS ===
	fmt.Println("\n--- COMMON PATTERNS ---")

	// 1. Timeout pattern
	timeout := time.After(1 * time.Second)
	work := make(chan bool)

	go func() {
		time.Sleep(500 * time.Millisecond)
		work <- true
	}()

	select {
	case <-work:
		fmt.Println("Work completed within timeout")
	case <-timeout:
		fmt.Println("Work timed out")
	}

	// 2. Quit channel pattern
	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Received quit signal")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	time.Sleep(300 * time.Millisecond)
	quit <- true
	time.Sleep(100 * time.Millisecond)

	// 3. Semaphore pattern
	semaphore := make(chan bool, 2) // Allow 2 concurrent operations

	for i := 1; i <= 5; i++ {
		go func(id int) {
			semaphore <- true              // Acquire
			defer func() { <-semaphore }() // Release

			fmt.Printf("Worker %d acquired semaphore\n", id)
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("Worker %d released semaphore\n", id)
		}(i)
	}

	time.Sleep(2 * time.Second)

	// === RACE CONDITIONS ===
	fmt.Println("\n--- RACE CONDITIONS ---")

	// Demonstrate race condition (unsafe)
	var unsafeCounter int
	var wg2 sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			unsafeCounter++
		}()
	}

	wg2.Wait()
	fmt.Printf("Unsafe counter (race condition): %d\n", unsafeCounter)

	// === DEADLOCK PREVENTION ===
	fmt.Println("\n--- DEADLOCK PREVENTION ---")

	// Example of potential deadlock (commented out)
	/*
		ch1 := make(chan bool)
		ch2 := make(chan bool)

		go func() {
			ch1 <- true
			<-ch2
		}()

		go func() {
			ch2 <- true
			<-ch1
		}()
	*/

	fmt.Println("Deadlock example commented out to prevent program hanging")

	// === BEST PRACTICES ===
	fmt.Println("\n--- CONCURRENCY BEST PRACTICES ---")
	fmt.Println("1. Don't communicate by sharing memory; share memory by communicating")
	fmt.Println("2. Use channels for communication between goroutines")
	fmt.Println("3. Use mutexes for protecting shared state")
	fmt.Println("4. Always close channels when done sending")
	fmt.Println("5. Use select for non-blocking operations")
	fmt.Println("6. Use context for cancellation and timeouts")
	fmt.Println("7. Use sync.WaitGroup to wait for goroutines")
	fmt.Println("8. Avoid goroutine leaks by ensuring they can exit")
	fmt.Println("9. Use buffered channels to prevent blocking")
	fmt.Println("10. Test concurrent code thoroughly")

	// === COMMON PITFALLS ===
	fmt.Println("\n--- COMMON PITFALLS TO AVOID ---")
	fmt.Println("❌ Forgetting to close channels")
	fmt.Println("❌ Not handling channel blocking")
	fmt.Println("❌ Race conditions on shared variables")
	fmt.Println("❌ Goroutine leaks")
	fmt.Println("❌ Deadlocks from circular dependencies")
	fmt.Println("❌ Not using proper synchronization")
	fmt.Println("❌ Sharing mutable state without protection")
	fmt.Println("❌ Not using context for cancellation")

	fmt.Println("\nGo's concurrency primitives make concurrent programming safe and efficient!")
}
