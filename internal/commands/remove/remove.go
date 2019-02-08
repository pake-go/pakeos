package remove

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

// The remove command is used to remove a file or directory.
type remove struct {
	args []string
}

// New returns an instance of the remove command for removing files or directories.
func New(args []string) pakelib.Command {
	return &remove{
		args: args,
	}
}

// Execute runs the remove action and returns any error it encounters.
func (r *remove) Execute(cfg *config.Config, logger *log.Logger) error {
	logger.Printf("Removing %s\n", r.args[0])

	pathToDelete := r.args[0]
	if expandedPath, err := pathutil.Expand(pathToDelete); err == nil {
		pathToDelete = expandedPath
	}
	return os.RemoveAll(pathToDelete)
}

// RemoveValidator represents an object used to check if a line is valid for the remove
// command.
type RemoveValidator struct {
}

// CanHandle reports if the given line represents source code that can be handled by the
// remove command.
func (rv *RemoveValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "remove ")
}

// ValidateArgs checks to see if the given arguments are valid arguments for the remove
// commands..
func (rv *RemoveValidator) ValidateArgs(args []string) error {
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
