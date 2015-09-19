package todo

import (
	"sort"
	"testing"
)

func TestSortTodos(t *testing.T) {
	todos := []Todo{
		NewTodo(2, "", "dummy2"),
		NewTodo(1, "1-1", "dummy1-1-1"),
		NewTodo(1, "1", "dummy1-1"),
		NewTodo(1, "2", "dummy2-1"),
		NewTodo(1, "", "dummy1"),
	}
	sort.Sort(ByOrder(todos))

	expectations := []string{"dummy1", "dummy1-1", "dummy1-1-1", "dummy2", "dummy2-1"}
	for i, expectation := range expectations {
		if todos[i].Title != expectation {
			t.Errorf("todos[%d].Title expected: %q, actual: %q", i, expectation, todos[i].Title)
		}
	}
}
