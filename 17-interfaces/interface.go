package main

import (
	"fmt"
)

// --- 1. WHAT IS AN INTERFACE? ---

// In Go, an interface is a contract that specifies what methods a type must have.
// It enables polymorphism - different types can be used interchangeably if they
// implement the same interface.

// PaymentProcessor defines the behavior needed to process payments.
// Any type that has BOTH Pay() and Refund() methods automatically satisfies this interface.
type PaymentProcessor interface {
	// Pay processes a payment of the given amount.
	Pay(amount float32) string

	// Refund processes a refund of the given amount to the specified account.
	// Returns a confirmation message.
	Refund(amount float32, account string) string
}

// --- 2. WHY USE INTERFACES? ---

// Imagine we have different payment gateways (Stripe, PayPal, Razorpay).
// Instead of writing separate code for each, we can define a common interface
// and write our business logic to work with ANY type that implements it.

// PaymentService uses dependency injection to work with any PaymentProcessor.
// This makes our code flexible, testable, and follows the Dependency Inversion Principle.
type PaymentService struct {
	// The service depends on the interface, not concrete implementations.
	processor PaymentProcessor
}

// ProcessPayment delegates the actual work to the injected processor.
// The service doesn't care WHICH processor it is - only that it satisfies the interface.
func (s *PaymentService) ProcessPayment(amount float32) string {
	return s.processor.Pay(amount)
}

// ProcessRefund delegates the refund work to the injected processor.
func (s *PaymentService) ProcessRefund(amount float32, account string) string {
	return s.processor.Refund(amount, account)
}

// --- 3. CONCRETE IMPLEMENTATIONS ---

// Stripe implements the PaymentProcessor interface.
type Stripe struct{}

func (s Stripe) Pay(amount float32) string {
	// In reality, this would call Stripe's API
	return fmt.Sprintf("Charged $%.2f using Stripe", amount)
}

func (s Stripe) Refund(amount float32, account string) string {
	return fmt.Sprintf("Refunded $%.2f to %s via Stripe", amount, account)
}

// PayPal also implements the PaymentProcessor interface.
type PayPal struct{}

func (p PayPal) Pay(amount float32) string {
	return fmt.Sprintf("Charged $%.2f using PayPal", amount)
}

func (p PayPal) Refund(amount float32, account string) string {
	return fmt.Sprintf("Refunded $%.2f to %s via PayPal", amount, account)
}

// Razorpay only implements Pay() - NOT Refund().
// Therefore, it does NOT satisfy the PaymentProcessor interface.
type Razorpay struct{}

func (r Razorpay) Pay(amount float32) string {
	return fmt.Sprintf("Charged $%.2f using Razorpay", amount)
}

// Note: Razorpay does NOT have a Refund method.
// If we tried to use Razorpay where a PaymentProcessor is needed,
// the compiler would give an error: "cannot use Razorpay{} as PaymentProcessor value"

// --- 4. USING INTERFACES IN PRACTICE ---

func main() {
	fmt.Println("--- Interface Polymorphism Demo ---")

	// Create different payment processors
	stripe := Stripe{}
	paypal := PayPal{}
	// razorpay := Razorpay{}

	// Create services using different processors
	stripeService := PaymentService{processor: stripe}
	paypalService := PaymentService{processor: paypal}

	// The same service code works with different processors!
	fmt.Println(stripeService.ProcessPayment(99.99))
	fmt.Println(paypalService.ProcessPayment(149.50))
	fmt.Println()

	fmt.Println("--- Refunds ---")
	fmt.Println(stripeService.ProcessRefund(99.99, "customer@example.com"))
	fmt.Println(paypalService.ProcessRefund(149.50, "customer@example.com"))
	fmt.Println()

	// --- 5. INTERFACE SATISFACTION RULES ---

	// A type only needs to implement ALL methods in the interface to satisfy it.
	// Let's demonstrate what works and what doesn't:

	// ✅ Stripe satisfies PaymentProcessor (has both Pay and Refund)
	var _ PaymentProcessor = stripe // This line compiles - stripe implements the interface
	fmt.Println("✅ Stripe satisfies PaymentProcessor interface")

	// ✅ PayPal satisfies PaymentProcessor (has both Pay and Refund)
	var _ PaymentProcessor = paypal // This line compiles
	fmt.Println("✅ PayPal satisfies PaymentProcessor interface")

	// ❌ Razorpay does NOT satisfy PaymentProcessor (missing Refund method)
	// Uncommenting the line below would cause a compile error:
	// var _ PaymentProcessor = razorpay
	fmt.Println("❌ Razorpay does NOT satisfy PaymentProcessor interface (missing Refund method)")
	fmt.Println()

	// --- 6. BENEFITS OF THIS APPROACH ---

	fmt.Println("--- Benefits ---")
	fmt.Println("1. Loose coupling: PaymentService doesn't need to know about specific gateways")
	fmt.Println("2. Easy testing: We can inject a mock/fake processor for tests")
	fmt.Println("3. Extensibility: Adding new payment gateways requires no changes to PaymentService")
	fmt.Println("4. Compile-time safety: If a type doesn't implement all interface methods, we get a clear error")

	// Demonstrate easy testing with a mock processor
	fmt.Println("--- Testing with a Mock Processor ---")
	mock := MockProcessor{}
	testService := PaymentService{processor: mock}
	fmt.Println(testService.ProcessPayment(42.00))
	fmt.Println(testService.ProcessRefund(42.00, "test@example.com"))
}

// --- 7. MOCK IMPLEMENTATION FOR TESTING ---

// MockProcessor is useful for testing - it lets us simulate behavior without
// calling real payment gateways.
type MockProcessor struct{}

func (m MockProcessor) Pay(amount float32) string {
	return fmt.Sprintf("[MOCK] Payment of $%.2f processed", amount)
}

func (m MockProcessor) Refund(amount float32, account string) string {
	return fmt.Sprintf("[MOCK] Refund of $%.2f to %s processed", amount, account)
}