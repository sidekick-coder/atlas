package root

import (
	"github.com/charmbracelet/bubbles/key"
)

type GlobalKeyMap struct {
	Quit      key.Binding
}

var GlobalBindings = GlobalKeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

func (m *model) GetGlobalBindings() []key.Binding {
	return []key.Binding{
		GlobalBindings.Quit,
	}
}
