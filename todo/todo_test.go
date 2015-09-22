package todo

import (
	"testing"
)

func TestGetOrders(t *testing.T) {
	todo := NewTodo(1, "", "dummy1")
	if todo.GetOrders()[0] != todo.Order {
		t.Errorf("todo.getOrders()[0] expected: %d, actual: %d", todo.Order, todo.GetOrders()[0])
	}
}

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
