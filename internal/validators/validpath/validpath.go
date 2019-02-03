package validpath

import (
	"io/ioutil"
	"os"

	"github.com/pake-go/pakeos/internal/utils/pathutil"
)

func Valid(path string) bool {
	path, _ = pathutil.Expand(path)
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
