package input

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

func (i *Input) Update(msg tea.Msg) tea.Cmd {
	return chain.Update(msg, i.HandleKeypress)
}

