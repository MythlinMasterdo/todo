package command

import (
	"testing"

	"github.com/naoty/todo/todo"
)

func TestDeleteTodos(t *testing.T) {
	todos := []todo.Todo{
		todo.NewTodo(1, "", "dummy1"),
		todo.NewTodo(2, "", "dummy2"),
		todo.NewTodo(3, "", "dummy3"),
	}
	delete := newTodoDeleteProcess("2", "3")
	result, _ := delete(todos)

	if len(result) != 1 {
		t.Errorf("len(result) expected: 1, actual: %d", len(result))
	}
}

func TestDeleteSubTodos(t *testing.T) {
	todos := []todo.Todo{
		todo.NewTodo(1, "", "dummy1"),
		todo.NewTodo(1, "1", "dummy1-1"),
		todo.NewTodo(1, "1-1", "dummy1-1-1"),
	}
	delete := newTodoDeleteProcess("1-1", "1-1-1")
	result, _ := delete(todos)

	if len(result) != 1 {
		t.Errorf("len(result) expected: 1, actual: %d", len(result))
	}
}

func TestDeleteAndReorderTodos(t *testing.T) {
	todos := []todo.Todo{
		todo.NewTodo(1, "", "dummy1"),
		todo.NewTodo(2, "", "dummy2"),
	}
	delete := newTodoDeleteProcess("1")
	result, _ := delete(todos)

	if result[0].ID != "1" {
		t.Errorf("result[0].ID expected: %q, actual: %q", "1", result[0].ID)
	}
}
