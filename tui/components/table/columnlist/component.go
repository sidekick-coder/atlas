package columnlist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/sidepeeck"
	"github.com/sidekick-coder/atlas/tui/components/table/column"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

type Component struct {
	sidepeeck sidepeeck.Component
	column    *column.Feature
}

func Create() *Component {
	return &Component{
		sidepeeck: *sidepeeck.Create(),
	}
}

func (c *Component) Open() {
	c.sidepeeck.Open()
	c.sidepeeck.OnRender(c.Render)
}

func (c *Component) Close() {
	c.sidepeeck.Close()
}

func (c *Component) IsOpen() bool {
	return c.sidepeeck.IsOpen()
}

func (c *Component) SetColumn(column *column.Feature) {
	c.column = column
}

func (c *Component) Init() tea.Cmd {
	return chain.Init(c.sidepeeck.Init)
}

func (c *Component) Dispose() tea.Cmd {
	return chain.Dispose(c.sidepeeck.Dispose)
}
