package entrylist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/entrycontroller"
)

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(
		msg,
		c.dialog.Update,
		c.list.Update,
		c.HandleMessage,
		chain.OnKey(c.HadleBinding),
	)
}

func (c *Component) HandleMessage(msg tea.Msg) tea.Cmd {
	if _, ok := msg.(entrycontroller.UpdatedMsg); ok {
		c.loader.Load()
	}

	return nil
}
