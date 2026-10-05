# 📝 Go Basics Cheatsheet
*(Syntax, Scopes, Control Flow, Slices & Maps)*

---

## 1. Variables & Types

### Three Ways to Declare
```go
// 1. Short declaration (inside functions - most common)
name := "Rishi"
port := 8080
isActive := true

// 2. Explicit type declaration (with 'var')
var timeout int = 30
var price float64 = 19.99

// 3. Zero Values (declared without initial value)
var count int    // Defaults to 0
var title string // Defaults to "" (empty string)
var ready bool   // Defaults to false
```

### Type Conversions (Must be explicit)
```go
a := 10
b := 3.5

// Go does not allow mixing types directly:
// sum := a + b // ❌ Error!

// Cast explicitly:
sum := float64(a) + b // 13.5
intSum := a + int(b)  // 13
```

### Constants & Simple Enums (`iota`)
```go
const AppName = "MyApp"
const MaxRetries = 3

// Enums using iota (auto-increments: 0, 1, 2)
const (
    StatusPending = iota // 0
    StatusActive         // 1
    StatusCompleted      // 2
)
```

---

## 2. Scopes in Go

Go does not have `public` or `private` keywords. It uses capitalization:

| Rule | Visibility | Example |
| :--- | :--- | :--- |
| **Capitalized** | **Public** (Exported to other packages) | `MaxLimit`, `CalculateTotal` |
| **Lowercase** | **Private** (Only inside current package) | `dbPassword`, `validateUser` |

### Block Scope & Reassignment Trap
```go
count := 10

if true {
    // ⚠️ Trap: ':=' here creates a NEW variable inside this block
    count := 20
    fmt.Println(count) // Prints 20
}
fmt.Println(count) // Prints 10 (outer variable was untouched!)

// ✅ Correct: use '=' to update the existing variable
if true {
    count = 20
}
fmt.Println(count) // Prints 20
```

---

## 3. Control Flow

### A. `if / else`
```go
// Standard if / else
if score >= 90 {
    fmt.Println("Grade A")
} else if score >= 80 {
    fmt.Println("Grade B")
} else {
    fmt.Println("Grade C")
}

// If with Short Statement (variable only exists inside the if/else block)
if length := len(name); length > 5 {
    fmt.Println("Long name:", length)
}
```

---

### B. `switch`
* No `break` needed (Go stops automatically at the end of each case).

```go
// 1. Matching values
role := "admin"
switch role {
case "admin":
    fmt.Println("Full access")
case "editor", "author": // Match multiple values
    fmt.Println("Edit access")
default:
    fmt.Println("Viewer access")
}

// 2. Tagless switch (alternative to long if/else chains)
age := 20
switch {
case age < 13:
    fmt.Println("Child")
case age < 20:
    fmt.Println("Teen")
default:
    fmt.Println("Adult")
}
```

---

### C. `for` Loops (Go has no `while`)

```go
// 1. Standard loop (like C)
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// 2. While-style loop
n := 1
for n < 10 {
    n *= 2
}

// 3. Infinite loop (use 'break' to stop)
for {
    if shouldStop {
        break
    }
}
```

---

## 4. Slices (Dynamic Lists)

Slices are Go's dynamic lists (like Python lists).

### Creating Slices
```go
// Direct literal
fruits := []string{"apple", "banana", "cherry"}

// Using make() with initial length and capacity
numbers := make([]int, 0, 10)
```

### Common Slice Operations
```go
// Append items
fruits = append(fruits, "orange")
fruits = append(fruits, "grape", "mango") // Append multiple

// Length
fmt.Println(len(fruits)) // Number of elements

// Slicing (same as Python [start:end])
sub := fruits[0:2] // First 2 items (indices 0 and 1)

// Loop with index and value
for idx, fruit := range fruits {
    fmt.Printf("Index %d: %s\n", idx, fruit)
}

// Loop values only (ignore index using '_')
for _, fruit := range fruits {
    fmt.Println(fruit)
}
```

---

## 5. Maps (Key-Value Dictionaries)

Maps store key-value pairs (like Python dicts).

### Creating Maps
```go
// Using make()
userAges := make(map[string]int)

// Literal map
scores := map[string]int{
    "alice": 95,
    "bob":   88,
}
```

### Common Map Operations
```go
// Set / Update
userAges["rishi"] = 25

// Read
fmt.Println(userAges["rishi"]) // 25

// Delete a key
delete(userAges, "rishi")

// Check if a key exists (The "Comma-ok" check)
age, exists := userAges["alex"]
if !exists {
    fmt.Println("User not found!")
} else {
    fmt.Println("Age is:", age)
}

// Loop through key and value
for name, age := range userAges {
    fmt.Printf("%s is %d years old\n", name, age)
}
```
