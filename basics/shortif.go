
package main

import "fmt"
import "math"

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(2, 5, 28),
		pow(3, 4, 82),
	)
}
