
package main

import "fmt"

// interface values can be thought of as a tuple of a value and concrete type.
//   (value, type)
// An interface value holds a vlaue of a specific underlying concrete type.

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I
	var t *T
	var f F = 5.454

	i = t
	describe(i)
	i.M()

	i = f
	describe(i)
	i.M()

	t = &T{"Bonjour"}
	i = t
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
