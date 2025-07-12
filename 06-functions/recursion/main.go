package main

import (
	"fmt"
	"strings"
)

// === RECURSION IN GO ===

// 1. Basic recursion - Factorial
func factorial(n int) int {
	// Base case
	if n <= 1 {
		return 1
	}
	// Recursive case
	return n * factorial(n-1)
}

// 2. Fibonacci sequence
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 3. Optimized Fibonacci with memoization
func fibonacciMemo(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}

	if val, exists := memo[n]; exists {
		return val
	}

	result := fibonacciMemo(n-1, memo) + fibonacciMemo(n-2, memo)
	memo[n] = result
	return result
}

// 4. Sum of digits
func sumOfDigits(n int) int {
	if n == 0 {
		return 0
	}
	return n%10 + sumOfDigits(n/10)
}

// 5. Power function
func power(base, exponent int) int {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	return base * power(base, exponent-1)
}

// 6. Optimized power function (fast exponentiation)
func fastPower(base, exponent int) int {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}

	if exponent%2 == 0 {
		half := fastPower(base, exponent/2)
		return half * half
	} else {
		return base * fastPower(base, exponent-1)
	}
}

// 7. Greatest Common Divisor (GCD)
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// 8. Binary search
func binarySearch(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		return binarySearch(arr, target, left, mid-1)
	} else {
		return binarySearch(arr, target, mid+1, right)
	}
}

// 9. Tree traversal structures
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Pre-order traversal (Root -> Left -> Right)
func preOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{root.Value}
	result = append(result, preOrderTraversal(root.Left)...)
	result = append(result, preOrderTraversal(root.Right)...)
	return result
}

// In-order traversal (Left -> Root -> Right)
func inOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	result = append(result, inOrderTraversal(root.Left)...)
	result = append(result, root.Value)
	result = append(result, inOrderTraversal(root.Right)...)
	return result
}

// Post-order traversal (Left -> Right -> Root)
func postOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	result = append(result, postOrderTraversal(root.Left)...)
	result = append(result, postOrderTraversal(root.Right)...)
	result = append(result, root.Value)
	return result
}

// 10. Tree depth
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// 11. Palindrome checker
func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	if s[0] != s[len(s)-1] {
		return false
	}

	return isPalindrome(s[1 : len(s)-1])
}

// 12. Reverse string
func reverseString(s string) string {
	if len(s) <= 1 {
		return s
	}

	return reverseString(s[1:]) + string(s[0])
}

// 13. Count occurrences in array
func countOccurrences(arr []int, target int) int {
	if len(arr) == 0 {
		return 0
	}

	count := 0
	if arr[0] == target {
		count = 1
	}

	return count + countOccurrences(arr[1:], target)
}

// 14. Subset generation
func generateSubsets(arr []int) [][]int {
	if len(arr) == 0 {
		return [][]int{{}}
	}

	first := arr[0]
	rest := arr[1:]

	subsetsWithoutFirst := generateSubsets(rest)
	subsetsWithFirst := make([][]int, len(subsetsWithoutFirst))

	for i, subset := range subsetsWithoutFirst {
		newSubset := make([]int, len(subset)+1)
		newSubset[0] = first
		copy(newSubset[1:], subset)
		subsetsWithFirst[i] = newSubset
	}

	result := make([][]int, 0, len(subsetsWithoutFirst)*2)
	result = append(result, subsetsWithoutFirst...)
	result = append(result, subsetsWithFirst...)

	return result
}

// 15. Permutations
func generatePermutations(arr []int) [][]int {
	if len(arr) <= 1 {
		return [][]int{arr}
	}

	var result [][]int

	for i, num := range arr {
		rest := make([]int, 0, len(arr)-1)
		rest = append(rest, arr[:i]...)
		rest = append(rest, arr[i+1:]...)

		subPermutations := generatePermutations(rest)

		for _, perm := range subPermutations {
			newPerm := make([]int, len(perm)+1)
			newPerm[0] = num
			copy(newPerm[1:], perm)
			result = append(result, newPerm)
		}
	}

	return result
}

// 16. Tower of Hanoi
func towerOfHanoi(n int, source, destination, auxiliary string) {
	if n == 1 {
		fmt.Printf("Move disk 1 from %s to %s\n", source, destination)
		return
	}

	towerOfHanoi(n-1, source, auxiliary, destination)
	fmt.Printf("Move disk %d from %s to %s\n", n, source, destination)
	towerOfHanoi(n-1, auxiliary, destination, source)
}

// 17. Maze solver
type Maze [][]int

func solveMaze(maze Maze, x, y, targetX, targetY int) bool {
	// Check bounds and obstacles
	if x < 0 || x >= len(maze) || y < 0 || y >= len(maze[0]) || maze[x][y] == 1 {
		return false
	}

	// Check if we reached the target
	if x == targetX && y == targetY {
		return true
	}

	// Mark current cell as visited
	maze[x][y] = 1

	// Try all four directions
	if solveMaze(maze, x+1, y, targetX, targetY) ||
		solveMaze(maze, x-1, y, targetX, targetY) ||
		solveMaze(maze, x, y+1, targetX, targetY) ||
		solveMaze(maze, x, y-1, targetX, targetY) {
		return true
	}

	// Backtrack
	maze[x][y] = 0
	return false
}

// 18. N-Queens problem
func solveNQueens(n int) [][]string {
	var result [][]string
	board := make([][]string, n)
	for i := range board {
		board[i] = make([]string, n)
		for j := range board[i] {
			board[i][j] = "."
		}
	}

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			// Found a solution
			solution := make([]string, n)
			for i := range board {
				solution[i] = strings.Join(board[i], "")
			}
			result = append(result, solution)
			return
		}

		for col := 0; col < n; col++ {
			if isValidQueenPlacement(board, row, col, n) {
				board[row][col] = "Q"
				backtrack(row + 1)
				board[row][col] = "."
			}
		}
	}

	backtrack(0)
	return result
}

func isValidQueenPlacement(board [][]string, row, col, n int) bool {
	// Check column
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}

	// Check diagonal (top-left to bottom-right)
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}

	// Check diagonal (top-right to bottom-left)
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}

	return true
}

// 19. Tail recursion examples
func factorialTailRecursive(n, accumulator int) int {
	if n <= 1 {
		return accumulator
	}
	return factorialTailRecursive(n-1, n*accumulator)
}

func fibonacciTailRecursive(n, a, b int) int {
	if n == 0 {
		return a
	}
	return fibonacciTailRecursive(n-1, b, a+b)
}

// 20. Mutual recursion
func isEven(n int) bool {
	if n == 0 {
		return true
	}
	return isOdd(n - 1)
}

func isOdd(n int) bool {
	if n == 0 {
		return false
	}
	return isEven(n - 1)
}

// 21. String manipulation recursion
func removeCharacter(s string, char rune) string {
	if len(s) == 0 {
		return ""
	}

	if rune(s[0]) == char {
		return removeCharacter(s[1:], char)
	}

	return string(s[0]) + removeCharacter(s[1:], char)
}

// 22. List operations
func reverseSlice(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	return append(reverseSlice(arr[1:]), arr[0])
}

func findMax(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	maxRest := findMax(arr[1:])
	if arr[0] > maxRest {
		return arr[0]
	}
	return maxRest
}

// 23. Advanced recursion with complex data structures
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func (ll *LinkedList) Length() int {
	if ll == nil {
		return 0
	}
	return 1 + ll.Next.Length()
}

func (ll *LinkedList) Sum() int {
	if ll == nil {
		return 0
	}
	return ll.Value + ll.Next.Sum()
}

func (ll *LinkedList) Contains(value int) bool {
	if ll == nil {
		return false
	}
	if ll.Value == value {
		return true
	}
	return ll.Next.Contains(value)
}

func main() {
	fmt.Println("=== GO RECURSION COMPREHENSIVE GUIDE ===")

	// === BASIC RECURSION ===
	fmt.Println("\n--- BASIC RECURSION ---")
	fmt.Printf("factorial(5) = %d\n", factorial(5))
	fmt.Printf("fibonacci(10) = %d\n", fibonacci(10))

	// Memoized Fibonacci
	memo := make(map[int]int)
	fmt.Printf("fibonacciMemo(10) = %d\n", fibonacciMemo(10, memo))

	/*
		JavaScript comparison:
		function factorial(n) {
			if (n <= 1) return 1;
			return n * factorial(n - 1);
		}

		function fibonacci(n) {
			if (n <= 1) return n;
			return fibonacci(n - 1) + fibonacci(n - 2);
		}
	*/

	// === MATHEMATICAL RECURSION ===
	fmt.Println("\n--- MATHEMATICAL RECURSION ---")
	fmt.Printf("sumOfDigits(12345) = %d\n", sumOfDigits(12345))
	fmt.Printf("power(2, 5) = %d\n", power(2, 5))
	fmt.Printf("fastPower(2, 10) = %d\n", fastPower(2, 10))
	fmt.Printf("gcd(48, 18) = %d\n", gcd(48, 18))

	// === ARRAY RECURSION ===
	fmt.Println("\n--- ARRAY RECURSION ---")
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 7
	index := binarySearch(arr, target, 0, len(arr)-1)
	fmt.Printf("binarySearch for %d in %v: index %d\n", target, arr, index)

	countArr := []int{1, 2, 3, 2, 2, 4, 2}
	fmt.Printf("countOccurrences of 2 in %v: %d\n", countArr, countOccurrences(countArr, 2))

	// === TREE RECURSION ===
	fmt.Println("\n--- TREE RECURSION ---")
	// Build a binary tree
	root := &TreeNode{Value: 1}
	root.Left = &TreeNode{Value: 2}
	root.Right = &TreeNode{Value: 3}
	root.Left.Left = &TreeNode{Value: 4}
	root.Left.Right = &TreeNode{Value: 5}
	root.Right.Right = &TreeNode{Value: 6}

	fmt.Printf("Pre-order traversal: %v\n", preOrderTraversal(root))
	fmt.Printf("In-order traversal: %v\n", inOrderTraversal(root))
	fmt.Printf("Post-order traversal: %v\n", postOrderTraversal(root))
	fmt.Printf("Max depth: %d\n", maxDepth(root))

	// === STRING RECURSION ===
	fmt.Println("\n--- STRING RECURSION ---")
	fmt.Printf("isPalindrome('racecar'): %t\n", isPalindrome("racecar"))
	fmt.Printf("isPalindrome('hello'): %t\n", isPalindrome("hello"))
	fmt.Printf("reverseString('hello'): %s\n", reverseString("hello"))
	fmt.Printf("removeCharacter('hello world', 'l'): %s\n", removeCharacter("hello world", 'l'))

	// === COMBINATORIAL RECURSION ===
	fmt.Println("\n--- COMBINATORIAL RECURSION ---")
	nums := []int{1, 2, 3}
	subsets := generateSubsets(nums)
	fmt.Printf("Subsets of %v:\n", nums)
	for _, subset := range subsets {
		fmt.Printf("  %v\n", subset)
	}

	perms := generatePermutations([]int{1, 2, 3})
	fmt.Printf("Permutations of [1, 2, 3]: (showing first 3)\n")
	for i, perm := range perms {
		if i >= 3 {
			break
		}
		fmt.Printf("  %v\n", perm)
	}

	// === TOWER OF HANOI ===
	fmt.Println("\n--- TOWER OF HANOI ---")
	fmt.Println("Tower of Hanoi with 3 disks:")
	towerOfHanoi(3, "A", "C", "B")

	// === MAZE SOLVING ===
	fmt.Println("\n--- MAZE SOLVING ---")
	maze := Maze{
		{0, 0, 0, 1},
		{1, 1, 0, 1},
		{0, 0, 0, 0},
		{0, 1, 1, 0},
	}

	// Create a copy for solving
	mazeCopy := make(Maze, len(maze))
	for i := range maze {
		mazeCopy[i] = make([]int, len(maze[i]))
		copy(mazeCopy[i], maze[i])
	}

	canSolve := solveMaze(mazeCopy, 0, 0, 2, 3)
	fmt.Printf("Can solve maze from (0,0) to (2,3): %t\n", canSolve)

	// === N-QUEENS ===
	fmt.Println("\n--- N-QUEENS ---")
	solutions := solveNQueens(4)
	fmt.Printf("N-Queens solutions for 4x4 board: %d solutions\n", len(solutions))
	if len(solutions) > 0 {
		fmt.Println("First solution:")
		for _, row := range solutions[0] {
			fmt.Printf("  %s\n", row)
		}
	}

	// === TAIL RECURSION ===
	fmt.Println("\n--- TAIL RECURSION ---")
	fmt.Printf("factorialTailRecursive(5) = %d\n", factorialTailRecursive(5, 1))
	fmt.Printf("fibonacciTailRecursive(10) = %d\n", fibonacciTailRecursive(10, 0, 1))

	// === MUTUAL RECURSION ===
	fmt.Println("\n--- MUTUAL RECURSION ---")
	fmt.Printf("isEven(4): %t\n", isEven(4))
	fmt.Printf("isOdd(4): %t\n", isOdd(4))
	fmt.Printf("isEven(7): %t\n", isEven(7))
	fmt.Printf("isOdd(7): %t\n", isOdd(7))

	// === LIST OPERATIONS ===
	fmt.Println("\n--- LIST OPERATIONS ---")
	numbers := []int{1, 5, 3, 9, 2, 8}
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Reversed: %v\n", reverseSlice(numbers))
	fmt.Printf("Max: %d\n", findMax(numbers))

	// === LINKED LIST RECURSION ===
	fmt.Println("\n--- LINKED LIST RECURSION ---")
	ll := &LinkedList{Value: 1}
	ll.Next = &LinkedList{Value: 2}
	ll.Next.Next = &LinkedList{Value: 3}
	ll.Next.Next.Next = &LinkedList{Value: 4}

	fmt.Printf("Linked list length: %d\n", ll.Length())
	fmt.Printf("Linked list sum: %d\n", ll.Sum())
	fmt.Printf("Contains 3: %t\n", ll.Contains(3))
	fmt.Printf("Contains 5: %t\n", ll.Contains(5))

	// === RECURSION BEST PRACTICES ===
	fmt.Println("\n--- RECURSION BEST PRACTICES ---")
	fmt.Println("1. Always define a base case to prevent infinite recursion")
	fmt.Println("2. Ensure the recursive case moves toward the base case")
	fmt.Println("3. Use memoization for optimization when needed")
	fmt.Println("4. Consider tail recursion for better performance")
	fmt.Println("5. Be aware of stack overflow for deep recursion")
	fmt.Println("6. Use iterative solutions for simple cases")
	fmt.Println("7. Recursion is ideal for tree and graph traversal")
	fmt.Println("8. Good for divide-and-conquer algorithms")
	fmt.Println("9. Useful for backtracking problems")
	fmt.Println("10. Test with small inputs first")

	// === PERFORMANCE COMPARISON ===
	fmt.Println("\n--- PERFORMANCE COMPARISON ---")
	fmt.Printf("Regular factorial(10): %d\n", factorial(10))
	fmt.Printf("Tail recursive factorial(10): %d\n", factorialTailRecursive(10, 1))

	// Note: For large values, tail recursion and memoization show significant benefits
	fmt.Println("\nNote: For large values, use memoization or iterative approaches")
	fmt.Printf("Example: fibonacci(40) would be very slow without memoization\n")
}
