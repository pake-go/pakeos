package validpath

import (
	"runtime"
	"testing"
)

func TestValid_invalidpath(t *testing.T) {
	var path string
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		path = "/dev/null/hi"
	} else {
		path = "hello?"
	}

	if Valid(path) {
		t.Errorf("%s should not be a valid path", path)
	}
}

func TestValid_validpath(t *testing.T) {
	path := "test"
	if !Valid(path) {
		t.Errorf("%s should be a valid path", path)
	}
}
