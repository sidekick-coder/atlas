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

type Item struct {
	Values map[string]string
}

type Component struct {
	width    int
	height   int
	cursor int
	onSelect func(cursor int) tea.Cmd
	items    []Item

	columns  []Column
	columnSelection selection.Feature
	itemSelection selection.Feature
}

func Create() *Component {
	return &Component{
		width:    100,
		height:   100,
		cursor: 0,
		items:    []Item{},

		columns:  []Column{},
		columnSelection: *selection.Create(),
		itemSelection: *selection.Create(),
	}
}

func (c *Component) OnSelect(f func(cursor int) tea.Cmd) *Component {
	c.onSelect = f
	return c
}

func (c *Component) SetColumns(columns []Column) {
	c.columns = columns
	c.columnSelection.SetTotal(len(c.columns))
}

func (c *Component) SetItems(items []Item) {
	c.items = items
	c.itemSelection.SetTotal(len(items))
}

func (c *Component) Init() tea.Cmd {
	return chain.Init(c.LoadBindings)
}

