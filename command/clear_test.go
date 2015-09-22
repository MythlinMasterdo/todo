package command

import (
	"testing"

	"github.com/naoty/todo/todo"
)

func TestClearTodos(t *testing.T) {
	todos := todo.Todos{
		todo.Todo{ID: "1", ParentID: "", Title: "dummy1", Done: false},
		todo.Todo{ID: "2", ParentID: "", Title: "dummy2", Done: true},
		todo.Todo{ID: "3", ParentID: "", Title: "dummy3", Done: true},
	}
	clear := newTodoClearProcess()
	result, _ := clear(todos)

	if len(result) != 1 {
		t.Errorf("len(result) expected: 1, actual: %d", len(result))
	}

	if result[0].ID != "1" {
		t.Errorf("result[0].ID expected: %q, actual: %q", "1", result[0].ID)
	}
}

func TestClearAndCompactTodos(t *testing.T) {
	todos := todo.Todos{
		todo.Todo{ID: "1", ParentID: "", Title: "dummy1", Done: true},
		todo.Todo{ID: "2", ParentID: "", Title: "dummy2", Done: false},
	}
	clear := newTodoClearProcess()
	result, _ := clear(todos)

	if result[0].ID == "2" {
		t.Errorf("result[0].ID not expected: %q, actual: %q", "2", "2")
	}
}
