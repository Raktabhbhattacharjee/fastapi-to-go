# Go Enums & Generics Cheat Sheet

A practical guide to how Go handles **Enums** and **Generics**, especially coming from Python (FastAPI/Pydantic) and C++.

---

## Part 1: Enums in Go (No `enum` Keyword!)

In Python/FastAPI, you use Python's `Enum` class. In Go, there is **no `enum` keyword**.  
Instead, Go uses **Custom Types** + **`const` blocks**.

### 1. String Enums (The FastAPI / JSON API Pattern)
This is what you will use 90% of the time when building REST APIs (matches JSON strings):

#### In FastAPI (Python):
```python
from enum import Enum

class OrderStatus(str, Enum):
    PENDING = "pending"
    CONFIRMED = "confirmed"
    CANCELLED = "cancelled"
```

#### In Go:
```go
// 1. Define custom type based on string
type OrderStatus string

// 2. Define the allowed constants
const (
    StatusPending   OrderStatus = "pending"
    StatusConfirmed OrderStatus = "confirmed"
    StatusCancelled OrderStatus = "cancelled"
)

// Usage:
type Order struct {
    ID     string      `json:"id"`
    Status OrderStatus `json:"status"` // Type-safe!
}
```

---

### 2. Integer Enums with `iota` (Auto-Incrementing: 0, 1, 2...)
For internal state, roles, or flags where you want numbers instead of strings, Go provides the **`iota`** keyword (an auto-incrementing counter):

```go
type Role int

const (
    RoleUser       Role = iota // 0
    RoleModerator              // 1 (auto-incremented)
    RoleAdmin                  // 2 (auto-incremented)
    RoleSuperAdmin             // 3 (auto-incremented)
)
```

---

## Part 2: Generics in Go (Type Parameters `[T any]`)

### What Problem Do Generics Solve?

Before Go 1.18, if you wanted a function to print a slice of `int`s and a slice of `string`s, you had to write **two duplicate functions**:
```go
func PrintInts(items []int) { ... }
func PrintStrings(items []string) { ... }
```

With **Generics**, you write **one function** with a type placeholder `[T any]`:

```go
// T stands for "any Type"
func PrintSlice[T any](items []T) {
    for _, item := range items {
        fmt.Println(item)
    }
}

// Go figures out the type automatically:
PrintSlice([]int{1, 2, 3})            // T is int
PrintSlice([]string{"Alice", "Bob"})  // T is string
```

---

### The #1 Real-World Production Use Case: Generic API Envelopes

In FastAPI, you often return a standardized response schema:
```python
# FastAPI Pydantic: Generic Response Envelope
class APIResponse(BaseModel, Generic[T]):
    status: str
    data: T
```

In Go, you do the **exact same thing** with a generic struct:

```go
// T can be a User, an Order, a Slice of Products, etc.
type APIResponse[T any] struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    Data    T      `json:"data"`
}

func main() {
    // 1. Response wrapping a single User struct:
    userResp := APIResponse[User]{
        Success: true,
        Message: "User found",
        Data:    User{Name: "Rishi"},
    }

    // 2. Response wrapping a list of Orders ([]Order):
    orderListResp := APIResponse[[]Order]{
        Success: true,
        Message: "Orders fetched",
        Data:    []Order{{ID: "1"}, {ID: "2"}},
    }
}
```

---

## Part 3: Type Constraints (Restricting `T`)

Sometimes `T` cannot be literally "any" type. For example, if you want to compare values (`==`), or do math (`+`).

### 1. `any`
Accepts **any type whatsoever** (int, string, struct, pointer, slice):
```go
func Identity[T any](val T) T {
    return val
}
```

### 2. `comparable`
Accepts types that support `==` and `!=` (numbers, strings, booleans, pointers):
```go
// Checks if an element exists in a slice:
func Contains[T comparable](items []T, target T) bool {
    for _, item := range items {
        if item == target { // Needs 'comparable' to use ==
            return true
        }
    }
    return false
}
```

### 3. Custom Constraint (Type Unions `int | float64`)
When you want to do math, restrict `T` to numeric types using the pipe `|` symbol:

```go
// Define a constraint that only allows numeric types:
type Number interface {
    int | int64 | float32 | float64
}

func Add[T Number](a, b T) T {
    return a + b // Safe to use '+' because T is guaranteed to be a number!
}
```

---

## Summary Cheat Table

| Concept | Go Syntax | Notes |
| :--- | :--- | :--- |
| **String Enum** | `type Status string; const (...)` | Ideal for JSON API models (FastAPI style) |
| **Integer Enum**| `type Role int; const ( User = iota )` | Auto-increments from 0, 1, 2... |
| **Generic Function** | `func Process[T any](item T)` | Write once, works for any type |
| **Generic Struct** | `type Envelope[T any] struct { Data T }`| The standard API response wrapper |
| **Equality Check** | `[T comparable]` | Allows using `==` and `!=` on `T` |
| **Numeric Math** | `[T int \| float64]` | Allows using `+`, `-`, `*`, `/` on `T` |
