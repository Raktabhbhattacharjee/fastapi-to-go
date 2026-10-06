package main

import "fmt"

func main() {
	fmt.Println("=========================================")
	fmt.Println("📚 LESSON 03: GO DATA STRUCTURES")
	fmt.Println("=========================================")

	// ------------------------------------------
	// 1. Arrays vs Slices
	// ------------------------------------------
	fmt.Println("\n--- 1. Arrays (Fixed Size, Rare in Production) ---")
	// Fixed size: [3]int is a completely different type from [4]int in Go!
	var fixedArr [3]int = [3]int{10, 20, 30}
	fmt.Println("Fixed Array:", fixedArr)
	arrayLengthDemo()

	fmt.Println("\n--- 2. Slices (Dynamic, 90% of Production Go) ---")
	slice() // Calls the slice() function from slices.go
	// Literal slice (no size inside brackets [])
	numbers := []int{1, 2, 3}
	fmt.Printf("Slice: %v (len=%d, cap=%d)\n", numbers, len(numbers), cap(numbers))

	// Appending elements (like Python list.append)
	numbers = append(numbers, 4, 5)
	fmt.Printf("After append: %v (len=%d, cap=%d)\n", numbers, len(numbers), cap(numbers))

	// ⭐️ Production Best Practice: Pre-allocating slice capacity with make()
	// make([]Type, length, capacity)
	// Pre-allocating avoids continuous memory re-allocations when adding many items!
	fastSlice := make([]string, 0, 5)
	fastSlice = append(fastSlice, "GET", "POST", "PUT")
	fmt.Printf("Pre-allocated slice: %v (len=%d, cap=%d)\n", fastSlice, len(fastSlice), cap(fastSlice))

	// Slicing (half-open range [low:high], same as Python)
	subSlice := fastSlice[0:2]
	fmt.Println("Sub-slice [0:2]:", subSlice)

	// ------------------------------------------
	// 2. Maps (Hash Tables / Dictionaries)
	// ------------------------------------------
	fmt.Println("\n--- 3. Maps (Key-Value Stores) ---")

	// Creating a map with make(map[KeyType]ValueType)
	userRoles := make(map[string]string)
	userRoles["rishi"] = "admin"
	userRoles["alex"] = "developer"
	userRoles["sam"] = "viewer"

	// Reading a value
	fmt.Println("Role of rishi:", userRoles["rishi"])

	// Deleting a key
	delete(userRoles, "sam")

	// ⭐️ The "Comma-Ok" Idiom (Check if a key exists)
	// In Python: if "guest" in user_roles:
	// In Go:
	role, exists := userRoles["guest"]
	if !exists {
		fmt.Println("Key 'guest' does NOT exist in map! (Default zero-value was:", role, ")")
	} else {
		fmt.Println("Guest role:", role)
	}

	// Range loop over a map (key, value)
	fmt.Println("\nIterating through map:")
	for username, role := range userRoles {
		fmt.Printf("  %s -> %s\n", username, role)
	}
}
