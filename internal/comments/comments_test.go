package comments

import "testing"

func TestIsValid_hashcomment(t *testing.T) {
	line := "# this is a comment"
	v := &Validator{}
	if !v.IsValid(line) {
		t.Errorf("Expected %s to be a valid comment", line)
	}
}

func TestIsValid_emptyline(t *testing.T) {
	line := ""
	v := &Validator{}
	if !v.IsValid(line) {
		t.Errorf("Expected %s to be a valid comment", line)
	}
}

func TestIsValid_notcomment(t *testing.T) {
	line := "this is not a comment"
	v := &Validator{}
	if v.IsValid(line) {
		t.Errorf("Expected %s to be an invalid comment", line)
	}
}
