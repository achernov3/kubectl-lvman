package table

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
)

type Standart struct{}

func NewStandart() ITable {
	return Standart{}
}

func (s Standart) RenderTable(data [][]string, header []string) error {
	border := tw.BorderNone
	settings := tw.Settings{
		Separators: tw.SeparatorsNone,
		Lines:      tw.LinesNone,
	}

	table := tablewriter.NewTable(os.Stdout, tablewriter.WithRenderer(renderer.NewBlueprint(tw.Rendition{Borders: border, Settings: settings})))

	table.Header(header)
	table.Renderer()

	for _, v := range data {
		err := table.Append(v)
		if err != nil {
			return fmt.Errorf("failed to append data: %w")
		}
	}

	err := table.Render()
	if err != nil {
		return fmt.Errorf("failed to render table: %w")
	}

	return nil
}
