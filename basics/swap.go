
package main

import (
	"fmt"
)

func swap(str1, str2 string) (string, string) {
	return str2, str1
}

func main() {
	a, b := swap("World!", "Hello")

	fmt.Println(a, b)
}
