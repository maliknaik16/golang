package main

import "fmt"

func fibonacci() func() int {
	a, b := -1, 1
	return func() int {
		c := a + b
		a = b
		b = c
		return c
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), " ")
	}
}
