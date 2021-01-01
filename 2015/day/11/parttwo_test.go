package main

import (
	"testing"
)

func TestPartTwo(t *testing.T) {
	tests := []TestCase{
		{"sample.txt", "ghjbbcdd"},
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

func BenchmarkPartTwo(b *testing.B) {
	input := ReadInput("input.txt")
	for i := 0; i < b.N; i++ {
		PartTwo(input)
	}
}
