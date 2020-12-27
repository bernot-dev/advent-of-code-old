package main

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []TestCase{
		{"sample.txt", 609043},
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

func TestFindHash(t *testing.T) {
	tests := []struct {
		key      string
		expected int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}
	for _, tc := range tests {
		t.Run(tc.key, func(t *testing.T) {
			actual := FindHash(tc.key, BeginsWithFiveZeroes)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Actual: %d", tc.expected, actual)
			}
		})
	}
}

func TestBeginsWithFiveZeroes(t *testing.T) {
	tests := []struct {
		input    [16]byte
		expected bool
	}{
		{md5.Sum([]byte("abcdef609043")), true},
		{md5.Sum([]byte("pqrstuv1048970")), true},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%x", tc.input), func(t *testing.T) {
			actual := BeginsWithFiveZeroes(&tc.input)
			if actual != tc.expected {
				t.Errorf("Expected: %t, Actual: %t", tc.expected, actual)
			}
		})
	}
}
