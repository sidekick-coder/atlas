package empty 

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	EntryList key.Binding
}

var Bindings = KeyMap{
	EntryList: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "entry list"),
	),
}

func (s *Screen) GetBindings() []key.Binding {
	return []key.Binding{
		Bindings.EntryList,
	}
}
