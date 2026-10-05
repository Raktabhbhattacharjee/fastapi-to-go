package main

import "fmt"

// CONCEPT 1: PACKAGE-LEVEL CONSTANTS
// Constants declared outside functions are accessible across the package.
// Capitalized name = Exported (Public to other packages)
// Lowercase name   = Unexported (Private to this package)
const BaseAPIURL = "https://api.example.com/v1"

func constExample() {

	// CONCEPT 2: SINGLE CONSTANT DECLARATION
	// Use 'const', the name, '=', and a compile-time literal value.
	// NOTE: You must use '=' (never ':=', which is only for local variables).
	const defaultPort = 8080

	// CONCEPT 3: GROUPED CONSTANT BLOCK
	// Group related constants inside parentheses to avoid repeating 'const'.
	const (
		host    = "localhost"
		timeout = 30 // Time in seconds
	)

	// CONCEPT 4: IMPLICIT VALUE REPETITION
	// If a constant in a block lacks a value, Go automatically copies 
	// the expression from the line directly above it.
	const (
		minConnections = 5
		maxConnections    // Inherits value 5 automatically
	)

	// CONCEPT 5: ENUM-LIKE AUTO-INCREMENTING CONSTANTS (iota)
	// 'iota' is a Go keyword that auto-increments by 1 starting from 0 inside a const block.
	const (
		StatusPending = iota // 0
		StatusActive         // 1
		StatusFailed         // 2
	)

	// Printing values to verify output
	fmt.Println("API URL:", BaseAPIURL)
	fmt.Println("Server:", host, defaultPort)
	fmt.Println("Connections (Min/Max):", minConnections, maxConnections)
	fmt.Println("Status Codes:", StatusPending, StatusActive, StatusFailed)
}