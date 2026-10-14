# Maps

## What this topic is
Maps in Go are built-in associative data types (also called hash tables or dictionaries in other languages). They provide an unordered collection of key-value pairs where you can look up values by their keys. Maps are reference types, meaning when you pass a map to a function, you're passing a reference to the same underlying data structure.

## When to use it
- When you need to look up values by keys (like a dictionary or hash table)
- Implementing caches, lookup tables, or memoization
- Counting occurrences of items (like word frequency)
- Representing sparse data
- When you need fast lookups, insertions, and deletions (average O(1) time complexity)
- Building sets (by using map[key]bool or map[key]struct{})

## Minimal syntax example
```go
package main

import "fmt"

func main() {
    // Creating maps
    userIDs := make(map[string]int)         // empty map
    capitals := map[string]string{          // map literal
        "France": "Paris",
        "Japan":  "Tokyo",
    }
    
    // Adding and updating
    userIDs["alice"] = 1001
    userIDs["bob"] = 1002
    userIDs["alice"] = 2001  // update
    
    // Accessing values (comma-ok idiom to check existence)
    if id, ok := userIDs["alice"]; ok {
        fmt.Println("Alice's ID:", id)
    } else {
        fmt.Println("Alice not found")
    }
    
    // Checking if key exists
    if _, exists := userIDs["charlie"]; !exists {
        fmt.Println("Charlie not in map")
    }
    
    // Getting map length
    fmt.Println("Number of users:", len(userIDs))
    
    // Deleting entries
    delete(userIDs, "bob")
    
    // Iterating over maps (order is randomized)
    for name, id := range userIDs {
        fmt.Printf("%s: %d\n", name, id)
    }
    
    // Different map types
    httpStatus := map[int]string{
        200: "OK",
        404: "Not Found",
        500: "Internal Server Error",
    }
    scores := map[string][]int{
        "Alice": {95, 87, 92},
        "Bob":   {88, 91},
    }
}
```

## Common mistakes
- Forgetting that maps are reference types (passing to functions passes a reference)
- Not understanding that map iteration order is randomized (not guaranteed to be consistent)
- Trying to use nil maps (will cause panic when assigning)
- Forgetting to check if a key exists when the zero value is a valid value (use comma-ok idiom)
- Assuming maps are comparable (they're not, except for comparing to nil)
- Not realizing that map values are not addressable (can't take address of map element)
- Confusing maps with slices or arrays
- Forgetting that you can't get the capacity of a map (only length with len())

## Command to run it
```bash
go run maps.go
```