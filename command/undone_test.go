package command

import (
	"testing"

	"github.com/naoty/todo/todo"
)

func TestUndoneTodos(t *testing.T) {
	todos := []todo.Todo{
		todo.Todo{ID: "1", ParentID: "", Title: "dummy1", Done: true},
		todo.Todo{ID: "2", ParentID: "", Title: "dummy2", Done: true},
		todo.Todo{ID: "3", ParentID: "", Title: "dummy3", Done: true},
	}
	undone := newTodoUndoneProcess("2", "3")
	result, _ := undone(todos)

	if result[0].Done != true {
		t.Errorf("result[0].Done expected: true, actual: false")
	}

	if result[1].Done != false {
		t.Errorf("result[1].Done expected: false, actual: true")
	}

	if result[2].Done != false {
		t.Errorf("result[2].Done expected: false, actual: true")
	}
}

func TestUndoneSubTodos(t *testing.T) {
	todos := []todo.Todo{
		todo.Todo{ID: "1", ParentID: "", Title: "dummy1", Done: true},
		todo.Todo{ID: "1-1", ParentID: "1", Title: "dummy1-1", Done: true},
		todo.Todo{ID: "1-1-1", ParentID: "1-1", Title: "dummy1-1-1", Done: true},
	}
	undone := newTodoUndoneProcess("1-1", "1-1-1")
	result, _ := undone(todos)

	if result[0].Done != true {
		t.Errorf("result[0].Done expected: true, actual: false")
	}

	if result[1].Done != false {
		t.Errorf("result[1].Done expected: false, actual: true")
	}

	if result[2].Done != false {
		t.Errorf("result[2].Done expected: false, actual: true")
	}
}
