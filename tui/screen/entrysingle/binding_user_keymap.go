package entrysingle

import (
	"charm.land/bubbles/v2/key"
)

func (s *Screen) GetUserKeymapBindings() []key.Binding {
	keymaps := s.App.Config().GetKeymapsByGroup("entry_single")

	bindings := []key.Binding{}

	for _, km := range keymaps {
		b := key.NewBinding(
			key.WithKeys(km.Keys...),
			key.WithHelp(km.Keys[0], km.Description),
		)

		bindings = append(bindings, b)
	}

	return bindings
}

