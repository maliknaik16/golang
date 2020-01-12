
package main

import (
	"fmt"
	"math/cmplx"
)

// The following types are supported by Go lang:
//   - bool
//   - string
//   - int, int8, int16, int32, int64
//   - uint, uint8, uint16, uint32, uint64, uintptr
//   - rune (alias for uint8)
//   - float32, float64
//   - complex64, complex128


var (
	IsDeveloper bool = true
	Text string = "Go Types"
	UnsignedInteger uint64 = 1 << 64 - 1
	ComplexNumber complex128 = cmplx.Sqrt(10 - 25i)
	FloatNumber float64 = 1515.887521451
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", IsDeveloper, IsDeveloper)
	fmt.Printf("Type: %T Value: %v\n", Text, Text)
	fmt.Printf("Type: %T Value: %v\n", UnsignedInteger, UnsignedInteger)
	fmt.Printf("Type: %T Value: %v\n", ComplexNumber, ComplexNumber)
	fmt.Printf("Type: %T Value: %v\n", FloatNumber, FloatNumber)
}
