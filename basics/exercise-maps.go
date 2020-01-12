package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	fields := strings.Fields(s)
	for _, str := range fields {
		elem, exists := result[str]
		if exists {
			result[str] = elem + 1
		}else{
			result[str] = 1
		}
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
