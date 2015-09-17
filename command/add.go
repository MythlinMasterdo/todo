package command

import (
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/naoty/todo/todo"
)

var Add = cli.Command{
	Name:  "add",
	Usage: "Add a TODO",
	Action: func(context *cli.Context) {
		status := ExecAdd(context)
		os.Exit(status)
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "parent, p",
			Usage: "Add the sub-TODO under a specified TODO",
		},
		cli.BoolFlag{
			Name:  "once, o",
			Usage: "Add the TODO only if it exists",
		},
	},
}

func ExecAdd(context *cli.Context) int {
	if len(context.Args()) == 0 {
		cli.ShowCommandHelp(context, "add")
		return 1
	}

	parentID := context.String("parent")
	if !todo.ValidateID(parentID) {
		log.Printf("invalid parent id: %q", parentID)
		return 1
	}

	order := getNextOrder(parentID)
	title := strings.Join(context.Args(), " ")
	add := newTodoAddProcess(order, parentID, title, context.Bool("once"))

	file := todo.OpenFile()
	err := file.Update(add)
	if err != nil {
		log.Println(err)
		return 1
	}

	return 0
}

func newTodoAddProcess(order int, parentID, title string, isOnce bool) todo.TodoProcess {
	return func(todos []todo.Todo) ([]todo.Todo, error) {
		if isOnce && hasTodo(todos, title) {
			return todos, nil
		}

		todo := todo.NewTodo(order, parentID, title)
		return append(todos, todo), nil
	}
}

func hasTodo(todos []todo.Todo, title string) bool {
	for _, todo := range todos {
		if todo.Title == title {
			return true
		}
	}
	return false
}

func getNextOrder(parentID string) int {
	file := todo.OpenFile()
	todos, _ := file.Read()
	siblings := todo.FilterTodos(todos, func(todo todo.Todo) bool {
		return todo.ParentID == parentID
	})
	return len(siblings) + 1
}
