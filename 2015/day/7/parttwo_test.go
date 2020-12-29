package main

import (
	"testing"
)

func TestPartTwo(t *testing.T) {
	tests := []TestCase{
		{"sample.txt", "a", 0},
	}
	for _, tc := range tests {
		t.Run(tc.inputFile, func(t *testing.T) {
			actual := PartTwo(ReadInput(tc.inputFile))
			if actual != tc.expected {
				t.Errorf("Expected: %v, Actual: %v", tc.expected, actual)
			}
		})
	}
}
