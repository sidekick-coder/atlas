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

	op := m.app.Config().GetArrayString("tui.open_screens")

	for _, s := range op {
		_, err := m.screen.Add(s)

		if err != nil {
			return toast.Error(err.Error())
		}
	}

	index, ok := m.app.Config().GetInt("tui.initial_screen_index")

	if ok {
		err := m.screen.SetCurrent(index)

		if err != nil {
			return toast.Error(err.Error())
		}
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
		action.Update,

		m.LoadHome,

		m.toaster.Update,
		m.footer.Update,
		m.screen.Update,
		m.toolbar.Update,

		m.HandleMessages,
		chain.OnKey(m.HandleBinding),
		keymaps.Update,
	)

	return m, cmd
}
