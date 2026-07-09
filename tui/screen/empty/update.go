package empty

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/utils"
)

func (s *Screen) Update(msg tea.Msg) tea.Cmd {
	return  utils.Chain(msg, s.list.HandleKeypress)
}
