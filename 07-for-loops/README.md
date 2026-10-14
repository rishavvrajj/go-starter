# For Loops

## What this topic is
The `for` loop is Go's only looping construct. It can be used in multiple ways: as a traditional for loop with init/condition/post, as a while-like loop with just a condition, or as an infinite loop. Go also uses `for` with `range` to iterate over arrays, slices, maps, channels, and strings.

## When to use it
- Repeating a block of code multiple times
- Iterating over collections (arrays, slices, maps)
- Processing characters in strings
- Reading from channels
- Any situation where you need to execute code repeatedly

## Minimal syntax example
```go
package main

import "fmt"

func main() {
    // Traditional for loop
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
    
    // While-like loop (condition only)
    count := 0
    for count < 5 {
        fmt.Println(count)
        count++
    }
    
    // Infinite loop (use break to exit)
    for {
        fmt.Println("Looping...")
        break // Exit after one iteration
    }
    
    // Range over slice
    fruits := []string{"apple", "banana", "cherry"}
    for index, fruit := range fruits {
        fmt.Printf("%d: %s\n", index, fruit)
    }
}
```

## Common mistakes
- Forgetting that Go doesn't have a `while` keyword (use `for` with condition only)
- Off-by-one errors in loop conditions
- Forgetting to increment/decrement loop variables in condition-only loops
- Not understanding that `range` on maps returns values in random order
- Using `range` on strings gives you runes (Unicode characters), not bytes
- Forgetting that you can use blank identifier `_` to ignore unwanted values from range

## Command to run it
```bash
go run for.go
```