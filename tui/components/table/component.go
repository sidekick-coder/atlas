package table

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/table/column"
	"github.com/sidekick-coder/atlas/tui/components/table/columnlist"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/selection"
)

type Column = column.Column

type Item struct {
	Values map[string]string
}

type Component struct {
	width    int
	height   int
	cursor int
	onSelect func(cursor int) tea.Cmd
	items    []Item

	column *column.Feature
	columnList columnlist.Component

	itemSelection selection.Feature
}

func Create() *Component {
	return &Component{
		width:    100,
		height:   100,
		cursor: 0,
		items:    []Item{},

		column: column.Create(),
		columnList: *columnlist.Create(),

		itemSelection: *selection.Create(),
	}
}

func (c *Component) OnSelect(f func(cursor int) tea.Cmd) *Component {
	c.onSelect = f
	return c
}

func (c *Component) SetColumns(columns []*column.Column) {
	c.column.SetColumns(columns) 
}

func (c *Component) SetItems(items []Item) {
	c.items = items
	c.itemSelection.SetTotal(len(items))
}

func (c *Component) Init() tea.Cmd {
	c.columnList.SetColumn(c.column)

	c.columnList.OnOpen(func() {
		c.UnloadBindings()
	})

	c.columnList.OnClose(func() {
		c.LoadBindings()
	})

	c.LoadBindings()
	c.columnList.Open()

	return chain.Init(c.columnList.Init)
}

func (c *Component) Dispose() tea.Cmd {
	return chain.Dispose(c.columnList.Dispose)
}

