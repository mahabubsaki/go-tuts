package main

import (
	"fmt"
	"sort"
)

// === MAPS IN GO ===

func main() {
	fmt.Println("=== GO MAPS COMPREHENSIVE GUIDE ===")

	// === BASIC MAP DECLARATION ===
	fmt.Println("\n--- BASIC MAP DECLARATION ---")

	// Different ways to declare maps
	var scores1 map[string]int
	fmt.Printf("Nil map: %v (len: %d)\n", scores1, len(scores1))

	// Map literal
	scores2 := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 92,
	}
	fmt.Printf("Map literal: %v (len: %d)\n", scores2, len(scores2))

	// Using make
	scores3 := make(map[string]int)
	fmt.Printf("Made map: %v (len: %d)\n", scores3, len(scores3))

	// Map with initial capacity (hint)
	scores4 := make(map[string]int, 10)
	fmt.Printf("Made map with capacity: %v (len: %d)\n", scores4, len(scores4))

	/*
		JavaScript comparison:
		// JavaScript objects and Maps
		const scores1 = {};
		const scores2 = {
			"Alice": 95,
			"Bob": 87,
			"Charlie": 92
		};
		const scores3 = new Map();
		const scores4 = new Map([
			["Alice", 95],
			["Bob", 87],
			["Charlie", 92]
		]);

		// JavaScript objects have prototype chain
		// Maps preserve insertion order
		// Maps can have any type as key
	*/

	// === MAP OPERATIONS ===
	fmt.Println("\n--- MAP OPERATIONS ---")

	// Initialize map
	ages := make(map[string]int)

	// Adding elements
	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35
	fmt.Printf("After adding: %v\n", ages)

	// Accessing elements
	fmt.Printf("Alice's age: %d\n", ages["Alice"])
	fmt.Printf("Non-existent key: %d\n", ages["David"]) // Returns zero value

	// Check if key exists
	age, exists := ages["Alice"]
	fmt.Printf("Alice exists: %t, age: %d\n", exists, age)

	age, exists = ages["David"]
	fmt.Printf("David exists: %t, age: %d\n", exists, age)

	// Updating elements
	ages["Alice"] = 31
	fmt.Printf("After updating Alice: %v\n", ages)

	// Deleting elements
	delete(ages, "Bob")
	fmt.Printf("After deleting Bob: %v\n", ages)

	// === MAP ITERATION ===
	fmt.Println("\n--- MAP ITERATION ---")

	colors := map[string]string{
		"red":    "#FF0000",
		"green":  "#00FF00",
		"blue":   "#0000FF",
		"yellow": "#FFFF00",
	}

	// Iterate over key-value pairs
	fmt.Println("Key-value pairs:")
	for color, hex := range colors {
		fmt.Printf("  %s: %s\n", color, hex)
	}

	// Iterate over keys only
	fmt.Println("Keys only:")
	for color := range colors {
		fmt.Printf("  %s\n", color)
	}

	// Iterate over values only
	fmt.Println("Values only:")
	for _, hex := range colors {
		fmt.Printf("  %s\n", hex)
	}

	// === MAP WITH DIFFERENT TYPES ===
	fmt.Println("\n--- MAP WITH DIFFERENT TYPES ---")

	// String to slice
	groups := map[string][]string{
		"fruits":     {"apple", "banana", "orange"},
		"vegetables": {"carrot", "broccoli", "spinach"},
		"grains":     {"rice", "wheat", "oats"},
	}

	fmt.Println("Groups:")
	for group, items := range groups {
		fmt.Printf("  %s: %v\n", group, items)
	}

	// Int to string
	httpStatus := map[int]string{
		200: "OK",
		404: "Not Found",
		500: "Internal Server Error",
	}

	fmt.Println("HTTP Status codes:")
	for code, message := range httpStatus {
		fmt.Printf("  %d: %s\n", code, message)
	}

	// === NESTED MAPS ===
	fmt.Println("\n--- NESTED MAPS ---")

	// Map of maps
	students := map[string]map[string]int{
		"Alice": {
			"Math":    95,
			"Science": 88,
			"English": 92,
		},
		"Bob": {
			"Math":    87,
			"Science": 94,
			"English": 78,
		},
	}

	fmt.Println("Student grades:")
	for student, grades := range students {
		fmt.Printf("  %s:\n", student)
		for subject, grade := range grades {
			fmt.Printf("    %s: %d\n", subject, grade)
		}
	}

	// Accessing nested map
	fmt.Printf("Alice's Math grade: %d\n", students["Alice"]["Math"])

	// === MAP FUNCTIONS ===
	fmt.Println("\n--- MAP FUNCTIONS ---")

	// Function that takes map
	printMap := func(m map[string]int) {
		for k, v := range m {
			fmt.Printf("  %s: %d\n", k, v)
		}
	}

	// Function that modifies map
	incrementValues := func(m map[string]int) {
		for k, v := range m {
			m[k] = v + 1
		}
	}

	// Function that returns map
	createCountMap := func(words []string) map[string]int {
		counts := make(map[string]int)
		for _, word := range words {
			counts[word]++
		}
		return counts
	}

	testMap := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("Original map:\n")
	printMap(testMap)

	incrementValues(testMap)
	fmt.Printf("After increment:\n")
	printMap(testMap)

	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	wordCounts := createCountMap(words)
	fmt.Printf("Word counts: %v\n", wordCounts)

	// === MAP COPYING ===
	fmt.Println("\n--- MAP COPYING ---")

	original := map[string]int{"a": 1, "b": 2, "c": 3}

	// Maps are reference types - assignment creates reference
	reference := original
	reference["a"] = 999
	fmt.Printf("Original after reference modification: %v\n", original)

	// Create a copy
	copyMap := make(map[string]int)
	for k, v := range original {
		copyMap[k] = v
	}

	copyMap["b"] = 888
	fmt.Printf("Original after copy modification: %v\n", original)
	fmt.Printf("Copy: %v\n", copyMap)

	// === MAP SORTING ===
	fmt.Println("\n--- MAP SORTING ---")

	population := map[string]int{
		"Tokyo":       13929286,
		"Delhi":       28514000,
		"Shanghai":    24256800,
		"São Paulo":   12252023,
		"Mexico City": 9209944,
	}

	// Sort by keys
	var keys []string
	for k := range population {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println("Sorted by city name:")
	for _, city := range keys {
		fmt.Printf("  %s: %d\n", city, population[city])
	}

	// Sort by values
	type cityPop struct {
		city string
		pop  int
	}

	var cities []cityPop
	for city, pop := range population {
		cities = append(cities, cityPop{city, pop})
	}

	sort.Slice(cities, func(i, j int) bool {
		return cities[i].pop > cities[j].pop // Descending order
	})

	fmt.Println("Sorted by population (descending):")
	for _, cp := range cities {
		fmt.Printf("  %s: %d\n", cp.city, cp.pop)
	}

	// === MAP OPERATIONS AND ALGORITHMS ===
	fmt.Println("\n--- MAP OPERATIONS AND ALGORITHMS ---")

	// Map intersection
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	map2 := map[string]int{"b": 2, "c": 4, "d": 5}

	intersection := make(map[string]int)
	for k, v := range map1 {
		if v2, exists := map2[k]; exists {
			intersection[k] = v + v2 // Sum the values
		}
	}
	fmt.Printf("Intersection: %v\n", intersection)

	// Map union
	union := make(map[string]int)
	for k, v := range map1 {
		union[k] = v
	}
	for k, v := range map2 {
		union[k] = v // map2 values override map1 values
	}
	fmt.Printf("Union: %v\n", union)

	// Map difference
	difference := make(map[string]int)
	for k, v := range map1 {
		if _, exists := map2[k]; !exists {
			difference[k] = v
		}
	}
	fmt.Printf("Difference (map1 - map2): %v\n", difference)

	// === MAP AS SET ===
	fmt.Println("\n--- MAP AS SET ---")

	// Using map as set
	set := make(map[string]bool)

	// Add elements
	set["apple"] = true
	set["banana"] = true
	set["cherry"] = true

	// Check membership
	fmt.Printf("Contains 'apple': %t\n", set["apple"])
	fmt.Printf("Contains 'grape': %t\n", set["grape"])

	// Alternative using empty struct (more memory efficient)
	set2 := make(map[string]struct{})
	set2["apple"] = struct{}{}
	set2["banana"] = struct{}{}

	// Check membership
	_, exists = set2["apple"]
	fmt.Printf("Set2 contains 'apple': %t\n", exists)

	// Set operations
	setA := map[string]struct{}{
		"a": {}, "b": {}, "c": {},
	}
	setB := map[string]struct{}{
		"b": {}, "c": {}, "d": {},
	}

	// Set intersection
	setIntersection := make(map[string]struct{})
	for k := range setA {
		if _, exists := setB[k]; exists {
			setIntersection[k] = struct{}{}
		}
	}

	fmt.Printf("Set A: %v\n", getKeys(setA))
	fmt.Printf("Set B: %v\n", getKeys(setB))
	fmt.Printf("Set intersection: %v\n", getKeys(setIntersection))

	// === MAP PERFORMANCE ===
	fmt.Println("\n--- MAP PERFORMANCE ---")

	// Maps have O(1) average case for access, insert, delete
	// Pre-allocate if approximate size is known
	largeMap := make(map[int]string, 1000)

	// Populate map
	for i := 0; i < 1000; i++ {
		largeMap[i] = fmt.Sprintf("value_%d", i)
	}

	fmt.Printf("Large map size: %d\n", len(largeMap))

	// === MAP GOTCHAS ===
	fmt.Println("\n--- MAP GOTCHAS ---")

	// Gotcha 1: Map iteration order is not guaranteed
	fmt.Println("Map iteration order is not guaranteed:")
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for i := 0; i < 3; i++ {
		fmt.Printf("  Iteration %d: ", i+1)
		for k := range m {
			fmt.Printf("%s ", k)
		}
		fmt.Println()
	}

	// Gotcha 2: Cannot take address of map element
	fmt.Println("Cannot take address of map element")
	// &m["a"] // This would cause compile error

	// Gotcha 3: Map of slices/maps requires initialization
	mapOfSlices := make(map[string][]int)
	mapOfSlices["numbers"] = []int{1, 2, 3} // Must initialize slice
	fmt.Printf("Map of slices: %v\n", mapOfSlices)

	// === CUSTOM KEY TYPES ===
	fmt.Println("\n--- CUSTOM KEY TYPES ---")

	// Struct as key (must be comparable)
	type Point struct {
		X, Y int
	}

	distances := map[Point]float64{
		{0, 0}: 0.0,
		{1, 0}: 1.0,
		{0, 1}: 1.0,
		{1, 1}: 1.414,
	}

	fmt.Println("Distances from origin:")
	for point, distance := range distances {
		fmt.Printf("  %v: %.3f\n", point, distance)
	}

	// === MAP VALIDATION ===
	fmt.Println("\n--- MAP VALIDATION ---")

	// Check if map is nil
	var nilMap map[string]int
	fmt.Printf("Nil map: %v\n", nilMap)
	fmt.Printf("Is nil: %t\n", nilMap == nil)

	// Reading from nil map returns zero value
	fmt.Printf("Reading from nil map: %d\n", nilMap["key"])

	// Writing to nil map causes panic
	// nilMap["key"] = 1 // This would panic

	// Safe map operations
	safeGet := func(m map[string]int, key string) (int, bool) {
		if m == nil {
			return 0, false
		}
		value, exists := m[key]
		return value, exists
	}

	value, exists := safeGet(nilMap, "key")
	fmt.Printf("Safe get: value=%d, exists=%t\n", value, exists)

	// === MAP UTILITIES ===
	fmt.Println("\n--- MAP UTILITIES ---")

	// Get all keys
	getMapKeys := func(m map[string]int) []string {
		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		return keys
	}

	// Get all values
	getMapValues := func(m map[string]int) []int {
		values := make([]int, 0, len(m))
		for _, v := range m {
			values = append(values, v)
		}
		return values
	}

	// Reverse map (swap keys and values)
	reverseMap := func(m map[string]int) map[int]string {
		reversed := make(map[int]string)
		for k, v := range m {
			reversed[v] = k
		}
		return reversed
	}

	sampleMap := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("Keys: %v\n", getMapKeys(sampleMap))
	fmt.Printf("Values: %v\n", getMapValues(sampleMap))
	fmt.Printf("Reversed: %v\n", reverseMap(sampleMap))

	// === MAP BEST PRACTICES ===
	fmt.Println("\n--- MAP BEST PRACTICES ---")
	fmt.Println("1. Initialize maps with make() or map literals")
	fmt.Println("2. Check for key existence with two-value assignment")
	fmt.Println("3. Use maps for O(1) lookups")
	fmt.Println("4. Remember iteration order is not guaranteed")
	fmt.Println("5. Use struct{} for sets to save memory")
	fmt.Println("6. Pre-allocate maps when size is known")
	fmt.Println("7. Maps are reference types")
	fmt.Println("8. Keys must be comparable types")
	fmt.Println("9. Check for nil maps before writing")
	fmt.Println("10. Use delete() to remove elements")
}

// Helper function to get keys from a set
func getKeys(set map[string]struct{}) []string {
	keys := make([]string, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	return keys
}
