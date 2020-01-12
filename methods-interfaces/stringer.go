
package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

// A Stringer is a type that can describe itself as a string.
//   type Stringer string {
//     String() string
//   }
// It is defined by the fmt package.
func main() {
	a := Person{"Arthur Dent", 41}
	z := Person{"Zaphod Beeblebrox", 9001}

	fmt.Println(a, z)
}
