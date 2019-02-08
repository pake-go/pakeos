package setdefault

import (
	"fmt"
	"log"
	"strings"

	pakelib "github.com/pake-go/pake-lib"
	"github.com/pake-go/pake-lib/config"
	"github.com/pake-go/pakeos/internal/validators/validconfig"
)

// The setdefault command is used to modify the config file for the long-term.
type setdefault struct {
	args []string
}

// New returns an instance of the setdefault command for modifying the config file until
// the next set command that modifies the same key.
func New(args []string) pakelib.Command {
	return &setdefault{
		args: args,
	}
}

// Execute runs the setdefault action and returns any error it encounters.
func (sd *setdefault) Execute(cfg *config.Config, logger *log.Logger) error {
	logger.Printf("Setting the default value for %s to %s\n", sd.args[0], sd.args[1])

	optionName := sd.args[0]
	optionValue := sd.args[1]
	cfg.SetPermanently(optionName, optionValue)
	return nil
}

// SetDefaultValidator represents the object used to check if a line is valid for the
// setdefault command.
type SetDefaultValidator struct {
}

// CanHandle reports if the given line represents source code that can be handled by the
// setdefault command.
func (sdv *SetDefaultValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "setdefault ")
}

// ValidateArgs checks to see if the given arguments are valid arguments for the
// setdefault command.
func (sdv *SetDefaultValidator) ValidateArgs(args []string) error {
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
