
package main

import "fmt"

func main() {
	m := make(map[string]int)

	// Insert a new element
	m["key"] = 12
	fmt.Println("The value:", m["key"])

	// Update the element
	m["key"] = 54
	fmt.Println("The value:", m["key"])

	// Delete the element
	delete(m, "key")
	fmt.Println("The value: ", m["key"])

	// Test whether key exists
	element, exists := m["key"]
	fmt.Printf("The value: %v and Element exists? %v\n", element, exists)
}
