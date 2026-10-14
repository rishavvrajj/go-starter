# Structs

## What this topic is
Structs in Go are composite data types that group together zero or more values of different types. They are similar to classes in object-oriented languages but without inheritance. Go uses composition and interfaces to achieve code reuse and polymorphism.

## When to use it
- When you need to group related data together
- Representing real-world entities (like User, Product, Order)
- Creating complex data structures
- When you want to attach methods to data types
- Building APIs or data transfer objects
- When you need named fields instead of relying on position (unlike tuples)

## Minimal syntax example
```go
package main

import (
    "fmt"
    "time"
)

// Define a struct
type User struct {
    ID        int
    Username  string
    Email     string
    CreatedAt time.Time
    IsActive  bool
}

// Method with pointer receiver (can modify the struct)
func (u *User) Deactivate() {
    u.IsActive = false
}

// Method with value receiver (read-only)
func (u User) IsAdult() bool {
    // In real code, you'd check age from a birthdate field
    return true // placeholder
}

// Constructor function (idiomatic in Go)
func NewUser(id int, username, email string) *User {
    return &User{
        ID:        id,
        Username:  username,
        Email:     email,
        CreatedAt: time.Now(),
        IsActive:  true,
    }
}

func main() {
    // Initialize using struct literal
    user1 := User{
        ID:        1,
        Username:  "john_doe",
        Email:     "john@example.com",
        CreatedAt: time.Now(),
        IsActive:  true,
    }
    
    // Initialize using constructor function
    user2 := NewUser(2, "jane_smith", "jane@example.com")
    
    // Access fields
    fmt.Printf("User1: %+v\n", user1)
    fmt.Printf("User2 ID: %d\n", user2.ID)
    
    // Call methods
    user2.Deactivate()
    fmt.Printf("User2 active: %v\n", user2.IsActive)
    
    // Embedded structs (composition)
    type Address struct {
        Street string
        City   string
    }
    
    type Employee struct {
        User    // embedded struct (promotes fields)
        Address // embedded struct
        Salary  float64
    }
    
    emp := Employee{
        User: User{ID: 3, Username: "bob", Email: "bob@example.com", IsActive: true},
        Address: Address{Street: "123 Main St", City: "Anytown"},
        Salary: 50000,
    }
    
    // Access promoted fields
    fmt.Printf("Employee name: %s\n", emp.Username) // from embedded User
    fmt.Printf("Employee city: %s\n", emp.City)     // from embedded Address
}
```

## Common mistakes
- Forgetting that structs are value types (assignment creates a copy)
- Not understanding the difference between pointer and value receivers
- Using structs when you need reference semantics (consider pointers or interfaces)
- Forgetting to initialize all fields (they get zero values)
- Using positional initialization for large structs (hard to maintain)
- Not realizing that embedded structs promote their fields to the outer struct
- Trying to compare structs that contain non-comparable types (like maps, functions)
- Not understanding zero values of struct fields (0, false, "", nil pointers, etc.)

## Command to run it
```bash
go run struct.go
```