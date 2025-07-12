package main

import "fmt"

// This file covers Go's comment system and documentation practices

// SINGLE-LINE COMMENTS:
// Use // for single-line comments (same as JavaScript)
// Comments can be at the end of a line or on their own line

/* MULTI-LINE COMMENTS:
   Use slash-star for multi-line comments (same as JavaScript)
   These can span multiple lines
   Often used for larger explanations or temporary code removal
*/

// DOCUMENTATION COMMENTS:
// Go has special documentation comments that start with the name being documented
// These are used by tools like 'go doc' and 'godoc'
// Similar to JSDoc in JavaScript, but simpler

// main demonstrates Go's comment and documentation system
func main() {
	fmt.Println("=== GO COMMENTS AND DOCUMENTATION ===")
	
	// This is a single-line comment
	fmt.Println("Learning about Go comments")
	
	/* This is a multi-line comment
	   that spans multiple lines
	   and explains complex concepts */
	
	demonstrateCommentTypes()
	demonstrateDocumentationStyle()
}

// demonstrateCommentTypes shows different ways to use comments in Go
// This is a documentation comment because it starts with the function name
func demonstrateCommentTypes() {
	fmt.Println("\n--- Comment Types ---")
	
	// TODO: This is a common comment convention
	// FIXME: Another common convention for marking issues
	// HACK: For marking temporary solutions
	// NOTE: For important notes
	
	var number int = 42        // End-of-line comment
	var message string = "Hi"  // Another end-of-line comment
	
	fmt.Printf("Number: %d, Message: %s\n", number, message)
	
	/* Sometimes you need to comment out
	   multiple lines of code for debugging
	   fmt.Println("This won't run")
	   fmt.Println("Neither will this")
	*/
}

// demonstrateDocumentationStyle shows Go's documentation conventions
// Documentation comments should be complete sentences
// They should start with the name of the item being documented
func demonstrateDocumentationStyle() {
	fmt.Println("\n--- Documentation Style ---")
	
	// DOCUMENTATION CONVENTIONS:
	// 1. Start with the name of the item
	// 2. Use complete sentences
	// 3. Keep it concise but informative
	// 4. Use proper grammar and punctuation
	
	fmt.Println("Go documentation conventions:")
	fmt.Println("1. Start comments with the item name")
	fmt.Println("2. Use complete sentences")
	fmt.Println("3. Be concise but informative")
	fmt.Println("4. Use proper grammar")
}

// ExampleFunction demonstrates how to write good documentation
// It takes a name parameter and returns a greeting message
// This documentation will be visible when using 'go doc'
func ExampleFunction(name string) string {
	return "Hello, " + name + "!"
}

// COMPARISON WITH JAVASCRIPT:
//
// JavaScript JSDoc:
// /**
//  * Adds two numbers together
//  * @param {number} a - First number
//  * @param {number} b - Second number
//  * @returns {number} Sum of a and b
//  */
// function add(a, b) { return a + b; }
//
// Go Documentation:
// // add takes two integers and returns their sum
// func add(a, b int) int { return a + b }
//
// Go's documentation is simpler and more natural language focused
// JavaScript's JSDoc is more structured with specific tags
// Both can be used to generate documentation automatically

// DOCUMENTATION TOOLS:
// Go: go doc, godoc, pkg.go.dev
// JavaScript: JSDoc, TypeDoc, etc.
//
// Go's tools are built into the language toolchain
// JavaScript tools are usually separate packages
