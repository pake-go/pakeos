package validpath

import (
	"io/ioutil"
	"os"
)

func Valid(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	var d []byte
	if err := ioutil.WriteFile(path, d, 0644); err == nil {
		os.Remove(path)
		return true
	}
	return false
}
