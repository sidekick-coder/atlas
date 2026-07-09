package empty 

import (
	"charm.land/bubbles/v2/key"
)

type ScreenKeyMap struct {
	EntryList key.Binding
}

var ScreenBindings = ScreenKeyMap{
	EntryList: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "entry list"),
	),
}

func (s *Screen) GetBindings() []key.Binding {
	return []key.Binding{
		ScreenBindings.EntryList,
	}
}
