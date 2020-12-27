package main

import (
	"fmt"
	"testing"
)

func TestPartTwo(t *testing.T) {
	tests := []TestCase{
		{"sample.txt", 6742839},
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

func TestBeginsWithSixZeroes(t *testing.T) {
	tests := []struct {
		input    [16]byte
		expected bool
	}{
		{[16]byte{}, true},
		{[16]byte{0, 0, 0, 0xFF}, true},
		{[16]byte{0, 0, 1}, false},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%x", tc.input), func(t *testing.T) {
			actual := BeginsWithSixZeroes(&tc.input)
			if actual != tc.expected {
				t.Errorf("Expected: %t, Actual: %t", tc.expected, actual)
			}
		})
	}
}
