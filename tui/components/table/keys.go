package table

import (
	"charm.land/bubbles/v2/key"
)

type Keymap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Next  key.Binding
	Prev  key.Binding
}

var Binding = Keymap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("k/up", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("j/down", "down"),
	),
	Next: key.NewBinding(
		key.WithKeys("l", "right"),
		key.WithHelp("l/right", "next"),
	),
	Prev: key.NewBinding(
		key.WithKeys("h", "left"),
		key.WithHelp("h/left", "prev"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
}

func (c *Component) GetBindigs() []key.Binding {
	bindings := []key.Binding{}

	bindings = append(bindings,
		Binding.Up,
		Binding.Down,
		Binding.Next,
		Binding.Prev,
		Binding.Enter,
	)

	return bindings
}
