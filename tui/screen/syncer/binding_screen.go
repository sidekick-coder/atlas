package syncer

import (
	"charm.land/bubbles/v2/key"
)

type ScreenKeyMap struct {
	Execute key.Binding
	ExecuteWithList key.Binding
	Clear   key.Binding
}

var ScreenBindings = ScreenKeyMap{
	Execute: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "Execute"),
	),
	ExecuteWithList: key.NewBinding(
		key.WithKeys("E"),
		key.WithHelp("E", "Execute detailed"),
	),
	Clear: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "Clear"),
	),
}

func (s *Screen) GetScreenBindigs() []key.Binding {
	return []key.Binding{
		ScreenBindings.Execute,
		ScreenBindings.ExecuteWithList,
		ScreenBindings.Clear,
	}
}
