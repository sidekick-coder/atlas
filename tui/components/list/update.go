package list

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/utils"
)

func (c *Component) HandleKeypress(msg tea.Msg) tea.Cmd {
	km, ok := msg.(tea.KeyMsg)

	if !ok {
		return nil
	}

	return utils.ChainKeypress(km, c.HandleSelection)
}

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return utils.Chain(msg, c.HandleKeypress)
}
