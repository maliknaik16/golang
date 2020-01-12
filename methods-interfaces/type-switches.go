
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice of %v is %v.\n", v, v*v)
	case string:
		fmt.Printf("%v is %v bytes long.\n", v, len(v))
	default:
		fmt.Printf("I don't know about this type %T!\n", v)
	}
}

func main() {
	do(21)
	do("Hello")
	do(32.58)
	do(false)
}
