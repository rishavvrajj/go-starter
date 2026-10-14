# Pointers

## What this topic is
Pointers in Go are variables that store memory addresses of other variables. They allow you to indirectly access and modify the value stored at that address. Pointers are essential for passing large data structures efficiently (by reference rather than by value) and for modifying variables inside functions.

## When to use it
- When you need to modify a variable inside a function
- When working with large data structures to avoid copying
- Implementing data structures like linked lists, trees, etc.
- When you need to create optional values (nil pointers indicate absence)
- Working with interfaces and method receivers
- Performance-critical code where avoiding copies matters

## Minimal syntax example
```go
package main

import "fmt"

func main() {
    // Declaration and initialization
    var p *int        // pointer to integer (initially nil)
    i := 42           // regular integer variable
    p = &i            // p now points to i (& gets the address)
    
    // Dereferencing (accessing the value at the address)
    fmt.Println("i =", i)         // 42
    fmt.Println("*p =", *p)       // 42 (value pointed to by p)
    
    // Modifying through pointer
    *p = 21                       // changes the value that p points to
    fmt.Println("After *p = 21:")
    fmt.Println("i =", i)         // 21 (i was changed!)
    fmt.Println("*p =", *p)       // 21
    
    // Pointer to struct
    type Person struct {
        Name string
        Age  int
    }
    
    alice := Person{Name: "Alice", Age: 30}
    alicePtr := &alice
    fmt.Println("alicePtr.Name =", alicePtr.Name) // automatic dereferencing
    alicePtr.Age = 31                             // modify through pointer
    fmt.Println("alice =", alice)                 // {Alice 31}
    
    // Function that modifies via pointer
    value := 10
    updateValue(&value)
    fmt.Println("After updateValue:", value) // 99
}

// Function that takes a pointer parameter
func updateValue(p *int) {
    *p = 99
}
```

## Common mistakes
- Forgetting to initialize pointers (they start as nil)
- Dereferencing a nil pointer (causes runtime panic)
- Not understanding that pointers hold addresses, not values
- Confusing `*` (pointer declaration/dereference) with `&` (address-of operator)
- Forgetting that when you pass a pointer to a function, you're passing a copy of the pointer (but it still points to the same memory)
- Not realizing that Go has automatic dereferencing for struct fields (pointer.Name works the same as (*pointer).Name)
- Using pointers when you don't need them (adds complexity without benefit)
- Forgetting that map lookups and channel operations don't need pointers

## Command to run it
```bash
go run pointers.go
```