package remove

import (
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/pake-go/pake-lib/config"
	"github.com/pake-go/pakeos/internal/utils/pathutil"
)

func TestExecute_fileexist(t *testing.T) {
	setup(t)
	defer teardown(t)

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	path := "test/testdirectory/testfile"

	r := &remove{[]string{path}}
	err := r.Execute(cfg, logger)
	if err != nil {
		t.Error(err)
	}

	if pathutil.Exists(path) {
		t.Errorf("Expected %s to be removed", path)
	}
}

func TestExecute_filedoesnotexist(t *testing.T) {
	setup(t)
	defer teardown(t)

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	path := "test/testdirectory/testfilenotexists"

	r := &remove{[]string{path}}
	err := r.Execute(cfg, logger)
	if err != nil {
		t.Error(err)
	}

	if pathutil.Exists(path) {
		t.Errorf("Expected %s to not exist", path)
	}
}

func TestExecute_direxist(t *testing.T) {
	setup(t)
	defer teardown(t)

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	path := "test/testdirectory"

	r := &remove{[]string{path}}
	err := r.Execute(cfg, logger)
	if err != nil {
		t.Error(err)
	}

	if pathutil.Exists(path) {
		t.Errorf("Expected %s to be removed", path)
	}
}

func TestExecute_dirnotexist(t *testing.T) {
	setup(t)
	defer teardown(t)

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	path := "test/testdirectorynotexists"

	r := &remove{[]string{path}}
	err := r.Execute(cfg, logger)
	if err != nil {
		t.Error(err)
	}

	if pathutil.Exists(path) {
		t.Errorf("Expected %s to not exist", path)
	}
}

func TestCanHandle_invalidline(t *testing.T) {
	line := "remove"
	rv := &RemoveValidator{}
	if rv.CanHandle(line) {
		t.Errorf("%s should not be valid", line)
	}
}

func TestCanHandle_validline(t *testing.T) {
	line := "remove "
	rv := &RemoveValidator{}
	if !rv.CanHandle(line) {
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
	rv := &RemoveValidator{}
	if rv.ValidateArgs(args) {
		t.Errorf("%+q should not be valid", args)
	}
}

func TestValidateArgs_validarg(t *testing.T) {
	args := []string{"hello"}
	rv := &RemoveValidator{}
	if !rv.ValidateArgs(args) {
		t.Errorf("%+q should be valid", args)
	}
}

func setup(t *testing.T) {
	err := os.Mkdir("test/testdirectory", 0755)
	if err != nil {
		t.Fatal(err)
	}
	newFile, err := os.Create("test/testdirectory/testfile")
	if err != nil {
		t.Fatal(err)
	}
	newFile.Close()
}

func teardown(t *testing.T) {
	err := os.RemoveAll("test/testdirectory")
	if err != nil {
		t.Fatal(err)
	}
}
