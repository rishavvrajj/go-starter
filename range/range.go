package main

import "fmt"

func main() {
	fmt.Println("Range keyword examples")
	fmt.Println("=====================")

	// 1. Range over slices/arrays
	fmt.Println("\n1. Range over slice:")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("  Index: %d, Value: %d\n", index, value)
	}

	// Same with array (just change the type to [5]int)
	fmt.Println("\n   Same with array:")
	numArray := [5]int{10, 20, 30, 40, 50}
	for index, value := range numArray {
		fmt.Printf("  Index: %d, Value: %d\n", index, value)
	}

	// 2. Range over maps
	fmt.Println("\n2. Range over map:")
	scores := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 92,
	}
	for name, score := range scores {
		fmt.Printf("  %s: %d\n", name, score)
	}
	// Note: Map iteration order is randomized

	// 3. Range over strings (Unicode aware)
	fmt.Println("\n3. Range over string (Unicode):")
	text := "Hello 世界" // Contains ASCII and Chinese characters
	for i, r := range text {
		fmt.Printf("  Byte index: %d, Character: %q, Unicode: U+%04X\n", i, r, r)
	}
	// Note: 'i' is the byte index, not character index

	// 4. Range with only index or only value (using blank identifier)
	fmt.Println("\n4. Range with blank identifier:")
	// Only need the index
	fmt.Println("   Only indices:")
	for i := range numbers {
		fmt.Printf("  Index: %d\n", i)
	}
	// Only need the value
	fmt.Println("   Only values:")
	for _, value := range numbers {
		fmt.Printf("  Value: %d\n", value)
	}

	// 5. Modifying elements while ranging (only works for slices/arrays/maps, not strings)
	fmt.Println("\n5. Modifying elements while ranging:")
	// For slice: we can modify the original slice
	fmt.Println("   Before:", numbers)
	for i := range numbers {
		numbers[i] *= 2 // double each value
	}
	fmt.Println("   After  :", numbers)

	// Reset for next example
	numbers = []int{10, 20, 30, 40, 50}

	// For map: we can modify the map
	fmt.Println("\n   Map modification:")
	fmt.Println("   Before:", scores)
	for key := range scores {
		scores[key] += 5 // add 5 to each score
	}
	fmt.Println("   After  :", scores)

	// 6. Range over nil map, slice, etc. (safe, produces zero iterations)
	fmt.Println("\n6. Range over nil:")
	var nilSlice []int
	var nilMap map[string]int
	fmt.Printf("   nilSlice length: %d, iterations: ", len(nilSlice))
	count := 0
	for range nilSlice {
		count++
	}
	fmt.Printf("%d\n", count)

	fmt.Printf("   nilMap length: %d, iterations: ", len(nilMap))
	count = 0
	for range nilMap {
		count++
	}
	fmt.Printf("%d\n", count)

	// 7. Range over channels (example with a simple channel)
	fmt.Println("\n7. Range over channel:")
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	ch <- 30
	close(ch) // Important: range on channel waits until it's closed
	fmt.Println("   Values from channel:")
	for value := range ch {
		fmt.Printf("  Received: %d\n", value)
	}

	// 8. Range over struct? (Not directly possible, but we can range over fields via reflection or by converting to map)
	// Instead, let's show ranging over a slice of structs
	fmt.Println("\n8. Range over slice of structs:")
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 30},
		{"Bob",   25},
		{"Charlie", 35},
	}
	for _, p := range people {
		fmt.Printf("  %s is %d years old\n", p.Name, p.Age)
	}

	// 9. Range with reassignment (reusing variables)
	fmt.Println("\n9. Range with reassignment:")
	// We can reuse i and v from previous loops
	i, v := 0, 0 // reset
	for i, v := range numbers {
		// i and v are new variables in this block scope
		// But we can't use them outside the loop if declared here
		// Instead, let's just use the loop variables
		fmt.Printf("  i=%d, v=%d\n", i, v)
	}
	// Outside the loop, i and v are the ones we reset above (0,0)
	fmt.Printf("   After loop: i=%d, v=%d\n", i, v)
}