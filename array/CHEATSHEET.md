# Go Slices Cheat Sheet (Python to Go)

A quick-reference guide for slices in Go. If you are coming from Python (`list`), slices are Go's dynamic lists.

---

## 1. Quick Comparison Table

| What you want to do | Python (`list`) | Go (`slice`) | Frequency |
| :--- | :--- | :--- | :--- |
| **Empty list** | `items = []` | `var items []string` | Daily (90%) |
| **List with values** | `nums = [1, 2, 3]` | `nums := []int{1, 2, 3}` | Daily (90%) |
| **Add item** | `items.append("a")` | `items = append(items, "a")` | Daily (90%) |
| **Add multiple items**| `items.extend(["b", "c"])` | `items = append(items, "b", "c")` | Daily |
| **Length (count)** | `len(items)` | `len(items)` | Daily (90%) |
| **Access element** | `items[0]` | `items[0]` | Daily (90%) |
| **Loop (values only)**| `for item in items:` | `for _, item := range items { ... }` | Daily (90%) |
| **Loop (index + val)**| `for i, item in enumerate(items):` | `for i, item := range items { ... }` | Daily |
| **Sub-slice** | `items[:2]` | `items[:2]` | Common |
| **Pre-allocate space**| *(not needed in Python)* | `items := make([]string, 0, 100)` | Performance |

---

## 2. The Core 4 Concepts You Actually Use

### ① Declaring & Appending
```go
// Empty slice ready to use (nil slice, length is 0)
var users []string

// append() adds items and returns the updated slice
users = append(users, "Alice")
users = append(users, "Bob", "Charlie")
```

### ② Looping with `range`
In Go, `range` gives both `(index, value)`. If you don't need the index, discard it with `_`:
```go
for _, user := range users {
    fmt.Println(user)
}
```

### ③ Sub-slicing `[start:end]`
Same rule as Python: `start` is included, `end` is **excluded**.
```go
scores := []int{10, 20, 30, 40}
topTwo := scores[:2]   // [10, 20]
rest   := scores[2:]   // [30, 40]
```

### ④ Pre-allocating with `make()`
Use this when you know in advance how many items you will add (e.g., mapping 100 database records):
```go
// make([]Type, length, capacity)
// length = 0 (starts empty)
// capacity = 50 (reserves space for 50 items up front)
results := make([]string, 0, 50)
```

---

## 3. What You Can Safely Ignore (For Now)

| Feature | Why you can ignore it right now |
| :--- | :--- |
| **Fixed-size Arrays (`[5]int`)** | Rarely used in application code. Always use `[]int` (no size in brackets). |
| **`cap()`** | Only needed when benchmarking or optimizing memory allocations. |
| **`copy()`** | Rarely needed in web APIs. |
| **Manual index deletion** (`append(s[:i], s[i+1:]...)`) | In real backends, you use a `map`, filter into a new slice, or delete rows in the DB. |
