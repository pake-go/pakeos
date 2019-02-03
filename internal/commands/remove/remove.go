package remove

import (
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

func (r *remove) Execute(cfg *config.Config) error {
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

func (rv *RemoveValidator) ValidateArgs(args []string) bool {
	for _, arg := range args {
		if !validpath.Valid(arg) {
			return false
		}
	}
	return len(args) == 1
}
