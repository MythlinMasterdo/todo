package todo

import (
	"fmt"
	"regexp"
)

type Todo struct {
	// ID is order concatenated to ParentID. The format is `<ParentID>-<order>`.
	ID string

	// ParentID is the ID of a parent todo. The value of todos without a parent
	// is "".
	ParentID string

	Title string
	Done  bool
}

func NewTodo(order int, parentID, title string) Todo {
	id := generateID(order, parentID)

	return Todo{
		ID:       id,
		ParentID: parentID,
		Title:    title,
		Done:     false,
	}
}

var pattern = regexp.MustCompile(`^([0-9]+-)*([0-9]+)?$`)

func ValidateID(id string) bool {
	matchedPositions := pattern.FindStringIndex(id)
	return len(matchedPositions) == 2
}

func generateID(order int, parentID string) string {
	if parentID == "" {
		return fmt.Sprintf("%d", order)
	} else {
		return fmt.Sprintf("%s-%d", parentID, order)
	}
}

func FilterTodos(todos []Todo, f func(Todo) bool) []Todo {
	newTodos := make([]Todo, 0)
	for _, todo := range todos {
		if f(todo) {
			newTodos = append(newTodos, todo)
		}
	}
	return newTodos
}

func ReorderTodos(todos []Todo, parentID string) []Todo {
	rootTodos := FilterTodos(todos, func(todo Todo) bool {
		return todo.ParentID == parentID
	})

	newTodos := make([]Todo, 0)
	for i, todo := range rootTodos {
		newTodo := NewTodo(i+1, todo.ParentID, todo.Title)
		newTodos = append(newTodos, newTodo)

		subTodos := ReorderTodos(todos, todo.ID)
		newTodos = append(newTodos, subTodos...)
	}
	return newTodos
}
