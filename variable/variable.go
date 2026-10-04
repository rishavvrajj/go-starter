package main

import "fmt"

func main() {
	fmt.Println("=== Variable Declaration and Initialization ===")

	// Basic variable declaration with type
	var name string = "golang"
	fmt.Printf("name (declared with type): %v\n", name)

	// Variable declaration with inferred type
	var language = "Go"
	fmt.Printf("language (inferred type): %v\n", language)

	// Short declaration (only inside functions)
	version := 1.22
	fmt.Printf("version (short declaration): %v\n", version)

	// Multiple variables in one line
	var x, y int = 10, 20
	fmt.Printf("x, y: %d, %d\n", x, y)

	// Multiple variables with inferred types
	a, b, c := true, 42, "text"
	fmt.Printf("a, b, c: %v, %d, %s\n", a, b, c)

	// Blank identifier (discard values)
	_, value := getPair()
	fmt.Printf("Discarded first, kept second: %d\n", value)

	fmt.Println()
	fmt.Println("=== Variable Scope ===")

	// Block scope
	if true {
		innerVar := "I'm in a block"
		fmt.Println(innerVar)
	}
	// fmt.Println(innerVar) // This would cause an error - not accessible outside block

	fmt.Println()
	fmt.Println("=== Zero Values ===")

	// Zero values for different types
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string
	var zeroPtr *int

	fmt.Printf("Zero int: %d\n", zeroInt)
	fmt.Printf("Zero float64: %f\n", zeroFloat)
	fmt.Printf("Zero bool: %v\n", zeroBool)
	fmt.Printf("Zero string: '%q'\n", zeroString)
	fmt.Printf("Zero pointer: %v\n", zeroPtr)

	fmt.Println()
	fmt.Println("=== Constants ===")

	const pi = 3.14159
	const statusOK = 200
	const (
		statusCreated = 201
		statusAccepted = 202
	)

	fmt.Printf("Pi: %.5f\n", pi)
	fmt.Printf("Status OK: %d\n", statusOK)
	fmt.Printf("Status Created: %d\n", statusCreated)
	fmt.Printf("Status Accepted: %d\n", statusAccepted)
}

func getPair() (int, int) {
	return 100, 200
}