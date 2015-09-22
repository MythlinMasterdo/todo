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

func (todos Todos) GroupBy(f func(Todo) string) map[string]Todos {
	groups := make(map[string]Todos)
	for _, todo := range todos {
		key := f(todo)
		groups[key] = append(groups[key], todo)
	}
	return groups
}

func (todos Todos) Compact() Todos {
	newTodos := make(Todos, 0)
	group := todos.GroupBy(func(todo Todo) string {
		return todo.ParentID
	})
	for _, _todos := range group {
		for i, todo := range _todos {
			newTodo := NewTodo(i+1, todo.ParentID, todo.Title)
			newTodo.Done = todo.Done
			newTodos = append(newTodos, newTodo)
		}
	}
	return newTodos
}
