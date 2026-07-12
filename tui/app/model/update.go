package model

import (
	"log/slog"

	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

func (m *model) LoadHome(msg tea.Msg) tea.Cmd {
	_, ok := msg.(tea.WindowSizeMsg)

	if !ok {
		return nil
	}

	if m.ready {
		return nil
	}

	hs, ok := m.app.Config().Get("tui.home_screen")

	if !ok {
		return m.AddScreenEmpty()
	}

	_, err := m.screen.Add(hs)

	if err != nil {
		return toast.Error(err.Error())
	}

	m.ready = true

	return nil
}

func (m *model) HandleMessages(msg tea.Msg) tea.Cmd {
	if wm, ok := msg.(tea.WindowSizeMsg); ok {
		slog.Info("window size changed", slog.Int("width", wm.Width), slog.Int("height", wm.Height))
		m.SetSize(wm.Width, wm.Height)
	}

	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := chain.Update(msg,
		key.HandleKeypress,

		m.HandleMessages,
		m.LoadHome,
		m.HandleActions,
		m.actionBindingMessageHandler,

		m.toaster.Update,
		m.footer.Update,
		m.screen.Update,
		m.toolbar.Update,

	    chain.OnKey(m.HandleBinding),
	)

	return m, cmd
}
