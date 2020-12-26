package helpers

import "testing"

func TestMin(t *testing.T) {
	tests := []struct {
		n        []int
		expected int
	}{
		{[]int{1, 2, 3}, 1},
		{[]int{5, 7, 1}, 1},
	}
	for _, tc := range tests {
		actual := Min(tc.n...)
		if actual != tc.expected {
			t.Errorf("Expected: %d, Actual: %d", tc.expected, actual)
		}
	}
}
func TestMax(t *testing.T) {
	tests := []struct {
		n        []int
		expected int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{5, 7, 1}, 7},
	}
	for _, tc := range tests {
		actual := Max(tc.n...)
		if actual != tc.expected {
			t.Errorf("Expected: %d, Actual: %d", tc.expected, actual)
		}
	}
}
