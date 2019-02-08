package setdefault

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

	sd := &setdefault{[]string{optionName, optionValue}}
	err := sd.Execute(cfg, logger)
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
	value, err = cfg.Get("overwrite")
	if err != nil {
		t.Error(err)
	}
	if value != "true" {
		t.Errorf("Expected true but got %s", value)
	}
}

func TestCanHandle_invalidline(t *testing.T) {
	line := "setdefault"
	sdv := &SetDefaultValidator{}
	if sdv.CanHandle(line) {
		t.Errorf("%+q should not be valid", line)
	}
}

func TestCanHandle_validline(t *testing.T) {
	line := "setdefault "
	sdv := &SetDefaultValidator{}
	if !sdv.CanHandle(line) {
		t.Errorf("%+q should be valid", line)
	}
}

func TestValidateArgs_invalidoptionname(t *testing.T) {
	args := []string{"o", "nil"}
	sdv := &SetDefaultValidator{}
	if sdv.ValidateArgs(args) == nil {
		t.Errorf("%+q should not be valid", args)
	}
}

func TestValidateArgs_invalidoptionvalue(t *testing.T) {
	args := []string{"overwrite", "nil"}
	sdv := &SetDefaultValidator{}
	if sdv.ValidateArgs(args) == nil {
		t.Errorf("%+q should not be valid", args)
	}
}

func TestValidateArgs_invalidoptionnamevalue(t *testing.T) {
	args := []string{"o", "t"}
	sdv := &SetDefaultValidator{}
	if sdv.ValidateArgs(args) == nil {
		t.Errorf("%+q should not be valid", args)
	}
}

func TestValidateArgs_validargs(t *testing.T) {
	args := []string{"overwrite", "true"}
	sdv := &SetDefaultValidator{}
	if sdv.ValidateArgs(args) != nil {
		t.Errorf("%+q should be valid", args)
	}
}
