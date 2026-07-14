package sidepeeck

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/app/screen"
)

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	if _, ok := msg.(screen.SizeMsg); ok {
		c.LoadSize()
	}
	return nil
}
