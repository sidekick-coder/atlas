package helpdialog

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/dialog"
	"github.com/sidekick-coder/atlas/tui/components/viewport"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

type Component struct {
	dialog   *dialog.Component
	viewport *viewport.Component
}

func Create() *Component {
	c := &Component{
		dialog:   dialog.Create().SetTitle("Help"),
		viewport: viewport.Create(),
	}

	return c
}

func (c *Component) Open() {
	c.dialog.Open()
	c.LoadBindings()
}

func (c *Component) Close() {
	c.dialog.Close()
	c.UnloadBindings()
}

func (c *Component) IsOpen() bool {
	return c.dialog.IsOpen()
}


func (c *Component) Init() tea.Cmd {
	c.dialog.OnRender(c.Render)
	return chain.Init(c.dialog.Init)
}
