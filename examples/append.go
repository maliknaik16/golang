
package main

import "fmt"

func customAppend(src []int, args...int) []int {
	srcLen := len(src)
	varLen := len(args)
	var sliceLen int = srcLen + varLen
	slice := make([]int, sliceLen)
	index := 0

	for _, v := range src {
		slice[index] = v
		index = index + 1
	}

	for _, v := range args {
		slice[index] = v
		index = index + 1
	}
	return slice
}

func main() {
	a := []int{1,2,3,4,5,6,7,8,9,10}
	a = customAppend(a[:2], a[4:]...)
	fmt.Println(a)
}
