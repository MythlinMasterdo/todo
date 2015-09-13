package command

import (
	"testing"

	"github.com/naoty/todo/todo"
)

func TestAddTodo(t *testing.T) {
	todo1 := todo.Todo{ID: "1", ParentID: "", Title: "dummy1", Done: false}
	add := newTodoAddProcess(2, "", "dummy2", false)

	result, _ := add([]todo.Todo{todo1})

	if len(result) != 2 {
		t.Errorf("len(result) expected: 2, actual: %d", len(result))
	}

	if result[0].ID != "1" {
		t.Errorf("result[0].ID expected: %q, actual: %q", "1", result[0].ID)
	}

	if result[1].ID != "2" {
		t.Errorf("result[1].ID expected: %q, actual: %q", "2", result[1].ID)
	}
}

func TestAddTodoOnce(t *testing.T) {
	todo1 := todo.Todo{ID: "1", ParentID: "", Title: "dummy1", Done: false}
	add := newTodoAddProcess(2, "", "dummy1", true)

	result, _ := add([]todo.Todo{todo1})

	if len(result) != 1 {
		t.Errorf("len(result) expected: 1, actual: %d", len(result))
	}

	if result[0].ID != "1" {
		t.Errorf("result[0].ID expected: %q, actual: %q", "1", result[0].ID)
	}
}

func TestAddSubTodoOnce(t *testing.T) {
	todo1 := todo.Todo{ID: "1", ParentID: "", Title: "dummy1", Done: false}
	add := newTodoAddProcess(1, "1", "dummy2", false)

	result, _ := add([]todo.Todo{todo1})

	if len(result) != 2 {
		t.Errorf("len(result) expected: 2, actual: %d", len(result))
	}

	if result[0].ID != "1" {
		t.Errorf("result[0].ID expected: %q, actual: %q", "1", result[0].ID)
	}

	if result[1].ID != "1-1" {
		t.Errorf("result[1].ID expected: %q, actual: %q", "1-1", result[1].ID)
	}

	if result[1].ParentID != "1" {
		t.Errorf("result[1].ID expected: %q, actual: %q", "1", result[1].ParentID)
	}
}
