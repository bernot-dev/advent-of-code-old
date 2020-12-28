package main

import "testing"

type TestCase struct {
	inputFile string
	expected  Solution
}

func TestNew(t *testing.T) {
	tests := []struct {
		input    string
		expected Instruction
	}{
		{
			"turn on 0,1 through 998,999",
			Instruction{
				command: "turn on",
				start:   Coordinate{0, 1},
				end:     Coordinate{998, 999},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actual := New(tc.input)
			if tc.expected != actual {
				t.Errorf("Expected: %v, Actual: %v", tc.expected, actual)
			}
		})
	}
}
