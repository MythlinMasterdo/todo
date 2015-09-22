package command

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/naoty/todo/todo"
	"github.com/naoty/todo/todoutil"
)

var Undone = cli.Command{
	Name:  "undone",
	Usage: "Undone TODOs",
	Action: func(context *cli.Context) {
		status := ExecUndone(context)
		os.Exit(status)
	},
}

func ExecUndone(context *cli.Context) int {
	if len(context.Args()) == 0 {
		cli.ShowCommandHelp(context, "undone")
		return 1
	}

	undone := newTodoUndoneProcess(context.Args()...)

	file := todo.OpenFile()
	err := file.Update(undone)
	if err != nil {
		log.Println(err)
		return 1
	}

	return 0
}

func newTodoUndoneProcess(ids ...string) todo.TodoProcess {
	return func(todos todo.Todos) (todo.Todos, error) {
		newTodos := make(todo.Todos, len(todos))
		for i, todo := range todos {
			newTodo := todo
			if todoutil.ContainsString(ids, todo.ID) {
				newTodo.Done = false
			}
			newTodos[i] = newTodo
		}
		return newTodos, nil
	}
}
