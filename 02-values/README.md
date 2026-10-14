# Values and Types

## What this topic is
Go is a statically typed language, meaning every variable has a type that is known at compile time. This topic covers the basic value types in Go: integers, floating-point numbers, booleans, strings, and runes (character codes).

## When to use it
- Understanding how to work with different data types
- Performing arithmetic operations
- Working with text and character data
- Using constants for fixed values
- Converting between types

## Minimal syntax example
```go
package main

import "fmt"

func main() {
    // Integers
    var age int = 25
    
    // Floating point
    var height float64 = 5.9
    
    // Boolean
    var isStudent bool = true
    
    // String
    var name string = "Go Developer"
    
    // Constants
    const Pi = 3.14159
    
    fmt.Printf("Name: %s, Age: %d, Height: %.1f, Student: %v\n", name, age, height, isStudent)
    fmt.Printf("Pi: %.2f\n", Pi)
}
```

## Common mistakes
- Forgetting that Go is statically typed (variables must have a type)
- Using integer division when you want floating-point results (e.g., 10/3 = 3, not 3.33)
- Confusing runes (int32) with bytes/strings
- Not understanding that string indexing returns bytes, not characters
- Forgetting to use the correct verb in Printf (%d for int, %f for float64, %s for string, %v for default format)

## Command to run it
```bash
go run values.go
```