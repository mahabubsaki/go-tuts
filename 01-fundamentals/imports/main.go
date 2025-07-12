package main

import (
	"fmt"
	"math"
	"strings"
)

// This file demonstrates different ways to import packages in Go
// and how Go's import system works

// IMPORT STYLES:
// 1. Single import: import "fmt"
// 2. Multiple imports with parentheses (shown above)
// 3. Named imports
// 4. Blank imports
// 5. Dot imports (not recommended)

func main2() {
	fmt.Println("=== GO IMPORT SYSTEM ===")

	// STANDARD LIBRARY IMPORTS:
	// Go has a rich standard library
	// No need to install external packages for basic functionality
	// JavaScript requires npm packages for many things that Go provides built-in

	demonstrateStandardLibrary()
	demonstrateImportConcepts()
}

func demonstrateStandardLibrary() {
	fmt.Println("\n--- Standard Library Examples ---")

	// fmt package - for formatted I/O
	fmt.Println("fmt.Println: Hello, World!")
	fmt.Printf("fmt.Printf: %s %d\n", "Number:", 42)

	// math package - mathematical functions
	fmt.Printf("math.Pi: %.2f\n", math.Pi)
	fmt.Printf("math.Sqrt(16): %.2f\n", math.Sqrt(16))
	fmt.Printf("math.Max(10, 20): %.2f\n", math.Max(10, 20))

	// strings package - string manipulation
	text := "Hello, Go Programming!"
	fmt.Printf("strings.ToUpper: %s\n", strings.ToUpper(text))
	fmt.Printf("strings.Contains: %t\n", strings.Contains(text, "Go"))
	fmt.Printf("strings.Replace: %s\n", strings.Replace(text, "Go", "JavaScript", 1))

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Math.PI, Math.sqrt(), Math.max()
	// Go: math.Pi, math.Sqrt(), math.Max()
	//
	// JavaScript: text.toUpperCase(), text.includes(), text.replace()
	// Go: strings.ToUpper(text), strings.Contains(text, "Go"), strings.Replace(...)
	//
	// Notice: In Go, string methods are functions that take strings as parameters
	// In JavaScript, strings have methods attached to them
}

func demonstrateImportConcepts() {
	fmt.Println("\n--- Import Concepts ---")

	// PACKAGE NAMING:
	// Package names are usually lowercase
	// Package names should be short and descriptive
	// Package path vs package name can be different

	fmt.Println("Package 'fmt' provides formatted I/O")
	fmt.Println("Package 'math' provides mathematical functions")
	fmt.Println("Package 'strings' provides string utilities")

	// UNUSED IMPORTS:
	// Go doesn't allow unused imports (compilation error)
	// JavaScript allows unused imports (just increases bundle size)
	// This keeps Go programs clean and efficient

	// IMPORT PATH:
	// Standard library packages have short names
	// Third-party packages use URLs: github.com/user/repo
	// Local packages use relative paths

	fmt.Println("\nGo's import system ensures:")
	fmt.Println("1. No unused code")
	fmt.Println("2. Clear dependencies")
	fmt.Println("3. Fast compilation")
	fmt.Println("4. Reproducible builds")
}

// EXAMPLE OF NAMED IMPORT:
// import (
//     f "fmt"              // Now use f.Println instead of fmt.Println
//     m "math"             // Now use m.Pi instead of math.Pi
// )

// EXAMPLE OF BLANK IMPORT:
// import (
//     _ "database/sql"     // Import for side effects only
// )

// EXAMPLE OF DOT IMPORT (NOT RECOMMENDED):
// import (
//     . "fmt"              // Now use Println directly without fmt prefix
// )

// COMPARISON WITH JAVASCRIPT MODULES:
// JavaScript ES6:
// import { useState, useEffect } from 'react';
// import * as React from 'react';
// import React, { Component } from 'react';
//
// JavaScript CommonJS:
// const fs = require('fs');
// const { join } = require('path');
//
// Go:
// import "fmt"
// import f "fmt"
// import _ "fmt"
// import . "fmt"

// MODULE SYSTEM:
// Go uses go.mod for dependency management
// JavaScript uses package.json
// Both support semantic versioning
// Go has built-in module management (go mod)
// JavaScript uses npm/yarn for package management
