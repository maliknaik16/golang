
package main

import "fmt"

type LatLong struct {
	Lat, Long float64
}

var m map[string]LatLong

func main() {
	m = make(map[string]LatLong)
	m["Bell Labs"] = LatLong{
		40.68433,
		-74.39967,
	}

	var a = map[string]LatLong{
		"Bell Labs": LatLong{
			40.68433,
			-74.39967,
		},
		"Google Labs": LatLong{
			37.42202,
			-122.08408,
		},
	}

	var b = map[string]LatLong{
		"Bell Labs": {40.68433, -74.39967,},
		"Google Labs": {37.42202, -122.08408,},
	}

	fmt.Println(m["Bell Labs"])
	fmt.Println(a)
	fmt.Println(b)
}
