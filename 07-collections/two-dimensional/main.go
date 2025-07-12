package main

import (
	"fmt"
	"math/rand"
	"time"
)

// === GO TWO-DIMENSIONAL ARRAYS/SLICES COMPREHENSIVE GUIDE ===

/*
TWO-DIMENSIONAL PHILOSOPHY:
- Go supports multi-dimensional arrays and slices
- 2D arrays: Fixed size, allocated on stack
- 2D slices: Dynamic size, more flexible
- Common for matrices, grids, tables, game boards

COMPARISON WITH JAVASCRIPT:
// JavaScript - 2D Array
const matrix = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9]
];

// JavaScript - Dynamic 2D Array
const grid = [];
for (let i = 0; i < 3; i++) {
  grid[i] = [];
  for (let j = 0; j < 3; j++) {
    grid[i][j] = i * 3 + j + 1;
  }
}

// Go - 2D Array
var matrix [3][3]int = [3][3]int{
  {1, 2, 3},
  {4, 5, 6},
  {7, 8, 9},
}

// Go - 2D Slice
grid := make([][]int, 3)
for i := 0; i < 3; i++ {
  grid[i] = make([]int, 3)
  for j := 0; j < 3; j++ {
    grid[i][j] = i*3 + j + 1
  }
}
*/

func main() {
	fmt.Println("=== GO TWO-DIMENSIONAL ARRAYS/SLICES COMPREHENSIVE GUIDE ===")

	// === 2D ARRAYS ===
	fmt.Println("\n1. TWO-DIMENSIONAL ARRAYS:")

	// Declaration and initialization
	var matrix [3][3]int

	// Initialize with values
	matrix = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Matrix (2D Array):")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	// Different initialization styles
	identity := [3][3]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}

	fmt.Println("\nIdentity Matrix:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", identity[i][j])
		}
		fmt.Println()
	}

	// === 2D SLICES ===
	fmt.Println("\n2. TWO-DIMENSIONAL SLICES:")

	// Creating 2D slice - Method 1
	grid := make([][]int, 3)
	for i := 0; i < 3; i++ {
		grid[i] = make([]int, 3)
	}

	// Fill with values
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			grid[i][j] = i*3 + j + 1
		}
	}

	fmt.Println("Grid (2D Slice):")
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%2d ", grid[i][j])
		}
		fmt.Println()
	}

	// Creating 2D slice - Method 2 (literal)
	colors := [][]string{
		{"red", "green", "blue"},
		{"yellow", "orange", "purple"},
		{"black", "white", "gray"},
	}

	fmt.Println("\nColors Grid:")
	for i := 0; i < len(colors); i++ {
		for j := 0; j < len(colors[i]); j++ {
			fmt.Printf("%-8s ", colors[i][j])
		}
		fmt.Println()
	}

	// === JAGGED ARRAYS (IRREGULAR 2D SLICES) ===
	fmt.Println("\n3. JAGGED ARRAYS (IRREGULAR 2D SLICES):")

	// Different row lengths
	jagged := [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
		{7, 8, 9, 10},
	}

	fmt.Println("Jagged Array:")
	for i := 0; i < len(jagged); i++ {
		fmt.Printf("Row %d: ", i)
		for j := 0; j < len(jagged[i]); j++ {
			fmt.Printf("%d ", jagged[i][j])
		}
		fmt.Println()
	}

	// === DYNAMIC 2D SLICES ===
	fmt.Println("\n4. DYNAMIC 2D SLICES:")

	// Start with empty slice
	dynamic := [][]int{}

	// Add rows dynamically
	for i := 0; i < 4; i++ {
		row := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			row[j] = j + 1
		}
		dynamic = append(dynamic, row)
	}

	fmt.Println("Dynamic 2D Slice:")
	for i := 0; i < len(dynamic); i++ {
		fmt.Printf("Row %d: %v\n", i, dynamic[i])
	}

	// === PASCAL'S TRIANGLE ===
	fmt.Println("\n5. PASCAL'S TRIANGLE:")

	pascalTriangle := func(n int) [][]int {
		triangle := make([][]int, n)

		for i := 0; i < n; i++ {
			triangle[i] = make([]int, i+1)
			triangle[i][0] = 1
			triangle[i][i] = 1

			for j := 1; j < i; j++ {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
			}
		}

		return triangle
	}

	pascal := pascalTriangle(6)
	fmt.Println("Pascal's Triangle:")
	for i := 0; i < len(pascal); i++ {
		// Print spaces for formatting
		for s := 0; s < len(pascal)-i-1; s++ {
			fmt.Print("  ")
		}

		for j := 0; j < len(pascal[i]); j++ {
			fmt.Printf("%3d ", pascal[i][j])
		}
		fmt.Println()
	}

	// === MATRIX OPERATIONS ===
	fmt.Println("\n6. MATRIX OPERATIONS:")

	// Matrix addition
	matrixAdd := func(a, b [][]int) [][]int {
		rows := len(a)
		cols := len(a[0])
		result := make([][]int, rows)

		for i := 0; i < rows; i++ {
			result[i] = make([]int, cols)
			for j := 0; j < cols; j++ {
				result[i][j] = a[i][j] + b[i][j]
			}
		}

		return result
	}

	// Matrix multiplication
	matrixMultiply := func(a, b [][]int) [][]int {
		aRows := len(a)
		aCols := len(a[0])
		bCols := len(b[0])

		result := make([][]int, aRows)
		for i := 0; i < aRows; i++ {
			result[i] = make([]int, bCols)
			for j := 0; j < bCols; j++ {
				for k := 0; k < aCols; k++ {
					result[i][j] += a[i][k] * b[k][j]
				}
			}
		}

		return result
	}

	// Example matrices
	matA := [][]int{
		{1, 2},
		{3, 4},
	}

	matB := [][]int{
		{5, 6},
		{7, 8},
	}

	fmt.Println("Matrix A:")
	printMatrix(matA)

	fmt.Println("\nMatrix B:")
	printMatrix(matB)

	fmt.Println("\nMatrix A + B:")
	printMatrix(matrixAdd(matA, matB))

	fmt.Println("\nMatrix A * B:")
	printMatrix(matrixMultiply(matA, matB))

	// === GAME BOARD EXAMPLE ===
	fmt.Println("\n7. GAME BOARD EXAMPLE:")

	// Tic-tac-toe board
	board := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}

	// Make some moves
	board[0][0] = "X"
	board[1][1] = "O"
	board[2][2] = "X"

	fmt.Println("Tic-Tac-Toe Board:")
	for i := 0; i < 3; i++ {
		fmt.Print("|")
		for j := 0; j < 3; j++ {
			fmt.Printf(" %s |", board[i][j])
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("-------------")
		}
	}

	// === MEMORY LAYOUT ===
	fmt.Println("\n8. MEMORY LAYOUT:")

	// 2D array - contiguous memory
	var contiguous [2][3]int
	contiguous[0][0] = 1
	contiguous[0][1] = 2
	contiguous[0][2] = 3
	contiguous[1][0] = 4
	contiguous[1][1] = 5
	contiguous[1][2] = 6

	fmt.Println("2D Array (contiguous memory):")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("contiguous[%d][%d] = %d (addr: %p)\n",
				i, j, contiguous[i][j], &contiguous[i][j])
		}
	}

	// 2D slice - potentially non-contiguous
	nonContiguous := make([][]int, 2)
	for i := 0; i < 2; i++ {
		nonContiguous[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			nonContiguous[i][j] = i*3 + j + 1
		}
	}

	fmt.Println("\n2D Slice (potentially non-contiguous):")
	for i := 0; i < 2; i++ {
		fmt.Printf("Row %d slice header: %p\n", i, &nonContiguous[i])
		for j := 0; j < 3; j++ {
			fmt.Printf("  nonContiguous[%d][%d] = %d (addr: %p)\n",
				i, j, nonContiguous[i][j], &nonContiguous[i][j])
		}
	}

	// === PERFORMANCE COMPARISON ===
	fmt.Println("\n9. PERFORMANCE COMPARISON:")

	// Large matrix operations
	size := 100

	// Time 2D array creation and access
	start := time.Now()
	var bigArray [100][100]int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			bigArray[i][j] = i + j
		}
	}
	fmt.Printf("2D Array creation and access: %v\n", time.Since(start))

	// Time 2D slice creation and access
	start = time.Now()
	bigSlice := make([][]int, size)
	for i := 0; i < size; i++ {
		bigSlice[i] = make([]int, size)
		for j := 0; j < size; j++ {
			bigSlice[i][j] = i + j
		}
	}
	fmt.Printf("2D Slice creation and access: %v\n", time.Since(start))

	// === REAL-WORLD EXAMPLES ===
	fmt.Println("\n10. REAL-WORLD EXAMPLES:")

	// Image processing (simplified)
	image := make([][]int, 5)
	for i := 0; i < 5; i++ {
		image[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			image[i][j] = rand.Intn(256) // Random pixel value
		}
	}

	fmt.Println("Image (pixel values):")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%3d ", image[i][j])
		}
		fmt.Println()
	}

	// Data table
	table := [][]string{
		{"Name", "Age", "City"},
		{"Alice", "30", "New York"},
		{"Bob", "25", "Los Angeles"},
		{"Charlie", "35", "Chicago"},
	}

	fmt.Println("\nData Table:")
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			fmt.Printf("%-10s ", table[i][j])
		}
		fmt.Println()
	}

	// Spreadsheet-like operations
	spreadsheet := [][]float64{
		{1.5, 2.3, 3.7},
		{4.1, 5.9, 6.2},
		{7.8, 8.4, 9.6},
	}

	fmt.Println("\nSpreadsheet:")
	for i := 0; i < len(spreadsheet); i++ {
		for j := 0; j < len(spreadsheet[i]); j++ {
			fmt.Printf("%6.2f ", spreadsheet[i][j])
		}
		fmt.Println()
	}

	// Calculate row and column sums
	fmt.Println("\nRow sums:")
	for i := 0; i < len(spreadsheet); i++ {
		sum := 0.0
		for j := 0; j < len(spreadsheet[i]); j++ {
			sum += spreadsheet[i][j]
		}
		fmt.Printf("Row %d: %.2f\n", i, sum)
	}

	fmt.Println("\nColumn sums:")
	for j := 0; j < len(spreadsheet[0]); j++ {
		sum := 0.0
		for i := 0; i < len(spreadsheet); i++ {
			sum += spreadsheet[i][j]
		}
		fmt.Printf("Col %d: %.2f\n", j, sum)
	}

	fmt.Println("\n=== END OF TWO-DIMENSIONAL GUIDE ===")
}

// Helper function to print matrix
func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Println()
	}
}

/*
RUNNING THE PROGRAM:
go run main.go

LEARNING POINTS:

1. 2D ARRAYS:
   - Fixed size: [rows][cols]type
   - Contiguous memory layout
   - Stack allocated (if not too large)
   - Good for small, fixed-size matrices

2. 2D SLICES:
   - Dynamic size: [][]type
   - More flexible than arrays
   - Each row can have different lengths (jagged)
   - Heap allocated

3. INITIALIZATION:
   - Literal: [][]int{{1,2},{3,4}}
   - make(): make([][]int, rows)
   - Loop initialization for dynamic content

4. MEMORY CONSIDERATIONS:
   - 2D arrays: contiguous memory
   - 2D slices: potentially fragmented
   - Consider cache locality for performance

5. COMMON PATTERNS:
   - Matrix operations
   - Game boards
   - Image processing
   - Data tables
   - Grids and maps

6. PERFORMANCE:
   - Arrays generally faster for fixed-size data
   - Slices more flexible but with overhead
   - Consider memory access patterns

This demonstrates comprehensive 2D array and slice usage in Go
for various data structures and algorithms.
*/
