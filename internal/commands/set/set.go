package set

import (
	"fmt"
	"log"
	"strings"

	pakelib "github.com/pake-go/pake-lib"
	"github.com/pake-go/pake-lib/config"
	"github.com/pake-go/pakeos/internal/validators/validconfig"
)

// The set command is used for modifying the config file temporarily.
type set struct {
	args []string
}

// New returns an instance of the set command for modifying the config file temporarily.
func New(args []string) pakelib.Command {
	return &set{
		args: args,
	}
}

// Execute runs the set action and returns any error it encounters.
func (s *set) Execute(cfg *config.Config, logger *log.Logger) error {
	logger.Printf("Temporarily setting %s to %s\n", s.args[0], s.args[1])

	optionName := s.args[0]
	optionValue := s.args[1]
	cfg.SetTemporarily(optionName, optionValue)
	return nil
}

// SetValidator represents the object used to check if a line is valid for the set
// command.
type SetValidator struct {
}

// CanHandle reports if the given line represents source code that can be handled by the
// set command.
func (sv *SetValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "set ")
}

// ValidateArgs checks to see if the given arguments are valid arguments for the set
// command.
func (sv *SetValidator) ValidateArgs(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Expected there to be 2 arguments, but got %d", len(args))
	}
	optionName := args[0]
	optionValue := args[1]
	if !validconfig.Valid(optionName, optionValue) {
		return fmt.Errorf("%s, %s is not a valid pair", optionName, optionValue)
	}
	return nil
}
