# Variadic Functions

## What this topic is
Variadic functions in Go are functions that can accept a variable number of arguments. They are declared using ellipsis (...) before the type of the last parameter. Inside the function, the variadic parameter behaves like a slice of that type.

## When to use it
- Functions that need to accept any number of arguments (like fmt.Print)
- When you want to provide a flexible API that can handle different numbers of inputs
- Mathematical operations that can work on any number of values (sum, max, min, etc.)
- String concatenation or formatting functions
- Any situation where the number of inputs is not fixed

## Minimal syntax example
```go
package main

import "fmt"

// Variadic function that sums integers
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

// Variadic function that concatenates strings with a separator
func join(separator string, parts ...string) string {
    result := ""
    for i, part := range parts {
        if i > 0 {
            result += separator
        }
        result += part
    }
    return result
}

func main() {
    // Call with different numbers of arguments
    fmt.Println(sum())           // 0
    fmt.Println(sum(1, 2, 3))    // 6
    fmt.Println(sum(1, 2, 3, 4, 5)) // 15
    
    // Pass a slice to a variadic function (using ... to expand)
    numbers := []int{10, 20, 30}
    fmt.Println(sum(numbers...)) // 60
    
    // String concatenation
    fmt.Println(join("-", "apple", "banana", "cherry")) // apple-banana-cherry
}
```

## Common mistakes
- Forgetting that variadic parameters must be the last parameter in the function signature
- Not realizing that inside the function, the variadic parameter is a slice
- Forgetting to use `...` when passing a slice to a variadic function
- Confusing variadic functions with optional parameters (Go doesn't have optional parameters)
- Not understanding that you can pass zero arguments to a variadic function
- Using variadic functions when a simple slice parameter would be clearer

## Command to run it
```bash
go run variadic.go
```