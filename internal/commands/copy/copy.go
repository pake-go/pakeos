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

// The copy command is used to copy a directory or file from one place to another.
type copy struct {
	args []string
}

// New returns an instance of the copy command for copying files/directories.
func New(args []string) pakelib.Command {
	return &copy{
		args: args,
	}
}

// Execute runs the copy action and returns any error it encounters.
func (c *copy) Execute(cfg *config.Config, logger *log.Logger) error {
	logger.Printf("Copying %s to %s\n", c.args[0], c.args[1])

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

// CopyValidator represents the object used to check if a line is valid for the copy
// command.
type CopyValidator struct {
}

// CanHandle reports if the given line represents source code that can be handled by the
// copy command.
func (cv *CopyValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "copy ")
}

// ValidateArgs checks to see if the given arguments are valid arguments for the copy
// command.
func (cv *CopyValidator) ValidateArgs(args []string) error {
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
