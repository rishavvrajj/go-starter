package main

import (
	"fmt"
	"time"
)

// --- 1. BASIC STRUCTS ---

// Customer represents a person who places an order.
type Customer struct {
	Name  string
	Phone string
}

// Order represents a purchase.
// Notice how it embeds 'Customer'. This is Go's way of doing composition instead of inheritance.
type Order struct {
	ID        string
	Amount    float32
	Status    string
	CreatedAt time.Time
	Customer  // Embedded struct (Anonymous field)
}

// --- 2. METHODS ---

// changeStatus uses a POINTER RECEIVER (*Order).
// This allows the method to modify the original struct.
// If we used (o Order), it would modify a copy and the original would remain unchanged.
func (o *Order) changeStatus(status string) {
	o.Status = status
}

// getAmount uses a VALUE RECEIVER (Order).
// This is appropriate for read-only methods because it doesn't need to modify the struct.
func (o Order) getAmount() float32 {
	return o.Amount
}

// --- 3. CONSTRUCTORS ---

// NewOrder is a "factory function".
// In Go, we often use functions starting with 'New' to initialize structs with defaults.
func NewOrder(id string, amount float32, status string) *Order {
	return &Order{
		ID:        id,
		Amount:    amount,
		Status:    status,
		CreatedAt: time.Now(),
	}
}

func main() {
	fmt.Println("--- 1. Struct Initialization ---")

	// Option A: Named fields (Recommended) - Clear and order-independent.
	order1 := Order{
		ID:     "1",
		Amount: 123.32,
		Status: "received",
	}
	fmt.Printf("Order 1: %+v\n", order1)

	// Option B: Positional fields - Must provide ALL fields in exact order.
	// Not recommended for large structs.
	order2 := Order{"2", 234.56, "shipped", time.Now(), Customer{"Jane", "555-0101"}}
	fmt.Printf("Order 2: %+v\n", order2)

	fmt.Println("\n--- 2. Zero Values ---")

	// If we omit fields, they get their "zero value" (empty string, 0, nil, etc.)
	order3 := Order{ID: "3"}
	fmt.Printf("Order 3 (incomplete): %+v\n", order3) // Amount will be 0, Status will be ""

	fmt.Println("\n--- 3. Methods & Receivers ---")

	order1.changeStatus("Paid") // Modifies order1 because of the pointer receiver
	fmt.Println("New Status:", order1.Status)
	fmt.Println("Amount:", order1.getAmount())

	fmt.Println("\n--- 4. Factory Functions ---")

	order4 := NewOrder("4", 50.00, "Pending")
	fmt.Printf("Order from factory: %+v\n", order4)

	fmt.Println("\n--- 5. Anonymous Structs ---")

	// Used for one-off data structures (e.g., JSON responses or test data)
	language := struct {
		Name   string
		IsGood bool
	}{"Golang", true}
	fmt.Printf("Anonymous Struct: %+v\n", language)

	fmt.Println("\n--- 6. Struct Embedding (Composition) ---")

	cust := Customer{Name: "Rishav", Phone: "123456789"}

	order5 := Order{
		ID:       "5",
		Amount:   99.99,
		Status:   "Confirmed",
		Customer: cust, // Injecting the Customer struct
	}

	// FIELD PROMOTION: Because Customer is embedded anonymously,
	// we can access Customer fields directly on the Order object.
	fmt.Println("Customer Name via Promotion:", order5.Name) // Instead of order5.Customer.Name
	fmt.Println("Customer Phone via Promotion:", order5.Phone)
}