package main

import "fmt"

// ==========================================
// 1. PACKAGE-LEVEL SCOPE (Outer Scope)
// ==========================================
// Identifiers declared here exist for the entire package.
// Capitalized = Exported (Public to other packages)
// Lowercase   = Unexported (Private to this package)
const BaseAPIURL = "https://api.example.com/v1"
const a = "package-level-hello"

func main() {
	fmt.Println("=========================================")
	fmt.Println("📚 LESSON 01: VALUES, VARIABLES & SCOPES")
	fmt.Println("=========================================")

	// ------------------------------------------
	// A. Basic Types & Values
	// ------------------------------------------
	fmt.Println("\n--- 1. Basic Types ---")
	fmt.Println("Integer:", 42)
	fmt.Println("Float:  ", 3.14)
	fmt.Println("String: ", "FastAPI to Go")
	fmt.Println("Boolean:", true)

	// ------------------------------------------
	// B. Variable Declarations
	// ------------------------------------------
	fmt.Println("\n--- 2. Variable Declarations ---")

	// 1. Short declaration (:=) - Go infers the type automatically
	name := "Rishi"
	port := 8080
	fmt.Printf("name: %s (type: %T), port: %d (type: %T)\n", name, name, port, port)

	// 2. Explicit type declaration (var)
	var maxConnections int64 = 10000
	fmt.Printf("maxConnections: %d (type: %T)\n", maxConnections, maxConnections)

	// 3. Multiple variable declaration on one line
	var x, y int = 10, 20
	fmt.Println("Multiple vars (x, y):", x, y)

	// 4. Zero values (Safe uninitialized memory)
	var count int
	var title string
	var isReady bool
	fmt.Printf("Zero values -> int: %d, string: %q, bool: %t\n", count, title, isReady)

	// ------------------------------------------
	// C. Variable Shadowing & Scopes
	// ------------------------------------------
	fmt.Println("\n--- 3. Variable Shadowing & Scopes ---")
	fmt.Println("Package-level 'a' before shadow:", a)

	// Declaring local 'a' shadows the package-level 'a' inside main()
	var a = "local-shadowed-a"
	fmt.Println("Local 'a' inside main():", a)

	// Block scope demonstration
	if true {
		inner := "scoped-to-if-block"
		fmt.Println("Inside block:", inner)
		// Reassigning outer variable (notice '=' instead of ':=')
		a = "reassigned-local-a"
	}
	// inner is not accessible here!
	fmt.Println("Local 'a' after block reassignment:", a)

	// ------------------------------------------
	// D. Constants & iota (Enums)
	// ------------------------------------------
	fmt.Println("\n--- 4. Constants & iota ---")
	const defaultTimeout = 30
	fmt.Println("API URL:", BaseAPIURL)
	fmt.Println("Timeout:", defaultTimeout)

	// iota generates auto-incrementing integers (0, 1, 2...)
	const (
		StatusPending = iota // 0
		StatusActive         // 1
		StatusFailed         // 2
	)
	fmt.Printf("Status Enums: Pending=%d, Active=%d, Failed=%d\n", StatusPending, StatusActive, StatusFailed)
}
