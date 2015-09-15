package command

import (
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/naoty/todo/todo"
)

var Rename = cli.Command{
	Name:  "rename",
	Usage: "Rename a TODO",
	Action: func(context *cli.Context) {
		status := ExecRename(context)
		os.Exit(status)
	},
}

func ExecRename(context *cli.Context) int {
	if len(context.Args()) < 2 {
		cli.ShowCommandHelp(context, "rename")
		return 1
	}

	title := strings.Join(context.Args()[1:], " ")
	rename := newTodoRenameProcess(context.Args()[0], title)

	file := todo.OpenFile()
	err := file.Update(rename)
	if err != nil {
		log.Println(err)
		return 1
	}

	return 0
}

func newTodoRenameProcess(id, title string) todo.TodoProcess {
	return func(todos []todo.Todo) ([]todo.Todo, error) {
		newTodos := make([]todo.Todo, len(todos))

		for i, todo := range todos {
			newTodo := todo
			if todo.ID == id {
				newTodo.Title = title
			}
			newTodos[i] = newTodo
		}

		return newTodos, nil
	}
}
