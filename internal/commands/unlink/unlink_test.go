package unlink

import (
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/pake-go/pake-lib/config"
	"github.com/pake-go/pakeos/internal/utils/pathutil"
)

func TestExecute_notsymlink(t *testing.T) {
	setup(t)
	defer teardown(t)

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	path := "test/testfile"

	u := &unlink{[]string{path}}
	err := u.Execute(cfg, logger)
	if err == nil {
		t.Errorf("%s should not be unlinked because it's not a symlink", path)
	}
}

func TestExecute_symlinkexist(t *testing.T) {
	setup(t)
	defer teardown(t)

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	path := "test/testsymlink"

	u := &unlink{[]string{path}}
	err := u.Execute(cfg, logger)
	if err != nil {
		t.Error(err)
	}
	if isSymlink, err := pathutil.IsSymlink(path); err == nil && isSymlink {
		t.Errorf("Expected %s to be unlinked", path)
	}
}

func TestExecute_symlinknotexist(t *testing.T) {
	setup(t)
	defer teardown(t)

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	path := "test/testsymlinknotexists"

	u := &unlink{[]string{path}}
	err := u.Execute(cfg, logger)
	if err == nil {
		t.Errorf("%s should not be unlinked because it doesn't exist", path)
	}
}

func TestCanHandle_invalidline(t *testing.T) {
	line := "unlink"

	uv := &UnlinkValidator{}
	if uv.CanHandle(line) {
		t.Errorf("%s should not be valid", line)
	}
}

func TestCanHandle_validline(t *testing.T) {
	line := "unlink "

	uv := &UnlinkValidator{}
	if !uv.CanHandle(line) {
		t.Errorf("%s should be valid", line)
	}
}

func TestValidateArgs_invalidarg(t *testing.T) {
	var args []string
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		args = []string{"/dev/null/hello"}
	} else {
		args = []string{"hello?"}
	}

	uv := &UnlinkValidator{}
	if uv.ValidateArgs(args) == nil {
		t.Errorf("%+q should not be valid", args)
	}
}

func TestValidateArgs_validarg(t *testing.T) {
	args := []string{"hello"}
	uv := &UnlinkValidator{}
	if uv.ValidateArgs(args) != nil {
		t.Errorf("%+q should be valid", args)
	}
}

func setup(t *testing.T) {
	newFile, err := os.Create("test/testfile")
	if err != nil {
		t.Error(err)
	}
	newFile.Close()
	err = os.Symlink("test/testfile", "test/testsymlink")
	if err != nil {
		t.Error(err)
	}
}

func teardown(t *testing.T) {
	err := os.Remove("test/testfile")
	if err != nil {
		t.Error(err)
	}
	symlinkPath := "test/testsymlink"
	if isSymlink, err := pathutil.IsSymlink(symlinkPath); err == nil && isSymlink {
		err = os.Remove(symlinkPath)
		if err != nil {
			t.Error(err)
		}
	}
}
