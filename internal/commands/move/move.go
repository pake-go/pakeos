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

type move struct {
	args []string
}

func New(args []string) pakelib.Command {
	return &move{
		args: args,
	}
}

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

type MoveValidator struct {
}

func (mv *MoveValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "move ")
}

func (mv *MoveValidator) ValidateArgs(args []string) bool {
	for _, arg := range args {
		if !validpath.Valid(arg) {
			return false
		}
	}
	return len(args) == 2
}
