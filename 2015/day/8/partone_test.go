package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []TestCase{
		{"sample.txt", 12},
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

func TestCodeCount(t *testing.T) {
	tests := []struct {
		inputFile string
		expected  []int
	}{
		{"sample.txt", []int{2, 5, 10, 6}},
	}
	for _, tc := range tests {
		t.Run(tc.inputFile, func(t *testing.T) {
			input := ReadInput(tc.inputFile)
			for i, s := range input {
				t.Run(s, func(t *testing.T) {
					actual := CodeCount(s)
					if actual != tc.expected[i] {
						t.Errorf("Expected: %v, Actual: %v", tc.expected[i], actual)
					}
				})
			}
		})
	}
}

func TestMemoryCountUnquote(t *testing.T) {
	tests := []struct {
		inputFile string
		expected  []int
	}{
		{"sample.txt", []int{0, 3, 7, 1}},
	}
	for _, tc := range tests {
		t.Run(tc.inputFile, func(t *testing.T) {
			input := ReadInput(tc.inputFile)
			for i, s := range input {
				t.Run(s, func(t *testing.T) {
					actual := MemoryCountUnquote(s)
					if actual != tc.expected[i] {
						t.Errorf("Expected: %v, Actual: %v", tc.expected[i], actual)
					}
				})
			}
		})
	}
}
