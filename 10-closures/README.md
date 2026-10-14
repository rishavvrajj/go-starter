# Closures

## What this topic is
A closure is a function value that references variables from outside its body. The function can access and modify those variables even after the outer function has finished executing. In Go, closures are created when you define a function inside another function and that inner function references variables from the outer function.

## When to use it
- Creating factory functions that return customized functions
- Implementing callbacks and event handlers
- Data encapsulation and privacy (creating private variables)
- Memoization (caching function results)
- Iterators and generators
- Partial application or currying of functions
- Maintaining state in functional programming patterns

## Minimal syntax example
```go
package main

import "fmt"

// Function that returns a closure (counter)
func makeCounter() func() int {
    count := 0 // This variable is captured by the closure
    return func() int {
        count++
        return count
    }
}

func main() {
    // Create a counter closure
    counter := makeCounter()
    
    // Use the closure - it maintains its state between calls
    fmt.Println(counter()) // 1
    fmt.Println(counter()) // 2
    fmt.Println(counter()) // 3
    
    // Create another independent counter
    counter2 := makeCounter()
    fmt.Println(counter2()) // 1 (starts fresh)
    fmt.Println(counter())  // 4 (continues from where it left off)
}
```

## Common mistakes
- Not understanding that closures capture variables by reference, not by value
- Expecting each closure to have its own copy of captured variables (they share the same variable)
- Forgetting that the captured variables persist as long as the closure is in use
- Creating closures in loops without understanding variable scope (all closures might reference the same loop variable)
- Overusing closures when simpler approaches would work
- Not realizing that closures can cause memory leaks if they hold references to large data structures unnecessarily

## Command to run it
```bash
go run closures.go
```