package model

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/action"
	"github.com/sidekick-coder/atlas/tui/components/toast"
	"github.com/sidekick-coder/atlas/tui/features/chain"
	"github.com/sidekick-coder/atlas/tui/features/key"
	"github.com/sidekick-coder/atlas/tui/features/keymaps"
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
		m.SetSize(wm.Width, wm.Height)
	}

	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := chain.Update(msg,
		key.HandleKeypress,

		m.LoadHome,

		m.toaster.Update,
		m.footer.Update,
		m.screen.Update,
		m.toolbar.Update,

		m.HandleMessages,
	    chain.OnKey(m.HandleBinding),
		keymaps.Update,
		action.Update,
	)

	return m, cmd
}
