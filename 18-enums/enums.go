package main

import (
	"fmt"
)

// --- 1. STRING-BASED ENUMS ---

// In Go, there's no built-in enum keyword, but we can create type-safe enums
// using custom types with const blocks.

// OrderStatus represents the possible states of an order.
// By defining it as a distinct type (not just string), we get compile-time safety.
type OrderStatus string

// Define the possible values as constants.
// Using string constants makes them readable and serializable (e.g., to JSON).
const (
	OrderReceived  OrderStatus = "received"
	OrderConfirmed OrderStatus = "confirmed"
	OrderPrepared  OrderStatus = "prepared"
	OrderShipped   OrderStatus = "shipped"
	OrderDelivered OrderStatus = "delivered"
	OrderCancelled OrderStatus = "cancelled"
)

// --- 2. INTEGER-BASED ENUMS WITH IOTA ---

// For numeric enums (like error codes, priorities, etc.), we use iota.
// iota is a preprocessor counter that starts at 0 and increments by 1 each line.

// PriorityLevel represents the urgency of a task.
// Using iota gives us automatic sequential values.
type PriorityLevel int

const (
	Low PriorityLevel = iota // 0
	Medium                   // 1
	High                     // 2
	Critical                 // 3
)

// --- 3. STRINGER INTERFACE FOR CUSTOM STRING REPRESENTATION ---

// Sometimes we want integer enums but still need readable string representations.
// We can implement the String() method from the fmt.Stringer interface.

// Category represents a product category.
type Category int

const (
	CategoryUnknown Category = iota
	CategoryElectronics
	CategoryBooks
	CategoryClothing
	CategoryHome
)

// String returns the string representation of a Category.
// This is called automatically when printing with %s or fmt.Print.
func (c Category) String() string {
	return [...]string{
		"Unknown",
		"Electronics",
		"Books",
		"Clothing",
		"Home",
	}[c]
}

// --- 4. USING ENUMS IN SWITCH STATEMENTS ---

// ProcessOrder handles an order based on its status.
// Using a switch on enum values is type-safe and exhaustive (if we handle all cases).
func ProcessOrder(status OrderStatus) string {
	switch status {
	case OrderReceived:
		return "Order received, preparing payment..."
	case OrderConfirmed:
		return "Payment confirmed, preparing shipment..."
	case OrderPrepared:
		return "Order prepared, awaiting shipment..."
	case OrderShipped:
		return "Order shipped, tracking updated..."
	case OrderDelivered:
		return "Order delivered successfully!"
	case OrderCancelled:
		return "Order cancelled by customer."
	default:
		// This case handles any unexpected values (though with typed constants, this shouldn't happen)
		return fmt.Sprintf("Unknown status: %s", status)
	}
}

// CalculateDiscount calculates discount based on priority.
// Demonstrates using integer enums in switches.
func CalculateDiscount(level PriorityLevel) float32 {
	switch level {
	case Low:
		return 0.0 // No discount
	case Medium:
		return 0.05 // 5% discount
	case High:
		return 0.10 // 10% discount
	case Critical:
		return 0.20 // 20% discount
	default:
		return 0.0
	}
}

// --- 5. ZERO VALUE AND VALIDATION ---

// The zero value of an enum is important to understand.
// For string-based enums, zero value is empty string "".
// For integer-based enums, zero value is 0 (usually the first constant).

// ValidateOrderStatus checks if a status is valid.
// This is useful when receiving status from external sources (e.g., JSON, user input).
func ValidateOrderStatus(status OrderStatus) bool {
	switch status {
	case OrderReceived, OrderConfirmed, OrderPrepared, OrderShipped, OrderDelivered, OrderCancelled:
		return true
	default:
		return false
	}
}

// --- 6. BEST PRACTICES DEMONSTRATION ---

func main() {
	fmt.Println("=== Go Enums Learning Guide ===")

	// --- String-based Enums ---
	fmt.Println("--- 1. String-based Enums ---")
	fmt.Printf("OrderReceived: %s\n", OrderReceived)
	fmt.Printf("OrderConfirmed: %s\n", OrderConfirmed)
	fmt.Printf("OrderPrepared: %s\n", OrderPrepared)
	fmt.Printf("OrderShipped: %s\n", OrderShipped)
	fmt.Printf("OrderDelivered: %s\n", OrderDelivered)
	fmt.Printf("OrderCancelled: %s\n", OrderCancelled)
	fmt.Println()

	// --- Using in Functions ---
	fmt.Println("--- 2. Using Enums in Functions ---")
	fmt.Println(ProcessOrder(OrderReceived))
	fmt.Println(ProcessOrder(OrderConfirmed))
	fmt.Println(ProcessOrder(OrderShipped))
	fmt.Println(ProcessOrder(OrderDelivered))
	fmt.Println()

	// --- Integer Enums with iota ---
	fmt.Println("--- 3. Integer Enums (iota) ---")
	fmt.Printf("Low: %d\n", Low)
	fmt.Printf("Medium: %d\n", Medium)
	fmt.Printf("High: %d\n", High)
	fmt.Printf("Critical: %d\n", Critical)
	fmt.Println()

	// --- Stringer Interface ---
	fmt.Println("--- 4. Stringer Interface (custom string conversion) ---")
	fmt.Printf("CategoryElectronics: %s\n", CategoryElectronics)
	fmt.Printf("CategoryBooks: %s\n", CategoryBooks)
	fmt.Printf("CategoryHome: %s\n", CategoryHome)
	fmt.Printf("Zero value Category: %s\n", CategoryUnknown) // Shows "Unknown"
	fmt.Println()

	// --- Using in Switches ---
	fmt.Println("--- 5. Enums in Switch Statements ---")
	fmt.Printf("Discount for Low priority: %.0f%%\n", CalculateDiscount(Low)*100)
	fmt.Printf("Discount for Medium priority: %.0f%%\n", CalculateDiscount(Medium)*100)
	fmt.Printf("Discount for High priority: %.0f%%\n", CalculateDiscount(High)*100)
	fmt.Printf("Discount for Critical priority: %.0f%%\n", CalculateDiscount(Critical)*100)
	fmt.Println()

	// --- Zero Value and Validation ---
	fmt.Println("--- 6. Zero Values and Validation ---")
	var zeroStatus OrderStatus // Zero value is empty string
	fmt.Printf("Zero value OrderStatus: '%s' (empty string)\n", zeroStatus)
	fmt.Printf("Is zero status valid? %t\n", ValidateOrderStatus(zeroStatus))
	fmt.Printf("Is OrderReceived valid? %t\n", ValidateOrderStatus(OrderReceived))
	fmt.Println()

	// --- Benefits Summary ---
	fmt.Println("--- 7. Why Use Enums in Go? ---")
	fmt.Println("✓ Type safety: Prevents mixing different status types")
	fmt.Println("✓ Readability: OrderConfirmed is clearer than \"confirmed\"")
	fmt.Println("✓ Refactoring safety: IDE can rename all usages")
	fmt.Println("✓ Exhaustive switches: Compiler helps catch missing cases")
	fmt.Println("✓ Zero value awareness: Understand default values")
	fmt.Println("✓ Stringer interface: Get readable names for integers")
	fmt.Println()

	// --- Comparison with "magic strings" ---
	fmt.Println("--- 8. Anti-pattern: Magic Strings (avoid this!) ---")
	// This is what we're avoiding:
	magicStatus := "confirmed" // No type safety, typo-prone
	// If we wrote "comfirmed" by mistake, compiler wouldn't catch it
	// With enums: OrderConfirmed vs typo would be compile error
	fmt.Printf("Magic string (unsafe): %s\n", magicStatus)
	fmt.Printf("Enum version (safe): %s\n", OrderConfirmed)
}