package install

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

type install struct {
	args []string
}

func New(args []string) pakelib.Command {
	return &install{
		args: args,
	}
}

func (i *install) Execute(cfg *config.Config, logger *log.Logger) error {
	logger.Printf("Installing %s\n", i.args[0])

	appRepo, appRepoErr := cfg.Get("appInstallationRepo")
	if appRepoErr != nil {
		return appRepoErr
	}
	appRepo, _ = pathutil.Expand(appRepo)
	distro, distroErr := cfg.Get("distro")
	if distroErr != nil {
		return distroErr
	}

	appName := i.args[0]
	scriptPath := filepath.Join(appRepo, "apps", appName, runtime.GOOS, distro, "install")
	cmd := exec.Command(scriptPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

type InstallValidator struct {
}

func (iv *InstallValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "install ")
}

func (iv *InstallValidator) ValidateArgs(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Expected there to be 1 argument, but got %d", len(args))
	}
	return nil
}
