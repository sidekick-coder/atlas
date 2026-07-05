package root

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/messages"
)

func (m *model) HandleScreeManagerMessages(msg tea.Msg) tea.Cmd {
	if as, ok := msg.(messages.AddScreen); ok {
		return m.AddScreen(as.Name, as.Options)
	}

	if rcs, ok := msg.(messages.ReplaceCurrentScreen); ok {
		return m.ReplaceScreen(m.currentIndex, rcs.Name, rcs.Options)
	}

	return nil
}
