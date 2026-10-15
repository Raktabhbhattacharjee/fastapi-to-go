# Go Pointers Cheat Sheet

A straightforward, memory-focused reference for pointers in Go, especially coming from C and C++.

---

## 1. Quick Concept Map (C/C++ vs Go)

| Concept | C / C++ | Go | Notes |
| :--- | :--- | :--- | :--- |
| **Address-of Operator** | `&num` | `&num` | Identical: gives memory address |
| **Pointer Type** | `int *p` | `*int` | In Go, type goes on the right |
| **Dereference Operator**| `*p = 5` | `*p = 5` | Identical: access/modify value at address |
| **Null Pointer** | `NULL` or `nullptr` | `nil` | Default zero-value for pointers |
| **Access Struct Field** | `ptr->field` | `ptr.field` | **No `->` in Go!** Go auto-dereferences |
| **Pointer Arithmetic** | `ptr++`, `ptr + 4` | ❌ **Not allowed** | Go prevents buffer overflows |
| **Freeing Memory** | `free(ptr)`, `delete`| ❌ **Automatic** | Go runtime uses Garbage Collection |

---

## 2. The Only Two Operators You Need

### ① The `&` Operator (Address-Of)
Gets the memory address where a variable is stored in RAM:
```go
num := 42
fmt.Println(&num) // Prints memory address (e.g., 0xc000014070)
```

### ② The `*` Operator (Two jobs)
1. **In Type Declarations**: Means "this variable holds an address of that type".
   ```go
   var ptr *int // ptr is a pointer to an int
   ```
2. **In Expressions (Dereferencing)**: Means "go to that address in RAM".
   ```go
   *ptr = 100 // writes 100 into the memory address
   fmt.Println(*ptr) // reads the value stored at that address
   ```

---

## 3. How Function Memory Works Under the Hood (The Stack)

Understanding how the call stack executes functions makes pointers instantly intuitive.

### Scenario A: Pass-by-Value (Copy)
```go
func ChangeVal(val int) {
    val = 99
}
```

**Memory execution:**
1. `main()` gets its own **Stack Frame**:
   ```text
   [main Stack Frame]
   num = 1 (Address: 0x1000)
   ```
2. `main()` calls `ChangeVal(num)`. Go creates a **NEW Stack Frame** on top:
   ```text
   [ChangeVal Stack Frame]
   val = 1 (Address: 0x2000) <-- A completely separate COPY in RAM!
   -----------------------
   [main Stack Frame]
   num = 1 (Address: 0x1000)
   ```
3. `val = 99` updates address `0x2000`.
4. `ChangeVal()` finishes. Its stack frame is destroyed.
5. In `main()`, address `0x1000` is still `1`!

---

### Scenario B: Pass-by-Pointer (Reference)
```go
func ChangeNum(num *int) {
    *num = 5
}
```

**Memory execution:**
1. `main()` sets up its Stack Frame:
   ```text
   [main Stack Frame]
   num = 1 (Address: 0x1000)
   ```
2. `main()` calls `ChangeNum(&num)` (passing address `0x1000`):
   ```text
   [ChangeNum Stack Frame]
   num = 0x1000 (points directly to main's address!)
   -----------------------
   [main Stack Frame]
   num = 1 (Address: 0x1000)
   ```
3. `*num = 5` says: *"Follow address `0x1000` and write `5` there"*.
   It modifies `main`'s memory directly!
4. `ChangeNum()` exits and its stack frame is destroyed.
5. In `main()`, `num` is now `5`!

---

## 4. Pointers with Structs (No `->` Operator!)

In C/C++, you had to use the arrow operator (`ptr->name`) for struct pointers.  
In Go, **you always use dot `.` notation**:

```go
type User struct {
    Name string
}

func UpdateName(u *User) {
    u.Name = "Rishi" // Go automatically does (*u).Name for you!
}

func main() {
    user := &User{Name: "Alice"}
    UpdateName(user)
    fmt.Println(user.Name) // "Rishi"
}
```

---

## 5. When Should You Use Pointers in Go?

Only use pointers for two reasons:
1. **You need to mutate data**: You want a function or method to modify the caller's variable.
2. **Performance / Avoiding large copies**: If a struct is very large (e.g., hundreds of fields), passing a pointer (`8 bytes`) is faster than copying the whole struct.

*(For small basic types like `int`, `bool`, `float64`, passing by value is usually faster and safer).*
