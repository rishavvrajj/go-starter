package main

import "fmt"

func main() {
	// Constants are declared with const and cannot be reassigned
	const name string = "rishav"
	// name = "updated" // This would cause a compile error

	fmt.Println(name)

	// Constants can be typed or untyped
	const pi = 3.14159 // untyped constant
	const piFloat64 float64 = 3.14159 // typed constant

	fmt.Printf("Pi (untyped): %v\n", pi)
	fmt.Printf("Pi (typed): %v\n", piFloat64)

	// Multiple constants using const block
	const (
		StatusOK   = 200
		StatusNotFound = 404
		StatusInternalServerError = 500
	)
	
	fmt.Printf("Status OK: %d\n", StatusOK)
	fmt.Printf("Status Not Found: %d\n", StatusNotFound)
	fmt.Printf("Status Internal Server Error: %d\n", StatusInternalServerError)
}