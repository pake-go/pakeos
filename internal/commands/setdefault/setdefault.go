package setdefault

import (
	"fmt"
	"log"
	"strings"

	pakelib "github.com/pake-go/pake-lib"
	"github.com/pake-go/pake-lib/config"
	"github.com/pake-go/pakeos/internal/validators/validconfig"
)

type setdefault struct {
	args []string
}

func New(args []string) pakelib.Command {
	return &setdefault{
		args: args,
	}
}

func (sd *setdefault) Execute(cfg *config.Config, logger *log.Logger) error {
	logger.Printf("Setting the default value for %s to %s\n", sd.args[0], sd.args[1])

	optionName := sd.args[0]
	optionValue := sd.args[1]
	cfg.SetPermanently(optionName, optionValue)
	return nil
}

type SetDefaultValidator struct {
}

func (sdv *SetDefaultValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "setdefault ")
}

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
