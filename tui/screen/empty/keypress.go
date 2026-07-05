package empty

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/messages"
)

func (s *Screen) HandleKeyPress(msg tea.Msg) tea.Cmd {
	km, ok := msg.(tea.KeyMsg)

	if !ok {
		return nil
	}

	if key.Matches(km, ScreenBindings.EntryList) {
		return messages.ReplaceScreenCmd(messages.ReplaceCurrentScreen{
			Name:    "entry_list",
			Options: nil,
		})
	}

	return nil
}
