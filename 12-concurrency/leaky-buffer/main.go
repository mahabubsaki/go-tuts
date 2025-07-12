package main

import (
	"fmt"
	"runtime"
	"time"
)

// === GO LEAKY BUFFER COMPREHENSIVE GUIDE ===

/*
LEAKY BUFFER PHILOSOPHY:
- A leaky buffer is a buffered channel that can handle overflow
- When the buffer is full, new items are dropped (leaked)
- Prevents blocking when the consumer is slower than producer
- Useful for scenarios where dropping data is acceptable
- Common in rate limiting, logging, and event processing

COMPARISON WITH JAVASCRIPT:
// JavaScript - No direct equivalent, but similar concepts:
// Throttling/debouncing
function throttle(func, delay) {
  let timeoutId;
  return function(...args) {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => func.apply(this, args), delay);
  };
}

// Ring buffer simulation
class RingBuffer {
  constructor(size) {
    this.buffer = new Array(size);
    this.head = 0;
    this.tail = 0;
    this.size = size;
  }

  push(item) {
    this.buffer[this.tail] = item;
    this.tail = (this.tail + 1) % this.size;
    if (this.tail === this.head) {
      this.head = (this.head + 1) % this.size; // Overflow, drop oldest
    }
  }
}

// Go - leaky buffer with channels
func leakyBuffer(bufferSize int) chan string {
  buffer := make(chan string, bufferSize)

  go func() {
    for {
      select {
      case item := <-buffer:
        // Process item
        processItem(item)
      default:
        // Buffer is empty, do nothing
      }
    }
  }()

  return buffer
}
*/

// === LEAKY BUFFER IMPLEMENTATION ===

type LeakyBuffer struct {
	buffer   chan interface{}
	dropped  int
	received int
	name     string
}

func NewLeakyBuffer(name string, size int) *LeakyBuffer {
	return &LeakyBuffer{
		buffer: make(chan interface{}, size),
		name:   name,
	}
}

func (lb *LeakyBuffer) Push(item interface{}) bool {
	lb.received++
	select {
	case lb.buffer <- item:
		return true
	default:
		lb.dropped++
		return false
	}
}

func (lb *LeakyBuffer) Pop() (interface{}, bool) {
	select {
	case item := <-lb.buffer:
		return item, true
	default:
		return nil, false
	}
}

func (lb *LeakyBuffer) Stats() (int, int, int) {
	return lb.received, lb.dropped, len(lb.buffer)
}

func (lb *LeakyBuffer) Name() string {
	return lb.name
}

// === ADVANCED LEAKY BUFFER WITH PROCESSING ===

type ProcessingLeakyBuffer struct {
	buffer    chan interface{}
	processor func(interface{})
	dropped   int
	processed int
	received  int
	name      string
	stopChan  chan bool
}

func NewProcessingLeakyBuffer(name string, size int, processor func(interface{})) *ProcessingLeakyBuffer {
	plb := &ProcessingLeakyBuffer{
		buffer:    make(chan interface{}, size),
		processor: processor,
		name:      name,
		stopChan:  make(chan bool),
	}

	go plb.process()
	return plb
}

func (plb *ProcessingLeakyBuffer) Push(item interface{}) bool {
	plb.received++
	select {
	case plb.buffer <- item:
		return true
	default:
		plb.dropped++
		return false
	}
}

func (plb *ProcessingLeakyBuffer) process() {
	for {
		select {
		case item := <-plb.buffer:
			plb.processor(item)
			plb.processed++
		case <-plb.stopChan:
			return
		}
	}
}

func (plb *ProcessingLeakyBuffer) Stop() {
	plb.stopChan <- true
}

func (plb *ProcessingLeakyBuffer) Stats() (int, int, int, int) {
	return plb.received, plb.dropped, plb.processed, len(plb.buffer)
}

// === RATE LIMITER USING LEAKY BUFFER ===

type RateLimiter struct {
	bucket   chan struct{}
	refillCh chan struct{}
	rate     time.Duration
	stopCh   chan bool
}

func NewRateLimiter(capacity int, rate time.Duration) *RateLimiter {
	rl := &RateLimiter{
		bucket:   make(chan struct{}, capacity),
		refillCh: make(chan struct{}),
		rate:     rate,
		stopCh:   make(chan bool),
	}

	// Fill bucket initially
	for i := 0; i < capacity; i++ {
		rl.bucket <- struct{}{}
	}

	// Start refill process
	go rl.refill()

	return rl
}

func (rl *RateLimiter) refill() {
	ticker := time.NewTicker(rl.rate)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			select {
			case rl.bucket <- struct{}{}:
				// Successfully added token
			default:
				// Bucket is full, drop token (leaky behavior)
			}
		case <-rl.stopCh:
			return
		}
	}
}

func (rl *RateLimiter) Allow() bool {
	select {
	case <-rl.bucket:
		return true
	default:
		return false
	}
}

func (rl *RateLimiter) Stop() {
	rl.stopCh <- true
}

// === CIRCULAR BUFFER IMPLEMENTATION ===

type CircularBuffer struct {
	buffer   []interface{}
	head     int
	tail     int
	size     int
	capacity int
	dropped  int
}

func NewCircularBuffer(capacity int) *CircularBuffer {
	return &CircularBuffer{
		buffer:   make([]interface{}, capacity),
		capacity: capacity,
	}
}

func (cb *CircularBuffer) Push(item interface{}) bool {
	if cb.size == cb.capacity {
		// Buffer is full, drop oldest item
		cb.head = (cb.head + 1) % cb.capacity
		cb.dropped++
	} else {
		cb.size++
	}

	cb.buffer[cb.tail] = item
	cb.tail = (cb.tail + 1) % cb.capacity
	return true
}

func (cb *CircularBuffer) Pop() (interface{}, bool) {
	if cb.size == 0 {
		return nil, false
	}

	item := cb.buffer[cb.head]
	cb.head = (cb.head + 1) % cb.capacity
	cb.size--
	return item, true
}

func (cb *CircularBuffer) Stats() (int, int, int) {
	return cb.size, cb.dropped, cb.capacity
}

// === LEAKY BUCKET ALGORITHM ===

type LeakyBucket struct {
	capacity   int
	tokens     int
	leakRate   time.Duration
	lastRefill time.Time
	mutex      chan bool
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity:   capacity,
		tokens:     capacity,
		leakRate:   leakRate,
		lastRefill: time.Now(),
		mutex:      make(chan bool, 1),
	}
}

func (lb *LeakyBucket) leak() {
	lb.mutex <- true
	defer func() { <-lb.mutex }()

	now := time.Now()
	elapsed := now.Sub(lb.lastRefill)

	// Calculate tokens to remove based on leak rate
	tokensToRemove := int(elapsed / lb.leakRate)
	if tokensToRemove > 0 {
		lb.tokens = max(0, lb.tokens-tokensToRemove)
		lb.lastRefill = now
	}
}

func (lb *LeakyBucket) AddTokens(tokens int) int {
	lb.leak()

	lb.mutex <- true
	defer func() { <-lb.mutex }()

	// Add tokens, but don't exceed capacity
	oldTokens := lb.tokens
	lb.tokens = min(lb.capacity, lb.tokens+tokens)

	// Return number of tokens that were dropped
	return (oldTokens + tokens) - lb.tokens
}

func (lb *LeakyBucket) ConsumeTokens(tokens int) bool {
	lb.leak()

	lb.mutex <- true
	defer func() { <-lb.mutex }()

	if lb.tokens >= tokens {
		lb.tokens -= tokens
		return true
	}
	return false
}

func (lb *LeakyBucket) CurrentTokens() int {
	lb.leak()

	lb.mutex <- true
	defer func() { <-lb.mutex }()

	return lb.tokens
}

// === HELPER FUNCTIONS ===

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// === MAIN FUNCTION ===

func main() {
	fmt.Println("=== GO LEAKY BUFFER COMPREHENSIVE GUIDE ===")

	// === BASIC LEAKY BUFFER ===
	fmt.Println("\n1. BASIC LEAKY BUFFER:")

	buffer := NewLeakyBuffer("test-buffer", 5)

	// Push items to buffer
	for i := 1; i <= 10; i++ {
		pushed := buffer.Push(fmt.Sprintf("item-%d", i))
		fmt.Printf("Pushed item-%d: %t\n", i, pushed)
	}

	// Check stats
	received, dropped, buffered := buffer.Stats()
	fmt.Printf("Stats - Received: %d, Dropped: %d, Buffered: %d\n", received, dropped, buffered)

	// Pop items from buffer
	fmt.Println("\nPopping items:")
	for i := 0; i < 7; i++ {
		item, ok := buffer.Pop()
		if ok {
			fmt.Printf("Popped: %v\n", item)
		} else {
			fmt.Println("Buffer is empty")
		}
	}

	// === PROCESSING LEAKY BUFFER ===
	fmt.Println("\n2. PROCESSING LEAKY BUFFER:")

	// Create processor function
	processor := func(item interface{}) {
		fmt.Printf("Processing: %v\n", item)
		time.Sleep(100 * time.Millisecond) // Simulate processing time
	}

	procBuffer := NewProcessingLeakyBuffer("proc-buffer", 3, processor)

	// Push items rapidly
	for i := 1; i <= 8; i++ {
		pushed := procBuffer.Push(fmt.Sprintf("task-%d", i))
		fmt.Printf("Pushed task-%d: %t\n", i, pushed)
		time.Sleep(50 * time.Millisecond)
	}

	// Wait for processing
	time.Sleep(1 * time.Second)

	// Check stats
	received, dropped, processed, buffered := procBuffer.Stats()
	fmt.Printf("Processing Stats - Received: %d, Dropped: %d, Processed: %d, Buffered: %d\n",
		received, dropped, processed, buffered)

	procBuffer.Stop()

	// === RATE LIMITER EXAMPLE ===
	fmt.Println("\n3. RATE LIMITER EXAMPLE:")

	// Create rate limiter (5 tokens, refill every 200ms)
	limiter := NewRateLimiter(5, 200*time.Millisecond)

	// Test rate limiting
	for i := 1; i <= 10; i++ {
		allowed := limiter.Allow()
		fmt.Printf("Request %d: %t\n", i, allowed)
		time.Sleep(100 * time.Millisecond)
	}

	limiter.Stop()

	// === CIRCULAR BUFFER EXAMPLE ===
	fmt.Println("\n4. CIRCULAR BUFFER EXAMPLE:")

	circular := NewCircularBuffer(4)

	// Fill buffer and overflow
	for i := 1; i <= 8; i++ {
		circular.Push(fmt.Sprintf("data-%d", i))
		size, dropped, capacity := circular.Stats()
		fmt.Printf("Pushed data-%d - Size: %d, Dropped: %d, Capacity: %d\n",
			i, size, dropped, capacity)
	}

	// Pop all items
	fmt.Println("\nPopping from circular buffer:")
	for {
		item, ok := circular.Pop()
		if !ok {
			break
		}
		fmt.Printf("Popped: %v\n", item)
	}

	// === LEAKY BUCKET ALGORITHM ===
	fmt.Println("\n5. LEAKY BUCKET ALGORITHM:")

	// Create leaky bucket (10 tokens capacity, leak 1 token per 100ms)
	bucket := NewLeakyBucket(10, 100*time.Millisecond)

	fmt.Printf("Initial tokens: %d\n", bucket.CurrentTokens())

	// Add tokens (simulate requests)
	for i := 1; i <= 5; i++ {
		dropped := bucket.AddTokens(3)
		fmt.Printf("Added 3 tokens, dropped: %d, current: %d\n",
			dropped, bucket.CurrentTokens())
		time.Sleep(50 * time.Millisecond)
	}

	// Wait for leaking
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("After leaking: %d tokens\n", bucket.CurrentTokens())

	// Consume tokens
	for i := 1; i <= 3; i++ {
		consumed := bucket.ConsumeTokens(2)
		fmt.Printf("Consume 2 tokens: %t, remaining: %d\n",
			consumed, bucket.CurrentTokens())
	}

	// === REAL-WORLD SCENARIOS ===
	fmt.Println("\n6. REAL-WORLD SCENARIOS:")

	// Scenario 1: High-frequency logging
	fmt.Println("\nScenario 1: High-frequency logging")

	logBuffer := NewLeakyBuffer("log-buffer", 5)

	// Simulate high-frequency log generation
	for i := 1; i <= 20; i++ {
		logEntry := fmt.Sprintf("LOG[%d]: Operation completed", i)
		pushed := logBuffer.Push(logEntry)
		if !pushed {
			fmt.Printf("Log entry %d dropped\n", i)
		}

		// Occasionally drain buffer
		if i%5 == 0 {
			fmt.Printf("Draining buffer at entry %d\n", i)
			for j := 0; j < 3; j++ {
				if entry, ok := logBuffer.Pop(); ok {
					fmt.Printf("  Logged: %v\n", entry)
				}
			}
		}
	}

	// Scenario 2: Event processing with backpressure
	fmt.Println("\nScenario 2: Event processing with backpressure")

	eventProcessor := func(event interface{}) {
		fmt.Printf("Processing event: %v\n", event)
		time.Sleep(200 * time.Millisecond) // Simulate slow processing
	}

	eventBuffer := NewProcessingLeakyBuffer("event-buffer", 3, eventProcessor)

	// Generate events rapidly
	for i := 1; i <= 10; i++ {
		event := fmt.Sprintf("Event-%d", i)
		pushed := eventBuffer.Push(event)
		if !pushed {
			fmt.Printf("Event %d dropped due to backpressure\n", i)
		}
		time.Sleep(50 * time.Millisecond)
	}

	// Wait for processing
	time.Sleep(2 * time.Second)
	eventBuffer.Stop()

	// === PERFORMANCE COMPARISON ===
	fmt.Println("\n7. PERFORMANCE COMPARISON:")

	// Compare different buffer types
	const iterations = 10000

	// Leaky buffer performance
	start := time.Now()
	testBuffer := NewLeakyBuffer("perf-test", 100)
	for i := 0; i < iterations; i++ {
		testBuffer.Push(i)
	}
	leakyDuration := time.Since(start)

	// Circular buffer performance
	start = time.Now()
	testCircular := NewCircularBuffer(100)
	for i := 0; i < iterations; i++ {
		testCircular.Push(i)
	}
	circularDuration := time.Since(start)

	fmt.Printf("Leaky buffer: %v for %d operations\n", leakyDuration, iterations)
	fmt.Printf("Circular buffer: %v for %d operations\n", circularDuration, iterations)

	// === LEAKY BUFFER CHARACTERISTICS ===
	fmt.Println("\n8. LEAKY BUFFER CHARACTERISTICS:")

	fmt.Println("✓ Prevents blocking when buffer is full")
	fmt.Println("✓ Drops data when overflow occurs")
	fmt.Println("✓ Useful for real-time systems")
	fmt.Println("✓ Implements backpressure handling")
	fmt.Println("✓ Maintains system responsiveness")
	fmt.Println("✓ Configurable buffer size")
	fmt.Println("✓ Non-blocking push operations")

	// === BEST PRACTICES ===
	fmt.Println("\n9. BEST PRACTICES:")

	fmt.Println("✓ Choose appropriate buffer size")
	fmt.Println("✓ Monitor drop rates")
	fmt.Println("✓ Implement proper error handling")
	fmt.Println("✓ Use for non-critical data")
	fmt.Println("✓ Consider data importance")
	fmt.Println("✓ Implement backpressure signals")
	fmt.Println("✓ Log dropped items if needed")

	// === WHEN TO USE LEAKY BUFFERS ===
	fmt.Println("\n10. WHEN TO USE LEAKY BUFFERS:")

	fmt.Println("✓ Real-time data processing")
	fmt.Println("✓ High-frequency logging")
	fmt.Println("✓ Event streaming")
	fmt.Println("✓ Rate limiting")
	fmt.Println("✓ Load shedding")
	fmt.Println("✓ Metrics collection")
	fmt.Println("✓ Network packet processing")

	// === WHEN NOT TO USE LEAKY BUFFERS ===
	fmt.Println("\n11. WHEN NOT TO USE LEAKY BUFFERS:")

	fmt.Println("✗ Critical data that cannot be lost")
	fmt.Println("✗ Financial transactions")
	fmt.Println("✗ User commands")
	fmt.Println("✗ Database operations")
	fmt.Println("✗ File operations")
	fmt.Println("✗ Security-related data")
	fmt.Println("✗ Configuration updates")

	// === MONITORING AND DEBUGGING ===
	fmt.Println("\n12. MONITORING AND DEBUGGING:")

	fmt.Println("Monitoring metrics:")
	fmt.Println("  - Buffer utilization")
	fmt.Println("  - Drop rate")
	fmt.Println("  - Processing rate")
	fmt.Println("  - Throughput")
	fmt.Println("  - Latency")

	fmt.Println("Debugging techniques:")
	fmt.Println("  - Track buffer statistics")
	fmt.Println("  - Log dropped items")
	fmt.Println("  - Monitor goroutine count")
	fmt.Println("  - Profile memory usage")
	fmt.Println("  - Analyze processing patterns")

	fmt.Printf("\nFinal goroutine count: %d\n", runtime.NumGoroutine())

	fmt.Println("\n=== END OF LEAKY BUFFER GUIDE ===")
}

/*
RUNNING THE PROGRAM:
go run main.go

LEARNING POINTS:

1. LEAKY BUFFER CONCEPT:
   - Buffered channel that drops data on overflow
   - Prevents blocking when consumer is slow
   - Non-blocking push operations
   - Configurable buffer size

2. IMPLEMENTATIONS:
   - Basic leaky buffer with channels
   - Processing leaky buffer with goroutines
   - Circular buffer with array backing
   - Rate limiter using token bucket
   - Leaky bucket algorithm

3. USE CASES:
   - Real-time data processing
   - High-frequency logging
   - Event streaming
   - Rate limiting
   - Load shedding
   - Metrics collection

4. CHARACTERISTICS:
   - Prevents system blocking
   - Implements backpressure handling
   - Maintains system responsiveness
   - Configurable behavior
   - Non-blocking operations

5. BEST PRACTICES:
   - Choose appropriate buffer size
   - Monitor drop rates
   - Use for non-critical data
   - Implement proper error handling
   - Log dropped items if needed

6. MONITORING:
   - Track buffer utilization
   - Monitor drop rates
   - Analyze processing patterns
   - Profile memory usage
   - Measure throughput and latency

This demonstrates comprehensive leaky buffer patterns in Go
for handling backpressure and overflow scenarios.
*/
