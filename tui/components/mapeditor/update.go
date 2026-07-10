package mapeditor

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)


func (c *Component) Open() {
	c.dialog.Open()
	c.LoadBindings()

	if c.onOpen != nil {
		c.onOpen()
	}
}

func (c *Component) Close() {
	c.dialog.Close()
	c.UnloadBindings()

	if c.onClose != nil {
		c.onClose()
	}
}

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(msg, chain.OnKey(c.HandleBindings))
}
