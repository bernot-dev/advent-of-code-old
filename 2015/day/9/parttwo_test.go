package main

import (
	"strings"
	"testing"

	"github.com/bernot-dev/advent-of-code/helpers"
)

func TestPartTwo(t *testing.T) {
	tests := []TestCase{
		{"sample.txt", 982},
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

func TestMaxDistance(t *testing.T) {
	tests := []struct {
		inputFile string
		cities    []string
		expected  int
	}{
		{"sample.txt", []string{"Belfast", "Dublin"}, 141},
		{"sample.txt", []string{"Dublin", "Belfast"}, 141},
		{"sample.txt", []string{"Belfast", "Dublin", "London"}, 982},
		{"sample.txt", []string{"Belfast", "London", "Dublin"}, 982},
		{"sample.txt", []string{"Dublin", "Belfast", "London"}, 982},
		{"sample.txt", []string{"Dublin", "London", "Belfast"}, 982},
		{"sample.txt", []string{"London", "Belfast", "Dublin"}, 982},
		{"sample.txt", []string{"London", "Dublin", "Belfast"}, 982},
	}
	for _, tc := range tests {
		t.Run(tc.inputFile, func(t *testing.T) {
			t.Run(strings.Join(tc.cities, ","), func(t *testing.T) {
				i := ReadInput(tc.inputFile)
				actual := i.Distance(helpers.Max, tc.cities)
				if actual != tc.expected {
					t.Errorf("Expected: %v, Actual: %v", tc.expected, actual)
				}
			})
		})
	}
}
