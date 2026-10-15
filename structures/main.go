package main

import (
	"fmt"
	"time"
)

// ============================================================
// STEP 1: DEFINE THE STRUCT (The Blueprint)
// ============================================================
// A struct groups related fields into a single custom type.
// Syntax: type <Name> struct { <field> <type> }
type Order struct {
	id        string
	amount    float32
	status    string
	CreatedAt time.Time // Defaults to zero-value (year 0001) if omitted
}

// ============================================================
// STEP 2: ATTACH A METHOD (The Receiver)
// ============================================================
// Go has no classes. We attach functions to structs using a "receiver":
//
//   func (o *Order) ChangeStatus(newStatus string)
//         ^^^^^^^^
//       The Receiver (tells Go which struct owns this method)
//
// 💡 Pointer receiver (*Order) vs Value receiver (Order):
// - *Order (with *) = POINTER receiver. Points to the real instance in RAM.
//   Modifications (o.status = ...) update the actual struct.
// - Order  (no *)   = VALUE receiver. Gets a duplicate copy.
//   Modifications would be lost as soon as the function exits.
func (o *Order) ChangeStatus(newStatus string) {
	o.status = newStatus // Modifies the original instance directly in memory
}

// ============================================================
// STEP 3: CONSTRUCTOR PATTERN (The "New..." Factory Function)
// ============================================================
// Go does NOT have constructors (like `__init__` in Python or `Order()` in C++).
// Instead, Go uses a clean, standard convention: a function named `New<StructName>`:
//   1. Takes the parameters needed to create the struct.
//   2. Can set defaults automatically (like CreatedAt = time.Now()).
//   3. Usually returns a pointer (*Order) to avoid copying memory.
func NewOrder(id string, amount float32, status string) *Order {
	return &Order{
		id:        id,
		amount:    amount,
		status:    status,
		CreatedAt: time.Now(), // auto-set the creation timestamp!
	}
}

func main() {
	// ============================================================
	// STEP 4: INSTANTIATE MANUALLY (Struct Literal)
	// ============================================================
	// Syntax: TypeName{ field: value }
	// Any fields not listed (like CreatedAt) automatically get safe zero values.
	myorder := Order{
		id:     "12asda",
		amount: 45.12,
		status: "received",
	}

	// ============================================================
	// STEP 5: CALL THE METHOD
	// ============================================================
	// Call methods using dot notation (.).
	// Go automatically passes the memory address (&myorder) behind the scenes.
	myorder.ChangeStatus("confirmed")

	fmt.Println("Full struct:", myorder)

	// ============================================================
	// STEP 6: ACCESS INDIVIDUAL FIELDS
	// ============================================================
	// Read or write individual fields using dot notation (.).
	fmt.Println("Order Amount:", myorder.amount)
	fmt.Println("Order Status:", myorder.status)

	// ============================================================
	// STEP 7: CREATE ANOTHER INSTANCE (Independent Memory)
	// ============================================================
	// myorder2 is a completely separate allocation in memory.
	// Changing myorder will never affect myorder2.
	myorder2 := Order{
		id:     "2",
		amount: 34.23,
		status: "pending",
	}
	fmt.Println("Second order:", myorder2)

	// ============================================================
	// STEP 8: DEFAULT ZERO-VALUES (Safe Initialization)
	// ============================================================
	// In Go, if you don't set a field, it automatically receives
	// its type's default "zero-value" (no random garbage memory like in C!):
	//   - string:   "" (empty string)
	//   - numbers:  0 or 0.0
	//   - bool:     false
	//   - time:     0001-01-01 00:00:00 (zero-time)
	//   - pointers: nil
	//
	// Notice above in myorder and myorder2: we never set `CreatedAt`,
	// so Go automatically filled it with its zero-value: 0001-01-01.

	// ============================================================
	// STEP 9: USING THE CONSTRUCTOR FUNCTION
	// ============================================================
	// Instead of manual struct literals, we call NewOrder().
	// It automatically fills in CreatedAt with the current time!
	myorder3 := NewOrder("3", 99.99, "delivered")
	fmt.Println("Constructor order (Step 9):", myorder3)

	// ============================================================
	// STEP 10: ANONYMOUS STRUCTS (One-Off Configs / Settings)
	// ============================================================
	// In FastAPI/Python, you might use Pydantic `BaseSettings` for database config.
	// In Go, if data is ONLY USED ONCE (like a DB config, test payload, or script setting),
	// you don't even need to create a named `type ... struct`.
	//
	// You define the shape AND assign the values all in one shot:
	//   variable := struct { <fields> }{ <values> }
	dbConfig := struct {
		host string
		port int
		user string
	}{
		host: "localhost",
		port: 5432,
		user: "postgres",
	}

	fmt.Printf("DB Config (Step 10): %s:%d (User: %s)\n", dbConfig.host, dbConfig.port, dbConfig.user)
}
