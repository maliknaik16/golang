
package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var s string = ""
	for i, b := range ip {
		if i == 3 {
			s = s + fmt.Sprintf("%v", b)
		}else{
			s = s + fmt.Sprintf("%v", b) + "."
		}
	}
	return s
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
