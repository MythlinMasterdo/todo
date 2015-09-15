package formatter

import (
	"fmt"
	"io"

	"github.com/naoty/todo/todo"
)

type NumberedFormatter struct {
	writer io.Writer
	Mode   Mode
}

func NewNumberedFormatter(w io.Writer, m Mode) *NumberedFormatter {
	return &NumberedFormatter{writer: w, Mode: m}
}

func (f *NumberedFormatter) Print(todos []todo.Todo, indent string) error {
	file := todo.OpenFile()

	for i, todo := range todos {
		if f.Mode == DONE && !todo.Done {
			continue
		}
		if f.Mode == UNDONE && todo.Done {
			continue
		}
		mark := NewMark(todo)
		fmt.Fprintf(f.writer, "%s%s %03d: %s\n", indent, mark, i+1, todo.Title)

		subTodos, err := file.ReadSubTodos(todo.ID)
		if err != nil {
			continue
		}

		f.Print(subTodos, indent+"  ")
	}

	return nil
}
