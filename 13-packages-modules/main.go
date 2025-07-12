package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

// === PACKAGES AND MODULES IN GO ===

// Package-level variables
var (
	PackageVersion = "1.0.0"
	PackageName    = "go-basics-packages"
)

// Package-level constants
const (
	DefaultTimeout = 30
	MaxRetries     = 3
)

// Exported function (starts with capital letter)
func PublicFunction() string {
	return "This is a public function"
}

// Unexported function (starts with lowercase letter)
func privateFunction() string {
	return "This is a private function"
}

// Exported struct
type PublicStruct struct {
	Name    string
	Version string
}

// Exported method
func (p *PublicStruct) GetInfo() string {
	return fmt.Sprintf("Package: %s, Version: %s", p.Name, p.Version)
}

// Unexported struct
type privateStruct struct {
	data string
}

// Exported interface
type Formatter interface {
	Format() string
}

// Implementation of interface
func (p *PublicStruct) Format() string {
	return fmt.Sprintf("[%s v%s]", p.Name, p.Version)
}

// Package initialization
func init() {
	fmt.Println("Package initialization: packages-modules")
}

// Demonstrate package structure
func demonstratePackageStructure() {
	fmt.Println("\n--- PACKAGE STRUCTURE ---")

	// Show current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v\n", err)
		return
	}
	fmt.Printf("Working directory: %s\n", wd)

	// Show package structure
	fmt.Println("\nTypical Go package structure:")
	fmt.Println("myproject/")
	fmt.Println("├── go.mod")
	fmt.Println("├── go.sum")
	fmt.Println("├── main.go")
	fmt.Println("├── README.md")
	fmt.Println("├── pkg/")
	fmt.Println("│   ├── util/")
	fmt.Println("│   │   └── helper.go")
	fmt.Println("│   └── models/")
	fmt.Println("│       └── user.go")
	fmt.Println("├── cmd/")
	fmt.Println("│   ├── server/")
	fmt.Println("│   │   └── main.go")
	fmt.Println("│   └── client/")
	fmt.Println("│       └── main.go")
	fmt.Println("├── internal/")
	fmt.Println("│   ├── config/")
	fmt.Println("│   │   └── config.go")
	fmt.Println("│   └── database/")
	fmt.Println("│       └── db.go")
	fmt.Println("└── tests/")
	fmt.Println("    └── integration/")
	fmt.Println("        └── api_test.go")
}

// Demonstrate module operations
func demonstrateModuleOperations() {
	fmt.Println("\n--- MODULE OPERATIONS ---")

	// Show common go mod commands
	fmt.Println("Common go mod commands:")
	fmt.Println("  go mod init <module-name>  - Initialize a new module")
	fmt.Println("  go mod tidy               - Add missing and remove unused modules")
	fmt.Println("  go mod download           - Download modules to local cache")
	fmt.Println("  go mod verify             - Verify dependencies have expected content")
	fmt.Println("  go mod graph              - Print module requirement graph")
	fmt.Println("  go mod why <module>       - Explain why a module is needed")
	fmt.Println("  go mod vendor             - Make vendored copy of dependencies")
	fmt.Println("  go mod edit               - Edit go.mod file")
	fmt.Println("  go list -m all            - List all modules")
	fmt.Println("  go get <module>           - Add or update module")
	fmt.Println("  go get -u <module>        - Update to latest version")
	fmt.Println("  go get <module@version>   - Get specific version")
	fmt.Println("  go clean -modcache        - Clean module cache")
}

// Demonstrate import patterns
func demonstrateImportPatterns() {
	fmt.Println("\n--- IMPORT PATTERNS ---")

	fmt.Println("1. Standard import:")
	fmt.Println("   import \"fmt\"")
	fmt.Println("   import \"os\"")
	fmt.Println()

	fmt.Println("2. Grouped imports:")
	fmt.Println("   import (")
	fmt.Println("       \"fmt\"")
	fmt.Println("       \"os\"")
	fmt.Println("   )")
	fmt.Println()

	fmt.Println("3. Aliased imports:")
	fmt.Println("   import (")
	fmt.Println("       f \"fmt\"")
	fmt.Println("       \"os\"")
	fmt.Println("   )")
	fmt.Println()

	fmt.Println("4. Blank imports (for side effects):")
	fmt.Println("   import (")
	fmt.Println("       _ \"github.com/lib/pq\"")
	fmt.Println("       \"database/sql\"")
	fmt.Println("   )")
	fmt.Println()

	fmt.Println("5. Dot imports (not recommended):")
	fmt.Println("   import (")
	fmt.Println("       . \"fmt\"")
	fmt.Println("   )")
	fmt.Println("   // Now can use: Println instead of fmt.Println")
}

// Demonstrate visibility rules
func demonstrateVisibility() {
	fmt.Println("\n--- VISIBILITY RULES ---")

	// Public (exported) - accessible from other packages
	fmt.Printf("Public function result: %s\n", PublicFunction())

	// Private (unexported) - only accessible within this package
	fmt.Printf("Private function result: %s\n", privateFunction())

	// Public struct
	pub := &PublicStruct{Name: "MyPackage", Version: "1.0.0"}
	fmt.Printf("Public struct: %s\n", pub.GetInfo())

	// Private struct
	priv := &privateStruct{data: "hidden"}
	fmt.Printf("Private struct data: %s\n", priv.data)

	// Constants and variables
	fmt.Printf("Package version: %s\n", PackageVersion)
	fmt.Printf("Default timeout: %d\n", DefaultTimeout)
}

// Demonstrate package documentation
func demonstrateDocumentation() {
	fmt.Println("\n--- PACKAGE DOCUMENTATION ---")

	fmt.Println("Package documentation best practices:")
	fmt.Println("1. Package comment should start with 'Package packagename'")
	fmt.Println("2. Every exported function should have a comment")
	fmt.Println("3. Comments should be complete sentences")
	fmt.Println("4. Use examples in comments")
	fmt.Println("5. Use godoc to generate documentation")
	fmt.Println()

	fmt.Println("Example documentation:")
	fmt.Println("// Package math provides basic mathematical operations.")
	fmt.Println("package math")
	fmt.Println()
	fmt.Println("// Add returns the sum of two integers.")
	fmt.Println("// Example: Add(2, 3) returns 5")
	fmt.Println("func Add(a, b int) int {")
	fmt.Println("    return a + b")
	fmt.Println("}")
}

// Demonstrate testing packages
func demonstrateTesting() {
	fmt.Println("\n--- TESTING PACKAGES ---")

	fmt.Println("Testing file structure:")
	fmt.Println("math/")
	fmt.Println("├── math.go")
	fmt.Println("├── math_test.go")
	fmt.Println("└── example_test.go")
	fmt.Println()

	fmt.Println("Example test file (math_test.go):")
	fmt.Println("package math")
	fmt.Println()
	fmt.Println("import \"testing\"")
	fmt.Println()
	fmt.Println("func TestAdd(t *testing.T) {")
	fmt.Println("    result := Add(2, 3)")
	fmt.Println("    expected := 5")
	fmt.Println("    if result != expected {")
	fmt.Println("        t.Errorf(\"Add(2, 3) = %d; want %d\", result, expected)")
	fmt.Println("    }")
	fmt.Println("}")
	fmt.Println()

	fmt.Println("Benchmark test:")
	fmt.Println("func BenchmarkAdd(b *testing.B) {")
	fmt.Println("    for i := 0; i < b.N; i++ {")
	fmt.Println("        Add(2, 3)")
	fmt.Println("    }")
	fmt.Println("}")
	fmt.Println()

	fmt.Println("Example test:")
	fmt.Println("func ExampleAdd() {")
	fmt.Println("    fmt.Println(Add(2, 3))")
	fmt.Println("    // Output: 5")
	fmt.Println("}")
}

// Demonstrate internal packages
func demonstrateInternalPackages() {
	fmt.Println("\n--- INTERNAL PACKAGES ---")

	fmt.Println("Internal packages are only importable by packages in the same subtree.")
	fmt.Println()
	fmt.Println("Example structure:")
	fmt.Println("myproject/")
	fmt.Println("├── go.mod")
	fmt.Println("├── main.go")
	fmt.Println("├── internal/")
	fmt.Println("│   ├── config/")
	fmt.Println("│   │   └── config.go")
	fmt.Println("│   └── database/")
	fmt.Println("│       └── db.go")
	fmt.Println("└── pkg/")
	fmt.Println("    └── api/")
	fmt.Println("        └── api.go")
	fmt.Println()
	fmt.Println("- main.go can import internal/config")
	fmt.Println("- pkg/api/api.go can import internal/config")
	fmt.Println("- External packages cannot import internal/config")
}

// Demonstrate vendoring
func demonstrateVendoring() {
	fmt.Println("\n--- VENDORING ---")

	fmt.Println("Vendoring stores dependencies in your project:")
	fmt.Println("go mod vendor")
	fmt.Println()
	fmt.Println("This creates a vendor/ directory:")
	fmt.Println("myproject/")
	fmt.Println("├── go.mod")
	fmt.Println("├── go.sum")
	fmt.Println("├── vendor/")
	fmt.Println("│   ├── modules.txt")
	fmt.Println("│   └── github.com/")
	fmt.Println("│       └── some/")
	fmt.Println("│           └── dependency/")
	fmt.Println("└── main.go")
	fmt.Println()
	fmt.Println("Benefits:")
	fmt.Println("- Ensures reproducible builds")
	fmt.Println("- Can work offline")
	fmt.Println("- Faster builds (no network)")
	fmt.Println("- Good for CI/CD")
}

// Demonstrate replace directive
func demonstrateReplace() {
	fmt.Println("\n--- REPLACE DIRECTIVE ---")

	fmt.Println("Replace directive in go.mod:")
	fmt.Println("replace github.com/old/module => github.com/new/module v1.2.3")
	fmt.Println("replace github.com/local/module => ./local/path")
	fmt.Println()
	fmt.Println("Use cases:")
	fmt.Println("- Fork a dependency")
	fmt.Println("- Use local development version")
	fmt.Println("- Fix security vulnerabilities")
	fmt.Println("- Work around broken dependencies")
}

// Demonstrate semantic versioning
func demonstrateSemanticVersioning() {
	fmt.Println("\n--- SEMANTIC VERSIONING ---")

	fmt.Println("Go modules use semantic versioning (semver):")
	fmt.Println("v1.2.3")
	fmt.Println("│ │ │")
	fmt.Println("│ │ └── Patch version (bug fixes)")
	fmt.Println("│ └──── Minor version (new features, backward compatible)")
	fmt.Println("└────── Major version (breaking changes)")
	fmt.Println()
	fmt.Println("Version constraints:")
	fmt.Println("- v1.2.3     (exact version)")
	fmt.Println("- v1.2       (latest patch of v1.2)")
	fmt.Println("- v1         (latest minor of v1)")
	fmt.Println("- latest     (latest version)")
	fmt.Println("- @commit    (specific commit)")
	fmt.Println("- @branch    (specific branch)")
}

// Demonstrate go.mod file
func demonstrateGoMod() {
	fmt.Println("\n--- GO.MOD FILE ---")

	fmt.Println("Example go.mod file:")
	fmt.Println("module github.com/user/project")
	fmt.Println()
	fmt.Println("go 1.21")
	fmt.Println()
	fmt.Println("require (")
	fmt.Println("    github.com/gorilla/mux v1.8.0")
	fmt.Println("    github.com/stretchr/testify v1.8.4")
	fmt.Println(")")
	fmt.Println()
	fmt.Println("require (")
	fmt.Println("    github.com/davecgh/go-spew v1.1.1 // indirect")
	fmt.Println("    github.com/pmezard/go-difflib v1.0.0 // indirect")
	fmt.Println("    gopkg.in/yaml.v3 v3.0.1 // indirect")
	fmt.Println(")")
	fmt.Println()
	fmt.Println("replace github.com/old/module => github.com/new/module v1.0.0")
	fmt.Println()
	fmt.Println("exclude github.com/broken/module v1.0.0")
	fmt.Println()
	fmt.Println("retract v1.0.1 // Contains serious bug")
}

// Demonstrate workspace mode
func demonstrateWorkspace() {
	fmt.Println("\n--- WORKSPACE MODE ---")

	fmt.Println("Go workspaces allow working with multiple modules:")
	fmt.Println("go work init")
	fmt.Println("go work use ./module1 ./module2")
	fmt.Println()
	fmt.Println("Creates go.work file:")
	fmt.Println("go 1.21")
	fmt.Println()
	fmt.Println("use (")
	fmt.Println("    ./module1")
	fmt.Println("    ./module2")
	fmt.Println(")")
	fmt.Println()
	fmt.Println("Benefits:")
	fmt.Println("- Work on multiple related modules")
	fmt.Println("- Test changes across modules")
	fmt.Println("- Easier development workflow")
}

// Demonstrate build tags
func demonstrateBuildTags() {
	fmt.Println("\n--- BUILD TAGS ---")

	fmt.Println("Build tags control compilation:")
	fmt.Println()
	fmt.Println("//go:build linux")
	fmt.Println("// +build linux")
	fmt.Println()
	fmt.Println("package main")
	fmt.Println()
	fmt.Println("Common build tags:")
	fmt.Println("- OS: linux, darwin, windows")
	fmt.Println("- Arch: amd64, 386, arm, arm64")
	fmt.Println("- Custom: debug, production")
	fmt.Println()
	fmt.Println("Building with tags:")
	fmt.Println("go build -tags debug")
	fmt.Println("go build -tags \"linux debug\"")
}

// Demonstrate CGO
func demonstrateCGO() {
	fmt.Println("\n--- CGO ---")

	fmt.Println("CGO allows calling C code from Go:")
	fmt.Println()
	fmt.Println("package main")
	fmt.Println()
	fmt.Println("/*")
	fmt.Println("#include <stdio.h>")
	fmt.Println("void hello() {")
	fmt.Println("    printf(\"Hello from C!\\n\");")
	fmt.Println("}")
	fmt.Println("*/")
	fmt.Println("import \"C\"")
	fmt.Println()
	fmt.Println("func main() {")
	fmt.Println("    C.hello()")
	fmt.Println("}")
	fmt.Println()
	fmt.Println("CGO considerations:")
	fmt.Println("- Increases build time")
	fmt.Println("- Reduces portability")
	fmt.Println("- Complicates deployment")
	fmt.Println("- Can disable with CGO_ENABLED=0")
}

// Demonstrate package analysis
func demonstratePackageAnalysis() {
	fmt.Println("\n--- PACKAGE ANALYSIS ---")

	// Simple example of parsing Go code
	src := `
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.ParseComments)
	if err != nil {
		fmt.Printf("Error parsing: %v\n", err)
		return
	}

	fmt.Printf("Package name: %s\n", f.Name.Name)
	fmt.Printf("Imports: %d\n", len(f.Imports))
	for _, imp := range f.Imports {
		fmt.Printf("  - %s\n", imp.Path.Value)
	}

	// Count declarations
	var funcCount, typeCount int
	ast.Inspect(f, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.FuncDecl:
			funcCount++
		case *ast.TypeSpec:
			typeCount++
		}
		return true
	})

	fmt.Printf("Functions: %d\n", funcCount)
	fmt.Printf("Types: %d\n", typeCount)
}

func main() {
	fmt.Println("=== GO PACKAGES AND MODULES COMPREHENSIVE GUIDE ===")

	// === PACKAGE BASICS ===
	fmt.Println("\n--- PACKAGE BASICS ---")

	fmt.Println("A package is a collection of Go source files in the same directory.")
	fmt.Println("The package name should match the directory name.")
	fmt.Println("Every Go file starts with a package declaration.")
	fmt.Println()

	/*
		JavaScript comparison:
		// JavaScript modules (ES6+)
		// math.js
		export function add(a, b) {
			return a + b;
		}

		export const PI = 3.14159;

		// main.js
		import { add, PI } from './math.js';
		console.log(add(2, 3)); // 5
		console.log(PI); // 3.14159

		// CommonJS (Node.js)
		// math.js
		function add(a, b) {
			return a + b;
		}

		module.exports = { add };

		// main.js
		const { add } = require('./math');
		console.log(add(2, 3)); // 5
	*/

	// === PACKAGE STRUCTURE ===
	demonstratePackageStructure()

	// === VISIBILITY RULES ===
	demonstrateVisibility()

	// === IMPORT PATTERNS ===
	demonstrateImportPatterns()

	// === MODULE SYSTEM ===
	fmt.Println("\n--- MODULE SYSTEM ---")

	fmt.Println("Go modules (introduced in Go 1.11) provide dependency management.")
	fmt.Println("A module is a collection of packages stored in a file tree.")
	fmt.Println("Each module is defined by a go.mod file.")
	fmt.Println()

	// === MODULE OPERATIONS ===
	demonstrateModuleOperations()

	// === GO.MOD FILE ===
	demonstrateGoMod()

	// === SEMANTIC VERSIONING ===
	demonstrateSemanticVersioning()

	// === PACKAGE DOCUMENTATION ===
	demonstrateDocumentation()

	// === TESTING PACKAGES ===
	demonstrateTesting()

	// === INTERNAL PACKAGES ===
	demonstrateInternalPackages()

	// === VENDORING ===
	demonstrateVendoring()

	// === REPLACE DIRECTIVE ===
	demonstrateReplace()

	// === WORKSPACE MODE ===
	demonstrateWorkspace()

	// === BUILD TAGS ===
	demonstrateBuildTags()

	// === CGO ===
	demonstrateCGO()

	// === PACKAGE ANALYSIS ===
	demonstratePackageAnalysis()

	// === PACKAGE NAMING CONVENTIONS ===
	fmt.Println("\n--- PACKAGE NAMING CONVENTIONS ---")

	fmt.Println("Package naming best practices:")
	fmt.Println("1. Use lowercase, single-word names")
	fmt.Println("2. Avoid underscores, hyphens, or mixed case")
	fmt.Println("3. Use descriptive names (http, json, time)")
	fmt.Println("4. Avoid generic names (util, common, misc)")
	fmt.Println("5. Package name should be a noun")
	fmt.Println("6. Don't use plural forms")
	fmt.Println()

	fmt.Println("Good package names:")
	fmt.Println("- http")
	fmt.Println("- json")
	fmt.Println("- time")
	fmt.Println("- strings")
	fmt.Println("- bytes")
	fmt.Println()

	fmt.Println("Avoid:")
	fmt.Println("- util")
	fmt.Println("- common")
	fmt.Println("- base")
	fmt.Println("- helper")
	fmt.Println("- lib")

	// === DEPENDENCY MANAGEMENT ===
	fmt.Println("\n--- DEPENDENCY MANAGEMENT ---")

	fmt.Println("Managing dependencies:")
	fmt.Println("1. Use go get to add dependencies")
	fmt.Println("2. Use go mod tidy to clean up")
	fmt.Println("3. Pin important dependencies to specific versions")
	fmt.Println("4. Use go mod vendor for reproducible builds")
	fmt.Println("5. Review dependencies regularly")
	fmt.Println("6. Use minimal dependencies")
	fmt.Println()

	fmt.Println("Dependency commands:")
	fmt.Println("go get github.com/gorilla/mux")
	fmt.Println("go get github.com/gorilla/mux@v1.8.0")
	fmt.Println("go get -u github.com/gorilla/mux")
	fmt.Println("go mod tidy")
	fmt.Println("go mod download")
	fmt.Println("go mod verify")

	// === CIRCULAR DEPENDENCIES ===
	fmt.Println("\n--- CIRCULAR DEPENDENCIES ---")

	fmt.Println("Go prevents circular dependencies at compile time.")
	fmt.Println("If package A imports B, then B cannot import A.")
	fmt.Println()
	fmt.Println("Solutions for circular dependencies:")
	fmt.Println("1. Extract common code to a third package")
	fmt.Println("2. Use interfaces to break the cycle")
	fmt.Println("3. Reorganize the code structure")
	fmt.Println("4. Use dependency injection")

	// === PACKAGE LAYOUT ===
	fmt.Println("\n--- STANDARD PACKAGE LAYOUT ---")

	fmt.Println("Standard Go project layout:")
	fmt.Println()
	fmt.Println("/")
	fmt.Println("├── cmd/")
	fmt.Println("│   └── myapp/")
	fmt.Println("│       └── main.go")
	fmt.Println("├── internal/")
	fmt.Println("│   ├── app/")
	fmt.Println("│   ├── pkg/")
	fmt.Println("│   └── config/")
	fmt.Println("├── pkg/")
	fmt.Println("│   └── api/")
	fmt.Println("├── api/")
	fmt.Println("├── web/")
	fmt.Println("├── configs/")
	fmt.Println("├── init/")
	fmt.Println("├── scripts/")
	fmt.Println("├── build/")
	fmt.Println("├── deployments/")
	fmt.Println("├── test/")
	fmt.Println("├── docs/")
	fmt.Println("├── tools/")
	fmt.Println("├── examples/")
	fmt.Println("├── third_party/")
	fmt.Println("├── githooks/")
	fmt.Println("├── assets/")
	fmt.Println("├── website/")
	fmt.Println("├── README.md")
	fmt.Println("├── LICENSE")
	fmt.Println("├── Makefile")
	fmt.Println("├── go.mod")
	fmt.Println("└── go.sum")

	// === PACKAGE BEST PRACTICES ===
	fmt.Println("\n--- PACKAGE BEST PRACTICES ---")
	fmt.Println("1. Keep packages focused and cohesive")
	fmt.Println("2. Use meaningful package names")
	fmt.Println("3. Minimize public API surface")
	fmt.Println("4. Document all exported identifiers")
	fmt.Println("5. Use internal packages for implementation details")
	fmt.Println("6. Organize code by feature, not by layer")
	fmt.Println("7. Avoid package-level state when possible")
	fmt.Println("8. Use init() functions sparingly")
	fmt.Println("9. Handle errors appropriately")
	fmt.Println("10. Write tests for all packages")

	// === MODULE BEST PRACTICES ===
	fmt.Println("\n--- MODULE BEST PRACTICES ---")
	fmt.Println("1. Use descriptive module names")
	fmt.Println("2. Follow semantic versioning")
	fmt.Println("3. Pin dependencies to specific versions")
	fmt.Println("4. Use replace directive for development")
	fmt.Println("5. Run go mod tidy regularly")
	fmt.Println("6. Commit go.sum to version control")
	fmt.Println("7. Use go mod verify in CI/CD")
	fmt.Println("8. Keep dependencies minimal")
	fmt.Println("9. Use retract for broken versions")
	fmt.Println("10. Document breaking changes")

	// === COMMON PITFALLS ===
	fmt.Println("\n--- COMMON PITFALLS TO AVOID ---")
	fmt.Println("❌ Circular import dependencies")
	fmt.Println("❌ Using dot imports (except in tests)")
	fmt.Println("❌ Generic package names (util, common)")
	fmt.Println("❌ Large packages with many responsibilities")
	fmt.Println("❌ Exposing internal implementation details")
	fmt.Println("❌ Not documenting exported identifiers")
	fmt.Println("❌ Ignoring go mod tidy warnings")
	fmt.Println("❌ Not pinning critical dependencies")
	fmt.Println("❌ Using replace in production")
	fmt.Println("❌ Committing vendor directory unnecessarily")

	fmt.Println("\nWell-organized packages and modules make Go code maintainable and reusable!")
}
