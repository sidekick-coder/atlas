package entrylist

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/app/program"
	"github.com/sidekick-coder/atlas/tui/components/list"
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

	if _, ok := msg.(list.MovedMsg); ok {
		current, exists := c.GetCurrent()

		return program.Command(ChangedMsg{
			Entry:  current,
			Exists: exists,
		})
	}

	return nil
}
