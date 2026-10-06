package main

import "fmt"

func main() {
	// ============================================================
	// 1. DECLARE & INITIALIZE (Go's dynamic list)
	// ============================================================
	var users []string          // Empty slice (len is 0) - standard way to start
	scores := []int{10, 20, 30} // Slice with initial values

	// ============================================================
	// 2. ADD ITEMS (append returns the updated slice)
	// ============================================================
	users = append(users, "Alice", "Bob")
	scores = append(scores, 40)

	// ============================================================
	// 3. READ & COUNT (len gives length, 0-indexed)
	// ============================================================
	fmt.Printf("Count: %d | First user: %s\n", len(users), users[0])

	// ============================================================
	// 4. LOOP (for range)
	// Use '_' to ignore index when you only need the value
	// ============================================================
	for _, user := range users {
		fmt.Println("User:", user)
	}

	// ============================================================
	// 5. SLICE A PORTION ([start:end], end is excluded)
	// ============================================================
	fmt.Println("Top 2 scores:", scores[:2])

	// ============================================================
	// 6. PRE-ALLOCATE (Performance: length 0, room for 100)
	// ============================================================
	fastList := make([]string, 0, 100)
	fastList = append(fastList, "Ready")
}