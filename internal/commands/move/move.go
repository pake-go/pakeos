package move

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	pakelib "github.com/pake-go/pake-lib"
	"github.com/pake-go/pake-lib/config"
	"github.com/pake-go/pakeos/internal/utils/pathutil"
	"github.com/pake-go/pakeos/internal/validators/validpath"
)

// The move command is used to move a directory or file from one place to another.
type move struct {
	args []string
}

// New returns an instance of the move command for moving files.
func New(args []string) pakelib.Command {
	return &move{
		args: args,
	}
}

// Execute runs the move action and returns any error it encounters.
func (m *move) Execute(cfg *config.Config, logger *log.Logger) error {
	logger.Printf("Moving %s to %s\n", m.args[0], m.args[1])

	sourcePath := m.args[0]
	if expandedPath, err := pathutil.Expand(sourcePath); err == nil {
		sourcePath = expandedPath
	}
	destinationPath := m.args[1]
	if expandedPath, err := pathutil.Expand(destinationPath); err == nil {
		destinationPath = expandedPath
	}

	destPathExists := pathutil.Exists(destinationPath)
	srcIsDir, err := pathutil.IsDir(sourcePath)
	if err != nil {
		return err
	}
	destIsDir, err := pathutil.IsDir(destinationPath)
	if err != nil && destPathExists {
		return err
	}

	if destPathExists && destIsDir {
		basename := filepath.Base(sourcePath)
		return os.Rename(sourcePath, filepath.Join(destinationPath, basename))
	} else if destPathExists && !srcIsDir && !destIsDir {
		return os.Rename(sourcePath, destinationPath)
	} else if destPathExists {
		return fmt.Errorf("%s is a file which cannot be overwritten", destinationPath)
	} else {
		return os.Rename(sourcePath, destinationPath)
	}
}

// MoveValidator represents the object used to check if a line is valid for the move
// command.
type MoveValidator struct {
}

// CanHandle reports if the given line represents source code that can be handled by the
// move command.
func (mv *MoveValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "move ")
}

// Validate args checks to see if the given arguments are valid arguments for the move
// command.
func (mv *MoveValidator) ValidateArgs(args []string) error {
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
