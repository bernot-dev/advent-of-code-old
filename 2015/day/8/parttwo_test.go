package main

import (
	"testing"
)

func TestPartTwo(t *testing.T) {
	tests := []TestCase{
		{"sample.txt", 19},
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

func TestMemoryCountQuote(t *testing.T) {
	tests := []struct {
		inputFile string
		expected  []int
	}{
		{"sample.txt", []int{6, 9, 16, 11}},
	}
	for _, tc := range tests {
		t.Run(tc.inputFile, func(t *testing.T) {
			input := ReadInput(tc.inputFile)
			for i, s := range input {
				t.Run(s, func(t *testing.T) {
					actual := MemoryCountQuote(s)
					if actual != tc.expected[i] {
						t.Errorf("Expected: %v, Actual: %v", tc.expected[i], actual)
					}
				})
			}
		})
	}
}
