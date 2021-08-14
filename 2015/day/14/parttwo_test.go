package main

import (
	"testing"
)

func TestPartTwo(t *testing.T) {
	tests := []TestCase{
		{"sample.txt", 1000, 689},
	}
	for _, tc := range tests {
		t.Run(tc.inputFile, func(t *testing.T) {
			actual := PartTwo(ReadInput(tc.inputFile), tc.length)
			if actual != tc.expected {
				t.Errorf("Expected: %v, Actual: %v", tc.expected, actual)
			}
		})
	}
}
