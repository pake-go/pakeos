package uninstall

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	pakelib "github.com/pake-go/pake-lib"
	"github.com/pake-go/pake-lib/config"
	"github.com/pake-go/pakeos/internal/utils/pathutil"
)

type uninstall struct {
	args []string
}

func New(args []string) pakelib.Command {
	return &uninstall{
		args: args,
	}
}

func (u *uninstall) Execute(cfg *config.Config, logger *log.Logger) error {
	appRepo, appRepoErr := cfg.Get("appInstallationRepo")
	if appRepoErr != nil {
		return appRepoErr
	}
	appRepo, _ = pathutil.Expand(appRepo)
	distro, distroErr := cfg.Get("distro")
	if distroErr != nil {
		return distroErr
	}

	appName := u.args[0]
	scriptPath := filepath.Join(appRepo, "apps", appName, runtime.GOOS, distro, "uninstall")
	cmd := exec.Command(scriptPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

type UninstallValidator struct {
}

func (uv *UninstallValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "uninstall ")
}

func (uv *UninstallValidator) ValidateArgs(args []string) bool {
	return len(args) == 1
}
