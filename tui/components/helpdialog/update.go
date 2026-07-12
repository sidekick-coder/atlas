package helpdialog

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(msg, chain.OnKey(c.HandleBindings), c.HandleView)
}
