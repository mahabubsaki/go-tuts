package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// === GO ALLOCATION WITH MAKE COMPREHENSIVE GUIDE ===

/*
MAKE FUNCTION PHILOSOPHY:
- make() is used for slices, maps, and channels only
- Creates and initializes the data structure
- Returns the initialized value (not a pointer)
- Different from new() which returns a pointer to zero value

COMPARISON WITH JAVASCRIPT:
// JavaScript - Collection creation
const arr = new Array(5);        // Array with length 5
const obj = {};                  // Empty object
const map = new Map();          // Map collection

// Go - Collection creation
slice := make([]int, 5)         // Slice with length 5
mapVal := make(map[string]int)  // Empty map
ch := make(chan int)            // Unbuffered channel
*/

func main() {
	fmt.Println("=== GO ALLOCATION WITH MAKE COMPREHENSIVE GUIDE ===")

	// === SLICE ALLOCATION ===
	fmt.Println("\n1. SLICE ALLOCATION:")

	// Basic slice creation
	slice1 := make([]int, 5)       // Length 5, capacity 5
	slice2 := make([]int, 3, 10)   // Length 3, capacity 10
	slice3 := make([]string, 0, 5) // Length 0, capacity 5

	fmt.Printf("slice1: %v, len=%d, cap=%d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2: %v, len=%d, cap=%d\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice3: %v, len=%d, cap=%d\n", slice3, len(slice3), cap(slice3))

	// Initialize slice elements
	for i := 0; i < len(slice1); i++ {
		slice1[i] = i * i
	}
	fmt.Printf("slice1 after init: %v\n", slice1)

	// Append to slice with capacity
	slice3 = append(slice3, "hello", "world", "go")
	fmt.Printf("slice3 after append: %v, len=%d, cap=%d\n", slice3, len(slice3), cap(slice3))

	// === MAP ALLOCATION ===
	fmt.Println("\n2. MAP ALLOCATION:")

	// Basic map creation
	map1 := make(map[string]int)
	map2 := make(map[int]string)
	map3 := make(map[string][]int)

	fmt.Printf("map1: %v, len=%d\n", map1, len(map1))
	fmt.Printf("map2: %v, len=%d\n", map2, len(map2))
	fmt.Printf("map3: %v, len=%d\n", map3, len(map3))

	// Initialize map elements
	map1["apple"] = 5
	map1["banana"] = 3
	map1["orange"] = 8

	map2[1] = "one"
	map2[2] = "two"
	map2[3] = "three"

	map3["numbers"] = []int{1, 2, 3}
	map3["squares"] = []int{1, 4, 9}

	fmt.Printf("map1 after init: %v\n", map1)
	fmt.Printf("map2 after init: %v\n", map2)
	fmt.Printf("map3 after init: %v\n", map3)

	// Map with initial capacity hint
	bigMap := make(map[string]int, 100)
	fmt.Printf("bigMap: %v, len=%d\n", bigMap, len(bigMap))

	// === CHANNEL ALLOCATION ===
	fmt.Println("\n3. CHANNEL ALLOCATION:")

	// Unbuffered channel
	ch1 := make(chan int)
	fmt.Printf("ch1: %v, len=%d, cap=%d\n", ch1, len(ch1), cap(ch1))

	// Buffered channels
	ch2 := make(chan string, 3)
	ch3 := make(chan bool, 1)

	fmt.Printf("ch2: %v, len=%d, cap=%d\n", ch2, len(ch2), cap(ch2))
	fmt.Printf("ch3: %v, len=%d, cap=%d\n", ch3, len(ch3), cap(ch3))

	// Send values to buffered channel
	ch2 <- "hello"
	ch2 <- "world"
	fmt.Printf("ch2 after sends: len=%d, cap=%d\n", len(ch2), cap(ch2))

	// Receive from channel
	msg1 := <-ch2
	msg2 := <-ch2
	fmt.Printf("Received: %s, %s\n", msg1, msg2)
	fmt.Printf("ch2 after receives: len=%d, cap=%d\n", len(ch2), cap(ch2))

	// Close channel
	close(ch2)

	// === COMPARISON: MAKE VS NEW ===
	fmt.Println("\n4. COMPARISON: MAKE VS NEW:")

	// Slice comparison
	sliceNew := new([]int)      // Pointer to nil slice
	sliceMake := make([]int, 5) // Actual slice with length 5

	fmt.Printf("new([]int): %v, is nil: %t\n", *sliceNew, *sliceNew == nil)
	fmt.Printf("make([]int, 5): %v, is nil: %t\n", sliceMake, sliceMake == nil)

	// Map comparison
	mapNew := new(map[string]int)   // Pointer to nil map
	mapMake := make(map[string]int) // Actual initialized map

	fmt.Printf("new(map): %v, is nil: %t\n", *mapNew, *mapNew == nil)
	fmt.Printf("make(map): %v, is nil: %t\n", mapMake, mapMake == nil)

	// Channel comparison
	chanNew := new(chan int)      // Pointer to nil channel
	chanMake := make(chan int, 1) // Actual channel

	fmt.Printf("new(chan): %v, is nil: %t\n", *chanNew, *chanNew == nil)
	fmt.Printf("make(chan): %v, is nil: %t\n", chanMake, chanMake == nil)

	// === SLICE CAPACITY AND GROWTH ===
	fmt.Println("\n5. SLICE CAPACITY AND GROWTH:")

	// Start with small slice
	numbers := make([]int, 0, 2)
	fmt.Printf("Initial: len=%d, cap=%d\n", len(numbers), cap(numbers))

	// Append and watch capacity growth
	for i := 1; i <= 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("After append %d: len=%d, cap=%d\n", i, len(numbers), cap(numbers))
	}

	// Pre-allocate if you know the size
	prealloc := make([]int, 0, 10)
	for i := 1; i <= 10; i++ {
		prealloc = append(prealloc, i)
	}
	fmt.Printf("Pre-allocated: len=%d, cap=%d\n", len(prealloc), cap(prealloc))

	// === MULTI-DIMENSIONAL SLICES ===
	fmt.Println("\n6. MULTI-DIMENSIONAL SLICES:")

	// 2D slice
	matrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		matrix[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			matrix[i][j] = i*4 + j + 1
		}
	}

	fmt.Println("2D slice:")
	for i := 0; i < len(matrix); i++ {
		fmt.Printf("Row %d: %v\n", i, matrix[i])
	}

	// 3D slice
	cube := make([][][]int, 2)
	for i := 0; i < 2; i++ {
		cube[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			cube[i][j] = make([]int, 2)
			for k := 0; k < 2; k++ {
				cube[i][j][k] = i*4 + j*2 + k + 1
			}
		}
	}

	fmt.Println("3D slice:")
	for i := 0; i < len(cube); i++ {
		fmt.Printf("Layer %d:\n", i)
		for j := 0; j < len(cube[i]); j++ {
			fmt.Printf("  Row %d: %v\n", j, cube[i][j])
		}
	}

	// === MAP PERFORMANCE ===
	fmt.Println("\n7. MAP PERFORMANCE:")

	// Map without capacity hint
	map1_perf := make(map[int]string)
	for i := 0; i < 1000; i++ {
		map1_perf[i] = fmt.Sprintf("value_%d", i)
	}
	fmt.Printf("Map without hint: len=%d\n", len(map1_perf))

	// Map with capacity hint (better performance)
	map2_perf := make(map[int]string, 1000)
	for i := 0; i < 1000; i++ {
		map2_perf[i] = fmt.Sprintf("value_%d", i)
	}
	fmt.Printf("Map with hint: len=%d\n", len(map2_perf))

	// === CHANNEL PATTERNS ===
	fmt.Println("\n8. CHANNEL PATTERNS:")

	// Worker pool pattern
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Start workers
	for w := 1; w <= 3; w++ {
		go func(id int) {
			for job := range jobs {
				result := job * 2
				results <- result
			}
		}(w)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}

	// Done channel pattern
	done := make(chan bool)
	go func() {
		fmt.Println("Working...")
		// Simulate work
		done <- true
	}()

	<-done
	fmt.Println("Work completed")

	// === MEMORY USAGE ===
	fmt.Println("\n9. MEMORY USAGE:")

	// Slice memory usage
	bigSlice := make([]int, 1000000)
	fmt.Printf("Slice size: %d bytes\n", int(unsafe.Sizeof(bigSlice)))
	fmt.Printf("Slice header size: %d bytes\n", int(unsafe.Sizeof([]int{})))
	fmt.Printf("Slice data size: %d bytes\n", len(bigSlice)*int(unsafe.Sizeof(int(0))))

	// Map memory usage
	bigMap2 := make(map[string]int)
	for i := 0; i < 1000; i++ {
		bigMap2[fmt.Sprintf("key_%d", i)] = i
	}
	fmt.Printf("Map size: %d bytes\n", int(unsafe.Sizeof(bigMap2)))

	// Channel memory usage
	bigChan := make(chan int, 1000)
	fmt.Printf("Channel size: %d bytes\n", int(unsafe.Sizeof(bigChan)))

	// === REFLECTION WITH MAKE ===
	fmt.Println("\n10. REFLECTION WITH MAKE:")

	// Create slice using reflection
	sliceType := reflect.TypeOf([]int{})
	sliceValue := reflect.MakeSlice(sliceType, 5, 10)

	// Set values using reflection
	for i := 0; i < 5; i++ {
		sliceValue.Index(i).SetInt(int64(i * 10))
	}

	fmt.Printf("Reflected slice: %v\n", sliceValue.Interface())

	// Create map using reflection
	mapType := reflect.TypeOf(map[string]int{})
	mapValue := reflect.MakeMap(mapType)

	// Set values using reflection
	mapValue.SetMapIndex(reflect.ValueOf("key1"), reflect.ValueOf(100))
	mapValue.SetMapIndex(reflect.ValueOf("key2"), reflect.ValueOf(200))

	fmt.Printf("Reflected map: %v\n", mapValue.Interface())

	// Create channel using reflection
	chanType := reflect.TypeOf(make(chan int))
	chanValue := reflect.MakeChan(chanType, 2)

	// Send value using reflection
	chanValue.Send(reflect.ValueOf(42))

	// Receive value using reflection
	received, ok := chanValue.Recv()
	fmt.Printf("Reflected channel recv: %v, ok: %t\n", received.Interface(), ok)

	// === BEST PRACTICES ===
	fmt.Println("\n11. BEST PRACTICES:")

	fmt.Println("✓ Use make() for slices, maps, and channels")
	fmt.Println("✓ Specify capacity when you know the approximate size")
	fmt.Println("✓ Use buffered channels for async communication")
	fmt.Println("✓ Pre-allocate slices if you know the final size")
	fmt.Println("✓ Use capacity hints for maps with many elements")
	fmt.Println("✗ Don't use make() for basic types (use new() or literals)")
	fmt.Println("✗ Don't forget to close channels when done")

	// === COMMON PATTERNS ===
	fmt.Println("\n12. COMMON PATTERNS:")

	// Builder pattern with slice
	builder := make([]string, 0, 10)
	builder = append(builder, "Hello")
	builder = append(builder, " ")
	builder = append(builder, "World")
	result := fmt.Sprintf("%s", builder)
	fmt.Printf("Builder result: %s\n", result)

	// Cache pattern with map
	cache := make(map[string]interface{})
	cache["user_1"] = "Alice"
	cache["user_2"] = "Bob"
	cache["config"] = map[string]int{"timeout": 30}

	fmt.Printf("Cache: %v\n", cache)

	// Event system with channels
	events := make(chan string, 5)

	// Event listener
	go func() {
		for event := range events {
			fmt.Printf("Event received: %s\n", event)
		}
	}()

	// Send events
	events <- "user_login"
	events <- "data_saved"
	events <- "user_logout"

	close(events)

	// Wait for events to be processed
	fmt.Println("Events sent")

	// === MEMORY EFFICIENCY ===
	fmt.Println("\n13. MEMORY EFFICIENCY:")

	// Efficient slice creation
	data := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		data = append(data, i)
	}
	fmt.Printf("Efficient slice: len=%d, cap=%d\n", len(data), cap(data))

	// Efficient map creation
	lookup := make(map[string]int, 1000)
	for i := 0; i < 1000; i++ {
		lookup[fmt.Sprintf("key_%d", i)] = i
	}
	fmt.Printf("Efficient map: len=%d\n", len(lookup))

	// Efficient channel creation
	pipeline := make(chan int, 100)
	fmt.Printf("Efficient channel: cap=%d\n", cap(pipeline))
	close(pipeline)

	fmt.Println("\n=== END OF MAKE ALLOCATION GUIDE ===")
}

/*
RUNNING THE PROGRAM:
go run main.go

LEARNING POINTS:

1. MAKE() FUNCTION:
   - Only for slices, maps, and channels
   - Returns initialized value (not pointer)
   - Creates ready-to-use data structures

2. SLICE ALLOCATION:
   - make([]T, length) - length and capacity equal
   - make([]T, length, capacity) - specify both
   - Zero values filled in allocated elements

3. MAP ALLOCATION:
   - make(map[K]V) - empty map
   - make(map[K]V, hint) - with capacity hint
   - Better performance with capacity hints

4. CHANNEL ALLOCATION:
   - make(chan T) - unbuffered channel
   - make(chan T, buffer) - buffered channel
   - Buffer size affects blocking behavior

5. PERFORMANCE TIPS:
   - Pre-allocate slices when size is known
   - Use capacity hints for maps
   - Choose appropriate buffer sizes for channels
   - Consider memory access patterns

6. COMMON PATTERNS:
   - Builder pattern with slices
   - Cache pattern with maps
   - Worker pools with channels
   - Event systems with channels

This demonstrates comprehensive make() allocation patterns in Go
for efficient data structure creation and initialization.
*/
