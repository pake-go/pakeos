package move

import (
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/pake-go/pake-lib/config"
	"github.com/pake-go/pakeos/internal/utils/pathutil"
)

func TestExecute_destexistsisfilesrcisfile(t *testing.T) {
	setup(t)
	defer teardown(t)

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	src := "test/testdirectory/testfile"
	dest := "test/testdirectoryexists/testfileexists"

	m := &move{[]string{src, dest}}
	err := m.Execute(cfg, logger)
	if err != nil {
		t.Error(err)
	}

	if !pathutil.Exists(dest) ||
		pathutil.Exists(src) {
		t.Errorf("Expected %s to be moved to %s", src, dest)
	}
}

func TestExecute_destexistsisfilesrcisdir(t *testing.T) {
	setup(t)
	defer teardown(t)

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	src := "test/testdirectory"
	dest := "test/testdirectoryexists/testfileexists"

	m := &move{[]string{src, dest}}
	err := m.Execute(cfg, logger)
	if err == nil {
		t.Errorf("Should not be able to move %s to %s", src, dest)
	}
}

func TestExecute_destexistsisdir(t *testing.T) {
	setup(t)
	defer teardown(t)

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	src := "test/testdirectory"
	dest := "test/testdirectoryexists"

	m := &move{[]string{src, dest}}
	err := m.Execute(cfg, logger)
	if err != nil {
		t.Error(err)
	}

	if !pathutil.Exists(dest) ||
		pathutil.Exists(src) {
		t.Errorf("Expected %s to be moved to %s", src, dest)
	}
}

func TestExecute_destnotexists(t *testing.T) {
	setup(t)
	defer teardown(t)

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	src := "test/testdirectory"
	dest := "test/testdirectorynotexists"

	m := &move{[]string{src, dest}}
	err := m.Execute(cfg, logger)
	if err != nil {
		t.Error(err)
	}

	if !pathutil.Exists(dest) ||
		pathutil.Exists(src) {
		t.Errorf("Expected %s to be moved to %s", src, dest)
	}
	_ = os.RemoveAll(dest)
}

func TestCanHandle_invalidline(t *testing.T) {
	line := "move"
	mv := &MoveValidator{}
	if mv.CanHandle(line) {
		t.Errorf("%s should not be valid", line)
	}
}

func TestCanHandle_validline(t *testing.T) {
	line := "move "
	mv := &MoveValidator{}
	if !mv.CanHandle(line) {
		t.Errorf("%s should be valid", line)
	}
}

func TestValidateArgs_oneinvalidargs(t *testing.T) {
	var args []string
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		args = []string{"hello", "/dev/null/hello"}
	} else {
		args = []string{"hello", "hello?"}
	}
	mv := &MoveValidator{}
	if mv.ValidateArgs(args) == nil {
		t.Errorf("%+q should not be valid", args)
	}
}

func TestValidateArgs_allinvalidargs(t *testing.T) {
	var args []string
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		args = []string{"/dev/null/hello", "/dev/null/hello"}
	} else {
		args = []string{"hello?", "hello?"}
	}
	mv := &MoveValidator{}
	if mv.ValidateArgs(args) == nil {
		t.Errorf("%+q should not be valid", args)
	}
}

func TestValidateArgs_validargs(t *testing.T) {
	args := []string{"hello", "bye"}
	mv := &MoveValidator{}
	if mv.ValidateArgs(args) != nil {
		t.Errorf("%+q should be valid", args)
	}
}

func setup(t *testing.T) {
	err := os.Mkdir("test/testdirectory", 0755)
	if err != nil {
		t.Fatal(err)
	}
	err = os.Mkdir("test/testdirectoryexists", 0755)
	if err != nil {
		t.Fatal(err)
	}
	newFile, err := os.Create("test/testdirectory/testfile")
	if err != nil {
		t.Fatal(err)
	}
	newFile.Close()
	newFileTwo, err := os.Create("test/testdirectoryexists/testfileexists")
	if err != nil {
		t.Fatal(err)
	}
	newFileTwo.Close()
}

func teardown(t *testing.T) {
	err := os.RemoveAll("test/testdirectory")
	if err != nil {
		t.Fatal(err)
	}
	err = os.RemoveAll("test/testdirectoryexists")
	if err != nil {
		t.Fatal(err)
	}
}
