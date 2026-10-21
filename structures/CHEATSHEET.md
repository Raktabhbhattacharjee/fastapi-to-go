# Go Structs Cheat Sheet

A quick, step-by-step reference for understanding structs, receivers, and instances in Go (especially coming from C/C++ and Python).

---

## 1. Quick Concept Map

| Concept | C / C++ | Go | Python |
| :--- | :--- | :--- | :--- |
| **Blueprint** | `typedef struct { ... } Order;` | `type Order struct { ... }` | `class Order:` / `@dataclass` |
| **Instance** | `Order o1 = { ... };` | `o1 := Order{ ... }` | `o1 = Order(...)` |
| **Method** | Member function in class | Receiver function: `func (o *Order) Method()` | `def method(self):` |
| **Access Field** | `o1.field` | `o1.field` | `o1.field` |
| **Pointer Access** | `ptr->field` | `ptr.field` *(no `->` in Go!)* | *(all objects are references)* |

---

## 2. Step-by-Step Breakdown

### Step 1: Defining a Struct
```go
type Order struct {
    id        string
    amount    float32
    status    string
    CreatedAt time.Time
}
```
* `type <Name> struct` defines a new custom data type.
* Notice: field name comes first, type comes second (`id string`).

---

### Step 2: Value Receiver vs. Pointer Receiver

When attaching a function to a struct:

```go
// 1. POINTER RECEIVER (*Order) -> Use when you want to MODIFY the struct
func (o *Order) ChangeStatus(newStatus string) {
    o.status = newStatus // Updates the real instance in RAM!
}

// 2. VALUE RECEIVER (Order) -> Use for READ-ONLY methods
func (o Order) GetAmount() float32 {
    return o.amount // Receives a copy; original is safe from changes
}
```

> **Rule of thumb**: If the method needs to change any field in the struct, or if the struct is large, use a pointer receiver (`*Order`).

---

### Step 3: Instantiating (Creating) Structs

#### Method A: Named Fields (Most Common)
```go
myOrder := Order{
    id:     "ord_123",
    amount: 50.00,
    status: "pending",
}
// Unmentioned fields (like CreatedAt) automatically get safe zero values (e.g. 0, "", false).
```

#### Method B: Zero-value Declaration
```go
var myOrder Order // Every field is initialized to its zero-value
myOrder.id = "ord_456"
```

#### Method C: Pointer to Struct
```go
orderPtr := &Order{id: "ord_789"}
// Go automatically dereferences pointers with dot notation:
orderPtr.status = "shipped" // No need for orderPtr->status or (*orderPtr).status
```

---

### Step 4: Memory Independence
Every struct variable created with `{}` or `var` has its own memory location:
```go
order1 := Order{id: "1"}
order2 := Order{id: "2"}

order1.status = "paid" // Does NOT affect order2 in any way
```

---

### Step 5: Default Zero-Values
If you omit a field when initializing a struct, Go automatically assigns safe zero-values (no random garbage memory):

| Type | Zero Value |
| :--- | :--- |
| `string` | `""` (empty string) |
| `int`, `float32`, `float64` | `0` or `0.0` |
| `bool` | `false` |
| `time.Time` | `0001-01-01 00:00:00 +0000 UTC` (zero-time) |
| pointer, slice, map | `nil` |

---

### Step 6: The Constructor Pattern (`New...`)
Go has no `__init__` or `constructor` keyword. Instead, use a factory function named `New<Type>`:

```go
func NewOrder(id string, amount float32, status string) *Order {
    return &Order{
        id:        id,
        amount:    amount,
        status:    status,
        CreatedAt: time.Now(), // automatically set defaults
    }
}

// Usage:
myOrder := NewOrder("ord_1", 29.99, "pending")
```
* Why return `*Order`? Returning a pointer avoids copying the struct in memory and allows you to immediately call pointer-receiver methods on it.

---

### Step 7: Anonymous Structs (One-Off Configs / Settings)
When you only need a struct **once** (like reading database credentials, script flags, or a quick test case), you don't need to define a global `type ... struct`. You can create an **anonymous struct** on the fly:

```go
// Define fields AND assign values in one expression:
dbConfig := struct {
    host string
    port int
    user string
}{
    host: "localhost",
    port: 5432,
    user: "postgres",
}

fmt.Println(dbConfig.host) // "localhost"
```
* **Python equivalent**: Like creating a temporary inline dictionary or one-off Pydantic `BaseSettings` that is only used inside a single initialization function.

---

### Step 8: Nested Structs & Embedding (Composition over Inheritance)

Go does not have classes or `extends`. Instead, it uses composition:

#### 1. Standard Nesting ("Has-A" relationship):
```go
type Customer struct {
    Name  string
    Email string
}

type Order struct {
    ID     string
    Client Customer // regular nested struct
}

// Access with chained dots:
// order.Client.Name
```

#### 2. Struct Embedding (Automatic Field Promotion):
```go
type Order struct {
    ID string
    Customer // Notice: NO field name! Just the type name.
}

// Customer's fields are automatically promoted to Order:
// order.Name  (Direct access without typing order.Customer.Name!)
```

#### 3. Custom Behavior (Method Shadowing / Overriding):
If the outer struct defines a method with the same name, it shadows the inner method:
```go
// Customer's default method:
func (c Customer) Greet() { fmt.Println("Hello from Customer") }

// Order's custom behavior:
func (o Order) Greet() { fmt.Println("Hello from Order!") }

// o.Greet()          -> Calls Order's custom Greet
// o.Customer.Greet() -> Calls Customer's original Greet
```

#### 4. Strictly One-Way Rules:
1. **Inner is isolated**: `Customer` has no knowledge of `Order` and cannot access `Order`'s fields.
2. **No Subtype Polymorphism**: `Order` is **not** a `Customer` in Go's type system. You cannot pass an `Order` to a function expecting `Customer` (must explicitly pass `order.Customer`).
3. **No Dynamic Dispatch**: If an inner method calls another inner method, it will never dynamically jump to an outer method.





