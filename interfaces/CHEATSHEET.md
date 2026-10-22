# Go Interfaces Cheat Sheet

A comprehensive, production-focused guide to understanding and using interfaces in Go (especially coming from Python, C++, and Java).

---

## 1. Quick Concept Map

| Concept | Python / FastAPI | Java / C++ | Go |
| :--- | :--- | :--- | :--- |
| **Definition** | `Protocol` / `ABC` | `interface` / pure virtual class | `type Name interface { ... }` |
| **Implementation** | Implicit (duck typing at runtime) | Explicit: `implements Interface` | **Implicit (duck typing at compile time!)** |
| **Dependency Injection** | `Depends(get_gateway)` | `@Inject` / Spring container | **Constructor injection into struct** |
| **Syntax Keyword** | `class ... (ABC):` | `implements` | **Zero keywords! Just write the method.** |

---

## 2. The Golden Rules of Go Interfaces

1. **Interfaces define BEHAVIOR, not Data**: They only contain method signatures (no fields, no state).
2. **Implicit Implementation**: You never write `implements`. If a struct has matching methods, Go satisfies the interface automatically.
3. **Exact Signature Match**: Method name, parameter types, and return types must match 100%.
4. **The `-er` Naming Convention**: Single-method interfaces end with `-er` (e.g. `Reader`, `Writer`, `Payer`, `Notifier`).

---

## 3. The 4 Production Pillars of Go Interfaces

### Pillar 1: "Accept Interfaces, Return Structs" (The #1 Go Proverb)
* **When accepting input**: Be as flexible as possible $\rightarrow$ accept an `interface`.
* **When returning output**: Be as concrete as possible $\rightarrow$ return a concrete `*struct`.
```go
// Takes flexible interface, returns concrete struct:
func NewPaymentService(gateway Payer) *PaymentService {
    return &PaymentService{gateway: gateway}
}
```

### Pillar 2: "Define Interfaces Where They Are Used" (Consumer, not Producer)
* In Java/C#, the library creating Razorpay also creates `IRazorpay`.
* In **Go**, the **caller (consumer)** defines the interface!
* `PaymentService` is the consumer that needs to make payments, so `PaymentService` defines `type Payer interface`. `Razorpay` doesn't need to know or care about `Payer`!

### Pillar 3: "Keep Interfaces Tiny" (1–2 Methods Max)
* The smaller the interface, the more powerful and reusable it is.
* Standard library examples:
  * `io.Reader`: 1 method (`Read`)
  * `io.Writer`: 1 method (`Write`)
  * `fmt.Stringer`: 1 method (`String`)
* If an interface has 10 methods, it becomes painful to swap or mock.

### Pillar 4: "Inject Via Constructor"
* Instead of creating dependencies secretly inside methods, pass them in through a `New...` factory function:
```go
razorpay := NewRazorpay()
paymentService := NewPaymentService(razorpay) // Injected cleanly!
```

---

## 4. Production Patterns

### Pattern 1: Swappable Infrastructure (The Repository Pattern)
Swap databases (Postgres, Mongo, In-Memory) without changing business logic:

```go
// 1. Contract defined by the consumer:
type UserRepository interface {
    SaveUser(user User) error
    GetUser(id string) (User, error)
}

// 2. Business logic only depends on the interface:
type UserService struct {
    repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
    return &UserService{repo: r}
}

// 3. Swap in main() based on environment:
// db := PostgresRepo{}  // in Production
// db := InMemoryRepo{}  // in Local/Test
// service := NewUserService(db)
```

---

### Pattern 2: Unit Testing with Mocks (Test without Real APIs or DBs)
In production, you never charge real credit cards or hit production databases during tests:

```go
type MockPaymentGateway struct {
    WasCalled bool
}

func (m *MockPaymentGateway) Pay(amount float32) {
    m.WasCalled = true // Record that payment was attempted
}

// In test file:
func TestPaymentFlow(t *testing.T) {
    mock := &MockPaymentGateway{}
    service := NewPaymentService(mock)

    service.MakePayment(50)

    if !mock.WasCalled {
        t.Errorf("Expected payment to be processed!")
    }
}
```

---

### Pattern 3: Standard Library Interfaces (Everywhere in Go)

1. **`fmt.Stringer`** (Like Python's `__str__`):
   ```go
   type Stringer interface {
       String() string
   }
   ```
   If your struct has a `.String() string` method, `fmt.Println(myStruct)` will automatically print your custom formatted string!

2. **`io.Reader`** and **`io.Writer`**:
   Used for files, network sockets, HTTP request bodies, and buffers. Anything with `Read()` can be processed by anything that takes an `io.Reader`.

---

### Summary Checklist

| What to do | Go Code |
| :--- | :--- |
| **Accept interfaces** | Functions & constructors take interfaces (`gateway Payer`) |
| **Return structs** | Constructors return concrete pointers (`*PaymentService`) |
| **Define where used** | The package that *needs* the tool defines the interface |
| **Keep it tiny** | 1 or 2 methods per interface (e.g. `Payer`) |
| **Constructor injection** | `NewService(tool Tool) *Service` |
