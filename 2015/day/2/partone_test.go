package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []TestCase{
		{"sample.txt", 58 + 43},
	}
	for _, tc := range tests {
		t.Run(tc.inputFile, func(t *testing.T) {
			actual := PartOne(ReadInput(tc.inputFile))
			if actual != tc.expected {
				t.Errorf("Expected: %v, Actual: %v", tc.expected, actual)
			}
		})
	}
}

func TestWrapping(t *testing.T) {
	tests := []struct {
		input    Present
		expected int
	}{
		{Present{2, 3, 4}, 58},
		{Present{1, 1, 10}, 43},
	}
	for _, tc := range tests {
		actual := tc.input.Wrapping()
		if actual != tc.expected {
			t.Errorf("Expected: %d, Actual: %d", tc.expected, actual)
		}
	}
}
