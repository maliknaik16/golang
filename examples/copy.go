

package main

import "fmt"

func customCopy(dest []int, src []int) int {
	count := 0
	for i, v := range src {
		dest[i] = v
		count = count + 1
	}
	return count
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	a := make([]int, 3)

	n := customCopy(a, s[1:4])
	fmt.Println("Number of elements copied: ", n)
	fmt.Println(a)
}
