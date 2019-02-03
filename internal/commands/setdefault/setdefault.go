package setdefault

import (
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

func (sdv *SetDefaultValidator) ValidateArgs(args []string) bool {
	if len(args) != 2 {
		return false
	}
	optionName := args[0]
	optionValue := args[1]
	return validconfig.Valid(optionName, optionValue)
}
