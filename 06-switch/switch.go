package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Switch Statements in Go ===")

	fmt.Println("\n1. Basic switch with integer:")
	i := 3
	switch i {
	case 1:
		fmt.Println("  One")
	case 2:
		fmt.Println("  Two")
	case 3:
		fmt.Println("  Three")
	default:
		fmt.Println("  Default case (value not 1, 2, or 3)")
	}

	fmt.Println("\n2. Switch with multiple values per case:")
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("  Weekend!")
	case time.Monday:
		fmt.Println("  Start of the work week")
	default:
		fmt.Println("  Weekday")
	}

	fmt.Println("\n3. Switch with no expression (like if-else chain):")
	score := 85
	switch {
	case score >= 90:
		fmt.Printf("  Score %d: Excellent\n", score)
	case score >= 80:
		fmt.Printf("  Score %d: Good\n", score)
	case score >= 70:
		fmt.Printf("  Score %d: Average\n", score)
	default:
		fmt.Printf("  Score %d: Needs improvement\n", score)
	}

	fmt.Println("\n4. Type switch:")
	var x interface{} = 42
	switch v := x.(type) {
	case int:
		fmt.Printf("  x is an integer: %d\n", v)
	case string:
		fmt.Printf("  x is a string: %s\n", v)
	case bool:
		fmt.Printf("  x is a boolean: %v\n", v)
	default:
		fmt.Printf("  x is another type: %T\n", v)
	}

	fmt.Println("\n5. Switch with fallthrough (use sparingly):")
	number := 2
	switch number {
	case 1:
		fmt.Print("  One ")
		fallthrough
	case 2:
		fmt.Print("Two ")
		fallthrough
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("  Default")
	}

	fmt.Println("\n6. Switch in initialization (short statement):")
	switch ch := getGradeChar('B'); ch {
	case 'A':
		fmt.Println("  Excellent grade")
	case 'B', 'C':
		fmt.Println("  Good grade")
	case 'D':
		fmt.Println("  Needs improvement")
	default:
		fmt.Println("  Invalid grade")
	}
}

func getGradeChar(score rune) rune {
	// Simple function to demonstrate switch in initialization
	switch score {
	case 'A', 'B', 'C', 'D', 'F':
		return score
	default:
		return '?'
	}
}