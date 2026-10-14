# Generics

## What this topic is
Generics in Go (introduced in Go 1.18) allow you to write functions and data structures that work with different types while maintaining type safety. They enable code reuse without sacrificing the compile-time type checking that makes Go robust.

## When to use it
- When you need to write functions that work with multiple types (like sorting, mapping, filtering)
- Creating reusable data structures (like stacks, queues, linked lists) that can hold any type
- When you want to avoid code duplication while maintaining type safety
- Building libraries or utilities that need to be flexible about the types they work with
- When interface{} would work but you want to preserve type information

## Minimal syntax example
```go
package main

import "fmt"

// Generic function that works with any type
func PrintAny[T any](items []T) {
    for _, item := range items {
        fmt.Print(item, " ")
    }
    fmt.Println()
}

// Generic function with constraints (only works with types that support +)
type Addable interface {
    int | int8 | int16 | int32 | int64 |
        uint | uint8 | uint16 | uint32 | uint64 | uintptr |
        float32 | float64 |
        string
}

func Add[T Addable](a, b T) T {
    return a + b
}

// Generic type (like a stack)
type Stack[T any] struct {
    elements []T
}

func (s *Stack[T]) Push(elem T) {
    s.elements = append(s.elements, elem)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.elements) == 0 {
        var zero T // zero value for type T
        return zero, false
    }
    index := len(s.elements) - 1
    elem := s.elements[index]
    s.elements = s.elements[:index]
    return elem, true
}

func main() {
    // Using generic functions
    numbers := []int{1, 2, 3}
    strings := []string{"a", "b", "c"}
    
    PrintAny(numbers)  // T inferred as int
    PrintAny(strings)  // T inferred as string
    
    // Using constrained generic function
    fmt.Println(Add(2, 3))         // 5 (int)
    fmt.Println(Add(2.5, 3.7))     // 6.2 (float64)
    fmt.Println(Add("hi", " there")) // "hi there" (string)
    
    // Using generic type
    intStack := Stack[int]{}
    intStack.Push(10)
    intStack.Push(20)
    fmt.Println(intStack.Pop()) // 20, true
}
```

## Common mistakes
- Forgetting that generic functions need type parameters in brackets: `func MyFunc[T any](param T)`
- Not understanding constraints (which types are allowed) and when to use them
- Confusing the `any` keyword (alias for `interface{}`) with specific types
- Not realizing that the zero value of a type parameter T is the zero value of its underlying type
- Forgetting that Go can often infer type arguments, so you don't always need to specify them
- Using generics when a simple interface or concrete type would be clearer
- Not understanding that generic types in method receivers work differently (the receiver can be generic, but method type parameters are not allowed)
- Expecting generics to work like in other languages (Go's approach is unique with type sets and constraints)

## Command to run it
```bash
go run generics.go
```