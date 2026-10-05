package main

import "fmt"

func main() {
	fmt.Println("=========================================")
	fmt.Println("📚 LESSON 02: CONTROL FLOW (IF, SWITCH, LOOPS)")
	fmt.Println("=========================================")

	// ------------------------------------------
	// 1. If / Else with Init Statement
	// ------------------------------------------
	fmt.Println("\n--- 1. If / Else with Init Statement ---")

	// Standard if/else
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	// ⭐️ The Go Idiom: Init statement before condition
	// 'status' is created and scoped ONLY inside this if/else block!
	if status := 200; status >= 200 && status < 300 {
		fmt.Printf("Request succeeded with HTTP %d\n", status)
	} else {
		fmt.Printf("Request failed with HTTP %d\n", status)
	}
	// 'status' does not leak out here!

	// ------------------------------------------
	// 2. Switch Statements (No 'break' needed!)
	// ------------------------------------------
	fmt.Println("\n--- 2. Switch Statements ---")

	// Value switch
	role := "editor"
	switch role {
	case "admin":
		fmt.Println("Access: Full System Admin")
	case "editor", "author": // Multiple values per case
		fmt.Println("Access: Content Creator")
	default:
		fmt.Println("Access: Read-Only")
	}

	// Tagless switch (clean alternative to long if/else chains)
	httpCode := 404
	switch {
	case httpCode == 200:
		fmt.Println("200 OK")
	case httpCode == 404:
		fmt.Println("404 Not Found")
	case httpCode >= 500:
		fmt.Println("5xx Server Error")
	}

	// ------------------------------------------
	// 3. For Loops (The ONLY loop in Go!)
	// ------------------------------------------
	fmt.Println("\n--- 3. For Loops ---")

	// A. Standard 3-part loop (like C)
	fmt.Print("Standard loop: ")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// B. While-style loop (no 'while' keyword in Go)
	n := 1
	fmt.Print("While-style loop: ")
	for n < 10 {
		fmt.Printf("%d ", n)
		n *= 2
	}
	fmt.Println()

	// C. Range loop over a slice (index, value)
	routes := []string{"/api/v1/health", "/api/v1/users", "/api/v1/login"}
	fmt.Println("\nRange loop over routes:")
	for idx, route := range routes {
		fmt.Printf("  Route [%d]: %s\n", idx, route)
	}

	// D. Range loop ignoring index using '_'
	fmt.Print("Routes only: ")
	for _, route := range routes {
		fmt.Printf("%s | ", route)
	}
	fmt.Println()
}
