package main

import (
	"testing"
)

func TestPartTwo(t *testing.T) {
	tests := []TestCase{
		{"sample2.txt", 2},
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

func TestRepeatedPair(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"xyxy", true},
		{"aabcdefgaa", true},
		{"aaa", false},
		{"aaab", false},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actual := RepeatedPair(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected: %t, Actual: %t", tc.expected, actual)
			}
		})
	}
}

func TestThirdWheel(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"xyx", true},
		{"abcdefeghi", true},
		{"efe", true},
		{"aaa", true},
		{"abba", false},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actual := ThirdWheel(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected: %t, Actual: %t", tc.expected, actual)
			}
		})
	}
}
