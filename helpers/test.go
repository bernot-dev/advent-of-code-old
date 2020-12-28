package helpers

import (
	"reflect"
	"testing"
)

func assertDeepEquals(t *testing.T, actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v\nActual: %v", expected, actual)
	}
}
