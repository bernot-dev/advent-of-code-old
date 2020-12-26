package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	inputFile string
	expected  Solution
}

func TestNewPresent(t *testing.T) {
	tests := []struct {
		input    string
		expected Present
	}{
		{"2x3x4", Present{2, 3, 4}},
		{"1x1x10", Present{1, 1, 10}},
	}
	for _, tc := range tests {
		actual := NewPresent(tc.input)
		assertDeepEquals(t, actual, tc.expected)
	}
}

func assertDeepEquals(t *testing.T, actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual: %v\nExpected: %v", actual, expected)
	}
}
