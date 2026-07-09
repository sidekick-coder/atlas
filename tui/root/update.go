package root

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

func (m *model) HandleWindow(msg tea.Msg) tea.Cmd {
	if wm, ok := msg.(tea.WindowSizeMsg); ok {
		m.SetSize(wm.Width, wm.Height)
		return m.AddScreenEmpty()
	}

	return nil
}

func (m *model) HandleScreenUpdate(msg tea.Msg) tea.Cmd {
	if len(m.screens) == 0 {
		return nil
	}

	s := m.screens[m.currentIndex]

	if s == nil {
		return nil
	}

	return s.Update(msg)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := chain.Update(msg,
		key.HandleKeypress,
		m.HandleWindow,
		m.HandleActions,
		m.actionBindingMessageHandler,
		m.HandleInput,
		m.HandleGlobalKeyMap,

		m.HandleScreeManagerKeypress,
		m.HandleScreeManagerMessages,

		m.HandleToastMessages,
		m.HandleScreenUpdate,
	)

	return m, cmd
}
