package main

import "fmt"

// This file demonstrates closures in Go:
// - What closures are and how they work
// - How closures capture variables from their surrounding scope
// - Practical examples of closures
// - Closures as function returns and parameters
// - Common closure patterns

func main() {
	fmt.Println("=== Understanding Closures in Go ===")

	// Basic closure example: a counter
	fmt.Println("\n1. Basic Counter Closure:")
	increment := makeCounter()
	fmt.Printf("   increment() = %d\n", increment()) // 1
	fmt.Printf("   increment() = %d\n", increment()) // 2
	fmt.Printf("   increment() = %d\n", increment()) // 3

	// Creating another independent counter
	fmt.Println("\n2. Independent Closures:")
	increment2 := makeCounter()
	fmt.Printf("   increment2() = %d\n", increment2()) // 1 (starts fresh)
	fmt.Printf("   increment() = %d\n", increment())   // 4 (continues from before)

	// Closure with parameters
	fmt.Println("\n3. Closure with Parameters:")
	addAdder := makeAdder(5)
	fmt.Printf("   addAdder(3) = %d\n", addAdder(3))  // 8 (5+3)
	fmt.Printf("   addAdder(10) = %d\n", addAdder(10)) // 15 (5+10)

	multiplier := makeMultiplier(3)
	fmt.Printf("   multiplier(4) = %d\n", multiplier(4)) // 12 (3*4)
	fmt.Printf("   multiplier(7) = %d\n", multiplier(7)) // 21 (3*7)

	// Closure capturing multiple variables
	fmt.Println("\n4. Closure Capturing Multiple Variables:")
	limitedAccumulator := makeLimitedAccumulator(10) // limit of 10
	for i := 0; i < 5; i++ {
		fmt.Printf("   After adding %d: %d\n", i, limitedAccumulator(i))
	}
	// Trying to exceed the limit
	fmt.Printf("   Trying to add 5 (would exceed limit): %d\n", limitedAccumulator(5))

	// Using closures as function parameters
	fmt.Println("\n5. Closures as Function Parameters:")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("   Original numbers: %v\n", numbers)
	transformed := transformNumbers(numbers, func(x int) int {
		return x * x
	})
	fmt.Printf("   After squaring: %v\n", transformed)

	transformed2 := transformNumbers(numbers, func(x int) int {
		return x + 10
	})
	fmt.Printf("   After adding 10: %v\n", transformed2)

	// Practical example: memoization with closure
	fmt.Println("\n6. Practical Example - Memoization:")
	fibonacci := memoizeFibonacci()
	fmt.Printf("   fibonacci(10) = %d\n", fibonacci(10)) // 55
	fmt.Printf("   fibonacci(15) = %d\n", fibonacci(15)) // 610
	fmt.Printf("   fibonacci(10) again = %d\n", fibonacci(10)) // 55 (cached)
}

// makeCounter returns a function that increments and returns a counter.
// This demonstrates a closure capturing the count variable.
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// makeAdder returns a function that adds a fixed value to its input.
// The closure captures the valueToAdd variable.
func makeAdder(valueToAdd int) func(int) int {
	return func(x int) int {
		return x + valueToAdd
	}
}

// makeMultiplier returns a function that multiplies its input by a fixed value.
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// makeLimitedAccumulator returns a function that accumulates values
// but stops accepting new values once a limit is reached.
func makeLimitedAccumulator(limit int) func(int) int {
	sum := 0
	return func(value int) int {
		if sum >= limit {
			return sum // already at or over limit
		}
		sum += value
		if sum > limit {
			sum = limit // cap at limit
		}
		return sum
	}
}

// transformNumbers applies a transformation function to each element in a slice.
// This demonstrates passing a closure as a function parameter.
func transformNumbers(numbers []int, transform func(int) int) []int {
	result := make([]int, len(numbers))
	for i, v := range numbers {
		result[i] = transform(v)
	}
	return result
}

// memoizeFibonacci returns a Fibonacci function that caches results.
// This demonstrates a practical use of closures for memoization.
func memoizeFibonacci() func(int) int {
	cache := map[int]int{0: 0, 1: 1} // seed the cache
	var fib func(int) int

	fib = func(n int) int {
		if val, found := cache[n]; found {
			return val
		}
		// Compute and cache the result
		val := fib(n-1) + fib(n-2)
		cache[n] = val
		return val
	}

	return fib
}