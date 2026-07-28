package entrymeta

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/entrycontroller"
)

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(
		msg,
		c.dialog.Update,
		c.keyValue.Update,
		chain.OnKey(c.HandleBindings),
		c.HandleEntryUpdate,
	)
}

func (c *Component) HandleEntryUpdate(msg tea.Msg) tea.Cmd {
	if m, ok := msg.(entrycontroller.UpdatedMsg); ok && m.Path == c.path {
		index := c.selection.GetCursor()

		c.Load()

		if index >= 0 && index < len(c.metas) {
			c.selection.SetCursor(index)
		}
	}

	return nil
}
