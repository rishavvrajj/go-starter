package main

import "fmt"

// greet returns a greeting message for the given name.
func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	fmt.Println("Hello, World!")

	// Using the greet function
	message := greet("Go Learner")
	fmt.Println(message)

	// Demonstrating multiple return values
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10/3 = %.2f\n", result)
	}
}

// divide returns the result of dividing a by b, or an error if b is zero.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}