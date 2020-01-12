
package main

import "fmt"

type I interface{}

func main() {
	// Can also be written as
	// `var i interface{}`
	var i I
	describe(i)

	i = 75
	describe(i)

	i = "Bonjour"
	describe(i)
}

// Can also be written as
// `func describe(i I) {`
func describe(i interface{})  {
	fmt.Printf("(%v, %T)\n", i, i)
}
