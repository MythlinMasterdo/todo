package todo

type Todos []Todo

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
