package command

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/naoty/todo/todo"
	"github.com/naoty/todo/todoutil"
)

var Delete = cli.Command{
	Name:  "delete",
	Usage: "Delete a TODO",
	Action: func(context *cli.Context) {
		status := ExecDelete(context)
		os.Exit(status)
	},
}

func ExecDelete(context *cli.Context) int {
	if len(context.Args()) == 0 {
		cli.ShowCommandHelp(context, "delete")
		return 1
	}

	delete := newTodoDeleteProcess(context.Args()...)

	file := todo.OpenFile()
	err := file.Update(delete)
	if err != nil {
		log.Println(err)
		return 1
	}

	return 0
}

func newTodoDeleteProcess(ids ...string) todo.TodoProcess {
	return func(todos []todo.Todo) ([]todo.Todo, error) {
		newTodos := make([]todo.Todo, 0)
		for _, todo := range todos {
			if todoutil.ContainsString(ids, todo.ID) {
				continue
			}
			newTodos = append(newTodos, todo)
		}
		newTodos = todo.ReorderTodos(newTodos, "")
		return newTodos, nil
	}
}
