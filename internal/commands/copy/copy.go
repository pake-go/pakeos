package copy

import (
	"fmt"
	"log"
	"os"
	"strings"

	pakelib "github.com/pake-go/pake-lib"
	"github.com/pake-go/pake-lib/config"
	"github.com/pake-go/pakeos/internal/utils/copyutil"
	"github.com/pake-go/pakeos/internal/utils/pathutil"
	"github.com/pake-go/pakeos/internal/validators/validpath"
)

type copy struct {
	args []string
}

func New(args []string) pakelib.Command {
	return &copy{
		args: args,
	}
}

func (c *copy) Execute(cfg *config.Config, logger *log.Logger) error {
	sourcePath := c.args[0]
	if expandedPath, err := pathutil.Expand(sourcePath); err == nil {
		sourcePath = expandedPath
	}
	destinationPath := c.args[1]
	if expandedPath, err := pathutil.Expand(destinationPath); err == nil {
		destinationPath = expandedPath
	}
	destPathExists := pathutil.Exists(destinationPath)

	overwrite, overwriteErr := cfg.Get("overwrite")
	if destPathExists && (overwriteErr != nil || overwrite == "false") {
		errMsg := "Destination path %s already exists and `overwrite` not set to true"
		err := fmt.Errorf(errMsg, destinationPath)
		return err
	}
	if destPathExists {
		if err := os.RemoveAll(destinationPath); err != nil {
			return err
		}
	}

	sourceIsDir, err := pathutil.IsDir(sourcePath)
	if err != nil {
		return err
	} else if sourceIsDir {
		return copyutil.CopyDir(sourcePath, destinationPath)
	} else {
		return copyutil.CopyFile(sourcePath, destinationPath)
	}
}

type CopyValidator struct {
}

func (cv *CopyValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "copy ")
}

func (cv *CopyValidator) ValidateArgs(args []string) bool {
	for _, arg := range args {
		if !validpath.Valid(arg) {
			return false
		}
	}
	return len(args) == 2
}
