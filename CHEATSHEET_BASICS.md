# Go Core Fundamentals Cheatsheet
*(Variables, Scopes, Control Flow & Data Structures)*

---

## 1. Variables & Zero Values

### A. Three Ways to Declare
```go
// 1. Short declaration (inside functions only - 90% of production Go)
name := "Rishi"
port := 8080
isActive := true

// 2. Explicit type declaration (used when zero-value or specific width is needed)
var maxConns int64 = 10000
var timeout float64 = 5.5

// 3. Declaration without initial value (guaranteed Zero Value)
var count int       // 0
var title string    // "" (empty string, NEVER null/nil)
var isReady bool    // false
var data []byte     // nil
```

### B. The Production Rule for Zero Values
In C, uninitialized memory contains dangerous garbage. In Go, types initialize safely to their **zero value**:
* `int`, `float` $\rightarrow$ `0`, `0.0`
* `bool` $\rightarrow$ `false`
* `string` $\rightarrow$ `""`
* Pointers, Slices, Maps, Channels, Interfaces $\rightarrow$ `nil`

---

## 2. Scopes in Go (Crucial!)

Go has 4 levels of scope. There are no `public` or `private` keywords!

| Scope Level | Rule | Example |
| :--- | :--- | :--- |
| **Package Exported** (Public) | Capitalized identifier (`PascalCase`) | `func HandleRequest()`, `var MaxLimit` — visible to other packages. |
| **Package Unexported** (Private) | Lowercase identifier (`camelCase`) | `func validate()`, `var dbConn` — visible only within the same package. |
| **Function Scope** | Declared inside a function | `func main() { x := 10 }` — visible only inside `main()`. |
| **Block / Statement Scope** | Declared inside `{ }` or in an `if`/`for` init | `if x := get(); x > 0 { ... }` — `x` ceases to exist outside that `if` block! |

### Short Variable Shadowing Pitfall (Beware!)
```go
x := 10
if true {
    x := 20 // ⚠️ Shadows the outer x! This creates a NEW local variable inside this block.
    fmt.Println(x) // Prints 20
}
fmt.Println(x) // Prints 10 (outer x was NOT modified!)

// Fix: use '=' instead of ':=' if you want to reassign the outer variable:
if true {
    x = 20
}
```

---

## 3. Control Flow

### A. `if / else` & The Go "Init Statement"
Go allows you to execute a short statement **before** the condition. The variable declared only exists within that `if/else` block.

```go
// Standard if / else
if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else {
    fmt.Println("C")
}

// ⭐️ PRODUCTION IDIOM: If with Init Statement (Used everywhere for errors & lookups)
// Pattern: if <init>; <condition> { ... }
if val, ok := cache["user:1"]; ok {
    // 'val' and 'ok' are ONLY in scope inside this if/else block!
    fmt.Println("Cache hit:", val)
} else {
    fmt.Println("Cache miss")
}
// 'val' does not leak into the rest of the function!
```

---

### B. `for` Loops (Go's ONLY Loop Keyword)
Go has **no `while`** or **`do-while`**. The `for` keyword handles every kind of loop.

#### 1. Standard 3-Component Loop (like C / C++)
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

#### 2. While-Style Loop
```go
count := 1
for count < 100 {
    count *= 2
}
```

#### 3. Infinite Loop (Production Servers / Workers)
```go
for {
    // Runs indefinitely until 'break' or 'return'
    if checkStopSignal() {
        break
    }
}
```

#### 4. The `for range` Loop (Like Python's `for item in items`)
Used to iterate through Slices, Maps, and Strings.

```go
// Over a Slice (index, value)
fruits := []string{"apple", "banana", "cherry"}
for idx, fruit := range fruits {
    fmt.Printf("%d: %s\n", idx, fruit)
}

// Ignore index using blank identifier `_`
for _, fruit := range fruits {
    fmt.Println(fruit)
}

// Over a Map (key, value)
users := map[string]int{"Alice": 25, "Bob": 30}
for name, age := range users {
    fmt.Printf("%s is %d\n", name, age)
}
```

---

### C. `switch` Statement (Modern & Clean)
Key differences from C/C++:
1. **No `break` needed!** Cases do not fall through automatically.
2. Conditions can be strings, comparisons, or expressions.

```go
// 1. Basic Switch
role := "admin"
switch role {
case "admin":
    fmt.Println("Full access")
case "editor", "moderator": // Multiple matches in one case
    fmt.Println("Limited edit access")
default:
    fmt.Println("Read-only")
}

// 2. Tagless Switch (Cleaner replacement for long if/else chains)
statusCode := 404
switch {
case statusCode >= 200 && statusCode < 300:
    fmt.Println("Success")
case statusCode >= 400 && statusCode < 500:
    fmt.Println("Client error")
case statusCode >= 500:
    fmt.Println("Server error")
}
```

---

## 4. Go Data Structures: What's Used in Production?

### Quick Summary:
* **Arrays (`[N]T`)**: Fixed-size, value-copied. **Rarely used in production** (mostly internal cryptographic keys/buffers).
* **Slices (`[]T`)**: Dynamic, backed by an array pointer. **Used 90% of the time**.
* **Maps (`map[K]V`)**: Hash tables. **Used 10% of the time** (caching, lookups, JSON dicts).

---

### A. Slices (`[]T`) — The King of Collections
A slice has 3 components under the hood:
1. Pointer to the underlying array
2. Length (`len`) — number of elements currently stored
3. Capacity (`cap`) — total space allocated before a reallocation is needed

```go
// 1. Literal declaration
nums := []int{10, 20, 30}

// 2. Pre-allocating with make() (⭐️ CRITICAL FOR HIGH-PERFORMANCE PRODUCTION CODE)
// make([]Type, length, capacity)
// Pre-allocating capacity avoids expensive reallocations when appending!
items := make([]string, 0, 100) // length 0, capacity for 100 items

// 3. Appending items (like Python list.append)
items = append(items, "first")
items = append(items, "second", "third") // can append multiple!

// 4. Slicing (half-open range [low:high], same as Python)
sub := items[0:2] // items at index 0 and 1
```

---

### B. Maps (`map[KeyType]ValueType`) — Fast Hash Tables
Maps are reference types. Keys must be comparable types (string, int, etc.).

```go
// 1. Literal declaration
scores := map[string]int{
    "alpha": 100,
    "beta":  85,
}

// 2. Allocating with make()
cache := make(map[string]string)

// 3. Insert or Update
cache["token"] = "xyz_123"

// 4. Delete a key
delete(cache, "token") // safe even if key doesn't exist

// 5. ⭐️ The "Comma-ok" Idiom (Check if key exists)
// In Python: if "token" in cache:
// In Go:
val, exists := cache["token"]
if !exists {
    fmt.Println("Key not found!")
} else {
    fmt.Println("Found:", val)
}
```

---

## 5. Production Cheatsheet Summary

| Task | Idiomatic Go Code |
| :--- | :--- |
| Check error / lookup | `if val, ok := m[key]; ok { ... }` |
| Fast dynamic list | `list := make([]T, 0, expectedSize)` |
| Dictionary / Map | `m := make(map[string]int)` |
| Safe uninitialized value | Rely on zero values (`0`, `""`, `false`, `nil`) |
| Public vs Private | `ExportedName` (Public) vs `unexportedName` (Private) |
| Iterate items | `for idx, item := range slice { ... }` |
| Iterate without index | `for _, item := range slice { ... }` |
