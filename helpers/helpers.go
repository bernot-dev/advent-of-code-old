package helpers

import (
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
