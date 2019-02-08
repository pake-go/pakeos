package uninstall

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

	u := &uninstall{[]string{"hello"}}
	err := u.Execute(cfg, logger)
	if err != nil {
		t.Error(err)
	}
}

func TestCanHandle_invalidline(t *testing.T) {
	line := "uninstall"
	uv := &UninstallValidator{}
	if uv.CanHandle(line) {
		t.Errorf("%s should not be valid", line)
	}
}

func TestCanHandle_validline(t *testing.T) {
	line := "uninstall "
	uv := &UninstallValidator{}
	if !uv.CanHandle(line) {
		t.Errorf("%s should be valid", line)
	}
}

func TestValidateArgs_invalidarg(t *testing.T) {
	args := []string{"fd", "die"}
	uv := &UninstallValidator{}
	if uv.ValidateArgs(args) == nil {
		t.Errorf("%+q should not be valid", args)
	}
}

func TestValidateArgs_validarg(t *testing.T) {
	args := []string{"fd"}
	uv := &UninstallValidator{}
	if uv.ValidateArgs(args) != nil {
		t.Errorf("%+q should be valid", args)
	}
}
