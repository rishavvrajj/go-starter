# Interfaces

## What this topic is
Interfaces in Go are contracts that specify what methods a type must have. They enable polymorphism - different types can be used interchangeably if they implement the same interface. A type implements an interface by implementing its methods; there's no explicit declaration needed.

## When to use it
- When you want to achieve polymorphism (different types behaving similarly)
- For dependency injection and loose coupling
- When writing testable code (using mock implementations)
- Creating abstractions that can work with different underlying types
- Building plugins or extensible systems
- When you need to group types by behavior rather than by data structure

## Minimal syntax example
```go
package main

import "fmt"

// Define an interface
type Speaker interface {
    Speak() string
}

// Concrete types that implement the interface
type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}

// Function that works with the interface type
func makeItSpeak(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    dog := Dog{}
    cat := Cat{}
    
    // Different types can be passed to the same function
    makeItSpeak(dog) // Woof!
    makeItSpeak(cat) // Meow!
    
    // Interface satisfaction is implicit - no need to declare implementation
    var _ Speaker = dog // This compiles if dog implements Speaker
    var _ Speaker = cat // This compiles if cat implements Speaker
}
```

## Common mistakes
- Forgetting that interface satisfaction is implicit (no implements keyword)
- Not understanding that a type must implement ALL methods in the interface
- Confusing interface values with the concrete types they hold
- Forgetting that interface methods can have pointer or value receivers
- Not realizing that you can compare interfaces to nil, but not interface-to-interface (except for comparable types)
- Using interfaces when concrete types would be simpler and clearer
- Forgetting that the zero value of an interface is nil
- Not understanding that calling a method on a nil interface causes a panic

## Command to run it
```bash
go run interface.go
```