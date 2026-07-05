package entrysingle 

import (
	"charm.land/bubbles/v2/key"
)

type KeyMap struct {
	Up   key.Binding
	Down key.Binding
	Enter key.Binding
}

var Bindings = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("k/up", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("j/down", "down"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
}

func (s *Screen) GetBindings() []key.Binding {
	return []key.Binding{
		Bindings.Up,
		Bindings.Down,
	}
}
