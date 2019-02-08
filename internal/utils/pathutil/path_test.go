package pathutil

import (
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/google/uuid"
	homedir "github.com/mitchellh/go-homedir"
)

func TestExists_pathexists(t *testing.T) {
	path := "path.go"
	if !Exists(path) {
		t.Errorf("path.go should exists")
	}
}

func TestExists_pathnotexists(t *testing.T) {
	path := ""
	if Exists(path) {
		t.Errorf("empty path (\"\") should not exists")
	}
}

func TestExpand_pathcanexpand(t *testing.T) {
	path := "~/Code"
	expandedDir, err := Expand(path)
	if err != nil {
		t.Error(err)
	}

	home, err := homedir.Dir()
	if err != nil {
		t.Error(err)
	}
	expected := fmt.Sprintf("%s/Code", home)
	if expandedDir != expected {
		t.Errorf("Expected %s but got %s", expected, expandedDir)
	}
}

func TestExpand_pathcannotexpand(t *testing.T) {
	path := "~foo/foo"
	if _, err := Expand(path); err == nil {
		t.Errorf("This path should not be expandable")
	}
}

func TestIsDir_pathisdir(t *testing.T) {
	var path string
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		path = "/"
	} else {
		path = "C:\\"
	}

	isDir, err := IsDir(path)
	if err != nil {
		t.Error(err)
	}
	if !isDir {
		t.Errorf("%s should be a directory", path)
	}
}

func TestIsDir_pathisnotdir(t *testing.T) {
	isDir, err := IsDir("path_test.go")
	if err != nil {
		t.Error(err)
	}
	if isDir {
		t.Errorf("path_test.go should not be a directory")
	}
}

func TestIsSymlink_pathissymlink(t *testing.T) {
	rawUUID := uuid.New()
	uuid := rawUUID.String()

	if linkErr := os.Symlink("path_test.go", uuid); linkErr != nil {
		t.Error(linkErr)
	}
	isSymlink, err := IsSymlink(uuid)
	if err != nil {
		t.Error(err)
	}
	if !isSymlink {
		t.Errorf("%s should be a valid symlink", uuid)
	}
	if removeErr := os.Remove(uuid); removeErr != nil {
		t.Error(removeErr)
	}
}

func TestIsSymlink_pathisnotsymlink(t *testing.T) {
	path := "path_test.go"
	isSymlink, err := IsSymlink(path)
	if err != nil {
		t.Error(err)
	}
	if isSymlink {
		t.Errorf("%s should not be a valid symlink", path)
	}
}
