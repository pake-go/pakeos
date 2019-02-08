package install

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/pake-go/pake-lib/config"
)

func TestExecute(t *testing.T) {
	logger := log.New(ioutil.Discard, "", 0)
	cfg := config.New()
	cfg.SetPermanently("appInstallationRepo", "test")
	cfg.SetPermanently("distro", "arch")

	i := &install{[]string{"hello"}}
	err := i.Execute(cfg, logger)
	if err != nil {
		t.Error(err)
	}
}

func TestCanHandle_invalidline(t *testing.T) {
	line := "install"
	iv := &InstallValidator{}
	if iv.CanHandle(line) {
		t.Errorf("%s should not be valid", line)
	}
}

func TestCanHandle_validline(t *testing.T) {
	line := "install "
	iv := &InstallValidator{}
	if !iv.CanHandle(line) {
		t.Errorf("%s should be valid", line)
	}
}

func TestValidateArgs_invalidarg(t *testing.T) {
	args := []string{"bat", "die"}
	iv := &InstallValidator{}
	if iv.ValidateArgs(args) {
		t.Errorf("%+q should not be valid", args)
	}
}

func TestValidateArgs_validarg(t *testing.T) {
	args := []string{"bat"}
	iv := &InstallValidator{}
	if !iv.ValidateArgs(args) {
		t.Errorf("%+q should be valid", args)
	}
}
