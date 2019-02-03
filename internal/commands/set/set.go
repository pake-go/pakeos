package set

import (
	"strings"

	pakelib "github.com/pake-go/pake-lib"
	"github.com/pake-go/pake-lib/config"
	"github.com/pake-go/pakeos/internal/validators/validconfig"
)

type set struct {
	args []string
}

func New(args []string) pakelib.Command {
	return &set{
		args: args,
	}
}

func (s *set) Execute(cfg *config.Config) error {
	optionName := s.args[0]
	optionValue := s.args[1]
	cfg.SetTemporarily(optionName, optionValue)
	return nil
}

type SetValidator struct {
}

func (sv *SetValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "set ")
}

func (sv *SetValidator) ValidateArgs(args []string) bool {
	if len(args) != 2 {
		return false
	}
	optionName := args[0]
	optionValue := args[1]
	return validconfig.Valid(optionName, optionValue)
}
