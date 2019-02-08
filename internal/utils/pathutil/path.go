package pathutil

import (
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

// Exists checks to see if the given path exists.
func Exists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

// Expand attempts to expand ~ or $HOME in the given path and return
// any error found during that process.
func Expand(path string) (string, error) {
	if strings.HasPrefix(path, "$HOME") {
		path = strings.Replace(path, "$HOME", "~", 1)
	}
	return homedir.Expand(path)
}

// IsDir checks to see if the given path is a directory that exists and
// return any error found.
func IsDir(path string) (bool, error) {
	file, err := os.Open(path)

	if err != nil {
		return false, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return false, err
	}
	if fileInfo.IsDir() {
		return true, nil
	} else {
		return false, nil
	}
}

// IsSymlihnk checks to see if the given path is a valid symlink that exists and
// return any error found.
func IsSymlink(path string) (bool, error) {
	fileInfo, err := os.Lstat(path)
	if err == nil {
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			return true, nil
		}
		return false, nil
	}
	return false, err
}
