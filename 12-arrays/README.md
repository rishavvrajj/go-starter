# Arrays

## What this topic is
Arrays in Go are fixed-size sequences of elements of the same type. The size of an array is part of its type, meaning [5]int and [10]int are different types. Arrays are value types - when assigned or passed to functions, they are copied.

## When to use it
- When you need a fixed-size collection of elements
- When performance is critical and you want to avoid heap allocation
- When you want to ensure a collection cannot change size
- As building blocks for slices (which are more commonly used)
- When working with multi-dimensional data where dimensions are fixed

## Minimal syntax example
```go
package main

import "fmt"

func main() {
    // Declaration and initialization
    var numbers [3]int              // [0 0 0]
    var letters [3]string           // ["" "" ""]
    var bools [2]bool               // [false false]
    
    // Array literals
    primes := [5]int{2, 3, 5, 7, 11}
    days := [7]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
    
    // Multi-dimensional array
    matrix := [2][3]int{
        {1, 2, 3},
        {4, 5, 6},
    }
    
    fmt.Println("numbers:", numbers)
    fmt.Println("primes:", primes)
    fmt.Println("days:", days)
    fmt.Println("matrix:", matrix)
    
    // Arrays are value types (assignment copies)
    a := [3]int{1, 2, 3}
    b := a // b is a copy of a
    b[0] = 999
    fmt.Println("a:", a) // [1 2 3] (unchanged)
    fmt.Println("b:", b) // [999 2 3]
}
```

## Common mistakes
- Forgetting that array size is part of the type ([3]int and [4]int are different types)
- Trying to resize arrays (use slices instead for dynamic size)
- Not understanding that arrays are value types (passing to functions copies the entire array)
- Confusing arrays with slices (they look similar but behave differently)
- Forgetting that array elements are zero-initialized when declared
- Trying to compare arrays of different sizes (won't compile)

## Command to run it
```bash
go run array.go
```