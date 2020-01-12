
package main

import "fmt"

// Index or value can be skipped by writing '_' in place of the variable name.
func main() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
