# Switch Statements

## What this topic is
Switch statements in Go provide a clean way to write long if-else chains when comparing a value against multiple possible values. Go's switch is more flexible than in many languages - it can work without an expression (like if-else chains), supports multiple values per case, and has type switches.

## When to use it
- When you have multiple conditions based on the same variable
- Implementing state machines
- Parsing input or commands
- Type checking with type switches
- When you want cleaner code than multiple if-else if statements

## Minimal syntax example
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    day := time.Now().Weekday()
    
    switch day {
    case time.Saturday, time.Sunday:
        fmt.Println("Weekend!")
    case time.Monday:
        fmt.Println("Start of work week")
    default:
        fmt.Println("Weekday")
    }
    
    // Switch without expression (like if-else chain)
    score := 85
    switch {
    case score >= 90:
        fmt.Println("Excellent!")
    case score >= 80:
        fmt.Println("Good!")
    default:
        fmt.Println("Keep practicing")
    }
}
```

## Common mistakes
- Forgetting that Go's switch automatically breaks after each case (no need for break statements)
- Using fallthrough when you don't actually need it (rarely needed in Go)
- Not realizing that switch cases can have multiple values separated by commas
- Forgetting that switch works with any comparable type (not just integers)
- Not knowing about type switches for checking interface types

## Command to run it
```bash
go run switch.go
```