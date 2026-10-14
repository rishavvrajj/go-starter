package main

import "fmt"

func main() {
	fmt.Println("=== For Loops in Go ===")

	fmt.Println("\n1. Basic for loop with condition (like while):")
	i := 1
	for i <= 3 {
		fmt.Println("  i =", i)
		i = i + 1
	}

	fmt.Println("\n2. Classic for loop with init, condition, post:")
	for j := 0; j < 3; j++ {
		fmt.Println("  j =", j)
	}

	fmt.Println("\n3. For loop as infinite loop (use break to exit):")
	count := 0
	for {
		if count >= 3 {
			break
		}
		fmt.Println("  count =", count)
		count++
	}

	fmt.Println("\n4. For loop with continue (skip even numbers):")
	for k := 0; k < 6; k++ {
		if k%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println("  odd k =", k)
	}

	fmt.Println("\n5. For loop over array:")
	numbers := [3]int{10, 20, 30}
	for index, value := range numbers {
		fmt.Printf("  numbers[%d] = %d\n", index, value)
	}

	fmt.Println("\n6. For loop over slice:")
	slice := []string{"apple", "banana", "cherry"}
	for _, fruit := range slice {
		fmt.Println("  fruit =", fruit)
	}

	fmt.Println("\n7. For loop over map:")
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, val := range m {
		fmt.Printf("  m[%s] = %d\n", key, val)
	}

	fmt.Println("\n8. For loop over string (runes):")
	for i, ch := range "Go" {
		fmt.Printf("  %d: '%c' (Unicode: %U)\n", i, ch, ch)
	}

	fmt.Println("\n9. For loop with range and blank index (values only):")
	for _, value := range []int{5, 10, 15} {
		fmt.Println("  value =", value)
	}

	fmt.Println("\n10. For loop with range and blank value (index only):")
	for index := range []string{"a", "b", "c"} {
		fmt.Println("  index =", index)
	}
}