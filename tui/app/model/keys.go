package model

import (
	tea "charm.land/bubbletea/v2"
	"github.com/sidekick-coder/atlas/tui/features/key"
)

type Keymap struct {
	Close key.Binding
}

var tags = []string{"global"}

var Bindings = Keymap{
	Close: key.CreateBinding("q", "<c-c>").
		SetDescription("Close the current screen").
		SetTags(tags...).
		SetHelp("q/crtl+c"),
}

func (m *model) GetBindings() []key.Binding {
	bindings := []key.Binding{}

	bindings = append(bindings, Bindings.Close)

	bindings = append(bindings, m.GetUserBindings()...)

	return bindings
}

func (m *model) LoadBindings() tea.Cmd {
	key.Register(m.GetBindings()...)

	return nil
}

func (m *model) UnloadBindings() {
	key.Unregister(m.GetBindings()...)
}

func (m *model) GetUserBindings() []key.Binding {
	bindings := []key.Binding{}

	keymaps := m.app.Config().GetKeymapsByGroup("global")

	for _, action := range keymaps {
		b := key.CreateBinding(action.Keys...).
			SetDescription(action.Description).
			SetTags(tags...).
			SetHelp(action.Keys[0])

		bindings = append(bindings, b)
	}

	return bindings
}

func (m *model) HandleBinding(km tea.KeyMsg) tea.Cmd {
	if key.Matches(Bindings.Close) {
		return tea.Quit
	}

	return nil
}
