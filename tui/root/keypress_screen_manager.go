package root

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/key"
)

func (m *model) HandleScreeManagerKeypress(msg tea.Msg) tea.Cmd {
	km, ok := msg.(tea.KeyMsg)

	if !ok {
		return nil
	}

	if key.Matches(km, ScreenBindings.Next) {
		nextIndex := (m.currentIndex + 1) % len(m.screens)
		m.currentIndex = nextIndex
		m.tabBar.SetCurrent(nextIndex)
		return nil
	}

	if key.Matches(km, ScreenBindings.Prev) {
		prevIndex := (m.currentIndex - 1 + len(m.screens)) % len(m.screens)
		m.currentIndex = prevIndex
		m.tabBar.SetCurrent(prevIndex)

		return nil
	}

	if key.Matches(km, ScreenBindings.Add) {
		m.AddScreen("empty", nil)

		return nil
	}

	return nil
}

