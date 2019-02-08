package uninstall

import (
	"fmt"
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

// The uninstall command is used to uninstall applications utilizing instructions in a
// specifically formatted app installation repo.
type uninstall struct {
	args []string
}

// New returns an instance of the uninstall command for uninstalling applications
// utilizing scripts in a specifically formatted app installation repo.
func New(args []string) pakelib.Command {
	return &uninstall{
		args: args,
	}
}

// Execute runs teh uninstall action and returns any error it encounters.
func (u *uninstall) Execute(cfg *config.Config, logger *log.Logger) error {
	logger.Printf("Uninstalling %s\n", u.args[0])

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

// UninstallValidator represents the object used to check if a line is valid for the
// uninstall command.
type UninstallValidator struct {
}

// CanHandle reports if the given line represents source code that can be handled by the
// install command.
func (uv *UninstallValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "uninstall ")
}

// ValidateArgs checks to see if the given arguments are valid arguments for the uninstall
// command.
func (uv *UninstallValidator) ValidateArgs(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Expected there to be 1 argument, but got %d", len(args))
	}
	return nil
}
