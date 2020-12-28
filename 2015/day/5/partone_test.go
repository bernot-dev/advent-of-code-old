package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []TestCase{
		{"sample.txt", 2},
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

func TestIsNice(t *testing.T) {
	rules := []Rule{ThreeVowels, DoubleLetter, NoBadPattern}
	tests := []struct {
		input    string
		expected bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actual := IsNice(tc.input, rules)
			if actual != tc.expected {
				t.Errorf("Expected: %t, Actual: %t", tc.expected, actual)
			}
		})
	}
}

func TestVowelCount(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"ugknbfddgicrmopn", 3},
		{"aaa", 3},
		{"jchzalrnumimnmhp", 3},
		{"haegwjzuvuyypxyu", 5},
		{"dvszwmarrgswjxmb", 1},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actual := VowelCount(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Actual: %d", tc.expected, actual)
			}
		})
	}
}

func TestDoubleLetter(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", true},
		{"dvszwmarrgswjxmb", true},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actual := DoubleLetter(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected: %t, Actual: %t", tc.expected, actual)
			}
		})
	}
}

func TestNoBadPattern(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", true},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", true},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actual := NoBadPattern(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected: %t, Actual: %t", tc.expected, actual)
			}
		})
	}
}
