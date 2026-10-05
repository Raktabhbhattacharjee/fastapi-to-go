package main

import (
	"fmt"
	"time"
)

// In FastAPI, you would write:
// class User(BaseModel):
//     id: int
//     name: str
//     role: str
//     is_active: bool
//
// In Go, we use a struct (just like C, but memory-safe):
type User struct {
	ID       int
	Name     string
	Role     string
	IsActive bool
}

// In Go, methods are attached to structs via "receivers":
// (u User) is a value receiver (like copying the struct).
// (u *User) is a pointer receiver (modifies the original struct in-place, zero copy).
func (u *User) PromoteToAdmin() {
	u.Role = "admin"
	fmt.Printf("[Update] %s is now an %s!\n", u.Name, u.Role)
}

func main() {
	fmt.Println("========================================")
	fmt.Println("🚀 Welcome to fastapi-to-go!")
	fmt.Println("========================================")

	// 1. Struct creation & Pointer mechanics
	// &User creates the struct and returns a pointer (*User) to it
	currentUser := &User{
		ID:       1,
		Name:     "Rishi",
		Role:     "engineer",
		IsActive: true,
	}

	fmt.Printf("User created: %+v (Memory Address: %p)\n", *currentUser, currentUser)
	currentUser.PromoteToAdmin()

	// 2. Slices (Go's dynamic array, like Python list or C++ std::vector)
	endpoints := []string{"/health", "/users", "/metrics"}
	fmt.Println("\nRegistered Routes:")
	for idx, route := range endpoints {
		fmt.Printf("  [%d] %s\n", idx, route)
	}

	// 3. Concurrency sneak peek (Goroutines)
	// Notice how we don't need 'async def' or 'await'
	fmt.Println("\nStarting a background task (Goroutine)...")
	done := make(chan bool) // A channel: how goroutines communicate

	go func() {
		fmt.Println("  [Worker] Running background job on another thread...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("  [Worker] Job complete!")
		done <- true // notify main
	}()

	// Wait for the goroutine to finish (like an explicit join/await)
	<-done
	fmt.Println("\n🎉 Day 1 Setup Verified & Ready!")
}
