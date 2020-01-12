
package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 3)
	printSlice(s)

	s = append(s, 2, 3, 5, 7, 11, 13)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%v cap=%v %v\n", len(s), cap(s), s)
}
