
package main

import (
	"testing"
)

// Source https://tutorialedge.net/golang/intro-testing-in-go/
func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 equals 4")
	}
}
