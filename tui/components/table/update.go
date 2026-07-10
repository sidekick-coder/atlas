package table

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)


func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(msg, c.columnList.Update, chain.OnKey(c.HandleBindings))
}
