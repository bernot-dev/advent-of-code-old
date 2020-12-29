package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputFile    string
	solutionWire string
	expected     Solution
}

func assertGatesEqual(t *testing.T, actual, expected Gate) {
	if actual.raw != expected.raw || !reflect.DeepEqual(actual.inputWireIDs, expected.inputWireIDs) || actual.outputWireID != expected.outputWireID {
		t.Errorf("Actual: %v\nExpected: %v", actual, expected)
	}
}

func TestNewGate(t *testing.T) {
	tests := []struct {
		input    string
		expected Gate
	}{
		{"123 -> x", Gate{
			raw:          "123 -> x",
			operator:     ops["SIGNAL"],
			inputWireIDs: []interface{}{uint16(123)},
			outputWireID: "x",
		}},
		{"456 -> y", Gate{
			raw:          "456 -> y",
			operator:     ops["SIGNAL"],
			inputWireIDs: []interface{}{uint16(456)},
			outputWireID: "y",
		}},
		{"x AND y -> d", Gate{
			raw:          "x AND y -> d",
			operator:     ops["AND"],
			inputWireIDs: []interface{}{"x", "y"},
			outputWireID: "d",
		}},
		{"x OR y -> e", Gate{
			raw:          "x OR y -> e",
			operator:     ops["OR"],
			inputWireIDs: []interface{}{"x", "y"},
			outputWireID: "e",
		}},
		{"x LSHIFT 2 -> f", Gate{
			raw:          "x LSHIFT 2 -> f",
			operator:     ops["LSHIFT"],
			inputWireIDs: []interface{}{"x", uint16(2)},
			outputWireID: "f",
		}},
		{"y RSHIFT 2 -> g", Gate{
			raw:          "y RSHIFT 2 -> g",
			operator:     ops["RSHIFT"],
			inputWireIDs: []interface{}{"y", uint16(2)},
			outputWireID: "g",
		}},
		{"NOT x -> h", Gate{
			raw:          "NOT x -> h",
			operator:     ops["NOT"],
			inputWireIDs: []interface{}{"x"},
			outputWireID: "h",
		}},
		{"NOT y -> i", Gate{
			raw:          "NOT y -> i",
			operator:     ops["NOT"],
			inputWireIDs: []interface{}{"y"},
			outputWireID: "i",
		}},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actual := NewGate(tc.input)
			assertGatesEqual(t, actual, tc.expected)
		})
	}
}

func TestOps(t *testing.T) {
	tests := []struct {
		op       string
		input    []uint16
		expected uint16
	}{
		{"SIGNAL", []uint16{1}, 1},
		{"AND", []uint16{1, 2}, 0},
		{"OR", []uint16{1, 2}, 3},
		{"NOT", []uint16{1}, 65534},
		{"LSHIFT", []uint16{1, 1}, 2},
		{"RSHIFT", []uint16{2, 1}, 1},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s %v", tc.op, tc.input), func(t *testing.T) {
			actual := ops[tc.op](tc.input...)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Actual: %d", tc.expected, actual)
			}
		})
	}
}
