# Enums

## What this topic is
Go doesn't have a built-in enum keyword, but we can create type-safe enums using custom types with const blocks. Enums (enumerations) are a way to define a set of named constants that represent related values, making code more readable and type-safe.

## When to use it
- When you have a fixed set of related constants (like status codes, states, options)
- When you want type safety to prevent using wrong values
- When you want to make code more readable (OrderConfirmed vs "confirmed")
- When you want to avoid "magic strings" or "magic numbers" in your code
- When you need to serialize/deserialize values (string-based enums work well with JSON)
- When you want exhaustive switch statements (compiler can help detect missing cases)

## Minimal syntax example
```go
package main

import "fmt"

// String-based enum (good for readability and serialization)
type OrderStatus string

const (
    OrderReceived  OrderStatus = "received"
    OrderConfirmed OrderStatus = "confirmed"
    OrderShipped   OrderStatus = "shipped"
    OrderDelivered OrderStatus = "delivered"
)

// Integer-based enum using iota (good for sequential values)
type Priority int

const (
    Low Priority = iota // 0
    Medium              // 1
    High                // 2
    Critical            // 3
)

// Implement Stringer interface for custom string representation of integer enums
func (p Priority) String() string {
    return [...]string{"Low", "Medium", "High", "Critical"}[p]
}

func main() {
    // Using string-based enum
    var status OrderStatus = OrderConfirmed
    fmt.Println("Order status:", status) // Output: confirmed
    
    // Using integer-based enum
    var p Priority = High
    fmt.Println("Priority:", p)        // Output: High (thanks to Stringer)
    fmt.Println("Priority value:", int(p)) // Output: 2
    
    // Using in switch statements
    switch status {
    case OrderReceived:
        fmt.Println("Processing new order")
    case OrderConfirmed:
        fmt.Println("Processing confirmed order")
    }
}
```

## Common mistakes
- Forgetting that you need to define a custom type (type OrderStatus string) for type safety
- Using untyped constants (const OrderReceived = "received") which don't provide enum-like type safety
- Not understanding that iota resets to 0 whenever the const keyword appears
- Forgetting to implement the Stringer interface if you want nice string representations of integer enums
- Using string-based enums when you actually need integer values (or vice versa)
- Not validating enum values when they come from external sources (like JSON)
- Trying to compare enums of different types (won't compile, which is good!)
- Forgetting that the zero value of an enum is meaningful (empty string for string-based, 0 for integer-based)

## Command to run it
```bash
go run enums.go
```