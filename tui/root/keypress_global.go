package root

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/bubbles/v2/key"
)

func (m *model) HandleGlobalKeyMap(msg tea.Msg) tea.Cmd {
	km, ok := msg.(tea.KeyMsg)

	if !ok {
		return nil
	}

	if key.Matches(km, GlobalBindings.Quit) {
		return tea.Quit
	}

	return nil
}
