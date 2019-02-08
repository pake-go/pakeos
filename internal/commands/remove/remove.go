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

type remove struct {
	args []string
}

func New(args []string) pakelib.Command {
	return &remove{
		args: args,
	}
}

func (r *remove) Execute(cfg *config.Config, logger *log.Logger) error {
	logger.Printf("Removing %s\n", r.args[0])

	pathToDelete := r.args[0]
	if expandedPath, err := pathutil.Expand(pathToDelete); err == nil {
		pathToDelete = expandedPath
	}
	return os.RemoveAll(pathToDelete)
}

type RemoveValidator struct {
}

func (rv *RemoveValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "remove ")
}

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
