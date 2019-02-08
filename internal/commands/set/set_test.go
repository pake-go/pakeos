package set

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/pake-go/pake-lib/config"
)

func TestExecute(t *testing.T) {
	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()

	optionName := "overwrite"
	optionValue := "true"

	s := &set{[]string{optionName, optionValue}}
	err := s.Execute(cfg, logger)
	if err != nil {
		t.Error(err)
	}

	value, err := cfg.Get("overwrite")
	if err != nil {
		t.Error(err)
	}
	if value != "true" {
		t.Errorf("Expected true but got %s", value)
	}

	cfg.SmartReset()
	value, err = cfg.Get("overwrite")
	if err != nil {
		t.Error(err)
	}
	if value != "true" {
		t.Errorf("Expected true but got %s", value)
	}

	cfg.SmartReset()
	_, err = cfg.Get("overwrite")
	if err == nil {
		t.Errorf("overwrite should no longer exist in config")
	}
}

func TestCanHandle_invalidline(t *testing.T) {
	line := "set"
	sv := &SetValidator{}
	if sv.CanHandle(line) {
		t.Errorf("%+q should not be valid", line)
	}
}

func TestCanHandle_validline(t *testing.T) {
	line := "set "
	sv := &SetValidator{}
	if !sv.CanHandle(line) {
		t.Errorf("%+q should be valid", line)
	}
}

func TestValidateArgs_invalidoptionname(t *testing.T) {
	args := []string{"o", "nil"}
	sv := &SetValidator{}
	if sv.ValidateArgs(args) == nil {
		t.Errorf("%+q should not be valid", args)
	}
}

func TestValidateArgs_invalidoptionvalue(t *testing.T) {
	args := []string{"overwrite", "nil"}
	sv := &SetValidator{}
	if sv.ValidateArgs(args) == nil {
		t.Errorf("%+q should not be valid", args)
	}
}

func TestValidateArgs_invalidoptionnamevalue(t *testing.T) {
	args := []string{"o", "t"}
	sv := &SetValidator{}
	if sv.ValidateArgs(args) == nil {
		t.Errorf("%+q should not be valid", args)
	}
}

func TestValidateArgs_validargs(t *testing.T) {
	args := []string{"overwrite", "true"}
	sv := &SetValidator{}
	if sv.ValidateArgs(args) != nil {
		t.Errorf("%+q should be valid", args)
	}
}
