package main

import "fmt"

func main() {
	fmt.Println("=== Basic Values and Types ===")

	// Integers
	fmt.Println("Integers:")
	fmt.Println("  1 + 1 =", 1+1)
	fmt.Println("  5 * 3 =", 5*3)
	fmt.Println("  10 / 3 =", 10/3) // Integer division
	fmt.Println("  10 % 3 =", 10%3) // Modulus
	fmt.Println()

	// Floating point
	fmt.Println("Floating point:")
	fmt.Println("  7.0 / 3.0 =", 7.0/3.0)
	fmt.Printf("  %%0.2f format: %.2f\n", 7.0/3.0)
	fmt.Println()

	// Booleans
	fmt.Println("Booleans:")
	fmt.Println("  true && false =", true && false)
	fmt.Println("  true || false =", true || false)
	fmt.Println("  !true =", !true)
	fmt.Println()

	// Strings
	fmt.Println("Strings:")
	name := "Rishav Raj"
	fmt.Println("  name =", name)
	fmt.Println("  len(name) =", len(name))
	fmt.Println("  name[0] =", string(name[0])) // First character as string
	fmt.Println("  name + \" Go\" =", name + " Go")
	fmt.Println()

	// Characters/runes
	fmt.Println("Characters and runes:")
	char := 'A' // rune (int32)
	fmt.Printf("  'A' = %v (type %T)\n", char, char)
	fmt.Printf("  'A' + 1 = %c\n", 'A'+1)
	fmt.Println()

	// Constants
	fmt.Println("Constants:")
	const Pi = 3.14159
	const (
		StatusOK   = 200
		StatusNotFound = 404
		MaxUsers   = 1000
	)
	fmt.Printf("  Pi = %.5f\n", Pi)
	fmt.Printf("  StatusOK = %d\n", StatusOK)
	fmt.Printf("  MaxUsers = %d\n", MaxUsers)
	fmt.Println()

	// Type conversions
	fmt.Println("Type conversions:")
	integer := 42
	float := float64(integer)
	fmt.Printf("  %d as float64: %v\n", integer, float)

	float2 := 3.14
	int2 := int(float2) // Truncates toward zero
	fmt.Printf("  %.2f as int: %d\n", float2, int2)
	fmt.Println()
}