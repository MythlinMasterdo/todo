package todo

import (
	"fmt"
)

type Todos []Todo

func (todos Todos) Find(id string) (Todo, error) {
	for _, todo := range todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return Todo{}, fmt.Errorf("TODO not found: %q", id)
}

func (todos Todos) Filter(f func(Todo) bool) Todos {
	newTodos := make(Todos, 0)
	for _, todo := range todos {
		if f(todo) {
			newTodos = append(newTodos, todo)
		}
	}
	return newTodos
}

func (todos Todos) Reorder(parentID string) Todos {
	rootTodos := todos.Filter(func(todo Todo) bool {
		return todo.ParentID == parentID
	})

	newTodos := make(Todos, 0)
	for i, todo := range rootTodos {
		newTodo := NewTodo(i+1, todo.ParentID, todo.Title)
		newTodos = append(newTodos, newTodo)

		subTodos := todos.Reorder(todo.ID)
		newTodos = append(newTodos, subTodos...)
	}
	return newTodos
}
