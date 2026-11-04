# Golang Error Handling: The Complete Cheatsheet

## 1. How Other Languages Handle Errors vs How Go Does It

| Language | Mechanism | How It Works | The Big Problem |
| :--- | :--- | :--- | :--- |
| **Python / JS / Java / C++** | **Exceptions** (`try / catch / throw`) | An error violently jumps up the call stack until an invisible `catch` block catches it. | **Hidden control flow:** You never know which function might blow up without reading all source code. |
| **C** | **Return Codes** (`-1`, `NULL`, `errno`) | Functions return numbers like `-1` on error. | You can't return real data and an error together cleanly; easy to forget checking. |
| **Go** | **Errors are Values** (`val, err`) | An error is a normal, explicit return value. No hidden jumps. | You type `if err != nil` frequently, but control flow is **100% visible and predictable**! |

---

## 2. The Core Go Error Philosophy

In Go:
> **Errors are NOT special magic. An error is just a normal interface value.**

The built-in `error` type in Go is defined as:
```go
type error interface {
    Error() string
}
```
Any struct or type in your entire program that has a method `Error() string` automatically counts as an error!

---

## 3. The 4 Essential Error Techniques

### 1. Sentinel Errors (Fixed Constants)
Used for standard, predictable conditions that callers need to compare against:
```go
import "errors"

var (
    ErrNotFound     = errors.New("user not found")
    ErrUnauthorized = errors.New("unauthorized access")
)
```

### 2. Formatted Errors (Dynamic Messages)
Used when the error needs runtime context (IDs, names, numbers):
```go
func findUser(id int) (*User, error) {
    if id <= 0 {
        return nil, fmt.Errorf("invalid user id: %d", id)
    }
    // ...
}
```

### 3. Custom Error Structs (Rich Context)
When you need to attach more metadata (HTTP status codes, timestamps, retry counts):
```go
type ValidationError struct {
    Field   string
    Problem string
}

// Implement the error interface:
func (v *ValidationError) Error() string {
    return fmt.Sprintf("field '%s' is invalid: %s", v.Field, v.Problem)
}
```

### 4. Error Wrapping & Unwrapping (Go 1.13+)
Wrap low-level errors with `%w` to add context without erasing the original root cause:
```go
// WRAP with %w:
if err := queryDB(); err != nil {
    return fmt.Errorf("getUser query failed: %w", err)
}

// CHECK with errors.Is():
if errors.Is(err, ErrNotFound) {
    // True even if ErrNotFound is buried deep inside 5 wrapped layers!
}

// EXTRACT with errors.As():
var valErr *ValidationError
if errors.As(err, &valErr) {
    fmt.Println("Invalid field was:", valErr.Field)
}
```

---

## 4. Panic & Recover (The Crash Protector)

### Rule: NEVER use `panic` for regular business errors!
* **Use Normal Errors (`error`) for:** Invalid passwords, file not found, network timeouts, bad JSON.
* **Use `panic` ONLY for:** Critical programmer bugs that make it impossible for the program to continue (e.g., failed to load config on startup, nil pointer dereference, array index out of bounds).

### How to Recover from a Panic:
```go
func safeWorker() {
    // defer + recover MUST be used together!
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Caught fatal crash:", r)
        }
    }()

    // Dangerous code that might panic:
    var ptr *int
    *ptr = 5 // nil pointer dereference!
}
```

---

## 5. Go Style Conventions for Errors

1. **Do NOT capitalize error strings:**
   * ✅ `errors.New("connection refused")`
   * ❌ `errors.New("Connection refused")`
   *(Why? Because error messages are often combined with other text: `fmt.Errorf("reading file: %w", err)`).*

2. **Do NOT end errors with punctuation:**
   * ✅ `errors.New("user not found")`
   * ❌ `errors.New("user not found.")`

3. **Handle errors immediately:**
   Keep the "happy path" on the left margin, and handle the error inside the `if`:
   ```go
   val, err := doSomething()
   if err != nil {
       return nil, err // Early return on failure
   }

   // Happy path continues cleanly without deep nesting:
   process(val)
   ```
