package root

import (
	"charm.land/bubbles/v2/key"
)

type GlobalKeyMap struct {
	Quit      key.Binding
	OpenSyncer key.Binding
}

var GlobalBindings = GlobalKeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	OpenSyncer: key.NewBinding(
		key.WithKeys("S"),
		key.WithHelp("S", "open syncer"),
	),
}

func (m *model) GetGlobalBindings() []key.Binding {
	return []key.Binding{
		GlobalBindings.Quit,
		GlobalBindings.OpenSyncer,
	}
}
