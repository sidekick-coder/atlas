package syncer

import (
	"charm.land/bubbles/v2/key"
)

type ScreenKeyMap struct {
	Enter	   key.Binding
}

var ScreenBindings = ScreenKeyMap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Enter"),
	),
}

func (s *Screen) GetScreenBindigs() []key.Binding {
	return []key.Binding{
		ScreenBindings.Enter,
	}
}
