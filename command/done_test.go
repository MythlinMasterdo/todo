package command

import (
	"testing"

	"github.com/naoty/todo/todo"
)

func TestDoneTodos(t *testing.T) {
	todos := []todo.Todo{
		todo.NewTodo(1, "", "dummy1"),
		todo.NewTodo(2, "", "dummy2"),
		todo.NewTodo(3, "", "dummy3"),
	}
	done := newTodoDoneProcess("2", "3")
	result, _ := done(todos)

	if result[0].Done != false {
		t.Errorf("result[0].Done expected: false, actual: true")
	}

	if result[1].Done != true {
		t.Errorf("result[1].Done expected: true, actual: false")
	}

	if result[2].Done != true {
		t.Errorf("result[2].Done expected: true, actual: false")
	}
}

func TestDoneSubTodos(t *testing.T) {
	todos := []todo.Todo{
		todo.NewTodo(1, "", "dummy1"),
		todo.NewTodo(1, "1", "dummy1-1"),
		todo.NewTodo(1, "1-1", "dummy1-1-1"),
	}
	done := newTodoDoneProcess("1-1", "1-1-1")
	result, _ := done(todos)

	if result[0].Done != false {
		t.Errorf("result[0].Done expected: false, actual: true")
	}

	if result[1].Done != true {
		t.Errorf("result[1].Done expected: true, actual: false")
	}

	if result[2].Done != true {
		t.Errorf("result[2].Done expected: true, actual: false")
	}
}
