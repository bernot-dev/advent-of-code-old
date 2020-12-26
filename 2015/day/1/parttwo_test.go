package main

import (
	"testing"
)

func TestPartTwo(t *testing.T) {
	tests := []TestCase{
		{"input.txt", 1783},
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

func TestEntersBasement(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{")", 1},
		{"()())", 5},
	}
	for _, tc := range tests {
		actual := EntersBasement(tc.input)
		if actual != tc.expected {
			t.Errorf("Expected: %d, Actual: %d", tc.expected, actual)
		}
	}
}
