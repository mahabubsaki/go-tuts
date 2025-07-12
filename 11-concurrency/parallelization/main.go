package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// === GO PARALLELIZATION COMPREHENSIVE GUIDE ===

/*
PARALLELIZATION PHILOSOPHY:
- Concurrency is about dealing with lots of things at once
- Parallelism is about doing lots of things at once
- Go provides tools for both concurrent and parallel programming
- GOMAXPROCS controls the number of OS threads
- Goroutines are multiplexed onto OS threads

COMPARISON WITH JAVASCRIPT:
// JavaScript - Single-threaded with event loop
// Web Workers for parallelism
const worker = new Worker('worker.js');
worker.postMessage({data: 'process this'});
worker.onmessage = function(e) {
  console.log('Result:', e.data);
};

// Go - Built-in parallelism
go func() {
  // This runs in parallel
  result := processData(data)
  resultChan <- result
}()
*/

// === WORK TYPES ===

type Task struct {
	ID     int
	Data   []int
	Result chan int
}

type WorkerPool struct {
	workers int
	tasks   chan Task
	wg      sync.WaitGroup
	stop    chan bool
	stopped bool
	mu      sync.Mutex
}

func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		workers: workers,
		tasks:   make(chan Task, workers*2),
		stop:    make(chan bool),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()

	for {
		select {
		case task := <-wp.tasks:
			// Process task
			sum := 0
			for _, v := range task.Data {
				sum += v
			}
			task.Result <- sum

		case <-wp.stop:
			fmt.Printf("Worker %d stopped\n", id)
			return
		}
	}
}

func (wp *WorkerPool) AddTask(task Task) {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	if !wp.stopped {
		wp.tasks <- task
	}
}

func (wp *WorkerPool) Stop() {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	if !wp.stopped {
		wp.stopped = true
		close(wp.stop)
		wp.wg.Wait()
	}
}

// === PARALLEL PROCESSING FUNCTIONS ===

func sequentialSum(data []int) int {
	sum := 0
	for _, v := range data {
		sum += v
	}
	return sum
}

func parallelSum(data []int, numWorkers int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / numWorkers
	if chunkSize == 0 {
		chunkSize = 1
		numWorkers = len(data)
	}

	results := make(chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numWorkers-1 {
			end = len(data)
		}

		go func(chunk []int) {
			sum := 0
			for _, v := range chunk {
				sum += v
			}
			results <- sum
		}(data[start:end])
	}

	totalSum := 0
	for i := 0; i < numWorkers; i++ {
		totalSum += <-results
	}

	return totalSum
}

func parallelMap(data []int, fn func(int) int, numWorkers int) []int {
	if len(data) == 0 {
		return []int{}
	}

	chunkSize := len(data) / numWorkers
	if chunkSize == 0 {
		chunkSize = 1
		numWorkers = len(data)
	}

	results := make([][]int, numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numWorkers-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(workerID int, chunk []int) {
			defer wg.Done()

			result := make([]int, len(chunk))
			for j, v := range chunk {
				result[j] = fn(v)
			}
			results[workerID] = result
		}(i, data[start:end])
	}

	wg.Wait()

	// Combine results
	var combined []int
	for _, result := range results {
		combined = append(combined, result...)
	}

	return combined
}

func parallelFilter(data []int, predicate func(int) bool, numWorkers int) []int {
	if len(data) == 0 {
		return []int{}
	}

	chunkSize := len(data) / numWorkers
	if chunkSize == 0 {
		chunkSize = 1
		numWorkers = len(data)
	}

	results := make([][]int, numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numWorkers-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(workerID int, chunk []int) {
			defer wg.Done()

			var filtered []int
			for _, v := range chunk {
				if predicate(v) {
					filtered = append(filtered, v)
				}
			}
			results[workerID] = filtered
		}(i, data[start:end])
	}

	wg.Wait()

	// Combine results
	var combined []int
	for _, result := range results {
		combined = append(combined, result...)
	}

	return combined
}

// === PIPELINE STAGES ===

func generateNumbers(count int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := 1; i <= count; i++ {
			out <- i
		}
	}()

	return out
}

func squareNumbers(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()

	return out
}

func filterEven(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
	}()

	return out
}

// === MAIN FUNCTION ===

func main() {
	fmt.Println("=== GO PARALLELIZATION COMPREHENSIVE GUIDE ===")

	// === RUNTIME INFORMATION ===
	fmt.Println("\n1. RUNTIME INFORMATION:")

	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("Number of Goroutines: %d\n", runtime.NumGoroutine())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

	// Set GOMAXPROCS to use all CPUs
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("GOMAXPROCS set to: %d\n", runtime.GOMAXPROCS(0))

	// === SEQUENTIAL VS PARALLEL COMPARISON ===
	fmt.Println("\n2. SEQUENTIAL VS PARALLEL COMPARISON:")

	// Generate test data
	data := make([]int, 1000000)
	for i := range data {
		data[i] = rand.Intn(1000)
	}

	// Sequential processing
	start := time.Now()
	seqSum := sequentialSum(data)
	seqDuration := time.Since(start)

	// Parallel processing
	start = time.Now()
	parSum := parallelSum(data, runtime.NumCPU())
	parDuration := time.Since(start)

	fmt.Printf("Sequential sum: %d (took %v)\n", seqSum, seqDuration)
	fmt.Printf("Parallel sum: %d (took %v)\n", parSum, parDuration)
	fmt.Printf("Speedup: %.2fx\n", float64(seqDuration)/float64(parDuration))

	// === WORKER POOL PATTERN ===
	fmt.Println("\n3. WORKER POOL PATTERN:")

	// Create worker pool
	pool := NewWorkerPool(runtime.NumCPU())
	pool.Start()

	// Submit tasks
	numTasks := 20
	results := make([]chan int, numTasks)

	for i := 0; i < numTasks; i++ {
		result := make(chan int, 1)
		results[i] = result

		taskData := make([]int, 1000)
		for j := range taskData {
			taskData[j] = rand.Intn(100)
		}

		task := Task{
			ID:     i,
			Data:   taskData,
			Result: result,
		}

		pool.AddTask(task)
	}

	// Collect results
	for i, result := range results {
		sum := <-result
		fmt.Printf("Task %d result: %d\n", i, sum)
	}

	pool.Stop()

	// === PARALLEL MAP OPERATION ===
	fmt.Println("\n4. PARALLEL MAP OPERATION:")

	smallData := make([]int, 20)
	for i := range smallData {
		smallData[i] = i + 1
	}

	// Sequential map
	start = time.Now()
	seqMapped := make([]int, len(smallData))
	for i, v := range smallData {
		seqMapped[i] = v * v
	}
	seqMapDuration := time.Since(start)

	// Parallel map
	start = time.Now()
	parMapped := parallelMap(smallData, func(x int) int { return x * x }, 4)
	parMapDuration := time.Since(start)

	fmt.Printf("Sequential map: %v (took %v)\n", seqMapped, seqMapDuration)
	fmt.Printf("Parallel map: %v (took %v)\n", parMapped, parMapDuration)

	// === PARALLEL FILTER OPERATION ===
	fmt.Println("\n5. PARALLEL FILTER OPERATION:")

	// Sequential filter
	start = time.Now()
	var seqFiltered []int
	for _, v := range smallData {
		if v%2 == 0 {
			seqFiltered = append(seqFiltered, v)
		}
	}
	seqFilterDuration := time.Since(start)

	// Parallel filter
	start = time.Now()
	parFiltered := parallelFilter(smallData, func(x int) bool { return x%2 == 0 }, 4)
	parFilterDuration := time.Since(start)

	fmt.Printf("Sequential filter: %v (took %v)\n", seqFiltered, seqFilterDuration)
	fmt.Printf("Parallel filter: %v (took %v)\n", parFiltered, parFilterDuration)

	// === PARALLEL PIPELINE ===
	fmt.Println("\n6. PARALLEL PIPELINE:")

	// Create pipeline
	numbers := generateNumbers(20)
	squares := squareNumbers(numbers)
	evens := filterEven(squares)

	// Collect results
	fmt.Print("Pipeline results: ")
	for result := range evens {
		fmt.Printf("%d ", result)
	}
	fmt.Println()

	// === FAN-OUT FAN-IN PATTERN ===
	fmt.Println("\n7. FAN-OUT FAN-IN PATTERN:")

	// Input channel
	input := make(chan int, 10)

	// Fan-out to multiple workers
	numWorkers := 3
	workers := make([]<-chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		workers[i] = func(workerID int) <-chan int {
			out := make(chan int)
			go func() {
				defer close(out)
				for n := range input {
					// Simulate work
					time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
					result := n * n
					fmt.Printf("Worker %d: %d -> %d\n", workerID, n, result)
					out <- result
				}
			}()
			return out
		}(i)
	}

	// Fan-in from workers
	fanIn := func(inputs ...<-chan int) <-chan int {
		out := make(chan int)
		var wg sync.WaitGroup

		for _, input := range inputs {
			wg.Add(1)
			go func(ch <-chan int) {
				defer wg.Done()
				for n := range ch {
					out <- n
				}
			}(input)
		}

		go func() {
			wg.Wait()
			close(out)
		}()

		return out
	}

	output := fanIn(workers...)

	// Send work
	go func() {
		defer close(input)
		for i := 1; i <= 10; i++ {
			input <- i
		}
	}()

	// Collect results
	var fanResults []int
	for result := range output {
		fanResults = append(fanResults, result)
	}

	fmt.Printf("Fan-out/Fan-in results: %v\n", fanResults)

	// === CPU-BOUND VS I/O-BOUND TASKS ===
	fmt.Println("\n8. CPU-BOUND VS I/O-BOUND TASKS:")

	// CPU-bound task
	cpuBoundTask := func() int {
		sum := 0
		for i := 0; i < 1000000; i++ {
			sum += i
		}
		return sum
	}

	// I/O-bound task (simulated)
	ioBoundTask := func() string {
		time.Sleep(100 * time.Millisecond)
		return "I/O operation completed"
	}

	// Run CPU-bound tasks in parallel
	start = time.Now()
	var cpuWg sync.WaitGroup
	cpuResults := make([]int, 4)

	for i := 0; i < 4; i++ {
		cpuWg.Add(1)
		go func(index int) {
			defer cpuWg.Done()
			cpuResults[index] = cpuBoundTask()
		}(i)
	}
	cpuWg.Wait()
	cpuDuration := time.Since(start)

	// Run I/O-bound tasks in parallel
	start = time.Now()
	var ioWg sync.WaitGroup
	ioResults := make([]string, 4)

	for i := 0; i < 4; i++ {
		ioWg.Add(1)
		go func(index int) {
			defer ioWg.Done()
			ioResults[index] = ioBoundTask()
		}(i)
	}
	ioWg.Wait()
	ioDuration := time.Since(start)

	fmt.Printf("CPU-bound tasks completed in: %v\n", cpuDuration)
	fmt.Printf("I/O-bound tasks completed in: %v\n", ioDuration)

	// === PARALLELIZATION STRATEGIES ===
	fmt.Println("\n9. PARALLELIZATION STRATEGIES:")

	fmt.Println("Data Parallelism:")
	fmt.Println("  - Split data into chunks")
	fmt.Println("  - Process chunks in parallel")
	fmt.Println("  - Combine results")

	fmt.Println("Task Parallelism:")
	fmt.Println("  - Different tasks run simultaneously")
	fmt.Println("  - Tasks may operate on different data")
	fmt.Println("  - Coordination via channels/sync primitives")

	fmt.Println("Pipeline Parallelism:")
	fmt.Println("  - Process data through stages")
	fmt.Println("  - Each stage runs in parallel")
	fmt.Println("  - Data flows through pipeline")

	// === PERFORMANCE CONSIDERATIONS ===
	fmt.Println("\n10. PERFORMANCE CONSIDERATIONS:")

	fmt.Println("✓ Consider overhead of goroutine creation")
	fmt.Println("✓ Balance number of workers with CPU cores")
	fmt.Println("✓ Use proper synchronization mechanisms")
	fmt.Println("✓ Avoid false sharing in parallel access")
	fmt.Println("✓ Consider memory allocation patterns")
	fmt.Println("✓ Profile and measure performance gains")

	// === BEST PRACTICES ===
	fmt.Println("\n11. BEST PRACTICES:")

	fmt.Println("✓ Use runtime.NumCPU() for CPU-bound tasks")
	fmt.Println("✓ Use more goroutines for I/O-bound tasks")
	fmt.Println("✓ Implement proper error handling")
	fmt.Println("✓ Use context for cancellation")
	fmt.Println("✓ Monitor resource usage")
	fmt.Println("✓ Consider using sync.Pool for object reuse")
	fmt.Println("✓ Profile your parallel code")

	// === COMMON PITFALLS ===
	fmt.Println("\n12. COMMON PITFALLS:")

	fmt.Println("✗ Creating too many goroutines")
	fmt.Println("✗ Not considering overhead")
	fmt.Println("✗ Race conditions")
	fmt.Println("✗ Deadlocks")
	fmt.Println("✗ Resource leaks")
	fmt.Println("✗ False sharing")
	fmt.Println("✗ Not measuring performance")

	fmt.Printf("\nFinal Goroutine count: %d\n", runtime.NumGoroutine())

	fmt.Println("\n=== END OF PARALLELIZATION GUIDE ===")
}

/*
RUNNING THE PROGRAM:
go run main.go

LEARNING POINTS:

1. PARALLELIZATION CONCEPTS:
   - Concurrency vs Parallelism
   - GOMAXPROCS controls OS threads
   - Goroutines are multiplexed onto threads
   - CPU-bound vs I/O-bound tasks

2. PARALLEL PATTERNS:
   - Worker Pool: Distribute work across workers
   - Fan-out/Fan-in: Distribute and collect work
   - Pipeline: Sequential processing stages
   - Data Parallelism: Split data across workers
   - Task Parallelism: Different tasks simultaneously

3. PERFORMANCE STRATEGIES:
   - Use runtime.NumCPU() for CPU-bound tasks
   - Use more goroutines for I/O-bound tasks
   - Consider overhead of goroutine creation
   - Balance workers with available resources

4. SYNCHRONIZATION:
   - sync.WaitGroup for waiting on multiple goroutines
   - Channels for communication
   - sync.Mutex for shared state
   - context.Context for cancellation

5. BEST PRACTICES:
   - Profile and measure performance
   - Consider overhead vs benefits
   - Use proper synchronization
   - Monitor resource usage
   - Implement error handling

This demonstrates comprehensive parallelization techniques in Go
for high-performance concurrent applications.
*/
