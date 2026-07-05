package root

import (
	"github.com/charmbracelet/bubbles/key"
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
