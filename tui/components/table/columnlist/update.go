package columnlist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	handlers := []func(tea.Msg) tea.Cmd{}

	if c.IsOpen() {
		handlers = append(handlers, c.HandleBindings)
	}

	return chain.Update(msg, handlers...)
}
