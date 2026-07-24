package entrymeta

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/entrycontroller"
)

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(msg, c.dialog.Update, chain.OnKey(c.HandleBindings), c.HandleUpdate)
}

func (c *Component) HandleUpdate(msg tea.Msg) tea.Cmd {
	if m, ok := msg.(entrycontroller.UpdatedMsg); ok && m.Path == c.path {
		c.Load()
	}

	return  nil
}
