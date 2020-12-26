package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []TestCase{
		{"sample.txt", -3},
		{"input.txt", 232},
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

func TestFloor(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actual := Floor(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Actual %d", tc.expected, actual)
			}
		})
	}
}
