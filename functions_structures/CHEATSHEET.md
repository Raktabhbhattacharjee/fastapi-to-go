# Go Functions Cheat Sheet

A comprehensive, clean guide to all function syntax, patterns, and conventions in Go (especially coming from C/C++ and Python).

---

## 1. Quick Concept Map

| Feature | Python | C / C++ | Go |
| :--- | :--- | :--- | :--- |
| **Declaration** | `def add(a: int, b: int) -> int:` | `int add(int a, int b) { ... }` | `func Add(a int, b int) int { ... }` |
| **Return Type Position** | At the end (`-> int`) | At the very front (`int`) | **At the end (`int`)** |
| **Multiple Returns** | Tuple `return x, y` | Struct or out-pointers | **Native `return x, err`** |
| **Error Handling** | `try / except` | Error codes / errno | **Comma-ok `if err != nil`** |
| **Method on Object** | `def method(self):` | Member function inside class | **Receiver `func (u *User) Method()`** |
| **Constructor** | `def __init__(self):` | `User::User()` | **Factory `func NewUser() *User`** |

---

## 2. Anatomy of a Go Function

```text
func Name(param1 Type, param2 Type) ReturnType {
     return value
}
```

### Type Shortcut:
If consecutive parameters share the same type, you only need to write the type once:
```go
// Instead of: a int, b int, c int
func Add(a, b, c int) int {
    return a + b + c
}
```

---

## 3. The 5 Core Flavors of Functions in Go

### Flavor 1: Basic Function (Single Return)
```go
func Greet(name string) string {
    return "Hello, " + name
}

// Call it:
msg := Greet("Rishi")
```

---

### Flavor 2: Multiple Return Values (The `(Result, Error)` Pattern)
Go's signature pattern for robust error handling without exceptions:

```go
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

// Standard usage in production:
result, err := Divide(10, 2)
if err != nil {
    log.Fatal("Division error:", err)
}
fmt.Println("Result:", result)
```

---

### Flavor 3: Variadic Functions (Arbitrary Number of Arguments)
Like `*args` in Python or `...` in C:

```go
// 'nums' is treated as a slice ([]int) inside the function
func Sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

// Usage:
Sum(1, 2)          // 3
Sum(1, 2, 3, 4, 5) // 15

// Or pass an existing slice using '...':
numbers := []int{10, 20, 30}
Sum(numbers...)    // 60
```

---

### Flavor 4: Methods (Functions with a Receiver)
Methods are functions attached to a `struct`. Look at the parentheses before the name:

#### A. Value Receiver (Read-Only Copy)
```go
func (u User) GetName() string {
    return u.Name // Receives a copy; cannot change original in RAM
}
```

#### B. Pointer Receiver (Mutates Real Memory)
```go
func (u *User) SetName(name string) {
    u.Name = name // 🎯 Directly modifies original struct in RAM!
}
```

---

### Flavor 5: Constructors / Factory Functions (`New...`)
Standard Go convention to initialize a struct and return a pointer:

```go
func NewUser(name string, age int) *User {
    return &User{
        Name: name,
        Age:  age,
    }
}

// Usage:
user := NewUser("Rishi", 22)
```

---

## 4. Anonymous Functions & Closures

Go treats functions as **first-class citizens** (you can store them in variables, pass them, or return them):

### Anonymous Function (Executed Immediately):
```go
func() {
    fmt.Println("I am running immediately!")
}()
```

### Stored in a Variable:
```go
double := func(x int) int {
    return x * 2
}
fmt.Println(double(5)) // 10
```

### Closure (Remembers outer scope variables):
```go
func Counter() func() int {
    count := 0
    return func() int {
        count++ // Remembers and modifies 'count'
        return count
    }
}

// Usage:
next := Counter()
fmt.Println(next()) // 1
fmt.Println(next()) // 2
fmt.Println(next()) // 3
```

---

## 5. Summary Cheat Table

| What you want to do | Go Syntax | How you call it |
| :--- | :--- | :--- |
| **Basic Calculation** | `func Add(a, b int) int` | `Add(1, 2)` |
| **No Return (Void)** | `func Log(msg string)` | `Log("hello")` |
| **Return Data + Error** | `func Fetch() (Data, error)` | `data, err := Fetch()` |
| **Constructor** | `func New...() *Struct` | `s := New...()` |
| **Read-Only Method** | `func (s Struct) Method()` | `s.Method()` |
| **Mutating Method** | `func (s *Struct) Method()` | `s.Method()` |
| **Variable Arguments** | `func Sum(nums ...int) int` | `Sum(1, 2, 3)` |
