
package main

import "fmt"

// A Slice is a dynamically-sized, flexible view into the elements of an array.
// The type `[]T` is a slice with elements of type 'T'.
// A slice is formed by specifying 2 indices as follows:
//    a[low : high]
// Slices are like references to arrays
func main() {
	names := [4]string {
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
}
