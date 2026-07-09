package root

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	handlers := []func(tea.Msg) tea.Cmd{}

	handlers = append(handlers,
		m.HandleActions,
		m.actionBindingMessageHandler,
		m.HandleInput,
		m.HandleGlobalKeyMap,

		m.HandleScreeManagerKeypress,
		m.HandleScreeManagerMessages,

		m.HandleToastMessages,
	)

	for _, handler := range handlers {
		cmd := handler(msg)

		if cmd != nil {
			return m, cmd
		}
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
		return m, m.AddScreenEmpty()
	case components.ToastExpiredMsg:
		components.GlobalToast.Hide()
		return m, nil

	case components.ToastShowMsg:
		cmd := components.GlobalToast.Show(msg.Message, msg.Level, 2*time.Second)
		return m, cmd
	}

	if len(m.screens) == 0 {
		return m, nil
	}

	cmd := m.screens[m.currentIndex].Update(msg)

	return m, cmd
}
