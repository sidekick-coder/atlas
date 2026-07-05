package root

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/messages"
)

func (m *model) HandleInput(msg tea.Msg) tea.Cmd {
	if im, ok := msg.(messages.Input); ok {
		m.input.SetTitle(im.Title)
		m.input.Open("")
		return nil
	}

	return m.input.HandleKeypress(msg)
}
