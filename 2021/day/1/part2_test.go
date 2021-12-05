package main

import (
	"strconv"
	"testing"
)

func TestPart2(t *testing.T) {
	type test struct {
		input    inputType
		expected string
	}

	tests := map[string]test{
		"sample": {
			input:    readInput("sample2.txt"),
			expected: strconv.Itoa(5),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := part2(tc.input)
			expected := tc.expected

			if expected != actual {
				t.Errorf("\nExpected: %s\nActual: %s\n", expected, actual)
			}
		})
	}
}
