package set

import (
	"fmt"
	"log"
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

func (s *set) Execute(cfg *config.Config, logger *log.Logger) error {
	logger.Printf("Temporarily setting %s to %s\n", s.args[0], s.args[1])

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
