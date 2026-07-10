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

	// if key.Matches(km, ScreenBindings.Next) {
	// 	nextIndex := (m.currentIndex + 1) % len(m.screens)
	// 	return m.SetCurrentScreen(nextIndex)
	// }
	//
	// if key.Matches(km, ScreenBindings.Prev) {
	// 	prevIndex := (m.currentIndex - 1 + len(m.screens)) % len(m.screens)
	// 	return m.SetCurrentScreen(prevIndex)
	// }

	if key.Matches(km, ScreenBindings.Add) {
		return m.AddScreenEmpty()
	}

	if key.Matches(km, ScreenBindings.Close) {
		if len(m.screens) > 0 {
			m.RemoveScreen(m.currentIndex)
		}

		return nil
	}

	return nil
}

