package main

import (
	"fmt"
	"strings"
)

// This file covers all logical operators in Go
// Logical operators work with boolean values and return boolean results

func main() {
	fmt.Println("=== GO LOGICAL OPERATORS - COMPLETE GUIDE ===")

	demonstrateBasicLogicalOperators()
	demonstrateShortCircuitEvaluation()
	demonstrateLogicalExpressions()
	demonstrateTruthTables()
	demonstrateLogicalOperatorPrecedence()
	demonstrateLogicalOperatorsWithComparisons()
	demonstrateComplexLogicalExpressions()
	demonstrateLogicalOperatorsBestPractices()
}

func demonstrateBasicLogicalOperators() {
	fmt.Println("\n--- BASIC LOGICAL OPERATORS ---")

	// Basic logical operators: &&, ||, !
	var a bool = true
	var b bool = false

	fmt.Printf("a = %t, b = %t\n", a, b)

	// AND operator (&&)
	fmt.Printf("a && b: %t\n", a && b)
	fmt.Printf("a && true: %t\n", a && true)
	fmt.Printf("b && false: %t\n", b && false)

	// OR operator (||)
	fmt.Printf("a || b: %t\n", a || b)
	fmt.Printf("a || false: %t\n", a || false)
	fmt.Printf("b || true: %t\n", b || true)

	// NOT operator (!)
	fmt.Printf("!a: %t\n", !a)
	fmt.Printf("!b: %t\n", !b)
	fmt.Printf("!!a: %t\n", !!a) // Double negation

	// Combining logical operators
	fmt.Printf("!a && b: %t\n", !a && b)
	fmt.Printf("a && !b: %t\n", a && !b)
	fmt.Printf("!a || !b: %t\n", !a || !b)

	// Logical operators with boolean variables
	var isLoggedIn bool = true
	var hasPermission bool = false
	var isAdmin bool = true

	fmt.Printf("\nUser status: isLoggedIn=%t, hasPermission=%t, isAdmin=%t\n",
		isLoggedIn, hasPermission, isAdmin)

	fmt.Printf("Can access: %t\n", isLoggedIn && (hasPermission || isAdmin))
	fmt.Printf("Must login: %t\n", !isLoggedIn)
	fmt.Printf("Access denied: %t\n", !isLoggedIn || (!hasPermission && !isAdmin))

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Same logical operators (&&, ||, !)
	// JavaScript: Truthy/falsy values (0, "", null, undefined)
	// JavaScript: Logical operators can return non-boolean values
	// Go: Logical operators only work with boolean values
	// Go: No automatic conversion to boolean
}

func demonstrateShortCircuitEvaluation() {
	fmt.Println("\n--- SHORT-CIRCUIT EVALUATION ---")

	// Short-circuit evaluation: second operand is not evaluated if result is determined

	// AND operator short-circuit
	fmt.Printf("Short-circuit AND (&&):\n")
	var x int = 5
	var y int = 0

	// First condition false, second won't be evaluated
	if x < 0 && y/x > 1 { // y/x would cause division by zero, but it's not evaluated
		fmt.Printf("This won't print\n")
	} else {
		fmt.Printf("First condition false, second condition not evaluated\n")
	}

	// Both conditions need to be checked
	if x > 0 && y == 0 {
		fmt.Printf("Both conditions evaluated: x > 0 and y == 0\n")
	}

	// OR operator short-circuit
	fmt.Printf("\nShort-circuit OR (||):\n")
	if x > 0 || y/x > 1 { // y/x would cause division by zero, but it's not evaluated
		fmt.Printf("First condition true, second condition not evaluated\n")
	}

	// Both conditions need to be checked
	if x < 0 || y == 0 {
		fmt.Printf("First condition false, second condition evaluated\n")
	}

	// Practical example: nil pointer check
	var ptr *int = nil

	// Safe way using short-circuit evaluation
	if ptr != nil && *ptr > 0 {
		fmt.Printf("Pointer value is positive\n")
	} else {
		fmt.Printf("Pointer is nil or value is not positive\n")
	}

	// Demonstrating evaluation order
	fmt.Printf("\nDemonstrating evaluation order:\n")

	// Function calls to show evaluation order
	if checkCondition("First") && checkCondition("Second") {
		fmt.Printf("Both conditions true\n")
	}

	fmt.Printf("---\n")

	if checkCondition("First") || checkCondition("Second") {
		fmt.Printf("At least one condition true\n")
	}

	// COMPARISON WITH JAVASCRIPT:
	// JavaScript: Same short-circuit behavior
	// JavaScript: && and || can return the actual values
	// Go: && and || always return boolean values
	// Go: More predictable behavior
}

func demonstrateLogicalExpressions() {
	fmt.Println("\n--- LOGICAL EXPRESSIONS ---")

	// Simple logical expressions
	var isWeekend bool = true
	var isHoliday bool = false
	var isWorkDay bool = !isWeekend && !isHoliday

	fmt.Printf("isWeekend: %t, isHoliday: %t, isWorkDay: %t\n",
		isWeekend, isHoliday, isWorkDay)

	// Complex logical expressions
	var age int = 25
	var hasLicense bool = true
	var hasInsurance bool = true
	var hasCar bool = false

	var canDrive bool = age >= 18 && hasLicense && hasInsurance && hasCar
	var canRentCar bool = age >= 21 && hasLicense && hasInsurance

	fmt.Printf("Age: %d, hasLicense: %t, hasInsurance: %t, hasCar: %t\n",
		age, hasLicense, hasInsurance, hasCar)
	fmt.Printf("Can drive own car: %t\n", canDrive)
	fmt.Printf("Can rent car: %t\n", canRentCar)

	// Logical expressions with string comparisons
	var username string = "admin"
	var password string = "secret123"
	var role string = "administrator"

	var isValidLogin bool = username != "" && password != "" && len(password) >= 8
	var isAdminUser bool = strings.ToLower(role) == "administrator" || strings.ToLower(role) == "admin"
	var hasFullAccess bool = isValidLogin && isAdminUser

	fmt.Printf("Username: '%s', Password length: %d, Role: '%s'\n",
		username, len(password), role)
	fmt.Printf("Valid login: %t\n", isValidLogin)
	fmt.Printf("Admin user: %t\n", isAdminUser)
	fmt.Printf("Full access: %t\n", hasFullAccess)

	// Logical expressions with numeric comparisons
	var score int = 85
	var attendance float64 = 0.9
	var assignments int = 8
	var totalAssignments int = 10

	var passingGrade bool = score >= 60
	var goodAttendance bool = attendance >= 0.8
	var completedAssignments bool = assignments >= totalAssignments*8/10
	var canPass bool = passingGrade && (goodAttendance || completedAssignments)

	fmt.Printf("Score: %d, Attendance: %.1f%%, Assignments: %d/%d\n",
		score, attendance*100, assignments, totalAssignments)
	fmt.Printf("Passing grade: %t\n", passingGrade)
	fmt.Printf("Good attendance: %t\n", goodAttendance)
	fmt.Printf("Completed assignments: %t\n", completedAssignments)
	fmt.Printf("Can pass: %t\n", canPass)
}

func demonstrateTruthTables() {
	fmt.Println("\n--- TRUTH TABLES ---")

	// AND truth table
	fmt.Printf("AND Truth Table (&&):\n")
	fmt.Printf("A     | B     | A && B\n")
	fmt.Printf("------|-------|-------\n")
	fmt.Printf("%-5t | %-5t | %t\n", true, true, true && true)
	fmt.Printf("%-5t | %-5t | %t\n", true, false, true && false)
	fmt.Printf("%-5t | %-5t | %t\n", false, true, false && true)
	fmt.Printf("%-5t | %-5t | %t\n", false, false, false && false)

	// OR truth table
	fmt.Printf("\nOR Truth Table (||):\n")
	fmt.Printf("A     | B     | A || B\n")
	fmt.Printf("------|-------|-------\n")
	fmt.Printf("%-5t | %-5t | %t\n", true, true, true || true)
	fmt.Printf("%-5t | %-5t | %t\n", true, false, true || false)
	fmt.Printf("%-5t | %-5t | %t\n", false, true, false || true)
	fmt.Printf("%-5t | %-5t | %t\n", false, false, false || false)

	// NOT truth table
	fmt.Printf("\nNOT Truth Table (!):\n")
	fmt.Printf("A     | !A\n")
	fmt.Printf("------|----\n")
	fmt.Printf("%-5t | %t\n", true, !true)
	fmt.Printf("%-5t | %t\n", false, !false)

	// XOR (exclusive OR) - not a built-in operator in Go
	fmt.Printf("\nXOR Truth Table (A && !B || !A && B):\n")
	fmt.Printf("A     | B     | A XOR B\n")
	fmt.Printf("------|-------|--------\n")
	fmt.Printf("%-5t | %-5t | %t\n", true, true, (true && !true) || (!true && true))
	fmt.Printf("%-5t | %-5t | %t\n", true, false, (true && !false) || (!true && false))
	fmt.Printf("%-5t | %-5t | %t\n", false, true, (false && !true) || (!false && true))
	fmt.Printf("%-5t | %-5t | %t\n", false, false, (false && !false) || (!false && false))

	// NAND (NOT AND)
	fmt.Printf("\nNAND Truth Table (!(A && B)):\n")
	fmt.Printf("A     | B     | NAND\n")
	fmt.Printf("------|-------|-----\n")
	fmt.Printf("%-5t | %-5t | %t\n", true, true, !(true && true))
	fmt.Printf("%-5t | %-5t | %t\n", true, false, !(true && false))
	fmt.Printf("%-5t | %-5t | %t\n", false, true, !(false && true))
	fmt.Printf("%-5t | %-5t | %t\n", false, false, !(false && false))

	// NOR (NOT OR)
	fmt.Printf("\nNOR Truth Table (!(A || B)):\n")
	fmt.Printf("A     | B     | NOR\n")
	fmt.Printf("------|-------|----\n")
	fmt.Printf("%-5t | %-5t | %t\n", true, true, !(true || true))
	fmt.Printf("%-5t | %-5t | %t\n", true, false, !(true || false))
	fmt.Printf("%-5t | %-5t | %t\n", false, true, !(false || true))
	fmt.Printf("%-5t | %-5t | %t\n", false, false, !(false || false))
}

func demonstrateLogicalOperatorPrecedence() {
	fmt.Println("\n--- LOGICAL OPERATOR PRECEDENCE ---")

	// Operator precedence: ! (highest), &&, || (lowest)
	var a bool = true
	var b bool = false
	var c bool = true

	fmt.Printf("a = %t, b = %t, c = %t\n", a, b, c)

	// Examples of precedence
	var result1 bool = !a && b || c
	var result2 bool = (!a && b) || c
	var result3 bool = !(a && b) || c

	fmt.Printf("!a && b || c: %t\n", result1)
	fmt.Printf("(!a && b) || c: %t\n", result2)
	fmt.Printf("!(a && b) || c: %t\n", result3)

	// More complex precedence examples
	var x bool = true
	var y bool = false
	var z bool = true

	fmt.Printf("\nx = %t, y = %t, z = %t\n", x, y, z)

	// Without parentheses
	var expr1 bool = !x || y && z
	fmt.Printf("!x || y && z: %t\n", expr1)

	// With parentheses to change precedence
	var expr2 bool = (!x || y) && z
	fmt.Printf("(!x || y) && z: %t\n", expr2)

	// Complex expression
	var complex bool = !x && y || !y && z
	fmt.Printf("!x && y || !y && z: %t\n", complex)

	// Step-by-step evaluation
	fmt.Printf("\nStep-by-step evaluation of !x && y || !y && z:\n")
	fmt.Printf("Step 1: !x = %t\n", !x)
	fmt.Printf("Step 2: !y = %t\n", !y)
	fmt.Printf("Step 3: !x && y = %t\n", !x && y)
	fmt.Printf("Step 4: !y && z = %t\n", !y && z)
	fmt.Printf("Step 5: (!x && y) || (!y && z) = %t\n", (!x && y) || (!y && z))

	// Best practice: use parentheses for clarity
	fmt.Printf("\nRecommended: Use parentheses for clarity:\n")
	fmt.Printf("Original: !x && y || !y && z\n")
	fmt.Printf("Clear: (!x && y) || (!y && z)\n")
}

func demonstrateLogicalOperatorsWithComparisons() {
	fmt.Println("\n--- LOGICAL OPERATORS WITH COMPARISONS ---")

	// Combining logical operators with comparison operators
	var age int = 25
	var income int = 50000
	var hasJob bool = true
	var creditScore int = 750

	fmt.Printf("Age: %d, Income: $%d, HasJob: %t, Credit Score: %d\n",
		age, income, hasJob, creditScore)

	// Loan eligibility example
	var ageEligible bool = age >= 18 && age <= 65
	var incomeEligible bool = income >= 30000
	var employmentEligible bool = hasJob
	var creditEligible bool = creditScore >= 650

	fmt.Printf("Age eligible: %t\n", ageEligible)
	fmt.Printf("Income eligible: %t\n", incomeEligible)
	fmt.Printf("Employment eligible: %t\n", employmentEligible)
	fmt.Printf("Credit eligible: %t\n", creditEligible)

	var loanApproved bool = ageEligible && incomeEligible && employmentEligible && creditEligible
	fmt.Printf("Loan approved: %t\n", loanApproved)

	// Grade calculation example
	var mathScore int = 85
	var englishScore int = 78
	var scienceScore int = 92

	fmt.Printf("\nScores - Math: %d, English: %d, Science: %d\n",
		mathScore, englishScore, scienceScore)

	var passingGrade bool = mathScore >= 60 && englishScore >= 60 && scienceScore >= 60
	var honorRoll bool = mathScore >= 90 || englishScore >= 90 || scienceScore >= 90
	var averageAbove80 bool = (mathScore+englishScore+scienceScore)/3 >= 80

	fmt.Printf("Passing all subjects: %t\n", passingGrade)
	fmt.Printf("Honor roll (90+ in any subject): %t\n", honorRoll)
	fmt.Printf("Average above 80: %t\n", averageAbove80)

	// Range checking
	var temperature int = 25
	var humidity int = 60

	fmt.Printf("\nTemperature: %d°C, Humidity: %d%%\n", temperature, humidity)

	var comfortableTemp bool = temperature >= 20 && temperature <= 26
	var comfortableHumidity bool = humidity >= 40 && humidity <= 70
	var comfortableConditions bool = comfortableTemp && comfortableHumidity

	fmt.Printf("Comfortable temperature: %t\n", comfortableTemp)
	fmt.Printf("Comfortable humidity: %t\n", comfortableHumidity)
	fmt.Printf("Comfortable conditions: %t\n", comfortableConditions)

	// String validation example
	var username string = "john_doe"
	var email string = "john@example.com"
	var password string = "SecurePass123"

	fmt.Printf("\nUsername: '%s', Email: '%s', Password length: %d\n",
		username, email, len(password))

	var validUsername bool = len(username) >= 3 && len(username) <= 20
	var validEmail bool = strings.Contains(email, "@") && strings.Contains(email, ".")
	var validPassword bool = len(password) >= 8 && containsUppercase(password) && containsDigit(password)

	fmt.Printf("Valid username: %t\n", validUsername)
	fmt.Printf("Valid email: %t\n", validEmail)
	fmt.Printf("Valid password: %t\n", validPassword)

	var validRegistration bool = validUsername && validEmail && validPassword
	fmt.Printf("Valid registration: %t\n", validRegistration)
}

func demonstrateComplexLogicalExpressions() {
	fmt.Println("\n--- COMPLEX LOGICAL EXPRESSIONS ---")

	// De Morgan's Laws demonstration
	var a bool = true
	var b bool = false

	fmt.Printf("a = %t, b = %t\n", a, b)

	// De Morgan's Law 1: !(A && B) = !A || !B
	var law1Left bool = !(a && b)
	var law1Right bool = !a || !b
	fmt.Printf("De Morgan's Law 1: !(a && b) = %t, !a || !b = %t, Equal: %t\n",
		law1Left, law1Right, law1Left == law1Right)

	// De Morgan's Law 2: !(A || B) = !A && !B
	var law2Left bool = !(a || b)
	var law2Right bool = !a && !b
	fmt.Printf("De Morgan's Law 2: !(a || b) = %t, !a && !b = %t, Equal: %t\n",
		law2Left, law2Right, law2Left == law2Right)

	// Distributive laws
	var c bool = true
	fmt.Printf("\na = %t, b = %t, c = %t\n", a, b, c)

	// Distributive law: A && (B || C) = (A && B) || (A && C)
	var distLeft bool = a && (b || c)
	var distRight bool = (a && b) || (a && c)
	fmt.Printf("Distributive: a && (b || c) = %t, (a && b) || (a && c) = %t, Equal: %t\n",
		distLeft, distRight, distLeft == distRight)

	// Complex business logic example
	var isWeekend bool = false
	var isHoliday bool = true
	var isEmergency bool = false
	var isOnCall bool = true
	var isManager bool = false

	fmt.Printf("\nBusiness logic variables:\n")
	fmt.Printf("Weekend: %t, Holiday: %t, Emergency: %t, OnCall: %t, Manager: %t\n",
		isWeekend, isHoliday, isEmergency, isOnCall, isManager)

	// Complex condition: Should work?
	var shouldWork bool = (!isWeekend && !isHoliday) ||
		(isEmergency && isOnCall) ||
		(isManager && isEmergency)

	fmt.Printf("Should work: %t\n", shouldWork)

	// Breaking down the complex condition
	var normalWorkDay bool = !isWeekend && !isHoliday
	var emergencyAndOnCall bool = isEmergency && isOnCall
	var managerInEmergency bool = isManager && isEmergency

	fmt.Printf("Normal work day: %t\n", normalWorkDay)
	fmt.Printf("Emergency and on call: %t\n", emergencyAndOnCall)
	fmt.Printf("Manager in emergency: %t\n", managerInEmergency)

	// Permission system example
	var isLoggedIn bool = true
	var userRole string = "editor"
	var resourceOwner bool = false
	var hasAdminPrivileges bool = false

	fmt.Printf("\nPermission system:\n")
	fmt.Printf("LoggedIn: %t, Role: %s, Owner: %t, Admin: %t\n",
		isLoggedIn, userRole, resourceOwner, hasAdminPrivileges)

	var canRead bool = isLoggedIn
	var canWrite bool = isLoggedIn &&
		(userRole == "editor" || userRole == "admin" || resourceOwner)
	var canDelete bool = isLoggedIn &&
		(userRole == "admin" || resourceOwner || hasAdminPrivileges)

	fmt.Printf("Can read: %t\n", canRead)
	fmt.Printf("Can write: %t\n", canWrite)
	fmt.Printf("Can delete: %t\n", canDelete)

	// Nested logical expressions
	var condition1 bool = (a && b) || (!a && c)
	var condition2 bool = (a || b) && (!b || c)
	var finalResult bool = condition1 && condition2

	fmt.Printf("\nNested expressions:\n")
	fmt.Printf("Condition 1: (a && b) || (!a && c) = %t\n", condition1)
	fmt.Printf("Condition 2: (a || b) && (!b || c) = %t\n", condition2)
	fmt.Printf("Final result: condition1 && condition2 = %t\n", finalResult)
}

func demonstrateLogicalOperatorsBestPractices() {
	fmt.Println("\n--- LOGICAL OPERATORS BEST PRACTICES ---")

	// Use parentheses for clarity
	var a, b, c bool = true, false, true

	// Unclear
	var unclear bool = !a && b || c

	// Clear
	var clear bool = (!a && b) || c

	fmt.Printf("Unclear: !a && b || c = %t\n", unclear)
	fmt.Printf("Clear: (!a && b) || c = %t\n", clear)

	// Extract complex conditions into variables
	var age int = 25
	var hasLicense bool = true
	var hasInsurance bool = true

	// Not recommended: complex inline condition
	if age >= 18 && hasLicense && hasInsurance {
		fmt.Printf("Eligible to drive\n")
	}

	// Recommended: extract to descriptive variables
	var isAdult bool = age >= 18
	var hasRequiredDocuments bool = hasLicense && hasInsurance
	var canDrive bool = isAdult && hasRequiredDocuments

	if canDrive {
		fmt.Printf("Can drive (with clear conditions)\n")
	}

	// Use short-circuit evaluation for performance and safety
	var ptr *int = nil

	// Safe: uses short-circuit evaluation
	if ptr != nil && *ptr > 0 {
		fmt.Printf("Pointer has positive value\n")
	}

	// Avoid deeply nested conditions
	var isLoggedIn bool = true
	var hasPermission bool = true
	var isActive bool = true

	// Not recommended: deeply nested
	if isLoggedIn {
		if hasPermission {
			if isActive {
				fmt.Printf("Access granted\n")
			}
		}
	}

	// Recommended: combined conditions
	if isLoggedIn && hasPermission && isActive {
		fmt.Printf("Access granted (combined)\n")
	}

	// Use early returns to reduce nesting
	fmt.Printf("Demonstrating early return pattern:\n")
	checkAccess(true, true, true)
	checkAccess(false, true, true)

	// Boolean function naming
	var user struct {
		name   string
		active bool
		admin  bool
	}
	user.name = "John"
	user.active = true
	user.admin = false

	// Good: boolean function names
	if isUserActive(user) && !isUserAdmin(user) {
		fmt.Printf("Regular active user: %s\n", user.name)
	}

	// Avoid comparing boolean values to true/false
	var isEnabled bool = true

	// Not recommended
	if isEnabled == true {
		fmt.Printf("Enabled (redundant comparison)\n")
	}

	// Recommended
	if isEnabled {
		fmt.Printf("Enabled (direct use)\n")
	}

	// For false check
	if !isEnabled {
		fmt.Printf("Not enabled\n")
	}

	// BEST PRACTICES SUMMARY:
	fmt.Println("\nBest Practices Summary:")
	fmt.Println("1. Use parentheses to make complex expressions clear")
	fmt.Println("2. Extract complex conditions into descriptive variables")
	fmt.Println("3. Use short-circuit evaluation for safety and performance")
	fmt.Println("4. Avoid deeply nested conditions")
	fmt.Println("5. Use early returns to reduce nesting")
	fmt.Println("6. Name boolean functions/variables clearly (is*, has*, can*)")
	fmt.Println("7. Don't compare boolean values to true/false")
	fmt.Println("8. Break down complex logical expressions")
	fmt.Println("9. Use De Morgan's laws to simplify negations")
	fmt.Println("10. Consider the logical operator precedence")
}

// Helper functions
func checkCondition(name string) bool {
	fmt.Printf("Checking condition: %s\n", name)
	return name == "First"
}

func checkAccess(loggedIn, hasPermission, isActive bool) {
	fmt.Printf("Checking access: loggedIn=%t, hasPermission=%t, isActive=%t\n",
		loggedIn, hasPermission, isActive)

	if !loggedIn {
		fmt.Printf("Access denied: not logged in\n")
		return
	}

	if !hasPermission {
		fmt.Printf("Access denied: no permission\n")
		return
	}

	if !isActive {
		fmt.Printf("Access denied: account not active\n")
		return
	}

	fmt.Printf("Access granted\n")
}

func isUserActive(user struct {
	name   string
	active bool
	admin  bool
}) bool {
	return user.active
}

func isUserAdmin(user struct {
	name   string
	active bool
	admin  bool
}) bool {
	return user.admin
}

func containsUppercase(s string) bool {
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			return true
		}
	}
	return false
}

func containsDigit(s string) bool {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			return true
		}
	}
	return false
}
