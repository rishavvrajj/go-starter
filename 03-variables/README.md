# Variables

## What this topic is
Variables in Go are used to store data that can be changed during program execution. This topic covers variable declaration, initialization, scope, and zero values.

## When to use it
- Storing data that needs to be modified
- Creating counters, flags, or temporary values
- Function parameters and return values
- Loop variables
- Any data that changes during program execution

## Minimal syntax example
```go
package main

import "fmt"

func main() {
    // Basic declaration with type
    var name string = "Alice"
    
    // Declaration with inferred type
    var age = 30
    
    // Short declaration (inside functions only)
    city := "New York"
    
    // Multiple variables
    var x, y int = 10, 20
    a, b, c := true, 42, "text"
    
    fmt.Printf("Name: %s, Age: %d, City: %s\n", name, age, city)
    fmt.Printf("x: %d, y: %d\n", x, y)
    fmt.Printf("a: %v, b: %d, c: %s\n", a, b, c)
}
```

## Common mistakes
- Using `var` outside of functions for short declaration (`:=` only works inside functions)
- Forgetting that variables have block scope
- Not understanding zero values (variables get their type's zero value when declared without initialization)
- Trying to re-declare a variable in the same scope (can use `:=` only if at least one new variable is declared)
- Confusing `=` (assignment) with `:=` (short declaration)

## Command to run it
```bash
go run variable.go
```