package validconfig

import (
	"regexp"

	"github.com/pake-go/pakeos/internal/validators/validpath"
)

var validOptionsByRegex = map[string]string{
	"distro":    "arch|debian|ubuntu",
	"overwrite": "true|false",
}

var validOptionsByFunc = map[string]func(string) bool{
	"appInstallationRepo": validpath.Valid,
}

func Valid(optionName, optionValue string) bool {
	if validateRegex, found := validOptionsByRegex[optionName]; found {
		isValid, _ := regexp.MatchString(validateRegex, optionValue)
		return isValid
	}
	if validateFunc, found := validOptionsByFunc[optionName]; found {
		return validateFunc(optionValue)
	}
	return false
}
