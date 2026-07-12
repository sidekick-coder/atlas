package list

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

func (c *Component) HandleKeypress(msg tea.Msg) tea.Cmd {
	km, ok := msg.(tea.KeyMsg)

	if !ok {
		return nil
	}

	return chain.Keypress(km, c.HandleSelection)
}

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(msg, c.HandleKeypress)
}
