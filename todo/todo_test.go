package todo

import (
	"testing"
)

func TestValidateID(t *testing.T) {
	if !ValidateID("") {
		t.Errorf("ValidateID(%q) expected: true, actual: false", "")
	}

	if !ValidateID("1") {
		t.Errorf("ValidateID(%q) expected: true, actual: false", "1")
	}

	if !ValidateID("1-2") {
		t.Errorf("ValidateID(%q) expected: true, actual: false", "1-2")
	}

	if !ValidateID("1-2-3") {
		t.Errorf("ValidateID(%q) expected: true, actual: false", "1-2-3")
	}

	if ValidateID("invalid") {
		t.Errorf("ValidateID(%q) expected: false, actual: true", "invalid")
	}
}
