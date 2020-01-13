
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
	if float64(err) < 0 {
		return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(err))
	} else {
		return ""
	}
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), ErrNegativeSqrt(x)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
