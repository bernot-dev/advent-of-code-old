package main

import (
	"testing"
)

func TestPartTwo(t *testing.T) {
	tests := []TestCase{
		{"sample.txt", 0},
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

func TestRibbon(t *testing.T) {
	tests := []struct {
		input    Present
		expected int
	}{
		{Present{2, 3, 4}, 34},
		{Present{1, 1, 10}, 14},
	}
	for _, tc := range tests {
		actual := tc.input.Ribbon()
		if actual != tc.expected {
			t.Errorf("Expected: %d, Actual: %d", tc.expected, actual)
		}
	}
}

func TestVolume(t *testing.T) {
	tests := []struct {
		input    Present
		expected int
	}{
		{Present{2, 3, 4}, 24},
		{Present{1, 1, 10}, 10},
	}
	for _, tc := range tests {
		actual := tc.input.Volume()
		if actual != tc.expected {
			t.Errorf("Expected: %d, Actual: %d", tc.expected, actual)
		}
	}
}
