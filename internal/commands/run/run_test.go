package run

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/pake-go/pake-lib/config"
)

func TestExecute_runclass(t *testing.T) {
	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	cmd := "test/hi.class"

	r := &run{[]string{cmd}}
	out := capturer.CaptureOutput(func() {
		err := r.Execute(cfg, logger)
		if err != nil {
			t.Error(err)
		}
	})
	if out != "Hello World\n" {
		t.Errorf("Expected Hello World\n but got %s", out)
	}
}

func TestExecute_runpy(t *testing.T) {
	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	cmd := "test/hi.py"

	r := &run{[]string{cmd}}
	out := capturer.CaptureOutput(func() {
		err := r.Execute(cfg, logger)
		if err != nil {
			t.Error(err)
		}
	})
	if out != "Hello World\n" {
		t.Errorf("Expected Hello World\n but got %s", out)
	}
}

func TestExecute_runjava(t *testing.T) {
	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	cmd := "test/hi.java"

	r := &run{[]string{cmd}}
	out := capturer.CaptureOutput(func() {
		err := r.Execute(cfg, logger)
		if err != nil {
			t.Error(err)
		}
	})
	if out != "Hello World\n" {
		t.Errorf("Expected Hello World\n but got %s", out)
	}
}

func TestExecute_runsh(t *testing.T) {
	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	cmd := "test/hi.sh"

	r := &run{[]string{cmd}}
	out := capturer.CaptureOutput(func() {
		err := r.Execute(cfg, logger)
		if err != nil {
			t.Error(err)
		}
	})
	if out != "Hello World\n" {
		t.Errorf("Expected Hello World\n but got %s", out)
	}
}

func TestExecute_runmultipleargs(t *testing.T) {
	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	executable := "python"
	cmd := "test/hi.py"

	r := &run{[]string{executable, cmd}}
	out := capturer.CaptureOutput(func() {
		err := r.Execute(cfg, logger)
		if err != nil {
			t.Error(err)
		}
	})
	if out != "Hello World\n" {
		t.Errorf("Expected Hello World\n but got %s", out)
	}
}

func TestCanHandle_invalidline(t *testing.T) {
	line := "run"
	rv := &RunValidator{}
	if rv.CanHandle(line) {
		t.Errorf("%s should not be valid", line)
	}
}

func TestCanHandle_validline(t *testing.T) {
	line := "run "
	rv := &RunValidator{}
	if !rv.CanHandle(line) {
		t.Errorf("%s should be valid", line)
	}
}

func TestValidateArgs_invalidarg(t *testing.T) {
	args := []string{}
	rv := &RunValidator{}
	if rv.ValidateArgs(args) {
		t.Errorf("%+q should not be valid", args)
	}
}

func TestValidateArgs_validarg(t *testing.T) {
	args := []string{"python", "hi.py"}
	rv := &RunValidator{}
	if !rv.ValidateArgs(args) {
		t.Errorf("%+q should be valid", args)
	}
}
