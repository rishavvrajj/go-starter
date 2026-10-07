package main

import "fmt"

// This file demonstrates variadic functions in Go:
// - Functions that accept a variable number of arguments
// - How to pass slices to variadic functions
// - Practical examples of variadic functions
// - Variadic functions with different types

// sum calculates the sum of a variable number of integers.
// The ...int parameter allows the function to accept zero or more int arguments.
func sum(nums ...int) int {
	total := 0

	// Range over the variadic parameter (which is a slice)
	for _, num := range nums {
		total += num
	}

	return total
}

// concat strings together with a separator.
// Demonstrates variadic functions with string parameters.
func concat(separator string, parts ...string) string {
	result := ""
	for i, part := range parts {
		if i > 0 {
			result += separator
		}
		result += part
	}
	return result
}

// printValues accepts any number of values of any type using interface{}.
// This shows how to create truly generic variadic functions.
func printValues(values ...interface{}) {
	for i, v := range values {
		fmt.Printf("  [%d] %v (type: %T)\n", i, v, v)
	}
}

// calculate accepts both regular parameters and variadic parameters.
// Regular parameters must come before variadic parameters.
func calculate(operation string, numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	switch operation {
	case "sum":
		total := 0.0
		for _, n := range numbers {
			total += n
		}
		return total
	case "avg":
		total := 0.0
		for _, n := range numbers {
			total += n
		}
		return total / float64(len(numbers))
	case "max":
		max := numbers[0]
		for _, n := range numbers[1:] {
			if n > max {
				max = n
			}
		}
		return max
	case "min":
		min := numbers[0]
		for _, n := range numbers[1:] {
			if n < min {
				min = n
			}
		}
		return min
	default:
		return 0
	}
}

func main() {
	fmt.Println("=== Variadic Functions in Go ===")

	// Basic usage of variadic function
	fmt.Println("\n1. Basic variadic function (sum):")
	result := sum(1, 2, 3, 4, 5)
	fmt.Printf("   sum(1, 2, 3, 4, 5) = %d\n", result)

	// Calling with no arguments (variadic functions can accept zero args)
	fmt.Println("\n2. Variadic function with no arguments:")
	result = sum()
	fmt.Printf("   sum() = %d\n", result)

	// Calling with a single argument
	fmt.Println("\n3. Variadic function with single argument:")
	result = sum(42)
	fmt.Printf("   sum(42) = %d\n", result)

	// Passing a slice to a variadic function (using ... to expand)
	fmt.Println("\n4. Passing a slice to variadic function:")
	nums := []int{10, 20, 30, 40}
	fmt.Printf("   nums = %v\n", nums)
	result = sum(nums...) // Note the ... to expand the slice
	fmt.Printf("   sum(nums...) = %d\n", result)

	// Variadic function with string parameters
	fmt.Println("\n5. Variadic function with strings (concat):")
	message := concat("-", "apple", "banana", "cherry")
	fmt.Printf("   concat('-', 'apple', 'banana', 'cherry') = %s\n", message)

	message2 := concat(" ", "Hello", "world", "from", "Go")
	fmt.Printf("   concat(' ', 'Hello', 'world', 'from', 'Go') = %s\n", message2)

	// Variadic function with interface{} (accepts any type)
	fmt.Println("\n6. Variadic function with interface{} (any type):")
	fmt.Println("   printValues(42, 'hello', 3.14, true):")
	printValues(42, "hello", 3.14, true)

	// Practical example: calculator function
	fmt.Println("\n7. Practical variadic function (calculator):")
	numbers := []float64{10.5, 20.3, 5.2, 8.7}
	fmt.Printf("   Numbers: %v\n", numbers)

	sumResult := calculate("sum", numbers...)
	fmt.Printf("   Sum: %.2f\n", sumResult)

	avgResult := calculate("avg", numbers...)
	fmt.Printf("   Average: %.2f\n", avgResult)

	maxResult := calculate("max", numbers...)
	fmt.Printf("   Maximum: %.2f\n", maxResult)

	minResult := calculate("min", numbers...)
	fmt.Printf("   Minimum: %.2f\n", minResult)

	// Direct function calls with variadic parameters
	fmt.Println("\n8. Direct variadic function calls:")
	fmt.Printf("   calculate('sum', 1, 2, 3, 4, 5) = %.2f\n",
		calculate("sum", 1, 2, 3, 4, 5))
	fmt.Printf("   calculate('avg', 10, 20, 30) = %.2f\n",
		calculate("avg", 10, 20, 30))
}