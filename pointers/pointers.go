package main

import "fmt"

// This file demonstrates pointers in Go:
// - What pointers are and how they work
// - Pointer declaration and initialization
// - Dereferencing pointers to access values
// - Pointers to structs
// - Pointers as function parameters
// - The difference between passing by value and passing by pointer
// - Common pointer patterns and pitfalls

func main() {
	fmt.Println("=== Understanding Pointers in Go ===")

	// Basic pointer declaration and initialization
	fmt.Println("\n1. Basic Pointer Declaration:")
	var p *int // pointer to integer
	fmt.Printf("   p (zero value): %v\n", p) // nil

	i := 42
	p = &i // p now points to i
	fmt.Printf("   i = %d\n", i)
	fmt.Printf("   p = %v (address of i)\n", p)
	fmt.Printf("   *p = %d (value pointed to by p)\n", *p) // dereferencing

	// Modifying value through pointer
	fmt.Println("\n2. Modifying Value Through Pointer:")
	*p = 21
	fmt.Printf("   After *p = 21: i = %d, *p = %d\n", i, *p)

	// Pointer to struct
	fmt.Println("\n3. Pointer to Struct:")
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "Alice", Age: 30}
	fmt.Printf("   person = %+v\n", person)

	personPtr := &person
	fmt.Printf("   personPtr = %v\n", personPtr)
	fmt.Printf("   personPtr.Name = %s\n", personPtr.Name) // automatic dereferencing
	fmt.Printf("   (*personPtr).Age = %d\n", (*personPtr).Age) // explicit dereferencing

	// Modifying struct through pointer
	personPtr.Age = 31
	fmt.Printf("   After modifying through pointer: person = %+v\n", person)

	// Pointers as function parameters
	fmt.Println("\n4. Pointers as Function Parameters:")
	value := 10
	fmt.Printf("   Before updateValue: value = %d\n", value)
	updateValue(&value)
	fmt.Printf("   After updateValue: value = %d\n", value)

	// Function that returns a pointer
	fmt.Println("\n5. Functions Returning Pointers:")
	var num *int
	num = createPointer(42)
	fmt.Printf("   createPointer(42) = %v\n", num)
	fmt.Printf("   *createPointer(42) = %d\n", *num)

	// Demonstrating nil pointers
	fmt.Println("\n6. Nil Pointers:")
	var nilPtr *int
	fmt.Printf("   nilPtr (uninitialized) = %v\n", nilPtr)
	if nilPtr == nil {
		fmt.Println("   nilPtr is nil (safe to check)")
	}

	// Show how to check if pointer is not nil before dereferencing
	i = 42
	validPtr := &i
	fmt.Printf("   validPtr = %v\n", validPtr)
	if validPtr != nil {
		fmt.Printf("   *validPtr = %d (safe to dereference)\n", *validPtr)
	}

	// Pointer vs value semantics
	fmt.Println("\n7. Pointer vs Value Semantics:")
	original := []int{1, 2, 3}
	fmt.Printf("   Original slice: %v\n", original)

	// Passing by value (copy of the slice header)
	sliceCopy := original
	sliceCopy[0] = 99
	fmt.Printf("   After modifying sliceCopy: original = %v, sliceCopy = %v\n", original, sliceCopy)

	// Passing by pointer (reference to the same underlying array)
	slicePtr := &original
	(*slicePtr)[0] = 1 // modify through pointer
	fmt.Printf("   After modifying through pointer: original = %v\n", original)

	// New function example
	fmt.Println("\n8. Practical Example - String Modifier:")
	text := "hello"
	fmt.Printf("   Original text: %s\n", text)
	modifyString(&text)
	fmt.Printf("   After modifyString: %s\n", text)

	// Function that returns a pointer to newly allocated memory
	fmt.Println("\n9. Functions Returning Pointers to New Memory:")
	newInt := newInteger(100)
	fmt.Printf("   newInteger(100) = %v\n", newInt)
	fmt.Printf("   *newInteger(100) = %d\n", *newInt)
}

// updateValue modifies the integer pointed to by p
func updateValue(p *int) {
	*p = 99
}

// createPointer returns a pointer to a new integer with the given value
func createPointer(value int) *int {
	return &value
}

// modifyString modifies the string pointed to by s
func modifyString(s *string) {
	*s = fmt.Sprintf("%s world", *s)
}

// newInteger allocates a new integer on the heap and returns a pointer to it
// Note: In Go, it's safe to return a pointer to a local variable
// because the compiler will allocate it on the heap if it detects escaping
func newInteger(value int) *int {
	return &value
}