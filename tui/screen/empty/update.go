package empty

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
)

func (s *Screen) Update(msg tea.Msg) tea.Cmd {
	return  chain.Update(msg, s.list.HandleKeypress)
}
