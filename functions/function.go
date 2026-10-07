package main

import "fmt"

// This file demonstrates various aspects of functions in Go:
// - Basic function definition and calling
// - Multiple return values
// - Function types and passing functions as parameters
// - Functions returning other functions (closures)
// - Anonymous functions

func main() {
	fmt.Println("=== Function Basics ===")

	// Calling a simple function that adds two integers
	result := add(3, 4)
	fmt.Printf("add(3, 4) = %d\n", result)

	// Function returning multiple values
	fmt.Println("\n=== Multiple Return Values ===")
	lang1, lang2, success := getLan()
	fmt.Printf("getLan() returns: %s, %s, %v\n", lang1, lang2, success)

	// Demonstrating blank identifier to ignore unwanted return values
	_, lang3, _ := getLan()
	fmt.Printf("Ignoring first and third values: %s\n", lang3)

	fmt.Println("\n=== Function Types and Parameters ===")

	// Defining an anonymous function inline
	double := func(x int) int {
		return x * 2
	}
	fmt.Printf("double(5) = %d\n", double(5))

	// Passing a function as a parameter to another function
	fmt.Println("Passing function as parameter:")
	processit(double)

	// Function that returns another function
	fmt.Println("\n=== Functions Returning Functions ===")
	squareMaker := makePowerFunction(2)
	cubeMaker := makePowerFunction(3)

	fmt.Printf("squareMaker(4) = %d\n", squareMaker(4)) // 4^2 = 16
	fmt.Printf("cubeMaker(3) = %d\n", cubeMaker(3))   // 3^3 = 27

	// Demonstrating closure - the returned function retains access to variables
	powerOfTwo := makePowerFunction(2)
	powerOfThree := makePowerFunction(3)
	fmt.Printf("powerOfTwo(5) = %d\n", powerOfTwo(5))  // 2^5 = 32
	fmt.Printf("powerOfThree(4) = %d\n", powerOfThree(4)) // 3^4 = 81
}

// add returns the sum of two integers.
func add(a int, b int) int {
	return a + b
}

// getLan returns two programming language names and a success boolean.
// This demonstrates a function with multiple return values of different types.
func getLan() (string, string, bool) {
	return "Go", "JavaScript", true
}

// processit takes a function as a parameter and calls it with the value 1.
// This demonstrates how functions can be passed as arguments to other functions.
func processit(fn func(int) int) {
	result := fn(1)
	fmt.Printf("  Function passed to processit returned: %d\n", result)
}

// makePowerFunction returns a function that raises its input to a specified power.
// This demonstrates functions that return other functions (closures).
func makePowerFunction(power int) func(int) int {
	return func(base int) int {
		result := 1
		for i := 0; i < power; i++ {
			result *= base
		}
		return result
	}
}