package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

// This file covers the string type in Go in complete detail
// Go strings are different from JavaScript strings in several important ways

func main() {
	fmt.Println("=== GO STRING TYPE - COMPLETE GUIDE ===")

	demonstrateStringBasics()
	demonstrateStringLiterals()
	demonstrateStringOperations()
	demonstrateStringComparison()
	demonstrateStringConversions()
	demonstrateUnicodeAndRunes()
	demonstrateStringPackage()
	demonstrateStringBestPractices()
}

func demonstrateStringBasics() {
	fmt.Println("\n--- STRING BASICS ---")

	// String type: string
	// Zero value is empty string ""
	// Strings are immutable in Go (like JavaScript)
	// Strings are UTF-8 encoded by default

	var stringZero string // zero value is ""
	var stringValue string = "Hello, Go!"
	var emptyString string = ""

	fmt.Printf("string - Zero value: '%s', Length: %d, Size: %d bytes\n",
		stringZero, len(stringZero), unsafe.Sizeof(stringZero))
	fmt.Printf("string - Value: '%s', Length: %d\n", stringValue, len(stringValue))
	fmt.Printf("string - Empty: '%s', Length: %d\n", emptyString, len(emptyString))

	// String literal
	var literalString = "Hello, World!" // inferred as string
	fmt.Printf("Literal string type: %T, value: '%s'\n", literalString, literalString)

	// String indexing (read-only)
	if len(stringValue) > 0 {
		fmt.Printf("First character: %c (byte value: %d)\n", stringValue[0], stringValue[0])
		fmt.Printf("Last character: %c (byte value: %d)\n",
			stringValue[len(stringValue)-1], stringValue[len(stringValue)-1])
	}

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: string is a primitive type
	// JavaScript: Strings are also immutable
	// JavaScript: string.length property
	// Go: len(string) function
	// JavaScript: string[index] for character access
	// Go: string[index] for byte access (not character!)
}

func demonstrateStringLiterals() {
	fmt.Println("\n--- STRING LITERALS ---")

	// Interpreted string literals (double quotes)
	var interpreted string = "Hello, World!\nThis is a new line."
	fmt.Printf("Interpreted string: %s\n", interpreted)

	// Raw string literals (backticks)
	var raw string = `Hello, World!
This is a raw string.
It can contain "quotes" and \n (literal backslash n).
It preserves all formatting.`
	fmt.Printf("Raw string: %s\n", raw)

	// Escape sequences in interpreted strings
	var escaped string = "Tab:\tNewline:\nBackslash:\\Quote:\""
	fmt.Printf("Escaped string: %s\n", escaped)

	// Unicode in strings
	var unicode string = "Hello, 世界! 🌍"
	fmt.Printf("Unicode string: %s\n", unicode)

	// String with hex escapes
	var hexString string = "Hello, \\x41\\x42\\x43" // ABC
	fmt.Printf("Hex string: %s\n", hexString)

	// String with octal escapes
	var octalString string = "Hello, \\101\\102\\103" // ABC
	fmt.Printf("Octal string: %s\n", octalString)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: 'single quotes', "double quotes", `template literals`
	// JavaScript: Template literals support ${expression}
	// Go: "interpreted", `raw`
	// Go: No template literal interpolation (use fmt.Sprintf)
}

func demonstrateStringOperations() {
	fmt.Println("\n--- STRING OPERATIONS ---")

	var str1 string = "Hello"
	var str2 string = "World"
	var str3 string = "Hello"

	// String concatenation
	var concatenated string = str1 + ", " + str2 + "!"
	fmt.Printf("Concatenation: %s + %s = %s\n", str1, str2, concatenated)

	// String length
	fmt.Printf("Length of '%s': %d\n", str1, len(str1))
	fmt.Printf("Length of '%s': %d\n", concatenated, len(concatenated))

	// String comparison
	fmt.Printf("'%s' == '%s': %t\n", str1, str2, str1 == str2)
	fmt.Printf("'%s' == '%s': %t\n", str1, str3, str1 == str3)
	fmt.Printf("'%s' != '%s': %t\n", str1, str2, str1 != str2)
	fmt.Printf("'%s' < '%s': %t\n", str1, str2, str1 < str2)
	fmt.Printf("'%s' > '%s': %t\n", str1, str2, str1 > str2)

	// String slicing
	var longString string = "Hello, Go Programming!"
	fmt.Printf("Original: '%s'\n", longString)
	fmt.Printf("Slice [0:5]: '%s'\n", longString[0:5])
	fmt.Printf("Slice [7:]: '%s'\n", longString[7:])
	fmt.Printf("Slice [:5]: '%s'\n", longString[:5])
	fmt.Printf("Slice [7:9]: '%s'\n", longString[7:9])

	// String immutability
	// longString[0] = 'h' // This would cause a compilation error
	// Strings are immutable in Go

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: string1 + string2 (concatenation)
	// JavaScript: string.length (property)
	// JavaScript: string.slice(), string.substring()
	// Go: string1 + string2 (concatenation)
	// Go: len(string) (function)
	// Go: string[start:end] (slicing)
}

func demonstrateStringComparison() {
	fmt.Println("\n--- STRING COMPARISON ---")

	// Lexicographical comparison
	var stringsList []string = []string{"apple", "banana", "cherry", "Apple", "Banana"}

	fmt.Println("String comparison (lexicographical):")
	for i := 0; i < len(stringsList)-1; i++ {
		for j := i + 1; j < len(stringsList); j++ {
			fmt.Printf("'%s' < '%s': %t\n", stringsList[i], stringsList[j], stringsList[i] < stringsList[j])
		}
	}

	// Case-sensitive comparison
	fmt.Printf("'Apple' == 'apple': %t\n", "Apple" == "apple")
	fmt.Printf("'Apple' < 'apple': %t\n", "Apple" < "apple")

	// Case-insensitive comparison using strings package
	fmt.Printf("strings.EqualFold('Apple', 'apple'): %t\n", strings.EqualFold("Apple", "apple"))
	fmt.Printf("strings.ToLower('Apple') == strings.ToLower('apple'): %t\n",
		strings.ToLower("Apple") == strings.ToLower("apple"))

	// Unicode comparison
	var unicode1 string = "café"
	var unicode2 string = "cafe"
	fmt.Printf("'%s' == '%s': %t\n", unicode1, unicode2, unicode1 == unicode2)

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Same string comparison behavior
	// JavaScript: string.localeCompare() for locale-specific comparison
	// Go: strings.Compare() for explicit comparison
}

func demonstrateStringConversions() {
	fmt.Println("\n--- STRING CONVERSIONS ---")

	// String to number conversion
	var numberStr string = "42"
	var floatStr string = "3.14159"
	var invalidStr string = "not a number"

	// String to integer
	if intVal, err := strconv.Atoi(numberStr); err == nil {
		fmt.Printf("String '%s' to int: %d\n", numberStr, intVal)
	}

	if intVal, err := strconv.ParseInt(numberStr, 10, 64); err == nil {
		fmt.Printf("String '%s' to int64: %d\n", numberStr, intVal)
	}

	// String to float
	if floatVal, err := strconv.ParseFloat(floatStr, 64); err == nil {
		fmt.Printf("String '%s' to float64: %f\n", floatStr, floatVal)
	}

	// Error handling
	if _, err := strconv.Atoi(invalidStr); err != nil {
		fmt.Printf("String '%s' to int: ERROR - %v\n", invalidStr, err)
	}

	// Number to string conversion
	var intNum int = 42
	var floatNum float64 = 3.14159

	fmt.Printf("Int %d to string: '%s'\n", intNum, strconv.Itoa(intNum))
	fmt.Printf("Float %f to string: '%s'\n", floatNum, strconv.FormatFloat(floatNum, 'f', 2, 64))

	// Boolean to string
	var boolVal bool = true
	fmt.Printf("Bool %t to string: '%s'\n", boolVal, strconv.FormatBool(boolVal))

	// Using fmt.Sprintf for conversions
	fmt.Printf("Using fmt.Sprintf: '%s'\n", fmt.Sprintf("Number: %d, Float: %.2f, Bool: %t",
		intNum, floatNum, boolVal))

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: parseInt(), parseFloat(), Number()
	// JavaScript: toString(), String()
	// Go: strconv package functions
	// Go: More explicit and error-handled conversions
}

func demonstrateUnicodeAndRunes() {
	fmt.Println("\n--- UNICODE AND RUNES ---")

	// Go strings are UTF-8 encoded
	var unicodeStr string = "Hello, 世界! 🌍"

	fmt.Printf("String: '%s'\n", unicodeStr)
	fmt.Printf("Byte length: %d\n", len(unicodeStr))
	fmt.Printf("Rune count: %d\n", utf8.RuneCountInString(unicodeStr))

	// Iterating over bytes
	fmt.Println("Iterating over bytes:")
	for i := 0; i < len(unicodeStr); i++ {
		fmt.Printf("  Byte %d: %d (%c)\n", i, unicodeStr[i], unicodeStr[i])
	}

	// Iterating over runes
	fmt.Println("Iterating over runes:")
	for i, r := range unicodeStr {
		fmt.Printf("  Rune at byte %d: %d (%c)\n", i, r, r)
	}

	// Working with individual runes
	var runeChar rune = '世'
	fmt.Printf("Rune '世': %d (Unicode code point U+%04X)\n", runeChar, runeChar)

	// Converting between string and runes
	var runes []rune = []rune(unicodeStr)
	fmt.Printf("String to runes: %v\n", runes)

	var backToString string = string(runes)
	fmt.Printf("Runes to string: '%s'\n", backToString)

	// Unicode categories
	fmt.Println("Unicode categories:")
	for _, r := range "Hello, 世界! 123" {
		fmt.Printf("  '%c': Letter=%t, Digit=%t, Space=%t, Punct=%t\n",
			r, unicode.IsLetter(r), unicode.IsDigit(r), unicode.IsSpace(r), unicode.IsPunct(r))
	}

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Strings are UTF-16 encoded
	// JavaScript: string.length gives code units, not characters
	// JavaScript: [...string] or Array.from(string) for proper character iteration
	// Go: Strings are UTF-8 encoded
	// Go: len(string) gives bytes, utf8.RuneCountInString() gives characters
	// Go: range over string gives proper character iteration
}

func demonstrateStringPackage() {
	fmt.Println("\n--- STRING PACKAGE FUNCTIONS ---")

	var text string = "  Hello, Go Programming World!  "

	// String manipulation functions
	fmt.Printf("Original: '%s'\n", text)
	fmt.Printf("ToUpper: '%s'\n", strings.ToUpper(text))
	fmt.Printf("ToLower: '%s'\n", strings.ToLower(text))
	fmt.Printf("Title: '%s'\n", strings.Title(text))
	fmt.Printf("TrimSpace: '%s'\n", strings.TrimSpace(text))
	fmt.Printf("Trim: '%s'\n", strings.Trim(text, " !"))

	// String searching
	fmt.Printf("Contains 'Go': %t\n", strings.Contains(text, "Go"))
	fmt.Printf("HasPrefix 'Hello': %t\n", strings.HasPrefix(strings.TrimSpace(text), "Hello"))
	fmt.Printf("HasSuffix 'World!': %t\n", strings.HasSuffix(strings.TrimSpace(text), "World!"))
	fmt.Printf("Index of 'Go': %d\n", strings.Index(text, "Go"))
	fmt.Printf("LastIndex of 'o': %d\n", strings.LastIndex(text, "o"))
	fmt.Printf("Count 'o': %d\n", strings.Count(text, "o"))

	// String replacement
	fmt.Printf("Replace 'Go' with 'JavaScript': '%s'\n", strings.Replace(text, "Go", "JavaScript", 1))
	fmt.Printf("ReplaceAll 'o' with '0': '%s'\n", strings.ReplaceAll(text, "o", "0"))

	// String splitting and joining
	var words []string = strings.Fields(strings.TrimSpace(text))
	fmt.Printf("Split into words: %v\n", words)

	var joined string = strings.Join(words, " | ")
	fmt.Printf("Join with ' | ': '%s'\n", joined)

	var csvData string = "apple,banana,cherry"
	var fruits []string = strings.Split(csvData, ",")
	fmt.Printf("Split CSV: %v\n", fruits)

	// String repetition
	fmt.Printf("Repeat 'Go' 3 times: '%s'\n", strings.Repeat("Go", 3))

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: string.toUpperCase(), string.toLowerCase()
	// JavaScript: string.trim(), string.trimStart(), string.trimEnd()
	// JavaScript: string.includes(), string.startsWith(), string.endsWith()
	// JavaScript: string.indexOf(), string.lastIndexOf()
	// JavaScript: string.replace(), string.replaceAll()
	// JavaScript: string.split(), array.join()
	// JavaScript: string.repeat()
	// Go: strings package with similar functionality
}

func demonstrateStringBestPractices() {
	fmt.Println("\n--- STRING BEST PRACTICES ---")

	// String concatenation performance
	fmt.Println("String concatenation methods:")

	// Method 1: Simple concatenation (inefficient for many strings)
	var result1 string = "Hello" + " " + "World" + "!"
	fmt.Printf("Simple concatenation: '%s'\n", result1)

	// Method 2: fmt.Sprintf (good for formatting)
	var result2 string = fmt.Sprintf("%s %s%s", "Hello", "World", "!")
	fmt.Printf("fmt.Sprintf: '%s'\n", result2)

	// Method 3: strings.Builder (efficient for many concatenations)
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")
	builder.WriteString("!")
	var result3 string = builder.String()
	fmt.Printf("strings.Builder: '%s'\n", result3)

	// Method 4: strings.Join (efficient for slices)
	var parts []string = []string{"Hello", "World", "!"}
	var result4 string = strings.Join(parts, " ")
	fmt.Printf("strings.Join: '%s'\n", result4)

	// String comparison best practices
	fmt.Println("\nString comparison best practices:")

	// Case-insensitive comparison
	var str1 string = "Hello"
	var str2 string = "HELLO"
	fmt.Printf("Case-insensitive: %t\n", strings.EqualFold(str1, str2))

	// Trimming before comparison
	var str3 string = "  hello  "
	var str4 string = "hello"
	fmt.Printf("Trim before compare: %t\n", strings.TrimSpace(str3) == str4)

	// Working with empty strings
	var emptyStr string = ""
	fmt.Printf("Empty string check: len(str) == 0: %t\n", len(emptyStr) == 0)
	fmt.Printf("Empty string check: str == \"\": %t\n", emptyStr == "")

	// Best practices summary
	fmt.Println("\n--- BEST PRACTICES SUMMARY ---")
	fmt.Println("1. Use strings.Builder for efficient concatenation of many strings")
	fmt.Println("2. Use fmt.Sprintf for formatted string creation")
	fmt.Println("3. Use strings.Join for joining slices of strings")
	fmt.Println("4. Use range loop for proper Unicode character iteration")
	fmt.Println("5. Use strings.EqualFold for case-insensitive comparison")
	fmt.Println("6. Always handle errors from string conversion functions")
	fmt.Println("7. Be aware of the difference between bytes and runes")
	fmt.Println("8. Use raw string literals for strings with many escape sequences")
	fmt.Println("9. Prefer strings package functions over manual string manipulation")
	fmt.Println("10. Consider using constant strings for fixed values")

	// Performance considerations
	fmt.Println("\n--- PERFORMANCE CONSIDERATIONS ---")
	fmt.Println("String concatenation performance (worst to best):")
	fmt.Println("  1. Multiple += operations (creates new string each time)")
	fmt.Println("  2. Single + operation (acceptable for few strings)")
	fmt.Println("  3. fmt.Sprintf (good for formatting)")
	fmt.Println("  4. strings.Join (efficient for slices)")
	fmt.Println("  5. strings.Builder (most efficient for many operations)")
}
