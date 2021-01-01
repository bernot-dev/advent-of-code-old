package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []TestCase{
		{"sample.txt", "ghjaabcc"},
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

func TestNextValid(t *testing.T) {
	tests := []struct {
		input, expected []byte
	}{
		{[]byte("ghijklmn"), []byte("ghjaabcc")},
		{[]byte("aaaaaaaa"), []byte("aaaaabcc")},
		{[]byte("ffeeaaaa"), []byte("ffeeaabc")},
		{[]byte("aaaabbcc"), []byte("aaaabbcd")},
		{[]byte("aaaabcba"), []byte("aaaabcbb")},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s", tc.input), func(t *testing.T) {
			NextValid(&tc.input)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("Expected: %s, Actual: %s", tc.input, tc.expected)
			}
		})
	}
}

func BenchmarkPartOne(b *testing.B) {
	input := ReadInput("input.txt")
	for i := 0; i < b.N; i++ {
		PartOne(input)
	}
}

func TestIncrement(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"a", "b"},
		{"az", "ba"},
		{"azz", "baa"},
		{"azzz", "baaa"},
		{"azzzz", "baaaa"},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actual := make([]byte, len(tc.input))
			copy(actual, tc.input)
			Increment(&actual, len(tc.input)-1)
			if !reflect.DeepEqual(actual, []byte(tc.expected)) {
				t.Errorf("Expected: %s, Actual: %s", tc.expected, string(actual))
			}
		})
	}
}

func TestHasStraight(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc", true},
		{"bcd", true},
		{"cde", true},
		{"xyz", true},
		{"abd", false},
		{"hijklmmn", true},
		{"abbcegjk", false},
		{"abcdefgh", true},
		{"abcdffaa", true},
		{"ghijklmn", true},
		{"ghjaabcc", true},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			bytes := []byte(tc.input)
			actual := hasStraight(&bytes)
			if actual != tc.expected {
				t.Errorf("Expected: %t, Actual: %t", tc.expected, actual)
			}
		})
	}
}

func TestContainsBannedLetters(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc", false},
		{"bcd", false},
		{"cde", false},
		{"xyz", false},
		{"abd", false},
		{"hijklmmn", true},
		{"abbcegjk", false},
		{"abcdefgh", false},
		{"abcdffaa", false},
		{"ghijklmn", true},
		{"ghjaabcc", false},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			bytes := []byte(tc.input)
			actual := containsBannedLetters(&bytes)
			if actual != tc.expected {
				t.Errorf("Expected: %t, Actual: %t", tc.expected, actual)
			}
		})
	}
}

func TestContainsTwoPair(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc", false},
		{"bcd", false},
		{"cde", false},
		{"xyz", false},
		{"abd", false},
		{"aaaaaaaa", false},
		{"hijklmmn", false},
		{"abbcegjk", false},
		{"abcdefgh", false},
		{"abcdffaa", true},
		{"ghijklmn", false},
		{"ghjaabcc", true},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			bytes := []byte(tc.input)
			actual := containsTwoPair(&bytes)
			if actual != tc.expected {
				t.Errorf("Expected: %t, Actual: %t", tc.expected, actual)
			}
		})
	}
}
