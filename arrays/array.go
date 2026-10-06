package main

import "fmt"

func main() {
	fmt.Println("Array section")

	// Declaration and initialization
	var nums [4]int
	nums[0] = 1
	fmt.Println("nums:", nums) // [1 0 0 0]
	fmt.Println("length:", len(nums)) // 4

	var vals [4]bool
	vals[2] = true
	fmt.Println("vals:", vals) // [false false true false]

	var name [3]string
	name[1] = "golang"
	fmt.Println("name:", name) // [ golang ]

	// Array literals
	numbs := [3]int{1, 2, 3}
	fmt.Println("numbs:", numbs) // [1 2 3]

	// Multi-dimensional array
	twoD := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println("twoD:", twoD) // [[3 4] [5 6]]

	// Arrays are value types: assigning copies the array
	nums2 := nums
	nums2[0] = 99
	fmt.Println("After changing nums2[0]:")
	fmt.Println("nums:", nums)   // [1 0 0 0] (unchanged)
	fmt.Println("nums2:", nums2) // [99 0 0 0]

	// Iteration with for loop
	fmt.Println("\nIterating with for loop:")
	for i := 0; i < len(numbs); i++ {
		fmt.Printf("index: %d, value: %d\n", i, numbs[i])
	}

	// Iteration with range
	fmt.Println("\nIterating with range:")
	for i, v := range numbs {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

	// Comparing arrays (only if same type and length)
	a := [2]int{1, 2}
	b := [2]int{1, 2}
	c := [2]int{2, 1}
	fmt.Println("\nArray comparison:")
	fmt.Println("a == b:", a == b) // true
	fmt.Println("a == c:", a == c) // false

	// Note: Arrays cannot be resized. Use slices for dynamic size.
}