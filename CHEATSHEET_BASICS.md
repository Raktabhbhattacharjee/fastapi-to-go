# ⚡ Go Foundations & Internals Cheatsheet
*(Everything before Functions & Structs: Types, Memory, Scopes, Control Flow, Collections & Gotchas)*

---

## 1. Type System & Memory Primitives

### A. The No-Implicit-Casting Rule (Huge contrast to C & Python)
In C, `int + float` auto-promotes. In Python, `1 + 2.5` yields `3.5`.
**In Go, implicit type conversion DOES NOT EXIST.** Every conversion must be explicit:
```go
var a int = 10
var b float64 = 20.5

// COMPILE ERROR: invalid operation: a + b (mismatched types int and float64)
// total := a + b 

// CORRECT:
total := float64(a) + b // 30.5 (type float64)
intTotal := a + int(b)  // 30   (type int - truncated)
```

### B. Integer Widths & Platform Specifics
| Type | Bits | Range | When to Use |
| :--- | :--- | :--- | :--- |
| `int` / `uint` | 32 or 64 (arch dependent) | $-2^{63}$ to $2^{63}-1$ on 64-bit | **Default for counts, loop counters, sizes.** |
| `int8` / `uint8` | 8 bits | -128 to 127 / 0 to 255 | Low-level binary protocols, raw bytes. |
| `byte` | 8 bits | Alias for `uint8` | Raw I/O data, network buffers, UTF-8 bytes. |
| `int32` / `rune` | 32 bits | Unicode code point | Processing individual characters (emojis, unicode). |
| `int64` / `uint64`| 64 bits | Massive numbers / timestamps | Unix timestamps (nanoseconds), DB primary keys. |
| `uintptr` | Pointer-sized uint | Memory address representation | Unsafe memory operations / syscalls. |

---

## 2. Strings, Bytes & Runes Internals (The Python Traps)

In Python, a string is a sequence of characters. **In Go, a string is an immutable, read-only slice of UTF-8 encoded bytes (`[]byte`).**

### A. `len()` Returns BYTES, Not Characters!
```go
s1 := "hello"
fmt.Println(len(s1)) // 5 bytes

s2 := "café"
fmt.Println(len(s2)) // 5 bytes! ('é' takes 2 bytes in UTF-8: 0xc3, 0xa9)

s3 := "🚀"
fmt.Println(len(s3)) // 4 bytes!
```

### B. Indexing `s[i]` Gives Bytes, Not Characters
```go
s := "café"
fmt.Println(s[3])        // Prints 195 (byte 0xc3), NOT 'é'!
fmt.Printf("%c\n", s[3]) // Prints 'Ã' (broken character)
```

### C. The Solution: `rune` (Unicode Code Point)
To iterate characters correctly (including multi-byte UTF-8 & emojis):
```go
// Method 1: Range loop over string auto-decodes runes!
for byteIndex, r := range "café 🚀" {
    fmt.Printf("Byte pos: %d -> Character: %c\n", byteIndex, r)
}

// Method 2: Convert to []rune slice
runes := []rune("café")
fmt.Println(len(runes))          // 4 (actual character count!)
fmt.Printf("%c\n", runes[3])     // 'é'
```

### D. String Mutability & Conversion
Strings in Go are **100% immutable**. To modify, convert to `[]byte` or `[]rune`:
```go
str := "hello"
// str[0] = 'H' // COMPILE ERROR: cannot assign to str[0]

b := []byte(str)
b[0] = 'H'
str = string(b) // "Hello" (Allocates new string memory)
```

---

## 3. Pointers to Primitives (C vs Go)

Unlike C, **Go has NO pointer arithmetic (`p++` is illegal)**, and you cannot have dangling pointers (escape analysis moves variables to the heap if a reference survives).

```go
count := 42

// 1. Get memory address with '&'
ptr := &count // Type is *int
fmt.Println("Address:", ptr) // e.g. 0xc0000180b0

// 2. Dereference with '*' to read or write
fmt.Println("Value:", *ptr) // 42
*ptr = 100                  // Modifies 'count' directly in memory
fmt.Println("Count is now:", count) // 100

// 3. The new() keyword
p := new(int) // Allocates memory for an int, zeroes it, returns *int
fmt.Println(*p) // 0
```

---

## 4. Advanced Constants & `iota` (Bitmasks & Enums)

Constants in Go are untyped until assigned, allowing arbitrary-precision math at compile time.

### A. Auto-increment & Skipping with `iota`
```go
const (
    _  = iota             // 0 (discarded using blank identifier)
    KB = 1 << (10 * iota) // 1 << (10 * 1) = 1024
    MB = 1 << (10 * iota) // 1 << (10 * 2) = 1048576
    GB = 1 << (10 * iota) // 1 << (10 * 3) = 1073741824
)
```

### B. Bitmask / Permission Flags (Production Pattern)
```go
const (
    ReadPermission    = 1 << iota // 1 (0001)
    WritePermission               // 2 (0010) - inherits expression!
    ExecutePermission             // 4 (0100)
)

userPerms := ReadPermission | WritePermission // 3 (0011)
hasWrite := (userPerms & WritePermission) != 0 // true
```

---

## 5. Scopes, Shadowing & Package Rules

### A. The 3 Scope Levels
1. **Package Scope (File-independent):**
   * Go has **NO file-private scope**. Every file in `package main` sees all package-level variables and constants across all other files in that same folder.
   * `:=` is **FORBIDDEN** at package scope. Must use `var` or `const`.
2. **Function Scope:** Variables declared inside a function body.
3. **Block Scope:** Any curly braces `{ ... }`, including `if`, `for`, `switch`.

### B. The Shadowing Trap in Loops / Blocks
```go
port := 8080
for i := 0; i < 1; i++ {
    port := 9000 // ⚠️ SHADOWING! Creates a new 'port' in this block.
    fmt.Println("Inside loop:", port) // 9000
}
fmt.Println("Outside loop:", port) // Still 8080!

// Fix: use '=' to mutate the outer variable:
port = 9000
```

---

## 6. Control Flow Mastery: Beyond the Basics

### A. Labeled `break` and `continue` (Breaking Nested Loops)
In C or Python, breaking an inner loop requires a boolean flag to escape the outer loop. In Go, use **labels**:

```go
OuterLoop:
for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
        if i == 2 && j == 2 {
            fmt.Println("Breaking out of BOTH loops!")
            break OuterLoop // Terminates the outer loop directly
        }
    }
}
```

### B. Switch: `fallthrough` and Type Switching
By default, Go cases do NOT fall through. Use explicit `fallthrough` if needed:
```go
num := 1
switch num {
case 1:
    fmt.Println("One")
    fallthrough // Forces execution of case 2 without checking its condition!
case 2:
    fmt.Println("Two")
}
// Outputs: "One" then "Two"
```

---

## 7. Deep-Dive: Slices & Memory Internals

A slice does NOT hold data itself. It is a 24-byte header (on 64-bit systems) containing:
1. `Data` pointer $\rightarrow$ points to an underlying array
2. `Len` $\rightarrow$ number of elements accessed (`len(s)`)
3. `Cap` $\rightarrow$ maximum capacity before reallocation (`cap(s)`)

### A. Reslicing Shares Memory (The Silent Mutation Gotcha!)
```go
original := []int{10, 20, 30, 40}
sub := original[1:3] // [20, 30]

sub[0] = 999 // ⚠️ Modifies the UNDERLYING array!
fmt.Println(original) // [10, 999, 30, 40] -> Original was modified!
```

### B. The Slice Memory Leak Gotcha
If you slice a tiny piece of a huge array, the garbage collector **cannot free the huge array**:
```go
hugeArray := make([]byte, 100_000_000) // 100MB
sub := hugeArray[:2]                   // 2 bytes, BUT holds reference to all 100MB!

// ⭐️ PRODUCTION FIX: Use copy() to isolate memory:
subClean := make([]byte, 2)
copy(subClean, hugeArray[:2]) // Now hugeArray can be garbage collected!
```

### C. Capacity Growth Strategy
When `append()` exceeds `cap`, Go allocates a new array (usually $2\times$ up to 256 elements, then $\approx 1.25\times$), copies the data, and discards the old array.
* **Always pre-allocate if you know the size:** `make([]T, 0, expectedCapacity)`

---

## 8. Deep-Dive: Maps Internals & Gotchas

### A. The Nil Map Panic (Crucial Production Gotcha!)
```go
var m map[string]int // Declared, but uninitialized (nil)

fmt.Println(m["key"]) // Safe: returns zero-value (0)
// m["key"] = 100     // 💥 RUNTIME PANIC: assignment to entry in nil map!

// ALWAYS initialize maps before writing:
m = make(map[string]int)
m["key"] = 100 // Safe!
```

### B. Map Iteration is Random by Design!
In Python 3.7+, dicts preserve insertion order. **In Go, map iteration order is deliberately randomized by the runtime:**
```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
for k := range m {
    fmt.Print(k, " ") // Order changes across different runs!
}
// If you need deterministic order: collect keys into a slice, sort it, and iterate over keys.
```

### C. Deleting Missing Keys is 100% Safe
```go
delete(m, "non-existent-key") // No error, no panic, does nothing safely.
```

---

## 9. Quick Summary Reference Table

| Concept | The Gotcha / Rule | Production Best Practice |
| :--- | :--- | :--- |
| **Type Conversion** | No auto-promotion (`int + float` fails) | Explicit casting: `float64(x) + y` |
| **Strings** | `len(s)` is byte count, not char count | Use `[]rune(s)` or `for _, r := range s` for UTF-8 |
| **String Mutation**| Strings are read-only immutable byte slices | Convert to `[]byte`, edit, cast back |
| **Pointers** | No pointer arithmetic (`p++` is illegal) | Use `*` to dereference, `&` to get address |
| **Constants** | `iota` auto-increments | Ideal for bitmasks (`1 << iota`) and enums |
| **Scopes** | All files in a package share package scope | Use `:=` locally; package scope requires `var`/`const` |
| **Loop Breaking** | Inner `break` only exits inner loop | Use labeled break: `break OuterLoop` |
| **Slice Sharing** | Sub-slices mutate the parent array | Use `copy()` if data must be independent |
| **Slice Pre-allocation** | Appending without capacity causes reallocations | `make([]T, 0, capacity)` |
| **Nil Maps** | Writing to an uninitialized map panics | Always allocate with `make(map[K]V)` |
| **Map Order** | Iteration order is randomized by the runtime | Sort slice of keys if deterministic order is needed |
