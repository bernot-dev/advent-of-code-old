package main

import (
	"fmt"
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := map[string][]TestCase{
		"sample.txt": {{0}},
	}
	for inputFile, ts := range tests {
		t.Run(inputFile, func(t *testing.T) {
			input := ReadInput(inputFile)
			for _, tc := range ts {
				t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
					actual := PartOne(input)
					if actual != tc.expected {
						t.Errorf("Expected: %v, Actual: %v", tc.expected, actual)
					}
				})
			}
		})
	}
}
