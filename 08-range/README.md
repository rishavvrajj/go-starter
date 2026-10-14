# Range Keyword

## What this topic is
The `range` keyword in Go is used to iterate over elements in various data structures including arrays, slices, maps, channels, and strings. It simplifies iteration by providing both index and value (or key and value) in each iteration.

## When to use it
- Iterating over slices or arrays when you need both index and value
- Iterating over maps to get key-value pairs
- Processing each character in a string (Unicode-aware)
- Receiving values from a channel until it's closed
- When you want clean, readable iteration code without managing indices manually

## Minimal syntax example
```go
package main

import "fmt"

func main() {
    // Range over slice
    numbers := []int{10, 20, 30}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
    
    // Range over map
    scores := map[string]int{"Alice": 95, "Bob": 87}
    for name, score := range scores {
        fmt.Printf("%s: %d us\n", name, score)
    }
    
    // Range over string (Unicode aware)
    text := "Hello 世界"
    for i, r := range text {
        fmt.Printf("%d: %q (U+%04X)\n", i, r, r)
    }
    
    // Only need values (ignore index with blank identifier)
    for _, value := range numbers {
        fmt.Println("Value:", value)
    }
    
    // Only need index (ignore value with blank identifier)
    for i := range numbers {
        fmt.Println("Index:", i)
    }
}
```

## Common mistakes
- Forgetting that `range` on strings gives you byte index and runes (Unicode characters), not byte index and bytes
- Not understanding that map iteration order is randomized (not guaranteed to be consistent)
- Trying to modify a string while ranging over it (strings are immutable in Go)
- Forgetting to close channels when ranging over them (range will block forever waiting for more values)
- Not realizing that `range` on nil slices, maps, or channels is safe and produces zero iterations
- Confusing the two values returned by `range` for different data structures:
  - Arrays/slices: index, value
  - Maps: key, value
  - Strings: byte index, rune value
  - Channels: value only
  - Maps with only index: for key := range map
  - Maps with only value: for _, value := range map

## Command to run it
```bash
go run range.go
```