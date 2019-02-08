package validconfig

import "testing"
import "github.com/lucasjones/reggen"

func TestValid_invalidconfigoptionname(t *testing.T) {
	optionName := ""
	optionValue := ""
	if Valid(optionName, optionValue) {
		t.Errorf("%s should not be valid", optionName)
	}
}

func TestValid_invalidconfigoptionvalue(t *testing.T) {
	optionName := "testing"
	optionValue := "false"

	if testValue, ok := validOptionsByRegex["testing"]; ok {
		if testValueTwo, ok := validOptionsByFunc["testing"]; ok {
			validOptionsByRegex["tesing"] = "true"
			validOptionsByFunc["testing"] = func(str string) bool {
				return false
			}
			if Valid(optionName, optionValue) {
				t.Errorf("%s should not be valid for %s", optionValue, optionName)
			}
			validOptionsByRegex["testing"] = testValue
			validOptionsByFunc["testing"] = testValueTwo
		} else {
			validOptionsByRegex["testing"] = "true"
			if Valid(optionName, optionValue) {
				t.Errorf("%s should not be valid for %s", optionValue, optionName)
			}
			validOptionsByRegex["testing"] = testValue
		}
	} else if testValue, ok := validOptionsByFunc["testing"]; ok {
		validOptionsByFunc["testing"] = func(str string) bool {
			return false
		}
		if Valid(optionName, optionValue) {
			t.Errorf("%s should not be valid for %s", optionValue, optionName)
		}
		validOptionsByFunc["testing"] = testValue
	} else {
		validOptionsByRegex["testing"] = "true"
		if Valid(optionName, optionValue) {
			t.Errorf("%s should not be valid for %s", optionValue, optionName)
		}
		delete(validOptionsByRegex, "testing")
	}
}

func TestValid_validconfigoption(t *testing.T) {
	for optionName, regex := range validOptionsByRegex {
		gen, err := reggen.NewGenerator(regex)
		if err != nil {
			t.Error(err)
		}

		optionValue := gen.Generate(1)
		if !Valid(optionName, optionValue) {
			t.Errorf("%s should be valid for %s", optionValue, optionName)
		}
	}

	if funcValidator, ok := validOptionsByFunc["testing"]; ok {
		validOptionsByFunc["testing"] = func(str string) bool {
			return str == "testing"
		}
		if !Valid("testing", "testing") {
			t.Errorf("testing should be valid for testing")
		}
		validOptionsByFunc["testing"] = funcValidator
	} else {
		validOptionsByFunc["testing"] = func(str string) bool {
			return str == "testing"
		}
		if !Valid("testing", "testing") {
			t.Errorf("testing should be valid for testing")
		}
		delete(validOptionsByFunc, "testing")
	}
}
