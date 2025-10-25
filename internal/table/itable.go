package table

type ITable interface {
	RenderTable(data [][]string, header []string) error
}
