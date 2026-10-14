# Slices

## What this topic is
Slices in Go are lightweight data structures that provide a view into an underlying array. They are dynamically-sized, flexible views into the elements of an array. Slices are much more commonly used than arrays because they can grow and shrink as needed.

## When to use it
- When you need a collection that can change size
- Most collection-related tasks in Go (the go-to data structure)
- When you want to avoid copying large data structures (slices are reference types)
- Building stacks, queues, or other dynamic data structures
- When working with sequences of data where the size isn't known at compile time

## Minimal syntax example
```go
package main

import "fmt"

func main() {
    // Creating slices
    numbers := []int{1, 2, 3, 4, 5}      // slice literal
    var names []string                   // zero value (nil, length 0)
    empty := []int{}                     // empty slice (not nil, length 0)
    made := make([]float64, 3, 5)        // length=3, capacity=5
    
    // Appending elements
    numbers = append(numbers, 6, 7)      // [1 2 3 4 5 6 7]
    names = append(names, "Alice", "Bob")
    
    // Slicing (creating views)
    firstThree := numbers[0:3]           // [1 2 3]
    lastTwo := numbers[len(numbers)-2:]  // last two elements
    allButFirst := numbers[1:]           // everything except first
    
    // Copying slices (to avoid sharing underlying array)
    original := []int{10, 20, 30}
    copy := make([]int, len(original))
    copy(copy, original)
    copy[0] = 999                        // original unchanged
    
    // Iteration
    for i, v := range numbers {
        fmt.Printf("%d: %d\n", i, v)
    }
    
    // Multidimensional slices (slice of slices)
    table := [][]int{
        {1, 2, 3},
        {4, 5, 6},
    }
}
```

## Common mistakes
- Not understanding the difference between length and capacity
- Forgetting that slices are reference types (modifying a slice affects the underlying array)
- Trying to append to a nil slice without realizing it works (nil slices can be appended to)
- Confusing slicing with copying (slicing creates a view, not a copy)
- Forgetting that slice operations like `numbers[2:4]` don't copy data
- Not realizing that when capacity is exceeded, append allocates a new underlying array
- Using slices when you actually need the immutability/fixed-size of arrays

## Command to run it
```bash
go run slice.go
```