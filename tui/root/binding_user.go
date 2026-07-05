package root

import (
	"charm.land/bubbles/v2/key"
)

func (m *model) GetUserBindings() []key.Binding {
	bindings := []key.Binding{}

	keymaps := m.app.Config().GetKeymapsByGroup("global")

	for _, action := range keymaps {
		bindings = append(bindings, key.NewBinding(
			key.WithKeys(action.Keys...),
			key.WithHelp(action.Keys[0], action.Description),
		))
	}

	return bindings
}
