package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		input    Input
		expected int
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}
	for _, tc := range tests {
		t.Run(string(tc.input), func(t *testing.T) {
			actual := PartOne(tc.input)
			if actual != Solution(tc.expected) {
				t.Errorf("Expected: %v, Actual: %v", tc.expected, actual)
			}
		})
	}
}
