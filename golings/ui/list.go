package ui

import (
	"io"

	"github.com/cleonte/golings/golings/exercises"
	"github.com/jedib0t/go-pretty/v6/table"
)

func PrintList(o io.Writer, exs []exercises.Exercise) {
	t := table.NewWriter()
	t.SetOutputMirror(o)
	t.AppendHeader(table.Row{"Name", "Path", "State"})
	for _, ex := range exs {
		t.AppendRow(table.Row{ex.Name, ex.Path, ex.State()})
	}
	t.Render()
}
