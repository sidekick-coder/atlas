package syncer

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

func (s *Screen) HandleScreenKeymaps(msg tea.Msg) tea.Cmd {
	km, ok := msg.(tea.KeyMsg)

	if !ok {
		return nil
	}

	if key.Matches(km, ScreenBindings.Execute) {
		s.ViewList = false
		s.Sync()
		return nil
	}
	if key.Matches(km, ScreenBindings.ExecuteWithList) {
		s.ViewList = true
		s.Sync()
		return nil
	}

	if key.Matches(km, ScreenBindings.Clear) {
		return func() tea.Msg {
			return EntryClear{}
		}
	}

	return nil
}
