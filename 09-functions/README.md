# Functions

## What this topic is
Functions in Go are first-class citizens that can be assigned to variables, passed as arguments, and returned from other functions. This topic covers function definition, parameters, return values (including multiple returns), function types, and closures.

## When to use it
- Organizing code into reusable blocks
- Breaking down complex problems into smaller, manageable pieces
- Creating abstractions and hiding implementation details
- Passing behavior as parameters (functional programming patterns)
- Creating factory functions that return other functions
- Event handlers and callbacks

## Minimal syntax example
```go
package main

import "fmt"

// Function with parameters and return value
func add(a int, b int) int {
    return a + b
}

// Function with multiple return values
func swap(x, y string) (string, string) {
    return y, x
}

// Function that returns another function (closure)
func makeGreeter(prefix string) func(name string) string {
    return func(name string) string {
        return prefix + " " + name
    }
}

func main() {
    // Calling functions
    sum := add(3, 4)
    fmt.Println("3 + 4 =", sum)
    
    // Multiple return values
    a, b := swap("hello", "world")
    fmt.Println(a, b) // world hello
    
    // Function returning function
    greeter := makeGreeter("Hello")
    fmt.Println(greeter("Go")) // Hello Go
}
```

## Common mistakes
- Forgetting that Go requires explicit return statements (no implicit return)
- Not naming return values when you plan to use naked return (can decrease readability)
- Forgetting that function types must match exactly when passing functions as parameters
- Not understanding that closures capture variables by reference, not by value
- Trying to return a function that references local variables without understanding closure semantics
- Forgetting that multiple return values must all be captured or explicitly ignored with blank identifier

## Command to run it
```bash
go run function.go
```