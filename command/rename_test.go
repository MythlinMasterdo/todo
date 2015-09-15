package command

import (
	"testing"

	"github.com/naoty/todo/todo"
)

func TestRenameTodo(t *testing.T) {
	todos := []todo.Todo{
		todo.NewTodo(1, "", "dummy1"),
	}
	rename := newTodoRenameProcess("1", "dummy2")
	result, _ := rename(todos)

	if result[0].Title != "dummy2" {
		t.Errorf("result[0].Title expected: %q, actual: %q", "dummy2", result[0].Title)
	}
}

func TestRenameSubTodo(t *testing.T) {
	todos := []todo.Todo{
		todo.NewTodo(1, "1", "dummy1-1"),
	}
	rename := newTodoRenameProcess("1-1", "dummy1-2")
	result, _ := rename(todos)

	if result[0].Title != "dummy1-2" {
		t.Errorf("result[0].Title expected: %q, actual: %q", "dummy2", result[0].Title)
	}
}
