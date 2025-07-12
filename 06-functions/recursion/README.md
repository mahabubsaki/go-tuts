# Recursion in Go

## Overview
Recursion is a programming technique where a function calls itself to solve a problem by breaking it down into smaller, similar subproblems. It's particularly useful for problems with recursive structures like trees, graphs, and mathematical sequences.

## Key Concepts

### 1. Base Case
Every recursive function must have a base case - a condition that stops the recursion to prevent infinite loops.

```go
func factorial(n int) int {
    if n <= 1 {  // Base case
        return 1
    }
    return n * factorial(n-1)  // Recursive case
}
```

### 2. Recursive Case
The recursive case is where the function calls itself with modified parameters, moving toward the base case.

### 3. Stack Overflow
Deep recursion can cause stack overflow. Go has a limited stack size, so consider iterative solutions for deep recursion.

## Common Patterns

### 1. Mathematical Recursion
```go
// Factorial
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

// Fibonacci
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

// Power function
func power(base, exponent int) int {
    if exponent == 0 {
        return 1
    }
    return base * power(base, exponent-1)
}
```

### 2. Array/Slice Recursion
```go
// Binary search
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

// Count occurrences
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
```

### 3. Tree Recursion
```go
type TreeNode struct {
    Value int
    Left  *TreeNode
    Right *TreeNode
}

// Tree traversal
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

// Tree depth
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
```

### 4. String Recursion
```go
// Palindrome check
func isPalindrome(s string) bool {
    if len(s) <= 1 {
        return true
    }
    
    if s[0] != s[len(s)-1] {
        return false
    }
    
    return isPalindrome(s[1 : len(s)-1])
}

// Reverse string
func reverseString(s string) string {
    if len(s) <= 1 {
        return s
    }
    
    return reverseString(s[1:]) + string(s[0])
}
```

## JavaScript Comparison

JavaScript recursion works similarly:

```javascript
// JavaScript factorial
function factorial(n) {
    if (n <= 1) return 1;
    return n * factorial(n - 1);
}

// JavaScript fibonacci
function fibonacci(n) {
    if (n <= 1) return n;
    return fibonacci(n - 1) + fibonacci(n - 2);
}

// JavaScript tree traversal
function inOrderTraversal(root) {
    if (!root) return [];
    
    return [
        ...inOrderTraversal(root.left),
        root.value,
        ...inOrderTraversal(root.right)
    ];
}

// JavaScript array operations
function reverseArray(arr) {
    if (arr.length <= 1) return arr;
    return [...reverseArray(arr.slice(1)), arr[0]];
}

// JavaScript string operations
function isPalindrome(s) {
    if (s.length <= 1) return true;
    if (s[0] !== s[s.length - 1]) return false;
    return isPalindrome(s.slice(1, -1));
}
```

## Advanced Patterns

### 1. Backtracking
```go
func solveNQueens(n int) [][]string {
    var result [][]string
    board := make([][]string, n)
    // Initialize board
    
    var backtrack func(row int)
    backtrack = func(row int) {
        if row == n {
            // Found solution
            return
        }
        
        for col := 0; col < n; col++ {
            if isValid(board, row, col) {
                board[row][col] = "Q"
                backtrack(row + 1)
                board[row][col] = "."  // Backtrack
            }
        }
    }
    
    backtrack(0)
    return result
}
```

### 2. Divide and Conquer
```go
func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    
    mid := len(arr) / 2
    left := mergeSort(arr[:mid])
    right := mergeSort(arr[mid:])
    
    return merge(left, right)
}
```

### 3. Tail Recursion
```go
func factorialTailRecursive(n, accumulator int) int {
    if n <= 1 {
        return accumulator
    }
    return factorialTailRecursive(n-1, n*accumulator)
}
```

### 4. Mutual Recursion
```go
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
```

## Optimization Techniques

### 1. Memoization
```go
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
```

### 2. Tail Call Optimization
Go doesn't automatically optimize tail calls, but you can structure recursive functions to be tail-recursive.

### 3. Iterative Conversion
For simple recursion, consider iterative solutions:

```go
// Recursive
func factorialRecursive(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorialRecursive(n-1)
}

// Iterative
func factorialIterative(n int) int {
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result
}
```

## Best Practices

1. **Always Define Base Case**: Prevent infinite recursion
2. **Progress Toward Base Case**: Ensure recursive calls move toward termination
3. **Use Memoization**: For overlapping subproblems
4. **Consider Stack Limits**: Go has limited stack size
5. **Test Small Inputs**: Start with simple cases
6. **Use Iterative for Simple Cases**: When recursion doesn't add clarity
7. **Tail Recursion**: Structure for potential optimization
8. **Avoid Deep Recursion**: Use iterative solutions for deep recursion
9. **Document Complexity**: Understand time and space complexity
10. **Profile Performance**: Measure and optimize as needed

## Common Pitfalls

1. **Missing Base Case**: Causes infinite recursion
2. **Stack Overflow**: Deep recursion exceeds stack limits
3. **Inefficient Recursion**: Recalculating same values (use memoization)
4. **Wrong Base Case**: Incorrect termination condition
5. **Parameter Modification**: Not progressing toward base case

## Performance Considerations

- **Time Complexity**: Can be exponential without optimization
- **Space Complexity**: Stack space for recursive calls
- **Memoization**: Trades space for time efficiency
- **Tail Recursion**: Can be optimized by compilers (not in Go)

## When to Use Recursion

### Good Cases:
- Tree and graph traversal
- Divide and conquer algorithms
- Mathematical sequences
- Backtracking problems
- Parsing nested structures

### Avoid When:
- Simple iterative solutions exist
- Deep recursion expected
- Performance is critical
- Stack space is limited

## Running the Code

```bash
cd go-basics/06-functions/recursion
go run main.go
```

This demonstrates comprehensive recursion patterns with practical examples, optimization techniques, and best practices for recursive programming in Go.
