package main

import (
	"strings"
	"testing"

	"github.com/bernot-dev/advent-of-code/helpers"
)

func TestPartOne(t *testing.T) {
	tests := []TestCase{
		{"sample.txt", 605},
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

func TestMinDistance(t *testing.T) {
	tests := []struct {
		inputFile string
		cities    []string
		expected  int
	}{
		{"sample.txt", []string{"Belfast", "Dublin"}, 141},
	}
	for _, tc := range tests {
		t.Run(tc.inputFile, func(t *testing.T) {
			i := ReadInput(tc.inputFile)
			actual := i.Distance(helpers.Min, tc.cities)
			if actual != tc.expected {
				t.Errorf("Expected: %v, Actual: %v", tc.expected, actual)
			}
		})
	}
}

func TestPermutations(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{"Belfast", "London"}, 2},
		{[]string{"London", "Dublin", "Belfast"}, 3},
	}
	for _, tc := range tests {
		t.Run(strings.Join(tc.input, ","), func(t *testing.T) {
			p := Permutations(tc.input)
			actual := len(p)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Actual: %d", tc.expected, actual)
			}
		})
	}
}
