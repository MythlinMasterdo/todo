package todo

import (
	"testing"
)

func TestReorderTodos(t *testing.T) {
	todos := Todos{
		NewTodo(1, "", "dummy1"),
		NewTodo(3, "", "dummy3"),
	}
	result := todos.Reorder("")

	if result[0].ID != "1" {
		t.Errorf("result[0].ID expected: %q, actual: %q", "1", result[0].ID)
	}

	if result[1].ID != "2" {
		t.Errorf("result[1].ID expected: %q, actual: %d", "2", result[1].ID)
	}
}

func TestReorderNestedTodos(t *testing.T) {
	todos := Todos{
		NewTodo(1, "", "dummy1"),
		NewTodo(1, "1", "dummy1-1"),
		NewTodo(3, "1", "dummy1-3"),
	}
	result := todos.Reorder("")

	if result[0].ID != "1" {
		t.Errorf("result[0].ID expected: %q, actual: %q", "1", result[0].ID)
	}

	if result[1].ID != "1-1" {
		t.Errorf("result[1].ID expected: %q, actual: %d", "1-1", result[1].ID)
	}

	if result[2].ID != "1-2" {
		t.Errorf("result[2].ID expected: %q, actual: %d", "1-2", result[2].ID)
	}
}
