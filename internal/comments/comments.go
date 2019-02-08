package comments

import "strings"

// Validator represents an object for checking if a comment is valid or
// not.
type Validator struct {
}

// IsValid checks if the given line is a valid comment or not.
func (v *Validator) IsValid(line string) bool {
	return strings.HasPrefix(line, "# ") || line == ""
}
