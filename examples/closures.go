
package main

import "fmt"

func fibonacci() func(int) {
	a, b := 0, 1
	return func(n int) {
		fmt.Print(a, " ", b)
		for i := 0; i < n; i++ {
			c := a + b
			fmt.Print(" ", c)
			a = b
			b = c
		}
	}
}

func main() {
	f := fibonacci()
	f(10)
}
