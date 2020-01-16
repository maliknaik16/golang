
package main

import "fmt"

type operation func(x, y int) int

func function1(m func(x, y int) int) {
	fmt.Println(m(10, 20))
}

func function2() func(int, int) int {
	var t operation = func(a, b int) int {
		return a / b
	}
	return t
}

func main() {
	printName := func(name string) {
		fmt.Println("Hello,", name)
	}

	var m operation = func(a int, b int) int {
		return a * b
	}

	var d operation = function2()

	printName("John Doe")
	fmt.Println(m(3, 4))
	function1(m)
	fmt.Println(d(81, 9))
}
