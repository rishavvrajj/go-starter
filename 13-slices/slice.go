package main

import "fmt"

func main() {
	fmt.Println("Slice section")

	// 1. Creating slices
	// Using make: make([]type, length, capacity)
	var nums = make([]int, 2, 4) // length=2, capacity=4, elements are zero values (0)
	fmt.Println("nums after make:", nums, "len:", len(nums), "cap:", cap(nums))

	// Using slice literal
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("fruits:", fruits, "len:", len(fruits), "cap:", cap(fruits))

	// Nil slice (zero value for slice type)
	var nilSlice []int
	fmt.Println("nilSlice == nil:", nilSlice == nil, "len:", len(nilSlice), "cap:", cap(nilSlice))

	// Empty slice (not nil, but length 0)
	emptySlice := []int{}
	fmt.Println("emptySlice == nil:", emptySlice == nil, "len:", len(emptySlice), "cap:", cap(emptySlice))

	// 2. Appending to slices
	fmt.Println("\n--- Appending ---")
	nums = append(nums, 10, 20, 30) // append multiple values
	fmt.Println("After appending 10,20,30:", nums, "len:", len(nums), "cap:", cap(nums))

	// Appending a slice (using ... to expand)
	moreNums := []int{40, 50}
	nums = append(nums, moreNums...)
	fmt.Println("After appending moreNums:", nums, "len:", len(nums), "cap:", cap(nums))

	// 3. Slicing (creating a view of an existing slice)
	fmt.Println("\n--- Slicing ---")
	s1 := fruits[1:3] // from index 1 (inclusive) to 3 (exclusive) -> ["banana", "cherry"]
	fmt.Println("fruits[1:3]:", s1)

	s2 := fruits[:2] // first two elements
	fmt.Println("fruits[:2]:", s2)

	s3 := fruits[1:] // from index 1 to the end
	fmt.Println("fruits[1:]:", s3)

	// 4. Modifying a slice modifies the underlying array (if within capacity)
	fmt.Println("\n--- Modification ---")
	s4 := nums[2:4] // shares underlying array with nums
	fmt.Println("Before modifying s4: nums=", nums, "s4=", s4)
	s4[0] = 99 // changes the third element of nums
	fmt.Println("After s4[0]=99: nums=", nums, "s4=", s4)

	// 5. Copying slices (to avoid sharing underlying array)
	fmt.Println("\n--- Copying ---")
	orig := []int{1, 2, 3, 4, 5}
	copySlice := make([]int, len(orig))
	copy(copySlice, orig) // dst, src
	fmt.Println("Original:", orig)
	fmt.Println("Copy:", copySlice)
	// Modify copy, original unchanged
	copySlice[0] = 999
	fmt.Println("After changing copySlice[0]:")
	fmt.Println("Original:", orig)
	fmt.Println("Copy:", copySlice)

	// 6. Iteration
	fmt.Println("\n--- Iteration ---")
	fmt.Println("Using for loop with index:")
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("  index: %d, value: %s\n", i, fruits[i])
	}

	fmt.Println("Using range:")
	for i, v := range fruits {
		fmt.Printf("  index: %d, value: %s\n", i, v)
	}

	// 7. Variadic functions and append
	fmt.Println("\n--- Variadic ---")
	numbers := []int{1, 2, 3}
	// We can pass a slice to a variadic function by expanding with ...
	sum := add(numbers...) // same as add(1,2,3)
	fmt.Println("Sum of", numbers, "=", sum)

	// 8. Multidimensional slices (slice of slices)
	fmt.Println("\n--- Multidimensional slice ---")
	// Create a 2D slice (3 rows, 4 columns)
	table := make([][]int, 3)
	for i := range table {
		table[i] = make([]int, 4)
		for j := range table[i] {
			table[i][j] = i*4 + j + 1 // fill with 1,2,3,...
		}
	}
	fmt.Println("2D slice:")
	for _, row := range table {
		fmt.Println("  ", row)
	}

	// 9. Built-in functions: len, cap, append, copy
	fmt.Println("\n--- Built-in functions ---")
	fmt.Println("len(nums) =", len(nums))
	fmt.Println("cap(nums) =", cap(nums))
	// Note: cap is the maximum length the slice can grow to without reallocating
}

// Variadic function example
func add(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}