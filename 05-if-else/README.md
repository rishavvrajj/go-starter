# If/Else Statements

## What this topic is
Conditional statements in Go allow you to execute different code blocks based on whether a condition is true or false. Go's if/else statements are straightforward and similar to other languages, with the addition of an optional initialization statement.

## When to use it
- Making decisions in your code based on conditions
- Validating input data
- Controlling program flow
- Error handling (checking if an error is not nil)
- Implementing business logic that depends on certain criteria

## Minimal syntax example
```go
package main

import "fmt"

func main() {
    age := 20
    
    // Basic if-else
    if age >= 18 {
        fmt.Println("You are an adult")
    } else {
        fmt.Println("You are a minor")
    }
    
    // If with initialization statement
    if score := 85; score >= 90 {
        fmt.Println("Excellent!")
    } else if score >= 80 {
        fmt.Println("Good job!")
    } else {
        fmt.Println("Keep practicing")
    }
}
```

## Common mistakes
- Forgetting braces around code blocks (required in Go, even for single statements)
- Using `=` instead of `==` for comparison (assignment vs equality)
- Not understanding that variables declared in the initialization statement are only in scope within the if/else block
- Complex conditions that are hard to read (consider extracting to a boolean variable)
- Forgetting to handle all cases in if-else chains

## Command to run it
```bash
go run ifelse.go
```