package root

import (
	"charm.land/bubbles/v2/key"
)

type ScreenManagerKeyMap struct {
	Add  key.Binding
	Next key.Binding
	Prev key.Binding
}

var ScreenBindings = ScreenManagerKeyMap{
	Add: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "add screen"),
	),
	Next: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next screen"),
	),
	Prev: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("shift+tab", "previous screen"),
	),
}

func (m *model) GetScreenManagerBindings() []key.Binding {
	return []key.Binding{
		ScreenBindings.Add,
		ScreenBindings.Next,
		ScreenBindings.Prev,
	}
}
