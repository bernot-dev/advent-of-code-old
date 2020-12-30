package helpers

import (
	"reflect"
	"regexp"
	"testing"
)

func TestMapStringSubmatches(t *testing.T) {
	tests := []struct {
		re        *regexp.Regexp
		str       string
		expected  map[string]string
		expectErr bool
	}{
		{
			regexp.MustCompile(`(?P<a>b)(?P<c>d)(?P<e>f)`),
			"bdf",
			map[string]string{"a": "b", "c": "d", "e": "f"},
			false,
		},
	}
	for _, tc := range tests {
		actual, err := MapNamedSubmatches(tc.re, tc.str)
		if !tc.expectErr && err != nil {
			t.Errorf("Unexpected error, %s", err)
		}
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("Expected: %v, Actual: %v", tc.expected, actual)
		}
	}
}
