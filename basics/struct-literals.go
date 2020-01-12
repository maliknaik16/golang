
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{2, 3}   // type Vertex
	v2 = Vertex{X: 10}  // Y: 0
	v3 = Vertex{}      // X: 0 and Y: 0
	p = &Vertex{5, 4}   // type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
