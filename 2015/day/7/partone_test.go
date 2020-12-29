package main

import (
	"fmt"
	"testing"
	"time"
)

func TestPartOne(t *testing.T) {
	tests := []TestCase{
		{"input.txt", "a", 16076},
		{"sample.txt", "d", 72},
		{"sample.txt", "e", 507},
		{"sample.txt", "f", 492},
		{"sample.txt", "g", 114},
		{"sample.txt", "h", 65412},
		{"sample.txt", "i", 65079},
		{"sample.txt", "x", 123},
		{"sample.txt", "y", 456},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s-%s", tc.inputFile, tc.solutionWire), func(t *testing.T) {
			input := ReadInput(tc.inputFile)
			actual := PartOne(input, tc.solutionWire)
			if actual != tc.expected {
				t.Errorf("Expected: %v, Actual: %v", tc.expected, actual)
			}
		})
	}
}

func TestSignal(t *testing.T) {
	ch := signal(1)
	time.Sleep(time.Millisecond)
	select {
	case <-ch:
	default:
		t.Errorf("Channel blocked")
	}
	select {
	case <-ch:
	default:
		t.Errorf("Channel does not provide multiple values")
	}
}
