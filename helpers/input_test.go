package helpers

import (
	"os"
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	inputFile, err := os.Open("string.txt")
	if err != nil {
		t.Error(err)
	}
	defer inputFile.Close()
	actual := String(inputFile)
	expected := "abc"
	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestStringLines(t *testing.T) {
	inputFile, err := os.Open("stringlines.txt")
	if err != nil {
		t.Error(err)
	}
	defer inputFile.Close()
	actual := StringLines(inputFile)
	expected := []string{"a", "b", "c"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
