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
		SetDescription("close").
		SetTags(tags...).
		SetHelp("q"),
}

var UserBindings = []key.Binding{}

func (m *model) GetBindings() []key.Binding {
	return []key.Binding{
		Bindings.Close,
	}
}

func (m *model) LoadBindings() tea.Cmd {
	key.Register(m.GetBindings()...)

	return nil
}

func (m *model) UnloadBindings() {
	key.Unregister(m.GetBindings()...)

	UserBindings = []key.Binding{}
}

func (m *model) HandleBinding(km tea.KeyMsg) tea.Cmd {
	if key.Matches(Bindings.Close) {
		return tea.Quit
	}

	return nil
}
