package entry 

import (
	"github.com/charmbracelet/bubbles/key"
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
	bindings := []key.Binding{}

	bindings = append(bindings, Bindings.Up)
	bindings = append(bindings, Bindings.Down)
	bindings = append(bindings, Bindings.Enter)

	bindings = append(bindings, s.GetKeymapBindings()...)

	return bindings
}
