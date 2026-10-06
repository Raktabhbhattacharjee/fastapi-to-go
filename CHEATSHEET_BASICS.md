# Go Basics Cheat Sheet: Python to Go

A personal reference for a Python developer who is already comfortable with C and C++. It focuses on Go's everyday syntax, static types, collections, and control flow.

## Mental-model shifts

| Python / FastAPI | Go |
| --- | --- |
| `str`, `int`, `float`, `bool` | `string`, `int`, `float64`, `bool` |
| `list[str]` | `[]string` |
| `dict[str, int]` | `map[string]int` |
| `None` | `nil` for pointers, slices, maps, and similar reference-like values |
| Implicit numeric conversion is common | Numeric conversions must be explicit |

## 1. Variables and Types

### Declaring Variables
```go
// Go infers the type from the value.
name := "Rishi"
port := 8080
isActive := true

// var lets you state the type explicitly.
var timeout int = 30
var price float64 = 19.99

// Variables without a value receive the type's zero value.
var count int    // 0
var title string // ""
var ready bool   // false
```

### Zero Values and `nil`
```go
var count int       // 0
var title string    // ""
var ready bool      // false
var tags []string   // nil slice; len(tags) is 0
var metadata map[string]string // nil map; initialize before writing to it

metadata = make(map[string]string)
metadata["source"] = "api"
```

### Converting Types
```go
a := 10
b := 3.5

// Go does not combine different numeric types automatically.
// sum := a + b // ❌ Error!

// Convert one value before using it.
sum := float64(a) + b // 13.5
intSum := a + int(b)  // 13
```

### Constants
```go
const AppName = "MyApp"
const MaxRetries = 3
```

### Operators

| Purpose | Go operators | Example |
| --- | --- | --- |
| Arithmetic | `+`, `-`, `*`, `/`, `%` | `remainder := 10 % 3` |
| Comparison | `==`, `!=`, `<`, `<=`, `>`, `>=` | `isAdult := age >= 18` |
| Logic | `&&`, `||`, `!` | `allowed := isAdmin || isOwner` |
| Assignment | `=`, `+=`, `-=`, `*=`, `/=` | `count += 1` |

## 2. Control Flow

### A. `if / else`
```go
// Standard if / else chain.
score := 92
grade := ""
if score >= 90 {
    grade = "A"
} else if score >= 80 {
    grade = "B"
} else {
    grade = "C"
}

// A short statement can prepare a value for the condition.
name := "Rishi"
isLongName := false
if length := len(name); length > 5 {
    isLongName = true
}
```

---

### B. `switch`

Go finishes each case automatically, so `break` is usually unnecessary.

```go
// Match a value.
role := "admin"
access := ""
switch role {
case "admin":
    access = "full"
case "editor", "author": // Match either value.
    access = "edit"
default:
    access = "view"
}

// A tagless switch is useful for a sequence of conditions.
personAge := 20
group := ""
switch {
case personAge < 13:
    group = "child"
case personAge < 20:
    group = "teen"
default:
    group = "adult"
}
```

---

### C. `for` Loops

Go uses `for` for every kind of loop; it has no separate `while` keyword.

```go
// Standard three-part loop.
total := 0
for i := 0; i < 5; i++ {
    total += i
}

// Condition-only loop.
n := 1
for n < 10 {
    n *= 2
}

// Infinite loop; use break to exit it.
for {
    if shouldStop {
        break
    }
}
```


## 3. Slices

A slice is a flexible, ordered collection. It is similar to a Python list.

### Creating Slices
```go
// A slice literal.
fruits := []string{"apple", "banana", "cherry"}

// make creates a slice with length 0 and room for 10 elements.
numbers := make([]int, 0, 10)
```

### Common Slice Operations
```go
// Add items.
fruits = append(fruits, "orange")
fruits = append(fruits, "grape", "mango") // Append multiple

// Number of elements.
count := len(fruits)

// Take a range: start is included; end is excluded.
sub := fruits[0:2] // Elements at indices 0 and 1.

// Iterate with both index and value.
positions := make(map[string]int)
for idx, fruit := range fruits {
    positions[fruit] = idx
}

// Ignore the index with _ when only the value matters.
latestFruit := ""
for _, fruit := range fruits {
    latestFruit = fruit
}
```


## 4. Maps

A map stores values under unique keys. It is similar to a Python dictionary.

### Creating Maps
```go
// Create an empty map.
userAges := make(map[string]int)

// Create a map with initial values.
scores := map[string]int{
    "alice": 95,
    "bob":   88,
}
```

### Common Map Operations
```go
// Add or update an entry.
userAges["rishi"] = 25

// Read a value.
rishiAge := userAges["rishi"] // 25

// Remove an entry.
delete(userAges, "rishi")

// The comma-ok form tells you whether a key exists.
age, exists := userAges["alex"]
nextAge := 0
if !exists {
    age = 0
} else {
    nextAge = age + 1
}

// Iterate through keys and values.
copiedAges := make(map[string]int)
for name, age := range userAges {
    copiedAges[name] = age
}
```
