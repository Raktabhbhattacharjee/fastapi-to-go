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

func main() {
	// ============================================================
	// STEP 3: INSTANTIATE (Create an actual instance in memory)
	// ============================================================
	// Syntax: TypeName{ field: value }
	// Any fields not listed (like CreatedAt) automatically get safe zero values.
	myorder := Order{
		id:     "12asda",
		amount: 45.12,
		status: "received",
	}

	// ============================================================
	// STEP 4: CALL THE METHOD
	// ============================================================
	// Call methods using dot notation (.).
	// Go automatically passes the memory address (&myorder) behind the scenes.
	myorder.ChangeStatus("confirmed")

	fmt.Println("Full struct:", myorder)

	// ============================================================
	// STEP 5: ACCESS INDIVIDUAL FIELDS
	// ============================================================
	// Read or write individual fields using dot notation (.).
	fmt.Println("Order Amount:", myorder.amount)
	fmt.Println("Order Status:", myorder.status)

	// ============================================================
	// STEP 6: CREATE ANOTHER INSTANCE (Independent Memory)
	// ============================================================
	// myorder2 is a completely separate allocation in memory.
	// Changing myorder will never affect myorder2.
	myorder2 := Order{
		id:     "2",
		amount: 34.23,
		status: "pending",
	}

	fmt.Println("Second order:", myorder2)
}
