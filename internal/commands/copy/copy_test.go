package copy

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/pake-go/pake-lib/config"
	"github.com/pake-go/pakeos/internal/utils/pathutil"
	"github.com/udhos/equalfile"
)

func TestExecute_copydiroverwrite(t *testing.T) {
	setup(t)
	defer teardown(t)

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()
	cfg.SetPermanently("overwrite", "true")

	src := "test/testdirectory"
	dest := "test/destdirectory"
	err := os.Mkdir(dest, 0755)
	if err != nil {
		t.Fatal(err)
	}

	c := &copy{[]string{src, dest}}
	err = c.Execute(cfg, logger)
	if err != nil {
		t.Error(err)
	}

	if !pathutil.Exists(fmt.Sprintf("%s/testfile", dest)) {
		t.Errorf("Expected %s to be copied to %s", src, dest)
	}
	_ = os.RemoveAll(dest)
}

func TestExecute_copydirnooverwrite(t *testing.T) {
	setup(t)
	defer teardown(t)

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()
	cfg.SetPermanently("overwrite", "false")

	src := "test/testdirectory"
	dest := "test/destdirectory"
	_ = os.RemoveAll(dest)
	c := &copy{[]string{src, dest}}
	err := c.Execute(cfg, logger)
	if err != nil {
		t.Error(err)
	}

	if !pathutil.Exists(fmt.Sprintf("%s/testfile", dest)) {
		t.Errorf("Expected %s to be copied to %s", src, dest)
	}
	_ = os.RemoveAll(dest)
}

func TestExecute_copydirdestexist(t *testing.T) {
	setup(t)
	defer teardown(t)

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()
	cfg.SetPermanently("overwrite", "false")

	src := "test/testdirectory"
	dest := "test/destdirectory"
	err := os.Mkdir(dest, 0755)
	if err != nil {
		t.Fatal(err)
	}

	c := &copy{[]string{src, dest}}
	err = c.Execute(cfg, logger)
	if err == nil {
		t.Errorf("Should not be able to copy %s to %s", src, dest)
	}
	_ = os.RemoveAll(dest)
}

func TestExecute_copyfileoverwrite(t *testing.T) {
	setup(t)
	defer teardown(t)

	err := os.Mkdir("test/destdirectory", 0755)
	if err != nil {
		t.Fatal(err)
	}

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()
	cfg.SetPermanently("overwrite", "true")

	src := "test/testdirectory/testfile"
	dest := "test/destdirectory/testfile"
	fileHandle, err := os.Create(dest)
	if err != nil {
		t.Fatal(err)
	}
	fileHandle.Close()

	c := &copy{[]string{src, dest}}
	err = c.Execute(cfg, logger)
	if err != nil {
		t.Error(err)
	}

	cmp := equalfile.New(nil, equalfile.Options{})
	if equal, err := cmp.CompareFile(src, dest); err != nil || !equal {
		t.Errorf("Expected %s to be copied to %s", src, dest)
	}

	_ = os.RemoveAll("test/destdirectory")
}

func TestExecute_copyfilenooverwrite(t *testing.T) {
	setup(t)
	defer teardown(t)

	err := os.Mkdir("test/destdirectory", 0755)
	if err != nil {
		t.Fatal(err)
	}

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()
	cfg.SetPermanently("overwrite", "false")

	src := "test/testdirectory/testfile"
	dest := "test/destdirectory/destfile"
	c := &copy{[]string{src, dest}}
	_ = c.Execute(cfg, logger)

	cmp := equalfile.New(nil, equalfile.Options{})
	if equal, err := cmp.CompareFile(src, dest); err != nil || !equal {
		t.Errorf("Expected %s to be copied to %s", src, dest)
	}

	_ = os.RemoveAll("test/destdirectory")
}

func TestExecute_copyfiledestexist(t *testing.T) {
	setup(t)
	defer teardown(t)

	err := os.Mkdir("test/destdirectory", 0755)
	if err != nil {
		t.Fatal(err)
	}

	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()
	cfg.SetPermanently("overwrite", "false")

	src := "test/testdirectory/testfile"
	dest := "test/destdirectory/testfile"
	fileHandle, err := os.Create(dest)
	if err != nil {
		t.Fatal(err)
	}
	fileHandle.Close()

	c := &copy{[]string{src, dest}}
	err = c.Execute(cfg, logger)
	if err == nil {
		t.Errorf("Should not be able to copy %s to %s", src, dest)
	}
	_ = os.RemoveAll("test/destdirectory")
}

func TestCanHandle_invalidline(t *testing.T) {
	line := "copy"
	cv := &CopyValidator{}
	if cv.CanHandle(line) {
		t.Errorf("%s should not be valid", line)
	}
}

func TestCanHandle_validline(t *testing.T) {
	line := "copy "
	cv := &CopyValidator{}
	if !cv.CanHandle(line) {
		t.Errorf("%s should be valid", line)
	}
}

func TestValidateArgs_oneinvalidarg(t *testing.T) {
	var args []string
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		args = []string{"hello", "/dev/null/hello"}
	} else {
		args = []string{"hello", "hello?"}
	}
	cv := &CopyValidator{}
	if cv.ValidateArgs(args) {
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
	cv := &CopyValidator{}
	if cv.ValidateArgs(args) {
		t.Errorf("%+q should not be valid", args)
	}
}

func TestValidateArgs_validargs(t *testing.T) {
	args := []string{"hello", "bye"}
	cv := &CopyValidator{}
	if !cv.ValidateArgs(args) {
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
	newFile.WriteString("This is a test file\n")
	newFile.Close()
}

func teardown(t *testing.T) {
	err := os.RemoveAll("test/testdirectory")
	if err != nil {
		t.Fatal(err)
	}
}
