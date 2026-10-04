package main

import "fmt"

func main() {
	fmt.Println("=== If/Else Statements in Go ===")

	fmt.Println("\n1. Basic if-else:")
	age := 7
	if age >= 18 {
		fmt.Printf("Age %d: Person is an adult\n", age)
	} else if age >= 12 {
		fmt.Printf("Age %d: Person is a teenager\n", age)
	} else {
		fmt.Printf("Age %d: Person is a child\n", age)
	}

	fmt.Println("\n2. If with logical operators:")
	role := "admin"
	hasPermission := false
	if role == "admin" || hasPermission {
		fmt.Println("Access granted: user is admin or has permission")
	} else {
		fmt.Println("Access denied")
	}

	fmt.Println("\n3. Short statement before condition (initiative statement):")
	if age := 20; age >= 18 {
		fmt.Printf("Age %d from initiative statement: Adult\n", age)
	} else {
		fmt.Printf("Age %d from initiative statement: Minor\n", age)
	}
	// age is not accessible here (out of scope)

	fmt.Println("\n4. Nested if statements:")
	score := 85
	if score >= 0 {
		if score >= 90 {
			fmt.Printf("Score %d: Excellent (A)\n", score)
		} else if score >= 80 {
			fmt.Printf("Score %d: Good (B)\n", score)
		} else {
			fmt.Printf("Score %d: Needs improvement (C or below)\n", score)
		}
	} else {
		fmt.Println("Invalid score")
	}

	fmt.Println("\n5. Error handling pattern:")
	result, err := divide(10, 0)
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
	} else {
		fmt.Printf("Division result: %.2f\n", result)
	}
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}