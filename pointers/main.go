package main

import "fmt"

// ============================================================
// STEP 1: PASS BY VALUE (A Copy in Memory)
// ============================================================
// When passing a regular type (int), Go creates a COPY in a new stack frame.
// Modifying 'val' here only changes the copy. The original in main() is untouched!
func ChangeNumByValue(val int) {
	val = 99 // Only changes the local copy!
}

// ============================================================
// STEP 2: PASS BY POINTER / REFERENCE (The Memory Address)
// ============================================================
// 'num' is of type '*int' (pointer to an int).
// It does NOT hold the number 1; it holds the MEMORY ADDRESS of 'num' from main().
// We use '*' to DEREFERENCE: go to that address and modify the actual value!
func ChangeNum(num *int) {
	// *num means: "go to the memory address and change what's inside"
	*num = 5

	fmt.Println("Inside ChangeNum:")
	fmt.Println("  Address received in num:      ", num)
	fmt.Println("  Value at that address (*num): ", *num)
}

func main() {
	// ============================================================
	// STEP 3: WHAT IS A POINTER? (& Operator)
	// ============================================================
	// A regular variable holds data (1).
	// '&' is the "Address-of" operator: it gives the location in RAM.
	num := 1

	fmt.Println("Initial state in main:")
	fmt.Println("  Value of num:        ", num)
	fmt.Println("  Memory address (&num):", &num)
	fmt.Println("--------------------------------------------------")

	// ============================================================
	// STEP 4: CALLING BY VALUE (Original doesn't change)
	// ============================================================
	ChangeNumByValue(num)
	fmt.Println("After ChangeNumByValue(num):", num) // Still 1!
	fmt.Println("--------------------------------------------------")

	// ============================================================
	// STEP 5: CALLING WITH A POINTER (Original changes!)
	// ============================================================
	// ⚠️ IMPORTANT: ChangeNum expects '*int' (an address).
	// So we MUST pass '&num' (address-of num), NOT just 'num'!
	ChangeNum(&num)

	fmt.Println("--------------------------------------------------")
	fmt.Println("After ChangeNum(&num) in main:", num) // Changed to 5!
}