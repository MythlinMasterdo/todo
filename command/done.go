package command

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/naoty/todo/todo"
	"github.com/naoty/todo/todoutil"
)

var Done = cli.Command{
	Name:  "done",
	Usage: "Done TODOs",
	Action: func(context *cli.Context) {
		status := ExecDone(context)
		os.Exit(status)
	},
}

func ExecDone(context *cli.Context) int {
	if len(context.Args()) == 0 {
		cli.ShowCommandHelp(context, "done")
		return 1
	}

	done := newTodoDoneProcess(context.Args()...)

	file := todo.OpenFile()
	err := file.Update(done)
	if err != nil {
		log.Println(err)
		return 1
	}

	return 0
}

func newTodoDoneProcess(ids ...string) todo.TodoProcess {
	return func(todos []todo.Todo) ([]todo.Todo, error) {
		newTodos := make([]todo.Todo, len(todos))

		for i, todo := range todos {
			newTodo := todo
			if todoutil.ContainsString(ids, todo.ID) {
				newTodo.Done = true
			}
			newTodos[i] = newTodo
		}

		return newTodos, nil
	}
}
