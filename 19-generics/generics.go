package main

import (
	"fmt"
)

// --- 1. WHY GENERICS? ---

// Before Go 1.18, if we wanted to write a function that worked on multiple types,
// we had two options:
//   a) Duplicate the function for each type (violates DRY)
//   b) Use interface{} (losing type safety and requiring type assertions)
//
// Generics solve this by allowing us to write type-safe, reusable code.

// --- 2. BASIC GENERIC FUNCTION ---

// PrintAny prints any slice of elements, regardless of their type.
// The `any` constraint is an alias for `interface{}` and accepts any type.
// Note: We still lose some type safety inside the function (we can't call type-specific methods).
func PrintAny[T any](items []T) {
	for _, item := range items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

// --- 3. CONSTRAINED GENERIC FUNCTIONS ---

// Sometimes we want to restrict the types our generic function accepts.
// We can do this by defining a constraint (as an interface) or using a type set directly.

// Addable is a constraint for types that support the + operator.
// We define it as an interface with a type set: each term is a type that satisfies the constraint.
type Addable interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64 |
		string
}

// Add returns the sum of two values of the same type.
// The type T must satisfy the Addable constraint (i.e., support +).
func Add[T Addable](a, b T) T {
	return a + b
}

// --- 4. USING COMPARABLE ---

// The comparable constraint is predefined and matches all types that can be compared
// with == and != (e.g., integers, floats, strings, pointers, structs containing only comparable types, etc.).

// Index returns the index of the first occurrence of v in s, or -1 if not found.
// It works for any comparable type.
func Index[T comparable](s []T, v T) int {
	for i, item := range s {
		if item == v {
			return i
		}
	}
	return -1
}

// --- 5. CUSTOM CONSTRAINT INTERFACES ---

// We can define our own constraint interfaces for more complex requirements.
// For example, a constraint for types that have an Area() method.

// Shape is a constraint for types that have an Area() method returning float64.
type Shape interface {
	Area() float64
}

// TotalArea calculates the sum of areas of a slice of shapes.
// It works for any type that satisfies the Shape constraint (i.e., has Area() method).
func TotalArea[S Shape](shapes []S) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

// --- 5.5. EXAMPLE SHAPES FOR DEMONSTRATION ---

// Define some shapes that implement Area() method
type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

// Rectangle implements Shape
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Circle implements Shape
func (c Circle) Area() float64 {
	return 3.14159 * c.radius * c.radius
}

// --- 6. GENERIC TYPES ---

// We can also define generic types, like a stack or a linked list.

// Stack is a generic stack (LIFO) that can hold any type.
type Stack[T any] struct {
	elements []T
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(elem T) {
	s.elements = append(s.elements, elem)
}

// Pop removes and returns the top element from the stack.
// It returns a zero value and false if the stack is empty.
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

// --- 7. TYPE INFERENCE ---

// Go can often infer type arguments, so we don't always have to specify them explicitly.
// For example, when calling a generic function, if the compiler can deduce T from the function arguments,
// we can omit the type arguments.

// --- 8. LIMITATIONS AND BEST PRACTICES ---

// - Generics cannot be used in method receivers (but the receiver type can be a generic type, as in Stack[T] above).
// - Avoid overusing generics: if a function only ever works with one or two types, generics might add unnecessary complexity.
// - The zero value of a type parameter T is the zero value of its underlying type (e.g., 0 for int, "" for string, nil for pointers, etc.).

func main() {
	fmt.Println("=== Go Generics Learning Guide ===")

	// --- 1. Basic Generic Function ---
	fmt.Println("--- 1. Basic Generic Function (PrintAny) ---")
	numbers := []int{1, 2, 3, 4, 5}
	strings := []string{"apple", "banana", "cherry"}
	booleans := []bool{true, false, true}

	PrintAny(numbers)  // T is inferred as int
	PrintAny(strings)  // T is inferred as string
	PrintAny(booleans) // T is inferred as bool
	fmt.Println()

	// --- 2. Constrained Generic Function ---
	fmt.Println("--- 2. Constrained Generic Function (Add) ---")
	fmt.Println("Add(2, 3):", Add(2, 3))                               // T inferred as int
	fmt.Println("Add(2.5, 3.7):", Add(2.5, 3.7))                       // T inferred as float64
	fmt.Println("Add(\"hello\", \" world\"):", Add("hello", " world")) // T inferred as string
	fmt.Println()

	// --- 3. Using Comparable ---
	fmt.Println("--- 3. Using Comparable (Index) ---")
	intSlice := []int{10, 20, 30, 40, 50}
	strSlice := []string{"go", "java", "python", "rust"}

	fmt.Println("Index of 30 in intSlice:", Index(intSlice, 30))
	fmt.Println("Index of \"java\" in strSlice:", Index(strSlice, "java"))
	fmt.Println("Index of 99 in intSlice (not found):", Index(intSlice, 99))
	fmt.Println()

	// --- 4. Custom Constraint Interface ---
	fmt.Println("--- 4. Custom Constraint Interface (Shape) ---")
	shapes := []Shape{
		Rectangle{width: 10, height: 5},
		Circle{radius: 7},
		Rectangle{width: 8, height: 8},
	}
	fmt.Printf("Total area of shapes: %.2f\n", TotalArea(shapes))
	fmt.Println()

	// --- 5. Generic Types ---
	fmt.Println("--- 5. Generic Types (Stack) ---")
	// Stack of strings
	stringStack := Stack[string]{}
	stringStack.Push("first")
	stringStack.Push("second")
	stringStack.Push("third")

	fmt.Printf("String stack after pushes: %v\n", stringStack.elements)
	if elem, ok := stringStack.Pop(); ok {
		fmt.Printf("Popped from string stack: %s\n", elem)
	}
	if elem, ok := stringStack.Pop(); ok {
		fmt.Printf("Popped from string stack: %s\n", elem)
	}
	fmt.Printf("String stack after pops: %v\n", stringStack.elements)
	fmt.Println()

	// Stack of integers (as in the original example)
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	fmt.Printf("Int stack: %v\n", intStack.elements)
	if elem, ok := intStack.Pop(); ok {
		fmt.Printf("Popped from int stack: %d\n", elem)
	}
	fmt.Println()

	// --- 6. Type Inference ---
	fmt.Println("--- 6. Type Inference ---")
	// We don't need to specify T when calling PrintAny because the compiler can infer it from the argument.
	PrintAny([]float64{1.1, 2.2, 3.3}) // T inferred as float64

	// However, sometimes we must specify it explicitly, e.g., when the type parameter
	// is not used in the function parameters (less common) or to disambiguate.
	// Example: a function that returns a generic type but takes no parameters of that type.
	// We'll skip that for brevity.

	fmt.Println()

	// --- 7. Limitations and Best Practices ---
	fmt.Println("--- 7. Limitations and Best Practices ---")
	fmt.Println("✓ Generics provide type safety without code duplication")
	fmt.Println("✓ Use constraints to limit types to those that support required operations")
	fmt.Println("✓ The `comparable` constraint is useful for equality checks")
	fmt.Println("✓ Define custom constraints for complex requirements (like methods)")
	fmt.Println("✓ Generic types (like Stack[T]) are useful for reusable data structures")
	fmt.Println("✓ The compiler can often infer type arguments, reducing verbosity")
	fmt.Println("✗ Avoid generics when a simple concrete type or interface{} would suffice")
	fmt.Println("✗ Generics cannot be used in method receivers (but the receiver can be a generic type)")
	fmt.Println()

	// --- 8. Comparison with Pre-Generics Approach ---
	fmt.Println("--- 8. Comparison with Pre-Generics Approach ---")
	fmt.Println("Before generics, to have a stack for multiple types, we might have used:")
	fmt.Println("  type Stack struct { elements []interface{} }")
	fmt.Println("But this loses type safety and requires type assertions.")
	fmt.Println("With generics, Stack[T] is type-safe and efficient.")
}
