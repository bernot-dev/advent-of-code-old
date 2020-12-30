package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []TestCase{
		{"sample.txt", 174274},
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

func TestLookAndSay(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{"1", "11"},
		{"11", "21"},
		{"21", "1211"},
		{"1211", "111221"},
		{"111221", "312211"},
	}
	for _, tc := range tests {
		t.Run(tc.s, func(t *testing.T) {
			actual := string(LookAndSay([]byte(tc.s)))
			if actual != tc.expected {
				t.Errorf("Expected %s, Actual: %s", tc.expected, actual)
			}
		})
	}
}
