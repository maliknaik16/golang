
package main

import (
	"fmt"
)

func factorial(n int) int {
	var fact int = 1
	for i := n; i > 1; i-- {
		fact = fact * i
	}

	return fact
}

func main() {
	fmt.Printf("The factorial of 5 is %v\n", factorial(5))
}
