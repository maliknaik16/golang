
package main

import "fmt"

// A Struct is a collection of fields
type Vector3 struct {
	X float32
	Y float32
	Z float32
}

func main() {
	// If you want to omit the declaration then use the following:
	// p := &v
	// Instead of p = &v
	var p *Vector3
	v := Vector3{1.0, 2.0, 1.5}

	p = &v
	p.Z = 1.25
	(*p).X = 1.75
	fmt.Println(v)
}
