package comments

import "strings"

type Validator struct {
}

func (v *Validator) IsValid(line string) bool {
	return strings.HasPrefix(line, "# ") || line == ""
}
