package main

import (
	"testing"
)

func TestPartTwo(t *testing.T) {
	tests := []struct {
		input    Input
		expected int
	}{
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}
	for _, tc := range tests {
		t.Run(string(tc.input), func(t *testing.T) {
			actual := PartTwo(tc.input)
			if actual != Solution(tc.expected) {
				t.Errorf("Expected: %v, Actual: %v", tc.expected, actual)
			}
		})
	}
}
