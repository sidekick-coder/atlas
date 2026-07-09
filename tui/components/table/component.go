package table

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/selection"
)

type Column struct {
	Label string 
	Field string
	Width int
}

type Component struct {
	width    int
	height   int
	cursor int
	onSelect func(cursor int) tea.Cmd
	items    []string

	columns  []Column
	columnSelection selection.Feature
}

func Create() *Component {
	return &Component{
		width:    100,
		height:   100,
		cursor: 0,
		items:    []string{},

		columns:  []Column{},
		columnSelection: *selection.Create(),
	}
}

func (c *Component) OnSelect(f func(cursor int) tea.Cmd) *Component {
	c.onSelect = f
	return c
}

func (c *Component) SetColumns(columns []Column) {
	c.columns = columns
}

func (c *Component) Init() tea.Cmd {
	c.columnSelection.SetTotal(len(c.columns))
	return chain.Init(c.LoadBindings)
}

