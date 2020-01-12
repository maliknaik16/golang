
package main

import "fmt"

// A type assertion provides access to an interface value's underlying concrete value.
//   t := i.(T)
// Testing whether the interface holds a specific type:
//   t, ok := i.(T)
func main() {
	var i interface{} = "Hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
}
