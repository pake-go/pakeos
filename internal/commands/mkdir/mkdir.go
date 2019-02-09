package mkdir

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

// The mkdir command is used to create a directory
type mkdir struct {
	args []string
}

// New returns an instance of the mkdir command for mkaing files or directories
func New(args []string) pakelib.Command {
	return &mkdir{
		args: args,
	}
}

// Execute runs the mkdir action and returns any error it encounters
func (m *mkdir) Execute(cfg *config.Config, logger *log.Logger) error {
	logger.Printf("Creating %s\n", m.args[0])

	pathToCreate := m.args[0]
	if expandedPath, err := pathutil.Expand(pathToCreate); err == nil {
		pathToCreate = expandedPath
	}
	return os.MkdirAll(pathToCreate, 0755)
}

// MkdirValidator represents an object used to check if a line is valid for the mkdir
// command.
type MkdirValidator struct {
}

// CanHandle reports if the given line represents source code that can be handled by the
// mkdir command.
func (mv *MkdirValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "mkdir ")
}

func (mv *MkdirValidator) ValidateArgs(args []string) error {
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
