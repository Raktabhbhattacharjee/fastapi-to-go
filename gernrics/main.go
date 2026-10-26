package main

import "fmt"

// ============================================================
// 1. ENUMS IN GO (Custom Types + Constants)
// ============================================================
// Go has NO 'enum' keyword. We use custom types with const blocks.

// A. String Enum (Used in 90% of REST APIs / JSON):
type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusConfirmed OrderStatus = "confirmed"
	StatusDelivered OrderStatus = "delivered"
)

// B. Integer Enum with 'iota' (Auto-increments: 0, 1, 2...):
type Role int

const (
	RoleUser       Role = iota // 0
	RoleModerator              // 1
	RoleAdmin                  // 2
)

// ============================================================
// 2. GENERICS IN GO (Type Parameters [T any])
// ============================================================

// Generic Function: Works with slices of ANY type (int, string, bool, etc.)
// [T any] means: T can be any type
func PrintSlice[T any](items []T) {
	for _, item := range items {
		fmt.Printf("%v ", item)
	}
	fmt.Println()
}

// Generic Function with constraint: [T comparable] allows using ==
func Contains[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

// ============================================================
// 3. GENERIC STRUCT (The FastAPI Response Envelope Pattern)
// ============================================================
// Standard JSON wrapper where Data can hold ANY type:
type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type User struct {
	Name string
}

func main() {
	// ==========================================
	// Testing Enums:
	// ==========================================
	currentStatus := StatusConfirmed
	fmt.Println("Order Status Enum: ", currentStatus)
	fmt.Println("Role Enum (Admin): ", RoleAdmin) // Prints 2 (from iota)

	// ==========================================
	// Testing Generic Functions:
	// ==========================================
	fmt.Print("Ints:    ")
	PrintSlice([]int{10, 20, 30})

	fmt.Print("Strings: ")
	PrintSlice([]string{"Alice", "Bob", "Charlie"})

	// Testing Contains (comparable):
	names := []string{"apple", "banana", "orange"}
	fmt.Println("Contains 'banana':", Contains(names, "banana")) // true
	fmt.Println("Contains 'grape': ", Contains(names, "grape"))  // false

	// ==========================================
	// Testing Generic Struct (API Response):
	// ==========================================
	// 1. Wrapping a single User:
	userRes := APIResponse[User]{
		Success: true,
		Message: "User profile loaded",
		Data:    User{Name: "Rishi"},
	}
	fmt.Printf("API Response 1: %+v\n", userRes)

	// 2. Wrapping a slice of ints:
	intListRes := APIResponse[[]int]{
		Success: true,
		Message: "Scores loaded",
		Data:    []int{95, 88, 92},
	}
	fmt.Printf("API Response 2: %+v\n", intListRes)
}
// so basiclly in nutshell genrics if i see lets say i have a function but that function can only use 
// int type i cant pass string cause this is not dynamically types langugae throw me an error 
// i just need to change the paratmer of the function and then 
// 