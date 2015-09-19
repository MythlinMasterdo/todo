package command

import (
	"testing"

	"github.com/naoty/todo/todo"
)

func TestRightMoveTodo(t *testing.T) {
	todos := []todo.Todo{
		todo.NewTodo(1, "", "dummy1"),
		todo.NewTodo(2, "", "dummy2"),
		todo.NewTodo(3, "", "dummy3"),
		todo.NewTodo(4, "", "dummy4"),
	}

	move := newTodoMoveProcess("1", "3")
	result, _ := move(todos)

	expectations := []string{"dummy2", "dummy3", "dummy1", "dummy4"}
	for i, expectation := range expectations {
		if result[i].Title != expectation {
			t.Errorf("result[%d].Title expected: %q, actual: %q", i, expectation, result[i].Title)
		}
	}
}

func TestLeftMoveTodo(t *testing.T) {
	todos := []todo.Todo{
		todo.NewTodo(1, "", "dummy1"),
		todo.NewTodo(2, "", "dummy2"),
		todo.NewTodo(3, "", "dummy3"),
		todo.NewTodo(4, "", "dummy4"),
	}

	move := newTodoMoveProcess("3", "1")
	result, _ := move(todos)

	expectations := []string{"dummy3", "dummy1", "dummy2", "dummy4"}
	for i, expectation := range expectations {
		if result[i].Title != expectation {
			t.Errorf("result[%d].Title expected: %q, actual: %q", i, expectation, result[i].Title)
		}
	}
}
