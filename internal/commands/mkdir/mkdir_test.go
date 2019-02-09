package mkdir

import (
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/pake-go/pake-lib/config"
	"github.com/pake-go/pakeos/internal/utils/pathutil"
)

func TestExecute_destexists(t *testing.T) {
	setup(t)
	defer teardown(t)

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	path := "test/testdirectory"

	m := &mkdir{[]string{path}}
	err := m.Execute(cfg, logger)
	if err != nil {
		t.Error(err)
	}

	if !pathutil.Exists(path) {
		t.Errorf("Expected %s to be created", path)
	}
}

func TestExecute_destnotexist(t *testing.T) {
	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	path := "test/testdirectory"

	m := &mkdir{[]string{path}}
	err := m.Execute(cfg, logger)
	if err != nil {
		t.Error(err)
	}

	if !pathutil.Exists(path) {
		t.Errorf("Expected %s to be created", path)
	}
	_ = os.RemoveAll(path)
}

func TestCanHandle_invalidline(t *testing.T) {
	line := "mkdir"

	mv := &MkdirValidator{}
	if mv.CanHandle(line) {
		t.Errorf("%s should not be valid", line)
	}
}

func TestCanHandle_validline(t *testing.T) {
	line := "mkdir "

	mv := &MkdirValidator{}
	if !mv.CanHandle(line) {
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
	mv := &MkdirValidator{}
	if mv.ValidateArgs(args) == nil {
		t.Errorf("%+q should not be valid", args)
	}
}

func TestValidateArgs_validarg(t *testing.T) {
	args := []string{"hello"}

	mv := &MkdirValidator{}
	if mv.ValidateArgs(args) != nil {
		t.Errorf("%+q should be valid", args)
	}
}

func setup(t *testing.T) {
	err := os.Mkdir("test/testdirectory", 0755)
	if err != nil {
		t.Fatal(err)
	}
}

func teardown(t *testing.T) {
	err := os.RemoveAll("test/testdirectory")
	if err != nil {
		t.Fatal(err)
	}
}
