
package main

import "fmt"

// Golang doesn't support pointer arithmetic
func main() {
	var p *int
	i, j := 49, 625

	// point to i
	p = &i

	// read i through the pointer
	fmt.Println(*p)

	// set i through the pointer
	*p = 22

	// print the new value of i
	fmt.Println(i)

	// point to j
	p = &j

	//read j through the pointer
	fmt.Println(*p)

	// divide j through the pointer
	*p = *p / 25

	// print the new value of j
	fmt.Println(j)
}
