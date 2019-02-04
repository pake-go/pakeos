package link

import (
	"fmt"
	"log"
	"os"
	"strings"

	pakelib "github.com/pake-go/pake-lib"
	"github.com/pake-go/pake-lib/config"
	"github.com/pake-go/pakeos/internal/utils/pathutil"
	"github.com/pake-go/pakeos/internal/validators/validpath"
)

type link struct {
	args []string
}

func New(args []string) pakelib.Command {
	return &link{
		args: args,
	}
}

func (l *link) Execute(cfg *config.Config, logger *log.Logger) error {
	logger.Printf("Linking %s to %s\n", l.args[0], l.args[1])

	sourcePath := l.args[0]
	if expandedPath, err := pathutil.Expand(sourcePath); err == nil {
		sourcePath = expandedPath
	}
	destinationPath := l.args[1]
	if expandedPath, err := pathutil.Expand(destinationPath); err == nil {
		destinationPath = expandedPath
	}
	destPathExists := pathutil.Exists(destinationPath)

	overwrite, overwriteErr := cfg.Get("overwrite")
	if overwriteErr != nil || overwrite == "false" {
		errMsg := "Destination path %s already exists and `overwrite` not set to true"
		return fmt.Errorf(errMsg, destinationPath)
	}
	if destPathExists {
		if err := os.RemoveAll(destinationPath); err != nil {
			return err
		}
	}
	return os.Symlink(sourcePath, destinationPath)
}

type LinkValidator struct {
}

func (lv *LinkValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "link ")
}

func (lv *LinkValidator) ValidateArgs(args []string) bool {
	for _, arg := range args {
		if !validpath.Valid(arg) {
			return false
		}
	}
	return len(args) == 2
}
