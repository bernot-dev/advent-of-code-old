package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []TestCase{
		{"sample.txt", 1000, 1120},
	}
	for _, tc := range tests {
		t.Run(tc.inputFile, func(t *testing.T) {
			actual := PartOne(ReadInput(tc.inputFile), 1000)
			if actual != tc.expected {
				t.Errorf("Expected: %v, Actual: %v", tc.expected, actual)
			}
		})
	}
}
