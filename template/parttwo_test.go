package main

import (
	"fmt"
	"testing"
)

func TestPartTwo(t *testing.T) {
	tests := map[string][]TestCase{
		"sample.txt": {0},
	}
	for inputFile, ts := range tests {
		t.Run(inputFile, func(t *testing.T) {
			input := ReadInput(inputFile)
			for _, tc := range ts {
				t.Run(fmt.Printf("%v", tc), func(t *testing.T) {
					actual := PartTwo(input)
					if actual != tc.expected {
						t.Errorf("Expected: %v, Actual: %v", tc.expected, actual)
					}
				})
			}
		})
	}
}
