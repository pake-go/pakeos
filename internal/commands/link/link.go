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

// The link command is used to create a symlnk for a file or directory.
type link struct {
	args []string
}

// New returns an instance of the link command for creating symlinks.
func New(args []string) pakelib.Command {
	return &link{
		args: args,
	}
}

// Execute runs the link action and returns any error it contains.
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
	destPathExists, err := pathutil.IsSymlink(destinationPath)
	if err != nil {
		destPathExists = pathutil.Exists(destinationPath)
	}

	overwrite, overwriteErr := cfg.Get("overwrite")
	if destPathExists && (overwriteErr != nil || overwrite == "false") {
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

// LinkValidator represents the object used to check if a line is valid for the link
// command.
type LinkValidator struct {
}

// CanHandle reports if the given line represents source code that can be handled by the
// link command.
func (lv *LinkValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "link ")
}

// ValidateArgs checks to see if the given arguments are valid arguments for the link
// command.
func (lv *LinkValidator) ValidateArgs(args []string) error {
	for _, arg := range args {
		if !validpath.Valid(arg) {
			return fmt.Errorf("%s is not a valid argument", arg)
		}
	}
	if len(args) != 2 {
		return fmt.Errorf("Expected there to be 2 arguments, but got %d", len(args))
	}
	return nil
}
