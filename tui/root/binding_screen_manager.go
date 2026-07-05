package root

import (
	"github.com/charmbracelet/bubbles/key"
)

type ScreenManagerKeyMap struct {
	Next key.Binding
	Prev key.Binding
}

var ScreenBindings = ScreenManagerKeyMap{
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
		ScreenBindings.Next,
		ScreenBindings.Prev,
	}
}
