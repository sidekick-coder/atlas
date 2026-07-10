package root

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

func (m *model) HandleWindow(msg tea.Msg) tea.Cmd {
	wm, ok := msg.(tea.WindowSizeMsg)

	if !ok {
		return nil
	}

	m.SetSize(wm.Width, wm.Height)

	hs, ok := m.app.Config().Get("tui.home_screen")

	if ok {
		return m.AddScreen(hs)
	}

	return m.AddScreenEmpty()
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
		m.HandleGlobalKeyMap,

		m.HandleScreenUpdate,
		m.HandleScreeManagerKeypress,
		m.HandleScreeManagerMessages,

		m.HandleToastMessages,
	)

	return m, cmd
}
