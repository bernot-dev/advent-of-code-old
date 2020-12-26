package helpers

import (
	"io/ioutil"
	"os"
)

// String returns the contents of the file as a single string
func String(r *os.File) string {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return string(b)
}
