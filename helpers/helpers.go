package helpers

import (
	"bufio"
	"io/ioutil"
	"os"
)

// String returns the contents of the file as a single string
func String(f *os.File) string {
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// StringLines return the contents of the file with each line as an element of a slice
func StringLines(f *os.File) (lines []string) {
	lines = make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}
