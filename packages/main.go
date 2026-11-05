package main

import (
	"fmt"
	"fastapi-to-go/packages/auth" // <--- Import your package!
)

func main() {
	fmt.Println("=== FASTAPI STYLE APP IN GO ===")

	// 1. Call the public function from auth package:
	user, err := auth.RegisterUser("rishi", "rishi@example.com", "supersecret123")
	if err != nil {
		fmt.Println("Registration failed:", err)
		return
	}

	// 2. Read public fields (Username, Email):
	fmt.Println("User registered successfully!")
	fmt.Println("Username:", user.Username)
	fmt.Println("Email:", user.Email)

	// 3. Try verifying the password through the public method:
	matched := user.VerifyPassword("supersecret123")
	fmt.Println("Password matched?", matched)

	// NOTE: If you tried to do: fmt.Println(user.password)
	// The Go compiler will refuse to compile because 'password' is private!
}