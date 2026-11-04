package main

import (
	"errors"
	"fmt"
	"time"
)

// =============================================================================
// GO ERROR HANDLING: STEP-BY-STEP TUTORIAL
// =============================================================================
// Philosophy: In Go, ERRORS ARE VALUES, NOT EXCEPTIONS!
// - In Python/JS/Java: You "throw" an exception and hope something "catches" it.
//   It breaks the normal execution flow and jumps up the call stack.
// - In Go: A function simply returns the error as its last return value:
//   (result, err). If err == nil, all good. If err != nil, handle it!
// =============================================================================

// -----------------------------------------------------------------------------
// STEP 1: Sentinel Errors (Standard Predefined Errors)
// -----------------------------------------------------------------------------
// Use errors.New() for constant, predictable errors.
var (
	ErrNotFound       = errors.New("resource not found")
	ErrInsufficientBal = errors.New("insufficient balance")
	ErrNegativeAmount  = errors.New("amount must be greater than zero")
)

// -----------------------------------------------------------------------------
// STEP 2: Returning an Error from a Function
// -----------------------------------------------------------------------------
// Idiom: The error is ALWAYS the last return value!
func withdrawMoney(balance float64, amount float64) (float64, error) {
	if amount <= 0 {
		return balance, ErrNegativeAmount
	}
	if amount > balance {
		return balance, ErrInsufficientBal
	}

	remaining := balance - amount
	return remaining, nil // nil means "No error occurred! Everything succeeded."
}

// -----------------------------------------------------------------------------
// STEP 3: Under the Hood: What is the 'error' Type?
// -----------------------------------------------------------------------------
// The built-in error type is just a simple 1-method interface!
//
//   type error interface {
//       Error() string
//   }
//
// That's it! Any struct that implements Error() string is an error!
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// STEP 4: Custom Error Structs (Rich Error Context)
// -----------------------------------------------------------------------------
// When you need more than just a string (e.g., HTTP status, timestamp, userID)
type APIError struct {
	StatusCode int
	Message    string
	Timestamp  time.Time
}

// Implement the error interface by adding the Error() string method:
func (e *APIError) Error() string {
	return fmt.Sprintf("[%s] HTTP %d: %s", e.Timestamp.Format("15:04:05"), e.StatusCode, e.Message)
}

func fetchUserProfile(userID int) (string, error) {
	if userID <= 0 {
		return "", &APIError{
			StatusCode: 400,
			Message:    "Invalid User ID provided",
			Timestamp:  time.Now(),
		}
	}
	if userID == 404 {
		return "", &APIError{
			StatusCode: 404,
			Message:    "User does not exist in database",
			Timestamp:  time.Now(),
		}
	}
	return "User: Rishi", nil
}

// -----------------------------------------------------------------------------
// STEP 5: Error Wrapping with %w (Adding Context without losing original error)
// -----------------------------------------------------------------------------
// Go 1.13 introduced '%w' in fmt.Errorf:
// It wraps a low-level error inside a higher-level explanation.
func queryDatabase(id int) error {
	// Simulating a raw database error:
	return ErrNotFound
}

func getUserService(id int) error {
	err := queryDatabase(id)
	if err != nil {
		// Wrap the error using %w!
		return fmt.Errorf("userService failed for user %d: %w", id, err)
	}
	return nil
}

// -----------------------------------------------------------------------------
// STEP 6: Inspecting Errors with errors.Is() and errors.As()
// -----------------------------------------------------------------------------
// - errors.Is(err, target): Checks if 'err' contains 'target' anywhere in the wrapped chain!
// - errors.As(err, &targetStruct): Casts/unwraps into a custom error struct.
func demoErrorInspection() {
	fmt.Println("--- STEP 5 & 6: ERROR WRAPPING & INSPECTION ---")

	err := getUserService(42)
	if err != nil {
		fmt.Println("Full Error:", err)

		// Check if the underlying root cause was ErrNotFound:
		if errors.Is(err, ErrNotFound) {
			fmt.Println("-> Root cause identified: The resource was missing in the DB!")
		}
	}

	// Testing custom APIError with errors.As():
	_, apiErr := fetchUserProfile(0)
	if apiErr != nil {
		var myErr *APIError
		// Extract the custom struct from the error:
		if errors.As(apiErr, &myErr) {
			fmt.Println("\nExtracted Custom APIError struct:")
			fmt.Printf("HTTP Status Code: %d\n", myErr.StatusCode)
			fmt.Printf("Error Message: %s\n", myErr.Message)
		}
	}
}

// -----------------------------------------------------------------------------
// STEP 7: Panic & Recover (When catastrophe strikes)
// -----------------------------------------------------------------------------
// - Panic is NOT for normal errors (file missing, bad input, wrong password).
// - Panic is for UNRECOVERABLE BUGS (nil pointer dereference, out-of-bounds slice).
// - Recover stops a panic and keeps the server alive (like a crash protector).
func safeDivision(a, b int) (result int, err error) {
	// Setup a crash-recovery guard using defer:
	defer func() {
		if r := recover(); r != nil {
			// Convert the panic into a normal error!
			err = fmt.Errorf("recovered from catastrophic panic: %v", r)
		}
	}()

	result = a / b // If b == 0, Go panics!
	return result, nil
}

func main() {
	fmt.Println("==================================================")
	fmt.Println("   GOLANG ERROR HANDLING: THE RIGHT WAY")
	fmt.Println("==================================================")

	// --- 1. Basic if err != nil ---
	fmt.Println("\n--- 1. BASIC ERROR CHECKING ---")
	balance, err := withdrawMoney(100.0, 150.0)
	if err != nil {
		fmt.Println("Withdraw failed:", err)
	} else {
		fmt.Printf("Withdraw success! Remaining balance: $%.2f\n", balance)
	}

	// Successful withdraw:
	balance, err = withdrawMoney(100.0, 40.0)
	if err != nil {
		fmt.Println("Withdraw failed:", err)
	} else {
		fmt.Printf("Withdraw success! Remaining balance: $%.2f\n", balance)
	}

	// --- 2. Custom Error Structs ---
	fmt.Println("\n--- 2. CUSTOM ERROR STRUCTS ---")
	profile, err := fetchUserProfile(404)
	if err != nil {
		fmt.Println("Fetch error:", err)
	} else {
		fmt.Println("Profile:", profile)
	}

	// --- 3. Error Wrapping & Inspection ---
	fmt.Println()
	demoErrorInspection()

	// --- 4. Panic & Recover ---
	fmt.Println("\n--- 4. PANIC & RECOVER (Crash Guard) ---")
	res, err := safeDivision(10, 0)
	if err != nil {
		fmt.Println("Caught error cleanly:", err)
	} else {
		fmt.Println("Result:", res)
	}

	fmt.Println("\nProgram continued executing smoothly without crashing!")
}
