package main

import "fmt"

// ============================================================
// STEP 1: DEFINE THE INTERFACE (The Contract / Rule)
// ============================================================
// An interface defines BEHAVIOR (what an object can DO), not data.
// By convention, single-method interfaces often end in "-er" (e.g. Payer, Reader).
// Rule: Anyone who has a method matching this exact signature qualifies!
type Payer interface {
	Pay(amount float32)
}

// ============================================================
// STEP 2: THE CONSUMER / SERVICE (Accepts the Interface)
// ============================================================
// The PaymentService doesn't care which company processes the payment.
// It only holds the `Payer` interface (Dependency Injection!).
type PaymentService struct {
	gateway Payer // 👈 Can hold Razorpay, Stripe, or a Mock test gateway!
}

func (p PaymentService) MakePayment(amount float32) {
	// Calls the Pay method on whatever gateway is plugged in:
	p.gateway.Pay(amount)
}

// ============================================================
// STEP 3: PROVIDERS / IMPLEMENTATIONS (Concrete Structs)
// ============================================================
// Notice: We NEVER write "implements Payer".
// In Go, if a struct has the matching method, it implements the interface AUTOMATICALLY!

// Provider 1: Razorpay
type Razorpay struct{}

func (r Razorpay) Pay(amount float32) {
	fmt.Printf("💳 Razorpay: Processed ₹%.2f successfully\n", amount)
}

// Provider 2: Stripe
type Stripe struct{}

func (s Stripe) Pay(amount float32) {
	fmt.Printf("💵 Stripe: Processed $%.2f successfully\n", amount)
}

// Provider 3: FakeGateway (Used in Production Unit Tests!)
type MockGateway struct{}

func (m MockGateway) Pay(amount float32) {
	fmt.Printf("🧪 Test Mock: Simulated payment of %.2f (no real money charged)\n", amount)
}

func main() {
	// ============================================================
	// STEP 4: PLUG-AND-PLAY IN MAIN (Dependency Injection)
	// ============================================================

	// 1. Pay with Razorpay:
	razorpay := Razorpay{}
	payment1 := PaymentService{gateway: razorpay}
	payment1.MakePayment(500)

	// 2. Easily switch to Stripe (Zero changes to PaymentService logic!):
	stripe := Stripe{}
	payment2 := PaymentService{gateway: stripe}
	payment2.MakePayment(120.50)

	// 3. Testing with Mock Gateway (Production testing pattern):
	mock := MockGateway{}
	testPayment := PaymentService{gateway: mock}
	testPayment.MakePayment(999)
}
