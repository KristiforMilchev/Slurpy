package implementations

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type TablePrinter struct {
}

func (tablePrinter *TablePrinter) Print(rows []table.Row, columns table.Row) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(columns)

	for _, row := range rows {
		t.AppendRow(row)
	}

	t.Render()
}
