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
