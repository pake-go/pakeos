package move

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
	return os.Rename(sourcePath, destinationPath)
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
