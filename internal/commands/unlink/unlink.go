package unlink

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

type unlink struct {
	args []string
}

func New(args []string) pakelib.Command {
	return &unlink{
		args: args,
	}
}

func (u *unlink) Execute(cfg *config.Config, logger *log.Logger) error {
	logger.Printf("Unlinking %s\n", u.args[0])

	pathToUnlink := u.args[0]
	if expandedPath, err := pathutil.Expand(pathToUnlink); err == nil {
		pathToUnlink = expandedPath
	}

	isSymlink, err := pathutil.IsSymlink(pathToUnlink)
	if err != nil {
		return err
	} else if isSymlink {
		return os.Remove(pathToUnlink)
	} else {
		return fmt.Errorf("%s is not a symlink, cannot unlink!", pathToUnlink)
	}
}

type UnlinkValidator struct {
}

func (uv *UnlinkValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "unlink ")
}

func (uv *UnlinkValidator) ValidateArgs(args []string) error {
	for _, arg := range args {
		if !validpath.Valid(arg) {
			return fmt.Errorf("%s is not a valid argument", arg)
		}
	}
	if len(args) != 1 {
		return fmt.Errorf("Expected there to be 1 argument, but got %d", len(args))
	}
	return nil
}
