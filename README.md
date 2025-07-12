# Go Basics - Complete Learning Guide

A comprehensive guide to learning Go programming language from basics to intermediate level, with detailed examples and JavaScript comparisons.

## 📚 Table of Contents


### 🔰 Fundamentals
- [01-fundamentals](./01-fundamentals/)  
  - [program-structure](./01-fundamentals/program-structure/)  
  - [imports](./01-fundamentals/imports/)  
  - [comments](./01-fundamentals/comments/)  
  - [formatting](./01-fundamentals/formatting/)  
  - [names](./01-fundamentals/names/)  
  - [semicolons](./01-fundamentals/semicolons/)  
- [02-data-types](./02-data-types/)  
  - [integers](./02-data-types/integers/)  
  - [floating-point](./02-data-types/floating-point/)  
  - [complex](./02-data-types/complex/)  
  - [boolean](./02-data-types/boolean/)  
  - [string](./02-data-types/string/)  
- [03-variables](./03-variables/)  
  - [declarations](./03-variables/declarations/)  
  - [constants](./03-variables/constants/)  
  - [scope](./03-variables/scope/)  
  - [init-function](./03-variables/init-function/)  
- [04-operators](./04-operators/)  
  - [arithmetic](./04-operators/arithmetic/)  
  - [comparison](./04-operators/comparison/)  
  - [logical](./04-operators/logical/)  
  - [bitwise](./04-operators/bitwise/)  
- [05-control-flow](./05-control-flow/)  
  - [if-else](./05-control-flow/if-else/)  
  - [switch](./05-control-flow/switch/)  
  - [loops](./05-control-flow/loops/)  


### 🔧 Core Concepts
- [06-functions](./06-functions/)  
  - [basic](./06-functions/basic/)  
  - [parameters](./06-functions/parameters/)  
  - [returns](./06-functions/returns/)  
  - [variadic](./06-functions/variadic/)  
  - [closures](./06-functions/closures/)  
  - [recursion](./06-functions/recursion/)  
  - [types](./06-functions/types/)  
  - [defer](./06-functions/defer/)  
- [07-methods-interfaces](./07-methods-interfaces/)  
  - [methods](./07-methods-interfaces/methods/)  
  - [interfaces](./07-methods-interfaces/interfaces/)  
  - [blank-identifier](./07-methods-interfaces/blank-identifier/)  
  - [type-assertions](./07-methods-interfaces/type-assertions/)  
  - [conversions](./07-methods-interfaces/conversions/)  
- [07-collections](./07-collections/)  
  - [arrays](./07-collections/arrays/)  
  - [slices](./07-collections/slices/)  
  - [maps](./07-collections/maps/)  
  - [append](./07-collections/append/)  
  - [printing](./07-collections/printing/)  
  - [two-dimensional](./07-collections/two-dimensional/)  
- [08-pointers](./08-pointers/)  
  - [allocation-new](./08-pointers/allocation-new/)  
  - [allocation-make](./08-pointers/allocation-make/)  
- [09-structs](./09-structs/)  
  - [basic](./09-structs/basic/)  
  - [advanced](./09-structs/advanced/)  
  - [composite-literals](./09-structs/composite-literals/)  
  - [embedding](./09-structs/embedding/)  


### 🚀 Advanced Topics
- [10-methods](./10-methods/) *(legacy, see 07-methods-interfaces)*
- [10-pointers](./10-pointers/) *(legacy, see 08-pointers)*
- [11-error-handling](./11-error-handling/)
- [12-concurrency](./12-concurrency/)
  - [channels-of-channels](./12-concurrency/channels-of-channels/)  
  - [parallelization](./12-concurrency/parallelization/)  
  - [leaky-buffer](./12-concurrency/leaky-buffer/)  
- [13-error-handling](./13-error-handling/) *(legacy, see 11-error-handling)*
- [13-packages-modules](./13-packages-modules/)
- [14-standard-library](./14-standard-library/)

### 🌐 Web Development
- [17-web-server](./17-web-server/)

### 🚀 Projects
- [15-project](./15-project/) - Web API with Database
- [16-project](./16-project/) - Microservices with Concurrency

---
**All topics and subtopics now have README.md files. Legacy/duplicate folders are noted.**
## 🎯 Features

### ✅ Comprehensive Coverage
- **Complete Go Language**: Every major Go concept from basics to advanced
- **Real-world Examples**: Practical code examples you can run immediately
- **JavaScript Comparisons**: Side-by-side comparisons for JavaScript developers
- **Best Practices**: Industry-standard coding practices and patterns
- **Common Pitfalls**: What to avoid and why

### 🔍 Learning Approach
- **Progressive Complexity**: Start simple, build complexity gradually
- **Hands-on Examples**: Every concept includes working code
- **Pattern Recognition**: Learn common Go patterns and idioms
- **Error Handling**: Proper error handling throughout
- **Performance Considerations**: Memory and performance insights

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or higher installed
- Basic programming knowledge (any language)
- Text editor or IDE

### Installation
1. Clone this repository:
```bash
git clone <repository-url>
cd go-basics
```

2. Verify Go installation:
```bash
go version
```

3. Run any example:
```bash
cd 01-basics
go run main.go
```

## � How to Use This Guide

### 🎓 For Beginners
1. Start with [01-basics](./01-basics/) to understand Go syntax
2. Follow the numerical order (01-14) for structured learning
3. Run each example and experiment with the code
4. Read the JavaScript comparisons to understand differences

### 👨‍💻 For Experienced Developers
1. Jump to specific topics of interest
2. Focus on the "Advanced" sections in each module
3. Study the patterns and best practices
4. Use as a reference for Go idioms

### 📚 For JavaScript Developers
1. Pay special attention to the JavaScript comparison sections
2. Note the differences in memory management and concurrency
3. Understand Go's static typing vs JavaScript's dynamic typing
4. Learn Go's error handling vs JavaScript's try-catch

## 🗂️ Project Structure

```
go-basics/
├── 01-basics/
│   └── main.go                 # Hello World and basic syntax
├── 02-data-types/
│   ├── basic/
│   │   └── main.go            # Basic data types
│   └── advanced/
│       └── main.go            # Advanced type concepts
├── 03-variables/
│   └── main.go                # Variable declaration and scope
├── 04-operators/
│   └── main.go                # All operator types
├── 05-control-flow/
│   ├── conditionals/
│   │   └── main.go            # if/else and switch
│   └── loops/
│       └── main.go            # for loops and control
├── 06-functions/
│   ├── basic/
│   │   └── main.go            # Function basics
│   └── advanced/
│       └── main.go            # Advanced function features
├── 07-methods-interfaces/
│   ├── methods/
│   │   └── main.go            # Method declarations
│   └── interfaces/
│       └── main.go            # Interface implementation
├── 08-collections/
│   ├── arrays/
│   │   └── main.go            # Array operations
│   ├── slices/
│   │   └── main.go            # Slice manipulation
│   └── maps/
│       └── main.go            # Map operations
├── 09-structs/
│   ├── basic/
│   │   └── main.go            # Struct basics
│   └── advanced/
│       └── main.go            # Advanced struct patterns
├── 10-pointers/
│   └── main.go                # Pointer usage
├── 11-error-handling/
│   └── main.go                # Error handling patterns
├── 12-concurrency/
│   └── main.go                # Goroutines and channels
├── 13-packages-modules/
│   └── main.go                # Package organization
├── 14-standard-library/
│   └── main.go                # Standard library overview
├── 15-project/
│   └── main.go                # Project 15 overview
├── 16-project/
│   └── main.go                # Project 16 overview
└── README.md                  # This file
```

## 🔧 Running Examples

### Single File
```bash
cd 01-basics
go run main.go
```

### With Build
```bash
cd 01-basics
go build -o basics main.go
./basics
```

### All Examples
```bash
# Run all examples in sequence
for dir in */; do
    if [ -f "$dir/main.go" ]; then
        echo "Running $dir"
        cd "$dir"
        go run main.go
        cd ..
    fi
done
```

## 📝 Key Learning Points

### 🎯 Go vs JavaScript

| Concept | Go | JavaScript |
|---------|----|-----------| 
| **Typing** | Static, strong typing | Dynamic, weak typing |
| **Memory** | Manual pointers, GC | Automatic memory management |
| **Concurrency** | Goroutines, channels | Async/await, promises |
| **Error Handling** | Explicit error returns | Try-catch exceptions |
| **Compilation** | Compiled to binary | Interpreted/JIT |
| **Package System** | Built-in modules | npm/ES6 modules |

### 🚀 Go Strengths
- **Performance**: Compiled binary, efficient runtime
- **Concurrency**: First-class goroutines and channels
- **Simplicity**: Clean syntax, minimal keywords
- **Reliability**: Strong typing, explicit error handling
- **Tooling**: Excellent built-in tools (go fmt, go test, etc.)
- **Deployment**: Single binary deployment

### 🔄 Common Patterns
- **Error Handling**: `if err != nil { return err }`
- **Resource Cleanup**: `defer` statements
- **Concurrency**: Channel-based communication
- **Interface Design**: Small, focused interfaces
- **Package Organization**: Clear separation of concerns

## 🎓 Learning Path

### Week 1: Foundations
- [ ] 01-basics: Syntax and basic concepts
- [ ] 02-data-types: Type system understanding
- [ ] 03-variables: Variable management
- [ ] 04-operators: Expression evaluation
- [ ] 05-control-flow: Program flow control

### Week 2: Core Concepts
- [ ] 06-functions: Function design and usage
- [ ] 07-methods-interfaces: OOP concepts in Go
- [ ] 08-collections: Data structure manipulation
- [ ] 09-structs: Custom type creation

### Week 3: Advanced Topics
- [ ] 10-pointers: Memory management
- [ ] 11-error-handling: Robust error handling
- [ ] 12-concurrency: Concurrent programming
- [ ] 13-packages-modules: Code organization

### Week 4: Practical Application
- [ ] 14-standard-library: Standard library mastery
- [ ] Build a small project using learned concepts
- [ ] Review and practice patterns
- [ ] Explore Go ecosystem

## 🔍 Best Practices Covered

### Code Quality
- ✅ Proper error handling
- ✅ Resource management with defer
- ✅ Clear variable naming
- ✅ Function design principles
- ✅ Interface segregation

### Performance
- ✅ Efficient memory usage
- ✅ Proper goroutine management
- ✅ Channel patterns
- ✅ Avoiding common pitfalls
- ✅ Profiling considerations

### Maintainability
- ✅ Package organization
- ✅ Documentation practices
- ✅ Test-driven development
- ✅ Code formatting
- ✅ Dependency management

## 🎯 Next Steps

After completing this guide, you should:

1. **Build Projects**: Create real applications using Go
2. **Explore Libraries**: Study popular Go libraries and frameworks
3. **Learn Advanced Topics**: Reflect, generics, CGO, build constraints
4. **Read Go Source**: Study Go standard library source code
5. **Join Community**: Participate in Go community and contribute

## 📚 Additional Resources

### Official Documentation
- [Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)

### Books
- "The Go Programming Language" by Alan Donovan and Brian Kernighan
- "Go in Action" by William Kennedy
- "Concurrency in Go" by Katherine Cox-Buday

### Online Resources
- [Go Tour](https://tour.golang.org/)
- [Go Playground](https://play.golang.org/)
- [Go Blog](https://blog.golang.org/)

## 🤝 Contributing

Found an error or want to improve an example? Contributions are welcome!

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License. See LICENSE file for details.

---

**Happy Learning! 🚀**

*Master Go programming with this comprehensive guide designed for developers who want to understand not just the "how" but also the "why" behind Go's design decisions.*

### 02. Data Types
- Basic types (bool, string, numeric)
- Integer types (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64)
- Floating-point types (float32, float64)
- Complex types (complex64, complex128)
- Type declarations and aliases
- Type conversion and casting

### 03. Variables
- Variable declarations (var, :=)
- Zero values
- Constants and iota
- Scope and visibility

### 04. Operators
- Arithmetic operators
- Comparison operators
- Logical operators
- Bitwise operators
- Assignment operators

### 05. Control Structures
- if/else statements
- switch statements
- for loops
- defer statement
- goto statement

### 06. Functions
- Function declarations
- Parameters and return values
- Named return values
- Variadic functions
- Anonymous functions and closures
- Function types

### 07. Collections
- Arrays
- Slices
- Maps
- Strings and runes

### 08. Pointers
- Pointer basics
- Pointer arithmetic (limited)
- new() function
- Address-of and dereference operators

### 09. Structs
- Struct definition
- Struct literals
- Embedded structs
- Struct tags

### 10. Methods
- Method definition
- Value vs pointer receivers
- Method sets

### 11. Interfaces
- Interface definition
- Interface implementation
- Empty interface
- Type assertions
- Type switches

### 12. Packages
- Package declaration
- Imports
- Exported vs unexported identifiers
- Package initialization

### 13. Error Handling
- Error interface
- Error creation
- Error handling patterns
- Custom errors

### 14. Concurrency
- Goroutines
- Channels
- Channel directions
- Select statement
- sync package

### 15. Modules
- Go modules system
- go.mod file
- Dependency management
- Module versioning

### 16. Standard Library
- fmt package
- strings package
- strconv package
- time package
- io package
- os package
- net/http package
- json package

## 🎯 Learning Path
1. Start with fundamentals and work through each folder in order
2. Each file contains complete examples with detailed explanations
3. Run each example to see the output
4. Compare with JavaScript concepts where mentioned
5. Practice by modifying the examples

## 🚀 Getting Started
```bash
# Navigate to any topic folder
cd go-basics/01-fundamentals

# Run any Go file
go run filename.go
```

## 📖 JavaScript Comparisons
Throughout this course, you'll see comparisons with JavaScript to help you understand Go concepts better, since you're already familiar with JavaScript.

Happy Learning! 🎉
