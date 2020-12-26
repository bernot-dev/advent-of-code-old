package helpers

import (
	"os"
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
