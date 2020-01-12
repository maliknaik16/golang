
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:3]
	printSlice("d", d)
}

func printSlice(str string, arr []int) {
	fmt.Printf("%s len=%v cap=%v %v\n", str, len(arr), cap(arr), arr)
}
