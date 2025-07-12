package main

import "fmt"

// This file covers the basic structure of a Go program
// and how packages work in Go

// PACKAGE DECLARATION:
// Every Go file must start with a package declaration
// The 'main' package is special - it creates an executable program
// Other packages create libraries that can be imported

// COMPARISON WITH JAVASCRIPT:
// In JavaScript, you don't need to declare a package
// But in Node.js, you have modules (CommonJS/ES6 modules)
// Go's package system is more explicit and structured

// IMPORTS:
// The import statement brings in external packages
// Similar to JavaScript's import/require, but with some differences:
// - Go imports are resolved at compile time
// - You can only import what you use (unused imports cause compilation errors)
// - Go has a standard library that's much more comprehensive than JavaScript's

// MAIN FUNCTION:
// The main() function is the entry point of the program
// Similar to JavaScript's main execution, but more explicit
// In JavaScript, code runs from top to bottom
// In Go, only the main() function in the main package runs automatically

func main() {
	// This is where your program starts execution
	fmt.Println("Welcome to Go Programming!")
	fmt.Println("This is your first Go program")
	
	// Let's demonstrate some basic concepts
	demonstrateBasicConcepts()
}

// FUNCTION DECLARATION:
// Functions in Go use the 'func' keyword
// Similar to JavaScript's function keyword, but with type information
// JavaScript: function name() { }
// Go: func name() { }

func demonstrateBasicConcepts() {
	// COMMENTS:
	// Single-line comments use // (same as JavaScript)
	/* Multi-line comments use /* */ /* (same as JavaScript) */
	
	fmt.Println("Go is a statically typed language")
	fmt.Println("JavaScript is dynamically typed")
	fmt.Println("Go compiles to machine code")
	fmt.Println("JavaScript runs on a virtual machine (V8, etc.)")
	
	// GO CHARACTERISTICS VS JAVASCRIPT:
	// 1. Go is compiled, JavaScript is interpreted (mostly)
	// 2. Go has static typing, JavaScript has dynamic typing
	// 3. Go has garbage collection (like JavaScript)
	// 4. Go has built-in concurrency, JavaScript has async/await
	// 5. Go has pointers, JavaScript has references
	// 6. Go has structs, JavaScript has objects
	// 7. Go has interfaces, JavaScript has duck typing
}

// PACKAGE VISIBILITY:
// Functions, variables, and types that start with a capital letter are exported (public)
// Functions, variables, and types that start with a lowercase letter are private to the package

// ExportedFunction can be called from other packages
func ExportedFunction() {
	fmt.Println("This function is exported (public)")
}

// privateFunction can only be called within this package
func privateFunction() {
	fmt.Println("This function is private to this package")
}

// COMPILATION AND EXECUTION:
// To run this program:
// go run 01-program-structure.go
//
// To compile this program:
// go build 01-program-structure.go
// This creates an executable file
//
// COMPARISON WITH JAVASCRIPT:
// JavaScript: node filename.js (interpreted)
// Go: go run filename.go (compiled and executed)
// Go: go build filename.go then ./filename (compiled to executable)
