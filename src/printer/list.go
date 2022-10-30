package printer

import (
	"io"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/mauricioabreu/golings/src/exercises"
)

func PrintList(o io.Writer, exs []exercises.Exercise) {
	t := table.NewWriter()
	t.SetOutputMirror(o)
	t.AppendHeader(table.Row{"Name", "Path"})
	for _, ex := range exs {
		t.AppendRow(table.Row{ex.Name, ex.Path})
	}
	t.Render()
}
