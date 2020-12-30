package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	inputFile string
	expected  Solution
}

func assertDeepEquals(t *testing.T, actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
