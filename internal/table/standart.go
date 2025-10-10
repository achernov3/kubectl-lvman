package table

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
)

type Standart struct{}

func NewStandart() ITable {
	return Standart{}
}

func (s Standart) RenderTable(data [][]string, header []string) {
	border := tw.BorderNone
	settings := tw.Settings{
		Separators: tw.SeparatorsNone,
		Lines:      tw.LinesNone,
	}

	table := tablewriter.NewTable(os.Stdout, tablewriter.WithRenderer(renderer.NewBlueprint(tw.Rendition{Borders: border, Settings: settings})))

	table.Header(header)
	table.Renderer()

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
