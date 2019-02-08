package link

import (
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/pake-go/pake-lib/config"
)

func TestExecute_destoverwrite(t *testing.T) {
	setup(t)
	defer teardown(t)

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()
	cfg.SetPermanently("overwrite", "true")

	src := "test/testfile"
	dest := "test/destfile"
	err := os.Symlink(src, dest)
	if err != nil {
		t.Fatal(err)
	}

	l := &link{[]string{src, dest}}
	err = l.Execute(cfg, logger)
	if err != nil {
		t.Error(err)
	}
	_ = os.Remove(dest)
}

func TestExecute_destnooverwrite(t *testing.T) {
	setup(t)
	defer teardown(t)

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()
	cfg.SetPermanently("overwrite", "false")

	src := "test/testfile"
	dest := "test/destfile"

	l := &link{[]string{src, dest}}
	err := l.Execute(cfg, logger)
	if err != nil {
		t.Error(err)
	}
	_ = os.Remove(dest)
}

func TestExecute_destexistsnooverwrite(t *testing.T) {
	setup(t)
	defer teardown(t)

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()
	cfg.SetPermanently("overwrite", "false")

	src := "test/testfile"
	dest := "test/destfile"
	err := os.Symlink(src, dest)
	if err != nil {
		t.Fatal(err)
	}

	l := &link{[]string{src, dest}}
	err = l.Execute(cfg, logger)
	if err == nil {
		t.Errorf("Should not be able to create a symlink from %s to %s", src, dest)
	}
	_ = os.Remove(dest)
}

func TestCanHandle_invalidline(t *testing.T) {
	line := "link"
	lv := &LinkValidator{}
	if lv.CanHandle(line) {
		t.Errorf("%s should not be valid", line)
	}
}

func TestCanHandle_validline(t *testing.T) {
	line := "link "
	lv := &LinkValidator{}
	if !lv.CanHandle(line) {
		t.Errorf("%s should be valid", line)
	}
}

func TestValidateArgs_oneinvalidarg(t *testing.T) {
	var args []string
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		args = []string{"/dev/null/hello", "bye"}
	} else {
		args = []string{"hello?", "bye"}
	}
	lv := &LinkValidator{}
	if lv.ValidateArgs(args) == nil {
		t.Errorf("%+q should not be valid", args)
	}
}

func TestValidArgs_invalidargs(t *testing.T) {
	var args []string
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		args = []string{"/dev/null/hello", "/dev/null/bye"}
	} else {
		args = []string{"hello?", "bye?"}
	}
	lv := &LinkValidator{}
	if lv.ValidateArgs(args) == nil {
		t.Errorf("%+q should not be valid", args)
	}
}

func TestValidArgs_validargs(t *testing.T) {
	args := []string{"hello", "bye"}
	lv := &LinkValidator{}
	if lv.ValidateArgs(args) != nil {
		t.Errorf("%+q should be valid", args)
	}
}

func setup(t *testing.T) {
	newFile, err := os.Create("test/testfile")
	if err != nil {
		t.Fatal(err)
	}
	newFile.Close()
}

func teardown(t *testing.T) {
	err := os.RemoveAll("test/testfile")
	if err != nil {
		t.Fatal(err)
	}
}
