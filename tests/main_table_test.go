
package main

import (
	"testing"
	"fmt"
)

func TestTableCalculate(t *testing.T) {
	tests := []struct {
		input int
		expected int
	}{
		{2, 4},
		{1000, 1002},
		{9999, 10001},
		{447, 449},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error(fmt.Sprintf("Test Failed: %v Inputted, %v expected, received: %v", test.input, test.expected, output))
		}
	}
}
