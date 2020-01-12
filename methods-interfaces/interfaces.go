
package main

import(
	"fmt"
	"math"
)

// An interface type is defined as a set of method signatures.
type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var a Abser
	v := Vertex{3, 4}
	f := MyFloat(-58.866)

	a = &v
	a = f
	//a = &v

	fmt.Println(a.Abs())
}
