package table

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/utils"
)


func (c *Component) Update(msg tea.Msg) tea.Cmd {
	return utils.Chain(msg, c.columnList.Update, c.HandleBindings)
}
