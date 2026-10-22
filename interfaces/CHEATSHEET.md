# Go Interfaces Cheat Sheet

A comprehensive, production-focused guide to understanding and using interfaces in Go (especially coming from Python, C++, and Java).

---

## 1. Quick Concept Map

| Concept | Python / FastAPI | Java / C++ | Go |
| :--- | :--- | :--- | :--- |
| **Definition** | `Protocol` / `ABC` | `interface` / pure virtual class | `type Name interface { ... }` |
| **Implementation** | Implicit (duck typing at runtime) | Explicit: `implements Interface` | **Implicit (duck typing at compile time!)** |
| **Dependency Injection** | `Depends(get_gateway)` | `@Inject` / Spring container | **Struct holding the interface** |
| **Syntax Keyword** | `class ... (ABC):` | `implements` | **Zero keywords! Just write the method.** |

---

## 2. The Golden Rules of Go Interfaces

1. **Interfaces define BEHAVIOR, not Data**: They only contain method signatures (no fields, no state).
2. **Implicit Implementation**: You never write `implements`. If a struct has matching methods, Go satisfies the interface automatically.
3. **Exact Signature Match**: Method name, parameter types, and return types must match 100%.
4. **The `-er` Naming Convention**: Single-method interfaces end with `-er` (e.g. `Reader`, `Writer`, `Payer`, `Notifier`).

---

## 3. The 3-Step Production Recipe

### Step 1: Define the Contract (The Interface)
```go
type Payer interface {
    Pay(amount float32)
}
```

### Step 2: Accept the Interface in your Service
```go
type PaymentService struct {
    gateway Payer // 👈 Holds ANY struct that satisfies Payer
}

func (p PaymentService) MakePayment(amount float32) {
    p.gateway.Pay(amount)
}
```

### Step 3: Implement the Structs (Providers)
```go
type Razorpay struct{}
func (r Razorpay) Pay(amount float32) {
    fmt.Println("Paying via Razorpay:", amount)
}

type Stripe struct{}
func (s Stripe) Pay(amount float32) {
    fmt.Println("Paying via Stripe:", amount)
}
```

### Step 4: Inject in `main()`
```go
func main() {
    service := PaymentService{gateway: Razorpay{}}
    service.MakePayment(100)

    // Swap to Stripe with ZERO changes to PaymentService:
    service = PaymentService{gateway: Stripe{}}
    service.MakePayment(200)
}
```

---

## 4. Production Patterns

### Pattern 1: Unit Testing with Mocks (Test without Real APIs or DBs)
In production, you never charge real credit cards or hit production databases during tests. Interfaces make testing effortless:

```go
// Fake mock gateway for tests
type MockPaymentGateway struct {
    WasCalled bool
}

func (m *MockPaymentGateway) Pay(amount float32) {
    m.WasCalled = true // Record that payment was attempted
}

// In test file:
func TestPaymentFlow(t *testing.T) {
    mock := &MockPaymentGateway{}
    service := PaymentService{gateway: mock}

    service.MakePayment(50)

    if !mock.WasCalled {
        t.Errorf("Expected payment to be processed!")
    }
}
```

---

### Pattern 2: "Accept Interfaces, Return Structs" (The #1 Go Proverb)
* **When accepting input**: Be as flexible as possible (accept an `interface`).
* **When returning output**: Be as concrete and specific as possible (return a concrete `struct` or pointer to `struct`).

```go
// Good Go practice:
func ProcessPayment(p Payer) { ... } // Accepts ANY Payer

func NewRazorpay() *Razorpay { ... }  // Returns concrete *Razorpay
```

---

### Pattern 3: Standard Library Interfaces (Everywhere in Go)

Go’s standard library is built on simple 1-method interfaces:

1. **`fmt.Stringer`** (Like Python's `__str__`):
   ```go
   type Stringer interface {
       String() string
   }
   ```
   If your struct has a `.String() string` method, `fmt.Println(myStruct)` will automatically print your custom string!

2. **`io.Reader`** and **`io.Writer`**:
   Used for files, network sockets, HTTP request bodies, and buffers. Anything with `Read()` can be processed by anything that takes an `io.Reader`.

---

### Summary Checklist

| What to do | Go Code |
| :--- | :--- |
| **Define interface** | `type Notifier interface { Send(msg string) }` |
| **Implement interface** | Simply attach `func (s SMS) Send(msg string) { ... }` |
| **Inject dependency** | `type Service struct { n Notifier }` |
| **Check compliance** | Go compiler checks it automatically at compile-time! |
