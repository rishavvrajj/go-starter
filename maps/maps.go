package main

import "fmt"

func main() {
	fmt.Println("Maps section")

	// 1. Creating maps
	// Using make
	userIds := make(map[string]int) // map[string]int
	fmt.Println("Empty map:", userIds)

	// Using map literal
	capitals := map[string]string{
		"France": "Paris",
		"Japan":  "Tokyo",
		"India":  "New Delhi",
	}
	fmt.Println("Capitals:", capitals)

	// Nil map (zero value)
	var nilMap map[string]int
	fmt.Println("nilMap == nil:", nilMap == nil)

	// 2. Adding and updating entries
	fmt.Println("\n--- Adding/Updating ---")
	userIds["alice"] = 1001
	userIds["bob"] = 1002
	userIds["charlie"] = 1003
	fmt.Println("After adding:", userIds)

	// Updating existing key
	userIds["alice"] = 2001 // update value
	fmt.Println("After updating alice:", userIds)

	// 3. Accessing values
	fmt.Println("\n--- Accessing ---")
	// Basic access (returns zero value if key missing)
	fmt.Println("userIds[\"bob\"]:", userIds["bob"]) // 1002
	fmt.Println("userIds[\"dave\"]:", userIds["dave"]) // 0 (zero value for int)

	// Comma ok idiom to check if key exists
	if id, ok := userIds["bob"]; ok {
		fmt.Println("bob's ID:", id)
	} else {
		fmt.Println("bob not found")
	}

	if _, ok := userIds["dave"]; !ok {
		fmt.Println("dave not in map")
	}

	// 4. Map length
	fmt.Println("\n--- Length ---")
	fmt.Println("len(userIds):", len(userIds)) // 3

	// 5. Deleting entries
	fmt.Println("\n--- Deleting ---")
	delete(userIds, "bob")
	fmt.Println("After deleting bob:", userIds)
	fmt.Println("len after delete:", len(userIds)) // 2

	// 6. Clearing a map (Go 1.21+)
	fmt.Println("\n--- Clearing ---")
	clear(userIds) // removes all entries
	fmt.Println("After clear:", userIds)
	fmt.Println("len after clear:", len(userIds)) // 0

	// 7. Iterating over maps
	fmt.Println("\n--- Iteration ---")
	// Repopulate for iteration demo
	userIds["alice"] = 1001
	userIds["charlie"] = 1003
	userIds["dave"] = 1004

	fmt.Println("Iterating over userIds:")
	for key, value := range userIds {
		fmt.Printf("  %s: %d\n", key, value)
	}
	// Note: Map iteration order is randomized (not guaranteed)

	// 8. Nested maps
	fmt.Println("\n--- Nested Maps ---")
	// Map of maps
	studentRecords := map[string]map[string]string{
		"alice": {
			"name":  "Alice Smith",
			"grade": "A",
			"city":  "New York",
		},
		"bob": {
			"name":  "Bob Jones",
			"grade": "B+",
			"city":  "Boston",
		},
	}
	fmt.Println("Student records:", studentRecords)

	// Access nested values
	if record, ok := studentRecords["alice"]; ok {
		if grade, ok := record["grade"]; ok {
			fmt.Println("Alice's grade:", grade)
		}
	}

	// 9. Different map types
	fmt.Println("\n--- Different Types ---")
	// map[int]string (int keys)
	httpStatusText := map[int]string{
		200: "OK",
		404: "Not Found",
		500: "Internal Server Error",
	}
	fmt.Println("HTTP status 404:", httpStatusText[404])

	// map[string][]string (slice values)
	teams := map[string][]string{
		"developers": {"Alice", "Bob", "Charlie"},
		"designers":  {"David", "Eve"},
		"managers":   {"Frank"},
	}
	fmt.Println("Developers team:", teams["developers"])

	// 10. Maps as function parameters/returns
	fmt.Println("\n--- Functions ---")
	original := map[string]int{"a": 1, "b": 2}
	copied := copyMap(original)
	fmt.Println("Original:", original)
	fmt.Println("Copy:", copied)

	// Modify copy, original unchanged (maps are reference types, but we copied)
	copied["c"] = 3
	fmt.Println("After modifying copy:")
	fmt.Println("Original:", original)
	fmt.Println("Copy:", copied)

	// 11. Zero value map behavior
	fmt.Println("\n--- Zero Value ---")
	var emptyMap map[string]int
	// emptyMap is nil, so we can't assign to it
	// fmt.Println(emptyMap["key"]) // panic: assignment to entry in nil map
	fmt.Println("emptyMap == nil:", emptyMap == nil)

	// Make it non-nil to use
	emptyMap = make(map[string]int)
	emptyMap["test"] = 42
	fmt.Println("After make:", emptyMap)

	// 12. Comparing maps (only == nil works)
	fmt.Println("\n--- Comparison ---")
	m1 := map[string]int{"a": 1}
	m2 := map[string]int{"a": 1}
	fmt.Println("Maps can only be compared to nil, not to each other")
	fmt.Println("m1 == nil:", m1 == nil) // false
	var mNil map[string]int
	fmt.Println("mNil == nil:", mNil == nil) // true
	// To check if two maps have same content, you must compare manually
	same := mapsEqual(m1, m2)
	fmt.Println("m1 and m2 have same content:", same) // true
}

// copyMap returns a shallow copy of the map
func copyMap(m map[string]int) map[string]int {
	if m == nil {
		return nil
	}
	result := make(map[string]int, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

// mapsEqual compares two maps for equality (same keys and values)
func mapsEqual(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, av := range a {
		if bv, ok := b[k]; !ok || bv != av {
			return false
		}
	}
	return true
}