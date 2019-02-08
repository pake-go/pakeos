package validconfig

import (
	"regexp"

	"github.com/pake-go/pakeos/internal/validators/validpath"
)

// ValidOptionsByRegex defines a map of valid option names and the corresponding
// regex for validating the option value.
var validOptionsByRegex = map[string]string{
	"distro":    "arch|debian|ubuntu",
	"overwrite": "true|false",
}

// ValidOptionsByFunc defines a map of valid option names and the corresponding
// function for validating the option value.
var validOptionsByFunc = map[string]func(string) bool{
	"appInstallationRepo": validpath.Valid,
}

// Valid checks to see if the given option name and option value is a valid pair
// for the configuration.
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
