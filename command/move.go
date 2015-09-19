package command

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/codegangsta/cli"
	"github.com/naoty/todo/todo"
)

var Move = cli.Command{
	Name:  "move",
	Usage: "move a TODO",
	Action: func(context *cli.Context) {
		status := ExecMove(context)
		os.Exit(status)
	},
}

func ExecMove(context *cli.Context) int {
	if len(context.Args()) < 2 {
		cli.ShowCommandHelp(context, "move")
		return 1
	}

	move := newTodoMoveProcess(context.Args()[0], context.Args()[1])

	file := todo.OpenFile()
	err := file.Update(move)
	if err != nil {
		log.Println(err)
		return 1
	}

	return 0
}

func newTodoMoveProcess(fromID, toID string) todo.TodoProcess {
	return func(todos []todo.Todo) ([]todo.Todo, error) {
		movedTodo, err := findTodo(todos, fromID)
		targetTodo, err := findTodo(todos, toID)
		if err != nil {
			return nil, err
		}

		isRightMove := (movedTodo.Order < targetTodo.Order)

		newTodos := make([]todo.Todo, len(todos))
		for i, t := range todos {
			newTodo := t

			if t.Order == movedTodo.Order {
				newTodo = todo.NewTodo(targetTodo.Order, t.ParentID, t.Title)
				newTodo.Done = t.Done
			}

			if isRightMove {
				if t.Order > movedTodo.Order && t.Order <= targetTodo.Order {
					newTodo = todo.NewTodo(t.Order-1, t.ParentID, t.Title)
					newTodo.Done = t.Done
				}
			} else {
				if t.Order >= targetTodo.Order && t.Order < movedTodo.Order {
					newTodo = todo.NewTodo(t.Order+1, t.ParentID, t.Title)
					newTodo.Done = t.Done
				}
			}

			newTodos[i] = newTodo
		}

		sort.Sort(todo.ByOrder(newTodos))

		return newTodos, nil
	}
}

func findTodo(todos []todo.Todo, id string) (todo.Todo, error) {
	for _, todo := range todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return todo.Todo{}, fmt.Errorf("TODO not found: %q", id)
}
