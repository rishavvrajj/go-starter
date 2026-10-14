# Constants

## What this topic is
Constants in Go are values that cannot be changed after they are created. They are declared with the `const` keyword and can be typed or untyped.

## When to use it
- Values that should never change (like mathematical constants)
- Configuration values that are fixed
- Enumerated values (using iota)
- Performance optimization (constants are evaluated at compile time)
- Making code more readable by giving names to magic numbers

## Minimal syntax example
```go
package main

import "fmt"

func main() {
    // Simple constant
    const Pi = 3.14159
    
    // Typed constant
    const PiFloat64 float64 = 3.14159
    
    // Constant block
    const (
        StatusOK = 200
        StatusNotFound = 404
        StatusInternalServerError = 500
    )
    
    fmt.Printf("Pi: %.5f\n", Pi)
    fmt.Printf("Status OK: %d\n", StatusOK)
}
```

## Common mistakes
- Trying to reassign a constant (results in compile error)
- Confusing constants with variables
- Not understanding the difference between typed and untyped constants
- Forgetting that constants must be assignable (can't be the result of a function call at runtime)
- Using constants when you actually need variables

## Command to run it
```bash
go run constant.go
```