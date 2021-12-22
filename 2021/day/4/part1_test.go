package main

import "testing"

func TestPart1(t *testing.T) {
	type test struct {
		inputFile string
		expected  string
	}

	tests := map[string]test{
		"sample": {"sample.txt", "4512"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			numbers, cards := readInput(tc.inputFile)
			actual := Part1(numbers, cards)
			if actual != tc.expected {
				t.Errorf("Expected: %q, Actual: %q", tc.expected, actual)
			}
		})
	}
}
